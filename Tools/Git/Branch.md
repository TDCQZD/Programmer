# 分支管理 - Branch

## 创建分支
* `git branch` : 查看当前分支
* `git branch <name>` : 创建分支
* `git checkout <name>`:切换到 name分支
* `git checkout -b dev`:创建并切换到dev分支
    - `git checkout`命令加上-b参数表示创建并切换，相当于以下两条命令：
        * `git branch dev`
        * `git checkout dev`
## 合并分支
* `git merge`：命令用于合并指定分支到当前分支
* `git branch -d dev`:删除dev分支

## 解决冲突
- 当Git无法自动合并分支时，就必须首先解决冲突。解决冲突后，再提交，合并完成。

- 解决冲突就是把Git合并失败的文件手动编辑为我们希望的内容，再提交。

- 用git log --graph命令可以看到分支合并图。
    * `git log --graph --pretty=oneline --abbrev-commit`:查看分支的合并情况

## 分支管理策略
通常，合并分支时，如果可能，Git会用Fast forward模式，但这种模式下，删除分支后，会丢掉分支信息。

如果要强制禁用Fast forward模式，Git就会在merge时生成一个新的commit，这样，从分支历史上就可以看出分支信息。

* `git merge --no-ff -m "merge with no-ff" dev`
    - `--no-ff`:禁用Fast forward
    - `-m`:commit描述
### 分支策略
- master分支应该是非常稳定的，也就是仅用来发布新版本
- dev分支是不稳定的,版本发布时，再把dev分支合并到master上，在master分支发布0版本；
## Bug分支
在Git中，每个bug都可以通过一个新的临时分支来修复，修复后，合并分支，然后将临时分支删除。
1. 把当前工作现场“储藏”起来，等以后恢复现场后继续工作
    - `git stash` 
2. 首先确定要在哪个分支上修复bug，假定需要在master分支上修复，就从master创建临时分支：
    - `git checkout master`
    - `git checkout -b issue-101`
3. 修复bug
4. 修复完成后，切换到master分支，并完成合并，最后删除issue-101分支
    - `git checkout master`
    - `git merge --no-ff -m "merged bug fix 101" issue-101`
5. 切回dev分支 
    - `git checkout dev`
    - `git status`
    - `git stash list`:查看之前工作现场
6. 恢复工作现场
    - `git stash apply`:恢复后，stash内容并不删除，你需要用`git stash drop`来删除
    - `git stash pop`:恢复的同时把stash内容删除
    - ` git stash apply stash@{0}`:恢复指定的stash

## Feature分支
每添加一个新功能，最好新建一个feature分支，在上面开发，完成后，合并，最后，删除该feature分支。

如果要丢弃一个没有被合并过的分支，可以通过`git branch -D <name>`强行删除。

## 多人协作
* `git remote`:查看远程库
* `git remote -v`:显示详细的信息

* 查看远程库信息，使用git remote -v；

* 本地新建的分支如果不推送到远程，对其他人就是不可见的；

* 从本地推送分支，使用git push origin branch-name，如果推送失败，先用git pull抓取远程的新提交；

* 在本地创建和远程分支对应的分支，使用git checkout -b branch-name origin/branch-name，本地和远程分支的名称最好一致；

* 建立本地分支和远程分支的关联，使用git branch --set-upstream branch-name origin/branch-name；

* 从远程抓取分支，使用git pull，如果有冲突，要先处理冲突

### 推送分支
推送分支，就是把该分支上的所有本地提交推送到远程库。推送时，要指定本地分支，这样，Git就会把该分支推送到远程库对应的远程分支上.

* master分支是主分支，因此要时刻与远程同步；

* dev分支是开发分支，团队所有成员都需要在上面工作，所以也需要与远程同步；

* bug分支只用于在本地修复bug，就没必要推到远程了，除非老板要看看你每周到底修复了几个bug；

* feature分支是否推到远程，取决于是否同事合作在上面开发。


## Rebase
`git rebase`:
- rebase操作可以把本地未push的分叉提交历史整理成直线；

- rebase的目的是使得我们在查看历史提交的变化时更容易，因为分叉的提交需要三方对比。