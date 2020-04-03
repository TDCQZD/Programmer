Java集合类足一种特别有用的工具类，可用于存储数量不等的对象，并可以实现常用的数据结构，如栈、队列等。除此之外,Java 集合还可用于保存具存映射关系的关联数组。Java 集合大致可以分为为 Set 、List、Queue 和 Map 四种体系，其中 Set 代表无序、不可重复的集合;List代表有序、可重复的集合;而Map 则代表具有映射关系的集合;Java 5又增加了Queue 体系集合，代表一种队列集合实现。

一、Java集合概述

**Java引入集合的目的：**为了保存数量不确定的数据，以及保存具有映射关系的数据（也被称为关联数组），Java 提供了集合类。

集合类主要负责保存、盛装其他数据，因此集合类也被称为容器类。所有的集合类都位于 java.util包下，后来为了处理多线程环境的并发安全问题 ，Java 5还在 java.util.concurrent 包下提供了一些多线程支持的集合类。

集合类和数组不一样，数组元素既可以是基本类型的值，也可以是对象（实际上保存的是对象的引用变量 \);而集合里只能保存对象（实际上只是保存对象的引用变量，但通常习惯上认为集合里保存的是对象）。

Java 的集合类主要由两个接口派生而出： Collection 和 Map ,  Collection和Map是Java集合框架的根接口，这两个接口又包含了一些子接口或实现类。如下图所示是 Collection 接口、子接口及其实现类的继承树。![](/assets/集合.png)二、Collection接口

Collection接口

Collection接口是List、Set和Queue接口的父接口，该接口里定义的方法既可用于操作Set集合，也可用于操作List和Queue 集合。Collection 接口里定义了如下操作集合元素的方法。

\(1\)boolean add \(Object o\) 该方法用于向集合里添加一个元素。如果集合对象被添加操作改变了,则返回true。

\(2\)boolean addAll \(Collection c\):该方法把集合c里的所有元素添加到指定集合里。如果集合对象被添加操作改变,则返回true 。

\(3\)void clear\(\) :淸除集合里的所有元素，将集合长度变为0。

\(4\)boolean contains\(Object o\): 返回集合里是否包含指定元索。

\(5\)boolean containsAll\(Collection c\): 返回集合里是否包含集合c里的所有元素。

\(6\)boolean isEmpty\(\): 返回集合是否为空。当集合长度为0时返回true，否则返回false。

\(7\)Iterator iterator\(\) : 返冋一个 Iterator 对象，用遍历集介里的元索。

\(8\)boolean remove\(Object o\): 刪除集合中的指定元素o,当集合中了一个或多个元素o时,该方法只删除第一个符合条件的元素，该方法将返回 true 。

\(9\)boolean removeAll\(Collection c\): 从集合中删除集合c里包含的所有元素（相当于用调用该方法的集合减集合 c\), 如果删除了一个或一个以上的元素 ， 则该方法返回 true 。

\(10\)boolean retainAll\(Collection c\): 从集合中删除集合c里不包含的元索（相当于把调用该方法的集合变成该集合和集合c的交集\)，如果该操作改变了调用该方法的集合，则该方法返回 true 。

\(11\)int size\(\) :该方法返回集合里元索的个数。

\(12\)Object\[\] toArray\(\): 该方法把集合转换成一个数组，所冇的集合元索变成对应的数组元素。

```
public class CollectionTest{
    public static void main(String[] args){
        Collection c = new ArrayList();
        // 添加元素
        c.add("孙悟空");
        // 虽然集合里不能放基本类型的值，但Java支持自动装箱
        c.add(6);
        System.out.println("c集合的元素个数为:" + c.size()); // 输出2
        // 删除指定元素
        c.remove(6);
        System.out.println("c集合的元素个数为:" + c.size()); // 输出1
        // 判断是否包含指定字符串
        System.out.println("c集合的是否包含\"孙悟空\"字符串:"
            + c.contains("孙悟空")); // 输出true
        c.add("轻量级Java EE企业应用实战");
        System.out.println("c集合的元素：" + c);
        Collection books = new HashSet();
        books.add("轻量级Java EE企业应用实战");
        books.add("疯狂Java讲义");
        System.out.println("c集合是否完全包含books集合？"
            + c.containsAll(books)); // 输出false
        // 用c集合减去books集合里的元素
        c.removeAll(books);
        System.out.println("c集合的元素：" + c);
        // 删除c集合里所有元素
        c.clear();
        System.out.println("c集合的元素：" + c);
        // 控制books集合里只剩下c集合里也包含的元素
        books.retainAll(c);
        System.out.println("books集合的元素:" + books);
    }
}
```

三、操作集合元素

3.1Iterator接口遍历集合元素

Iterator接口主要用于遍历（即迭代访问 ）Collection集合中的元索，Iterator 对象也被称为迭代器。Iterator 接口隐藏了各种 Collection 实现类的底层细节，向应用程序提供遍历 Collection 集合元素的统一编程接口。Iterator 接口甩定义了如下4个方法。

\(1\)boolean hasNext\(\):如果被迭代的集合元索还没有被遍历完，则返回true

\(2\)Object next\(\):返回集合里的下一个元索。

\(3\)void remove\(\):删除集合里上一次next方法返回的元素。

\(4\)void forEachRemaining \( Consumer action \)，这是Java8为Iterator 新增的默认方法，该方法可使用Lambda 表达式来遍历集合元素.

Iterator接口遍历集合元素：

```
public class IteratorTest{
    public static void main(String[] args){
        // 创建集合、添加元素的代码与前一个程序相同
        Collection books = new HashSet();
        books.add("轻量级Java EE企业应用实战");
        books.add("疯狂Java讲义");
        books.add("疯狂Android讲义");
        // 获取books集合对应的迭代器
        Iterator it = books.iterator();
        while(it.hasNext())
        {
            // it.next()方法返回的数据类型是Object类型，因此需要强制类型转换
            String book = (String)it.next();
            System.out.println(book);
            if (book.equals("疯狂Java讲义"))
            {
                // 从集合中删除上一次next方法返回的元素
                it.remove();
            }
            // 对book变量赋值，不会改变集合元素本身
            book = "测试字符串";   //①
        }
        System.out.println(books);
    }
}
```

使用Lambad表达式遍历Iterator:

```
public class IteratorEach{
    public static void main(String[] args){
        // 创建集合、添加元素的代码与前一个程序相同
        Collection books = new HashSet();
        books.add("轻量级Java EE企业应用实战");
        books.add("疯狂Java讲义");
        books.add("疯狂Android讲义");
        // 获取books集合对应的迭代器
        Iterator it = books.iterator();
        // 使用Lambda表达式（目标类型是Comsumer）来遍历集合元素
        it.forEachRemaining(obj -> System.out.println("迭代集合元素：" + obj));
    }
}
```

3.2foreach遍历集合元素

```
public class ForeachTest{
    public static void main(String[] args){
        // 创建集合、添加元素的代码与前一个程序相同
        Collection books = new HashSet();
        books.add(new String("轻量级Java EE企业应用实战"));
        books.add(new String("疯狂Java讲义"));
        books.add(new String("疯狂Android讲义"));
        for (Object obj : books)
        {
            // 此处的book变量也不是集合元素本身
            String book = (String)obj;
            System.out.println(book);
            if (book.equals("疯狂Android讲义"))
            {
                // 下面代码会引发ConcurrentModificationException异常
                books.remove(book);     //①
            }
        }
        System.out.println(books);
    }
}
```

3.3Java8新增Predicate删除集合元素。

Java8为Collection集合新增了一个removeIf\(Predicate filter\)方法，该方法将会批量删除符号filter条件的所有元素。

