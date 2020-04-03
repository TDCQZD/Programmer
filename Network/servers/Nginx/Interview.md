# Interview
## 1、nginx的负载均衡如何配置
在nginx里面配置一个upstream，然后把相关的服务器ip都配置进去。然后采用轮询的方案，然后在nginx里面的配置项里，proxy-pass指向这个upstream，这样就能实现负载均衡。
> https://blog.csdn.net/Rodgexue/article/details/79976610


