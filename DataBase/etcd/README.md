# etcd

## 什么是etcd
Etcd 按照官方介绍

> Etcd is a distributed, consistent key-value store for shared configuration and service discovery

是一个分布式的，一致的 key-value 存储，主要用途是共享配置和服务发现。

etcd 是 CoreOS 团队于 2013 年 6 月发起的开源项目，它的目标是构建一个高可用的分布式键值（key-value）数据库，基于 Go 语言实现。在分布式系统中，各种服务的配置信息的管理分享，**服务的发现**是一个很基本同时也是很重要的问题。CoreOS 项目就希望基于 etcd 来解决这一问题。

> etcd 目前在 github.com/coreos/etcd 进行维护。

受到 Apache ZooKeeper 项目和 doozer 项目的启发，etcd 在设计的时候重点考虑了下面四个要素：

- 简单：良好的定义，面向用户的API (gRPC)
- 安全：带有可选客户端证书认证的自动TLS
- 快速：支持并发 1k/s 的写操作
- 可靠：支持分布式结构，基于 Raft 的一致性算法

**etcd是一个高可用的键值存储系统，主要用于共享配置和服务发现。**etcd是由CoreOS开发并维护的，灵感来自于 ZooKeeper 和 Doozer，它使用Go语言编写，并通过Raft一致性算法处理日志复制以保证强一致性。Raft是一个来自Stanford的新的一致性算法，适用于分布式系统的日志复制，Raft通过选举的方式来实现一致性，在Raft中，任何一个节点都可能成为Leader。Google的容器集群管理系统Kubernetes、开源PaaS平台Cloud Foundry和CoreOS的Fleet都广泛使用了etcd。

etcd 集群的工作原理基于 raft 共识算法 (The Raft Consensus Algorithm)。etcd 在 0.5.0 版本中重新实现了 raft 算法，而非像之前那样依赖于第三方库 go-raft 。raft 共识算法的优点在于可以在高效的解决分布式系统中各个节点日志内容一致性问题的同时，也使得集群具备一定的容错能力。即使集群中出现部分节点故障、网络故障等问题，仍可保证其余大多数节点正确的步进。甚至当更多的节点（一般来说超过集群节点总数的一半）出现故障而导致集群不可用时，依然可以保证节点中的数据不会出现错误的结果。 

- [etcd的安装](./notes/install.md)
- [etcd集群的静态搭建](./notes/static-cluster.md)
- [etcd集群的动态搭建](./notes/dynamic-cluster.md)
- [etcd集群的member操作](./notes/member.md)
- [etcd操作](./notes/operating.md)
- [etcd 中的raft协议](./notes/raft.md )
- [etcd的存储分析](./notes/storage.md )
- [解析etcd设计原理](./notes/design-principle.md)

## 项目实战
[etcd-manage项目](./etcd-manage.md)
## 参考资料
* https://github.com/etcd-io/etcd
* https://www.chaindesk.cn/witbook/36
