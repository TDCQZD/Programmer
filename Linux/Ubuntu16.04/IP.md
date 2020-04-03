# IP

## 设置静态IP
1. `vi /etc/network/interfaces`

添加内容：
```
auto eth0
iface eth0 inet static
address 192.168.8.100    
netmask 255.255.255.0
gateway 192.168.8.2
dns-nameserver 119.29.29.29
```
> ifocnfig 查看网卡名称
2. 重启网络：sudo /etc/init.d/networking restart