```
public class PredicateTest2{
    public static void main(String[] args){
        // 创建books集合、为books集合添加元素的代码与前一个程序相同。
        Collection books = new HashSet();
        books.add(new String("轻量级Java EE企业应用实战"));
        books.add(new String("疯狂Java讲义"));
        books.add(new String("疯狂iOS讲义"));
        books.add(new String("疯狂Ajax讲义"));
        books.add(new String("疯狂Android讲义"));
        // 统计书名包含“疯狂”子串的图书数量
        System.out.println(calAll(books , ele->((String)ele).contains("疯狂")));
        // 统计书名包含“Java”子串的图书数量
        System.out.println(calAll(books , ele->((String)ele).contains("Java")));
        // 统计书名字符串长度大于10的图书数量
        System.out.println(calAll(books , ele->((String)ele).length() > 10));
    }
    public static int calAll(Collection books , Predicate p)
    {
        int total = 0;
        for (Object obj : books)
        {
            // 使用Predicate的test()方法判断该对象是否满足Predicate指定的条件
            if (p.test(obj))
            {
                total ++;
            }
        }
        return total;
    }
}
```

3.4Java8新增Stream聚集操作。

Stream是一个通用的流接口。Sream提供了大量的方法进行聚集操作，这些方法既可以是“中间的“\(imermcdiaic\)，也可以是”末端的”。\(imermcdiaic\)。

\(1\)中间方法：中间操作允许流保持打开状态，并允许直接调用后续方法.上面程序中的map\(\)方法就是中间方法。中间方法的的返回值是另外一个流。

\(2\)末端方法：木端方法是对流的最终操作。当对某个Stream执行末端方法后，该流将会被“诮耗”且不再可用。

除此之外，关于流的方法还有如下两个特征.

\(1\)有状态的方法：这种方法会给流增加一些新的属性,比如元素的唯-性、元素的最大数量、保证元素以排序的方式被处理等.有状态的方法往往需要更大的性能开销，

\(2\)短路方法： 短路方法可以尽的结束对流的操作.不必检查所有的元素.。

Stream对元素进行批量操作：

```
public class CollectionStream{
    public static void main(String[] args){
        // 创建books集合、为books集合添加元素的代码与8.2.5小节的程序相同。
        Collection books = new HashSet();
        books.add(new String("轻量级Java EE企业应用实战"));
        books.add(new String("疯狂Java讲义"));
        books.add(new String("疯狂iOS讲义"));
        books.add(new String("疯狂Ajax讲义"));
        books.add(new String("疯狂Android讲义"));
        // 统计书名包含“疯狂”子串的图书数量
        System.out.println(books.stream()
            .filter(ele->((String)ele).contains("疯狂"))
            .count()); // 输出4
        // 统计书名包含“Java”子串的图书数量
        System.out.println(books.stream()
            .filter(ele->((String)ele).contains("Java") )
            .count()); // 输出2
        // 统计书名字符串长度大于10的图书数量
        System.out.println(books.stream()
            .filter(ele->((String)ele).length() > 10)
            .count()); // 输出2
        // 先调用Collection对象的stream()方法将集合转换为Stream,
        // 再调用Stream的mapToInt()方法获取原有的Stream对应的IntStream
        books.stream().mapToInt(ele -> ((String)ele).length())
            // 调用forEach()方法遍历IntStream中每个元素
            .forEach(System.out::println);// 输出8  11  16  7  8
    }
}
```

四、Set集合

Set集合中元素是无序和不可重复的。

4.1HashSet类

HashSet是Set接口的典型实现，大多数时候使用 Set 集合时就是使用这个实现类。HashSet按Hash  
算法来存储集合中的元素，因此具有很好的存取和查找性能。

HashSet 具有以下特点。

\(1\)不能保证元素的排列顺序，顺序可能与添加顺序不同，顺序也有可能发生变化。

\(2\)HashSet不是同步的,如果多个线程同时访问一个 HashSet, 假设有两个或者两个以上线程同时修改了HashSet集合时，则必须通过代码来保证其同步。

\(3\)集合元素值可以是null。

当向HashSet集合中存入一个元素时,HashSet会调用该对象的hashCode\(\)方法来得到该对象的hashCode值，然后根据该hashCode值决定该对象在HashSet中的存储位置。如果有两个元素通过equals\(\)方法比较返回true,但它们的 hashCode\(\)方法返回值不相等， HashSet 将会把它们存储在不同的位置，依然可以添加成功。也就是说， HashSet 集合判断两个元素相等的标准是两个对象通过equals\(\)方法比较相等，并且两个对象的 hashCode\(\)方法返回值也相等。

```
// 类A的equals方法总是返回true，但没有重写其hashCode()方法
class A{
    public boolean equals(Object obj)
    {
        return true;
    }
}
// 类B的hashCode()方法总是返回1，但没有重写其equals()方法
class B{
    public int hashCode()
    {
        return 1;
    }
}
// 类C的hashCode()方法总是返回2，且重写其equals()方法总是返回true
class C{
    public int hashCode()
    {
        return 2;
    }
    public boolean equals(Object obj)
    {
        return true;
    }
}
public class HashSetTest{
    public static void main(String[] args){
        HashSet books = new HashSet();
        // 分别向books集合中添加两个A对象，两个B对象，两个C对象
        books.add(new A());
        books.add(new A());
        books.add(new B());
        books.add(new B());
        books.add(new C());
        books.add(new C());
        System.out.println(books);
    }
}
```

**hashCode\(\)**

把一个对象放入HashSet中时，如果需要重写该对象对应类的equals\(\)方法。则也应该重写其hashCode\(\)方法。规则是：如果两个对象通过equals\(\)方法比较返回true,这两个对象的hashCode也应该相同。

如果两个对象通过 equals\(\)方法比较返回 true ，但这两个对象的 hashCode \(\)方法返回不同的 hashCode值时，这将导致 HashSet 会把这两个对象保存在 Hash 表的不同位置，从而使两个对象都可以添加成功，这就与 Set 集合的规则冲突了。

如果两个对象的 hashCode \(\)方法返回的 hashCode 值相同，但它们通过 equals \(\)方法比较返回 false时将更麻烦：因为两个对象的 hashCode 值相同， HashSet 将试图把它们保存在同一个位置，但又不行（否则将只剩下一个对象），所以实际上会在这个位置用链式结构来保存多个对象：而 HashSet 访问集合元素时也是根据元素的 hashCode 值来快速定位的，如果 HashSet 中两个以上的元素具有相同的 hashCode值，将会导致性能下降。

**如果需要把某个类的对象保存到 HashSet 集合中，重写这个类的 equals \(\)方法和hashCode \(\)方法时，应该尽量保证两个对象通过 equals \(\)方法比较返回 true 时，它们的hashCode \(\)方法返回值也相等。**

hashCode（）方法的基本规则：

\(1\)在程序运行过程中，同一个对象多次调用 hashCode\(\)方法应该返回相同的值。

\(2\)当两个对象通过 equals \(\)方法比较返回 true 时，这两个对象的 hashCode \(\)方法应返回相等的值。

\(3\)对象中用作 equals\(\)方法比较标准的实例变都应该用于计算 hashCode 值。

重写 hashCode\(\)方法的一般步骤：

\(1\)\)把对象内每个有意义的实例变量\(即毎个参与equals\(\)方法比较标准的实例变量\)计算出一个int 类型的 hashCode 值。

![](/assets/hashcode.png)

\(2\)用第1步计算出来的多个 hashCode 值组合计兑出一个 hashCode 值返回

```
return fl.hashCode() + (int)f2;
```

为了避免直接相加产生偶然相等（两个对象的 fl 、 f 2实例变量 并不相等，但它们的 hashCode 的和恰好相等），可以通过为各实例变量 的 hashCode 值乘以任意一个质数后再相加。

```
return fl.hashCode() * 19 + (int)f2 * 31;
```

如果向 HashSet 中添加一个吋变对象后，后面程序修改了该可变对象的实例变量，则可能导致它与集合中的其他元索相同（即两个对象通过 equals\(\)方法比较返回 true ,两个对象的 hashCode 值也相等），这就有可能导致HashSet中包含两个相同的对象。如下代码，

