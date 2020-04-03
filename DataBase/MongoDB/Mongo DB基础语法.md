## 数据库 (database)

**创建数据库**
```
use DATABASE_NAME
```
如果数据库不存在，则创建数据库，否则切换到指定数据库。

示例：
```
> use runoob
switched to db runoob
> db
runoob
> 
```
**查看所有数据库**
```
> show dbs
admin   0.000GB
config  0.000GB
local   0.000GB
> 
```

刚创建的数据库 runoob 并不在数据库的列表中， 要显示它，我们需要向 runoob 数据库插入一些数据。

```
> db.product.insert({"name":"iphone","category":"cellphone","value":"$8000"})
WriteResult({ "nInserted" : 1 })
> show dbs
admin   0.000GB
config  0.000GB
ebay    0.000GB
local   0.000GB
>
```

MongoDB 中默认的数据库为 test，如果你没有创建新的数据库，集合将存放在 test 数据库中。

**注意: 在 MongoDB 中，集合只有在内容插入后才会创建! 就是说，创建集合(数据表)后要再插入一个文档(记录)，集合才会真正创建。**

**删除数据库**
```
db.dropDatabase()
```
删除当前数据库，默认为 test，你可以使用 db 命令查看当前数据库名。

## 集合
### 查看已存在的集合
```
show collections
show tables
```
### 创建集合
```
db.createCollection(name, options)
```
参数说明：
- name: 要创建的集合名称
- options: 可选参数, 指定有关内存大小及索引的选项

options 可以是如下参数：
字段 | 类型 | 描述 
--------|-----|-------		
capped|	布尔|	（可选）如果为 true，则创建固定集合。固定集合是指有着固定大小的集合，当达到最大值时，它会自动覆盖最早的文档。当该值为 true 时，必须指定 size 参数。
autoIndexId|	布尔|	（可选）如为 true，自动在 _id 字段创建索引。默认为 false。
size|	数值|	（可选）为固定集合指定一个最大值（以字节计）。如果 capped 为 true，也需要指定该字段。
max	|数值|	（可选）指定固定集合中包含文档的最大数量。

在插入文档时，MongoDB 首先检查固定集合的 size 字段，然后检查 max 字段。

```
> db.createCollection("mycol", { capped : true, autoIndexId : true, size : 6142800, max : 10000 } )
```



创建固定集合 mycol，整个集合空间大小 6142800 KB, 文档最大个数为 10000 个。
```
> db.createCollection("mycol", { capped : true, autoIndexId : true, size : 6142800, max : 10000 } )
{ "ok" : 1 }
>
```
> 在 MongoDB 中，通常不需要专门创建集合；当你插入一些文档时，MongoDB 会自动创建集合。
### 删除集合

语法格式：
```
db.COLLECTION_NAME.drop()
```
参数说明：无

返回值:如果成功删除选定集合，则 drop() 方法返回 true，否则返回 false。


## 文档

### 插入文档
MongoDB 使用 insert() 或 save() 方法向集合中插入文档，语法如下：

```
db.COLLECTION_NAME.insert(document)
db.COLLECTION_NAME.save(document)
```
这里的 document 是一个文档对象，如：
> document=({name: 'iPhone', category: 'cellphone', value: 5000 });


如果该集合不在该数据库中， MongoDB 会自动创建该集合并插入文档。

`db.COLLECTION_NAME.save(document)` 命令。如果不指定 _id 字段 save() 方法类似于 insert() 方法。如果指定 _id 字段，则会更新该 _id 的数据。


**3.2 版本后还有以下几种语法可用于插入文档:**

 `db.collection.insertOne()`:向指定集合中插入一条文档数据
 `db.collection.insertMany()`:向指定集合中插入多条文档数据

