# 项目系统架构

## 微服务
* [go-micro 微服务开发中文手册] https://githgoub.com/lixiangyun/micro_doc_ZH_CN
## RESTful架构
* [理解RESTful架构] http://www.ruanyifeng.com/blog/2011/09/restful.html
* [RESTful API 设计指南] http://www.ruanyifeng.com/blog/2014/05/restful_api.html
* [基于Go语言构建企业级的RESTful-API服务]https://github.com/confucianzuoyuan/go-tutorials/tree/master/%E5%9F%BA%E4%BA%8EGo%E8%AF%AD%E8%A8%80%E6%9E%84%E5%BB%BA%E4%BC%81%E4%B8%9A%E7%BA%A7%E7%9A%84RESTful-API%E6%9C%8D%E5%8A%A1
## 高并发架构
* [亿级商品详情页架构演进技术解密] https://www.jianshu.com/p/d8ade6430bc3
* [设计Ticket Master] htt ps://mp.weixin.qq.com/s?__biz=MzU2OTUyNzk1NQ==&mid=2247491120&amp;idx=1&amp;sn=40ef45e108992d16942864998e2ff5e8&source=41#wechat_redirect

### 高并发考虑的问题
最重要是有高并发场景（大公司未必有这种场景），在这种场景下，如何提高吞吐量？如何降低响应时间？

1. 对频繁操作的表行，要加上乐观锁，防止写丢失。mysql不支持行锁，需要在程序中加入。hibernate自带version字段。

2. 表设计上尽量采用分库，减少一台数据库在高并发环境下，性能低下。

3. 事务加到service层，可以保证同时回滚。

4. 数据库中某些字段可以确定唯一的情况下，最好加入唯一索引，防止数据插入重复。对查询的字段加上索引。

5. 对方法的并发操作，加上synchronized.
## 支付系统
## 有高负载、高并发开发设计
## 如何突破程序性能瓶颈？

## 系统安全考虑的问题？
 - 预防CSRF攻击

- 确保输入过滤

- 避免XSS攻击

- 避免SQL注入

- 存储密码

- 加密和解密数据

## 分布式+集群简介

- 分布式：不同的多台服务器上面部署不同的服务模块（工程），他们之间通过Rpc/Rmi之间通信和调用，对外提供服务和组内协作。
 
- 集群：不同的多台服务器上面部署相同的服务模块，通过分布式调度软件进行统一的调度，对外提供服务和访问。


### 分布式系统（distributed system）
 由多台计算机和通信的软件组件通过计算机网络连接（本地网络或广域网）组成。分布式系统是建立在网络之上的软件系统。正是因为软件的特性，所以分布式系统具有高度的内聚性和透明性。因此，网络和分布式系统之间的区别更多的在于高层软件（特别是操作系统），而不是硬件。分布式系统可以应用在在不同的平台上如：Pc、工作站、局域网和广域网上等。
