# Chapter 6 Java常用类

一、 java日期处理类

1、Date类

Date类是jdk给我们提高的标准日期类，在java.util包下；

```
import java.util.Date;

public class TestDate {

    public static void main(String[] args) {
        Date date=new Date();
        System.out.println("当前日期："+date);
    }
}
```

2、Calendar类

Calendar是日历类，也是java.util包下的，功能比较强大，能获取到年月日时分秒的具体值；

```
public class TestCalendar {

    public static void main(String[] args) {
        Calendar calendar=Calendar.getInstance();
        System.out.println(calendar.get(Calendar.YEAR));
        System.out.println(calendar.get(Calendar.MONTH)+1); // 月份从0开始 要+1

        System.out.println("现在是："+calendar.get(Calendar.YEAR)+"年"
                +(calendar.get(Calendar.MONTH)+1)+"月"
                +calendar.get(Calendar.DAY_OF_MONTH)+"日"
                +calendar.get(Calendar.HOUR_OF_DAY)+"时"
                +calendar.get(Calendar.MINUTE)+"分"
                +calendar.get(Calendar.SECOND)+"秒");
    }
}
```

3、SimpleDateFormat类

SimpleDateFormat类主要是用作日期类型转换用的，在java.text包下：

```
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;

public class TestSimpleDateFormat {

    /**
     * 将日期对象格式化为指定格式的日期字符串
     * @param date 传入的日期对象
     * @param format 格式
     * @return
     */
    public static String formatDate(Date date,String format){
        String result="";
        SimpleDateFormat sdf=new SimpleDateFormat(format);
        if(date!=null){
            result=sdf.format(date);
        }
        return result;
    }

    /**
     * 将日期字符串转换成一个日期对象 
     * @param dateStr 日期字符串
     * @param format 格式
     * @return
     * @throws ParseException 
     */
    public static Date formatDate(String dateStr,String format) throws ParseException{
        SimpleDateFormat sdf=new SimpleDateFormat(format);
        return sdf.parse(dateStr);
    }

    public static void main(String[] args) throws ParseException {
        Date date=new Date();

        System.out.println(formatDate(date,"yyyy-MM-dd"));
        System.out.println(formatDate(date,"yyyy-MM-dd HH:mm:ss"));
        System.out.println(formatDate(date,"yyyy年MM月dd日HH时mm分ss秒"));

        String dataStr="1989-11-02 18:01:41";
        Date date2=formatDate(dataStr,"yyyy-MM-dd HH:mm:ss");
        System.out.println(formatDate(date2,"yyyy-MM-dd HH:mm:ss"));
    }
}
```

二、 String  StringBuffer **StringBuffer**

String ：字符串常量，适用于少量的字符串操作的情况

StringBuffer： 字符串变量（线程安全），适用于单线程下在字符缓冲区进行大量操作的情况

StringBuilder： 字符串变量（非线程安全）适用多线程下在字符缓冲区进行大量操作的情况

在效率上：StringBuilder &gt;StringBuffer &gt; String；

String最慢的原因：String为字符串常量，而StringBuilder和StringBuffer均为字符串变量，即String对象一旦创建之后该对象是不可更改的，但后两者的对象是变量，是可以更改的

三、 Math类

Math类是一个数学工具类方法，里面有很多静态工具方法；方便开发者直接调用；

在Math类中提供了众多数学函数方法，主要包括三角函数方法、指数函数方法、取整函数方法、取最大值、最小值以及平均值函数方法，这些方法都被定义为static形式，所以在程序中应用比较简便。可以使用如下形式调用：

```
Math.数学方法
```

在Math类中除了函数方法之外还存在一些常用数学常量，如圆周率、E等。这些数学常量作为Math类的成员变量出现，调用起来也很简单。可以使用如下形式调用：

```
Math.PI
Math.E
```

1、max方法 求最大值；

2、min方法 求最小值；

3、round方法 四舍五入；

4、pow方法 求次幂；

```
public class TestMath {

    public static void main(String[] args) {
        System.out.println("最大值："+Math.max(1, 2));
        System.out.println("最小值："+Math.min(1, 2));
        System.out.println("四舍五入："+Math.round(4.5));
        System.out.println("3的4次方："+Math.pow(3, 4));
    }
}
```