```
class R{
    int count;
    public R(int count){
        this.count = count;
    }
    public String toString(){
        return "R[count:" + count + "]";
    }
    public boolean equals(Object obj){
        if(this == obj)
            return true;
        if (obj != null && obj.getClass() == R.class)
        {
            R r = (R)obj;
            return this.count == r.count;
        }
        return false;
    }
    public int hashCode(){
        return this.count;
    }
}
public class HashSetTest2{
    public static void main(String[] args){
        HashSet hs = new HashSet();
        hs.add(new R(5));
        hs.add(new R(-3));
        hs.add(new R(9));
        hs.add(new R(-2));
        // 打印HashSet集合，集合元素没有重复
        System.out.println(hs);
        // 取出第一个元素
        Iterator it = hs.iterator();
        R first = (R)it.next();
        // 为第一个元素的count实例变量赋值
        first.count = -3;     // ①
        // 再次输出HashSet集合，集合元素有重复元素
        System.out.println(hs);
        // 删除count为-3的R对象
        hs.remove(new R(-3));    // ②
        // 可以看到被删除了一个R元素
        System.out.println(hs);
        System.out.println("hs是否包含count为-3的R对象？"
            + hs.contains(new R(-3))); // 输出false
        System.out.println("hs是否包含count为-2的R对象？"
            + hs.contains(new R(-2))); // 输出false
    }
}
```

4.2LinkedHashSet类

HashSet还有一个子类LinkedHashSet,LinkedHashSet 集合也是根据元素的 hashCode值来决定元素的存储位置，它同时使用链表维护元索的次序，这样使得元素看起来是以插入的顺序保存的。也就是说，当遍历 LinkedHashSet 集合里的元素时， LinkedHashSet 将会按元索的添加顺序来访问集合里的元素。

LinkedHashSet 需要维护元素的插入顺序，因此性能略低。HashSet的性能，但在迭代访问 Set 里的全部元素时将有很好的性能，因为它以链表来维护内部顺序。

```
public class LinkedHashSetTest{
    public static void main(String[] args){
        LinkedHashSet books = new LinkedHashSet();
        books.add("疯狂Java讲义");
        books.add("轻量级Java EE企业应用实战");
        System.out.println(books);
        // 删除 疯狂Java讲义
        books.remove("疯狂Java讲义");
        // 重新添加 疯狂Java讲义
        books.add("疯狂Java讲义");
        System.out.println(books);
    }
}
```

4.3TreeSet类

TreeSet是SortedSet接口的实现类，TreeSet可以确保集合元素处于排序状态。TreeSet并不是根据元素的插入顺序进行排序的，而是根据元素实际值的大小来进行排序的。

HashSet集合采用 hash 算法来决定元素的存储位置不同， TreeSet 采用红黑树的数据结构来存储集合元素。 TreeSet 支持两种排序方法：自然排序和定制排序。在默认情况下，TreeSet 采用自然排序。

1、自然排序

TreeSet会调用集合元素的compareTo\(Object  obj\)方法来比较元素之间的大小关系，然后将集合元素按升序排列，这种方式就是自然排序。

Java 提供了一个 Comparable 接口，该接口里定义了一个 compareTo\(Object obj \)方法，该方法返回一个整数值，实现该接口的类必须实现该方法，实现了该接口的类的对象就可以比较大小。当一个对象调用该方法与另一个对象进行比较时，例如 objl.compareTo\(Obj2\),如果该方法返回0,则表明这两个对象相等：如果该方法返回一个正整数，则表明 objl 大于 obj2;如果该方法返回一个负整数，则表明 objl小于obj2。

Java的一些常用类已经实现了 Comparable 接口，并提供了比较大小的标准。实现Comparable 接口的常用类。

\(1\)BigDecimal 、 Biglnteger 以及所有的数值型对应的包装类：按它们对应的数值大小进行比较。

\(2\)Character :按字符的 UNICODE 值进行比较。

\(3\)Boolean :  true 对应的包装类实例大于 false 对应的包装类实例。

\(4\)String :按字符串屮字符的 UNICODE 值进行比较。

\(5\)Date 、 Time :后面的时间、日期比前面的时间、日期大。

如果试阁把一个对象添加到TreeSet时，则该对象的类必须实现Comparable接口，否则程序将会抛出异常。

TreeSet只能添加同一种类型的对象，否则会ClassCastException异常。

2、定制排序

TreeSet的自然排序是根据集合元素的大小， TreeSet 将它们以升序排列。如果需要实现定制排序，例如以降序排列，则可以通过 Comparator 接口的帮助。该接口里包含一个 int compared\(T o1,T o2\)方法，该方法用于比较o1和o2的大小：如果该方法返回正整数，则表明o1大于o2;如果该方法返回0,则表明o1等于o2 ; 如果该方法返回负整数，则表明o1小于o2。

如果需要实现定制排序，则需要在创建 TreeSet 集合对象时，提供一个 Comparator 对象与该 TreeSet集合关联，由该 Comparator 对象负责集合元素的排序逻辑。由于 Comparator 是一个函数式接口，因此可使用Lambda表达式来代替Comparator 对象。

```
class M{
    int age;
    public M(int age)
    {
        this.age = age;
    }
    public String toString()
    {
        return "M[age:" + age + "]";
    }
}
public class TreeSetTest4{
    public static void main(String[] args)
{
        // 此处Lambda表达式的目标类型是Comparator
        TreeSet ts = new TreeSet((o1 , o2) ->
        {
            M m1 = (M)o1;
            M m2 = (M)o2;
            // 根据M对象的age属性来决定大小，age越大，M对象反而越小
            return m1.age > m2.age ? -1
                : m1.age < m2.age ? 1 : 0;
        });
        ts.add(new M(5));
        ts.add(new M(-3));
        ts.add(new M(9));
        System.out.println(ts);
    }
}
```

4.4EnumSet类

EnumSet是一个专为枚举类设计的集合类，EnumSet中的所存元素都必须是指定枚举类型的枚举类，该枚举类把在创建EnumSet时显式或隐式地指定。EnumSet的集合元素是有序的，EnumSet以枚举值在 EnumSet类内的定义顺序来决定集合元素的顺序。

EnumSet在内部以位向量的形式存储，这种存储形式非常紧凑、高效，因此 EnumSet对象占用内存很小，而且运行效率很好。尤其是进行批量操作（如调用 containsAll\(\)和 retainAll\(\)方法）时，如果其参数也是EnumSet 集合，则该批量操作的执行速度也非常快。

EnumSet集合不允许加入 null元素，如果试图插入null元素，EnumSet 将抛出 NullPointerException异常。如果只是想判断 EnumSet 是否包含null元素或试图删除null元素都不会抛出异常，只是删除操作将返回false ，因为没有任何 null 元素被刪除。

EnumSet 类没有暴露任何构造器来创建该类的实例，程序应该通过它提供的类方法来创建EnumSet对象。 EnumSet 类它提供了如下常用的类方法来创建 EnumSet 对象。

\(1\) EnumSet allOf\(Class elementType \):创建一个包含指定枚举类里所有枚举值的 EnumSet 集合。

\(2\)EnumSet complementOfl\(EnumSet s \):创建一个其元素类型与指定 EnumSet 里元素类型相同的EnumSet 集合，新 EnumSet 集合包含原 EnumSet 集合所不包含的、此枚举类剩下的枚举值（即新 EnumSet 集合和原 EnumSet 集含的集合元素加起来就是该枚举类的所有枚举值\)。

\(3\)EnumSet copyOf \( Collection c \):使用一个普通集合来创建  EnumSet 集含，

\(4\)EnumSet copyOf\(EnumSet s \):创建一个与指定 EnumSet 具有相同元素类型、相同集合元素的EnumSet 集合。

\(5\)EnumSet noneOf\(Class elementType \):创建一个元素类型为指定枚举类型的空 EnumSet ,

\(6\)EnumSet of\(E first ,  E ... rest \):创建一个包含一个或多个枚举值的 EnumSet 集合，传入的多个枚举值必须属于同—个枚举类。

\(7\)EnumSet range\(E from，E to \):创建一个包含从 from 枚举值到 to 枚举值范围内所有枚举值的EnumSet 集含。

