# BFS 

宽度优先搜索算法（又称广度优先搜索）是最简便的图的搜索算法之一，这一算法也是很多重要的图的算法的原型。Dijkstra单源最短路径算法和Prim最小生成树算法都采用了和宽度优先搜索类似的思想。其别名又叫BFS，属于一种盲目搜寻法，目的是系统地展开并检查图中的所有节点，以找寻结果。换句话说，它并不考虑结果的可能位置，彻底地搜索整张图，直到找到结果为止。

## 概述
广度优先搜索算法（Breadth-First-Search，缩写为 BFS），是一种利用队列实现的搜索算法.

## 代码模板
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
    ...
```

## 适用场景

## 案列



## BFS & DFS区别
* BFS 的重点在于队列，而 DFS 的重点在于递归。这是它们的本质区别