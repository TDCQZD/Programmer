# git patch
Git 打补丁

## patch 和diff 的区别
Git 提供了两种补丁方案，一是用git diff生成的UNIX标准补丁.diff文件，二是git format-patch生成的Git专用.patch 文件。
- .diff文件只是记录文件改变的内容，不带有commit记录信息,多个commit可以合并成一个diff文件。
- .patch文件带有记录文件改变的内容，也带有commit记录信息,每个commit对应一个patch文件。

## patch使用  
### git patch
1. 创建patch 文件的常用命令行
**`*`某次提交（含）之前的几次提交：**
```
git format-patch 【commit sha1 id】-n
```
n指从sha1 id对应的commit开始算起n个提交。
eg：
```
git format-patch  2a2fb4539925bfa4a141fe492d9828d030f7c8a8 -2
```
**`*`某个提交的patch：**
```
git format-patch 【commit sha1 id】 -1
```
eg：
```
git format-patch  2a2fb4539925bfa4a141fe492d9828d030f7c8a8 -1
```
**`*`某两次提交之间的所有patch:**
```
git format-patch 【commit sha1 id】..【commit sha1 id】
``` 
eg：
```
git format-patch  2a2fb4539925bfa4a141fe492d9828d030f7c8a8..89aebfcc73bdac8054be1a242598610d8ed5f3c8
```
2. 应用patch
**检查patch是否能正常打入:**
```
git apply --check 【path/to/xxx.patch】
```
**打入patch:**
```
git apply 【path/to/xxx.patch】

```
或者
```
git  am 【path/to/xxx.patch】
```

### git diff
1. 创建diff文件
```
git diff  【commit sha1 id】 【commit sha1 id】 >  【diff文件名】
```
eg：
```
git diff  2a2fb4539925bfa4a141fe492d9828d030f7c8a8  89aebfcc73bdac8054be1a242598610d8ed5f3c8 > patch.diff
```

2. 应用diff
**检查diff是否能正常打入:**
```
git apply --check 【path/to/xxx.diff】
```
**打入diff:**
```
git apply 【path/to/xxx.diff】
```

### url.patch
1. 通过pull requests 获取patch 文件[路径后面加`.patch`生成patch]
```
https://github.com/nknorg/nkn/pull/426
https://github.com/nknorg/nkn/pull/426.patch  
https://patch-diff.githubusercontent.com/raw/nknorg/nkn/pull/426.patch
```
2. 下载patch文件
``` 
curl https://patch-diff.githubusercontent.com/raw/nknorg/nkn/pull/427.patch >> 427.patch
```
3. 使用patch 文件
```
git am 427.patch
git am *.patch
```

## 冲突解决
