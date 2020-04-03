# Java数组

## 一、数组的简介

数组是具有相同数据类型的一组数据的集合。在Java中同样可以将数组看做为一个对象，虽然基本数据类型不是对象，但由基本数据类型组成的数组却是对象。

根据数组的维数将数组分为一维数组和多位数组\(包括二维数组\)

## 二、一维数组

1、数组定义

数组作为对象允许使用new关键字进行内存分配。在使用数组之前，必须首先定义数组变量所属的类型，即声明数组，声明一维数组有两种形式，语法如下所示：

```
数组元素类型  数组名字[ ];
数组元素类型[ ]  数组名字;
```

数组元素类型：决定了数组的数据类型，它可以是Java中任意的数据类型，包括基本数据类型和非基本数据类型。  
数组名字：为一个合法的标识符  
符号“\[ \]”：指明该变量是一个数组类型变量，单个“\[ \]”表示要创建的数组是一维数组

```
public class Demo18 {

    public static void main(String[] args) {
        // 定义一个数组
        int []arr;
        // 定义一个数组
        int arr2[];
    }
}
```

2、内存分配

\(1\)先声明数组，再使用new运算符进行内存分配

```
int []arr;
arr=new int [5]
```

\(2\)声明的同时为数组分配内存

```
int arr2[]=new int [10]
```

![](/assets/数组内存分析.jpg)

主要分为 栈区和堆区 :

栈区主要存放引用地址

堆区主要存放大块的实际数据，比如对象，数组；

3、数组初始化（数组赋值）

数组可以与基本数据类型一样进行初始化操作，数组的初始化可分别初始化数组中每个元素。数组的初始化有两种形式。  
两种方法初始化一维数组，示例代码如下所示：

```
int arr[] = new int[]{1,2,3,5,25};    //第一种初始化方式
int arr2[] = {34,23,12,6};        //第二种初始化方式
```

数组的初始化方式是：把数据类型包括在大括号之内，中间用逗号分开数组元素的值，系统自动为数组分配一定的空间。

第一种初始化方式，创建5个元素的数组，其值依次为1、2、3、5、25；

第二种初始化方式，创建4个元素的数组，其值依次为34、23、12、6。

\(1\)静态初始化

格式 arrayName=new type\[\]{element1,element2,element3...}

实例：int arr\[\]=new int\[\]{1,2,3};

实例代码：

```
public class Demo18 {

    public static void main(String[] args) {
        // 定义一个数组，并且静态初始化
        int arr[]=new int[]{1,2,3};

        // 普通的遍历数组方式
        for(int i=0;i<arr.length;i++){
            System.out.println(arr[i]);
        }

        System.out.println("---------------");
        // foreach方式
        for(int j:arr){
            System.out.println(j);
        }
    }
}
```

\(2\)动态初始化

格式 arrayName=new type\[length\];

实例：int arr2\[\]=new int\[3\]

```
public class Demo18 {

    public static void main(String[] args) {
        // 定义一个数组，然后动态初始化，长度是3
        int arr[]=new int[3];
        // int类型的数组，默认是0
        for(int i=0;i<arr.length;i++){
            System.out.println(arr[i]);
        }
    }
}
```

## 三、二维数组

1、数组定义

声明二维数组的方法有两种，语法如下所示：

```
数组元素类型 数组名字[ ][ ];
数组元素类型[ ][ ] 数组名字;
```

数组元素类型：决定了数组的数据类型，它可以是Java中任意的数据类型，包括基本数据类型和非基本数据类型。  
数组名字：为一个合法的标识符  
符号“\[ \]”：指明该变量是一个数组类型变量，两个“\[ \]”表示要创建的数组是二维数组。

2、数组初始化（数组赋值）

二维数组的初始化同一维数组初始化类似，同样可以使用大括号完成二维数组的初始化。语法如下所示：

```
type arrayname[][] = {value1,value2…valuen};
type：数组数据类型
arrayname：数组名称，一个合法的标识符
value：数组中各元素的值
```

二维数组动态初始化 格式：

arrayName=new type\[length\]\[length\];

实例：int \[\]\[\]arr2=new int\[3\]\[3\];

静态初始化实例：

```
public class Demo19 {

    public static void main(String[] args) {
        int [][]arr=new int[][]{{1,2,3},{4,5,6},{7,8,9}};

        for(int i=0;i<arr.length;i++){
            for(int j=0;j<arr[i].length;j++){
                System.out.print(arr[i][j]+" ");
            }
            System.out.println();
        }
    }
}
```

动态初始化实例：

```
public class Demo19 {

    public static void main(String[] args) {
        int [][]arr=new int[3][3];

        for(int i=0;i<arr.length;i++){
            for(int j=0;j<arr[i].length;j++){
                System.out.print(arr[i][j]+" ");
            }
            System.out.println();
        }
    }
}
```

3、 二维数组的内存分析

![](/assets/二维数组的内存分析.jpg)

## 四、数组操作

1、遍历数组

遍历数组就是获取数组中的每个元素。通常遍历数组都是使用for循环来实现。遍历一维数组很简单，也很好理解。  
遍历二维数组要比遍历一维数组要麻烦一些，需使用双层for循环，还要通过数组的length属性获得数组的长度。

