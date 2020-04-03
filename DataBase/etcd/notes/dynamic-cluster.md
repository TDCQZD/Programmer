# etcd集群的动态搭建
* https://www.chaindesk.cn/witbook/36/526
## 什么是集群的Discovery
静态配置前提是在搭建集群之前已经提前知道各节点的信息，而实际应用中可能存在预先并不知道各节点ip的情况，不过很多时候，你只知道你要搭建一个多大(包含多少节点)的集群，但是并不能事先知道这几个节点的ip，从而无法使用--initial-cluster参数。

这个时候，你就需要使用discovery的方式来搭建etcd集群。discovery方式有两种：etcd discovery和DNS discovery。

> 在一些案例中，集群伙伴的 IP 可能无法提前知道。当使用云提供商或者网络使用 DHCP 时比较常见。在这些情况下，相比指定静态配置，使用使用已经存在的 etcd 集群来启动一个新的。我们称这个过程为"发现"。