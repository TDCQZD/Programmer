# install
## java
1. 下载并解压JDK
```
tar -zxvf jdk-8u201-linux-x64.tar.gz
mv jdk1.8.0_201/ jdk # 重命名文件夹
```
2. 配置环境变量
```
vim /etc/profile
```

添加如下内容：JAVA_HOME根据实际目录来
```
export JAVA_HOME=/opt/jdk
export CLASSPATH=$JAVA_HOME/lib/
export PATH=$PATH:$JAVA_HOME/bin
```
重启机器或执行命令使编译环境生效
```
source /etc/profile
```
查看安装情况
```
java -version
```
## kafka
* https://www.apache.org/dyn/closer.cgi?path=/kafka/2.2.0/kafka_2.12-2.2.0.tgz
```
> tar -xzf kafka_2.12-2.2.0.tgz
> cd kafka_2.12-2.2.0
```
## zookper
* https://zookeeper.apache.org/doc/r3.4.14/zookeeperStarted.html