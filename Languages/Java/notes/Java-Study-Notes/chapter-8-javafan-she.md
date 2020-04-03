Java Reflection

Reflection（反射）是被视为动态语言的关键，反射机制允许程序在执行期借助于Reflection API取得任何类的内部信息，并能直接操作任意对象的内部属性及方法。

Java反射机制提供的功能

1. 在运行时判断任意一个对象所属的类
2. 在运行时构造任意一个类的对象
3. 在运行时判断任意一个类所具有的成员变量和方法
4. 在运行时调用任意一个对象的成员变量和方法
5. 生成动态代理

反射相关的主要API：

1. java.lang.Class:代表一个类
2. java.lang.reflect.Method:代表类的方法
3. java.lang.reflect.Field:代表类的成员变量
4. java.lang.reflect.Constructor:代表类的构造方法

Class 类

在Object类中定义了以下的方法，此方法将被所有子类继承：

```
● public final Class getClass()
```

以上的方法返回值的类型是一个Class类，此类是Java反射的源头，实际上所谓反射从程序的运行结果来看也很好理解，即：可以通过对象反射求出类的名称。

![](/assets/Java反射.png)

一个 Class 对象包含了特定某个类的有关信息。

Class本身也是一个类

Class 对象只能由系统建立对象

一个类在 JVM 中只会有一个Class实例

一个Class对象对应的是一个加载到JVM中的一个.class文件

每个类的实例都会记得自己是由哪个 Class 实例所生成

通过Class可以完整地得到一个类中的完整结构

实例化Class类对象

1）前提：若已知具体的类，通过类的class属性获取，该方法

最为安全可靠，程序性能最高

实例：Class clazz = String.class;

2）前提：已知某个类的实例，调用该实例的getClass\(\)方法获

取Class对象

实例：Class clazz = “www.atguigu.com”.getClass\(\);

3）前提：已知一个类的全类名，且该类在类路径下，可通过

Class类的静态方法forName\(\)获取，可能抛出ClassNotFoundException

实例：Class clazz = Class.forName\(“java.lang.String”\);

4）其他方式\(不做要求\)

```
ClassLoader cl = this.getClass().getClassLoader();
Class clazz4 = cl.loadClass(“类的全类名”);
```

2.1.创建类的对象：调用Class对象的newInstance\(\)方法

要 求： 1）类必须有一个无参数的构造器。

2）类的构造器的访问权限需要足够。

难道没有无参的构造器就不能创建对象了吗？

不是！只要在操作的时候明确的调用类中的构造方法，并将参数传递进去之后，才可以实例化操作。步骤如下：

1）通过Class类的getDeclaredConstructor\(Class … parameterTypes\)取得本类的指定形参类型的构造器

2）向构造器的形参中传递一个对象数组进去，里面包含了构造器中所需的各个参数。

3）通过Constructor实例化对象。

使用反射可以取得：

1.实现的全部接口

```
public Class <?>[] getInterfaces()
```

确定此对象所表示的类或接口实现的接口。

2.所继承的父类

```
public Class <? Super T>  getSuperclass()
```

返回表示此 Class 所表示的实体（类、接口、基本类型）的父类的 Class。

3.全部的构造器

```
public Constructor <T> [] getConstructors()
```

返回此 Class 对象所表示的类的所有public构造方法。

```
public Constructor <T> [] getDeclaredConstructors()
```

返回此 Class 对象表示的类声明的所有构造方法。

Constructor类中：

取得修饰符: public int getModifiers\(\);

取得方法名称: public String getName\(\);

取得参数的类型：public Class &lt;?&gt; \[\] getParameterTypes\(\);

4.全部的方法

public Method\[\] getDeclaredMethods\(\)

返回此Class对象所表示的类或接口的全部方法

public Method\[\] getMethods\(\)

返回此Class对象所表示的类或接口的public的方法

Method类中：

```
public Class<?>getReturnType()取得全部的返回值
public Class<?>[] getParameterTypes()取得全部的参数
public int getModifiers()取得修饰符
public Class<?>[] getExceptionTypes()取得异常信息
```

