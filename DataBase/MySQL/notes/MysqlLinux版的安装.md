# MysqlLinux版的安装
> 版本 mysql5.5 rpm安装
## 安装
1. 官网下载地址：http://dev.mysql.com/downloads/mysql/

2. 拷贝&解压缩

3. 检查工作

    - 检查当前系统是否安装过mysql

    执行安装命令前，先执行查询命令`rpm -qa|grep mysql`如果存在mysql-libs的旧版本包如下：执行卸载命令：`rpm -e --nodeps  mysql-libs`
    - 检查/tmp文件夹权限

        由于mysql安装过程中，会通过mysql用户在/tmp目录下新建tmp_db文件，所以请给/tmp较大的权限。执行 ： `chmod -R 777 /tmp`
4. 安装

    在mysql的安装文件目录下执行：
    ```
    rpm -ivh MySQL-client-5.5.54-1.linux2.6.x86_64.rpm
    rpm -ivh MySQL-server-5.5.54-1.linux2.6.x86_64.rpm
    ```
## 操作

1. 查看MySQL安装版本:`mysqladmin --version`

2. mysql服务的启+停
```
$ service mysql start # 启动
$ service mysql stop # 关闭
$ service mysql restart # 重启
```
3. 登录
```
$ mysql # 无密码登录
$ /usr/bin/mysqladmin -u root password 123456 # 设置root账户密码
$ mysql -u root -p # 密码登录
 Enter password:
```
4. MySQL的安装位置

5. 在linux下查看安装目录  
```
ps -ef|grep mysql
```
|参数|	路径|	解释|	备注|
|------|------|------|------|
|--basedir| 	/usr/bin| 	相关命令目录|	mysqladmin mysqldump等命令|
|--datadir|	/var/lib/mysql/ 	|mysql数据库文件的存放路径	 |
|--plugin-dir|	/usr/lib64/mysql/plugin|	mysql插件存放路径|	
|--log-error|	/var/lib/mysql/jack.atguigu.err|	mysql错误日志路径	|
|--pid-file|	/var/lib/mysql/jack.atguigu.pid|	进程pid文件	|
|--socket|	/var/lib/mysql/mysql.sock|	本地连接时用的unix套接字文件|	 
|	|/usr/share/mysql| 	配置文件目录|	mysql脚本及配置文件|
|	|/etc/init.d/mysql|	服务启停相关脚本|	

6. 自启动mysql服务
```
$ chkconfig mysql on  #设置开机自启动
$ chkconfig -list|grep mysql # 查询
$ ntsysv # 
```
7. 修改字符集问题

- 1、查看字符集
```
show variables like 'character%'; 
show variables like '%char%';
```

- 2、修改my.cnf

    在`/usr/share/mysql/` 中找到`my.cnf`的配置文件，拷贝其中的`my-huge.cnf` 到 `/etc/`  并命名为`my.cnf `然后修改my.cnf:
    ```
    [client]
    default-character-set=utf8
    [mysqld]
    character_set_server=utf8
    character_set_client=utf8
    collation-server=utf8_general_ci
    [mysql]
    default-character-set=utf8
    ``` 
 
 - 3、重新启动mysql
    
    > 但是原库的设定不会发生变化，参数修改之对新建的数据库生效
- 4、已生成的库表字符集如何变更

    修改数据库的字符集
    ```
    mysql> alter database mytest character set 'utf8';
    ```
    修改数据表的字符集
    ```
    mysql> alter table user convert to  character set 'utf8';
    ```
    > 但是原有的数据如果是用非'utf8'编码的话，数据本身不会发生改变。