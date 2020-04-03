# Mongo DB异常处理

## 安装中遇到的错误

**错误一：**
```
./mongod: error while loading shared libraries: libcurl.so.4: cannot open shared object file: No such file or directory 
```

原因：在启动./mongod 的时候缺少 `libcurl.so `库

解决方法：
```
apt-get install libcurl4-openssl-dev
```
**错误二：** 

```
./mongod: error while loading shared libraries: libnetsnmpmibs.so.30: cannot open shared object file: No such file or directory
```
原因：缺少 snmpd 相关的东西 。客户端服务端都给它装了。

解决方法：
```
apt-get install snmpd snmp snmp-mibs-downloader
```

