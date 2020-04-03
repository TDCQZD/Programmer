# DataBase

## NOSQL
### MongoDB
MongoDB 是一个基于分布式文件存储的数据库。由 C++ 语言编写。旨在为 WEB 应用提供可扩展的高性能数据存储解决方案。

MongoDB 是一个介于关系数据库和非关系数据库之间的产品，是非关系数据库当中功能最丰富，最像关系数据库的。

MongoDB 旨在为WEB应用提供可扩展的高性能数据存储解决方案。

MongoDB 将数据存储为一个文档，数据结构由键值(key=>value)对组成。MongoDB 文档类似于 JSON 对象。字段值可以包含其他文档，数组及文档数组。
### LevelDB
LevelDB是一个由Google公司所研发的键／值对（Key/Value Pair）嵌入式数据库管理系统编程库.

### BoltDB
Bolt 是一个纯键值存储的 Go 数据库，启发自 Howard Chu 的 LMDB. 它旨在为那些无须一个像 Postgres 和 MySQL 这样有着完整数据库服务器的项目，提供一个简单，快速和可靠的数据库。

由于 Bolt 意在用于提供一些底层功能，简洁便成为其关键所在。它的 API 并不多，并且仅关注值的获取和设置。仅此而已。


BoltDB是一个嵌入式key/value的数据库，即只需要将其链接到你的应用程序代码中即可使用BoltDB提供的API来高效的存取数据。而且BoltDB支持完全可序列化的ACID事务，让应用程序可以更简单的处理复杂操作。

BoltDB设计源于LMDB，具有以下特点：

- 直接使用API存取数据，没有查询语句；
- 支持完全可序列化的ACID事务，这个特性比LevelDB强；
- 数据保存在内存映射的文件里。没有wal、线程压缩和垃圾回收；
- 通过COW技术，可实现无锁的读写并发，但是无法实现无锁的写写并发，这就- 注定了读性能超高，但写性能一般，适合与读多写少的场景。
最后，BoltDB使用Golang开发，而且被应用于influxDB项目作为底层存储。
### CouchDB

CouchDB是由Apache软件基金会开发的一个开源数据库。重点是易于使用，拥抱网络。它是一个NoSQL文档存储数据库。它使用JSON，存储数据（文档），使用java脚本作为其查询语言来转换文档，使用http协议为api访问文档，使用Web浏览器查询索引。它是一个多主应用程序在2005年发布，它成为一个apache项目在2008年。

CouchDB是一个开源的NoSQL数据库，专注于易用性。 它由Apache开发，完全兼容web。 CouchDB使用JSON来存储数据，使用JavaScript作为查询语言来转换文档，使用MapReduce和HTTP作为API。

CouchDB提供易于使用的复制，可以在数据库和机器之间复制，共享和同步数据。
### RocksDB

### TiDB

## 请解释关系型数据库概念及主要特点？

概念：

关系型数据库是支持采用了关系模型的数据库，简单来说，关系模型就是指二维表模型，而一个关系数据库就是由二维表及其之间的联系所组成的一个数据组织。

特点：

最大的特点就是事务的一致性。

优点：容易理解、使用方便、易于维护、支持SQL。

缺点：

1. 高并发读写需求：网站的用户并发非常高，往往达到每秒上万次读写请求，对于传统关系型数据库来说，硬盘I/O是一个很大的瓶颈。

2. 海量数据的高效读写：对于数据量巨大的网站来说，关系型数据库的查询效率非常低。

3. 固定的表结构。

 

## 请说出关系型数据库的典型产品、特点及应用场景？

1.SQLserver

特点：真正的客户机/服务器体系结构

      图形化用户界面

      丰富的编程接口工具

      与Windows NT完全集成

      具有很好的伸缩性

应用场景：

主机为Windows系统，主要应用于web网站的建设，承载中小型web后台数据。

2.MySQL

特点：体积小，总拥有成本低，开放源代码，可运行在多数系统平台上，轻量级易扩展。

应用场景：广泛的应用在Internet上的中小型网站中。

3.Oracle

特点：跨平台运行，安全性方面，性能最高。对硬件要求高，价格昂贵。

应用场景：大部分国企事业单位都用Oracle，在电信行业占用最大的份额。

 

## 请解释非关系型数据库概念及主要特点？

1.  使用键值对存储数据，且结构不固定

2. 一般不支持ACID特性。

3. 基于键值对，数据没有耦合性，容易扩展。

4. 不提供SQL支持，学习和使用成本较高。


## 请说出非关系型数据库的典型产品、特点及应用场景？

MongoDB

特点：1.高性能，易部署，易使用。

      2.面向集合存储，易存储对象类型的数据。

      3.模式自由

      4.自动处理碎片，以支持云计算层次的扩展性。

应用场景：

      网站数据：mongodb非常适合实时的插入，更新与查询。

      缓存：适合作为信息基础设施的缓存层

      大尺寸、低价值的数据

      高伸缩性的场景

Redis

特点：1.性能极高，能支持超过100k+每秒的读写频率

      2.丰富的数据类型

      3.所有操作都是原子性的

使用场景：

      少量的数据存储，高速读写访问

SQLlite

特点：

1.嵌入式的，零配置，无需安装和管理配置

2.ACID事务

3.存储在单一磁盘文件中的一个完整的数据库。

应用场景：

1.需要数据库的小型桌面软件。

2.需要数据库的手机软件。

3.作为数据容器的应用场景。

## Redis、MongoDB、LevelDB主流 NoSQL 数据库的对比

