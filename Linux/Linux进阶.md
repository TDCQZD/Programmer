# 面试题
1.  分布式

把整个系统拆分，积分、订单、库存、支付系统，把不同的服务部署不同的机器上，对外提供统一的整体服务。 就是分布式

2. 集群

扩容：纵向（自身扩展，8G——> 16G,性能有限） 横向（多个自身,1个8G -> 多个8G,集群）
在不同的机器部署相同的服务
3. 负载均衡

路由算法，平均分配

## 面试：
1. 查看后台服务器
``` 
ps -ef|grep mysql
netstat -anp|grep 3306 
lsof -i:3306
ss -tnl|grep 3306
```
```
 ps -ef|grep -n mysql # 位置
 ps -ef|grep mysql|grep -v grep #筛选,从小到大
 ps -ef|grep mysql|grep -v grep|sort -nk2 #筛选,从小到大
 ps -ef|grep mysql|grep -v grep|sort -nrk2 #筛选,从大到小
 ps -ef|grep mysql|grep -v grep|wc -l # 筛选,统计
 ps -ef|grep mysql|grep -v grep|sort -nrk2|head -n 1  # 极限值
 ps -ef|grep mysql|grep -v grep|sort -nrk2|head -n 1|awk '{print$1$2$3$4}' # 极限值
 ps -ef|grep mysql|grep -v grep|sort -nrk2|head -n 1|awk '{print$2}' # 极限值PID
```
``` 删除
ps -ef|grep mysql|grep -v grep|sort -nrk2|head -n 1|awk '{print$2}'|xargs kill -9
```
2. find 和 grep区别

find：查找搜索 解决有无的问题
grep：逐行扫描  解决有多少的问题

3. `|`

管道符, 上一个命令集结果集作为下一个数据集
