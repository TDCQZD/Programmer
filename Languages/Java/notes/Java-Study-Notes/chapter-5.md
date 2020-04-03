# **一、Java注释**

在调试代码，或者是有些代码暂时不用，我们可以给代码加注释，这时候代码就不会被执行；

## 1、单行注释 //

主要用于代码注释，以及语句备注;

```
public class Note {

    public static void main(String[] args) {
        System.out.println("java大爷，你好！"); // 这个是hellWorld备注
        // System.out.println("java大爷，你好！");
    }
}
```

## 2、多行注释： /\* ... \*/

```
public class Note {

    public static void main(String[] args) {
        System.out.println("java大爷，你好！"); // 这个是hellWorld备注
        // System.out.println("java大爷，你好！");


        /*System.out.println("第一行");
        System.out.println("第二行");
        System.out.println("第三行");*/
    }
}
```

## 3、文档注释： /\*\* ..\*/

我们对类和方法上加注释 ，要用文档注释；

```
/**
 * Java注释
 * @author user
 *
 */
public class Note {

    /**
     * 主方法
     * @param args
     */
    public static void main(String[] args) {
        System.out.println("java大爷，你好！"); // 这个是hellWorld备注
        // System.out.println("java大爷，你好！");


        /*System.out.println("第一行");
        System.out.println("第二行");
        System.out.println("第三行");*/
    }
}
```

# 二、**Java标识符命名**

标识符定义：Java语言中，对于变量，常量，函数，语句块也有名字，我们统统称之为Java标识符。  
标识符作用：标识符是用来给类、对象、方法、变量、接口和自定义数据类型命名的；

标识符命名规则：Java标识符由数字，字母和下划线\(\_\)，美元符号（$）组成。

在Java中是区分大小写的，而且还要求首位不能是数字。最重要的是，Java关键字不能当作Java标识符

# 三、**Java中的关键字**

![](/assets/Java中的关键字.jpg)

# 四、Java四类八种基本数据类型

![](/assets/Java四类八种基本数据类型.png)

Java数据类型可以分为 基本数据类型和引用数据类型。基本数据类型又细分为 数值型，字符型，布尔类型；

数值型又细分为整数类型和浮点型；

## 1、整数类型

在java中，根据位数，数据范围，可以分为byte,short,int,long类型；

![](/assets/整数类型jpg.jpg)

![](/assets/byte数据范围表示.jpg)

## 2、浮点类型

在Java中 根据位数和数据范围，浮点类型有float,double类型；

![](/assets/浮点类型.jpg)

## 3、字符类型

字符类型主要用来存储单个字符；字符型常量有3种表示形式；

\(1\)直接通过单个字符来指定字符型常量，如'A'，‘B’，'5'；

\(2\)通过转义字符表示特殊字符型常量，如'\n'，'\'；

有些字符属于特殊字符，不能直接使用，所以需要转义；

![](/assets/转义字符.jpg)

\(3\)直接使用Unicode值来表示字符型常量，如'\u66f9'，'\u950b'；

关于Unicode，全世界有这么多语言文字，我们可以用Unicode表示出来，我们主要要学会用Unicode来表示中文；

