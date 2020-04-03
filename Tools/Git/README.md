# 代码管理工具

## Git
Git 是一个开源的分布式版本控制系统，用于敏捷高效地处理任何或小或大的项目。

### 使用指南
* 安装教程
* 仓库管理 - repository
    - 本地仓库
    - 远程仓库
* 分支管理 - Branch
* 标签管理 - tag
* GitHub
* 自定义Git
### 参考资料
* https://www.liaoxuefeng.com/wiki/896043488029600
* https://www.runoob.com/git/git-tutorial.html
* [Git 完整命令手册](https://git-scm.com/docs)
* [Git 简明指南](https://www.runoob.com/manual/git-guide/)


## SVN

## Git 与 SVN 区别
Git 不仅仅是个版本控制系统，它也是个内容管理系统(CMS)，工作管理系统等。

Git 与 SVN 区别点：

1. Git 是分布式的，SVN 不是：这是 Git 和其它非分布式的版本控制系统，例如 SVN，CVS 等，最核心的区别。

2. Git 把内容按元数据方式存储，而 SVN 是按文件：所有的资源控制系统都是把文件的元信息隐藏在一个类似 .svn、.cvs 等的文件夹里。

3. Git 分支和 SVN 的分支不同：分支在 SVN 中一点都不特别，其实它就是版本库中的另外一个目录。

4. Git 没有一个全局的版本号，而 SVN 有：目前为止这是跟 SVN 相比 Git 缺少的最大的一个特征。

5. Git 的内容完整性要优于 SVN：Git 的内容存储使用的是 SHA-1 哈希算法。这能确保代码内容的完整性，确保在遇到磁盘故障和网络问题时降低对版本库的破坏。
