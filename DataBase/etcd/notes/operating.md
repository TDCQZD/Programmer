# etcd操作
用户主要通过输入或获取密钥的值与etcd进行交互。使用etcdctl（一个用于与etcd服务器交互的命令行工具）来完成此操作。
> 这里描述的概念应该适用于gRPC API或客户端库API。

通过etcdctl 配合help参数查看，可以使用的操作:
```
$ etcdctl help
NAME:
	etcdctl - A simple command line client for etcd3.

USAGE:
	etcdctl [flags]

VERSION:
	3.3.0+git

API VERSION:
	3.3


COMMANDS:
	alarm disarm		Disarms all alarms
	alarm list		Lists all alarms
	auth disable		Disables authentication
	auth enable		Enables authentication
	check datascale		Check the memory usage of holding data for different workloads on a given server endpoint.
	check perf		Check the performance of the etcd cluster
	compaction		Compacts the event history in etcd
	defrag			Defragments the storage of the etcd members with given endpoints
	del			Removes the specified key or range of keys [key, range_end)
	elect			Observes and participates in leader election
	endpoint hashkv		Prints the KV history hash for each endpoint in --endpoints
	endpoint health		Checks the healthiness of endpoints specified in `--endpoints` flag
	endpoint status		Prints out the status of endpoints specified in `--endpoints` flag
	get			Gets the key or a range of keys
	help			Help about any command
	lease grant		Creates leases
	lease keep-alive	Keeps leases alive (renew)
	lease list		List all active leases
	lease revoke		Revokes leases
	lease timetolive	Get lease information
	lock			Acquires a named lock
	make-mirror		Makes a mirror at the destination etcd cluster
	member add		Adds a member into the cluster
	member list		Lists all members in the cluster
	member remove		Removes a member from the cluster
	member update		Updates a member in the cluster
	migrate			Migrates keys in a v2 store to a mvcc store
	move-leader		Transfers leadership to another etcd cluster member.
	put			Puts the given key into the store
	role add		Adds a new role
	role delete		Deletes a role
	role get		Gets detailed information of a role
	role grant-permission	Grants a key to a role
	role list		Lists all roles
	role revoke-permission	Revokes a key from a role
	snapshot restore	Restores an etcd member snapshot to an etcd directory
	snapshot save		Stores an etcd node backend snapshot to a given file
	snapshot status		Gets backend snapshot status of a given file
	txn			Txn processes all the requests in one transaction
	user add		Adds a new user
	user delete		Deletes a user
	user get		Gets detailed information of a user
	user grant-role		Grants a role to a user
	user list		Lists all users
	user passwd		Changes password of user
	user revoke-role	Revokes a role from a user
	version			Prints the version of etcdctl
	watch			Watches events stream on keys or prefixes

OPTIONS:
      --cacert=""				verify certificates of TLS-enabled secure servers using this CA bundle
      --cert=""					identify secure client using this TLS certificate file
      --command-timeout=5s			timeout for short running command (excluding dial timeout)
      --debug[=false]				enable client-side debug logging
      --dial-timeout=2s				dial timeout for client connections
  -d, --discovery-srv=""			domain name to query for SRV records describing cluster endpoints
      --discovery-srv-name=""			service name to query when using DNS discovery
      --endpoints=[127.0.0.1:2379]		gRPC endpoints
  -h, --help[=false]				help for etcdctl
      --hex[=false]				print byte strings as hex encoded strings
      --insecure-discovery[=true]		accept insecure SRV records describing cluster endpoints
      --insecure-skip-tls-verify[=false]	skip server certificate verification
      --insecure-transport[=true]		disable transport security for client connections
      --keepalive-time=2s			keepalive time for client connections
      --keepalive-timeout=6s			keepalive timeout for client connections
      --key=""					identify secure client using this TLS key file
      --password=""				password for authentication (if this option is used, --user option shouldn't include password)
      --user=""					username[:password] for authentication (prompt if password is not supplied)
  -w, --write-out="simple"			set the output format (fields, json, protobuf, simple, table)


```
## 查找版本
etcdctl版本和Server API版本可以用于查找用于在etcd上执行各种操作的适当命令。