```
> db.person.insert({"name":"Bob","age":"21","city":"Beijing"})
WriteResult({ "nInserted" : 1 })
> db.person.save({"name":"Alice","age":"25","city":"Shanghai"})
WriteResult({ "nInserted" : 1 })
> db.person.insertOne({"name":"Cary","age":"23","city":"Shenzhen"})
{
	"acknowledged" : true,
	"insertedId" : ObjectId("5bec01ce246d8d2ee94bdb21")
}
> db.person.insertMany([{"name":"Jack","age":"13","city":"杭州"},{"name":"Amry","age":"23","city":"广州"}])
{
	"acknowledged" : true,
	"insertedIds" : [
		ObjectId("5bec02c9246d8d2ee94bdb22"),
		ObjectId("5bec02c9246d8d2ee94bdb23")
	]
}
> 

```
### 查询文档——基础

MongoDB 查询数据的语法格式如下：
```
db.collection.find(query, projection)
```
参数说明：
- query ：可选，使用查询操作符指定查询条件
- projection ：可选，使用投影操作符指定返回的键。查询时返回文档中所有键值， 只需省略该参数即可（默认省略）。

如果你需要以易读的方式来读取数据，可以使用 pretty() 方法，语法格式如下：
```
>db.collection.find().pretty()
```
除了 find() 方法之外，还有一个 `findOne()` 方法，它只返回一个文档。
```
db.person.findOne()
```
> 注意不能使用`db.person.findOne().pretty()`,TypeError: db.person.findOne(...).pretty is not a function 


示例：
```
> db.person.findOne()
{
	"_id" : ObjectId("5bec014f246d8d2ee94bdb1e"),
	"name" : "Bob",
	"age" : "21",
	"city" : "Beijing"
}

> db.person.find()
{ "_id" : ObjectId("5bec014f246d8d2ee94bdb1e"), "name" : "Bob", "age" : "21", "city" : "Beijing" }
{ "_id" : ObjectId("5bec0182246d8d2ee94bdb1f"), "name" : "Alice", "age" : "25", "city" : "Shanghai" }
{ "_id" : ObjectId("5bec01ce246d8d2ee94bdb21"), "name" : "Cary", "age" : "23", "city" : "Shenzhen" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb22"), "name" : "Jack", "age" : "13", "city" : "杭州" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb23"), "name" : "Amry", "age" : "23", "city" : "广州" }
> db.person.find().pretty()
{
	"_id" : ObjectId("5bec014f246d8d2ee94bdb1e"),
	"name" : "Bob",
	"age" : "21",
	"city" : "Beijing"
}
{
	"_id" : ObjectId("5bec0182246d8d2ee94bdb1f"),
	"name" : "Alice",
	"age" : "25",
	"city" : "Shanghai"
}
{
	"_id" : ObjectId("5bec01ce246d8d2ee94bdb21"),
	"name" : "Cary",
	"age" : "23",
	"city" : "Shenzhen"
}
{
	"_id" : ObjectId("5bec02c9246d8d2ee94bdb22"),
	"name" : "Jack",
	"age" : "13",
	"city" : "杭州"
}
{
	"_id" : ObjectId("5bec02c9246d8d2ee94bdb23"),
	"name" : "Amry",
	"age" : "23",
	"city" : "广州"
}
> 

```

**指定 projection**

若不指定 projection，则默认返回所有键，指定 projection 格式如下，有两种模式:
```
db.collection.find(query, {title: 1, by: 1}) // inclusion模式 指定返回的键，不返回其他键
db.collection.find(query, {title: 0, by: 0}) // exclusion模式 指定不返回的键,返回其他键
```
1.  _id 键默认返回，需要主动指定 _id:0 才会隐藏

2. 只能全1或全0，除了在inclusion模式时可以指定_id为0
    ```
    db.collection.find(query, {_id:0, title: 1, by: 1}) 
    ```
