# MySql

## MySQL的CURD操作
### 表操作
```
# 查看表结构
DESCRIBE orderitems;
# 创建表
DESCRIBE customers;
CREATE TABLE customers1
(
cust_id 			INT  			NOT NULL AUTO_INCREMENT, # AUTO_INCREMENT自增
cust_name			CHAR(50)		NOT NULL,
cust_state		    CHAR(5)			NOT NULL,
PRIMARY KEY (cust_id)
) ENGINE = INNODB; #引擎类型 不支持文本搜索

# 更新表
ALTER TABLE vendors
ADD/DROP  vend_phone CHAR(20); # 添加/删除列

# 定义外键
ALTER TABLE orderitems
ADD CONSTRAINT fk_orderitems_orders #外键名称  
FOREIGN KEY(order_num) #KEY
REFERENCES orders(order_num); #关联表

# 删除表
DROP TABLE customers1;

# 重命名表
RENAME TABLE customers1 TO customers2;#重命名一个表
RENAME TABLE customers1 TO customers2,vendors1 TO vendors1,;#重命名多个表
```
### 更新删除数据
```
# 修改数据

UPDATE customers
SET cust_email = 'elemer@fudd.com'
WHERE cust_id = 10005;

# 删除数据
DELETE FROM customers
WHERE cust_id = 10006;
```
### 插入数据
```
# 插入完整的行
INSERT INTO customers
VALUES(NULL,
    'Pep E. LaPew',
    '100 Main street',
    'Los Angeles',
    'CA',
    '90046',
    'USA',
    NUll,
    NULL
    );# 不安全
DESCRIBE customers
INSERT INTO customers(cust_name,
                    cust_address,
                    cust_city,
                    cust_state,
                    cust_zip,
                    cust_country,
                    cust_contact,
                    cust_email
                    )
VALUES('Pep E. LaPew1',
    '100 Main street',
    'Los Angeles',
    'CA',
    '900462',
    'USA',
    NUll,
    NULL
    );# 安全
```
### 数据查询-子查询
- 注意：限制有歧义的列名
```
#单行子查询实操
SELECT last_name,job_id,salary
FROM employees
WHERE job_id = (
        SELECT job_id
        FROM employees
        WHERE employee_id = 141
        )
AND salary > (
        SELECT salary
        FROM employees
        WHERE employee_id = 143
        );

```
### 数据查询
- DISTINCT 数据去重
- ORDER BY 子句排序 :ASC（ascend）: 升序 ,DESC（descend）: 降序
- UNION ALL 操作符返回两个查询的结果集的并集。对于两个结果集的重复部分，不去重。 
- 凡是出现所有，考虑使用外连接，join on
- sql 99 语法：left/right outer join ... on 
- HAVING和WHERE的差别这里有另一种理解方法，WHERE在数据分组前进行过滤，HAVING在数据分组后进行过滤。这是一个重要的区别，WHERE排除的行不包括在分组中。这可能会改变计算值，从而影响HAVING子句中基于这些值过滤掉的分组。
```
# union:求并集数据
# 实现满外连接
SELECT employee_id,last_name,department_name
FROM employees e LEFT JOIN departments d
ON e.`department_id` = d.`department_id`
UNION ALL
SELECT employee_id,last_name,department_name
FROM employees e RIGHT JOIN departments d
ON e.`department_id` = d.`department_id`
WHERE e.`department_id` IS NULL
#实现内连接：
#查询员工的id,name,部门名称
SELECT employee_id,last_name,department_name
FROM employees e JOIN departments d
ON e.`department_id` = d.`department_id`;
#右外连接
SELECT employee_id,last_name,department_name
FROM employees e RIGHT JOIN departments d
ON e.`department_id` = d.`department_id`;
# HAVING和WHERE的区别
SELECT vend_id, COUNT(*) AS num_prods
FROM products
WHERE prod_price >= 10
GROUP BY vend_id
HAVING COUNT(*) >= 2;
```
https://blog.csdn.net/TDCQZD/article/details/82707721
### 笛卡尔集 
笛卡儿积（cartesian product) 由没有联结条件的表关系返回的结果为笛卡儿积。检索出的行的数目将是第一个表中的行数乘以第二个表中的行数。
```
SELECT vend_name, prod_name, prod_price
FROM vendors, products
ORDER BY vend_name,prod_name;# 缺少where,产生笛卡尔集
```
1) 笛卡尔集会在下面条件下产生:
- 省略连接条件
- 连接条件无效
- 所有表中的所有行互相连接
2)为了避免笛卡尔集， 可以在 WHERE 加入有效的连接条件。
## MySQL事务
事务处理（transaction processing）可以用来维护数据库的完整性，它
保证成批的MySQL操作要么完全执行，要么完全不执行。

