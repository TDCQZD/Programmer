# JAVA**注解**

**元注解：**

元注解的作用就是负责注解其他注解。Java5.0定义了4个标准的meta-annotation类型，它们被用来提供对其它 annotation类型作说明。Java5.0定义的元注解：  
　　　　1.@Target,  
　　　　2.@Retention,  
　　　　3.@Documented,  
　　　　4.@Inherited  
　　这些类型和它们所支持的类在java.lang.annotation包中可以找到。下面我们看一下每个元注解的作用和相应分参数的使用说明。**@Target：**

@Target说明了Annotation所修饰的对象范围：Annotation可被用于 packages、types（类、接口、枚举、Annotation类型）、类型成员（方法、构造方法、成员变量、枚举值）、方法参数和本地变量（如循环变量、catch参数）。在Annotation类型的声明中使用了target可更加明晰其修饰的目标。

**作用：用于描述注解的使用范围（即：被描述的注解可以用在什么地方）**

**　取值\(ElementType\)有：**

1.CONSTRUCTOR:用于描述构造器  
　　　　2.FIELD:用于描述域  
　　　　3.LOCAL\_VARIABLE:用于描述局部变量  
　　　　4.METHOD:用于描述方法  
　　　　5.PACKAGE:用于描述包  
　　　　6.PARAMETER:用于描述参数  
　　　　7.TYPE:用于描述类、接口\(包括注解类型\) 或enum声明

注解Table 可以用于注解类、接口\(包括注解类型\) 或enum声明,而注解NoDBColumn仅可用于注解类的成员变量。

**Retention：**

**　@Retention**定义了该Annotation被保留的时间长短：某些Annotation仅出现在源代码中，而被编译器丢弃；而另一些却被编译在class文件中；编译在class文件中的Annotation可能会被虚拟机忽略，而另一些在class被装载时将被读取（请注意并不影响class的执行，因为Annotation与class在使用上是被分离的）。使用这个meta-Annotation可以对 Annotation的“生命周期”限制。

**作用：表示需要在什么级别保存该注释信息，用于描述注解的生命周期（即：被描述的注解在什么范围内有效）**

**　取值（RetentionPoicy）有：**

1.SOURCE:在源文件中有效（即源文件保留）  
　　　　2.CLASS:在class文件中有效（即class保留）  
　　　　3.RUNTIME:在运行时有效（即运行时保留）

Retention meta-annotation类型有唯一的value作为成员，它的取值来自java.lang.annotation.RetentionPolicy的枚举类型值。

**@Documented:**

**@**Documented用于描述其它类型的annotation应该被作为被标注的程序成员的公共API，因此可以被例如javadoc此类的工具文档化。Documented是一个标记注解，没有成员。

**　@Inherited：**

@Inherited 元注解是一个标记注解，@Inherited阐述了某个被标注的类型是被继承的。如果一个使用了@Inherited修饰的annotation类型被用于一个class，则这个annotation将被用于该class的子类。

注意：@Inherited annotation类型是被标注过的class的子类所继承。类并不从它所实现的接口继承annotation，方法并不从它所重载的方法继承annotation。

当@Inherited annotation类型标注的annotation的Retention是RetentionPolicy.RUNTIME，则反射API增强了这种继承性。如果我们使用java.lang.reflect去查询一个@Inherited annotation类型的annotation时，反射代码检查将展开工作：检查class和其父类，直到发现指定的annotation类型被发现，或者到达类继承结构的顶层。

**自定义注解：**

使用@interface自定义注解时，自动继承了java.lang.annotation.Annotation接口，由编译程序自动完成其他细节。在定义注解时，不能继承其他的注解或接口。@interface用来声明一个注解，其中的每一个方法实际上是声明了一个配置参数。方法的名称就是参数的名称，返回值类型就是参数的类型（返回值类型只能是基本类型、Class、String、enum）。可以通过default来声明参数的默认值。

**定义注解格式：**  
　　public @interface 注解名 {定义体}

**注解参数的可支持数据类型：**

1.所有基本数据类型（int,float,boolean,byte,double,char,long,short\)  
　　　　2.String类型  
　　　　3.Class类型  
　　　　4.enum类型  
　　　　5.Annotation类型  
　　　　6.以上所有类型的数组

Annotation类型里面的参数该怎么设定:  
　　第一,只能用public或默认\(default\)这两个访问权修饰.例如,String value\(\);这里把方法设为defaul默认类型；　 　  
　　第二,参数成员只能用基本类型byte,short,char,int,long,float,double,boolean八种基本数据类型和 String,Enum,Class,annotations等数据类型,以及这一些类型的数组.例如,String value\(\);这里的参数成员就为String;　　  
　　第三,如果只有一个参数成员,最好把参数名称设为"value",后加小括号.例:下面的例子FruitName注解就只有一个参数成员。

### 1.1 概述 {#11-概述}

注解\(Annotation\)在JDK1.5之后增加的一个新特性，注解的引入意义很大，有很多非常有名的框架，比如Hibernate、Spring等框架中都大量使用注解。注解作为程序的元数据嵌入到程序。注解可以被解析工具或编译工具解析，此处注意注解不同于注释\(comment\)。

当一个接口直接继承java.lang.annotation.Annotation接口时，仍是接口，而并非注解。要想自定义注解类型，只能通过@interface关键字的方式，其实通过该方式会隐含地继承.Annotation接口。1.2 API 摘要

所有与Annotation相关的API摘要如下：

\(1\). 注解类型\(Annotation Types\) API

| 注解类型 | 含义 |
| :--- | :--- |
| Documented | 表示含有该注解类型的元素\(带有注释的\)会通过javadoc或类似工具进行文档化 |
| Inherited | 表示注解类型能被自动继承 |
| Retention | 表示注解类型的存活时长 |
| Target | 表示注解类型所适用的程序元素的种类 |

\(2\). 枚举\(Enum\) API

| 枚举 | 含义 |
| :--- | :--- |
| ElementType | 程序元素类型，用于Target注解类型 |
| RetentionPolicy | 注解保留策略，用于Retention注解类型 |

\(3\). 异常和错误 API

| 异常/错误 | 含义 |
| :--- | :--- |
| AnnotationTypeMismatchException | 当注解经过编译\(或序列化\)后，注解类型改变的情况下，程序视图访问该注解所对应的元素，则抛出此异常 |
| IncompleteAnnotationException | 当注解经过编译\(或序列化\)后，将其添加到注解类型定义的情况下，程序视图访问该注解所对应的元素，则抛出此异常。 |
| AnnotationFormatError | 当注解解析器试图从类文件中读取注解并确定注解出现异常时，抛出该错误 |



