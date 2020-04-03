## 端口
```
netstat -ano                    列出所有端口
netstat -aon |findstr 8080
netstat -aon|findstr "49157"    查看被占用端口对应的PID
tasklist |findstr "9472"
taskkill /f /t /im main.exe     终止进程
```
