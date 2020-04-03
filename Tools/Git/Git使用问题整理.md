# Git 使用遇到的问题整理
>github https://github.com/git/git 

## 问题
1. efrror: RPC failed; curl 56 OpenSSL SSL_read: error:140943FC:SSL routines:ss
可能是上传大小限制： 

执行如下命令
```
git config http.postBuffer 524288000
```
2. curl 56 OpenSSL SSL_read:SSL_ERROR_sysCALL
```
git config http.sslVerify "false"
```