三角函数：

```
public class TrigonometricFunction {
    public static void main(String[] args) {
        // 取90度的正弦
        System.out.println("90度的正弦值：" + Math.sin(Math.PI / 2));
        System.out.println("0度的余弦值：" + Math.cos(0)); // 取0度的余弦
        // 取60度的正切
        System.out.println("60度的正切值：" + Math.tan(Math.PI / 3));
        // 取2的平方根与2商的反正弦
        System.out.println("2的平方根与2商的反弦值："
                + Math.asin(Math.sqrt(2) / 2));
        // 取2的平方根与2商的反余弦
        System.out.println("2的平方根与2商的反余弦值："
                + Math.acos(Math.sqrt(2) / 2));
        System.out.println("1的反正切值：" + Math.atan(1)); // 取1的反正切
        // 取120度的弧度值
        System.out.println("120度的弧度值：" + Math.toRadians(120.0));
        // 取π/2的角度
        System.out.println("π/2的角度值：" + Math.toDegrees(Math.PI / 2));
    }
}
```

指数函数：

```
public class ExponentFunction {
    public static void main(String[] args) {
        System.out.println("e的平方值：" + Math.exp(2)); // 取e的2次方
        // 取以e为底2的对数
        System.out.println("以e为底2的对数值：" + Math.log(2));
        // 取以10为底2的对数
        System.out.println("以10为底2的对数值：" + Math.log10(2));
        System.out.println("4的平方根值：" + Math.sqrt(4)); // 取4的平方根
        System.out.println("8的立方根值：" + Math.cbrt(8)); // 取8的立方根
        System.out.println("2的2次方值：" + Math.pow(2, 2)); // 取2的2次方
    }
}
```

取整函数：

```
public class IntFunction {
    public static void main(String[] args) {
        // 返回第一个大于等于参数的整数
        System.out.println("使用ceil()方法取整：" + Math.ceil(5.2));
        // 返回第一个小于等于参数的整数
        System.out.println("使用floor()方法取整：" + Math.floor(2.5));
        // 返回与参数最接近的整数
        System.out.println("使用rint()方法取整：" + Math.rint(2.7));
        // 返回与参数最接近的整数
        System.out.println("使用rint()方法取整：" + Math.rint(2.5));
        // 将参数加上0.5后返回最接近的整数
        System.out.println("使用round()方法取整：" + Math.round(3.4f));
        // 将参数加上0.5后返回最接近的整数，并将结果强制转换为长整型
        System.out.println("使用round()方法取整：" + Math.round(2.5));
    }
}
```

取最大值、最小值、绝对值函数：

```
public class AnyFunction {
    public static void main(String[] args) {
        System.out.println("4和8较大者:" + Math.max(4, 8)); 
         // 取两个参数的最小值
        System.out.println("4.4和4较小者：" + Math.min(4.4, 4));
        System.out.println("-7的绝对值：" + Math.abs(-7)); // 取参数的绝对值
    }
}
```

四、数字格式化——DecimalFormat类

DecimalFormat是NumberFormat的一个子类，用于格式化十进制数字。它可以将一些数字格式化为整数、浮点数、科学计数法、百分数等。通过使用该类可以为要输出的数字加上单位或控制数字的精度。一般情况下可以在实例化DecimalFormat对象时传递数字格式，也可以通过DecimalFormat类中的applyPattern\(\)方法来实现数字格式化。![](/assets/DecimalFormat.png)

