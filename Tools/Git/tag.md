# 标签管理 - tag

## 创建标签
1. 切换到需要打标签的分支上
    - `git branch`
    - `git checkout master`
2. 打标签
    -  `git tag v1.0`
3. 查看所有标签
    - `git tag`
> 默认标签是打在最新提交的commit上的

### 创建带有说明的标签

` git tag -a v0.1 -m "version 0.1 released" 1094adb`
    - `-a`指定标签名
    - `-m`指定说明文字

### 补打标签
1. 找到历史提交的commit id
2. 打标签
    - `git tag v0.9 f52c633`
3. 查看标签信息
    - `git show <tagname>`

## 操作标签
### 删除
* 删除本地标签:`git tag -d v0.1`
> 因为创建的标签都只存储在本地，不会自动推送到远程。所以，打错的标签可以在本地安全删除
* 删除远程标签
    1. 如果标签已经推送到远程，要删除远程标签就麻烦一点，先从本地删除：
        - `git tag -d v0.9`

    2. 然后，从远程删除。删除命令也是push，但是格式如下：
        -  `git push origin :refs/tags/v0.9`

### 推送
* 推送某个标签到远程:`git push origin <tagname>`
* 一次性推送全部尚未推送到远程的本地标签:`git push origin --tags`
