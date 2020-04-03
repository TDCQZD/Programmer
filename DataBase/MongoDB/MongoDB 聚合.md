# MongoDB 聚合
MongoDB中聚合(aggregate)主要用于处理数据(诸如统计平均值,求和等)，并返回计算后的数据结果。有点类似sql语句中的 count(*)。

## aggregate() 方法
MongoDB中聚合的方法使用aggregate()。

### 语法
aggregate() 方法的基本语法格式如下所示：
```
>db.COLLECTION_NAME.aggregate(AGGREGATE_OPERATION)
```