示例：
```
> db.person.find({"name":"Bob"}, {city: 1, by: 1}) 
{ "_id" : ObjectId("5bec014f246d8d2ee94bdb1e"), "city" : "Beijing" }
> db.person.find({"name":"Bob"}, {city: 0, by: 0})
{ "_id" : ObjectId("5bec014f246d8d2ee94bdb1e"), "name" : "Bob", "age" : "21" }
> db.person.find({"name":"Bob"}, {_id:0, city: 1, by: 1})
{ "city" : "Beijing" }
>
```

### 更新文档

MongoDB 使用 update() 和 save() 方法来更新集合中的文档

**update() 方法**

update() 方法用于更新已存在的文档。语法格式如下：
```
db.collection.update(
   <query>,
   <update>,
   {
     upsert: <boolean>,
     multi: <boolean>,
     writeConcern: <document>
   }
)
```
参数说明：

- query : update的查询条件，类似sql update查询内where后面的。
- update : update的对象和一些更新的操作符（如$,$inc...）等，也可以理解为sql update查询内set后面的
- upsert : 可选，这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入。
- multi : 可选，mongodb 默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新。
- writeConcern :可选，抛出异常的级别。

示例：
```
> db.person.find({"age":"23"})
{ "_id" : ObjectId("5bec01ce246d8d2ee94bdb21"), "name" : "Cary", "age" : "23", "city" : "Shenzhen" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb23"), "name" : "Amry", "age" : "23", "city" : "广州" }
> db.person.update({"age":"23"},{$set:{"city": "北京"}})
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
> db.person.find({"age":"23"})
{ "_id" : ObjectId("5bec01ce246d8d2ee94bdb21"), "name" : "Cary", "age" : "23", "city" : "北京" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb23"), "name" : "Amry", "age" : "23", "city" : "广州" }
> db.person.update({"age":"23"},{$set:{"city": "北京"}},{multi:true})
WriteResult({ "nMatched" : 2, "nUpserted" : 0, "nModified" : 1 })
> db.person.find({"age":"23"})
{ "_id" : ObjectId("5bec01ce246d8d2ee94bdb21"), "name" : "Cary", "age" : "23", "city" : "北京" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb23"), "name" : "Amry", "age" : "23", "city" : "北京" }
> 
```

**save() 方法**
save() 方法通过传入的文档来替换已有文档。语法格式如下：
```
db.collection.save(
   <document>,
   {
     writeConcern: <document>
   }
)
```
参数说明：
- document : 文档数据。
- writeConcern :可选，抛出异常的级别。
    ```
    WriteConcern.NONE:没有异常抛出
    WriteConcern.NORMAL:仅抛出网络错误异常，没有服务器错误异常
    WriteConcern.SAFE:抛出网络错误异常、服务器错误异常；并等待服务器完成写操作。
    WriteConcern.MAJORITY: 抛出网络错误异常、服务器错误异常；并等待一个主服务器完成写操作。
    WriteConcern.FSYNC_SAFE: 抛出网络错误异常、服务器错误异常；写操作等待服务器将数据刷新到磁盘。
    WriteConcern.JOURNAL_SAFE:抛出网络错误异常、服务器错误异常；写操作等待服务器提交到磁盘的日志文件。
    WriteConcern.REPLICAS_SAFE:抛出网络错误异常、服务器错误异常；等待至少2台服务器完成写操作。
    ```
