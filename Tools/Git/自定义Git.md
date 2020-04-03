# 自定义Git

## Git显示颜色
- `git config --global color.ui true`
## 忽略特殊文件
在Git工作区的根目录下创建一个特殊的.gitignore文件，然后把要忽略的文件名填进去，Git就会自动忽略这些文件。
> 所有配置文件可以直接在线浏览：https://github.com/github/gitignore


忽略文件的原则是：
- 忽略操作系统自动生成的文件，比如缩略图等；
- 忽略编译生成的中间文件、可执行文件等，也就是如果一个文件是通过另一个文件自动生成的，那自动生成的文件就没必要放进版本库，比如Java编译产生的.class文件；
- 忽略你自己的带有敏感信息的配置文件，比如存放口令的配置文件。

总结：
- 忽略某些文件时，需要编写.gitignore；

- .gitignore文件本身要放到版本库里，并且可以对.gitignore做版本管理
    * `git check-ignore`
## 配置别名
`git config --global alias.st status`:
- `--global`参数是全局参数
- 配置Git的时候，加上--global是针对当前用户起作用的，如果不加，那只针对当前的仓库起作用。
> 每个仓库的Git配置文件都放在.git/config文件中

> 别名就在[alias]后面，要删除别名，直接把对应的行删掉即可

> 当前用户的Git配置文件放在用户主目录下的一个隐藏文件.gitconfig中
## 搭建Git服务器
1. 安装git:` $ sudo apt-get install git`
2. 创建一个git用户，用来运行git服务:` $ sudo adduser git`
3. 创建证书登录：
    收集所有需要登录的用户的公钥，就是他们自己的id_rsa.pub文件，把所有公钥导入到`/home/git/.ssh/authorized_keys`文件里，一行一个。

4. 初始化Git仓库：

    先选定一个目录作为Git仓库，假定是/srv/sample.git，在/srv目录下输入命令：
    ```
    $ sudo git init --bare sample.git
    ```
    Git就会创建一个裸仓库，裸仓库没有工作区，因为服务器上的Git仓库纯粹是为了共享，所以不让用户直接登录到服务器上去改工作区，并且服务器上的Git仓库通常都以.git结尾。然后，把owner改为git：
    ```
    $ sudo chown -R git:git sample.git
    ```
5. 禁用shell登录：

    出于安全考虑，第二步创建的git用户不允许登录shell，这可以通过编辑/etc/passwd文件完成。找到类似下面的一行：
    ```
    git:x:1001:1001:,,,:/home/git:/bin/bash
    ```
    改为：
    ```
    git:x:1001:1001:,,,:/home/git:/usr/bin/git-shell
    ```
    这样，git用户可以正常通过ssh使用git，但无法登录shell，因为我们为git用户指定的git-shell每次一登录就自动退出。

6. 克隆远程仓库：

    现在，可以通过git clone命令克隆远程仓库了，在各自的电脑上运行：
    ```
    $ git clone git@server:/srv/sample.git
    Cloning into 'sample'...
    warning: You appear to have cloned an empty repository.
    ```

> 要方便管理公钥，用[Gitosis](https://github.com/sitaramc/gitolite)；

> 要像SVN那样变态地控制权限，用[Gitolite](https://github.com/sitaramc/gitolite)。