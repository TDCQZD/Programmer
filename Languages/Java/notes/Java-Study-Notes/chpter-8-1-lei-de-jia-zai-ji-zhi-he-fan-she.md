一、类的加载、连接和初始化

1.1JVM和类

当调用java 命令运行某个 Java 程序时，该命令将会启动一个 Java 虚拟机进程，不管该 Java 程序有多么复杂，该程序启动了多少个线程，它们都处于该 Java 虚拟机进程里。

同一个 JVM的所有线程、所有变最都处于同一个进程里，它们都使用该 JVM 进程的内存区。

当系统出现以下几种情况时 ， JVM 进程将被终止。

\(1\)程序运行到最后正常结朿。

\(2\)程序运行到使用 System.exit \(\)或 Runtime.getRuntime\(\).exit\(\)代码处结束程序。

\(3\)程序执行过程中遇到未捕获的异常或错误而结束。

\(4\)程序所在平台强制结束了JVM进程。

当Java程序运行结束后，JVM进程结束，该进程在内存中的状态将会丢失。

1.2类的加载

当稈序主动使用某个类时，如果该类还未被加载到内存中，则系统会通过加载、连接、初始化三个步骤来对该类进行初始化。如果没有意外,JVM将会连续完成这三个步骤，所以有时也把这三个步骤统称为类加载或类初始化。

类加栽指的是将类的class文件读入内存，并为之创建一个java.lang.class对象也就是说，当程序中使用任何类时，系统都会为之建立一个java.lang.class 对象。

类的加载由类加载器完成，类加载器通常由JVM提供，这些类加载器也是前面所有程序运行的基础， JVM 提供的这些类加载器通常被称为系统类加栽器。除此之外，开发者可以通过继承 ClassLoader基类来创建自己的类加载器。

通过使用不同的类加载器，可以从不同来源加载类的二进制数据，通常有如下几种来源。

\(1\)从本地文件系统加载class文件,这是前面绝大部分示例程序的类加载方式。

\(2\)从JAR包加载class文件，这种方式也是很常见的,JDBC编程时用到的数据库驱动类就放在JAR文件中,JVM可以从 JAR文件屮直接加载该class文件。

\(3\)通过网络加载 class 文件。

\(4\) 把 一 个Java源文件动态编译，并执行加载。

类加战器通常无须等到“首次使用”该类时才加载该类， Java 虚拟机规范允许系统预先加载某些类。

1.3类的连接

当类被加载之后，系统为之生成一个对应的Class对象，接着将会进入连接阶段，连接阶段负责把类的二进制数据合并到JRE中。类连接又可分为如下三个阶段。

\(1\)验证：验证阶段用于检验被加载的类是否有正确的内部结构，并和其他类协调一致。

\(2\)准备：类准备阶段则负责为类的类变量分配内存，并设H默认初始值。

\(3\)解析：将类的二进制数据中的符号引用替换成直接引用。

1.4类的初始化

在类的初始化阶段，虚拟机负责对类进行初始化，主要就是对类变最进行初始化。在Java类中对类变景指定初始值有两种方式：①声明类变量时指定初始值：②使用静态初始化块为类变量指定初始值。

JVM 初始化一个类包含如下儿个步骤。

\(1\)假如这个类还没有被加载和连接，则程序先加载并连接该类。

\(2\)假如该类的茛接父类还没有被初始化，则先初始化其直接父类。

\(3\)假如类中有初始化语句，则系统依次执行这些初始化语句。

当执行第2个步骤时，系统对直接父类的初始化步骤也遵循此步骤1~3;如果该直接父类又有直接父类，则系统再次重复这三个步骤来先初始化这个父类……依此类推，所以 JVM 谅先初始化的总是java.lang.Object 类。当程序主动使用任何一个类时，系统会保证该类以及所有父类（包括直接父类和间接父类）都会被初始化。

1.5类初始化的时机

当Java程序首次通过下而6种方式来使用某个类或接口时，系统就会初始化该类或接口。

\(1\)创建类的实例。为某个类创建实例的方式包栝：使用 new 操作符來创建实例，通过反射来创建实例，通过反序列化的方式来创建类实例。

\(2\)调用某个类的类方法（静态方法）。

\(3\)访问某个类或接口的类变量，或为该类变量赋值。

\(4\)使用反射方式来强制创建某个类或接口对应的java.lang.Class 对象。例如代码 ：

```
Class.forName("Person")
```

如果系统还未初始化 Person 类，则这行代码将会导致该person类被初始化，并返回 Person 类对应的 java . Iang.Class 对象。

\(5\)初始化某个类的子类。当初始化某个类的子类时，该子类的所有父类都会被初始化。

\(6\) 直接使用java.exe命令来运行某个主类。当运行某个主类时，程序会先初始化该主类。

除此之外，下面的几种情形需要特别指出

对于一个 final 型的类变量 ，如果该类变量的值在编译时就可以确定下来，那么这个类变量相当于“宏变量”。 Java 编译器会在编译时直接把这个类变量 出现的地方替换成它的值，因此即使程序使用该静态类变量，也不会导致该类的初始化。

