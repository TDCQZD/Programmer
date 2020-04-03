# history
## History命令语法: 
```
$ history [n]
$ history [-c]
$ history [-raw] histfiles
```
参数:
* n   :数字,要列出最近的 n 笔命令列表
* -c  :将目前的shell中的所有 history 内容全部消除
* -a  :将目前新增的history 指令新增入 histfiles 中,若没有加 histfiles ,则预设写入 ~/.bash_history
* -r  :将 histfiles 的内容读到目前这个 shell 的 history 记忆中
* -w  :将目前的 history 记忆内容写入 histfiles
## 常用history命令

- 只列出最近10条记录:
```
$ history 10 (注,history和10中间有空格)
```
- 使用命令记录号码执行命令,执行历史清单中的第99条命令
```
$!99 (!和99中间没有空格)
```
- 重复执行上一个命令
```
$!!
```
- 执行最后一次以rpm开头的命令(!?  ?代表的是字符串,这个String可以随便输,Shell会从最后一条历史命令向前搜索,最先匹配的一条命令将会得到执行。)
```
$!rpm
```
- 逐屏列出所有的历史记录:
```
$ history | more
```
- 立即清空history当前所有历史命令的记录
```
$ history -c
```
除了使用history命令,在 shell 或 GUI 终端提示下,你也可以使用上下方向键来翻阅命令历史(向下箭头会向前翻阅),直到你找到所需命令为止。