### Redis
Redis 是一个开源的使用ANSI C语言编写、支持网络、可基于内存亦可持久化的日志型、Key-Value数据库，并提供多种语言的API.

* 数据格式

      - 字符串(String), 哈希(Hash/Map), 列表(list), 集合(sets) 和 有序集合(sorted sets)五种类型
* 性能

      - Redis数据库完全在内存中，因此处理速度非常快，每秒能执行约11万集合，每秒约81000+条记录

      - Redis的数据能确保一致性——所有Redis操作是原子性（Atomicity，意味着操作的不可再分，要么执行要么不执行）的，这保证了如果两个客户端同时访问的Redis服务器将获得更新后的值
* 持久化

      - 通过定时快照（snapshot）和基于语句的追加（AppendOnlyFile，aof）两种方式，redis可以支持数据持久化——将内存中的数据存储到磁盘上，方便在宕机等突发情况下快速恢复。
* 优势

      - 非常丰富的数据结构；
      - Redis提供了事务的功能，可以保证一串 命令的原子性，中间不会被任何操作打断；
      -  数据存在内存中，读写非常的高速，可以达到10w/s的频率。
* 缺点

      - 持久化功能体验不佳——通过快照方法实现的话，需要每隔一段时间将整个数据库的数据写到磁盘上，代价非常高；而aof方法只追踪变化的数据，类似于mysql的binlog方法，但追加log可能过大，同时所有操作均要重新执行一遍，恢复速度慢；

      - 由于是内存数据库，所以，单台机器，存储的数据量，跟机器本身的内存大小。虽然redis本身有key过期策略，但是还是需要提前预估和节约内存。如果内存增长过快，需要定期删除数据。

* 适用场景

      适用于数据变化快且数据库大小可遇见（适合内存容量）的应用程序
### MongoDB
MongoDB 是一个高性能，开源，无模式的文档型数据库，开发语言是C++。它在许多场景下可用于替代传统的关系型数据库或键/值存储方式。

* 数据格式

      在 MongoDB 中，文档是对数据的抽象，它的表现形式就是我们常说的 BSON（Binary JSON ）。

      BSON 是一个轻量级的二进制数据格式。MongoDB 能够使用 BSON，并将 BSON 作为数据的存储存放在磁盘中。

* 性能

      MongoDB 目前支持的存储引擎为内存映射引擎。当 MongoDB 启动的时候，会将所有的数据文件映射到内存中，然后操作系统会托管所有的磁盘操作。这种存储引擎有以下几种特点：
      - MongoDB 中关于内存管理的代码非常精简，毕竟相关的工作已经有操作系统进行托管。

      - MongoDB 服务器使用的虚拟内存将非常巨大，并将超过整个数据文件的大小。不用担心，操作系统会去处理这一切。


      另外，MongoDB 提供了全索引支持：包括文档内嵌对象及数组。Mongo的查询优化器会分析查询表达式，并生成一个高效的查询计划。通常能够极大的提高查询的效率。
* 持久化

      MongoDB 在1.8版本之后开始支持 journal，就是我们常说的 redo log，用于故障恢复和持久化。

      当系统启动时，MongoDB 会将数据文件映射到一块内存区域，称之为Shared view，在不开启 journal 的系统中，数据直接写入shared view，然后返回，系统每60s刷新这块内存到磁盘，这样，如果断电或down机，就会丢失很多内存中未持久化的数据。

      当系统开启了 journal 功能，系统会再映射一块内存区域供 journal 使用，称之为 private view，MongoDB 默认每100ms刷新 privateView 到 journal，也就是说，断电或宕机，有可能丢失这100ms数据，一般都是可以忍受的，如果不能忍受，那就用程序写log吧
* 优势

      - 强大的自动化 shading 功能（更多戳这里）；

      - 全索引支持，查询非常高效；

      -  面向文档（BSON）存储，数据模式简单而强大。

      -  支持动态查询，查询指令也使用JSON形式的标记，可轻易查询文档中内嵌的对象及数组。

      -  支持 javascript 表达式查询，可在服务器端执行任意的 javascript函数
* 缺点

      1. 单个文档大小限制为16M，32位系统上，不支持大于2.5G的数据；

      2. 对内存要求比较大，至少要保证热数据（索引，数据及系统其它开销）都能装进内存；

      3. 非事务机制，无法保证事件的原子性
* 适用场景

      1. 适用于实时的插入、更新与查询的需求，并具备应用程序实时数据存储所需的复制及高度伸缩性；

      2. 非常适合文档化格式的存储及查询；

      3. 高伸缩性的场景：MongoDB 非常适合由数十或者数百台服务器组成的数据库。

      4. 对性能的关注超过对功能的要求
### LevelDB
LevelDB 是由谷歌重量级工程师（Jeff Dean 和 Sanjay Ghemawat）开发的开源项目，它是能处理十亿级别规模 key-value 型数据持久性存储的程序库，开发语言是C++。

除了持久性存储，LevelDB 还有一个特点是 —— 写性能远高于读性能
* 数据格式
      - slice
* 性能
* 持久化
* 优势

      1. 操作接口简单，基本操作包括写记录，读记录和删除记录，也支持针对多条操作的原子批量操作；

      2. 写入性能远强于读取性能，

      3. 数据量增大后，读写性能下降趋平缓。
* 缺点
      1. 随机读性能一般;

      2. 对分布式事务的支持还不成熟。而且机器资源浪费率高。
* 适用场景
      适用于对写入需求远大于读取需求的场景