```
class MyTest
{
    static
    {
        System.out.println("静态初始化块...");
    }
    // 使用一个字符串直接量为static final的类变量赋值
    static final String compileConstant = "疯狂Java讲义";
}
public class CompileConstantTest
{
    public static void main(String[] args)
    {
        // 访问、输出MyTest中的compileConstant类变量
        System.out.println(MyTest.compileConstant);   // ①
    }
}
```

反之，如果final 修饰的类变量的值不能在编译时确定下来，则必须等到运行时才可以确定类变量的值，如果通过该类来访问它的类变量，则会导致该类被初始化。

当ClassLoader类的LoadClass\(\)方法来加载某个类时，该方法只是加载该类，并不会执行该类的初始化，使用Class的forName\(\)静态方法才会  
导致强制初始化该类.

```
class Tester
{
    static
    {
        System.out.println("Tester类的静态初始化块...");
    }
}
public class ClassLoaderTest
{
    public static void main(String[] args)
        throws ClassNotFoundException
    {
        ClassLoader cl = ClassLoader.getSystemClassLoader();
        // 下面语句仅仅是加载Tester类
        cl.loadClass("Tester");
        System.out.println("系统加载Tester类");
        // 下面语句才会初始化Tester类
        Class.forName("Tester");
    }
}
```

二、类的加载器

类加载器负责将.class文件加载到内存中，并为之生成对应的java.lang.Class对象。

2.1类加载器简介

类加载器负责加载所有的类，系统为所有被栽入内存中的类生成一个java.lang.Class实例。一旦一  
个类被战入 JVM 中，同一个类就不会被再次载入了。

现在的问题是，怎么样才算“同一个类”？

正如一个对象有一个唯一的标识一样，一个载入 JVM 的炎也有一个唯一的标识。在 Java 中，一个  
类用其全限定类名（包括包名和类名）作为标识：在JVM中，一个类用其全限定类名和其类加载器  
作为唯一标识。例如，如果在 pg 的包中有一个名为 Person 的类，被类加栽器ClassLoader的实例 kl  
负责加载，则该Person类对应的Class对象在JVM中表示为（ Person 、 pg 、 kl \)。这意味若两个类加载  
器加载的同名类：（ Person 、 pg 、 kl \)和（ Person 、 pg 、 k2\)是不同的、它们所加载的类也是完全不同、  
互不兼容的。

![](/assets/JVM类加载器.png)

当 JVM 启动时，会形成由三个类加载器组成的初始类加载器层次结构。

Bootstrap ClassLoader :根类加载器。

Extension ClassLoader :扩展类加载器。

System ClassLoader :系统类加载器。

Bootstrap ClassLoader 被称为引导（也称为原始或根）类加载器，它负责加载 Java 的核心类。在Sun的JVM中，执行java. exe命令时，使用- Xbootclasspath 选项或使用- D 选项衍定sun.boot.class.path系统属性值可以指定加载附加的类。

根类加载器非常特殊，它并不是java.lang.ClassLoader的子类，而是有JVM自身实现的。

```
public class BootstrapTest
{
    public static void main(String[] args)
    {
        // 获取根类加载器所加载的全部URL数组
        URL[] urls = sun.misc.Launcher.
        getBootstrapClassPath().getURLs();
        // 遍历、输出根类加载器加载的全部URL
        for (int i = 0; i < urls.length; i++)
        {
            System.out.println(urls[i].toExternalForm());
        }
    }
}
```

Extension ClassLoader 被称为扩展类加载器，它负责加载JRE的扩展目录\(%JAVA\_HOME%jre/lib/ext或者由java.ext.dirs系统属性指定的目录\)中JAR包的类.通过这种方式，就可以为Java扩展核心类以外的新功能，只要把自己开发的类打包成JAR文件，然后放入JAVA\_HOME/jre/lib/ext路径即可。

System ClassLoader被称为系统（应用）类加载器它负责在JVM启动时加载来自java命令的-classpath选项、java.class.path系统属性，或CLASSPATH环境变量所指定的JAR包和类路径。程序可以通过ClassLoader的静态方法getSystemClassLoader\(\)来获取系统类加载器。如果没有特别指定，则用户自定义的类加载器都以类加载器作为父加载器。

2.2 类加载机制

JVM 的类加载机制主要有如下三种。

\(1\)全盘负责。所谓全盘负责，就是当一个类加载器负责加载某个 Class 时，该 Class 所依赖的和引用的其他Class 也将由该类加载器负责载入，除非显式使用另外一个类加载器来载入。