示例：
```
> db.person.find({"name" : "Cary"}) 
{ "_id" : ObjectId("5bec01ce246d8d2ee94bdb21"), "name" : "Cary", "age" : "23", "city" : "北京" }
> db.person.save({ "_id" : ObjectId("5bec01ce246d8d2ee94bdb21"), "name" : "Cary", "age" : "33", "city" : "芜湖" })
WriteResult({ "nMatched" : 1, "nUpserted" : 0, "nModified" : 1 })
> db.person.find({"name" : "Cary"})
{ "_id" : ObjectId("5bec01ce246d8d2ee94bdb21"), "name" : "Cary", "age" : "33", "city" : "芜湖" }

```
**在3.2版本开始，MongoDB提供以下更新集合文档的方法：**
```
db.collection.updateOne() 向指定集合更新单个文档
db.collection.updateMany() 向指定集合更新多个文档
```
示例：
```
# 更新单个文档
> db.person.find({"name" : "Alice"}) 
{ "_id" : ObjectId("5bec0182246d8d2ee94bdb1f"), "name" : "Alice", "age" : "25", "city" : "Shanghai" }
> db.person.updateOne({"name":"Alice"},{$set:{"city":"上海"}}) 
{ "acknowledged" : true, "matchedCount" : 1, "modifiedCount" : 1 }
> db.person.find({"name" : "Alice"}) 
{ "_id" : ObjectId("5bec0182246d8d2ee94bdb1f"), "name" : "Alice", "age" : "25", "city" : "上海" }
# 更新多个文档
> db.person.find()
{ "_id" : ObjectId("5bec014f246d8d2ee94bdb1e"), "name" : "Bob", "age" : "21", "city" : "Beijing" }
{ "_id" : ObjectId("5bec0182246d8d2ee94bdb1f"), "name" : "Alice", "age" : "25", "city" : "上海" }
{ "_id" : ObjectId("5bec01ce246d8d2ee94bdb21"), "name" : "Cary", "age" : "33", "city" : "芜湖" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb22"), "name" : "Jack", "age" : "13", "city" : "杭州" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb23"), "name" : "Amry", "age" : "23", "city" : "北京" }
> 
> db.person.updateMany({"age":{$lt:"22"}},{$set:{"city":"北京"}})
{ "acknowledged" : true, "matchedCount" : 2, "modifiedCount" : 2 }
> db.person.find()
{ "_id" : ObjectId("5bec014f246d8d2ee94bdb1e"), "name" : "Bob", "age" : "21", "city" : "北京" }
{ "_id" : ObjectId("5bec0182246d8d2ee94bdb1f"), "name" : "Alice", "age" : "25", "city" : "上海" }
{ "_id" : ObjectId("5bec01ce246d8d2ee94bdb21"), "name" : "Cary", "age" : "33", "city" : "芜湖" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb22"), "name" : "Jack", "age" : "13", "city" : "北京" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb23"), "name" : "Amry", "age" : "23", "city" : "北京" }
> 

```
### 删除文档

在执行remove()函数前先执行find()命令来判断执行的条件是否正确，这是一个比较好的习惯。

**remove() 方法的基本语法格式如下所示：**
```
db.collection.remove(
   <query>,
   <justOne>
)
```
**如果MongoDB 是 2.6 版本以后的，语法格式如下：**
```
db.collection.remove(
   <query>,
   {
     justOne: <boolean>,
     writeConcern: <document>
   }
)
```
参数说明：

- query :（可选）删除的文档的条件。
- justOne : （可选）如果设为 true 或 1，则只删除一个文档。
- writeConcern :（可选）抛出异常的级别。

**删除第一条找到的记录可以设置 justOne 为 1**
```
>db.COLLECTION_NAME.remove(DELETION_CRITERIA,1)
```
示例：
```
> db.person.find()
{ "_id" : ObjectId("5bec014f246d8d2ee94bdb1e"), "name" : "Bob", "age" : "21", "city" : "北京" }
{ "_id" : ObjectId("5bec0182246d8d2ee94bdb1f"), "name" : "Alice", "age" : "25", "city" : "上海" }
{ "_id" : ObjectId("5bec01ce246d8d2ee94bdb21"), "name" : "Cary", "age" : "33", "city" : "芜湖" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb22"), "name" : "Jack", "age" : "13", "city" : "北京" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb23"), "name" : "Amry", "age" : "23", "city" : "北京" }
> db.person.remove({"city" : "北京"},1)
WriteResult({ "nRemoved" : 1 })
> db.person.find()
{ "_id" : ObjectId("5bec0182246d8d2ee94bdb1f"), "name" : "Alice", "age" : "25", "city" : "上海" }
{ "_id" : ObjectId("5bec01ce246d8d2ee94bdb21"), "name" : "Cary", "age" : "33", "city" : "芜湖" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb22"), "name" : "Jack", "age" : "13", "city" : "北京" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb23"), "name" : "Amry", "age" : "23", "city" : "北京" }
> 

```
**删除所有数据**
```
> db.COLLECTION_NAME.remove({})
```
示例：
```
> db.product.find()
{ "_id" : ObjectId("5bebf992246d8d2ee94bdb1a"), "name" : "iphone", "category" : "cellphone", "value" : "$8000" }
> db.product.remove({})
WriteResult({ "nRemoved" : 1 })
> db.product.find()
> 
```

