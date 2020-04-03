# Docker
Docker 是一个开源的应用容器引擎，基于 Go 语言 并遵从Apache2.0协议开源。

Docker 可以让开发者打包他们的应用以及依赖包到一个轻量级、可移植的容器中，然后发布到任何流行的 Linux 机器上，也可以实现虚拟化。

容器是完全使用沙箱机制，相互之间不会有任何接口（类似 iPhone 的 app）,更重要的是容器性能开销极低。
## Docker的应用场景
- Web 应用的自动化打包和发布。

- 自动化测试和持续集成、发布。

- 在服务型环境中部署和调整数据库或其他的后台应用。

- 从头编译或者扩展现有的OpenShift或Cloud Foundry平台来搭建自己的PaaS环境。
## Docker 的优点
1. 简化程序：

    Docker 让开发者可以打包他们的应用以及依赖包到一个可移植的容器中，然后发布到任何流行的 Linux 机器上，便可以实现虚拟化。Docker改变了虚拟化的方式，使开发者可以直接将自己的成果放入Docker中进行管理。方便快捷已经是 Docker的最大优势，过去需要用数天乃至数周的	任务，在Docker容器的处理下，只需要数秒就能完成。

2. 避免选择恐惧症：

    如果你有选择恐惧症，还是资深患者。Docker 帮你打包你的纠结！比如 Docker 镜像；Docker 镜像中包含了运行环境和配置，所以 Docker 可以简化部署多种应用实例工作。比如 Web 应用、后台应用、数据库应用、大数据应用比如 Hadoop 集群、消息队列等等都可以打包成一个镜像部署。

3. 节省开支：

    一方面，云计算时代到来，使开发者不必为了追求效果而配置高额的硬件，Docker 改变了高性能必然高价格的思维定势。Docker 与云的结合，让云空间得到更充分的利用。不仅解决了硬件管理的问题，也改变了虚拟化的方式。
## Docker 架构
Docker 使用客户端-服务器 (C/S) 架构模式，使用远程API来管理和创建Docker容器。

Docker 容器通过 Docker 镜像来创建。

容器与镜像的关系类似于面向对象编程中的对象与类。

|Docker|	面向对象|
|------|------|
|容器	|对象|
|镜像	|类|

![](./notes/images/docker1.png)

|Docker工具|描述|
|------|------|
|Docker 镜像(Images)|Docker 镜像是用于创建 Docker 容器的模板。|
|Docker 容器(Container)|容器是独立运行的一个或一组应用。
|Docker 客户端(Client)|Docker 客户端通过命令行或者其他工具使用 Docker API (https://docs.docker.com/reference/api/docker_remote_api) 与 Docker 的守护进程通信。
|Docker 主机(Host)|一个物理或者虚拟的机器用于执行 Docker 守护进程和容器。
|Docker 仓库(Registry)|Docker 仓库用来保存镜像，可以理解为代码控制中的代码仓库。Docker Hub(https://hub.docker.com)| 提供了庞大的镜像集合供使用。
|Docker Machine|Docker Machine是一个简化Docker安装的命令行工具，通过一个简单的命令行即可在相应的平台上安装Docker，比如VirtualBox、 Digital Ocean、Microsoft Azure。
## Docker 工具-三剑客
- docker-machine

    docker-machine就是docker公司官方提出的，用于在各种平台上快速创建具有docker服务的虚拟机的技术，甚至可以通过指定driver来定制虚拟机的实现原理（一般是virtualbox）。
- docker-compose

    dcoker-compose技术，就是通过一个.yml配置文件，将所有的容器的部署方法、文件映射、容器连接等等一系列的配置写在一个配置文件里，最后只需要执行`docker-compose up`命令就会像执行脚本一样的去一个个安装容器并自动部署他们。
- docker-swarm

    swarm是基于docker平台实现的集群技术，可以通过几条简单的指令快速的创建一个docker集群，接着在集群的共享网络上部署应用，最终实现分布式的服务。




## 参考资料
* [docker官网]：http://www.docker.com
* [docker中文网站]：https://www.docker-cn.com/
* [Docker Hub官网] https://hub.docker.com/
* https://www.runoob.com/docker/docker-resources.html
