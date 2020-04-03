# 动态规划 - Dynamic Programming

## 概述
动态规划（英语：Dynamic programming，简称 DP）是一种在数学、管理科学、计算机科学、经济学和生物信息学中使用的，通过把原问题分解为相对简单的子问题的方式求解复杂问题的方法。

动态规划常常适用于有重叠子问题和最优子结构性质的问题，动态规划方法所耗时间往往远少于朴素解法。

动态规划背后的基本思想非常简单。大致上，若要解一个给定问题，我们需要解其不同部分（即子问题），再根据子问题的解以得出原问题的解。动态规划往往用于优化递归问题，例如斐波那契数列，如果运用递归的方式来求解会重复计算很多相同的子问题，利用动态规划的思想可以减少计算量。

通常许多子问题非常相似，为此动态规划法试图仅仅解决每个子问题一次，具有天然剪枝的功能，从而减少计算量：一旦某个给定子问题的解已经算出，则将其记忆化存储，以便下次需要同一个子问题解之时直接查表。这种做法在重复子问题的数目关于输入的规模呈指数增长时特别有用。

## DP解法思路：
1. 递推: 递归+记忆化 —> 递推
2. 状态的定义：`opt[n], dp[n], fib[n]` （通常定义为数组）
3. 状态转移方程：`opt[n] = best_of(opt[n-1], opt[n-2], …)`
4. 最优子结构

## 代码模板
```python
# 状态定义
dp = new int[m+1][n-1]

# 初始状态
dp[0][0] = x
dp[0][1] = y

# DP状态推导
for i = 0; i <=n; ++i {
    for j = 0; j <=m; ++j {
        ....
        dp[i][j] = min { dp[i-1][j],dp[i][j-1], etc..}
    }  
}
# 最优先
return dp[m][n]
```
## 适用场景

## 案列
1. 斐波那契数列列
```java
int fib(int n) {
    return n <= 1 ? n : fib(n - 1) + fib(n - 2);
}
```
2. 翻书问题或者走台阶问题
- 共有n个台阶，每次只能上1个台阶或者2个台阶，共有多少种方法爬完台阶。
- 共有n页书，每次只能翻1页或者2页书，共有多少种方法翻完全书。
```python
# f(n)为翻完全书的方法
# 递归写法
# 时间复杂度为指数级
def f(n):
    if n == 1:
        return 1
    if n == 2:
        return 2
    if n > 2:
        return f(n - 1) + f(n - 2)

# 迭代写法，或者叫循环写法
# 时间复杂度为O(N)
# res数组缓存了之前计算的结果
def f(n):
    res = [0 for i in range(n + 1)]
    res[1] = 1
    res[2] = 2
    for i in range(3, n+1):
        res[i] = res[i - 1] + res[i - 2]
    return res[n]


# 使用缓存
# 时间复杂度为O(N)
cache = {}
def fib(n):
    if n not in cache.keys():
        cache[n] = _fib(n)
    return cache[n]

def _fib(n):
    if n == 1 or n == 2:
        return n
    else:
        return fib(n-1) + fib(n-2)
# 使用两个变量
# 时间复杂度为O(N)，空间复杂度为O(1)
def climbStaris(self, n):
    x, y = 1, 1
    for _ in range(1, n):
        x, y = y, x + y
    return y 
```
## 动态规划 vs 回溯 vs 贪心算法
- 回溯(递归)  — 重复计算
- 贪心算法    — 永远局部最优
- 动态规划    — 记录局部最优子结构 / 多种记录值
## 参考资料
* [动态编程](http://www.baike.com/wiki/%E5%8A%A8%E6%80%81%E7%BC%96%E7%A8%8B)