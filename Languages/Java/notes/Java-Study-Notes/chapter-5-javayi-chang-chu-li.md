**异常是什么？**

在程序中，错误可能产生于程序员没有预料到的各种情况，或者是因为超出了程序员控制之外的环境因素，如用户的坏数据、试图打开一个根本不存在的文件等。在Java中这种在程序运行时可能出现的一些错误称为异常。异常是一个在程序执行期间发生的事件，它中断了正在执行的程序的正常指令流。

**异常的分类？**

Java类库中的每个包中都定义了异常类，所有这些类都是Throwable类的子类。Throwable类派生了两个子类，分别是Exception和Error类。

Error类及其子类用来描述Java运行系统中的内部错误以及资源耗尽的错误，这类错误比较严重。

Exception类称为非致命性类，可以通过捕捉处理使程序继续执行。Exception类又根据错误发生的原因分为RuntimeException异常和RuntimeException之外的异常。

![](/assets/异常.png)

异常产生后，如果不做任何处理，程序就会被终止。

**Exception 和 RuntimeException 区别？**

RuntimeException异常是程序运行过程中产生的异常。

Exception是检查型异常，在程序中必须使用try...catch进行处理；

RuntimeException是非检查型异常，例如NumberFormatException，可以不使用try...catch进行处理，但是如果产生异常，则异常将由JVM进行处理；

RuntimeException最好也用try...catch捕获；

**如何捕获异常？**

Java语言的异常捕获结构由try、catch和finally 3部分组成。其中，try语句块存放的是可能发生异常的Java语句；catch程序块在try语句块之后，用来激发被捕获的异常；finally语句块是异常处理结构的最后执行部分，无论try块中的代码如何退出，都将执行finally块。

try、catch捕获异常：

```
public class ExceptionDemo {

    public static void main(String[] args) {
        String str="123a";
        try{
            int a=Integer.parseInt(str);          
        }catch(NumberFormatException e){
            e.printStackTrace();
        }
        System.out.println("继续执行");
    }
}
```

try...cacth...finally：

```
public class Demo2 {

    public static void testFinally(){
        String str="123a";
        try{
            int a=Integer.parseInt(str);
            System.out.println(a);
        }catch(Exception e){
            e.printStackTrace();
            System.out.println("exception");
            return;
        }finally{
            System.out.println("finally end");
        }
        System.out.println("end");
    }

    public static void main(String[] args) {
        testFinally();
    }
}
```

**Java常见异常**

![](/assets/常见异常.png)

**如何自定义异常类？**

在程序中使用自定义异常类，大体可分为以下几个步骤：

（1）创建自定义异常类。

（2）在方法中通过throw关键字抛出异常对象。

（3）如果在当前抛出异常的方法中处理异常，可以使用try-catch语句捕获并处理，否则在方法的声明处通过throws关键字指明要抛出给方法调用者的异常，继续进行下一步操作。

（4）在出现异常方法的调用者中捕获并处理异常。

```
/**
 * 自定义异常，继承自Exception
 * @author user
 *
 */
public class CustomException extends Exception{

    public CustomException(String message) {
        super(message);
    }

}
public class TestCustomException {

    public static void test()throws CustomException{
        throw new CustomException("自定义异常");
    }

    public static void main(String[] args) {
        try {
            test();
        } catch (CustomException e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
    }
}
```

**如何在方法中抛出异常？**

1、使用throws声明抛出异常

throws关键字通常被应用在声明方法时，用来指定方法可能抛出的异常。多个异常可使用逗号分隔。

2、使用throw语句抛出异常

throw关键字通常用于方法体中，并且抛出一个异常对象。程序在执行到throw语句时立即终止，它后面的语句都不执行。通过throw抛出异常后，如果想在上一级代码中来捕获并处理异常，则需要在抛出异常的方法中使用throws关键字在方法的声明中指明要抛出的异常；如果要捕捉throw抛出的异常，则必须使用try-catch语句。

throws表示当前方法不处理异常，而是交给方法的调用出去处理；