```
enum Season{
    SPRING,SUMMER,FALL,WINTER
}
public class EnumSetTest{
    public static void main(String[] args){
        // 创建一个EnumSet集合，集合元素就是Season枚举类的全部枚举值
        EnumSet es1 = EnumSet.allOf(Season.class);
        System.out.println(es1); // 输出[SPRING,SUMMER,FALL,WINTER]
        // 创建一个EnumSet空集合，指定其集合元素是Season类的枚举值。
        EnumSet es2 = EnumSet.noneOf(Season.class);
        System.out.println(es2); // 输出[]
        // 手动添加两个元素
        es2.add(Season.WINTER);
        es2.add(Season.SPRING);
        System.out.println(es2); // 输出[SPRING,WINTER]
        // 以指定枚举值创建EnumSet集合
        EnumSet es3 = EnumSet.of(Season.SUMMER , Season.WINTER);
        System.out.println(es3); // 输出[SUMMER,WINTER]
        EnumSet es4 = EnumSet.range(Season.SUMMER , Season.WINTER);
        System.out.println(es4); // 输出[SUMMER,FALL,WINTER]
        // 新创建的EnumSet集合的元素和es4集合的元素有相同类型，
        // es5的集合元素 + es4集合元素 = Season枚举类的全部枚举值
        EnumSet es5 = EnumSet.complementOf(es4);
        System.out.println(es5); // 输出[SPRING]
    }
}
```

4.5各种Set实现类的性能分析

HashSet和 TreeSet是Set的两个典型实现，到底如何选择 HashSet 和 TreeSet 呢？  HashSet 的性能总是比 TreeSet 好（特別是最常用的添加、查询元素等操作），因为 TreeSet 需要额外的红黑树算法来维护集合元素的次序。只有当需要一个保持排序的 Set 时，才应该使用 TreeSet ,否则都应该使用 HashSetHashSet 还有一个子类： LinkedHashSet ，对于普通的插入、删除操作 ， LinkedHashSet 比  HashSet要略微慢一点，这是由维护链表所带来的额外开销造成的，但由于有了链表，遍历 LinkedHashSet 会更快。

EnumSet 是所存 Set实现类中性能最好的，但它只能保存同一个枚举类的枚举值作为集合元素。必须指出的是，Set 的三个实现类 HashSet 、 TreeSet 和 EnumSet 都是线程不安全的。如果有多个线程同时访问一个 Set 集合，并且有超过一个线程修改了该 Set 集合，则必须手动保证该 Set 集合的同步性。通常可以通过 Collections 工具类的 synchronizedSortedSet 方法来“包装”该 Set 集合。此操作最好在创建时进行，以防止对Set集合的意外非同步访问。

五、List集合

List集合代表一个元素有序 、可重复的集合，集合中每个元素都有其对应的顺序索引。List集合允许使用元素 ，可以通过索引来访问指定位置的集合元素。List集合默认按元素的添加顺序设置元素的索引。

5.1List接口

List作为Collection接口的子接口 ， 当然可以使用Collection 接口里的全部方法。而且由于List是有序集合，因此List集合里增加了一些根据索引来操作集合元素的方法。

\(1\)void add\(int index,Object element \):将元素element 插入到List集合的index处。

\(2\)boolean addAll\(int index，Collection c \):将集合c所包含的所有元素都插入到List集合的index处.

\(3\)Object get\(int index \):返回集合 index 索引处的元索。

\(4\)int indexOf\( Object o \):返回对象o在 List 集合中第一次出现的位置索引。

\(5\)int lastIndexOf\(Object o \) : 返回对象o在 List 集合中最后一次出现的位置索引。

\(6\)Object remove\(int index \):删除并返回index索引处的元素。

\(7\)Object set\(int index ,  Object element \):将 index 索引处的元素替换成element对象，返回被替换的旧元素，

\(8\)List subList\(int fromIndex,int toIndex\):返回从索引fromIndex\(包含）到索引toIndex\(不包含）处所有集合元素组成的子集合。

所有的List实现类都可以调用这些方法来操作集合元素与Set集合相比，List 增加了根据索引来插入、替换和删除集合元素的方法。除此之外 ，Java 8还为 List 接口添加了如下两个默认方法。

\(1\)void replaceAll\(UnaryOperator operator \):根据operator 指定的计算规则重新设置 List 集合的所有元素。

\(2\)void sort\(Comparator c \) : 根据 Comparator 参数对 List 集合的元素排序。

```
public class ListTest{
    public static void main(String[] args){
        List books = new ArrayList();
        // 向books集合中添加三个元素
        books.add(new String("轻量级Java EE企业应用实战"));
        books.add(new String("疯狂Java讲义"));
        books.add(new String("疯狂Android讲义"));
        System.out.println(books);
        // 将新字符串对象插入在第二个位置
        books.add(1 , new String("疯狂Ajax讲义"));
        for (int i = 0 ; i < books.size() ; i++ ){
            System.out.println(books.get(i));
        }
        // 删除第三个元素
        books.remove(2);
        System.out.println(books);
        // 判断指定元素在List集合中位置：输出1，表明位于第二位
        System.out.println(books.indexOf(new String("疯狂Ajax讲义"))); //①
        //将第二个元素替换成新的字符串对象
        books.set(1, new String("疯狂Java讲义"));
        System.out.println(books);
        //将books集合的第二个元素（包括）
        //到第三个元素（不包括）截取成子集合
        System.out.println(books.subList(1 , 2));
    }
}
```

5.2Listlterator接口

与Set只提供了一个iterator\(\)方法不同，List还额外提供了一个listIterator方法，该方法返回一个Listlterator 对象，ListUerator接口继承Iterator接口，提供了专门操作 List 的方法 .ListIterator 接口在Iterator 接口基础上增加了如下方法。

\(1\)boolean hasPrevious\(\): 返回该迭代器关联的集合足否还有上一个元素。

\(2\)Object previous\(\): 返冋该迭代器的上一个元素。

\(3\)void add\(Object o\): 在指定位赏插入一个元素。

Listlterator与普通的Iterator 进行对比， Listlterator 增加了向前迭代的功能\(Iterator只能向后迭代\) ，而且LisIterator 还可通过add\(\)方法向List集合中添加元素（Iterator 只能删除元素）。

```
public class ListIteratorTest
{
    public static void main(String[] args)
    {
        String[] books = {
            "疯狂Java讲义", "疯狂iOS讲义",
            "轻量级Java EE企业应用实战"
        };
        List bookList = new ArrayList();
        for (int i = 0; i < books.length ; i++ )
        {
            bookList.add(books[i]);
        }
        ListIterator lit = bookList.listIterator();
        while (lit.hasNext())
        {
            System.out.println(lit.next());
            lit.add("-------分隔符-------");
        }
        System.out.println("=======下面开始反向迭代=======");
        while(lit.hasPrevious())
        {
            System.out.println(lit.previous());
        }
    }
}
```

5.3ArrayList和Vector实现类

ArrayList和Vector 类都是基于数组实现的 List类，所以ArrayList和Vector类封装了一个动态的、允许再分配的Object\[\] 数姐， ArrayList或Vector对象使用initialCapacity参数来设置该数组的长度，当向ArrayList或Vector中添加元素超出了该数组的长度时，它们的 initialCapacity会自动增加。

如果向 ArrayList或 Vector 集合中添加大量的元素时 ，可使用 ensureCapacity\(int minCapacity\)方法一次性地增加initialCapacity 。这可以减少分配的次数，从而提高性能。

如果开始就知道ArrayList或 Vector集合需要保存多少个元索，则可以在创建它们时就指定initialCapacity大小。如果创建空的ArrayList或Vector 集合时不指定 initialCapacity  参数，则 Object\[\]数组的长度默认为10 。

除此之外， ArrayList 和 Vector 还提供了两个方法来重新分配Object\[\]数组，

\(1\)void ensureCapacity\(int minCapacity \):将 ArrayList  或 Vector  集合的 Object\[\]数组长度增加大于或等于minCapacity值。

\(2\)void trimToSize\(\): 调整ArrayList或Vector集合的Object\[\]数组长度为当前元素的个数。调用该方法时减少ArrayList或Vector集合对象占用的存储空间。

ArrayList和Vector的最大的区别：ArrayListid 线程不安全的，当多个线程访问同一个ArrayList 集合时，如果有超过一个线程修改了ArrayList 集合，则程序必须手动保证该集合的同步性。

但Vector 集合则是线程安全的，无须程序保证该集合的同步性。因为 Vector 是线程安全的，所以 Vector的性能比 ArrayList 的性能要低。实际，即使需要保证 List 集合线程安全，也同样不推荐使用 Vector实现类。

Vector 还提供了一个 Stack 子类，它用于模拟“栈”这种数据结构，“栈”通常是指“后进先出”\(LIFO\)的容器。最后“ push ”进找的元素，将最先被“ pop ”出栈。与Java中的其他集合一样，进栈出栈的都是 Object ,因此从栈中取出元素后必须进行类型转换，除非你只是使用Object具有的操作。所以 Stack类里提供以下方法。

\(1\)Object peek \(\) : 返回“栈”的第一个元素，但并不将该元素“ pop ”出栈。

\(2\)Object pop \(\) : 返冋“栈”的第一个元素，并将该元素“ pop ”出栈•

\(3\)void puh \( Object item \) : 将一个元素“ push ”进栈，最后一个进“栈”的元素总是位于“栈”顶.

需要指出的是，由于 Stack 继承了  Vector ,因此它也是一个非常古老的 Java 集合类，它同样是线程安全的、性能较差的，因此应该尽量少用 Stack 类。如采程序需要使用“栈”这种数据结构，则可以使用 ArrayDeque 。

5.4固定长度的List

Arrays工具类里提供了asList\(Object... a\)方法，该方法可以把一个数组或指定个数的对象转换成一个List集合，这个List集合既不是ArrayList实现类的实例，也不是Vector实现类的实例，而足Arrays的内部类ArrayList的实例。

Arrays.ArrayList是一个固定长度的List集合，程序只能遍历访问该集合里的元素，不可增加、删除该集合里的元索。

```
public class FixedSizeList{
    public static void main(String[] args){
        List fixedList = Arrays.asList("疯狂Java讲义"
            , "轻量级Java EE企业应用实战");
        // 获取fixedList的实现类，将输出Arrays$ArrayList
        System.out.println(fixedList.getClass());
        // 使用方法引用遍历集合元素
        fixedList.forEach(System.out::println);
        // 试图增加、删除元素都会引发UnsupportedOperationException异常
        fixedList.add("疯狂Android讲义");
        fixedList.remove("疯狂Java讲义");
    }
}
```

六、Queue集合

Queue用于模拟队列这种数据结构，队列通常是指“先进先出”\(FIFO\)的容器。队列的头部保存在在队列中存放时间最长的元素,队列的头部保存在队列中存放时间最短的元素。新元索插入（ offer \)到队列的尾部，访问元素 \( poll \)操作会返回队列头部的元素。通常，队列不允许随机访问队列中的元素。

Queue 接口中定义了如下几个方法。

\(1\)void add\(Object e\): 将指定元素加入此队列的尾部。

\(2\)Object element\(\):获取队列头部的元素，但是不删除该元素。

\(3\)boolean ofTeKObjecte \):将指定元素加入此队列的尾部。当使用有容最限制的队列时，此方法通常比 add\(Object e \)方法更好。

