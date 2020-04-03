Java提供了一种原程序中的元素关联任何信息和任何元数据的途径和方法,称为注解。

从JDK5开始，Java增加了对元数据（MetaData）的支持，也就是Annotation（即注解）。

一、基本注解（JDK内置系统注解）

JDK内置系统注解：

```
@Override
@Deprecated
@suppressWarings
@SafeVarargs(JDK7新增)
@FunctionalInterface（JDK8新增）
```

1.1限定重写父类方法：@Override

@Override 用来指定方法重载，它可以强制一个子类必须覆盖父类的方法。

```
public class Fruit
{
    public void info()
    {
        System.out.println("水果的info方法...");
    }
}
class Apple extends Fruit
{
    // 使用@Override指定下面方法必须重写父类方法
    @Override
    public void inf0()
    {
        System.out.println("苹果重写水果的info方法...");
    }
}
```

1.2标记已过时@Deprecated

@Deprecated用于表示某个程序元素（类、方法等）已过时，当其它程序使用已过时的类、方法时，编译器将会给出警告。

```
class Apple
{
    // 定义info方法已过时
    @Deprecated
    public void info()
    {
        System.out.println("Apple的info方法");
    }
}
public class DeprecatedTest
{
    public static void main(String[] args)
    {
        // 下面使用info方法时将会被编译器警告
        new Apple().info();
    }
}
```

1.3 抑制编译器警告@suppressWarings

@suppressWarings指示被该Annotation修饰的程序元素（以及该程序元素中的所有子元素）取消显示指定的编译器警告。

@suppressWarings会一直作用于该程序元素的所有子元素。

```
// 关闭整个类里的编译器警告
@SuppressWarnings(value="unchecked")
public class SuppressWarningsTest
{
    public static void main(String[] args)
    {
        List<String> myList = new ArrayList();     // ①
    }
}
```

1.4 Java7的”堆污染“警告@SafeVarargs

由于泛型擦除，当把一个不带泛型的对象赋给一个带泛型的变量时，会导致运行时异常，Java中把引起这种错误的原因称为“堆污染”。

```
public class ErrorUtils
{
    @SafeVarargs
    public static void faultyMethod(List<String>... listStrArray)
    {
        // Java语言不允许创建泛型数组，因此listArray只能被当成List[]处理
        // 此时相当于把List<String>赋给了List，已经发生了“擦除”
        List[] listArray = listStrArray;
        List<Integer> myList = new ArrayList<Integer>();
        myList.add(new Random().nextInt(100));
        // 把listArray的第一个元素赋为myList
        listArray[0] = myList;
        String s = listStrArray[0].get(0);
    }
}
```

Java 7使用如下三种方式来“抑制”“堆污染”警告。

\(1\)使用@SafeVarargs 修饰引发该警告的方法或构造器。

\(2\)使用@SuppressWamings \(" unchecked "\)修饰。

\(3\)编译时使用- Xlint:varargs 选项。

很明显，第三种方式一般比较少用，通常可以选择第一种或第二种方式，尤其是使用 @SafeVarargs

修饰引发该警告的方法或构造器，它是Java7专门为抑制“堆污染”警告提供的。

1.5@FunctionalInterface

@FunctionalInterface用来指定某个接口必须是函数式接口。

Java8规定：如果接口中只有一个抽象方法（可以包含多个默认方法或多个static方法），该接口就是函数式接口

```
@FunctionalInterface
public interface FunInterface
{
    static void foo()
    {
        System.out.println("foo类方法");
    }
    default void bar()
    {
        System.out.println("bar默认方法");
    }
    void test(); // 只定义一个抽象方法

    void abc();
}
```

1.6Java8新增的重复注解

在Java8以前，同一个程序元素前最多只能使用一个相同类型的Annotation，如果需要在同一个元素前  
使用多个相同类型的 Annotation ,则必须使用 Annotation “容器”。  
Java 8允许使用多个相同类型的 Annotation 来修饰同一个类。

