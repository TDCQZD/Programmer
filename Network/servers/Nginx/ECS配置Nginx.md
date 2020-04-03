# 阿里云ECS搭建nginx服务器

```
$ apt-get install nginx
```
```
$ apt-add-repository ppa:nginx/development
$ apt-get install nginx
```
```
$ sudo /etc/init.d/nginx start # 开启nginx服务器
$ sudo /etc/init.d/nginx stop  # 停止nginx服务器
$ sudo /etc/init.d/nginx restart # 重启
$ sudo /etc/init.d/nginx reload
$ sudo /etc/init.d/nginx force-reload
$ sudo /etc/init.d/nginx upgrade
$ sudo /etc/init.d/nginx status
$ sudo /etc/init.d/nginx configtest
$ sudo /etc/init.d/nginx rotate
```