\(4\)Object peek \(\) : 获取队列头部的元素,但是不删除该元素。如果此队列为空，则返回 null 。

\(5\)Object poll \(\):获取队列头部的元素，并删除该元素。如果此队列为空，则返回 null 。

\(6\)Object remove \(\) : 获取队列头部的元素，并删除该元素。

Queue接口有一个PriorityQueue实现类和一个Deque接口，Deque 代表一个“双端队列”，双端队列可以同时从两端来添加、删除元素，因此 Deque 的实现类既可当成队列使用，也当成栈使用。 Java为Deque提供了ArrayDeque和LinkedList两个实现类。

6.1 PriorityQueue  实现类

PriorityQueue 是一个比较标准的队列实现类。之所以说它是比较标准的队列实现，而不是绝对标准的队列实现，是因为 PriorityQueue 保存队列元素的顺序并不是按加入队列的顺序，而是按队列元素的大小进行重新排序。因此当调用 peek \(\)方法或者 poll \(\)方法取出队列中的元素时，并不是取出最先进入队列的元素，而是取出队列中最小的元素。从这个意义上来说 ， PriorityQueue 已经违反了队列的最基本用法。

```
public class PriorityQueueTest{
    public static void main(String[] args){
        PriorityQueue pq = new PriorityQueue();
        // 下面代码依次向pq中加入四个元素
        pq.offer(6);
        pq.offer(-3);
        pq.offer(20);
        pq.offer(18);
        // 输出pq队列，并不是按元素的加入顺序排列
        System.out.println(pq); // 输出[-3, 6, 20, 18]
        // 访问队列第一个元素，其实就是队列中最小的元素：-3
        System.out.println(pq.poll());
    }
}
```

PriorityQueue不允许插入null元素，它还需要对队列元素进行排序，PriorityQueue的元素有两种排序方式，

\(1\)自然排序:采用自然顺序的 PriorityQueue集合中的元素必须实现了Comparable接口，而且应该是同一个类的多个实例，否则可能导致ClassCastException异常。

\(2\)定制排序：创建PriorityQueue 队列时，传入一个Comparator对象，该对象负责对队列中的所有元素进行排序。采用定制排序时不要求队列元素实现 Comparable 接口。

PriorityQueue 队列对元素的要求与TreeSet 对元素的要求基本一致。

6.2Deque 接口与 ArrayDeque 实现类

Deque接口是Queue接口的子接口，它代表一个双端队列，Deque接口里定义了一些双端队列的方法，这些方法允许从两端来操作队列的元素。

\(1\)void addFirst\(Object e\)将指定元素插入该双端队列的开头。

\(2\)void addLast&lt;Object e\):将指定元素插入该双端队列的末尾。

\(3\)Iterator descendingherator\(\):返回该双端队列对应的迭代器，该迭代器将以逆向顺序来迭代队列中的元素。

\(4\)Object getFirst\(\):获取但不删除双端队列的第一个元素。

\(5\)Object gctLast\(\):获取但不删除双端队列的最后一个元素。

\(6\)boolean offerFirst\(Object e\):将指定元素插入该双端队列的开头.

\(7\)boolean offerLast\(Object e\)：将指定元素插入该双端队列的未尾.

\(8\)Object peekFirst\(\):获取但不删除该双端队列的第一个元素；如果此双端队列为空，则返回null.

\(9\)Object pcckLast\(\):获取但不删除该双端队列的最后一个元素：如果此双端队列为空，则返回null.

\(10\)Object pollFirst\(\):获取并删除该双端队列的第一个元素：如果此双端队列为空，则返回null.

\(11\)Object pollast\(\):获取并删除该双端队列的最后一个元素：如果此双端队列为空，则返回null,