这里我们要学会看unicode汉字编码表，这里给下转载的帖子[http://blog.java1234.com/blog/articles/125.html](http://blog.java1234.com/blog/articles/125.html)

## 4、布尔类型

布尔类型只有true和false两个值； 用关键字boolean定义。

## 5、 基本类型的类型转换

基本类型的类型转换分为自动类型转换和强制类型转换

**自动类型转换**

条件：\(1\)转换前后的数据类型兼容；\(2\)转换后的数据类型的表示范围要比转换前的大；

**强制类型转换**

数据类型不兼容，一种不安全的转换；

# 五、**Java运算符与表达式**

## 1、赋值运算符      符号 "=" （赋值） 我们开发的时候，可以先定义变量，然后再赋值，分两个步骤；也可以定义变量的同时再赋值，一个操作；

```
public class Demo06 {

    public static void main(String[] args) {
        // 定义变量a
        int a;
        // 给变量a赋值
        a=1;
        System.out.println("a="+a);

        // 定义变量a，并且给a赋值
        int a2=2;
        System.out.println("a2="+a2);
    }
}
```

## 2、算数运算符

在java中 算数运算符有：

+\(加\)；

—\(减\)；

\* \(乘\)；

/ \(除\)；

%（取模）；取模其实就是求余数

## 3、自增与自减运算符

符号：++（自增）； --\(自减\)

重点掌握 a++ 和++a的区别

a++是先操作后加1；

++a是先加1后操作；

## 4、逻辑运算符

符号：

&& 与 ；&& 与  前后两个操作数必须都是true才返回true,否则返回false

& 不短路与 ； & 不短路与 表达式都会执行到

\|\| 或； \|\| 或 只要两个操作数中有一个是true，就返回true，否则返回false

\|不短路或 ；\| 不短路 或 表达式都会执行到

!非；! 非 ，如果操作数为true，返回false，否则，返回true

^异或；^ 异或 ，当两个操作数不相同时返回true，返回false

```
public class Demo09 {

    public static void main(String[] args) {
        // && 与  前后两个操作数必须都是true才返回true,否则返回false
                boolean b1=(5<3)&&(4>5);
                System.out.println("b1="+b1);

                // & 不短路与
                boolean b2=(5<3)&(4>5);
                System.out.println("b2="+b2);

                // 一般都用&& 短路  
                // 原因：效率高

                // || 或 只要两个操作数中有一个是true，就返回true，否则返回false
                boolean b3=(2<3)||(4>5);
                System.out.println("b3="+b3);

                // | 不短路 或
                boolean b4=(2<3)|(4>5);
                System.out.println("b4="+b4);

                // ! 非 ，如果操作数为true，返回false，否则，返回true
                boolean b5=!(3<4);
                System.out.println("b5="+b5);

                // ^ 异或 ，当两个操作数不相同时返回true，返回false
                boolean b6=(5>4)^(4>5);
                System.out.println("b6="+b6);
    }
}
```

## 5、关系运算符

符号：

&gt; 大于；

&lt; 小于；

&gt;=大于等于；

&lt;=小于等于；

==等于；

!=不等于；

## 6、三目运算符

格式：  \(表达式\)?表达式为true返回值:表达式为false返回值

```
public class Demo11 {

    public static void main(String[] args) {
        // 三目运算符
        String s=2>3?"表达式为真":"表达式为假";
        System.out.println(s);

        boolean b=2>3?true:false;
        System.out.println(b);
    }
}
```

# 六、**Java控制语句**

## 1、选择语句

### 在Java中 选择语句可以用if else 和switch来实现； \(1\)if else

if else可以直接用 条件满足执行里面的代码；或者多条件 可以用if else... else if ... else；

```
public class Demo12 {

    public static void main(String[] args) {
        int a=-1;
        // if语句
        // 多行注释快捷方式  ctrl+shift+/
        if(a>0){
            System.out.println(a+"是正数");
        }

        // if...else语句
        if(a>0){
            System.out.println(a+"是正数");
        }else{
            System.out.println(a+"不是正数");
        }

        // if...else if...else
        if(a>0){
            System.out.println(a+"是正数");
        }else if(a<0){
            System.out.println(a+"是负数");
        }else{
            System.out.println(a+"是0");
        }
    }
}
```

### \(2\)switch语句 用于条件很多的情况；

在switch在jdk1.6或者jdk1.6以下版本，仅支持整型类型；jdk1.7开始支持字符串；

整型类型示例：

```
public class Demo13 {

    public static void main(String[] args) {
        System.out.println("请输入一个数字：");
        // 定义一个系统输入对象
        // 自动导包  ctrl+shift+o
        Scanner scanner=new Scanner(System.in);
        int n=scanner.nextInt();
        scanner.close();
        switch(n){
            case 1:{
                System.out.println("用户输入的是1");
                break;
            }
            case 2:{
                System.out.println("用户输入的是2");
                break;
            }
            default:{
                System.out.println("用户输入的是其他数字");
            }
        }
    }
}
```

字符串示例：

```
public class Demo14 {

    public static void main(String[] args) {
        System.out.println("请输入一个字符串：");
        Scanner scanner=new Scanner(System.in);
        String str=scanner.next();
        scanner.close();
        switch(str){
            case "张三":{
                System.out.println("输入的是张三");
                break;
            }
            case "李四":{
                System.out.println("输入的是李四");
                break;
            }
            default:{
                System.out.println("用户输入的是其他字符串");
            }
        }
    }
}
```

## 2、while 和 do while循环语句

**while循环语句：**先定义i变量，while\(\) 括号里加条件判断；循环体里 我们输出i 以及i++；后面是不断的循环执行，直到while条件不满足为止，退出循环。

```
int i=1;
while(i<=10){
    System.out.print(i+" ");
    i++;
}
```

**do...while循环语句：**定义一个j变量，然后我们执行是执行输出语句，再j++；最后才是判断；直到while条件不满足为止，退出循环。

```
int j=1;
do{
    System.out.print(j+" ");
    j++;
}while(j<=10);
```

**while和do...while的区别**

从前面两个实例我们很容易看出，

while是先判断后执行 do...while是先执行后判断。

do..while是肯定会至少执行一次，while的话，不一定会执行。

```
public class Demo15 {

    public static void main(String[] args) {
        // 在控制台输出1到10
        // while循环语句
        int i=1;
        while(i<=10){
            System.out.print(i+" ");
            i++;
        }

        System.out.println("===================");

        // do....while 循环语句
        int j=1;
        do{
            System.out.print(j+" ");
            j++;
        }while(j<=10);

        System.out.println("===================");

        // while和do...while的区别
        // while是先判断后执行 do...while是先执行后判断


    }
}
```

## 3、for循环语句

for循环有点类似while循环语句；

int k=1;定义一个变量；

k&lt;=10是条件语句 ，满足条件就执行循环体，不满足就结束循环；

执行完循环体，再执行k++;

```
for(int k=1;k<=10;k++){
    System.out.print(k+" ");
}
```

**for循环嵌套**

假如在for循环中还有for循环，我们称之为循环的嵌套；

```
for(int m=0;m<10;m++){
    for(int n=0;n<10;n++){
        System.out.print("m="+m+"n="+n+" ");
    }
    System.out.println();
}
```

## 4、循环结构的控制

循环结构的控制

**1、break语句**

break语句使用场合主要是switch语句和循环结构。break语句，那么直接退出循环，执行循环结构下面的第一条语句。如果在多重嵌套循环中使用break语句，当执行break语句时，退出的是它所在的循环结构，对外层循环没有影响。

简而言之。跳出到上一层循环，不再执行循环\(结束当前的循环体\)

**2、continue语句**

continue跳出本次循环，继续执行下次循环\(结束正在执行的循环 进入下一个循环条件\) ；

**3、return语句**

return语句的话，比较狠，直接结束方法里的内容执行了。



# 



