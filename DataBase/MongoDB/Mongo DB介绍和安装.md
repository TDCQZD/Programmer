# Mongo DB
MongoDB 是一个基于分布式文件存储的数据库。由 C++ 语言编写。旨在为 WEB 应用提供可扩展的高性能数据存储解决方案。

MongoDB 是一个介于关系数据库和非关系数据库之间的产品，是非关系数据库当中功能最丰富，最像关系数据库的。
## 什么是MongoDB ?
- MongoDB 是由C++语言编写的，是一个基于分布式文件存储的开源数据库系统。

- 在高负载的情况下，添加更多的节点，可以保证服务器性能。

- MongoDB 旨在为WEB应用提供可扩展的高性能数据存储解决方案。

- MongoDB 将数据存储为一个文档，数据结构由键值(key=>value)对组成。MongoDB 文档类似于 JSON 对象。字段值可以包含其他文档，数组及文档数组。
## Mongo DB 特性
- 层级 Database-Collection-Document

- 灵活的类JSON数据存储，每条文档的字段可以完全不同

- 方便的即席查询（ad hoc queries）、索引（indexing）和实时聚合（aggregation）

- 使用update()命令可以实现替换完整的文档（数据）或者一些指定的数据字段

- MongoDB允许在服务端执行脚本

## Mongo DB 下载和安装

mongodb官网：[https://www.mongodb.com/download-center#community]()

### Linux平台安装MongoDB

**下载安装** 

直接用 curl 命令下载，以 64 位 Linux 平台为例：
```
curl -O https://fastdl.mongodb.org/linux/mongodb-linux-x86_64-ubuntu1604-4.0.4.tgz # 下载
tar -zxvf  mongodb-linux-x86_64-ubuntu1604-4.0.4.tgz # 解压
mv mongodb-linux-x86_64-ubuntu1604-4.0.4 mongodb # 将解压包拷贝到指定目录
```
**环境变量配置**

MongoDB 的可执行文件位于 bin 目录下，所以可以将其添加到 PATH 路径中

```
export PATH=$PATH:/opt/mongodb/bin
```
**创建数据库目录**

MongoDB的数据存储在data目录的db目录下，但是这个目录在安装过程不会自动创建，所以你需要手动创建data目录，并在data目录中创建db目录。

以下实例中我们将data目录创建于根目录下(/)。

> 注意：/data/db 是 MongoDB 默认的启动的数据库路径(--dbpath)。
```
mkdir -p /data/db
```
查看Mongo DB的存储:
```
cd /data/db/
ls
```
```
root@Aws:/data/db# ls
collection-0--5371242798992439506.wt  index-3--8531195751554862470.wt  storage.bson
collection-2--5371242798992439506.wt  index-5--5371242798992439506.wt  WiredTiger
collection-2--8531195751554862470.wt  index-6--5371242798992439506.wt  WiredTigerLAS.wt
collection-4--5371242798992439506.wt  journal                          WiredTiger.lock
diagnostic.data                       _mdb_catalog.wt                  WiredTiger.turtle
index-1--5371242798992439506.wt       mongod.lock                      WiredTiger.wt
index-3--5371242798992439506.wt       sizeStorer.wt
root@Aws:/data/db# 

```

## 启动Mongo服务和客户端

### 启动 mongo 服务 和 Mongo DB 连接
``` 
$ ./mongod  
```
> 注意：如果你的数据库目录不是/data/db，可以通过 --dbpath 来指定。

### MongoDB后台管理 Shell

如果你需要进入MongoDB后台管理，你需要先打开mongodb装目录的下的bin目录，然后执行mongo命令文件。

MongoDB Shell是MongoDB自带的交互式Javascript shell,用来对MongoDB进行操作和管理的交互式环境。

当你进入mongoDB后台后，它默认会链接到 test 文档（数据库）：
```
$ ./mongo
MongoDB shell version: 3.0.6
connecting to: test
Welcome to the MongoDB shell.
```
### Mongo DB 连接

**标准 URI 连接语法：**
```
mongodb://[username:password@]host1[:port1][,host2[:port2],...[,hostN[:portN]]][/[database][?options]]
```
**参数说明：**
- mongodb:// 这是固定的格式，必须要指定。
- username:password@ 可选项，用户名/密码。
- host1 必须指定至少一个host, host1 是这个URI唯一必须要填写的，它指定了要连接服务器的地址。如果要连接复制集，需要指定多个主机地址。
- portN 可选的指定端口，如果不填，默认为27017。
- /database 如果指定username:password@，连接并验证登陆指定数据库。若不指定，默认打开 test 数据库。
- ?options 连接选项。

**示例：**

``` 
$ ./mongo mongodb://admin:123456@localhost/test
mongodb:
```
