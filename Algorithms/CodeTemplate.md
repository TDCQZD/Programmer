# Code Template

1. recursion
```python
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

2. divide & conquer
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

3. backtrack
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

4. DP
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

5. DFS
```python
visited = set() # 定义visited集合记录节点的访问状态
def dfs(node, visited):
    visited.add(node)
    
    # process current node here. # 处理当前节点
    
    for next_node in node.children():  # 处理子节点
        if not next_node in visited:
            dfs(next_node, visited)
```

6. BFS 
```python
def bfs(graph, start, end):
    # 定义队列，存储备用节点
    queue = []
    queue.append([start])
    visited.add(start)
     # 判断、处理其余节点的状态
    while queue:
        node = queue.pop()
        visited.add(node)
        
        process(node)
        
        nodes = generate_related_nodes(node)
        queue.push(nodes)
        
    # other processing work
```

7. Binary Search

```python
left, right = 0, len(array)  - 1
while left <=  right;
    //mid =  (left + right) / 2
    mid = left + (right - left) / 2
    if array[mid] == target:
        // find the target
        break or return result
    elif array[mid] < target :
        left = mid + 1
    else:
        right = mid - 1
```
6. Quick Search
```
```