MyISAM 和 InnoDB 是两种最常使用的引擎。前者不支持明确的事务处理管理，而后者支持

事务处理需要知道的几个术语：
- 事务（ transaction ）指一组SQL语句；
- 回退（ rollback ）指撤销指定SQL语句的过程；
- 提交（ commit ）指将未存储的SQL语句结果写入数据库表；
- 保留点（ savepoint ）指事务处理中设置的临时占位符（place-holder），你可以对它发布回退（与回退整个事务处理不同）。
    - 保留点越多越好 可以在MySQL代码中设置任意多的保留点，越多越好。
        为什么呢？因为保留点越多，你就越能按自己的意愿灵活地进行回退。
    - 释放保留点 
        保留点在事务处理完成（执行一条 ROLLBACK 或COMMIT ）后自动释放。自MySQL 5以来，也可以用 RELEASESAVEPOINT 明确地释放保留点。
### ROLLBACK 
- MySQL的 ROLLBACK 命令用来回退（撤销）MySQL语句
- 事务处理用来管理 INSERT 、 UPDATE 和DELETE 语句。
- 你不能回退 SELECT 语句。（这样做也没有什么意义。）你不能回退 CREATE 或 DROP 操作。事务处理块中可以使用。这两条语句，但如果你执行回退，它们不会被撤销。
```
# ROLLBACK
SELECT *FROM orderitems;
START TRANSACTION;
DELETE FROM orderitems
SELECT *FROM orderitems;
ROLLBACK；
SELECT *FROM orderitems;

```
### commit
- 一般的MySQL语句都是直接针对数据库表执行和编写的。这就是所谓的隐含提交（implicit commit），即提交（写或保存）操作是自动进行的。
- 但是，在事务处理块中，提交不会隐含地进行。为进行明确的提交，使用 COMMIT 语句
```
# COMMIT
START TRANSACTION;
DELETE FROM orders WHERE order_num = 20010;
COMMIT;
```
注意：隐含事务关闭 当 COMMIT 或 ROLLBACK 语句执行后，事务会自
动关闭（将来的更改会隐含提交）。


## MySQL的基本优化策略——查询优化
### 单表使用索引
- 对于单键索引，尽量选择针对当前query过滤性更好的索引
- 在选择组合索引的时候，当前Query中过滤性最好的字段在索引字段顺序中，位置越靠前越好。
- 在选择组合索引的时候，尽量选择可以能够包含当前query中的where字句中更多字段的索引
- 在选择组合索引的时候，如果某个字段可能出现范围查询时，尽量把这个字段放在索引次序的最后面
- 书写sql语句时，尽量避免造成索引失效的情况
### 关联查询优化
1、保证被驱动表的join字段已经被索引
2、left join 时，选择小表作为驱动表，大表作为被驱动表。
3、inner join 时，mysql会自己帮你把小结果集的表选为驱动表。
4、子查询尽量不要放在被驱动表，有可能使用不到索引。
5、能够直接多表关联的尽量直接关联，不用子查询。
### 子查询优化
- 尽量不要使用not in  或者 not exists
    - 用left outer join  on  xxx is null 替代