\(2\)父类委托。所谓父类委托，则是先让 parent \(父）类加载器试图加载该 Class ,只有在父类加载器无法加载该类时才尝试从己的类路径中加载该类。

\(3\)缓存机制，缓存机制将会保证所有加载过的 Class 都会被缓存，当程序中需要使用某个 Class 时，类加战器先从缓存区中搜寻该 Class 。只有当缓存区中不存在该Class对象时，系统才会读取该类对应的二进制数据，并将其转换成 Class 对象，存入缓存区中。这就是为什么修改了Class 后，必须重新启动JVM ,程序所做的修改才会生效的原因。

```
public class ClassLoaderPropTest
{
    public static void main(String[] args)
        throws IOException
    {
        // 获取系统类加载器
        ClassLoader systemLoader = ClassLoader.getSystemClassLoader();
        System.out.println("系统类加载器：" + systemLoader);
        /*
        获取系统类加载器的加载路径——通常由CLASSPATH环境变量指定
        如果操作系统没有指定CLASSPATH环境变量，默认以当前路径作为
        系统类加载器的加载路径
        */
        Enumeration<URL> em1 = systemLoader.getResources("");
        while(em1.hasMoreElements())
        {
            System.out.println(em1.nextElement());
        }
        // 获取系统类加载器的父类加载器：得到扩展类加载器
        ClassLoader extensionLader = systemLoader.getParent();
        System.out.println("扩展类加载器：" + extensionLader);
        System.out.println("扩展类加载器的加载路径："
            + System.getProperty("java.ext.dirs"));
        System.out.println("扩展类加载器的parent: "
            + extensionLader.getParent());
    }
}
```

类加载器加载 Class 大致要经过如下8个步骤。

\(1\)检测此Class是否载入过（即在缓存区中是否有此 Class \),如果有则直接进入第8步，否则接着执行第2步。

\(2\)如果父类加载器不存在（如果没有父类加载器，则要么 parent —定是根类加载器，要么本身就是根类加载器），则跳到第4步执行;

如果父类加载器存在，则接着执行第3步。

\(3\)请求使用父类加载器去载入目标类，如果成功载入则跳到第8步，否则接着执行第5步。

\(4\)请求使用根类加载器去载入目标类，如果成功载入则跳到第8步，杏则跳到第7步。

\(5\)当前类加载器尝试寻找 Class 文件（从与此 ClassLoader 相关的类路径中寻找），如果找到则执行笫6步，如果找不到则跳到第7步。

\(6\)从文件中战入 Class ,成功载入后跳到第8步。

\(7\)抛出ClassNotFoundException 异常。

\(8\)返回对应的java.lang.Class对象。其中，第5、6步允许重写ClassLoader的findClass\(\)方法来实现自己的载入策略，甚至重写loadClass\(\)方法来实现自己的载入过程。

2.3自定义类加载器

通过扩展ClassLoader的子类，并重写ClassLoader所包含的方法来实现自定义的类加载器。

ClassLoader类有如下两个关键方法。

\(1\)loadClass\(String name, boolean resolve\):该方法为 ClassLoader 的入口点，根据指定名称来加载类，系统就是调用ClassLoader的该方法来获取指定类对应的Class对象。

\(2\)findClass\(String name\):根据指定名称来。

建议重写findClass\(\)方法而不是loadClass\(\)方法。

loadClass\(\)方法的执行步骤如下:

\(1\)用findLoadedClass\(String\):来检查是否已经加载类，如果已经加载则直接返回.

\(2\)在父类加载器上调用loadClass\(\)方法。如果父类加载器为null,则使用根类加载器来加载。

\(3\)调用findClass\(String\)方法查找找类。

从上面步骤中可以看出，重写findClass\(\)方法法可以避免覆盖默认类加载器的父类委托、缓冲机制两种策略：如果重写loadClass\(\)方法，则实现逻辑更为复杂。

在 ClassLoader 里还有一个核心方法：Class defineClass\(String name，byte\[\] b, int off, int len\)，该方法复杂将指定类的字节码文件（即Class文件，如Hello.class\)读入字节数组byte\[\] b内，并把它转换为Class对象，该字节码文件可以来源于文件、网络等。

defineClassO方法管理JVM的许多复杂的实现，它负责将字节码分析成运行时数据结构，并校验有效性等。自定义加载器时不需要重写该方法。

2.4 URLCIassLoader类

Java为ClassLoader提供了一个URLCIassLoader实现类，该类也是系统类加载 器和扩展类加载器的父类（此处的父类，就是指类与类之间的继承关系）。URLClassLoader功能比较强大，它既可以从本地文件系统获取二进制文件来加载类，也可以从远程主机获取二进制文件来加载类。

在应用程序中可以直接使用URLCIassLoader加载类，URLClassLoader类提供了如下两个构造器，

\(1\)URLClassLoader\(URL\[\] urls\):使用默认的父类加载器创建一个ClassLoader对象，该对象将从urls所指定的系列路径来查询并加载类。

\(2\)URLCIassLoader\(URL\[\] urls,ClassLoader parent\):使用指定的父类加战器创建一个 ClassLoader对象，其他功能与前一个构造器相同。

一旦得到了 URLClassLoader对象之后，就可以调用该对象的loadClassO方法来加载指定类。

