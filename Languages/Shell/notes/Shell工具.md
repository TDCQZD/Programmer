#  Shell工具
## cut
cut的工作就是“剪”，具体的说就是在文件中负责剪切数据用的。cut 命令从文件的每一行剪切字节、字符和字段并将这些字节、字符和字段输出。
### 基本用法
```
cut [选项参数]  filename
```
说明：默认分隔符是制表符

**选项参数说明**
* -f	列号，提取第几列
* -d	分隔符，按照指定分隔符分割列

## sed
sed是一种流编辑器，它一次处理一行内容。处理时，把当前处理的行存储在临时缓冲区中，称为“模式空间”，接着用sed命令处理缓冲区中的内容，处理完成后，把缓冲区的内容送往屏幕。接着处理下一行，这样不断重复，直到文件末尾。文件内容并没有改变，除非你使用重定向存储输出。
###	基本用法
```
sed [选项参数]  ‘command’  filename
```
**选项参数说明**

* -e	直接在指令列模式上进行sed的动作编辑。
###	命令功能描述

|命令|	功能描述|
|:------|:------|
|a| 新增，a的后面可以接字串，在下一行出现
|d|	删除
|s|	查找并替换 

## awk
一个强大的文本分析工具，把文件逐行的读入，以空格为默认分隔符将每行切片，切开的部分再进行分析处理。
###	基本用法e
```
awk [选项参数] ‘pattern1{action1}  pattern2{action2}...’ filename
```
- pattern：表示AWK在数据中查找的内容，就是匹配模式
- action：在找到匹配内容时所执行的一系列命令

**选项参数说明**

|选项参数|	功能
|:------|:------|
|-F|	指定输入文件折分隔符
|-v|	赋值一个用户定义变量
###	awk的内置变量

|变量|	说明|
|:------|:------|
|FILENAME|	文件名
|NR|	已读的记录数
|NF|	浏览记录的域的个数（切割后，列的个数）

## sort
sort命令是在Linux里非常有用，它将文件进行排序，并将排序结果标准输出。
###	基本语法
```
sort(选项)(参数)
```
|选项|	说明|
|:------|:------|
|-n|	依照数值的大小排序
|-r|	以相反的顺序来排序
|-t|	设置排序时所用的分隔字符
|-k|	指定需要排序的列
参数：指定待排序的文件列表