**现在官方推荐使用 deleteOne() 和 deleteMany() 方法:**

删除集合下全部文档：
```
db.collection.deleteMany({})
```
删除 status 等于 A 的全部文档：
```
db.collection.deleteMany({ status : "A" })
```
删除 status 等于 D 的一个文档：
```
db.collection.deleteOne( { status: "D" } )
```
示例：
```
> db.person.find()
{ "_id" : ObjectId("5bec0182246d8d2ee94bdb1f"), "name" : "Alice", "age" : "25", "city" : "上海" }
{ "_id" : ObjectId("5bec01ce246d8d2ee94bdb21"), "name" : "Cary", "age" : "33", "city" : "芜湖" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb22"), "name" : "Jack", "age" : "13", "city" : "北京" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb23"), "name" : "Amry", "age" : "23", "city" : "北京" }
>  db.person.deleteOne({"city" : "芜湖"})
{ "acknowledged" : true, "deletedCount" : 1 }
> db.person.find()
{ "_id" : ObjectId("5bec0182246d8d2ee94bdb1f"), "name" : "Alice", "age" : "25", "city" : "上海" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb22"), "name" : "Jack", "age" : "13", "city" : "北京" }
{ "_id" : ObjectId("5bec02c9246d8d2ee94bdb23"), "name" : "Amry", "age" : "23", "city" : "北京" }
> db.person.deleteMany({"city" : "北京" })
{ "acknowledged" : true, "deletedCount" : 2 }
> db.person.find()
{ "_id" : ObjectId("5bec0182246d8d2ee94bdb1f"), "name" : "Alice", "age" : "25", "city" : "上海" }
>  db.person.deleteMany({})
{ "acknowledged" : true, "deletedCount" : 1 }
> db.person.find()
> 
```
### 查询文档——进阶

MongoDB 与 RDBMS Where 语句比较：
操作 | 格式 | 范例 |RDBMS中的类似语句
--------|-----|-----|-------			
等于|	{<key>:<value>} |	db.col.find({"name":"Bob"}).pretty()	|where name = 'Bob'
小于|	{<key>:{$lt:<value>}}|	db.col.find({"likes":{$lt:50}}).pretty()|	where likes < 50
小于或等于|	{<key>:{$lte:<value>}}|	db.col.find({"likes":{$lte:50}}).pretty()|	where likes <= 50
大于|	{<key>:{$gt:<value>}}|	db.col.find({"likes":{$gt:50}}).pretty()|	where likes > 50
大于或等于|	{<key>:{$gte:<value>}}|	db.col.find({"likes":{$gte:50}}).pretty()|	where likes >= 50
不等于|	{<key>:{$ne:<value>}}|	db.col.find({"likes":{$ne:50}}).pretty()|	where likes != 50

**MongoDB AND 条件** 
MongoDB 的 find() 方法可以传入多个键(key)，每个键(key)以逗号隔开，即常规 SQL 的 AND 条件。

