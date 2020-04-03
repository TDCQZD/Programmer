# 分布式系统
- RPC和对象序列化
- 分布式内存缓存技术、分布式内存计算
- 分布式存储
- 分布式计算
- 全文检索
- 消息队列MQ
- 容器
## 分布式文件系统
- IPFS
- GFS
- HDFS
- Lustre 
- Ceph 
- GridFS 
- mogileFS
- TFS
- FastDF
- MooseFS
## 主流的分布式存储文件系统
- GFS
- AFS 
- Lustre
### GFS
GFS（Google file system）被称为谷歌文件系统，其性能、可扩展性、可靠性和可用性都收到了肯定，主要部件包括一个Master和n个chunkserver，和Chunk Server（数据块服务器）同时可以被多个客户Client访问。

不同于传统的文件系统，GFS不再将组建错误当成异常，而是将其看做一种常见情况予以处理。对待文件的大小一直是文件系统要考虑的问题，对于任何一种文件系统，成千上万的几KB的系统，总会压死内存，所以，对于大型的文件， 管理要高效，对于小型文件也要支持，但是并没有进行优化。

在GFS中chunk server大小被固定为64MB，这样的块规模比一般的文件系统的块规模要大得多，可以减少元数据metadata开销，减少Master的交互，但太大的块规模也会产生内部碎片，或者同一个Chunk中存在多个小文件可能产生访问热点。

GFS主要运行在大量运行Linux系统的普通机器上，从而降低了其硬件成本。但一系列冗余备份、快速恢复等技术保证其正常和高效运行，GFS也是实现非结构化数据的主要技术和文件系统。
### AFS
AFS是Andrew File System的简称。AFS将文件系统的可扩展性放在了设计和实践的首要位置，因此AFS拥有过很好的扩展性，能够轻松支持数百个节点，甚至数千个节点的分布式环境。AFS由卡内基美隆大学最初设计开发，目前已经相当成熟，用于研究和部分大型网络之上。 

AFS主要组建包括：Cells、AFS clients、基本存储单元Volumes、AFS servers和Volume replication。AFS实现的是模块化的，并不要求在每台服务器上运行所有服务器进程。AFS拥有良好可扩展性，客户端华村能够带来性能的提升和可用性的提高、AFS的缺点在于管理员界面友好性不足，需要更多的专业知识来支持AFS。
### Lustre平行分布式文件系统
Lustre是HP，Intel，Cluster File System公司联合美国能源部开发的Linux集群并行文件系统，名称来源于Linux和Clusters。同时Lustre也是一个遵循GPL许可协议的开源软件，Lustre也被称为平行分布式文件系统，常用于大型计算机集群和超级电脑中。

Lustre的主要组建包括：元数据服务器（Metadataservers， MDSs）、对象存储服务器（objectstorage servers， OSSs）和客户端。其中MDSs提供元数据服务，MGS管理服务器提供Lustre文件系统配置信息，OSS对象存储服务器expose块设备提供数据。

Lustre文件系统针对大文件读写进行了优化，能够提高性能的IO能力；在源数据独立存储、服务和网络失效的快速恢复、基于意图的分布式锁管理和系统可快速配置方面优异。

## 分布式内存缓存技术、分布式内存计算

缓存系统常用的缓存淘汰策略：
1. Least Frequently Used（LFU）策略--计算使用频率，优先淘汰最不常用的缓存条目，CPU的cache所采用的淘汰策略即为LFU策略；
2. Least Recently Used（LRU）策略--淘汰最近最少使用的条目；
3. Adaptive Replacement Cache（ARC）策略--该策略由两个LRU组成，第1个LRU包含的条目是最近只被使用过一次的条目，第2个LRU包含的是最近被使用过二次的条目；
4. 其他还有一些基于缓存时间的淘汰策略，比如淘汰存活时间超过5分钟的缓存条目。

