# 回溯 - Backtrace

## 概述
回溯算法实际上一个类似枚举的搜索尝试过程，主要是在搜索尝试过程中寻找问题的解，当发现已不满足求解条件时，就 “回溯” 返回，尝试别的路径。回溯法是一种选优搜索法，按选优条件向前搜索，以达到目标。但当探索到某一步时，发现原先选择并不优或达不到目标，就退回一步重新选择，这种走不通就退回再走的技术为回溯法，而满足回溯条件的某个状态的点称为 “回溯点”。许多复杂的，规模较大的问题都可以使用回溯法，有“通用解题方法”的美称。

**回溯算法的基本思想是：**从一条路往前走，能进则进，不能进则退回来，换一条路再试。


## 回溯算法框架
解决一个回溯问题，实际上就是一个决策树的遍历过程。你只需要思考 3 个问题：
1. 路径：也就是已经做出的选择。
2. 选择列表：也就是你当前可以做的选择。
3. 结束条件：也就是到达决策树底层，无法再做选择的条件。
## 代码模板
```python
result = []
def backtrack(path, lists):
     #recursion termintor 递归终止条件
    if lists in None:
        result.add(path)
        return
    # processing recursion inside a for loop for 循环里面的递归
    for list in lists:
    
        # make decisions 做选择
        path.add(list)

        # Recursive call 递归调用
        backtrack(path, lists)
        
        # Deselect 撤销选择
        path.remove(list)
```
> 其核心就是 for 循环里面的递归，在递归调用之前「做选择」，在递归调用之后「撤销选择」

## 适用场景
1. 求子集（subset）
2. 求排列（permutation）
3. 求组合（combination）

## 案列
1. 全排列
2. N皇后问题