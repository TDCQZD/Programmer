# Nginx
## Nginx 简介
Nginx（发音同engine x）是一个异步框架的 Web 服务器，也可以用作反向代理，负载平衡器 和 HTTP 缓存。该软件由 Igor Sysoev 创建，并于2004年首次公开发布。同名公司成立于2011年，以提供支持。Nginx 是一款免费的开源软件，根据类 BSD 许可证的条款发布。一大部分Web服务器使用 Nginx ，通常作为负载均衡器。[

## Nginx的特点 
- 更快：

    - 单次请求会得到更快的响应。
    - 在高并发环境下，Nginx 比其他 WEB 服务器有更快的响应。

- 高扩展性：

    Nginx 是基于模块化设计，由多个耦合度极低的模块组成，因此具有很高的扩展性。许多高流量的网站都倾向于开发符合自己业务特性的定制模块。

- 高可靠性：

    Nginx 的可靠性来自于其核心框架代码的优秀设计，模块设计的简单性。另外，官方提供的常用模块都非常稳定，每个 worker 进程相对独立，master 进程在一个 worker 进程出错时可以快速拉起新的 worker 子进程提供服务。

- 低内存消耗：

    - 一般情况下，10000个非活跃的 HTTP Keep-Alive 连接在 Nginx 中仅消耗 2.5MB 的内存，这是 Nginx 支持高并发连接的基础。
    - 单机支持10万以上的并发连接：理论上，Nginx 支持的并发连接上限取决于内存，10万远未封顶。

- 热部署:

    master 进程与 worker 进程的分离设计，使得 Nginx 能够提供热部署功能，即在 7x24 小时不间断服务的前提下，升级 Nginx 的可执行文件。当然，它也支持不停止服务就更新配置项，更换日志文件等功能。
- 最自由的 BSD 许可协议:

    这是 Nginx 可以快速发展的强大动力。BSD 许可协议不只是允许用户免费使用 Nginx ，它还允许用户在自己的项目中直接使用或修改 Nginx 源码，然后发布。




## 参考资料
* https://xuexb.github.io/learn-nginx/
* [nginx的负载均衡如何配置] https://blog.csdn.net/Rodgexue/article/details/79976610
* https://www.w3cschool.cn/nginx/
* 
