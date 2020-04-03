# Docker容器数据卷

##  是什么
Docker 容器产生的数据，如果不通过 docker commit 生成新的镜像，使得数据作为镜像的一部分保存下来，那么当容器删除后，数据自然也就没有了。那么如何在 Docker 容器中保证数据持久化呢？我们可以使用 Docker 容器数据卷。

>： Docker 容器数据卷有点类似 Redis 里面的 rdb 和 aof 文件。

容器数据卷就是目录或文件，存在于一个或多个容器中，由 Docker 挂载到容器，但不属于联合文件系统，因此可以绕过联合文件系统提供一些用于持续存储或共享数据的特性。

其设计目的就是数据的持久化，完全独立于容器的生存周期，因此 Docker 不会在容器删除的时候删除其挂载的数据卷

特点：

1. 数据卷可在容器之间共享或重用数据

2. 数据卷中的更改可以直接生效

3. 数据卷中的更改不会包含在镜像的更新中

4. 数据卷的生命周期一直持续到没有容器使用它为止

##  能干嘛
- 容器的持久化
- 容器间继承+共享数据
## 数据卷

### 容器内添加

#### 直接命令添加**
1. 命令
```
docker run -it -v /宿主机绝对路径目录:/容器内目录      镜像名
```   
2. 查看数据卷是否挂载成功
```  
docker inspect 容器ID
```  
3. 容器和宿主机之间数据共享

通过log查看

4. 容器停止退出后，主机修改后数据是否同步
 
**命令(带权限)**
```  
docker run -it -v /宿主机绝对路径目录:/容器内目录:ro 镜像名
```  
#### DockerFile添加
1. 根目录下新建mydocker文件夹并进入
2. 可在Dockerfile中使用`VOLUME`指令来给镜像添加一个或多个数据卷
```
 
VOLUME["/dataVolumeContainer","/dataVolumeContainer2","/dataVolumeContainer3"]
 
```
说明：
- 出于可移植和分享的考虑，用-v 主机目录:容器目录这种方法不能够直接在Dockerfile中实现。
- 由于宿主机目录是依赖于特定宿主机的，并不能够保证在所有的宿主机上都存在这样的特定目录。

3. File构建
```
$ cat dockerfile

# volume test
FROM centos
VOLUME ["/dataVolumeContainer1","/dataVolumeContainer2"]
CMD echo "finished,--------success1"
CMD /bin/bash

```
4. build后生成镜像
```
docker build -f /mydocker/dockerfile2 -t zzyy/centos .
```
获得一个新镜像zzyy/centos
5. run容器
```
docker run -it <images ID> /bin/bash
```
通过上述步骤，容器内的卷目录地址已经知道 对应的主机目录地址哪？？



> Docker挂载主机目录Docker访问出现cannot open directory .: Permission denied

> 解决办法：在挂载目录后多加一个--privileged=true参数即可
##  数据卷容器

###  是什么
命名的容器挂载数据卷，其它容器通过挂载这个(父容器)实现数据共享，挂载数据卷的容器，称之为数据卷容器

###  容器间传递共享(--volumes-from)
 结论：容器之间配置信息的传递，数据卷的生命周期一直持续到没有容器使用它为止