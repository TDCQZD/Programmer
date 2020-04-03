# Elasticsearch 
Elasticsearch 是一个建立在全文搜索引擎 Apache Lucene(TM) 基础上的搜索引擎，可以说 Lucene 是当今最先进，最高效的全功能开源搜索引擎框架。

Elasticsearch是基于Apache Lucene的搜索服务器。它由Shay Banon开发并于2010年发布。现在是由Elasticsearch BV负责维护。

Elasticsearch是一个实时分布式和开源的全文搜索和分析引擎。 它可以从RESTful Web服务接口访问，并使用模式少JSON(JavaScript对象符号)文档来存储数据。它是基于Java编程语言，这使Elasticsearch能够在不同的平台上运行。使用户能够以非常快的速度来搜索非常大的数据量。

## Elasticsearch的特性
Elasticsearch的一般特性如下 
- Elasticsearch可扩展高达PB级的结构化和非结构化数据。
- Elasticsearch可以用来替代MongoDB和RavenDB等做文档存储。
- Elasticsearch使用非标准化来提高搜索性能。
- Elasticsearch是受欢迎的企业搜索引擎之一，目前被许多大型组织使用，如- Wikipedia，The Guardian，StackOverflow，GitHub等。
- Elasticsearch是开放源代码，可在Apache许可证版本2.0下提供。
## Elasticsearch的主要概念
Elasticsearch的主要概念如下：

- 节点 - 它指的是Elasticsearch的单个正在运行的实例。单个物理和虚拟服务器容纳多个节点，这取决于其物理资源的能力，如RAM，存储和处理能力。

- 集群 - 它是一个或多个节点的集合。 集群为整个数据提供跨所有节点的集合索引和搜索功能。

- 索引 - 它是不同类型的文档和文档属性的集合。索引还使用分片的概念来提高性能。 例如，一组文档包含社交网络应用的数据。

- 类型/映射 - 它是共享同一索引中存在的一组公共字段的文档的集合。 例如，索引包含社交网络应用的数据，然后它可以存在用于用户简档数据的特定类型，另一类型可用于消息的数据，以及另一类型可用于评论的数据。

- 文档 - 它是以JSON格式定义的特定方式的字段集合。每个文档都属于一个类型并驻留在索引中。每个文档都与唯一标识符(称为UID)相关联。

- 碎片 - 索引被水平细分为碎片。这意味着每个碎片包含文档的所有属性，但包含的数量比索引少。水平分隔使碎片成为一个独立的节点，可以存储在任何节点中。主碎片是索引的原始水平部分，然后这些主碎片被复制到副本碎片中。

- 副本 - Elasticsearch允许用户创建其索引和分片的副本。 复制不仅有助于在故障情况下增加数据的可用性，而且还通过在这些副本中执行并行搜索操作来提高搜索的性能。

## Elasticsearch的优点
- Elasticsearch是基于Java开发的，这使得它在几乎每个平台上都兼容。
- Elasticsearch是实时的，换句话说，一秒钟后，添加的文档可以在这个引擎中搜索得到。
- Elasticsearch是分布式的，这使得它易于在任何大型组织中扩展和集成。
- 通过使用Elasticsearch中的网关概念，创建完整备份很容易。
- 与Apache Solr相比，在Elasticsearch中处理多租户非常容易。
- Elasticsearch使用JSON对象作为响应，这使得可以使用不同的编程语言调用Elasticsearch服务器。
- Elasticsearch支持几乎大部分文档类型，但不支持文本呈现的文档类型。
## Elasticsearch的缺点
- Elasticsearch在处理请求和响应数据方面没有多语言和数据格式支持(仅在JSON中可用)，与Apache Solr不同，Elasticsearch不可以使用CSV，XML等格式。

- Elasticsearch也有一些伤脑的问题发生，虽然在极少数情况下才会发生。

## Elasticsearch和RDBMS之间的比较
在Elasticsearch中，索引是类型的集合，因为数据库是RDBMS(关系数据库管理系统)中表的集合。每个表都是行的集合，就像每个映射都是JSON对象的Elasticsearch集合一样。

|Elasticsearch|	关系数据库|
|------|------|
|索引|	数据库|
|碎片|	碎片|
|映射|	表|
|字段|	字段|
|JSON对象|	元组|
## 参考资料
* https://es.xiaoleilu.com/
* http://m.yiibai.com/elasticsearch/
* http://www.ruanyifeng.com/blog/2017/08/elasticsearch.html