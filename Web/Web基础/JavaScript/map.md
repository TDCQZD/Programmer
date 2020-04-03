## map
Map 对象保存键值对。任何值(对象或者原始值) 都可以作为一个键或一个值。
### 语法
```
new Map([iterable])
```
### 参数
iterable
- Iterable 可以是一个数组或者其他 iterable 对象，其元素或为键值对，或为两个元素的数组。 每个键值对都会添加到新的 Map。null 会被当做 undefined。
### 描述
一个Map对象在迭代时会根据对象中元素的插入顺序来进行 — 一个  for...of 循环在每次迭代后会返回一个形式为[key，value]的数组。
**键的相等(Key equality)**
键的比较是基于 "SameValueZero" 算法：
- NaN 是与 NaN 相等的（虽然 NaN !== NaN），
- 剩下所有其它的值是根据 === 运算符的结果判断是否相等。
- 在目前的ECMAScript规范中，-0和+0被认为是相等的，尽管这在早期的草案中并不是这样

**Objects 和 maps 的比较**
Objects 和 Maps 类似的是，它们都允许你按键存取一个值、删除键、检测一个键是否绑定了值。因此（并且也没有其他内建的替代方式了）过去我们一直都把对象当成 Maps 使用。不过 Maps 和 Objects 有一些重要的区别，在下列情况里使用 Map 会是更好的选择：

- 一个Object的键只能是字符串或者 Symbols，但一个 Map 的键可以是任意值，包括函数、对象、基本类型。
- Map 中的键值是有序的，而添加到对象中的键则不是。因此，当对它进行遍历时，Map 对象是按插入的顺序返回键值。
- 你可以通过 size 属性直接获取一个 Map 的键值对个数，而 Object 的键值对个数只能手动计算。
- Map 可直接进行迭代，而 Object 的迭代需要先获取它的键数组，然后再进行迭代。
- Object 都有自己的原型，原型链上的键名有可能和你自己在对象上的设置的键名产生冲突。虽然 ES5 开始可以用 map = Object.create(null) 来创建一个没有原型的对象，但是这种用法不太常见。
- Map 在涉及频繁增删键值对的场景下会有些性能优势。

### 属性
- Map.length
    属性 length 的值为 0 。
- get Map[@@species]
    本构造函数用于创建派生对象。
- Map.prototype
    表示 Map 构造器的原型。 允许添加属性从而应用于所有的 Map 对象。
### Map 实例
所有的 Map 对象实例都会继承 Map.prototype。

**属性**
- Map.prototype.constructor
    返回一个函数，它创建了实例的原型。默认是Map函数。
- Map.prototype.size
    返回Map对象的键/值对的数量。
**方法**
- Map.prototype.clear()
    移除Map对象的所有键/值对 。
- Map.prototype.delete(key)
    如果 Map 对象中存在该元素，则移除它并返回 true；否则如果该元素不存在则返回 false
- Map.prototype.entries()
    返回一个新的 Iterator 对象，它按插入顺序包含了Map对象中每个元素的 [key, value] 数组。
- Map.prototype.forEach(callbackFn[, thisArg])
    按插入顺序，为 Map对象里的每一键值对调用一次callbackFn函数。如果为forEach提供了thisArg，它将在每次回调中作为this值。
- Map.prototype.get(key)
    返回键对应的值，如果不存在，则返回undefined。
- Map.prototype.has(key)
    返回一个布尔值，表示Map实例是否包含键对应的值。
- Map.prototype.keys()
    返回一个新的 Iterator对象， 它按插入顺序包含了Map对象中每个元素的键 。
- Map.prototype.set(key, value)
    设置Map对象中键的值。返回该Map对象。
- Map.prototype.values()
    返回一个新的Iterator对象，它按插入顺序包含了Map对象中每个元素的值 。
- Map.prototype[@@iterator]()
    返回一个新的Iterator对象，它按插入顺序包含了Map对象中每个元素的 [key, value] 数组。
### 示例