```
public class URLClassLoaderTest
{
    private static Connection conn;
    // 定义一个获取数据库连接方法
    public static Connection getConn(String url ,
        String user , String pass) throws Exception
    {
        if (conn == null)
        {
            // 创建一个URL数组
            URL[] urls = {new URL(
                "file:mysql-connector-java-5.1.30-bin.jar")};
            // 以默认的ClassLoader作为父ClassLoader，创建URLClassLoader
            URLClassLoader myClassLoader = new URLClassLoader(urls);
            // 加载MySQL的JDBC驱动，并创建默认实例
            Driver driver = (Driver)myClassLoader.
                loadClass("com.mysql.jdbc.Driver").newInstance();
            // 创建一个设置JDBC连接属性的Properties对象
            Properties props = new Properties();
            // 至少需要为该对象传入user和password两个属性
            props.setProperty("user" , user);
            props.setProperty("password" , pass);
            // 调用Driver对象的connect方法来取得数据库连接
            conn = driver.connect(url , props);
        }
        return conn;
    }
    public static void main(String[] args)throws Exception
    {
        System.out.println(getConn("jdbc:mysql://localhost:3306/mysql"
            , "root" , "32147"));
    }
}
```

三、通过反射查看类的信息

程序在运行时，动态获取对象和类的信息，称为反射。

3.1获取Class对象

每个类被加载之后，系统就会为该类生成—个对应的Class对象，通过该Class对象就可以访问到JVM中的这个类。在Java程序中获得Class对象通常有如下三种方式。

\(1\)使用Class类的forName\(String clazzName\)静态方法。该方法需要传入字符串参数，该字符串参数的值是某个类的全限定类名（必须添加完整包名）。

\(2\)调用某个类的class属性来获取该类对应的Class对象。例如，Person.class将会返回Person类对应的Class对象。

\(3\)调用某个对象的getClass\(\)方法。该方法是java.lang.Object类中的一个方法，所以所有的Java对象都可以调用该方法，该方法将会返回该对象所域类对应的Class对象。

对于第一种方式和第二种方式都是直接根据类来取得该类的 Class 对象，相比之下，第二种方式有如下两种优势：

\(1\)代码更安全。程序在编译阶段就可以检査需要访问的Class对象是否存在。

\(2\)程序性能更好。因为这种方式无须调用方法，所以性能更好。

3.2获取Class信息

\(1\)获取Class对应类所包含的构造器（参考API）

\(2\)获取Class对应类所包含的方法

\(3\)获取Class对应类所包含的成员变量

\(5\)获取Class对应类所包含的Annotation

\(6\)获取Class对应类所包含的内部类

\(7\)获取Class对应类所包含的外部类

\(8\)获取Class对应类所实现的接口

\(9\)获取Class对应类所继承的父类

\(10\)获取Class对应类的修饰符、所在包、类名等基本信息

\(11\)判断该类是否为接口、枚举、注解等类型

```
// 定义可重复注解
@Repeatable(Annos.class)
@interface Anno {}
@Retention(value=RetentionPolicy.RUNTIME)
@interface Annos {
    Anno[] value();
}
// 使用4个注解修饰该类
@SuppressWarnings(value="unchecked")
@Deprecated
// 使用重复注解修饰该类
@Anno
@Anno
public class ClassTest
{
    // 为该类定义一个私有的构造器
    private ClassTest()
    {
    }
    // 定义一个有参数的构造器
    public ClassTest(String name)
    {
        System.out.println("执行有参数的构造器");
    }
    // 定义一个无参数的info方法
    public void info()
    {
        System.out.println("执行无参数的info方法");
    }
    // 定义一个有参数的info方法
    public void info(String str)
    {
        System.out.println("执行有参数的info方法"
            + "，其str参数值：" + str);
    }
    // 定义一个测试用的内部类
    class Inner
    {
    }
    public static void main(String[] args)
        throws Exception
    {
        // 下面代码可以获取ClassTest对应的Class
        Class<ClassTest> clazz = ClassTest.class;
        // 获取该Class对象所对应类的全部构造器
        Constructor[] ctors = clazz.getDeclaredConstructors();
        System.out.println("ClassTest的全部构造器如下：");
        for (Constructor c : ctors)
        {
            System.out.println(c);
        }
        // 获取该Class对象所对应类的全部public构造器
        Constructor[] publicCtors = clazz.getConstructors();
        System.out.println("ClassTest的全部public构造器如下：");
        for (Constructor c : publicCtors)
        {
            System.out.println(c);
        }
        // 获取该Class对象所对应类的全部public方法
        Method[] mtds = clazz.getMethods();
        System.out.println("ClassTest的全部public方法如下：");
        for (Method md : mtds)
        {
            System.out.println(md);
        }
        // 获取该Class对象所对应类的指定方法
        System.out.println("ClassTest里带一个字符串参数的info()方法为："
            + clazz.getMethod("info" , String.class));
        // 获取该Class对象所对应类的上的全部注解
        Annotation[] anns = clazz.getAnnotations();
        System.out.println("ClassTest的全部Annotation如下：");
        for (Annotation an : anns)
        {
            System.out.println(an);
        }
        System.out.println("该Class元素上的@SuppressWarnings注解为："
            + Arrays.toString(clazz.getAnnotationsByType(SuppressWarnings.class)));
        System.out.println("该Class元素上的@Anno注解为："
            + Arrays.toString(clazz.getAnnotationsByType(Anno.class)));
        // 获取该Class对象所对应类的全部内部类
        Class<?>[] inners = clazz.getDeclaredClasses();
        System.out.println("ClassTest的全部内部类如下：");
        for (Class c : inners)
        {
            System.out.println(c);
        }
        // 使用Class.forName方法加载ClassTest的Inner内部类
        Class inClazz = Class.forName("ClassTest$Inner");
        // 通过getDeclaringClass()访问该类所在的外部类
        System.out.println("inClazz对应类的外部类为：" +
            inClazz.getDeclaringClass());
        System.out.println("ClassTest的包为：" + clazz.getPackage());
        System.out.println("ClassTest的父类为：" + clazz.getSuperclass());
    }
}
```

