# Docker 集群

## daemon
docker的基础服务，比如容器的创建、查看、停止、镜像的管理，其实都是由docker的守护进程(daemon)来实现的。

每次执行的docker指令其实都是通过向daemon发送请求来实现的。

daemon的运作（通信模式）主要有两种：
- 一种是通过unix套接字（默认，但只能在本地访问到，比较安全），
- 一种是通过监听tcp协议地址和端口来实现（这个可以实现在远程调用到docker服务）。
## 远程API
除了通过远程tcp协议访问远程主机上的docker服务外，docker还提供了一套基于HTTP的API，可以使用curl来实现操作远程主机上的docker服务，这为开发基于WEB的docker服务提供了便利。

## Docker 集群实现
- Docker Swarm
- Kubernetes

- Fleet
- Mesos

