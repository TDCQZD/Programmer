# 分治 - Divde & Conquer

## 概述

分治方法将问题划分为互不相交的子问题，递归的求解子问题，再将它们的解组合起来，求出原问题的解

分治算法可以分三步走：分解 -> 解决 -> 合并:
1. 分解原问题为结构相同的子问题。
2. 分解到某个容易求解的边界之后，进行递归求解。
3. 将子问题的解合并成原问题的解。

## 代码模板
```python
def divide_conquer(problem, param1, param2, ...):
    #recursion termintor 递归终止条件
    if problem in None:
        print_resut
        return
    # prepare data 准备数据，同时把当前的大问题拆分成子问题
    data = prepare_data(problem)
    subproblems = split_problem(problem, data)

    # conquer subproblems 处理不同的子问题
    subresult1 = self.divide_conquer(subproblems[0], p1, ...)
    subresult2 = self.divide_conquer(subproblems[1], p1, ...)
    subresult3 = self.divide_conquer(subproblems[2], p1, ...)
    ...

    # process and generate the final result 合并所有子结果
    result = process_result(subresult1, subresult2, subresult3, ...)
```
## 适用场景
## 案列
1. 归并排序
```python
def MergeSort(lst):   
    # 递归终止条件
    if len(lst) <= 1:
        return lst

    # 处理不同的子问题
    def merge(left, right):
        merged = []
        
        while left and right:
            merged.append(left.pop(0) if left[0] <= right[0] else right.pop(0))
        
        while left:
            merged.append(left.pop(0))
        
        while right:
            merged.append(right.pop(0))
        
        return merged

    # 把当前的大问题拆分成子问题    
    middle = int(len(lst) / 2) 
    left = MergeSort(lst[:middle])
    right = MergeSort(lst[middle:])
    return merge(left, right)
```