\(11\)Object pop\(\) \(栈方法）:pop出该双端队列所表示的栈的栈顶元素。相当于 rcmovcFirst\(\).

\(12\)void push\(Object e\)\(栈方法）：将一个元素push进该双端队列所表示的栈的栈顶.相当于addFirst\(e\).

\(13\)Object removeFirst\(\):获取并删除该双端队列的第一个元素•

\(14\)Object removeFirstOccurrence\(Object o\):删除该端队列的第一次出现的元索o.

\(15\)Object removeLast\(\):获取并删除该双端队列的最后一个元素.

\(16\)boolean removeLastOccurrcnce\(Objcct o\): 删除该双端队列的最后一次出现的元素o.

![](/assets/Queue.png)Deque接口提供了一个典型的实现类： ArrayDeque ,从该名称就可以看出，它是一个基于数组实现的双端队列，创建Deque时同样可指定一个 numElements 参数，该参数用于指定Object\[\]数组的长度；如果不指定 numElements 参数， Deque 底层数组的长度为16。

ArrayList 和 ArrayDeque 两个集合类的实现机制基本相似，它们的底层都采用一个动态的、可重分配的 Object \[\]数组来存储集合元素，当集合元素超出了该数组的容量时，系统会在底层重新分配一个 Object \[\]数组来存储集合元素.

ArrayDeque 当做栈使用代码：

```
public class ArrayDequeStack{
    public static void main(String[] args)
{
        ArrayDeque stack = new ArrayDeque();
        // 依次将三个元素push入"栈"
        stack.push("疯狂Java讲义");
        stack.push("轻量级Java EE企业应用实战");
        stack.push("疯狂Android讲义");
        // 输出：[疯狂Android讲义, 轻量级Java EE企业应用实战, 疯狂Java讲义]
        System.out.println(stack);
        // 访问第一个元素，但并不将其pop出"栈"，输出：疯狂Android讲义
        System.out.println(stack.peek());
        // 依然输出：[疯狂Android讲义, 疯狂Java讲义, 轻量级Java EE企业应用实战]
        System.out.println(stack);
        // pop出第一个元素，输出：疯狂Android讲义
        System.out.println(stack.pop());
        // 输出：[轻量级Java EE企业应用实战, 疯狂Java讲义]
        System.out.println(stack);
    }
}
```

ArrayDeque 当做队列使用代码：

```
public class ArrayDequeQueue{
    public static void main(String[] args){
        ArrayDeque queue = new ArrayDeque();
        // 依次将三个元素加入队列
        queue.offer("疯狂Java讲义");
        queue.offer("轻量级Java EE企业应用实战");
        queue.offer("疯狂Android讲义");
        // 输出：[疯狂Java讲义, 轻量级Java EE企业应用实战, 疯狂Android讲义]
        System.out.println(queue);
        // 访问队列头部的元素，但并不将其poll出队列"栈"，输出：疯狂Java讲义
        System.out.println(queue.peek());
        // 依然输出：[疯狂Java讲义, 轻量级Java EE企业应用实战, 疯狂Android讲义]
        System.out.println(queue);
        // poll出第一个元素，输出：疯狂Java讲义
        System.out.println(queue.poll());
        // 输出：[轻量级Java EE企业应用实战, 疯狂Android讲义]
        System.out.println(queue);
    }
}
```

6.3LinkedList  实现类

LinkedList类是List接口的实现，它是—个List集合，可以根据索引来随机访问集合中的元素。除此之外，LinkedList还实现Deque接口，可以被当成双端队列来使用，因为既可以被当成“栈”来使用，也可以当成队列使用。

```
public class LinkedListTest{
    public static void main(String[] args){
        LinkedList books = new LinkedList();
        // 将字符串元素加入队列的尾部
        books.offer("疯狂Java讲义");
        // 将一个字符串元素加入栈的顶部
        books.push("轻量级Java EE企业应用实战");
        // 将字符串元素添加到队列的头部（相当于栈的顶部）
        books.offerFirst("疯狂Android讲义");
        // 以List的方式（按索引访问的方式）来遍历集合元素
        for (int i = 0; i < books.size() ; i++ )
        {
            System.out.println("遍历中：" + books.get(i));
        }
        // 访问、并不删除栈顶的元素
        System.out.println(books.peekFirst());
        // 访问、并不删除队列的最后一个元素
        System.out.println(books.peekLast());
        // 将栈顶的元素弹出“栈”
        System.out.println(books.pop());
        // 下面输出将看到队列中第一个元素被删除
        System.out.println(books);
        // 访问、并删除队列的最后一个元素
        System.out.println(books.pollLast());
        // 下面输出：[轻量级Java EE企业应用实战]
        System.out.println(books);
    }
}
```

LinkedList与ArrayList 、 ArrayDeque 的实现机制完全不同， ArrayList 、 ArrayDeque 内部以数组的形式保存集合中的元素，因此随机访问集合元素时有较好的性能；而 LinkedList 内部以链表的形式来保存集合中的元素，因此随机访问集合元索时性能较差，但在插入、删除元索时性能比较出色（只需改变指针所指的地址即可）。需要指出的是，虽然 Vector也是以数组的形式来存储集合元素的，但因为它实现了线程冋步功能（而且实现机制也不好），所以各方面性能都比较差。

对于所有的内部基于数组的集合实现，例如 ArrayList 、 ArrayDeque 等，使用随机访问的性能比使用Iterator 迭代访问的性能要好，因为随机访问会被映射成对数组元素的访问.

6.4 各种线性表的性能分析

Java提供的List就是一个线性表接口，而ArrayList 、LinkedList 又是线性表的两种典型实现：基于数组的线性表和基于链的线性表。Queue代表了队列，Deque 代表了双端队列（既可作为队列使用，也可作为栈使用），

初学者吋以无须理会 ArrayList 和 LinkedList 之间的性能差异，只需要知道 LinkedList 集合不仅提供了List的功能，还提供了双端队列、栈的功能就行。但对于一个成熟的 Java 程序员，在一些性能非常敏感的地方,可能需要慎重选择哪个 List 实现。

一般来说，由于数组以一块连续内存区来保存所有的数组元素，所以数组在随机访问时性能最好，所有的内部以数组作为底层实现的集合在随机访问时性能都比较好：而内部以链表作为底层实现的集合在执行插入、删除操作时有较好的性能。但总体来说， ArrayList 的性能比 LinkedList 的性能要好，因此大部分时候都应该考虑使 ArrayList。

关于使用List集合有如下建议：

\(1\)如果需要遍历List集合元素，对于 ArrayList 、 Vector 集合，应该使用随机访问方法（get\)来遍历集合元素，这样性能更好 ： 对于 LinkedList 集合，则应该采用迭代器（ Iterator \)来遍历集合元素。

\(2\)如果需要插入 、删除操作来改变包含大量数据的 List 集合的大小 ，可考虑使用LinkedList 集合。使用 ArrayList 、 Vector 集合吋能耑耍经常重新分配内部数组的人小，效果可能较差。

\(3\)如果有多个线程盖要同时访问 List 集合中的元索，开发者可考虑使用 Collections 将集合包装成线稈安全的集合。

七、Map集合

Map 用于保存具有映射关系的数据，因此 Map 集合里保存着两组值，一组值用于保存Map里的key ,另外一组值用于保存 Map里的value，key和value都可以是任何引用类型的数据。 Map的key不允许重复，即同一个Map对象的任何两个key通过equals方法比较总是返回false。key和value之间存在单向一对一关系，即通过指定的 key ,总能找到唯一的、确定的 value 。从 Map 中取出数据时，只要给出指定的 key ,就可以取出对应的value 。如果把 Map 的两组值拆开来看，Map里的数据结构如图。

![](/assets/Map.png)

Map基本使用：

```
public class MapTest{
    public static void main(String[] args){
        Map map = new HashMap();
        // 成对放入多个key-value对
        map.put("疯狂Java讲义" , 109);
        map.put("疯狂iOS讲义" , 10);
        map.put("疯狂Ajax讲义" , 79);
        // 多次放入的key-value对中value可以重复
        map.put("轻量级Java EE企业应用实战" , 99);
        // 放入重复的key时，新的value会覆盖原有的value
        // 如果新的value覆盖了原有的value，该方法返回被覆盖的value
        System.out.println(map.put("疯狂iOS讲义" , 99)); // 输出10
        System.out.println(map); // 输出的Map集合包含4个key-value对
        // 判断是否包含指定key
        System.out.println("是否包含值为 疯狂iOS讲义 key："
            + map.containsKey("疯狂iOS讲义")); // 输出true
        // 判断是否包含指定value
        System.out.println("是否包含值为 99 value："
            + map.containsValue(99)); // 输出true
        // 获取Map集合的所有key组成的集合，通过遍历key来实现遍历所有key-value对
        for (Object key : map.keySet() )
        {
            // map.get(key)方法获取指定key对应的value
            System.out.println(key + "-->" + map.get(key));
        }
        map.remove("疯狂Ajax讲义"); // 根据key来删除key-value对。
        System.out.println(map); // 输出结果不再包含 疯狂Ajax讲义=79 的key-value对
    }
}
```

7.1HashMap和HashTable

HashMap和HashTable都是Map接口的典型实现类。两者的区别如下：

\(1\)Hashtable 是一个线程安全的Map实现，但HashMap 是线程不安全的实现，所以HashMap比Hashtable 的性能髙一点；但如果有多个线程访问同一个 Map 对象时，使用Hashtable实现类会更好。

\(2\)Hashtable 不允许使用null作为 key 和 value, 如果试阁把null值放进Hashtable中，将会引发NullPointerException异常：但HashMap可以使用null作为key或value 。

HashMap和HashTable 不能保证其中key-value对的顺序。它们判断两个Key相等的标准是：两个Key通过equals\(\)方法比较返回true，

两个Key的hashCode值也相等。

HashMap和HashTable包含一个containsValue\(\)方法，用于判断是否包含指定的value。

HashMap和HashTable判断两个value是否相等的方法是：只有两个对象通过equals\(\)方法比较返回true即可。

