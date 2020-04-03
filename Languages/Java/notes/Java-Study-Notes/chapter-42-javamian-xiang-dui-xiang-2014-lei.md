类是同一类事物的统称。类是封装对象的属性和行为的载体，反过来说具有相同**属性和行为**的一类实体被称为类。在Java语言中类中对象的行为通过方法定义，对象的属性通过成员变量来定义。

一、变量

1、成员变量

在Java中对象的属性称为成员变量，也可以称为属性。

2、成员方法

在Java语言中使用成员方法对应于类对象的行为。

定义成员方法的语法格式如下所示：

```
权限修饰符 返回值类型 方法名(参数类型 参数名){
    ...//方法体
    return 返回值;
}
```

3、权限修饰符

Java中的权限修饰符主要包括 private\(私有\)， package\(包访问权限\)，protected\(子类访问权限\)，public\(公共访问权限\)，这些修饰符控制着对类和类的成员变量以及成员方法的访问。

![](/assets/访问控制权限.jpg)

4、局部变量

局部变量是在方法被执行时创建，在方法执行结束时被销毁。局部变量在使用时必须进行赋值操作或被初始化，否则会出现编译错误。

5、局部变量的有效范围

局部变量是在方法被执行时创建，在方法执行结束时被销毁。局部变量在使用时必须进行赋值操作或被初始化，否则会出现编译错误。

![](/assets/局部变量.png)

二、方法

1、类的构造方法

概念：构造方法是一个特殊的方法，这个特殊方法用于创建实例时执行初始化操作；

在类中除了成员方法之外，还存在一种特殊类型的方法，那就是构造方法。构造方法是一个与类同名的方法，对象的创建就是通过构造方法完成的。每当类实例化一个对象时，类都会自动调用构造方法。

构造方法的特点如下：

\(1\)构造方法没有返回值。

\(2\)构造方法的名称要与本类的名称相同。

```
/**
 * 定义人类
 * @author user
 *
 */
public class People {

    // 定义属性 
    private String name; // 实例化对象时，默认值是null
    private int age; // 实例化对象时，默认值是0

    /**
     * 默认构造方法
     */
    People(){
        System.out.println("默认构造方法！");
    }



    public void say(){
        System.out.println("我叫："+name+"，我今年："+age);
    }

    public static void main(String[] args) {
        People people=new People();
        people.say();
    }
}
```

2、类的主方法

主方法是类的入口点，它定义了程序从何处开始；主方法提供对程序流向的控制，Java编译器通过主方法来执行程序。主方法的语法如下：

```
public static void main(String[]args){
 //方法体
}
```

在主方法的定义中可以看到主方法具有以下特性：

\(1\)主方法也是静态的，所以如要直接在主方法中调用其他方法，则该方法必须也是静态的。

\(2\)主方法没有返回值。

\(3\)主方法的形参为数组。其中args\[0\]~args\[n\]分别代表程序的第一个参数到第n个参数，可以使用args.length获取参数的个数。

3、方法的值传递：

值传递 在方法里改变变量的值 作用范围仅仅是方法里 对外面不影响：

```
public class Person {

    void speak(int age){
        System.out.println("我今年"+age+"岁了");
        age=24; // 作用范围是方法里
    }

    public static void main(String[] args) {
        Person xiaoli=new Person();
        int age=23;
        xiaoli.speak(age);
        System.out.println(age);
    }
}
```

4、 引用传递

引用传递，传递的是地址，对象里的属性在方法里值修改，对外面有影响，我们通过对象.属性可以获取到最新的数据；

```
class SanWei{
    int b; // 胸围
    int w; // 腰围
    int h; // 腿围
}

public class Person {

    void speak(int age,SanWei sanWei){
        System.out.println("我今年"+age+"岁了，我的三围是："+sanWei.b+","+sanWei.w+","+sanWei.h);
        age=24; // 作用范围是方法里
        sanWei.b=80; // 胸围改成80
    }

    public static void main(String[] args) {
        Person xiaoli=new Person();
        int age=23;
        SanWei sanWei=new SanWei();
        sanWei.b=90;
        sanWei.w=60;
        sanWei.h=90;
        // age传递的是值 ，sanWei是对象，传递的是引用(地址,c里叫指针)
        xiaoli.speak(age,sanWei);
        System.out.println(age);  
        System.out.println(sanWei.b);
    }
}
```

5、方法的重载

方法的重载就是在同一个类中允许同时存在一个以上的同名方法，只要这些方法的参数个数或类型不同即可。

所谓方法的重载，类里面有两个或者多个重名的方法，但是方法的参数个数、类型、顺序至少有一个不一样，这时候局构成方法重载；

```
public class Demo01 {

    int add(int a,int b){
        System.out.println("方法一");
        return a+b;
    }

    /**
     * 方法的重载，参数个数不一样
     * @param a
     * @param b
     * @param c
     * @return
     */
    int add(int a,int b,int c){
        System.out.println("方法二");
        return a+b+c;
    }

    /**
     * 方法的重载，参数的类型不一样
     * @param a
     * @param b
     * @return
     */
    int add(int a,String b){
        System.out.println("方法三");
        return a+Integer.parseInt(b);
    }

    /**
     * 参数类型个数一样，返回值类型不一样 不算重载，直接会报错，说方法重名
     * @param args
     */
    /*long add(int a,int b){
        return a+b;
    }*/

    public static void main(String[] args) {
        Demo01 demo=new Demo01();
        System.out.println(demo.add(1, 2));
        System.out.println(demo.add(1, 2,3));
        System.out.println(demo.add(1, "3"));
    }
}
```

注意： 假如参数个数和类型一样，返回值不一样，不能算重载，直接是编译出错，编译器认为是方法重复了。

6、 方法的递归

递归，就是程序调用自身，我们讲的是方法递归调用，也就是在方法里自己调用自己；

```
public class Demo03 {

    /**
     * 递归方式
     * @param n
     * @return
     */
    static long diGui(int n){
        if(n==1){
            return 1;
        }
        return diGui(n-1)*n;
    }


    public static void main(String[] args) {
        System.out.println(Demo03.diGui(5));
    }
}
```

注意：递归也是一种算法。

7、this关键字

this总是指向调用该方法的对象；this关键字可以调用成员变量和成员方法。事实上this引用就是对一个对象的引用。

this表示当前对象:

\(1\)使用this调用本类中的属性；

\(2\)使用this调用构造方法；

```
/**
 * 定义人类
 * @author user
 *
 */
public class People {

    // 定义属性 
    private String name; // 实例化对象时，默认值是null
    private int age; // 实例化对象时，默认值是0

    /**
     * 默认构造方法
     */
    People(){
        System.out.println("默认构造方法！");
    }

    /**
     * 有参数的构造方法
     * @param name
     * @param age
     */
    People(String name,int age){
        this(); // 调用无参数的构造方法
        System.out.println("调用的是有参数的构造方法");
        this.name=name;
        this.age=age;
    }

    public void say(){
        System.out.println("我叫："+name+"，我今年："+age);
    }

    public static void main(String[] args) {
        // People people=new People();
        People people=new People("张三",20);
        people.say();
    }
}
```



