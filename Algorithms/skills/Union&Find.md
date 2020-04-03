# 并查集 - union & find

## 概述
计算机科学中，并查集是一种树型的数据结构，用于处理一些不交集（Disjoint Sets）的合并及查询问题。有一个联合-查找算法（Union-find Algorithm）定义了两个用于此数据结构的操作：
* Find：确定元素属于哪一个子集。它可以被用来确定两个元素是否属于同一子集。
* Union：将两个子集合并成同一个集合。
* 由于支持这两种操作，一个不相交集也常被称为联合-查找数据结构（Union-find Data Structure）或合并-查找集合（Merge-find Set）。

为了更加精确的定义这些方法，需要定义如何表示集合。一种常用的策略是为每个集合选定一个固定的元素，称为代表，以表示整个集合。接着，Find(x)返回 x所属集合的代表，而 Union 使用两个集合的代表作为参数。

## 代码模板
```JavaScript
function MakeSet(x)
    x.parent := x

function Find(x)
    if x.parent == x
        return x
    else 
        return Find(x.parent)

function Union(x, y)
    xRoot := Find(x)
    yRoot := Find(y)
    xRoot.parent := yRoot 
```
并查集优化一：平衡性优化
```JavaScript

function MakeSet(x)
    x.parent := x
    x.rank := 0

function Union(x, y)
    xRoot := Find(x)
    yRoot := Find(y)
    if xRoot == yRoot
     return
    
    
    // x 和 y 不在同一个集合，合并他们
    if xRoot.rank < yRoot.rank
        xRoot.parent := yRoot
    else if xRoot.rank > yRoot.rank
        yRoot.parent := xRoot
    else
        yRoot.parent := xRoot
        xRoot.rank := xRoot.rank + 1
```
并查集优化二：java实现路径压缩
```java
public class QuickUnionUF{
    private int[] roots;

    public QuickUnionUF(int N){
        roots = new int[N];
        for (int i = 0; i < N; i++){
            root[i] = i;
        }
    }

    private int findRoot(int i){
        int root = i;
        while (root != roots[root])
            root = roots[root];
        while i != roots[i]{
            int tmp = roots[i];
            roots[i] = root;
            i = tmp;
        }
        return root;
    }

    public boolean connected(int p, int q){
        return findRoot(p) == findRoot(q);
    }

    public void union(int p, int q){
        int qroot = findRoot(q);
        int proot = findRoot(p);
        roots[proot] == qroot;
    }
}
```
## 适用场景
1. 并查集算法主要是解决图论中「动态连通性」问题的

## 案列
1. 很多使用 DFS 深度优先算法解决的问题，也可以用 Union-Find 算法解决