```
class A{
    int count;
    public A(int count)
    {
        this.count = count;
    }
    // 根据count的值来判断两个对象是否相等。
    public boolean equals(Object obj)
    {
        if (obj == this)
            return true;
        if (obj != null && obj.getClass() == A.class)
        {
            A a = (A)obj;
            return this.count == a.count;
        }
        return false;
    }
    // 根据count来计算hashCode值。
    public int hashCode()
    {
        return this.count;
    }
}
class B{
    // 重写equals()方法，B对象与任何对象通过equals()方法比较都返回true
    public boolean equals(Object obj)
    {
        return true;
    }
}
public class HashtableTest
{
    public static void main(String[] args)
    {
        Hashtable ht = new Hashtable();
        ht.put(new A(60000) , "疯狂Java讲义");
        ht.put(new A(87563) , "轻量级Java EE企业应用实战");
        ht.put(new A(1232) , new B());
        System.out.println(ht);
        // 只要两个对象通过equals比较返回true，
        // Hashtable就认为它们是相等的value。
        // 由于Hashtable中有一个B对象，
        // 它与任何对象通过equals比较都相等，所以下面输出true。
        System.out.println(ht.containsValue("测试字符串")); // ① 输出true
        // 只要两个A对象的count相等，它们通过equals比较返回true，且hashCode相等
        // Hashtable即认为它们是相同的key，所以下面输出true。
        System.out.println(ht.containsKey(new A(87563)));   // ② 输出true
        // 下面语句可以删除最后一个key-value对
        ht.remove(new A(1232));    //③
        System.out.println(ht);
    }
}
```

7.2LinkedHashMap

LinkedHashMap是HashMap的子类。LinkedHashMap 也使用双向链表来维护 key-value 对的次序（其实只需要考虑 key 的次序），该链表负责维护Map的迭代顺序，迭代顺序与key-value 对的插入顺序保持一致。

LinkedHashMap可以避免对HashMap、Hashtable里的 key-value 对进行排序（只要插入key-value对保持顺序即可），同时又避免使用 TreeMap 所增加的成本。

LinkedHashMap 需要维护元索的插入顺序， 因此性能略低于HashMap 的性能：但因为它以链表来维护内部顺序，所以在迭代访问Map里的全部元素时将有较好的性能。

```
public class LinkedHashMapTest
{
    public static void main(String[] args)
    {
        LinkedHashMap scores = new LinkedHashMap();
        scores.put("语文" , 80);
        scores.put("英文" , 82);
        scores.put("数学" , 76);
        // 调用forEach方法遍历scores里的所有key-value对
        scores.forEach((key, value) -> System.out.println(key + "-->" + value));
    }
}
```

7.3使用Properties读写属性文件

Properties类是HashTable的子类，用于处理属性文件（Windows操作台上的ini文件就是一种属性文件）。

Properties 类可以把Map对象和属性文件关联起来，从而可以把 Map对象中的 key-value 对写入属性文件中，也可以把属性文件中的“属性名 = 属性值 ”

加加载到 Map 对象中。由于属性文件里的属性名、 属性值只能是字符串类型，所以Properties里的key、value 都是字符中类型。该类提供了如下三个方法来修改 Properties 里的 key、value 值。

\(1\)String getProperty\(String key\): 获取Properties中指定属性名对应的属性值，类似于Map的get&lt;Object key\) 方法。

\(2\)String getProperty\(String key, String defauhValue\): 该方法与前一个方法基本相似。该方法多一个功能，如果 Properties 中存在指定的key时，则该方法指定默认值。

\(3\)Object setProperty\(String key, String value\): 设置属性值，类似于Hashtable的put\(\)方法。

除此之外，它还提供了两个读写属性文件的方法:

\(1\)void load \(InputStream inStream\): 从属性文件（以输入流表示）中加载 key-value 对，把加载到的key-value 对追加到 Properties 里（ Properties 是 Hashtable 的子类，它不保证 key-value 对之间的次序\)。

\(2\)void store\(OutputStream out， String comments \):将 Properties 中的 key-value 对输出到指定的属性文件（以输出流表示）中。

```
public class PropertiesTest
{
    public static void main(String[] args)
        throws Exception
    {
        Properties props = new Properties();
        // 向Properties中增加属性
        props.setProperty("username" , "yeeku");
        props.setProperty("password" , "123456");
        // 将Properties中的key-value对保存到a.ini文件中
        props.store(new FileOutputStream("a.ini")
            , "comment line");   //①
        // 新建一个Properties对象
        Properties props2 = new Properties();
        // 向Properties中增加属性
        props2.setProperty("gender" , "male");
        // 将a.ini文件中的key-value对追加到props2中
        props2.load(new FileInputStream("a.ini") );   //②
        System.out.println(props2);
    }
}
```

7.4WeakHashMap实现类

WeakHashMap与 HashMap的用法基本相似。与HashMap的区别在于，HashMap的key保留了对实际对象的强引用,这意味着只要该 HashMap 对象不被销毁，该HashMap的所有key所引用的对象就不会被垃圾回收， HashMap 也不会自动删除这些 key 所对应的 key-value 对；但 WeakHashMap 的 key只保留对实际对象的弱引用，这意味着如果 WeakHashMap 对象的 key 所引用的对象没有被其他强引用变量所引用，则这些 key 所引用的对象 可能被垃圾回收 ，WeakHashMap 也可能自动删除这些 key 所

对应的key-value对。WeakHashMap中的每个key对象只持有对实际对象的弱引用，因此，当垃圾回收了该key所对应的实际对象之后， WeakHashMap 会自动删除该key对应的 key-value对。

```
public class WeakHashMapTest
{
    public static void main(String[] args)
    {
        WeakHashMap whm = new WeakHashMap();
        // 将WeakHashMap中添加三个key-value对，
        // 三个key都是匿名字符串对象（没有其他引用）
        whm.put(new String("语文") , new String("良好"));
        whm.put(new String("数学") , new String("及格"));
        whm.put(new String("英文") , new String("中等"));
        //将 WeakHashMap中添加一个key-value对，
        // 该key是一个系统缓存的字符串对象。
        whm.put("java" , new String("中等"));    // ①
        // 输出whm对象，将看到4个key-value对。
        System.out.println(whm);
        // 通知系统立即进行垃圾回收
        System.gc();
        System.runFinalization();
        // 通常情况下，将只看到一个key-value对。
        System.out.println(whm);
    }
}
```

7.5IdentityHashMap实现类

IdentityHashMap实现机制与HashMap基本类似。但在处理两个Key相等时比较独特。当且仅当两个Key严格相等\(key1=key2\)时,IdentityHashMap认为两个key相等。

```
public class IdentityHashMapTest
{
    public static void main(String[] args)
    {
        IdentityHashMap ihm = new IdentityHashMap();
        // 下面两行代码将会向IdentityHashMap对象中添加两个key-value对
        ihm.put(new String("语文") , 89);
        ihm.put(new String("语文") , 78);
        // 下面两行代码只会向IdentityHashMap对象中添加一个key-value对
        ihm.put("java" , 93);
        ihm.put("java" , 98);
        System.out.println(ihm);
    }
}
```

7.6EnumMap实现类

EnumMap中所有的key都必须是单个枚举类的枚举值。创建EnumMap时必须显式或隐式指定它对应的枚举类。

枚举类的特征：

\(1\)EnumMap在内部以数组形式保存，所以这种实现形式非常紧凑、高效。

\(2\)EnumMap根据key的自然顺序（即枚平值在枚举类中的定义顺序）来维护 key - value 对的顺序。当程序通过 keySet \(\)、 entrySet\(\)、 values \(\)等方法遍历 EnumMap 时可以看到这种顺序。

\(3\)EnumMap 不允许使用 null 作为 key ，但允许使用 null 作为 value 。如果试图使用 null 作为 key时将抛出 NuHPointerException 异常。如果只是查询是否包含值为null的 key ,或只是刪除值为null 的 key ，都不会抛出异常。

与创建普通的 Map有所区別的是，创建EnumMap 时必须指定一个枚举类，从而将该 EnumMap 和指定枚举类关联起来。

