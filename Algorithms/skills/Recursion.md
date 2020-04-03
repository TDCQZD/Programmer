# 递归 - Recursion 

## 概述
在编程语言中，递归的本质是方法自己调用自己
> 递归其实和循环是非常像的，循环都可以改写成递归，递归未必能改写成循环，这是一个充分不必要的条件。

>递归是通过函数体来进行的循环

递归本质就是不断重复相同的事情。而不是去思考完整的调用栈，一级又一级，无从下手,主要关心的主要有三点：
- 返回值
- 调用单元做了什么
- 终止条件

## 递归两个条件：
* 递归出口(终止递归的条件)
* 递归表达式(规律)
> 技巧：在递归中常常是将问题切割成两个部分(1和整体的思想)，这能够让我们快速找到递归表达式(规律)
> 写递归的技巧： 明确一个函数的作用并相信它能完成这个任务，千万不要试图跳进细节
## 代码模板

``` python
def recursion(level, param1, param2, ...):  # level表示当前的递归层数
    # recusrsion terminator 递归终止条件
    if level > MAX_LEVEL:   # MAX_LEVEL终止条件的递归层数
        print_result
        return
    
    # process logic in current 执行level层的业务操作
    process_data(level, data...)
    
    # drill down 调用下一层底层操作
    self.recursion(level + 1, p1, ...)

    # reverse the current level status if needed 如果需要返回level层的状态值
    reverse_state(level)
```
## 适用场景
1. 各种数学问题如：N皇后，汉诺塔，阶乘，迷宫，球和篮子(Google编程大赛)等
2. 各种算法中如：快排，归并排序，二分查找，分治
3. 
## 注意事项
1. 递归处理并不都是高效的，在存在大量重复子问题上低效，比如斐波那契数列和归并排序
## 案列
1. 计算n! n!= 1*2*3*...*n
```python
def factorial(n):
    if n <= 1:
        return 1
    return factorial(n - 1) * n 
```
2. 斐波那契数列 `f(n) = f(n-1)+f(n-2)`
```python
def fib(n):
    if n == 0 or n == 2:
        return n
    return fib(n - 1) + fib(n - 2)
```
> 斐波那契数列使用递归比较低效，使用动态规划DP比较高效

3. 汉诺塔算法
```python
# 递归实现Hanoi塔
def Hanoi(fromPole, withPole, toPole, diskNum):
    if diskNum <= 1:
        print("moving disk from %s to %s" % (fromPole, toPole))
    else:
        Hanoi(fromPole, toPole, withPole, diskNum-1)
        print("moving disk from %s to %s" % (fromPole, toPole))
        Hanoi(withPole, fromPole, toPole, diskNum-1)

Hanoi('A', 'B', 'C', 3)
```
4. 二叉树遍历
```python
# 前序遍历
def PreOrder(rooot):
    if root:
    print(root.Val)
    PreOrder(root.Left)
    PreOrder(rooot.Right)

# 中序遍历
def InOrdor(rooot):
    if root:
    InOrdor(rooot.Left)
    print(root.Val)
    InOrdor(rooot.Right)
# 后序遍历
def PostOrder(rooot):
    if root:
    pre(root.Left)
    pre(root.Right)
    print(root.Val)
```

### 循环改写递归
* 求和:求和1+2+3+4+....+100
* 数组内部的最大值
* 冒泡排序递归写法