5.全部的Field

public Field\[\] getFields\(\)

返回此Class对象所表示的类或接口的public的Field。

public Field\[\] getDeclaredFields\(\)

返回此Class对象所表示的类或接口的全部Field。

Field方法中：

```
public int getModifiers() 以整数形式返回此Field的修饰符
public Class<?>getType() 得到Field的属性类型
public String getName() 返回Field的名称。
```

1. Annotation相关

```
get Annotation(Class  <T> annotationClass)
getDeclaredAnnotations()
```

7.泛型相关

获取父类泛型类型：Type getGenericSuperclass\(\)

泛型类型：ParameterizedType

获取实际的泛型类型参数数组：getActualTypeArguments\(\)

8.类所在的包 Package getPackage\(\)

三、通过反射调用类中的指定方法、指定属性

1.调用指定方法

通过反射，调用类中的方法，通过Method类完成。步骤：

1.通过Class类的getMethod\(String name,Class…parameterTypes\)方法取得一个Method对象，并设置此方法操作时所需要的参数类型。

2.之后使用Object invoke\(Object obj, Object\[\] args\)进行调用，并向方法中传递要设置的obj对象的参数信息。

![](/assets/Javaffs.png)

说明：

1.Object 对应原方法的返回值，若原方法无返回值，此时返回null

2.若原方法若为静态方法，此时形参Object obj可为null

3.若原方法形参列表为空，则Object\[\] args为null

4.若原方法声明为private,则需要在调用此invoke\(\)方法前，显式调用方法对象的setAccessible\(true\)方法，将可访问private的方法。

2.调用指定属性

在反射机制中，可以直接通过Field类操作类中的属性，通过Field类提供的set\(\)和get\(\)方法就可以完成设置和取得属性内容的操作。

public Field getField\(String name\) 返回此Class对象表示的类或接口的指定的public的Field。

public Field getDeclaredField\(String name\)返回此Class对象表示的类或接口的指定的Field。

在Field中：

public Object get\(Object obj\) 取得指定对象obj上此Field的属性内容

public void set\(Object obj,Object value\) 设置指定对象obj上此Field的属性内容

注：在类中属性都设置为private的前提下，在使用set\(\)和get\(\)方法时，首先要使用Field类中的setAccessible\(true\)方法将需要操作的属性设置为可以被外部访问。

public void setAccessible\(true\)访问私有属性时，让这个属性可见。

Java动态代理

动态代理是指客户通过代理类来调用其它对象的方法，并且是在程序运行时根据需要动态创建目标类的代理对象。

动态代理使用场合:

调试

远程方法调用

代理设计模式的原理:

使用一个代理将对象包装起来, 然后用该代理对象取代原始对象. 任何对原始对象的调用都要通过代理. 代理对象决定是否以及何时将方法调用转到原始对象上

Proxy ：专门完成代理的操作类，是所有动态代理类的父类。通过此类为一个或多个接口动态地生成实现类。

提供用于创建动态代理类和动态代理对象的静态方法

static  Class  &lt;?&gt;getProxyClass\(ClassLoader  loader, Class&lt;?&gt;  ...

interfaces\) 创建一个动态代理类所对应的Class对象

static   Object newProxyInstance\(ClassLoader  loader, Class &lt;?&gt;\[\] interfaces, InvocationHandler

h\) 直接创建一个动态代理对象

动态代理步骤

![](/assets/clipboard.png)

![](/assets/clipboardq.png)

3.通过Proxy的静态方法

```
newProxyInstance(ClassLoader loader, Class[] interfaces, InvocationHandler h) 创建一个Subject接口代理
 RealSubject target = new RealSubject();
// Create a proxy to wrap the original implementation
 DebugProxy proxy = new DebugProxy(target);
 // Get a reference to the proxy through the Subject interface
 Subject sub = (Subject) Proxy.newProxyInstance(
 Subject.class.getClassLoader(),
 new Class[] { Subject.class }, proxy);
```

4.通过 Subject代理调用RealSubject实现类的方法

```
 String info = sub.say(“Peter", 24);
 System.out.println(info);
```



