一、Java引入泛型的目的

Java集合有个缺点,把一个对象“丢进”集合里之后，集合就会“忘记”这个对象的数据类型，  
当再次取出该对象时，该对象的编译类型就变成了 Object类型（其运行时类型没变\)。

Java集合之所以被设计成这样，是因为集合的设计者不知道我们会用集合来保存什么类型的对象，  
所以他们把集合设计成能保存任何类型的对象，只要求具冇很好的通用性。但这样做带来如下两个问题：

1、集合对元素类型没有任何限制，这样叫能引发一些M题。例如，想创建一个只能保存Dog对象  
的集合，但程序也吋以轻姑地将Cat对象“丢”进去，所以可能引发异常。

2、由于把对象“丢进”集合时，集合丢失了对象的状态信息，集合只知道它盛装的是Object，因  
此取出集合元素后通常还需要进行强制类型转换。这种强制类型转换既增加了编程的复杂度，  
也可能引发ClassCastException异常。

增加了泛型支持后的集合，完全可以记住集合中元素的类型，并可以在编译时检查集合中元素的类型，如果试图向集合中添加不满足类型要求的对象，编译器就会提示错误。增加泛型后的集合，可以让  
代码更加简洁，程序更加健壮（Java泛型.可以保证如果程序在编译时没冇发出赘告，运行时就不会产生  
ClassCastException异常）。

二、泛型的使用

1、泛型的定义

在JDK 1.5版本以后，提出了泛型机制，允许程序在创建集合时指定集合元素的类型。

其语法如下：

```
类名<T>
其中，T代表一个类型的名称。
```

Java的参数化类型成为泛型。

所谓泛型，就是允许在定义类、接口、方法时使用类型形参，这个类型形参在声明变量、创建对象、调用方法时动态的指定（即传入实际的类型参数，也可称为类型实参）。

2、Java7以后的“菱形”语法

JDK7以后，Java允许在构造器之后不需要带完整的泛型信息，这种语法也被称为“菱形”语法（把两个尖括号\(&lt;&gt;\)并排放在一起非常想一个“菱形”）。

```
public static void main(String[] args)
    {
        // Java自动推断出ArrayList的<>里应该是String
        List<String> books = new ArrayList<>();
        books.add("疯狂Java讲义");
        books.add("疯狂Android讲义");
        // 遍历books集合，集合元素就是String类型
        books.forEach(ele -> System.out.println(ele.length()));
        // Java自动推断出HashMap的<>里应该是String , List<String>
        Map<String , List<String>> schoolsInfo = new HashMap<>();
        // Java自动推断出ArrayList的<>里应该是String
        List<String> schools = new ArrayList<>();
        schools.add("斜月三星洞");
        schools.add("西天取经路");
        schoolsInfo.put("孙悟空" , schools);
        // 遍历Map时，Map的key是String类型，value是List<String>类型
        schoolsInfo.forEach((key , value) -> System.out.println(key + "-->" + value));
    }
```

3、定义泛型接口、类

泛型的实质：允许在定义接口、类时声明类型形参，类型形参在整个接口、类体内可以当成类型使用，几乎所有可使用普通类型的地方都可以使用这种类型形参。

```
// 定义Apple类时使用了泛型声明
public class Apple<T>
{
    // 使用T类型形参定义实例变量
    private T info;
    public Apple(){}
    // 下面方法中使用T类型形参来定义构造器
    public Apple(T info)
    {
        this.info = info;
    }
    public void setInfo(T info)
    {
        this.info = info;
    }
    public T getInfo()
    {
        return this.info;
    }
    public static void main(String[] args)
    {
        // 由于传给T形参的是String，所以构造器参数只能是String
        Apple<String> a1 = new Apple<>("苹果");
        System.out.println(a1.getInfo());
        // 由于传给T形参的是Double，所以构造器参数只能是Double或double
        Apple<Double> a2 = new Apple<>(5.67);
        System.out.println(a2.getInfo());
    }
}
```

4、从泛型类派生子类

当创建了带泛型声明的接口、父类之后，可以为该接口创建实现类，或从该父类派生了子类，需耍需要指  
出的是，当使用这些接口父类时不能再包含类型形参，要传入具体的实参。

```
// 定义类A继承Apple类，Apple类不能跟类型形参  
public class A extends Apple<String>{ }
public class A extends Apple<T>{ }  //语法错误
```

三、类型通配符

1、定义

类型通配符是一个问号（？），将一个问号作为类型实参传给 List 集合，写作：List &lt;?&gt; \(意思趄元索类型未知的 List \)。这个问号（？）被称为通配符，它的元素类型可以匹配任何类型。

2、设定类型通配符的上限

程序不希望这个List &lt;?&gt; 是任何泛型的List的父类，只代表某一类泛型的父类。通过使用”？extend 具体类“来限制通配符的上限。

```
List <? extend Shape>
```

3、设定类型形参的上限

Java泛型不仅允许在使用通配符形参时设定上限，而且可以在定义类型形参时设定上限，用于表示传给该类型形参的实际类型要么是该上限类型，要么是该上限类型的子类。

