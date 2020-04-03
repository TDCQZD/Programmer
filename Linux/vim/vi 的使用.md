# vi 的使用
基本上 vi 共分为三种模式，分别是『一般模式』、『编辑模式』与『指令列命令模式』。 这三种模式的作用分别是：

- 一般模式：
    以 vi 打开一个档案就直接进入一般模式了(这是默认的模式)。在这个模式中， 你可以使用『上下左右』按键来移动光标，你可以使用『删除字符』或『删除整行』来处理档案内容， 也可以使用『复制、贴上』来处理你的文件数据。

- 编辑模式：
    在一般模式中可以进行删除、复制、贴上等等的动作，但是却无法编辑文件内容的！ 要等到你按下『i, I, o, O, a, A, r, R』等任何一个字母之后才会进入编辑模式。注意了！通常在 Linux 中，按下这些按键时，在画面的左下方会出现『 INSERT 或 REPLACE 』的字样，此时才可以进行编辑。而如果要回到一般模式时， 则必须要按下『Esc』这个按键即可退出编辑模式。

- 指令列命令模式：
    在一般模式当中，输入『 : / ? 』三个中的任何一个按钮，就可以将光标移动到最底下那一行。在这个模式当中， 可以提供你『搜寻资料』的动作，而读取、存盘、大量取代字符、离开 vi 、显示行号等等的动作则是在此模式中达成的！
可以将这三个模式想成底下的图标来表示：

![](./img/vi-mode.gif)
> 注意：一般模式可与编辑模式及指令列模式切换， 但编辑模式与指令列模式之间不可互相切换。
## 按键说明
## vim 的暂存档、救援回复与开启时的警告讯息
当我们在使用 vim 编辑时， vim 会在与被编辑的档案的目录下，再建立一个名为 `.filename.swp` 的档案。 比如说我们在上一个小节谈到的编辑 /tmp/`vitest/man.config` 这个档案时， vim 会主动的建立 `/tmp/vitest/.man.config.swp` 的暂存档，你对 `man.config` 做的动作就会被记录到这个 `.man.config.swp` 当中喔！如果你的系统因为某些原因断线了， 导致你编辑的档案还没有储存，这个时候 `.man.config.swp` 就能够发会救援的功能了！

由于 vim 的工作被不正常的中断，导致暂存盘无法藉由正常流程来结束， 所以暂存档就不会消失，而继续保留下来
```
[root@www vitest]# vim man.config
E325: ATTENTION  <==错误代码
Found a swap file by the name ".man.config.swp"  <==底下数行说明有暂存档的存在
          owned by: root   dated: Mon Jan 12 14:48:24 2009
         file name: /tmp/vitest/man.config  <==这个暂存盘属于哪个实际的档案？
          modified: no
         user name: root   host name: www.vbird.tsai
        process ID: 11539
While opening file "man.config"
             dated: Mon Jan 12 13:55:07 2009
底下说明可能发生这个错误的两个主要原因与解决方案！
(1) Another program may be editing the same file.
    If this is the case, be careful not to end up with two
    different instances of the same file when making changes.
    Quit, or continue with caution.

(2) An edit session for this file crashed.
    If this is the case, use ":recover" or "vim -r man.config"
    to recover the changes (see ":help recovery").
    If you did this already, delete the swap file ".man.config.swp"
    to avoid this message.

Swap file ".man.config.swp" already exists!底下说明你可进行的动作
[O]pen Read-Only, (E)dit anyway, (R)ecover, (D)elete it, (Q)uit, (A)bort:  
```
由于暂存盘存在的关系，因此 vim 会主动的判断你的这个档案可能有些问题，在上面的图示中 vim 提示两点主要的问题与解决方案，分别是这样的：

**问题一：可能有其他人或程序同时在编辑这个档案：**

由于 Linux 是多人多任务的环境，因此很可能有很多人同时在编辑同一个档案。如果在多人共同编辑的情况下， 万一大家同时储存，那么这个档案的内容将会变的乱七八糟！为了避免这个问题，因此 vim 会出现这个警告窗口！ 解决的方法则是：

- 找到另外那个程序或人员，请他将该 vim 的工作结束，然后你再继续处理。 

- 如果你只是要看该档案的内容并不会有任何修改编辑的行为，那么可以选择开启成为只读(O)档案， 亦即上述画面反白部分输入英文『 o 』即可，其实就是 [O]pen Read-Only 的选项啦！

**问题二：在前一个 vim 的环境中，可能因为某些不知名原因导致 vim 中断 (crashed)：**

这就是常见的不正常结束 vim 产生的后果。解决方案依据不同的情况而不同喔！常见的处理方法为：

- 如果你之前的 vim 处理动作尚未储存，此时你应该要按下『R』，亦即使用 (R)ecover 的项目， 此时 vim 会载入 .man.config.swp 的内容，让你自己来决定要不要储存！这样就能够救回来你之前未储存的工作。 不过那个 .man.config.swp 并不会在你结束 vim 后自动删除，所以你离开 vim 后还得要自行删除 .man.config.swp 才能避免每次打开这个档案都会出现这样的警告！

- 如果你确定这个暂存盘是没有用的，那么你可以直接按下『D』删除掉这个暂存盘，亦即 (D)elete it 这个项目即可。 此时 vim 会载入 man.config ，并且将旧的 .man.config.swp 删除后，建立这次会使用的新的 .man.config.swp 喔！

至于这个发现暂存盘警告讯息的画面中，有出现六个可用按钮，各按钮的说明如下：

* [O]pen Read-Only：打开此档案成为只读档， 可以用在你只是想要查阅该档案内容并不想要进行编辑行为时。一般来说，在上课时，如果你是登入到同学的计算机去看他的配置文件， 结果发现其实同学他自己也在编辑时，可以使用这个模式；

* (E)dit anyway：还是用正常的方式打开你要编辑的那个档案， 并不会载入暂存档的内容。不过很容易出现两个使用者互相改变对方的档案等问题！不好不好！

* (R)ecover：就是加载暂存盘的内容，用在你要救回之前未储存的工作。 不过当你救回来并且储存离开 vim 后，还是要手动自行删除那个暂存档喔！

* (D)elete it：你确定那个暂存档是无用的！那么开启档案前会先将这个暂存盘删除！ 这个动作其实是比较常做的！因为你可能不确定这个暂存档是怎么来的，所以就删除掉他吧！哈哈！

* (Q)uit：按下 q 就离开 vim ，不会进行任何动作回到命令提示字符。

* (A)bort：忽略这个编辑行为，感觉上与 quit 非常类似！ 也会送你回到命令提示字符就是啰！