查找版本的命令：
```
etcdctl --version
```
## 存储数据-put
etcd 的数据按照树形的结构组织，类似于 linux 的文件系统，也有目录和文件的区别，不过一般被称为 nodes。数据的 endpoint 都是以 /v3/keys 开头（v3 表示当前 API 的版本），比如 /v3/keys/names/cizixs。

要创建一个值，只要使用 PUT 方法在对应的 url endpoint 设置就行。如果对应的 key 已经存在， PUT 也会对 key 进行更新。

etcd中的数据，都是通过key-value进行存储，所以存储数据也理解为是存储一个秘钥。应用程序通过写入密钥将密钥存储到etcd集群中。每个存储的密钥都通过Raft协议复制到所有etcd集群成员，以实现一致性和可靠性。

存储一个key-value：
```
$ etcdctl put foo hanru
```
> 返回OK就是存储成功：
## 查询数据-get
```
$ etcdctl get foo
```
### 按照16进制读取
```
$ etcdctl get foo --hex
```
### 只查询value
```
$ etcdctl get foo --print-value-only
```
### 查询指定范围的key
```
$ etcdctl get foo foo3
```
> 命令范围从键foo到foo3
### 按照前缀查询
```
$ etcdctl get --prefix --limit=2 foo
```
> 命令范围的所有键前缀foo，将结果数量限制为2
### 大于等于
```
$ etcdctl get --from-key b
```
> 读取大于或等于键的字节值的键的命令b
### json格式返回查询数据
```
$ etcdctl --write-out="json" get foo
```
###  阅读过去版本的key
```
$ etcdctl get --write-out="json" --prefix key
$ etcdctl get --prefix key    # access the most recent versions of keys

# access the versions of keys at revision 4
$ etcdctl get --prefix --rev=4 key      

# access the versions of keys at revision 3
$ etcdctl get --prefix --rev=3 key

# access the versions of keys at revision 2
$ etcdctl get --prefix --rev=2 key

# access the versions of keys at revision 1
$ etcdctl get --prefix --rev=1 key
```
## 删除数据-del
应用程序可以从一个etcd集群中删除一个密钥或一系列密钥。

### 删除指定key-value
```
etcdctl del foo
```
删除密钥的命令foo，返回值是成功删除key-value的个数：

### 删除指定范围的key

```
etcdctl del foo foo3
```
这里是删除键，从命令foo到foo3：
###  删除指定key返回键值对信息
删除键值zoo1与删除的键值对的命令:
```
etcdctl del --prev-kv zoo1
```

###  按照前缀删除
删除具有前缀的键的命令zoo：
```
$ etcdctl del --prefix zoo
```
###  删除大于等于指定key
删除大于或等于键的字节值的键的命令b：
```
$ etcdctl del --from-key b
```
## 事务写入-txn
txn从标准输入中读取多个请求，将它们看做一个原子性的事务执行。事务是由条件列表，条件判断成功时的执行列表（条件列表中全部条件为真表示成功）和条件判断失败时的执行列表（条件列表中有一个为假即为失败）组成的。
```
$ etcdctl put flag 1
$ etcdctl txn -i 

compares:
# 输入以下内容，输入结束后，按两次回车
value("flag") = "1"

# 如果flag = 1，则执行put result true，当然你也可以执行get flag等操作。。
success requests (get, put, del):
put result true

# 如果flag != 1，则执行put result flase
failure requests (get, put, del):
put result false

# 运行结果，执行 success
SUCCESS

OK

# 最后我们可以get一下result的值
$ etcdctl get result
result
true
```
- etcdctl put flag 1设置flag为1
- etcdctl txn -i开启事务（-i表示交互模式：也可以使用 --interactive）
- 第2步输入命令后回车，终端显示出compares：
- 输入value("flag") = "1"，此命令是比较flag的值与1是否相等
- 第4步完成后输入回车，终端会换行显示，此时可以继续输入判断条件（前面说过事务由条件列表组成），再次输入回车表示判断条件输入完毕
- 第5步连续输入两个回车后，终端显示出success requests (get, put, delete):，表示下面输入判断条件为真时要执行的命令
- 与输入判断条件相同，连续两个回车表示成功时的执行列表输入完成
- 终端显示failure requests (get, put, delete):后输入条件判断失败时的执行列表
为了看起来简洁，此实例中条件列表和执行列表只写了一行命令，实际可以输入多行