```
public class Apple<T extends Number>
{
    T col;
    public static void main(String[] args)
    {
        Apple<Integer> ai = new Apple<>();
        Apple<Double> ad = new Apple<>();
        // 下面代码将引起编译异常，下面代码试图把String类型传给T形参
        // 但String不是Number的子类型，所以引发编译错误
        Apple<String> as = new Apple<>();        // ①
    }
}
```

四、泛型方法

1、定义泛型方法

泛型方法的格式：

```
[访问权限]<泛型>  返回类型 方法名([泛型标识 参数名称]) 抛出的异常
public class DAO {
public<E>  E get(int id, E e){
E result = null;
return result;
}}
```

```
public class RightTest
{
	// 声明一个泛型方法，该泛型方法中带一个T形参
	static <T> void test(Collection<? extends T> from , Collection<T> to)
	{
		for (T ele : from)
		{
			to.add(ele);
		}
	}
	public static void main(String[] args)
	{
		List<Object> ao = new ArrayList<>();
		List<String> as = new ArrayList<>();
		// 下面代码完全正常
		test(as , ao);
	}
}
```

泛型方法允许类型形参被用来表示方法的一个或多个参数之间的类型依赖关系，或者方法返回值与参数之间的类型依赖关系。如果没有这样的类型依赖关系，就不应该使用泛型方法。

2、泛型方法与类型通配符的区别

\(1\)通配符被设计用来支持灵活的子类化的。

\(2\)类型通配符既可以在方法签名中定义形参的类型，也可以用来定义变量的类型；但泛型方法中的类型形参必须在对应方法中显示声明。

3、设定通配符下限

为了表示两个参数至今的类型依赖，同时使用通配符、泛型参数来实现该方法。为了表达这种约束关系，Java允许设定通配符的下限：&lt;? super Type&gt;,这个通配符表示它必须是Type本身，或是Type的父类。

```
public class MyUtils
{
	// 下面dest集合元素类型必须与src集合元素类型相同，或是其父类
	public static <T> T copy(Collection<? super T> dest
		, Collection<T> src)
	{
		T last = null;
		for (T ele  : src)
		{
			last = ele;
			dest.add(ele);
		}
		return last;
	}
	public static void main(String[] args)
	{
		List<Number> ln = new ArrayList<>();
		List<Integer> li = new ArrayList<>();
		li.add(5);
		// 此处可准确的知道最后一个被复制的元素是Integer类型
		// 与src集合元素的类型相同
		Integer last = copy(ln , li);    // ①
		System.out.println(ln);
	}
}
```

Java集合框架中TreeSet&lt;E&gt;有一个构造器也用到了这种设定通配符下限的语法

```
//E是定义TreeSet类市的参数类型
TreeSet (Comparator <? super E > c)
```

TreeSet 会对集合中的元素按自然顺序或定制顺序进行排序。如果需要 TreeSet对集合中的所冇元素进行定制排序，则要求 TreeSet 对象有一个与之关联的 Comparator 对象。构造器中的参数 c 就是进行定制排序的 Comparatoi •对象。

Comparator 接口也是一个带泛型声明的接口：

```
public interface Comparator<T>
{
int compare(T fst, T snd);
)
```

通过这种带下限的通配符的语法，可以在创建 TreeSet 对象时灵活地选择合适的 Comparator 。假定需要创建一个 TreeSet &lt; String &gt;集合，并传入一个可以比较 String 大小的 Comparator ,这个 Comparator既可以是 Comparator &lt; String &gt;,也可以是 Comparator &lt; Object &gt;——只要尖括号里传入的类型是 String 的父类型（或它本身）即可。

```
public class TreeSetTest
{
	public static void main(String[] args)
	{
		// Comparator的实际类型是TreeSet的元素类型的父类，满足要求
		TreeSet<String> ts1 = new TreeSet<>(
			new Comparator<Object>()
		{
			public int compare(Object fst, Object snd)
			{
				return hashCode() > snd.hashCode() ? 1
					: hashCode() < snd.hashCode() ? -1 : 0;
			}
		});
		ts1.add("hello");
		ts1.add("wa");
		// Comparator的实际类型是TreeSet元素的类型，满足要求
		TreeSet<String> ts2 = new TreeSet<>(
			new Comparator<String>()
		{
			public int compare(String first, String second)
			{
				return first.length() > second.length() ? -1
					: first.length() < second.length() ? 1 : 0;
			}
		});
		ts2.add("hello");
		ts2.add("wa");
		System.out.println(ts1);
		System.out.println(ts2);
	}
}

```

4、泛型方法与方法重载

因为泛型既允许设定通配符的上限，也允许设定通配符的下限，从而允许在一个类里包含如下两个方法的定义。

![](/assets/泛型方法重载.png)

五、擦除与转换

擦除：当把一个具有泛型信息的对象赋给另一个没有泛型信息的变量时，所有在尖括号之间的类型信息都将被扔掉。

转换：对泛型而言，可以直接把一个List对象赋给一个List&lt;String&gt;对象，编译器仅仅提示“未经检查的转换”。但是试图取出集合的元素时，可能会发生运行时异常。