### 排序分组优化
当范围条件和group by 或者 order by  的字段出现二选一时 ，优先观察条件字段的过滤数量，如果过滤的数据足够多，而需要排序的数据并不多时，优先把索引放在范围字段上。反之，亦然。

**GROUP BY关键字优化**
group by 使用索引的原则几乎跟order by一致 ，唯一区别是groupby 即使没有过滤条件用到索引，也可以直接使用索引。
###  最后使用索引的手段：覆盖索引
什么是覆盖索引？
简单说就是，select 到 from 之间查询的列 <=使用的索引列+主键
## MySQL的索引原理
见MySQL索引.md

## 详述MySQL主从复制原理及配置主从的完整步骤
### 主从复制的原理如下：
主库开启binlog功能并授权从库连接主库，从库通过change master得到主库的相关同步信息,然后连接主库进行验证，主库IO线程根据从库slave线程的请求，从master.info开始记录的位置点向下开始取信息，
同时把取到的位置点和最新的位置与binlog信息一同发给从库IO线程，从库将相关的sql语句存放在relay-log里面，最终从库的sql线程将relay-log里的sql语句应用到从库上，至此整个同步过程完成，之后将是无限重复上述过程
###  完整步骤如下：
1、主库开启binlog功能，并进行全备，将全备文件推送到从库服务器上
2、show master status\G 记录下当前的位置信息及二进制文件名
3、登陆从库恢复全备文件
4、执行change master to 语句
5、执行start slave and show slave status\G

## MySQL引擎机制(InnoDB, Myisam)
对比项   | MyISAM | InnoDB
-------- | ----- | -----		
外键|	不支持|	支持
事务|	不支持|	支持
行表锁|	表锁，即使操作一条记录也会锁住整个表，不适合高并发的操作| 行锁,操作时只锁某一行，不对其它行有影响，适合高并发的操作
缓存|	只缓存索引，不缓存真实数据|	不仅缓存索引还要缓存真实数据，对内存要求较高，而且内存大小对性能有决定性的影响
关注点|	读性能|	并发写、事务、资源
默认安装|	Y|	Y
默认使用|	N |	Y
自带系统表使用|	Y|	N

## mysql的查询流程大致是：
mysql客户端通过协议与mysql服务器建连接，发送查询语句，先检查查询缓存，如果命中，直接返回结果，否则进行语句解析,也就是说，在解析查询之前，服务器会先访问查询缓存(query cache)——它存储SELECT语句以及相应的查询结果集。如果某个查询结果已经位于缓存中，服务器就不会再对查询进行解析、优化、以及执行。它仅仅将缓存中的结果返回给用户即可，这将大大提高系统的性能。
 
语法解析器和预处理：首先mysql通过关键字将SQL语句进行解析，并生成一颗对应的“解析树”。mysql解析器将使用mysql语法规则验证和解析查询；预处理器则根据一些mysql规则进一步检查解析数是否合法。
 
查询优化器当解析树被认为是合法的了，并且由优化器将其转化成执行计划。一条查询可以有很多种执行方式，最后都返回相同的结果。优化器的作用就是找到这其中最好的执行计划。。
 
然后，mysql默认使用的BTREE索引，并且一个大致方向是:无论怎么折腾sql，至少在目前来说，mysql最多只用到表中的一个索引。

## MySQL密码丢了，请找回？
```
 # pkill mysql               #先关闭mysql服务

 #使用--skip-grant-tables启动mysql，忽略授权登陆验证

# mysqld_safe --defaults-file=/etc/my.cnf --skip-grant-tables &

# mysql                   #此时再登陆，已经不需要密码了

mysql> update mysql.user set password=password('abc123') where user='root' and host="localhost";            #设置新的密码

mysql> flush privileges;

# mysql -uroot -pabc123     #再次用新设置的密码登陆即可
```