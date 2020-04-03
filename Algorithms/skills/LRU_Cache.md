# LRU Cache

## 什么是 LRU 算法
就是一种缓存淘汰策略。

LRU 缓存淘汰算法就是一种常用策略。LRU 的全称是 `Least Recently Used`，也就是说我们认为最近使用过的数据应该是「有用的」，很久都没用过的数据应该是无用的，内存满了就优先删那些很久没用过的数据。

## LRU Cache 
1. Least Recently Used（最近最少使用）
2. Hash Table + Double LinkedList（哈希表+双向链表）
3. get O(1) And set O(1) 

### LFU Cache
1. LFU - least frequently used（最近最不常用页面置换算法）
2. LRU - least recently usd（最近最少使用页面置换算法）

## LRU 算法描述
LRU 算法实际上是让你设计数据结构：首先要接收一个 `capacity` 参数作为缓存的最大容量，然后实现两个 API:
- 一个是 put(key, val) 方法存入键值对;
- 另一个是 get(key) 方法获取 key 对应的 val，如果 key 不存在则返回 -1。
- 注意:get 和 put 方法必须都是 O(1) 的时间复杂度.

示例：LRU 算法怎么工作
```
/* 缓存容量为 2 */
LRUCache cache = new LRUCache(2);
// 你可以把 cache 理解成一个队列
// 假设左边是队头，右边是队尾
// 最近使用的排在队头，久未使用的排在队尾
// 圆括号表示键值对 (key, val)

cache.put(1, 1);
// cache = [(1, 1)]
cache.put(2, 2);
// cache = [(2, 2), (1, 1)]
cache.get(1);       // 返回 1
// cache = [(1, 1), (2, 2)]
// 解释：因为最近访问了键 1，所以提前至队头
// 返回键 1 对应的值 1
cache.put(3, 3);
// cache = [(3, 3), (1, 1)]
// 解释：缓存容量已满，需要删除内容空出位置
// 优先删除久未使用的数据，也就是队尾的数据
// 然后把新的数据插入队头
cache.get(2);       // 返回 -1 (未找到)
// cache = [(3, 3), (1, 1)]
// 解释：cache 中不存在键为 2 的数据
cache.put(1, 4);    
// cache = [(1, 4), (3, 3)]
// 解释：键 1 已存在，把原始值 1 覆盖为 4
// 不要忘了也要将键值对提前到队头
```
## LRU 算法设计

分析上面的操作过程，要让 put 和 get 方法的时间复杂度为 O(1)，我们可以总结出 cache 这个数据结构必要的条件：查找快，插入快，删除快，有顺序之分。

因为显然 cache 必须有顺序之分，以区分最近使用的和久未使用的数据；而且我们要在 cache 中查找键是否已存在；如果容量满了要删除最后一个数据；每次访问还要把数据插入到队头。

那么，什么数据结构同时符合上述条件呢？哈希表查找快，但是数据无固定顺序；链表有顺序之分，插入删除快，但是查找慢。所以结合一下，形成一种新的数据结构：哈希链表。

LRU 缓存算法的核心数据结构就是哈希链表，双向链表和哈希表的结合体。

## 代码模板
``` python
class LRUCache(object):

    def _init_(selt, capacity):
        self.dic = collections.OrderedDict()
        self.remain = capacity

    def get(self, key):
        if key not in self.dic:
            return -1
        v = self.dic.pop(key)
        self.dic[key] =  # set key as the newest one
        return v
    
    def put(self, key, value):
        if key in self.dic:
            self.dic.pop(key)
        else:
            if self.remain > 0:
                self.remain -= 1
            else: # self.dic in full
                self.dic.popitem(last=False) # 推出最后元素
            self.dic[key] = value
````


## 案例
1. 设计和实现一个LRU Cache缓存机制
> https://github.com/labuladong/fucking-algorithm/blob/master/%E9%AB%98%E9%A2%91%E9%9D%A2%E8%AF%95%E7%B3%BB%E5%88%97/LRU%E7%AE%97%E6%B3%95.md

## 参考资料
* [中文页面](https://zh.wikipedia.org/wiki/快取文件置換機制)
* [英文页面](https://en.wikipedia.org/wiki/Cache_replacement_policies)