分布式缓存都采用Hash算法进行数据分片，将数量庞大的缓存项均匀分布到集群中的每个节点上，比如Redis3.0开始实现的分布式集群功能就采用了Hash算法，将缓存项均匀分布到16384个Slot上去。以Redis2.x为基础改造的Codis是国内分布式缓存开源的一个典范，出自豆瓣网。Memcache本身并没有提供集群功能，但很多客户端Driver实现了Hash算法分配逻辑，因此也可以看成是一种分布式缓存的解决方案。

内存计算产品：商业的SAP Hana、开源的VoltDB等。VoltDB是一种开源的高性能的内存关系型数据库，提供社区版和商业版，是一种NewSql，是一个借鉴并基于HSQL的分配内存数据库集群。

## 分布式缓存中间件
### Redis
redis是以key-value的形式存储数据，是一个非关系型的，分布式开源的（BSD协议），水平可扩展的缓存服务器。

redis用途：缓存（StackOverFlow），数据库（微博），消息中间件（微博）
### Memcached
memcached 是由 Danga Interactive 开发并使用 BSD 许可的一种通用的分布式内存缓存系统。也是以key-value的形式存储数据。由于它的工作机制是在内存中开辟一块空间，然后建立一个HashTable，Memcached自管理这些HashTable。
### SSDB
SSDB是一个快速的用来存储十亿级别列表数据的开源NoSql数据库，可以替代Redis数据库，是Redis的100倍容量。LevelDB网络支持，使用C/C++开发。与Redis完美兼容。

## 全文检索

- Lucence Core：Java编写的核心类库，提供全文检索功能的底层API与SDK。
- Solr：基于Lucence Core开发的高性能搜索服务，提供了REST API的高层封装接口，还提供了一个Web管理界面。
- PyLucene：一个Python版的Lucene Core的高仿实现。
- ElasticSearch：也是基于Lucence的分布式全文检索中间件。
## 分布式消息中间件
消息队列(Message Queue，简称MQ)。消息中间件作为实现分布式消息系统可拓展、可伸缩性的关键组件，具有高吞吐量、高可用等等优点。

常用的消息队列有RocketMQ、kafka、ActiveMQ、RabbitMQ、ZeroMQ等等 

- ActiveMQ、RabbitMQ与 kafka、RocketMQ有很大的区别就是前2个只支持主从模式，后2个是分布式消息系统，支持分布式。 
- 持久化消息比较： ZeroMQ不支持，ActiveMQ和RabbitMQ都支持。

### RocketMQ
RocketMQ 是一款开源的分布式消息系统，基于高可用分布式集群技术，提供低延时的、高可靠的消息发布与订阅服务。

### Kafka
Kafka是没有重发机制的消息队列。它以可水平扩展和高吞吐率而被广泛使用。目前越来越多的开源分布式处理系统如Cloudera、Apache Storm、Spark都支持与Kafka集成。kafka不完全符合jms规范，注重吞吐量，类似udp 和 tcp。kafka吞吐量大 ，对数据不是百分之百保证的，会有数据丢失，不是百分百送达。所以kafka适合大数据量流转， 比如日志数据 比如用作统计的数据

### ActiveMQ
ActiveMQ类似于ZemoMQ，它可以部署于代理模式和P2P模式。类似于RabbitMQ，它易于实现高级场景，而且只需付出低消耗。

### RabbitMQ
RabbitMQ是AMQP协议领先的一个实现，它实现了代理(Broker)架构，意味着消息在发送到客户端之前可以在中央节点上排队。此特性使得RabbitMQ易于使用和部署，适宜于很多场景如路由、负载均衡或消息持久化等，用消息队列只需几行代码即可搞定。但是，这使得它的可扩展性差，速度较慢，因为中央节点增加了延迟，消息封装后也比较大。

### ZeroMQ
ZeroMQ是一个非常轻量级的消息系统，专门为高吞吐量/低延迟的场景开发，在金融界的应用中经常可以发现它。

## 参考资料
* https://chaser520.iteye.com/blog/2428231