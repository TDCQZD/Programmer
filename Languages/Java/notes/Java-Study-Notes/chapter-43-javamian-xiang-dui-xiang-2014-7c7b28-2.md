1、包装类

每个基本类型都有一个对应的类；就是所谓的包装类；

![](/assets/保证类.jpg)

2、Object类

Object类是比较特殊的类，它是所有类的父类，是Java类层中的最高层类，实质上Java中任何一个类都是它的子类。当创建一个类时，总是在继承，除非某个类已经指定要从其他类继承，否则它就是从java.lang.Object类继承而来的，可见Java中的每个类都源于java.lang.Object类，如String、Integer等类都是继承于Object类；除此之外自定义的类也都继承于Object类。由于所有类都是Object子类，所以在定义类时，省略了extends Object关键字。  
Object类的常用方法:

```
(1)public String toString() 返回该对象的字符串表示。
(2)public boolean equals(Object obj) 指示其他某个对象是否与此对象“相等”
```

3、instanceof关键字

当在程序中执行向下转型操作时，如果父类对象不是子类对象的实例，就会发生ClassCastException异常，所以在执行向下转型之前需要养成一个良好的习惯，就是判断父类对象是否为子类对象的实例。这个判断通常使用instanceof操作符来完成。

作用：判断运算符前面的对象是否属于运算符后面的类。

格式：对象 instanceof 类 返回布尔类型

注意：instanceof 运算符前面操作符的编译类型要么与后面的类相同，要么与后面的类具有父子继承关系，否则会引起编译错误。

通过instanceof我们可以确保向下转型的不出问题；

```
public class Dog extends Animal{

    public void say(){
        System.out.println("我是一只狗");
    }

    public void f1(){
        System.out.println("汪汪...");
    }
}


public class Cat extends Animal{

    public void say(){
        System.out.println("我是一只猫");
    }

    public void f2(){
        System.out.println("我喜欢吃鱼");
    }


public class Test {

    public static void doSomeThing(Animal animal){
        animal.say();
        if(animal instanceof Dog){
            ((Dog) animal).f1();
        }else if(animal instanceof Cat){
            ((Cat) animal).f2();
        }
    }

    public static void main(String[] args) {
        Animal dog=new Dog();
        System.out.println("dog对象是否属于Animal类："+(dog instanceof Animal));
        System.out.println("dog对象是否属于Dog类："+(dog instanceof Dog));
        System.out.println("dog对象是否属于Cat类："+(dog instanceof Cat));

        doSomeThing(new Dog());
        doSomeThing(new Cat());
    }
}
```

4、final关键字

final是终结 完结的意思；

\(1\)使用final声明的类不能被继承；

\(2\)使用final声明的方法不能被子类覆盖，即定义为final的方法不能被重写；

\(3\)将方法定义为final类型可以防止任何子类修改该类的定义与实现方式，同时定义为final的方法执行效率要高于非final方法。

\(4\)使用final声明的变量不能被修改，即为常量；

5、static关键字

由static修饰的变量、常量和方法被称作静态变量、常量和方法。被声明为static的变量、常量和方法被称作静态成员。

静态成员属于类本身，区别于个别对象。

5、内部类

局部内部类就是在类的方法中定义的内部类，它的作用范围也是在这个方法体内。在一个类中使用内部类可以在内部类中直接存取其所在类的私有成员变量。

内部类定义：在类的内部定义的类；

内部类优点：可以方便的使用外部类的属性；

内部类缺点：破坏类的基本结构；

注意：慎用内部类；

```
/**
 * 外部类
 * @author user
 *
 */
public class Outer {

    private int a=1;

    /**
     * 内部类
     * @author user
     *
     */
    class Inner{
        public void show(){
            System.out.println(a); // 可以方便的额使用外部类的属性；
        }
    }

    public void show(){
        Inner inner=new Inner();
        inner.show();
    }

    public static void main(String[] args) {
        Outer outer=new Outer();
        outer.show();
    }
}
```

6、匿名内部类

匿名内部类 就是类的内部再定义类；

匿名内部类 这里指的是实例化内部对象 就是没有名字的内部类；

作用：假如某个类只使用一次，则可以使用匿名内部类；

```
public class Test {

    public void test(A a){
        a.a();
    }

    public static void main(String[] args) {
        Test t=new Test();
        t.test(new B());

        // 匿名内部类
        t.test(new A(){

            @Override
            public void a() {
                System.out.println("匿名内部类，一次性使用");
            }

        });
    }
}
```

7、静态内部类

在内部类前添加修饰符static，这个内部类就变为静态内部类。一个静态内部类中可以声明static成员，但是在非静态内部类中不可以声明静态成员。

静态内部类具有以下特点：

\(1\) 创建静态内部类的对象，不需要其外部类的对象。

\(2\)不能从静态内部类的对象中访问非静态外部类的对象。

\(3\)不可以使用外部类的非静态成员，所以静态内部类在程序开发中比较少见。

8、静态变量、常量和方法

由static修饰的变量、常量和方法被称作静态变量、常量和方法。被声明为static的变量、常量和方法被称作静态成员。

静态成员属于类所有，而不属于该类的单个实例，区别于个别对象。

静态成员不能直接访问非静态成员。

静态成员的调用:

```
类名.静态成员
```

9、String对象

\(1\)实例化String对象有两种方式，

a.直接赋值方式，创建的对象存放到字符串对象池里，假如存在的，就不会再创建；

b.new对象方式，每次都创建一个新的对象；

代码：

```
public class Demo1 {

    public static void main(String[] args) {
        // 实例化String的方式一,直接赋值方式
        String name1="张三";
        System.out.println("name1="+name1);

        // 实例化String的方式二,
        String name2=new String("李四");
        System.out.println("name2="+name2);
    }
}
```

\(2\)字符串的特性：不能改变字符串的内容；要想改变，只能通过指向一个新的内存地址；

\(3\)String类常用方法及基本使用

* char chartAt\(int index\) 返回指定索引处的char值

这里的index 是从0开始的；

```
public class Demo5 {

    public static void main(String[] args) {
        String name="张三";
        char ming=name.charAt(1);
        System.out.println(ming);

        String str="我是中国人";
        for(int i=0;i<str.length();i++){
            System.out.println(str.charAt(i));
        }
    }
}
```

* int length\(\) 返回字符串的长度；
* int indexOf\(int ch\) 返回指定字符在此字符串中第一次出现处的索引。

```
public class Demo06 {

    public static void main(String[] args) {
        // indexOf方法使用实例
        String str="abcdefghijdklmoqprstuds";
        System.out.println("d在字符串str中第一次出现的索引位置："+str.indexOf('d'));
        System.out.println("d在字符串str中第一次出现的索引位置,从索引4位置开始："+str.indexOf('d',4));
    }
}
```

* String substring\(int beginIndex\)   返回一个新的字符串，它是此字符串的一个子字符串。

```
public class Demo07 {

    public static void main(String[] args) {
        // substring方式读取
        String str="不开心每一天";
        String str2="不开心每一天，不可能";
        String newStr=str.substring(1);
        System.out.println(newStr);
        String newStr2=str2.substring(1, 6);
        System.out.println(newStr2);
    }
}
```

* public String toUpperCase\(\)  String 中的所有字符都转换为大写

```
package com.java1234.chap03.sec08;

public class Demo08 {

    public static void main(String[] args) {
        String str="I'm a boy!";
        String upStr=str.toUpperCase(); // 转成大写
        System.out.println(upStr);
        String lowerStr=upStr.toLowerCase(); // 转成小写
        System.out.println(lowerStr);
    }
}
```



