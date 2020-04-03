# DFS
## 概述
深度优先搜索算法（英语：Depth-First-Search，DFS）是一种用于遍历或搜索树或图的算法。沿着树的深度遍历树的节点，尽可能深的搜索树的分支。当节点v的所在边都己被探寻过，搜索将回溯到发现节点v的那条边的起始节点。这一过程一直进行到已发现从源节点可达的所有节点为止。如果还存在未被发现的节点，则选择其中一个作为源节点并重复以上过程，整个进程反复进行直到所有节点都被访问为止。属于盲目搜索。

深度优先搜索是图论中的经典算法，利用深度优先搜索算法可以产生目标图的相应拓扑排序表，利用拓扑排序表可以方便的解决很多相关的图论问题，如最大路径问题等等。

## 代码模板
DFS代码-递归写法：
```python 
visited = set() # 定义visited集合记录节点的访问状态
def dfs(node, visited):
    visited.add(node)
    
    # process current node here. # 处理当前节点
    process(node)

    # 处理子节点
    for next_node in node.children():  
        if not next_node in visited:
            dfs(next_node, visited)
```
DFS代码-非递归写法：
```python
def DFS(self, tree):
    if tree.root is None:
        return []

    visited, stack = [], [tree.root]

    while stack:
        node = stack.pop()
        visited.add(node)

        process(node)
        nodes = generate_related_node(node)
        stack.push(nodes)
        
    # other processing work
    ...
```
## 适用场景
## 案列
## BFS & DFS区别
* BFS 的重点在于队列，而 DFS 的重点在于递归。这是它们的本质区别