语法格式如下：
```
>db.collection.find({key1:value1, key2:value2}).pretty()
```
示例：
```
> db.person.find({"name":"Alice","age":"25"}).pretty()
{
	"_id" : ObjectId("5bec0182246d8d2ee94bdb1f"),
	"name" : "Alice",
	"age" : "25",
	"city" : "Shanghai"
}
> 
```
**MongoDB OR 条件**
MongoDB OR 条件语句使用了关键字 $or,语法格式如下：
```
>db.collection.find(
   {
      $or: [
         {key1: value1}, {key2:value2}
      ]
   }
).pretty()
```
示例：
```
> db.person.find({$or: [{"age":"25"}, {"age":"21"}]}).pretty()
{
	"_id" : ObjectId("5bec014f246d8d2ee94bdb1e"),
	"name" : "Bob",
	"age" : "21",
	"city" : "Beijing"
}
{
	"_id" : ObjectId("5bec0182246d8d2ee94bdb1f"),
	"name" : "Alice",
	"age" : "25",
	"city" : "Shanghai"
}
>
```
**AND 和 OR 联合使用**
示例：
```
> db.person.find({"age": {$lt:"25"}, $or: [{"city": "Beijing"},{"city": "杭州"}]}).pretty()
{
	"_id" : ObjectId("5bec014f246d8d2ee94bdb1e"),
	"name" : "Bob",
	"age" : "21",
	"city" : "Beijing"
}
{
	"_id" : ObjectId("5bec02c9246d8d2ee94bdb22"),
	"name" : "Jack",
	"age" : "13",
	"city" : "杭州"
}

```

## $type 操作符
$type操作符是基于BSON类型来检索集合中匹配的数据类型，并返回结果。
类型 | 数字 | 备注 |
--------|-----|-------		
Double|	1|
String|	2|	 
Object|	3|	 
Array|	4|	 
Binary data|	5|	 
Undefined|	6|	已废弃。
Object id|	7|	 
Boolean|	8|	 
Date|	9|	 
Null|	10|	 
Regular Expression|	11|	 
JavaScript|	13|	 
Symbol|	14|	 
JavaScript (with scope)|	15|	 
32-bit integer|	16|	 
Timestamp|	17|	 
64-bit integer|	18|	 
Min key|	255	|Query with -1.|
Max key|	127|
MongoDB 中可以使用的类型如下表所示：
如果想获取 "col" 集合中 title 为 String 的数据，你可以使用以下命令：
```
db.col.find({"title" : {$type : 2}})
db.col.find({"title" : {$type : 'string'}})
```

## 排序
在 MongoDB 中使用 sort() 方法对数据进行排序，sort() 方法可以通过参数指定排序的字段，并使用 1 和 -1 来指定排序的方式，其中 1 为升序排列，而 -1 是用于降序排列。

**语法**

sort()方法基本语法如下所示：
```
>db.COLLECTION_NAME.find().sort({KEY:1})
```
示例：
```
> db.person.find()
{ "_id" : ObjectId("5bec1a0bb5e34da10f176049"), "name" : "Bob", "age" : "21", "city" : "Beijing" }
{ "_id" : ObjectId("5bec1a13b5e34da10f17604a"), "name" : "Alice", "age" : "25", "city" : "Shanghai" }
{ "_id" : ObjectId("5bec1a1ab5e34da10f17604b"), "name" : "Cary", "age" : "23", "city" : "Shenzhen" }
{ "_id" : ObjectId("5bec1a20b5e34da10f17604c"), "name" : "Jack", "age" : "13", "city" : "杭州" }
{ "_id" : ObjectId("5bec1a20b5e34da10f17604d"), "name" : "Amry", "age" : "23", "city" : "广州" }
> db.person.find({},{"age":1,_id:0}).sort({"age":-1})
{ "age" : "25" }
{ "age" : "23" }
{ "age" : "23" }
{ "age" : "21" }
{ "age" : "13" }
> db.person.find({},{"age":1,_id:0}).sort({"age":1})
{ "age" : "13" }
{ "age" : "21" }
{ "age" : "23" }
{ "age" : "23" }
{ "age" : "25" }
> 

```
### Limit() 方法

