接口：一种特殊的“抽象类”，没有普通方法，由全局常量和公共的抽象方法所组成；

1、接口的定义

接口定义用关键字 interface

```
* 定义一个接口A
 * @author user
 *
 */
public interface A {

    /**
     * 全局常量
     */
    public static final String TITLE="www.java1234.com";

    /**
     * 定义一个抽象方法 abstract可以省略  
     */
    public abstract void a();
}
```

注意点：由于接口里的方法都是抽象的，所以abstract可以省略，实际开发一般都是省略的，开发者的习惯；

2、实现接口 可以实现一个或者多个接口

实现多个接口，中间用逗号隔开，并且实现所有抽象方法；

```
/**
 * 定义一个接口B
 * @author user
 *
 */
public interface B {

    /**
     * 全局常量
     */
    public static final String TITLE2="java知识分享网";

    /**
     * 定义一个抽象方法 abstract可以省略  
     */
    public abstract void b();
}
```

```
    public class Test implements A,B{
    @Override
    public void a() {
        System.out.println("a方法");
    }

    @Override
    public void b() {
        System.out.println("b方法");
    }

    public static void main(String[] args) {
        Test t=new Test();
        t.a();
        t.b();
        System.out.println(Test.TITLE);
        System.out.println(Test.TITLE2);
    }


}
```

3、继承类和实现接口 先继承，后实现接口

```
public class C {

    public void c(){
        System.out.println("c方法");
    }
}
```

```
public class Test extends C implements A,B{

    @Override
    public void a() {
        System.out.println("a方法");
    }

    @Override
    public void b() {
        System.out.println("b方法");
    }

    public static void main(String[] args) {
        Test t=new Test();
        t.a();
        t.b();
        t.c();
        System.out.println(Test.TITLE);
        System.out.println(Test.TITLE2);
    }


}
```