```
public class Rakel { // 创建类
    public static void main(String[] args) { // 主方法
        // 定义String型数组str
        String str[] = new String[] { "ab", "cd", "ef", "yz" };
        Arrays.sort(str); // 将数组进行排序
        // 在指定的范围内搜索元素"cd"的索引位置
        int index = Arrays.binarySearch(str, 0, 2, "cd");
        System.out.println("cd的索引位置是：" + index); // 将索引输出
    }
}
```

2、填充替换数组元素

数组中的元素定义完成后，可通过Arrays类的静态方法fill\(\)来对数组中的元素进行替换。该方法通过各种重载形式可完成任意类型的数组元素的替换。fill\(\)方法有两种参数类型。下面以int型数组为例介绍fill\(\)方法的使用方法。

```
(1)fill(int[] a,int value)将指定的值分配给int型数组的每个元素
(2)fill(int[] a,int fromindex,int toindex,int value)将指定的值分配给int型数组指定范围内的每个元素

```

```
public class Swap { // 创建类
    public static void main(String[] args) { // 主方法
        int arr[] = new int[5]; // 创建int型数组
        Arrays.fill(arr, 8); // 使用同一个值对数组进行填充
        for (int i = 0; i < arr.length; i++) { // 循环遍历数组中的元素
            // 将数组中的元素依次输出
            System.out.println("第" + i + "个元素是：" + arr[i]);
        }
    }
}
```

```
public class Displace { // 创建类
    public static void main(String[] args) { // 主方法
        int arr[] = new int[] { 45, 12, 2, 10 }; // 定义并初始化int型数组arr
        Arrays.fill(arr, 1, 2, 8); // 使用fill方法对数组进行初始化
        for (int i = 0; i < arr.length; i++) { // 循环遍历数组中元素
            // 将数组中的每个元素输出
            System.out.println("第" + i + "个元素是：" + arr[i]);
        }
    }
}
```

3、对数组进行排序

通过Arrays类的静态sort\(\)方法可实现对数组排序，sort\(\)方法提供了许多种重载形式，可对任意类型数组进行升序排序。语法如下所示：

```
Arrays.sort(object)
object：指进行排序的数组名称
返回值：排序后的数组
```

```
public class Taxis { // 创建类
    public static void main(String[] args) { // 主方法
        int arr[] = new int[] { 23, 42, 12, 8 }; // 声明数组
        Arrays.sort(arr); // 将数组进行排序
        for (int i = 0; i < arr.length; i++) { // 循环遍历排序后的数组
            System.out.println(arr[i]); // 将排序后数组中的各个元素输出
        }
    }
}
```

4、复制数组

Arrarys类的copyOf\(\)方法与copyOfRange\(\)方法可实现对数组的复制。copyOf\(\)方法是复制数组至指定长度，copyOfRange\(\)方法则将指定数组的指定长度复制到一个新数组中。

```
copyof(arr,int newlength)
copyofRange(arr,int fromindex,int toindex)
```

```
public class Cope { // 创建类
    public static void main(String[] args) { // 主方法
        int arr[] = new int[] { 23, 42, 12, }; // 定义数组
        int newarr[] = Arrays.copyOf(arr, 5); // 复制数组arr
        for (int i = 0; i < newarr.length; i++) { // 循环变量复制后的新数组
            System.out.println(newarr[i]); // 将新数组输出
        }
    }
}
```

```
public class Repeat { // 创建类
    public static void main(String[] args) { // 主方法
        int arr[] = new int[] { 23, 42, 12, 84, 10 }; // 定义数组
        int newarr[] = Arrays.copyOfRange(arr, 0, 3); // 复制数组
        for (int i = 0; i < newarr.length; i++) { // 循环遍历复制后的新数组
            System.out.println(newarr[i]); // 将新数组中的每个元素输出
        }
    }
}
```

5、数组查询

Arrays类的binarySearch\(\)方法，可使用二分搜索法来搜索指定数组，以获得指定对象。该方法返回要搜索元素的索引值。binarySearch\(\)方法提供了多种重载形式，用于满足各种类型数组的查找需要。

binarySearch\(\)方法有两种参数类型。

```
（1）binarySearch(Object[].Object key)
（2）binarySearch(Object[].,int fromIndex , int toIndex,Object key)
```

```
public class Example { // 创建类
    public static void main(String[] args) { // 主方法
        int ia[] = new int[] { 1, 8, 9, 4, 5 }; // 定义int型数组ia
        Arrays.sort(ia); // 将数组进行排序
        int index = Arrays.binarySearch(ia, 4); // 查找数组ia中元素4的索引位置
        System.out.println("4的索引位置是：" + index); // 将索引输出
    }
}
```

```
public class Rakel { // 创建类
    public static void main(String[] args) { // 主方法
        // 定义String型数组str
        String str[] = new String[] { "ab", "cd", "ef", "yz" };
        Arrays.sort(str); // 将数组进行排序
        // 在指定的范围内搜索元素"cd"的索引位置
        int index = Arrays.binarySearch(str, 0, 2, "cd");
        System.out.println("cd的索引位置是：" + index); // 将索引输出
    }
}
```

## 五、数组排序

1、冒泡排序

2、快速排序

