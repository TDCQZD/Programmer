package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/boltdb/bolt"
)

const dbFile = "blockchain.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"

/*Blockchain 表示区块链，保存区块*/
type Blockchain struct {
	tip []byte
	db  *bolt.DB
}

/*Iterator 返回BlockchainIterator*/
func (bc *Blockchain) Iterator() *BlockchainIterator {
	return &BlockchainIterator{bc.tip, bc.db}
}

/*CreateBlockchain 创建一个新的区块链数据库*/
func CreateBlockchain(address string) *Blockchain {
	if dbExists() {
		fmt.Println("Blockchain already exists.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {

		cbTX := NewCoinbaseTX(address, genesisCoinbaseData)
		genesisBlock := NewGenesisBlock(cbTX)
		b, err := tx.CreateBucket([]byte(blocksBucket))
		if err != nil {
			log.Panic(err)
		}
		err = b.Put(genesisBlock.Hash, genesisBlock.SerializeBlock())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"), genesisBlock.Hash)
		if err != nil {
			log.Panic(err)
		}
		tip = genesisBlock.Hash

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc := &Blockchain{tip, db}
	return bc

}

/*NewBlockchain 创建并返回一个带创世区块是区块链*/
func NewBlockchain() *Blockchain {
	if dbExists() == false {
		fmt.Println("No existing blockchain found. Create one first.")
		os.Exit(1)
	}

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Panic(err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		tip = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	bc := &Blockchain{tip, db}
	return bc
}

/*MineBlock 给区块链中添加新区块*/
func (bc *Blockchain) MineBlock(transactions []*Transaction) {
	var lastBlockHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastBlockHash = b.Get([]byte("l"))
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
	newBlock := NewBlock(transactions, lastBlockHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.SerializeBlock())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		if err != nil {
			log.Panic(err)
		}

		bc.tip = newBlock.Hash
		return nil
	})
	if err != nil {
		log.Panic(err)
	}
}

/*FindTransaction 按ID查找交易*/
func (bc *Blockchain) FindTransaction(ID []byte) (Transaction, error) {
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, tx := range block.Transactions {
			if bytes.Compare(tx.ID, ID) == 0 {
				return *tx, nil
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}

	return Transaction{}, errors.New("Transaction is not found")
}

/*FindUnspentTransactions 返回包含未使用输出的交易列表*/
func (bc *Blockchain) FindUnspentTransactions(pubKeyHash []byte) []Transaction {
	var unspentTxs []Transaction
	spentTXOS := make(map[string][]int)
	bci := bc.Iterator()

	for {
		block := bci.Next()

		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {
				if spentTXOS[txID] != nil {
					for _, spentOut := range spentTXOS[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}

				if out.IsLockedWithKey(pubKeyHash) {
					unspentTxs = append(unspentTxs, *tx)

				}
			}
			if tx.IsCoinbaseTX() == false {
				for _, in := range tx.Vin {
					if in.UsesKey(pubKeyHash) {
						inTxID := hex.EncodeToString(in.Txid)
						spentTXOS[inTxID] = append(spentTXOS[inTxID], in.Vout)
					}
				}
			}
		}

		if len(block.PrevBlockHash) == 0 {
			break
		}

	}
	return unspentTxs
}

/*FindUTXO 查找并返回所有未使用的交易输出 */
func (bc *Blockchain) FindUTXO(pubKeyHash []byte) []TXOutput {
	var UTXOs []TXOutput
	unspentTxs := bc.FindUnspentTransactions(pubKeyHash)
	for _, tx := range unspentTxs {
		for _, out := range tx.Vout {
			if out.IsLockedWithKey(pubKeyHash) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

/*FindSpendableOutputs 查找并返回未使用的输出以在输入中引用*/
func (bc *Blockchain) FindSpendableOutputs(pubKeyHash []byte, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	unspentTXs := bc.FindUnspentTransactions(pubKeyHash)
	accumulated := 0

Work:
	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Vout {
			if out.IsLockedWithKey(pubKeyHash) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)

				if accumulated > amount {
					break Work
				}
			}
		}
	}
	return accumulated, unspentOutputs

}

/*SignTransaction 签署交易的输入*/
func (bc *Blockchain) SignTransaction(tx *Transaction, privKey ecdsa.PrivateKey) {

	prevTXs := make(map[string]Transaction)

	for _, vin := range tx.Vin {
		prevTX, err := bc.FindTransaction(vin.Txid)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}

	tx.Sign(privKey, prevTXs)
}

/*VerifyTransaction 验证交易输入签名*/
func (bc *Blockchain) VerifyTransaction(tx *Transaction) bool {

	prevTXs := make(map[string]Transaction)

	for _, vin := range tx.Vin {
		prevTX, err := bc.FindTransaction(vin.Txid)
		if err != nil {
			log.Panic(err)
		}
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}

	return tx.Verify(prevTXs)

}
func dbExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}