3.3Java8新增的方法参数反射

Java8在java.lang.reflect 下新增了一个Executable抽象基类，该对象代表可执行的类成员，该类派生了Constructor、Method 两个子类。

Executable 基类提供了大量方法来获取修饰该方法或构造器的注解信息；还提供了isVarArgs\(\)方法用于判断该方法或构造器是否包含数量可变的形参，以及通过 getModifiers \(\)方法来获取该方法或构造器的修饰符。除此之外， Executable 提供了如下两个方法来获取该方法或参数的形参个数及形参名。

```
(1)intgetParameterCount ():获取该构造器或方法的形参个数。
(2)Parameter [] getParameters ():获取该构造器或方法的所有形参。
```

上面第二个方法返回了一个 Parameter \[\]数组，Parameter也是Java8新增的 API,每个Parameter对象代表方法或构造器的一个参数。Parameter也提供了大量方法来获取声明该参数的泛型信息，还提供了如下常用方法来获取参数信息。

```
(1)getModifiers(): 获取修饰该形参的修饰符。
(2)String getName():获取形参名。
(3)TypegetParameterizedType():获取带泛型的形参类型。
(4)Class <?> getType ():获取形参类型。
(5)boolean isNamePresent ():该方法返回该类的 class 文件中是否包含了方法的形参名信息。
(6)boolean isVarArgsO :该方法用于判断该参数是否为个数可变的形参。
```

需要指出的足，使用 javac 命令编译 Java 源文件时，默认生成的 class 文件并不包含方法的形参名信息，因此调用 isNamePresent\(\)方法将会返回 false ,调用 getName \(\)方法也不能得到该参数的形参名。如果希望javac命令编译Java源文件时可以保留形参信息，则需要为该命令指定- parameters 选项。

```
class Test
{
    public void replace(String str, List<String> list){}
}
public class MethodParameterTest
{
    public static void main(String[] args)throws Exception
    {
        // 获取String的类
        Class<Test> clazz = Test.class;
        // 获取String类的带两个参数的replace()方法
        Method replace = clazz.getMethod("replace"
            , String.class, List.class);
        // 获取指定方法的参数个数
        System.out.println("replace方法参数个数：" + replace.getParameterCount());
        // 获取replace的所有参数信息
        Parameter[] parameters = replace.getParameters();
        int index = 1;
        // 遍历所有参数
        for (Parameter p : parameters)
        {
            if (p.isNamePresent())
            {
                System.out.println("---第" + index++ + "个参数信息---");
                System.out.println("参数名：" + p.getName());
                System.out.println("形参类型：" + p.getType());
                System.out.println("泛型类型：" + p.getParameterizedType());
            }
        }
    }
}
```

四、通过反射生成并操作对象