```
public class Demo1 {

    /**
     * 把异常向外面抛
     * @throws NumberFormatException
     */
    public static void testThrows()throws NumberFormatException{
        String str="123a";
        int a=Integer.parseInt(str);
        System.out.println(a);
    }

    public static void main(String[] args) {
        try{
            testThrows();  
            System.out.println("here");
        }catch(Exception e){
            System.out.println("我们在这里处理异常");
            e.printStackTrace();
        }
        System.out.println("I'm here");
    }
}
```

throw表示直接抛出一个异常；

```
public class Demo2 {

    public static void testThrow(int a) throws Exception{
        if(a==1){
            // 直接抛出一个异常类
            throw new Exception("有异常");
        }
        System.out.println(a);
    }

    public static void main(String[] args) {
        try {
            testThrow(1);
        } catch (Exception e) {
            // TODO Auto-generated catch block
            e.printStackTrace();
        }
    }
}
```

**异常的处理机制**

1、使用try-catch捕获异常

2、Java7以后可以多异常捕获

3、异常类存在继承体系，父类异常处理的catch块要放在子类异常处理的catch块后，否则会出现编译异常。简称先处理小异常再处理大异常。

4、获取异常信息

如果程序需要在catch块中访问异常对象的相关信总，则可以通过访问catch块的后异常形参来获  
得。当Java运行时决定调用某个catch块来处理该异常对象时，会将异常对象赋给catch块后的异常参  
数，程序即可通过该参数来获得异常的相关信息。

所有的异常对象都包含了如下几个常用方法。

\(1\)getMessageO:返回该异常的详细描述字符串。

\(2\) printStackTraceO:将该异常的跟踪栈信总输出到标准错误输出。

\(3\)printStackTrace\(PrintStreams\):将该异常的跟踪栈信息输出到指定输出流。

\(4\)getStackTraceO:返回该异常的跟踪栈信息。

5、使用finally回收物理资源（比如数据库连接、网络连接和磁盘信息等）

6、异常处理的嵌套

finally 块中也包含了一个完整的异常处理流程，这种在 try 块 、catch  
块或 finally 块中包含完整的异常处理流程的情形被称为异常处理的嵌套。

异常处理流程代码可以放在任何能放可执行性代码的地方，因此完整的异常处理流程既可放在 try  
块里，也可放在 catch 块里，还可放在 finally 块里。

异常处理嵌奁的深度没有很明确的限制，但通常没有必要使用超过两层的嵌套异常处理，层次太深

的嵌套异常处理没有太大必要， 而且导致程序可读性降低。

7、Java 7自动关闭资源的try语句

Java 7增强了  try 语句的功能——它允许在 try 关键字后紧跟一对圆括号，圆括号可以声明、初始化一个  
或多个资源,此处的资源指的足那些必须在程序结束时 M 式关闭的资源（比如数据库连接、 W 络连接等），  
try 语句在该语句结束时自动关闭这些资源。

需要指出的是，为了保证 try 语句可以正常关闭资源，这些资源实现类必须实现 AutoCloseable 或  
Closeable 接口，实现这两个接口就必须实现 close \(\)方法。

```
public static void main(String[] args)
		throws IOException
	{
		try (
			// 声明、初始化两个可关闭的资源
			// try语句会自动关闭这两个资源。
			BufferedReader br = new BufferedReader(
				new FileReader("AutoCloseTest.java"));
			PrintStream ps = new PrintStream(new
				FileOutputStream("a.txt")))
		{
			// 使用两个资源
			System.out.println(br.readLine());
			ps.println("庄生晓梦迷蝴蝶");
		}
	}
```

**异常的处理原则**

在程序中使用异常，可以捕获程序中的错误，但是异常的使用也要遵循一定的规则，下面是异常类的几项使用原则：

1、不要过多的使用异常，这样会增加系统的负担。

2、在方法中使用try-catch捕获异常时，要对异常作出处理。

3、try-catch语句块的范围不要太大，这样不利于对异常的分析。

4、一个方法被覆盖时，覆盖它的方法必须抛出相同的异常或子异常。

5、避免使用catch All语句

/؏Z

/؏Z

