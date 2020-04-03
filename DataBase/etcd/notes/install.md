# etcd的安装
- 下载etcd安装包安装
- 源码编译安装
- 下载etcd二进制
## 下载etcd安装包安装
### 1. 打开终端，输入以下命令：
```
// 因为包比较大，需要耐心等待一会儿，而且此处需要翻墙
$ go get -v go.etcd.io/etcd
```
下载成功后，可以在src目录下看到一个go.etcd.io目录：
 
 此时在bin目录下已经有了etcd命令.
### 2. etcd 交互
 其实etcd支持多种交互方式，包括etcdctl，http和客户端连接等等。

etcd 对外通过 HTTP API 对外提供服务，这种方式方便测试（通过 curl 或者其他工具就能和 etcd 交互），也很容易集成到各种语言中（每个语言封装 HTTP API 实现自己的 client 就行）。

etcdctl 是一个命令行客户端，它能提供一些简洁的命令，供用户直接跟 etcd 服务打交道，而无需基于 HTTP API 方式。这在某些情况下将很方便，例如用户对服务进行测试或者手动修改数据库内容。我们也推荐在刚接触 etcd 时通过 etcdctl 命令来熟悉相关的操作，这些操作跟 HTTP API 实际上是对应的。

接下来我们就安装etcdctl( a simple command line client.)：
```
$ go get -v go.etcd.io/etcd/etcdctl
```
然后我们可以在bin目录下看见etcdctl命令


### 检查安装的版本
然后输入命令查看一下安装的版本：
```
$ etcdctl version
```

> etcd 是服务主文件，etcdctl 是提供给用户的命令客户端

## 源码编译安装
GOPATH/src目录下创建文件夹：go.etcd.io：

然后通过cd命令进入该目录：
```
$ cd $GOPATH/src/go.etcd.io
```
然后执行以下命令，从github上拷贝etcd的安装包到该目录下：
```
$ git clone https://github.com/coreos/etcd.git
```
拷贝好后在go.etcd.io目录下有个子目录etcd，进入该子目录：
```
$ cd etcd
```
然后执行build方法，生成etcd和etcdctl命令：
```
$ ./build
```
## 下载etcd二进制
还有另外一种下载安装的方法：

直接下载etcd二进制 （包含etcd、etcdctl） https://github.com/coreos/etcd/releases

# 启动etcd
要想启动etcd服务很简单，执行etcd命令即可：
```
$ etcd
```
当然你也可以自己build生成etcd和etcdctl命令：

首先进入etcd目录：
```
$ cd $GOPATH/src/go.etcd.io/etcd
```
然后执行build命令：
```
$ ./build
```
此时就会在etcd中生成一个bin目录，里面有刚刚build出来的etcd和etcdctl命令.