在MongoDB中读取指定数量的数据记录，可以使用MongoDB的Limit方法，limit()方法接受一个数字参数，该参数指定从MongoDB中读取的记录条数。

语法
limit()方法基本语法如下所示：
```
>db.COLLECTION_NAME.find().limit(NUMBER)
```
示例：
```
> db.person.find()
{ "_id" : ObjectId("5bec1a0bb5e34da10f176049"), "name" : "Bob", "age" : "21", "city" : "Beijing" }
{ "_id" : ObjectId("5bec1a13b5e34da10f17604a"), "name" : "Alice", "age" : "25", "city" : "Shanghai" }
{ "_id" : ObjectId("5bec1a1ab5e34da10f17604b"), "name" : "Cary", "age" : "23", "city" : "Shenzhen" }
{ "_id" : ObjectId("5bec1a20b5e34da10f17604c"), "name" : "Jack", "age" : "13", "city" : "杭州" }
{ "_id" : ObjectId("5bec1a20b5e34da10f17604d"), "name" : "Amry", "age" : "23", "city" : "广州" }
>  db.person.find().limit(3)
{ "_id" : ObjectId("5bec1a0bb5e34da10f176049"), "name" : "Bob", "age" : "21", "city" : "Beijing" }
{ "_id" : ObjectId("5bec1a13b5e34da10f17604a"), "name" : "Alice", "age" : "25", "city" : "Shanghai" }
{ "_id" : ObjectId("5bec1a1ab5e34da10f17604b"), "name" : "Cary", "age" : "23", "city" : "Shenzhen" }

```
### Skip() 方法
skip()方法来跳过指定数量的数据，skip方法同样接受一个数字参数作为跳过的记录条数。

语法
skip() 方法脚本语法格式如下：
```
>db.COLLECTION_NAME.find().limit(NUMBER).skip(NUMBER)
```
> 注:skip()方法默认参数为 0 。
示例：
```
> db.person.find()
{ "_id" : ObjectId("5bec1a0bb5e34da10f176049"), "name" : "Bob", "age" : "21", "city" : "Beijing" }
{ "_id" : ObjectId("5bec1a13b5e34da10f17604a"), "name" : "Alice", "age" : "25", "city" : "Shanghai" }
{ "_id" : ObjectId("5bec1a1ab5e34da10f17604b"), "name" : "Cary", "age" : "23", "city" : "Shenzhen" }
{ "_id" : ObjectId("5bec1a20b5e34da10f17604c"), "name" : "Jack", "age" : "13", "city" : "杭州" }
{ "_id" : ObjectId("5bec1a20b5e34da10f17604d"), "name" : "Amry", "age" : "23", "city" : "广州" }
>  db.person.find().limit(3).skip(2)
{ "_id" : ObjectId("5bec1a1ab5e34da10f17604b"), "name" : "Cary", "age" : "23", "city" : "Shenzhen" }
{ "_id" : ObjectId("5bec1a20b5e34da10f17604c"), "name" : "Jack", "age" : "13", "city" : "杭州" }
{ "_id" : ObjectId("5bec1a20b5e34da10f17604d"), "name" : "Amry", "age" : "23", "city" : "广州" }
> 

```
> skip(), limilt(), sort()三个放在一起执行的时候，执行的顺序是先 sort(), 然后是 skip()，最后是显示的 limit()。
## 索引
MongoDB使用 createIndex() 方法来创建索引。

> 注意在 3.0.0 版本前创建索引方法为 db.collection.ensureIndex()，之后的版本使用了 db.collection.createIndex() 方法，ensureIndex() 还能用，但只是 createIndex() 的别名。

**语法**
createIndex()方法基本语法格式如下所示：
```
>db.collection.createIndex(keys, options)
```
语法中 Key 值为你要创建的索引字段，1 为指定按升序创建索引，如果你想按降序来创建索引指定为 -1 即可.

