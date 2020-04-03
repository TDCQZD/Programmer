# 仓库 - repository

## 仓库操作
* 初始化一个Git仓库: `git init`
    - 查看隐藏目录：`ls -ah` 

* 添加文件到Git仓库:
    - `git add <file>`:提交的所有修改放到暂存区（Stage）
    - `git commit -m <message>`:一次性把暂存区的所有修改提交到分支

* 添加远程库: `git remote add origin <github repository address>`
    - 关联url位置：`.git/config:url`
* 推送master分支内容:`git push -u origin master`

* 克隆远程仓库 `git clone <github repository address>`

## 版本回退 `git reset --hard <commit_id>`
* 历史记录:
    - `git log`:显示从最近到最远的提交日志:  
    - `git log --pretty=oneline` 
* `git reset --hard <parameter>`:把当前版本回退到上一个版本 
    - HEAD :表示当前版本
    - HEAD^ :上一个版本
    - HEAD^^:上上一个版本
    - HEAD~100：往上100个版本
* `git reset --hard <commit_id>`:回到指定的某个版本
* `git reflog`:记录每一次命令
## 修改操作
* `git status`:查看状态
* `git checkout -- file`:丢弃工作区file的修改
* `git reset HEAD <file>`:把暂存区的修改撤销掉（unstage），重新放回工作区 
* `git reset`:命令既可以回退版本，也可以把暂存区的修改回退到工作区
* `git rm <file>`:从版本库中删除该文件
* `git checkout -- <file>`:是用版本库里的版本替换工作区的版本
    >  注意：从来没有被添加到版本库就被删除的文件，是无法恢复的！