然后我们同样也可以通过执行etcd命令来启动：
```
$ ./bin/etcd
```
运行 etcd，将默认组建一个两个节点的集群
```
2019-04-13 11:16:22.473850 I | etcdmain: etcd Version: 3.3.0+git
2019-04-13 11:16:22.474048 I | etcdmain: Git SHA: Not provided (use ./build instead of go build)
2019-04-13 11:16:22.474109 I | etcdmain: Go Version: go1.10.7
2019-04-13 11:16:22.474165 I | etcdmain: Go OS/Arch: linux/amd64
2019-04-13 11:16:22.474202 I | etcdmain: setting maximum number of CPUs to 1, total number of available CPUs is 1
2019-04-13 11:16:22.474245 W | etcdmain: no data-dir provided, using default data-dir ./default.etcd
2019-04-13 11:16:22.474653 N | etcdmain: the server is already initialized as member before, starting as etcd member...
2019-04-13 11:16:22.478167 I | embed: name = default
2019-04-13 11:16:22.478226 I | embed: data dir = default.etcd
2019-04-13 11:16:22.478266 I | embed: member dir = default.etcd/member
2019-04-13 11:16:22.478301 I | embed: heartbeat = 100ms
2019-04-13 11:16:22.478334 I | embed: election = 1000ms
2019-04-13 11:16:22.478365 I | embed: snapshot count = 100000
2019-04-13 11:16:22.478410 I | embed: advertise client URLs = http://localhost:2379
2019-04-13 11:16:22.478452 I | embed: initial advertise peer URLs = http://localhost:2380
2019-04-13 11:16:22.478487 I | embed: initial cluster = 
2019-04-13 11:16:22.620368 I | etcdserver: restarting member 8e9e05c52164694d in cluster cdf818194e3a8c32 at commit index 16
{"level":"info","ts":1555125382.6209052,"caller":"raft/raft.go:712","msg":"8e9e05c52164694d became follower at term 5"}
{"level":"info","ts":1555125382.6220214,"caller":"raft/raft.go:389","msg":"newRaft 8e9e05c52164694d [peers: [], term: 5, commit: 16, applied: 0, lastindex: 16, lastterm: 5]"}
2019-04-13 11:16:22.870259 W | auth: simple token is not cryptographically signed
2019-04-13 11:16:22.872570 I | etcdserver: starting server... [version: 3.3.0+git, cluster version: to_be_decided]
2019-04-13 11:16:22.880841 I | embed: listening for peers on 127.0.0.1:2380
2019-04-13 11:16:22.882063 I | etcdserver/membership: added member 8e9e05c52164694d [http://localhost:2380] to cluster cdf818194e3a8c32
2019-04-13 11:16:22.882226 N | etcdserver/membership: set the initial cluster version to 3.3
2019-04-13 11:16:22.882313 I | etcdserver/api: enabled capabilities for version 3.3
{"level":"info","ts":1555125384.522473,"caller":"raft/raft.go:922","msg":"8e9e05c52164694d is starting a new election at term 5"}
{"level":"info","ts":1555125384.5225365,"caller":"raft/raft.go:725","msg":"8e9e05c52164694d became candidate at term 6"}
{"level":"info","ts":1555125384.5225604,"caller":"raft/raft.go:820","msg":"8e9e05c52164694d received MsgVoteResp from 8e9e05c52164694d at term 6"}
{"level":"info","ts":1555125384.5225787,"caller":"raft/raft.go:777","msg":"8e9e05c52164694d became leader at term 6"}
{"level":"info","ts":1555125384.5225928,"caller":"raft/node.go:330","msg":"raft.node: 8e9e05c52164694d elected leader 8e9e05c52164694d at term 6"}
2019-04-13 11:16:24.523009 I | etcdserver: published {Name:default ClientURLs:[http://localhost:2379]} to cluster cdf818194e3a8c32
2019-04-13 11:16:24.523188 E | etcdmain: forgot to set Type=notify in systemd service file?
2019-04-13 11:16:24.523214 I | embed: ready to serve client requests
2019-04-13 11:16:24.524229 N | embed: serving insecure client requests on 127.0.0.1:2379, this is strongly discouraged!


```
> 默认 2379 端口处理客户端的请求，2380 端口用于集群各成员间的通信。

说明：

|名称|	说明|
|------|------|
|name = default|	name表示节点名称，默认为default
|data dir = default.etcd|	指定节点的数据存储目录，保存日志和快照等，默认为当前工作目录default.etcd/|
|heartbeat = 100ms|	leader发送心跳到followers的间隔时间
|election = 1000ms|	重新投票的超时时间，如果follow在该时间间隔没有收到心跳包，会触发重新投票，默认为1000ms
|advertise client URLs = http://localhost:2379|	对外提供服务的地址：比如 http://ip:2379,http://127.0.0.1:2379 ，客户端会连接到这里和 etcd 交互|
|initial advertise peer URLs = http://localhost:2380|	和集群中其他节点通信|
##  测试
既然etcd是一个分布式的存储系统，我们先简单测试一下，存储一个key-value：
```
$ etcdctl put foo hanru
```
打印出ok，就表示存储成功，要注意先启动etcd服务：