总结上面的事务，要做的事情就是flag为1时设置result为true，否则设置result为false

事务执行完成后查看result值为true
## 查看键更改watch
应用程序可以观察一个键或一系列键来监视任何更新。
```
etcdctl watch foo
```
> 当 foo 的数值改变（ put 方法）的时候，watch 会收到通知
## 租约lease
### 赠与租约
应用程序可以为etcd集群中的密钥授予租约。当一个密钥附加到租约中时，其生命周期必然与租期的生命周期相关，后者又由生存时间（TTL）管理。每个租约都有一个在授予时间由应用程序指定的最小生存时间（TTL）值。租约的实际TTL值至少是最小TTL，由etcd集群选择。租约的TTL过期后，租约到期并删除所有连接的密钥。

以下是授予租约的命令：
```
# grant a lease with 300 second TTL
$ etcdctl lease grant 300
lease 2deb68788ab64ebd granted with TTL(300s)
> 注意这个租约的时间是自己设置，秒为单位，时间到了自动就销毁了。所以300s后，这个租约就不在了。。
# 查看lease list
$ etcdctl lease list
found 1 leases
2deb68788ab64ebd

# attach key foo to lease 2deb68788ab64ebd
$ etcdctl put name hanru --lease=2deb68788ab64ebd
OK
```
### 撤销租约
应用程序通过租赁ID撤销租赁。撤销租约将删除其所有连接的密钥。
在合约有效时长内手动撤销：
```
$  etcdctl lease revoke 2deb68788ab64ebd
lease 2deb68788ab64ebd revoked

$  etcdctl lease list
found 0 leases
$  etcdctl get name

# empty response since foo is deleted due to lease revocation
```
撤销租约后，key-value也随之没有了。
### 保持租约活着
应用程序可以通过刷新其TTL来保持租约活着，因此不会过期。

假设我们完成了以下一系列操作：
```
$ etcdctl lease grant 50
lease 2deb68788ab64ec4 granted with TTL(50s)
```
以下是保持同一租约有效的命令：
```
$ etcdctl lease keep-alive 2deb68788ab64ec4
lease 2deb68788ab64ec4 keepalived with TTL(50)
lease 2deb68788ab64ec4 keepalived with TTL(50)
lease 2deb68788ab64ec4 keepalived with TTL(50)
...
```
### 获取租赁信息
应用程序可能想要了解租赁信息，以便它们可以续订或检查租赁是否仍然存在或已过期。应用程序也可能想知道特定租约所附的密钥。

假设我们完成了以下一系列操作：
```
# grant a lease with 500 second TTL
$ etcdctl lease grant 500
lease 2deb68788ab64ec6 granted with TTL(500s)

# 查看 lease list
$ etcdctl lease list
found 1 leases
2deb68788ab64ec6

# attach key name to lease 2deb68788ab64ec6
$ etcdctl put name hanru --lease=2deb68788ab64ec6
OK

# attach key name2 to lease 2deb68788ab64ec6
$ etcdctl put name2 wangergou  --lease=2deb68788ab64ec6
OK
```
以下是获取有关租赁信息的命令：
```
$ etcdctl lease timetolive 2deb68788ab64ec6
lease 2deb68788ab64ec6 granted with TTL(500s), remaining(417s)
etcd_caozuo24
```
以下是获取有关租赁信息的命令以及租赁附带的密钥：
```
$ etcdctl lease timetolive --keys 2deb68788ab64ec6
lease 2deb68788ab64ec6 granted with TTL(500s), remaining(215s), attached keys([name name2])

# if the lease has expired  it will give the below response:
lease 2deb68788ab64ec6 already expired
```