示例：
```
>db.col.createIndex({"title":1})
>
```
createIndex() 方法中你也可以设置使用多个字段创建索引（关系型数据库中称作复合索引）。
```
>db.col.createIndex({"title":1,"description":-1})
>
```
createIndex() 接收可选参数，可选参数列表如下：
Parameter | Type | Description |
--------|-----|-------			
background|	Boolean|	建索引过程会阻塞其它数据库操作，background可指定以后台方式创建索引，即增加 "background" 可选参数。 "background" 默认值为false。|
unique|	Boolean|	建立的索引是否唯一。指定为true创建唯一索引。默认值为false.|
name|	string|	索引的名称。如果未指定，MongoDB的通过连接索引的字段名和排序顺序生成一个索引名称。|
dropDups|	Boolean|	3.0+版本已废弃。在建立唯一索引时是否删除重复记录,指定 true 创建唯一索引。默认值为 false.|
sparse|	Boolean|	对文档中不存在的字段数据不启用索引；这个参数需要特别注意，如果设置为true的话，在索引字段中不会查询出不包含对应字段的文档.。默认值为 false.
expireAfterSeconds|	integer|	指定一个以秒为单位的数值，完成 TTL设定，设定集合的生存时间。
v|	index version|	索引的版本号。默认的索引版本取决于mongod创建索引时运行的版本。
weights|	document|	索引权重值，数值在 1 到 99,999 之间，表示该索引相对于其他索引字段的得分权重。
default_language|	string|	对于文本索引，该参数决定了停用词及词干和词器的规则的列表。 默认为英语
language_override|	string|	对于文本索引，该参数指定了包含在文档中的字段名，语言覆盖默认的language，默认值为 language.

## MongoDB 备份(mongodump)与恢复(mongorestore)

### MongoDB数据备份

在Mongodb中我们使用mongodump命令来备份MongoDB数据。该命令可以导出所有数据到指定目录中。

mongodump命令可以通过参数指定导出的数据量级转存的服务器。

**语法**
mongodump命令脚本语法如下：

```
>mongodump -h dbhost -d dbname -o dbdirectory

```
参数说明：
- -h：
    MongDB所在服务器地址，例如：127.0.0.1，当然也可以指定端口号：127.0.0.1:27017

- -d：
    需要备份的数据库实例，例如：test

- -o：
    备份的数据存放位置，例如：c:\data\dump，当然该目录需要提前建立，在备份完成后，系统自动在dump目录下建立一个test目录，这个目录里面存放该数据库实例的备份数据。

示例：
```
root@Aws:~# mongodump -h 127.0.0.1 -d ebay -o ~
2018-11-14T21:08:01.725+0800	writing ebay.person to 
2018-11-14T21:08:01.725+0800	writing ebay.product to 
2018-11-14T21:08:01.726+0800	done dumping ebay.person (5 documents)
2018-11-14T21:08:01.726+0800	done dumping ebay.product (0 documents)
root@Aws:~# ls
ebay  ethereum  goproject  software
root@Aws:~# 
```
### MongoDB 恢复(mongorestore)

mongodb使用 mongorestore 命令来恢复备份的数据。

**语法**

mongorestore命令脚本语法如下：
```
>mongorestore -h <hostname><:port> -d dbname <path>
```
- --host <:port>, -h <:port>：
    MongoDB所在服务器地址，默认为： localhost:27017

- --db , -d ：
    需要恢复的数据库实例，例如：test，当然这个名称也可以和备份时候的不一样，比如test2

- --drop：
    恢复的时候，先删除当前数据，然后恢复备份的数据。就是说，恢复后，备份后添加修改的数据都会被删除，慎用哦！

- <path>：
    mongorestore 最后的一个参数，设置备份数据所在位置，例如：c:\data\dump\test。

    你不能同时指定 <path> 和 --dir 选项，--dir也可以设置备份目录。

- --dir：
    指定备份的目录

    你不能同时指定 <path> 和 --dir 选项。