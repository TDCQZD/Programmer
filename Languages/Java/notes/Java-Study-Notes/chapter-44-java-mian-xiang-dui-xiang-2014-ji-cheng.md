类与类之间同样具有关系，如一个百货公司类与销售员类相联系，类之间这种关系被称为关联。关联是描述两个类之间的一般二元关系，例如一个百货公司类与销售员类就是一个关联，再比如学生类以及教师类也是一个关联。两个类之间的关系有很多种，继承是关联中的一种。

1、继承定义  
定义：子类能够继承父类的属性和方法；  
注意点：Java中只支持单继承；私有方法不能继承；

2、继承的基本使用

Java的继承通过extends 关键字来实现。

父类：

```
/**
 * 动物类
 * @author user
 *
 */
public class Animal {

    private String name; // 姓名
    private int age;  // 年龄

    public String getName() {
        return name;
    }
    public void setName(String name) {
        this.name = name;
    }
    public int getAge() {
        return age;
    }
    public void setAge(int age) {
        this.age = age;
    }

    public void say(){
        System.out.println("我是一个动物，我叫："+this.name+",我的年龄是："+this.age);
    }
}
```

子类：

```
/**
 * 定义Dog类，继承自Animal类
 * @author user
 *
 */
public class Dog extends Animal{

    public static void main(String[] args) {
        Dog dog=new Dog();
        dog.setName("Pick");
        dog.setAge(1);
        dog.say();
    }
}
```

3、方法重写

所谓方法的重写 我们可以在子类中根据实际业务把父类的方法重写；

方法重写需要注意：

\(1\)方法名相间、形参列表相同

\(2\)子类方法返回值类型应比父类方法返回值类型更小或相等，子类方法声明抛出的异常类应比父类方法声  
明抛出的异常类更小或相等

\(3\)子类方法的访问 权限应比父类方法的访问权限更大或相等.

尤其需要指出的是， 覆盖方法和被覆盖盖方法要么都是类方法，要么都是实例方法，不能一个是类方法，另一个是实例方法.

```
/**
 * 定义一个Cat类，继承自Animal
 * @author user
 *
 */
public class Cat extends Animal{

    /**
     * 重写父类的say方法
     */
    public void say(){
        System.out.println("我是一个猫，我叫："+this.getName()+",我的年龄是："+this.getAge());
    }

    public static void main(String[] args) {
        Cat cat=new Cat();
        cat.setName("Mini");
        cat.setAge(2);
        cat.say();
    }
}
```

4、对象实例过程

对象实例化 先实例化调用父类构造方法，再调用子类实例化构造方法；

```
/**
 * 动物类
 * @author user
 *
 */
public class Animal {

    private String name; // 姓名
    private int age;  // 年龄


    /**
     * 无参父类构造方法
     */
    public Animal() {
        System.out.println("无参父类构造方法");
    }

    /**
     * 有参父类构造方法
     * @param name 姓名
     * @param age 年龄
     */
    public Animal(String name,int age) {
        System.out.println("有参父类构造方法");
        this.name=name;
        this.age=age;
    }

    public String getName() {
        return name;
    }
    public void setName(String name) {
        this.name = name;
    }
    public int getAge() {
        return age;
    }
    public void setAge(int age) {
        this.age = age;
    }

    public void say(){
        System.out.println("我是一个动物，我叫："+this.name+",我的年龄是："+this.age);
    }
}


/**
 * 定义一个Cat类，继承自Animal
 * @author user
 *
 */
public class Cat extends Animal{

    private String address;


    public String getAddress() {
        return address;
    }

    public void setAddress(String address) {
        this.address = address;
    }

    public Cat() {
        super();
        System.out.println("子类无参构造方法");
    }

    public Cat(String name, int age,String address) {
        super(name, age);
        this.address=address;
        System.out.println("子类有参构造方法");
    }

    /**
     * 重写父类的say方法
     */
    public void say(){
        super.say(); // 调用父类的say()方法
        System.out.println("我是一个猫，我叫："+this.getName()+",我的年龄是："+this.getAge()+"，我来自："+this.getAddress());
    }

    public static void main(String[] args) {
        Cat cat=new Cat("Mini",2,"火星");
        /*cat.setName("Mini");
        cat.setAge(2);*/
        cat.say();
    }
}
```

5、super关键字

super关键字用于限定该对象调用它从父类继承得到的变量实例和方法。