```
enum Season
{
    SPRING,SUMMER,FALL,WINTER
}
public class EnumMapTest
{
    public static void main(String[] args)
    {
        // 创建EnumMap对象，该EnumMap的所有key都是Season枚举类的枚举值
        EnumMap enumMap = new EnumMap(Season.class);
        enumMap.put(Season.SUMMER , "夏日炎炎");
        enumMap.put(Season.SPRING , "春暖花开");
        System.out.println(enumMap);
    }
}
```

7.7各种Map实现类的性能分析

对于Map的常用实现类而言，HashMap 和 Hashtable 的实现机制几乎一样，但由于 Hashtable

足一个古老的、线程安全的集合，因此 HashMap 通常比 Hashtable 要快。

TreeMap 通常比 HashMap 、 Hashtable 要慢（尤其在插入、删除 key-value 对时更慢），因为 TreeMap

底层采用红黑树来管理 key-value 对（红黑树的每个节点就足一个key-value对）。

使用 TreeMap 的一个好处： TreeMap 中的 key-value 对总是处于有序状态，无须专门进行排序操作。

当TreeMap被填充之后, 就可以调用KeySet\(\). 取得由key组成的 Set ，然后使用toArrayO 方法生成 key

的数组，接下来使用 Arrays 的 binarySearch（）方法在已排序的数组中快速地查询对象。

对对于一般的应用场景，应考虑使用 HashMap, 因为 HashMap 正是为快速查询设计的。

\(HashMap 底层其实也是采用数组来存储 key-value 对）。如果程序需要一个总是排好序的 Map 时，则

可以考虑使用 TreeMap 。

LinkedHashMap 比 HashMap 慢一点，因为它需要维护链表来保持Map中 key-value 时的添加顺序。

IdentityHashMap 性能没有特别出色之处，因为它采用与 HashMap 基本相似的实现，只是它使用==而不

是 equals（）方法判断元索相等。 EnumMap 的性能最好，但它只能使用同一个枚举类的枚举值作为 key 。

八、Collections操作集合的工具类

8.1排序操作

Collection提供了如下常用的类方法用于对List集合元素进行排序：

\(1\)void reverse\(Listlist\): 反转指定 List 集合中元素的顺序.

\(2\)void shuffle\(List list\): 对List集合元素进行随机排序\(shuffle 方法模拟了 “ 洗牌”动作\).

\(3\)void sort\(List list\) : 根据元素的自然顺序对指定 List 集合的元素 按升序进行排序，

\(4\)void sort\(List list，Comparator c \):根据指定 Comparator 产生的顺序对 List 集合元素进行排序。

\(5\)void swap\(List list ， int i ， int j\) : 将指定 List集合中的i处元素和j处元素进行交换。

\(6\)void rotate\(Listst list, int distance \):当distance为正数时，将List集合的后 distance个元素 “ 整体 ”

移到前面；当distance 为负数时，将 list 集合的前 distance 个元素 “ 整体 ” 移到后面。该方法不会改变集合的长度。

```
public class SortTest
{
    public static void main(String[] args)
    {
        ArrayList nums = new ArrayList();
        nums.add(2);
        nums.add(-5);
        nums.add(3);
        nums.add(0);
        System.out.println(nums); // 输出:[2, -5, 3, 0]
        Collections.reverse(nums); // 将List集合元素的次序反转
        System.out.println(nums); // 输出:[0, 3, -5, 2]
        Collections.sort(nums); // 将List集合元素的按自然顺序排序
        System.out.println(nums); // 输出:[-5, 0, 2, 3]
        Collections.shuffle(nums); // 将List集合元素的按随机顺序排序
        System.out.println(nums); // 每次输出的次序不固定
    }
}
```

8.2查找、替换操作

Collection提供了如下常用的类方法用于查找、替换集合元素的类方法：

\(1\)int binarySearch\(List list， Object key\):使用二分搜索法搜索指定的 List 集合，以获得指走对象在 List

集合中的索引。如果要使该方法可以正常工作，则必须保证 List 中的元素已经处于有序状态。

\(2\) Objectmax\(Collection coll\): 根据元素的自然顺序，返回给定集合中的最大元素，

\(3\) Objectmax\(Conection coll,Comparator comp } : 根据Comparator指定的顺序，返回给定集合中的最大元素。

\(4\) Object min\(Collection coll\): 根据元素的自然顺序，返回给定集合中的最小元素。

\(5\) Object min\(Collection coll， Comparator comp \) : 根 据 Comparator  指定的顺序，返回给定集合中的最小元素。

\(6\) void fill\(List list ,Object obj\(\): 使用指定元素 obj 替换指定 List 集合中的所有元素。

\(7\) int frequency\(Conection c, Object o\): 返回报定集合中指定元素的出现次数。

\(8\) int indexOfSubList\(List source，List target\): 返回子 List对象在父 List 对象中第一次出现的位置索

引：如果父 List 中没有出现这样的子 List , 则 返 回-1.

\(9\) int  lastIndexOfSubList\(List source, List target \) : 返回子 List  对象在父 List  对象中最后一次出现的

位置索引：如果父List中没有出现这样的子List , 则 返 回-1.

\(10\) boolean replaceAll\(List list, Object oldVal, Object newVal \):使用一个新值newVal好换List对象的

所而旧值oldVal 。

```
public class SearchTest
{
    public static void main(String[] args)
    {
        ArrayList nums = new ArrayList();
        nums.add(2);
        nums.add(-5);
        nums.add(3);
        nums.add(0);
        System.out.println(nums); // 输出:[2, -5, 3, 0]
        System.out.println(Collections.max(nums)); // 输出最大元素，将输出3
        System.out.println(Collections.min(nums)); // 输出最小元素，将输出-5
        Collections.replaceAll(nums , 0 , 1); // 将nums中的0使用1来代替
        System.out.println(nums); // 输出:[2, -5, 3, 1]
        // 判断-5在List集合中出现的次数，返回1
        System.out.println(Collections.frequency(nums , -5));
        Collections.sort(nums); // 对nums集合排序
        System.out.println(nums); // 输出:[-5, 1, 2, 3]
        //只有排序后的List集合才可用二分法查询，输出3
        System.out.println(Collections.binarySearch(nums , 3));
    }
}
```

8.3同步控制

Collection提供了多个synchronizedXxx\(\)方法。该方法可以将指定集合包装成线程同步的集合，从而解决多线程并发访问集合时的线程安全问题。

```
public class SynchronizedTest
{
    public static void main(String[] args)
    {
        // 下面程序创建了四个线程安全的集合对象
        Collection c = Collections
            .synchronizedCollection(new ArrayList());
        List list = Collections.synchronizedList(new ArrayList());
        Set s = Collections.synchronizedSet(new HashSet());
        Map m = Collections.synchronizedMap(new HashMap());
    }
}
```

8.4设置不可变集合

Collection提供了如下三类方法来返回一个不可变的集合。

Collections类中提供了多个synchronizedXxx \(\)方法，该方法吋以将指定集合装成线程同步的集合，

从而可以解决多线程并发访问集合时的线程安全问题。

Java 中常用的集合框架中的实现类HashSet 、TreeSct 、ArrayList 、ArrayDeque 、LinkedList、HashMap

和 TrceMap 都是线程不安全的 。如果有多个线程访问它们，而且有超过一个的线程试图修改它们，则

存在线程安全的问题。Collections 提供了多个类方法把它们包装成线程同步的集合。

```
public class UnmodifiableTest
{
    public static void main(String[] args)
    {
        // 创建一个空的、不可改变的List对象
        List unmodifiableList = Collections.emptyList();
        // 创建一个只有一个元素，且不可改变的Set对象
        Set unmodifiableSet = Collections.singleton("疯狂Java讲义");
        // 创建一个普通Map对象
        Map scores = new HashMap();
        scores.put("语文" , 80);
        scores.put("Java" , 82);
        // 返回普通Map对象对应的不可变版本
        Map unmodifiableMap = Collections.unmodifiableMap(scores);
        // 下面任意一行代码都将引发UnsupportedOperationException异常
        unmodifiableList.add("测试元素");   //①
        unmodifiableSet.add("测试元素");    //②
        unmodifiableMap.put("语文" , 90);   //③
    }
}
```