```
public class DecimalFormatSimpleDemo {
    // 使用实例化对象时设置格式化模式
    static public void SimgleFormat(String pattern, double value) {
        // 实例化DecimalFormat对象
        DecimalFormat myFormat = new DecimalFormat(pattern); 
        String output = myFormat.format(value); // 将数字进行格式化
        System.out.println(value + " " + pattern + " " + output);
    }

    // 使用applyPattern()方法对数字进行格式化
    static public void UseApplyPatternMethodFormat(String pattern, double value) {
        DecimalFormat myFormat=new DecimalFormat();//实例化DecimalFormat对象
        myFormat.applyPattern(pattern); // 调用applyPatten()方法设置格式化模板
        System.out
                .println(value + " " + pattern + " " + myFormat.format(value));
    }

    public static void main(String[] args) {
        SimgleFormat("###,###.###", 123456.789);// 调用静态SimgleFormat()方法
        SimgleFormat("00000000.###kg", 123456.789); // 在数字后加上单位
        // 按照格式模板格式化数字，不存在的位以0显示
        SimgleFormat("000000.000", 123.78);
        // 调用静态UseApplyPatternMethodFormat()方法
        UseApplyPatternMethodFormat("#.###%", 0.789); // 将数字转换为百分数形式
        // 将小数点后格式化为两位
        UseApplyPatternMethodFormat("###.##", 123456.789);
        // 将数字转化为千分数形式
        UseApplyPatternMethodFormat("0.00\u2030", 0.789);
    }
}
```

五、获取随机数

1、Math.random\(\)

在Math类中存在一个random\(\)方法，用于产生随机数字，这个方法默认生成大于等于0.0小于1.0的double型随机数，即0&lt;=Math.random\(\)&lt;1.0，虽然Math.random\(\)方法只可以产生0~1之间的double型数字，其实只要在Math.random\(\)语句上稍加处理，就可以使用这个方法产生任意范围的随机数。

```
public class MathRandomChar {
    // 定义获取任意字符之间的随机字符
    public static char GetRandomChar(char cha1, char cha2) {
        return (char) (cha1 + Math.random() * (cha2 - cha1 + 1));
    }

    public static void main(String[] args) {
        // 获取a~z之间的随机字符
        System.out.println("任意小写字符" + GetRandomChar('a', 'z'));
        // 获取A~Z之间的随机字符
        System.out.println("任意大写字符" + GetRandomChar('A', 'Z'));
        // 获取0~9之间的随机字符
        System.out.println("0到9任意数字字符" + GetRandomChar('0', '9'));
    }
}
```

2、实例化Random

在Java中还提供了一种可以获取随机数的方式，那就是java.util.Random类。可以通过实例化一个Random对象创建一个随机数生成器。

语法如下：

```
Random r=new Random();
其中，r是指Random对象。
```

```
public class RandomDemo {
    public static void main(String[] args) {
        Random r = new Random(); // 实例化一个Random类
        // 随机产生一个整数
        System.out.println("随机产生一个整数:" + r.nextInt());
        // 随机产生一个大于等于0小于10的整数
        System.out.println("随机产生一个大于等于0小于10的整数：" + r.nextInt(10));
        // 随机产生一个布尔型的值
        System.out.println("随机产生一个布尔型的值：" + r.nextBoolean());
        // 随机产生一个双精度型的值
        System.out.println("随机产生一个双精度型的值：" + r.nextDouble());
        // 随机产生一个浮点型的值
        System.out.println("随机产生一个浮点型的值：" + r.nextFloat());
        // 随机产生一个概率密度为高斯分布的双精度值
        System.out.println("随机产生一个概率密度为高斯分布的双精度值："
                + r.nextGaussian());
    }
}
```

六、Arrays类

Arrays类主要是封装了很多操作数组的工具方法，方便开发者直接调用；

常用的方法；

1、toString\(\)方法 返回指定数组内容的字符串表示形式；

2、sort\(\)方法 对指定的类型数组按数字升序进行排序；

3、binarySearch\(\)方法 使用二分搜索法来搜索指定类型数组，以获取指定值；

4、fill\(\)方法，将指定类型值分配给指定类型数组的每个元素；

```
import java.util.Arrays;
 
public class TestArrays {
 
    public static void main(String[] args) {
        int arr[]={1,7,3,8,2};
        System.out.println(arr);
        System.out.println("以字符串形式输出数组："+Arrays.toString(arr));
        Arrays.sort(arr); // 给数组排序
        System.out.println("排序后的数组："+Arrays.toString(arr));
        System.out.println(Arrays.binarySearch(arr, 1));
        Arrays.fill(arr, 0); // 将指定内容填充到数组中
        System.out.println("填充数组后的字符串："+Arrays.toString(arr));
    }
}
```

七、

八、

九、