Class对象可以获得该类里的方法（由Method对象表示）、构造器（由Constructor对象表示\)、成员变量\(由Field对象表示），这三个类都位于java.lang.reflect包下，并实现了java.lang.reflect.Member接口。程序以通过Method对象来执行对应的方法，通过Constructor对象来调用对应的构造器创建实例，能通过Field对象直接访问并修改对象的成员变墩值。

4.1创建对象

通过反射来生成对象有如下两种方式。

\(1\)使用Class对象的newInstance\(\)方法来创建该Class对象对应类的实例，这种方式要求该Class对象的对应类有默认构造器，而执行 newInstance\(\)方法时实际上是利用默认构造器来创建该类的实例。

\(2\)先使用Class对象获取指定的Constructor对象，再调用Constructor对象的newlnstance\(\)方法来创建该Class对象对应类的实例。通过这种方式可以选择使用指定的构造器来创建实例。

通过第一种方式来创建对象是比较常见的情形，因为在很多 Java EE 框架中都需要根据配置文件信息来创建Java 对象，从配置文件读取的只是某个类的字符串类名，程序需要根据该字符串来创建对应的实例，就必须使用反射。

```
public class ObjectPoolFactory
{
    // 定义一个对象池,前面是对象名，后面是实际对象
    private Map<String ,Object> objectPool = new HashMap<>();
    // 定义一个创建对象的方法，
    // 该方法只要传入一个字符串类名，程序可以根据该类名生成Java对象
    private Object createObject(String clazzName)
        throws InstantiationException
        , IllegalAccessException , ClassNotFoundException
    {
        // 根据字符串来获取对应的Class对象
        Class<?> clazz = Class.forName(clazzName);
        // 使用clazz对应类的默认构造器创建实例
        return clazz.newInstance();
    }
    // 该方法根据指定文件来初始化对象池，
    // 它会根据配置文件来创建对象
    public void initPool(String fileName)
        throws InstantiationException
        , IllegalAccessException ,ClassNotFoundException
    {
        try(
            FileInputStream fis = new FileInputStream(fileName))
        {
            Properties props = new Properties();
            props.load(fis);
            for (String name : props.stringPropertyNames())
            {
                // 每取出一对key-value对，就根据value创建一个对象
                // 调用createObject()创建对象，并将对象添加到对象池中
                objectPool.put(name ,
                    createObject(props.getProperty(name)));
            }
        }
        catch (IOException ex)
        {
            System.out.println("读取" + fileName + "异常");
        }

    }
    public Object getObject(String name)
    {
        // 从objectPool中取出指定name对应的对象。
        return objectPool.get(name);
    }

    public static void main(String[] args)
        throws Exception
    {
        ObjectPoolFactory pf = new ObjectPoolFactory();
        pf.initPool("obj.txt");
        System.out.println(pf.getObject("a"));      // ①
        System.out.println(pf.getObject("b"));      // ②
    }
}
```

4.2调用方法

当获得某个类对应的 Class 对象后，就可以通过该 Class 对象的 getMethods \(\)方法或者 getMethod \(\)方法来获取全部方法或指定方法——这两个方法的返回值是 Method 数组，或者 Method 对象。  
每个 Method 对象对应一个方法，获得 Method 对象后，程序就可通过该 Method 来调用它对应的方法。在 Method 里包含一个 invoke \(\)方法，该方法的签名如下。：

```
Object invoke ( Objectobj ,  Object ... args ):
```

该方法中的obj是执行该方法的主调，后面的args是执行该方法时传入该方法的实参。

```
public class ExtendedObjectPoolFactory
{
	// 定义一个对象池,前面是对象名，后面是实际对象
	private Map<String ,Object> objectPool = new HashMap<>();
	private Properties config = new Properties();
	// 从指定属性文件中初始化Properties对象
	public void init(String fileName)
	{
		try(
			FileInputStream fis = new FileInputStream(fileName))
		{
			config.load(fis);
		}
		catch (IOException ex)
		{
			System.out.println("读取" + fileName + "异常");
		}
	}
	// 定义一个创建对象的方法，
	// 该方法只要传入一个字符串类名，程序可以根据该类名生成Java对象
	private Object createObject(String clazzName)
		throws InstantiationException
		, IllegalAccessException , ClassNotFoundException
	{
		// 根据字符串来获取对应的Class对象
		Class<?> clazz =Class.forName(clazzName);
		// 使用clazz对应类的默认构造器创建实例
		return clazz.newInstance();
	}
	// 该方法根据指定文件来初始化对象池，
	// 它会根据配置文件来创建对象
	public void initPool()throws InstantiationException
		,IllegalAccessException , ClassNotFoundException
	{
		for (String name : config.stringPropertyNames())
		{
			// 每取出一对key-value对，如果key中不包含百分号（%）
			// 这就标明是根据value来创建一个对象
			// 调用createObject创建对象，并将对象添加到对象池中
			if (!name.contains("%"))
			{
				objectPool.put(name ,
					createObject(config.getProperty(name)));
			}
		}
	}
	// 该方法将会根据属性文件来调用指定对象的setter方法
	public void initProperty()throws InvocationTargetException
		,IllegalAccessException,NoSuchMethodException
	{
		for (String name : config.stringPropertyNames())
		{
			// 每取出一对key-value对，如果key中包含百分号（%）
			// 即可认为该key用于控制调用对象的setter方法设置值，
			// %前半为对象名字，后半控制setter方法名
			if (name.contains("%"))
			{
				// 将配置文件中key按%分割
				String[] objAndProp = name.split("%");
				// 取出调用setter方法的参数值
				Object target = getObject(objAndProp[0]);
				// 获取setter方法名:set + "首字母大写" + 剩下部分
				String mtdName = "set" +
				objAndProp[1].substring(0 , 1).toUpperCase()
					+ objAndProp[1].substring(1);
				// 通过target的getClass()获取它实现类所对应的Class对象
				Class<?> targetClass = target.getClass();
				// 获取希望调用的setter方法
				Method mtd = targetClass.getMethod(mtdName , String.class);
				// 通过Method的invoke方法执行setter方法，
				// 将config.getProperty(name)的值作为调用setter的方法的参数
				mtd.invoke(target , config.getProperty(name));
			}
		}
	}
	public Object getObject(String name)
	{
		// 从objectPool中取出指定name对应的对象。
		return objectPool.get(name);
	}
	public static void main(String[] args)
		throws Exception
	{
		ExtendedObjectPoolFactory epf = new ExtendedObjectPoolFactory();
		epf.init("extObj.txt");
		epf.initPool();
		epf.initProperty();
		System.out.println(epf.getObject("a"));
	}
}


```

4.3访问成员变量值

通过Class对象的 getFields \(\)或 getField\(\)方法可以获取该类所包括的全部成员变量或指定成员变最。

Field 提供了如下两组方法来读取或设置成员变量值。

\(1\)getXxx \( Object obj \) : 获取 obj 对象的该成员变量的值。此处的 Xxx对应8种基本类型，如果该  
成员变量的类型足引用类型，则取消 get 后面的 Xxx 。

\(2\)setXxx \( Object obj ， Xxx val \):将 obj 对象的该成员变量设置成 val值。此处的 Xxx 对应8种基本  
类型，如果该成员变最的类型是引用类型，则取消 set 后面的 Xxx 。

使用这两个方法可以随意地访问指定对象的所有成员变景，包括 private 修饰的成员变景。

```
class Person
{
	private String name;
	private int age;
	public String toString()
	{
		return "Person[name:" + name +
		" , age:" + age + " ]";
	}
}
public class FieldTest
{
	public static void main(String[] args)
		throws Exception
	{
		// 创建一个Person对象
		Person p = new Person();
		// 获取Person类对应的Class对象
		Class<Person> personClazz = Person.class;
		// 获取Person的名为name的成员变量
		// 使用getDeclaredField()方法表明可获取各种访问控制符的成员变量
		Field nameField = personClazz.getDeclaredField("name");
		// 设置通过反射访问该成员变量时取消访问权限检查
		nameField.setAccessible(true);
		// 调用set()方法为p对象的name成员变量设置值
		nameField.set(p , "Yeeku.H.Lee");
		// 获取Person类名为age的成员变量
		Field ageField = personClazz.getDeclaredField("age");
		// 设置通过反射访问该成员变量时取消访问权限检查
		ageField.setAccessible(true);
		// 调用setInt()方法为p对象的age成员变量设置值
		ageField.setInt(p , 30);
		System.out.println(p);
	}
}

```

4.4操作数组

在java.lang.reflect 包下还提供了一个Array类，Array 对象可以代表所有的数组。程序可以通过使  
用 Array 来动态地创建数组，操作数组元素等。

Array 提供了如下几类方法。

```
(1)static Object newInstance(Class<?> componentType ,int…length)
 创建一个具有指定的元素类型、
指定维度的新数组。
(2)static xxx getXxx(Object array, int index ) :
 返回array数组中第index个元素。其中xxx是各种基
本数据类型，如果数组元素是引用类型，则该方法变为get(Object array，int index )。
(3)static void setXxx(Object array，int index，xxx val)：将array数组中第index个元素的值设为 val 。
其中 XXX 是各种基本数据类型，如果数组元素是引用类型，则该方法变成 set(Object array，intindex, Object val ).
```

```
public class ArrayTest1
{
	public static void main(String args[])
	{
		try
		{
			// 创建一个元素类型为String ，长度为10的数组
			Object arr = Array.newInstance(String.class, 10);
			// 依次为arr数组中index为5、6的元素赋值
			Array.set(arr, 5, "疯狂Java讲义");
			Array.set(arr, 6, "轻量级Java EE企业应用实战");
			// 依次取出arr数组中index为5、6的元素的值
			Object book1 = Array.get(arr , 5);
			Object book2 = Array.get(arr , 6);
			// 输出arr数组中index为5、6的元素
			System.out.println(book1);
			System.out.println(book2);
		}
		catch (Throwable e)
		{
			System.err.println(e);
		}
	}
}

```

五、使用反射生成JDK动态代理

在Java的java.lang.reflect包下提供一个Proxy类和一个InvocationHandler个接口，通过使用这个类和接口可以生成JDK动态代理或动态代理对象

5.1使用 Proxy 和 InvocationHandler 创建动态代理

Proxy提供了用于创建动态代理类和代理对象的静态方法，它也是所有动态代理类的父类。如果在程序中为一个或多个接口动态地生成实现类，就可以使用 Proxy 来创建动态代理类；如果需要为一个或多个接口动态地创建实例，也可以使用 Proxy 来创建动态代理实例。

Proxy 提供了如下两个方法来创建动态代理类和动态代理实例。

\(1\)static Class &lt;?&gt;  getProxyClass\(ClassLoader loader ,  Class &lt;?&gt;...  interfaces \):

创建一个动态代理类所对应的 Class 对象，该代理类将实现 interfaces 所指定的多个接口 。第一个 ClassLoader 参数指定生成动态代理类的类加载器。

\(2\)static Object newProxylnstance\(ClassLoader loader ， Class &lt;?&gt;\[\]  interfaces ,  InvocationHandler h \):

直接创建一个动态代理对象，该代理对象的实现类实现了interfaces 指定的系列接口，执行代理对象的每个方法时都会被替换执行 InvocationHandler 对象的 invoke 方法。

实际上，即使采用第一个方法生成动态代理类之后，如果程序需要通过该代理类来创建对象，依然需要传入一个InvocationHandler对象.也就是说,系统生成的每个代理对象都有一个与之关联的InvocationHandler 对象，

```
interface Person
{
	void walk();
	void sayHello(String name);
}
class MyInvokationHandler implements InvocationHandler
{
	/*
	执行动态代理对象的所有方法时，都会被替换成执行如下的invoke方法
	其中：
	proxy：代表动态代理对象
	method：代表正在执行的方法
	args：代表调用目标方法时传入的实参。
	*/
	public Object invoke(Object proxy, Method method, Object[] args)
	{
		System.out.println("----正在执行的方法:" + method);
		if (args != null)
		{
			System.out.println("下面是执行该方法时传入的实参为：");
			for (Object val : args)
			{
				System.out.println(val);
			}
		}
		else
		{
			System.out.println("调用该方法没有实参！");
		}
		return null;
	}
}
public class ProxyTest
{
	public static void main(String[] args)
		throws Exception
	{
		// 创建一个InvocationHandler对象
		InvocationHandler handler = new MyInvokationHandler();
		// 使用指定的InvocationHandler来生成一个动态代理对象
		Person p = (Person)Proxy.newProxyInstance(Person.class.getClassLoader()
			, new Class[]{Person.class}, handler);
		// 调用动态代理对象的walk()和sayHello()方法
		p.walk();
		p.sayHello("孙悟空");
	}
}

```

5.2动态代理和AOP

六、反射和泛型

从JDK5以后,Java的Class 类增加了泛型功能，从而允许使用泛型来限制Class类，例如String.class的类型实际上是 Class &lt; String &gt;。如果 Class 对应的类暂时未知，则使用 Class &lt;?&gt;。通过在反射中使用泛型，可以避免使用反射生成的对象需要强制类型转换。

6.1泛型和Class 类

使用 Class &lt; T &gt;泛型可以避免强制类型转换。

```
public class CrazyitObjectFactory
{
	public static Object getInstance(String clsName)
	{
		try
		{
			// 创建指定类对应的Class对象
			Class cls = Class.forName(clsName);
			// 返回使用该Class对象所创建的实例
			return cls.newInstance();
		}
		catch(Exception e)
		{
			e.printStackTrace();
			return null;
		}
	}
}

```

```
public class CrazyitObjectFactory2
{
	public static <T> T getInstance(Class<T> cls)
	{
		try
		{
			return cls.newInstance();
		}
		catch(Exception e)
		{
			e.printStackTrace();
			return null;
		}
	}
	public static void main(String[] args)
	{
		// 获取实例后无须类型转换
		Date d = CrazyitObjectFactory2.getInstance(Date.class);
		JFrame f = CrazyitObjectFactory2.getInstance(JFrame.class);
	}
}

```

6.2使用反射来获取泛型信息

通过指定类对应的 Class 对象，可以获得该类里包含的所有成员变量 ,不管该成员变量是使用 private修饰，还是使用 public 修饰。获得了成员变量对应的 Field 对象后，就可以很容易地获得该成员变量的数据类型，即使用如卜代码即 nj 获得指定成员变量的类型。

```
/ / 获取成员变量的类型
Class<?> a = f.getType();
```

但这种方式只对普通类型的成员变量有效。如果该成员变里的类型是有泛型类型的类型，如Map&lt;String,Integer&gt; 类型，则不能准确地得到该成员变量的泛型参数。

为了获得指定成员变量的泛型类型，应先使用如下方法来获取该成员变量的泛型类型。

```
/ / 获取成员变量的泛型类型
Type gType = f .getGenericType();
```

然后将Type对象强制类型转换为ParameterizedType 对象，ParameterizedType代表被参数化的类型，也就足增加了泛型限制的类型。 ParameterizedType 类提供了如下两个方法。

```
(1)getRawTypeO :返回没有泛型信息的原始类型。
(2)getActualTypeArguments ():返回泛型参数的类型。
```

```
public class GenericTest
{
	private Map<String , Integer> score;
	public static void main(String[] args)
		throws Exception
	{
		Class<GenericTest> clazz = GenericTest.class;
		Field f = clazz.getDeclaredField("score");
		// 直接使用getType()取出的类型只对普通类型的成员变量有效
		Class<?> a = f.getType();
		// 下面将看到仅输出java.util.Map
		System.out.println("score的类型是:" + a);
		// 获得成员变量f的泛型类型
		Type gType = f.getGenericType();
		// 如果gType类型是ParameterizedType对象
		if(gType instanceof ParameterizedType)
		{
			// 强制类型转换
			ParameterizedType pType = (ParameterizedType)gType;
			// 获取原始类型
			Type rType = pType.getRawType();
			System.out.println("原始类型是：" + rType);
			// 取得泛型类型的泛型参数
			Type[] tArgs = pType.getActualTypeArguments();
			System.out.println("泛型信息是:");
			for (int i = 0; i < tArgs.length; i++)
			{
				System.out.println("第" + i + "个泛型类型是：" + tArgs[i]);
			}
		}
		else
		{
			System.out.println("获取泛型类型出错！");
		}
	}
}

```