```
// 指定该注解信息会保留到运行时
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.TYPE)
public @interface FkTags
{
    // 定义value成员变量，该成员变量可接受多个@FkTag注解
    FkTag[] value();
}
```

1.7Java8新增的类型注解Type Annotation

Java 8为ElementType枚举增加了 TYPE\_PARAMETER、TYPE\_USE两个枚举值，这样就允许定义枚举时使用@Target\(ElementType.TYPE\_USE\)修饰，这种注解被称为TypeAnnotation \(类型注解）,TypeAnnotation可用在任何用到类型的地方。

从Java8开始,Type Annotation可以在任何用到类型的地方使用.比如，允许在如下位置使用 Type Annotation。

\(1\)创违对象\(用new关键字创建\).

\(2\)类型转换。

\(3\)使用implements实现接口.

\(4\)伸用throws 声明抛出异常。

```
// 定义一个简单的Type Annotation，不带任何成员变量
@Target(ElementType.TYPE_USE)
@interface NotNull{}
// 定义类时使用Type Annotation
@NotNull
public class TypeAnnotationTest
    implements @NotNull /* implements时使用Type Annotation */ Serializable
{
    // 方法形参中使用Type Annotation
    public static void main(@NotNull String[] args)
        // throws时使用Type Annotation
        throws @NotNull FileNotFoundException
    {
        Object obj = "fkjava.org";
        // 强制类型转换时使用Type Annotation
        String str = (@NotNull String)obj;
        // 创建对象时使用Type Annotation
        Object win = new @NotNull JFrame("疯狂软件");
    }
    // 泛型中使用Type Annotation
    public void foo(List<@NotNull String> info){}
}
```

二、Java的元注解

JDK在java.lang.annotation包下提供了6个Meta Annotion\(元Annotation\)。除了@Repeation专门用于定义Java8新增的重复注解外，其它都用于修饰其他的Annotation定义。

1、@Retention\(\)

@Retention\(\)只能用于修饰Annotation定义，用于指定被修饰的Annotation可以保留多长时间，@Retention\(\)包含一个RetentionPolicy类型的value成员变量，所以使用@Retention\(\)时必须为该value或成员变量指定值。

\(1\)RetentionPolicy.CLASS   : 编译时会把Annotation记录到class中，运行时忽略；这是默认值

\(2\)RetentionPolicy.RUNTIME : 运行时存在，可以通过反射读取。

\(3\)RetentionPolicy.SOURCE :  Annotion只在源码显示，编译时会丢弃。

2、@Target

@Target只能用于修饰Annotation定义，用于指定被修饰的Annotation能够用于哪些程序单位，@Target包含一个value成员变量。

\(1\)ElementType.ANNOTATION\_TYPE:只修饰Annotion

\(2\)ElementType.CONSTRUCTOR:只修饰构造器

\(3\)ElementType.FIELD:只修饰成员变量

\(4\)ElementType.LOCAL\_VARIABLE:只修饰局部变量

\(5\)ElementType.METHOD:只修饰方法定义

\(6\)ElementType.PACKAGE:只修饰包定义

\(7\)ElementType.PARAMETER:只修饰参数

\(8\)ElementType.RYPE:可修饰类、接口（包括注解类型）或枚举定义

3、@Documented

@Documented用于指定被修饰的Annotation修饰的Annotation类将被Javadoc工具提取成文档。

```
@Retention(RetentionPolicy.RUNTIME)
@Target(ElementType.METHOD)
// 定义Testable Annotation将被javadoc工具提取
@Documented
public @interface Testable
{
}
```

```
public class MyTest
{
    // 使用@Test修饰info方法
    @Testable
    public void info()
    {
        System.out.println("info方法...");
    }
}
```

4、@Inherited

@Inherited 用于指定被修饰的Annotation将具有继承性，允许子类继承。

```
@Target(ElementType.TYPE)
@Retention(RetentionPolicy.RUNTIME)
@Inherited
public @interface Inheritable
{
}
```

```
// 使用@Inheritable修饰的Base类
@Inheritable
class Base
{
}
// TestInheritable类只是继承了Base类，
// 并未直接使用@Inheritable Annotiation修饰
public class InheritableTest extends Base
{
    public static void main(String[] args)
    {
        // 打印TestInheritable类是否具有@Inheritable修饰
        System.out.println(InheritableTest.class
            .isAnnotationPresent(Inheritable.class));
    }
}
```

三、自定义Annotation

1.定义Annotation

使用@interface关键字定义注解。可以修饰类、方法、变量、接口等。

通常把Annotation另放一行。

2.提取Annotation信息

使用Annotation修饰了类、方法、成员变虽等成员之后，这些 Annotation 不会自己生效，必须由开  
发者提供相应的工具来提取并处理 Annotation 信息。

Java使用 Annotation接口来代表程序元素前面的注解，该接口是所有注解的父接口。Java5在  
java.lang.reflect 包下新增了 AnnotatedElement接口 ，该接口代表程序中可以接受注解的程序元素。该接口主要有如下几个实现类：

```
(1)Class: 类定义。
(2)Constructor: 构造器定义.
(3)Field: 类的成员变请定义。
(4)Method: 类的方法定义。
(5)Package: 类的包定义。
```

java.lang.reflect 包下主要包含一些实现反射功能的类 ，从Java5开始 ，java.lang.reflect包所提供的反射API增加了读取运行时 Annotation的能力.只有当定义Annotation时使用了@Retention\(RetentionPolicy.RUNTIME\)修饰，该Annotation才会在运行时可见 ,JVM才会在装载 \*.class文件时读取保存在class文件中的 Annotation 。

AnnotatedElement接口是所有程序元素（如Class、Method 、Constructor 等）的父接口，所以程序  
通过反射获取了某个类的 AnnolatedElement 对象（如 Class 、 Method 、 Constructor 等）之后，程序就可  
以调用该对象的如下几个方法来访问 Annotation 信息。

\(1\)&lt;A extends Annotation〉 A getAnnotation\(Class&lt;A&gt; annotationClass \):返回该程序元索上存在的、  
指定类型的注解，如果该类型的注解不存在，则返回 null 。

\(2\)&lt;A extends Annotation〉 A getDeclaredAnnotation\(Class&lt;A&gt; annotationClass \):这垃 Java8 新增的方  
法，该方法尝试获取直接修饰该程序元索、指定类型的 Amiotation. 如果该类型的注解不存在，  
则返回null 。

\(3\)Annolation\[\]getAnnotations\(\) :返回该程序元索上存在的所有注解。

\(4\) Annotation\[\] getDeclaredAnnotations \(\):返回直接修饰该程序元素的所有 Annotation 。

\(5\) boolean isAnnotationPresent\(Class&lt;? extends Annotation 〉 annotationClass \):判断该程序元素上是否  
存在指定类型的注解，如果存在则返回 true, 否则返回 false 。

\(6\) &lt;A extends Annotation &gt; A\[\] getAnnotationsByType\(Class&lt;A\(\)annotationClass \):该方法的功能与前  
面介绍的 getAnnotion\(\)方法基本相似。但由于Java8增加了重复注解功能，因此需要使用该方  
法获取修饰该程序元素、指定类型的多个 Annotation 。

\(7\)&lt;A extends Annotation &gt;A\[\] getDeclaredAnnolationsByType\(Class&lt;A&gt; annotaUonClass &gt;:该方法的  
功能与前面介绍的 getDeclaredAnnotationsO 方法基本相似。但由于Java8增加了重复注解功能，  
因此需要使用该方法获取直接修饰该程序元素、指定类型的多个 Annotation.

