
---

IO\(输入/输出）是所有程序都必需的部分:使 用输入 机制，允许程序读取 外部数据（包括来自磁盘、光盘等存储设备的数据）、用户输入数据;使用输出机制，允许程序记录运行状态，将程序数据输出到磁盘、光盘等存储设备中。

Java的IO通过 java.io 包下的类和接口来支持，在 java.io 包下主要包括输入、输出两种IO流，每种输入、输出流又可分为字节流和字符流两大类，其中字节流以字节为单位来处理输入、输出操作，而字符流则以字符来处理输入、输出操作。除此之外，Java的IO流使用了一种装饰器设计模式，它将IO流分成底层节点流和上层处理流，其中节点流用于和底层的物理存储节点直接关联一不同的物理节点获取节点流的方式可能存在一定的差异，但程序可以把不同的物理节点流包装成统一的处理流，从而允许程序使用统一的输入、输出代码来读取不同的物理存储节点的资源。

一、File类

File类是java.io包下代表与平台无关的文件和目录，也就是说，如果希望在程序中操作文件和目录，都可以通过File类来完成.值得指出的是，不管是文件还是目录都是使用File来操作的，File能新建、删除、重命名文件和目录，File不能访问文件内容本身。如果需要访问文件内容本身，则需要使用输入/输出流.

1.1访问文件和目录

File 类可以使用文件路径字符串来创建File实例，该文件路径字符串既可以是绝对路径，也可以是相对路径。在默认情况下.系统总是依据用户的工作路径来解释相对路径,这个路径由系统属性“ userdir ”指定，通常也就是运行 Java 虚拟机时所在的路径。

一旦创建了File对象后，就可以调用File对象的方法来访问，File类提供了很多方法来搡作文件和目录，下面列出一些比较常用的方法.

1.访问文件名相关的方法

\(1\)String getName\(\):返回此File对象所表示的文件名或路径名（如果是路径，则返回最后一级子路径名）。

\(2\)String getPath \(\):返回此File对象所对应的路径名，

\(3\)File getAbsoluteFile\(\): 返 回 此File对象的绝对路径.

\(4\)Siring getAbsolutePath\(\): 返 回 此File对象所对应的绝对路径名.

\(5\)String getParent\(\): 返 回 此File对象所对应目录（最后一级子目录）的父目录名，

\(6\)boolean renameTo\(\)\( FilenewNanie\): 重 命 名 此File对象所对应的文件或目录，如果重命名成功，则返回true;否则返回false.

2.文件检测相关的方法

\(1\)boolean exists\(\) :判断File对象所对应的文件或目录是否存在，

\(2\)boolean canWrite\(\):判断File对象所对应的文件和目录是否可写.

\(3\)boolean canRead\(\): 判 断File对象所对应的文件和目录是否可读.

\(4\)boolean isFilc\(\):判断File对象所对应的是否是文件，而不是目录.

\(5\)boolean isDirectory\(\): 判 断File对象所对应的是否是目录，而不是文件。

\(6\)boolean isAbsolute\(\):判断File对象所对应的文件或目录是否是绝对路径.该方法消除了不同平台的差异，可以直接判断File对象是否为绝对路径。在 UNIX / Linux / BSD 等系统上，如果路径名开头是一条斜线（/），则衣明该File对柒对应一个绝对路径：在 Windows 等系统上，如果路径幵头是盘符，则说明它是—个绝对路径。

3.获取常规文件信息

\(1\)long lastModified\(\):返回文件的最后修改时间。

\(2\)long lenglh\(\):返回文件内容的长度.

4.文件操作相关的方法

\(1\)boolean createNewFile\(\):当此File对象所对应的文件不存在时，该方法将新建一个该File对象所指定的新文件，如果创建成功则返冋 true ;否则返回 false 。

\(2\)boolean delete\(\):删除File对象所对应的文件或路劲.

\(3\)static File createTempFile\(String prefix , String suffix \):在默认的临时文件目录中创建一个临时的空文件 . 使用给定前缀 、 系统生成的随机数和给定后缀作为文件名 .这是一个静态方法 ，可以直接通 过File类来调用 。 prefix参数必须至少是3字节长 。建议前缀使用一个短的、有意义的字符串。比如“hjb”或"mail”。suffix 参数可以为null, 在这种情况下 ，将使用默认的后缀 “. tmp ”。

\(3\)static File createTempFile\(String prefix ,  String suffix ，Filedirectory \):在  directory 所指定的目录中创建一个临时的空文件 。使用给定前缀 、系统生成的随机数和给定后缀作为文件名. 这是一个静态方法，可以直接通过File类来调用。

\(4\)void deleteOnExit\(\): 注册一个刪除钩子，指定当 Java 虚拟机沿出时.删除File对象所对应的文件和目录。

5.目录操作相关的方法

\(1\)boolean mkdir\(\):试阁创建一个File对象所对应的目录，如果创建成功.则返回true :否则返回false 。调用该方法时File对象必须对应一个路径，而不是一个文件。

\(2\)String\[\] list\(\) :列出File对象的所有子文件名和路径名，返回 String 数组.

\(3\)File\[\] listFile\(\):列出File对象的所有子文件和路径.返回File数组.

\(4\)static File\[\]  listRoots\(\) : 列出系统所有的根路径.这是一个静态方法，可以直接通过File类来调用.

```
public class FileTest
{
    public static void main(String[] args)
        throws IOException
    {
        // 以当前路径来创建一个File对象
        File file = new File(".");
        // 直接获取文件名，输出一点
        System.out.println(file.getName());
        // 获取相对路径的父路径可能出错，下面代码输出null
        System.out.println(file.getParent());
        // 获取绝对路径
        System.out.println(file.getAbsoluteFile());
        // 获取上一级路径
        System.out.println(file.getAbsoluteFile().getParent());
        // 在当前路径下创建一个临时文件
        File tmpFile = File.createTempFile("aaa", ".txt", file);
        // 指定当JVM退出时删除该文件
        tmpFile.deleteOnExit();
        // 以系统当前时间作为新文件名来创建新文件
        File newFile = new File(System.currentTimeMillis() + "");
        System.out.println("newFile对象是否存在：" + newFile.exists());
        // 以指定newFile对象来创建一个文件
        newFile.createNewFile();
        // 以newFile对象来创建一个目录，因为newFile已经存在，
        // 所以下面方法返回false，即无法创建该目录
        newFile.mkdir();
        // 使用list()方法来列出当前路径下的所有文件和路径
        String[] fileList = file.list();
        System.out.println("====当前路径下所有文件和路径如下====");
        for (String fileName : fileList)
        {
            System.out.println(fileName);
        }
        // listRoots()静态方法列出所有的磁盘根路径。
        File[] roots = File.listRoots();
        System.out.println("====系统所有根路径如下====");
        for (File root : roots)
        {
            System.out.println(root);
        }
    }
}
```

1.2文件过滤器

在File类的list\(\) 方法中可以接收一个 FilenameFilter 参数，通过该参数可以只列出符合条件的文件。这里的 FilenameFilter 接口和 javax.swing.filechooser 包下的 FileFilter 抽象类的功能非常相似，可以把FileFilter 当成 FilenameFilter 的实现类，但可能 Sun 在设计它们时产生了一些小小遗漏，所以没有让FileFilter  实现  FilenameFilter  接口。

FilenameFilter 接口里包含了一个 accept\(File dir ,  String name \)方法，该方法将依次对指定File的所有子目录或者文件进行迭代，如果该方法返回true，则list\(\)方法会列出该子目录或者文件。

FilenameFilter接口内只有一个抽象方法，因此该接口也是一个函数式接口，可使用Lambda表达式创建实现该接口的对象。

```
public class FilenameFilterTest
{
    public static void main(String[] args)
    {
        File file = new File(".");
        // 使用Lambda表达式（目标类型为FilenameFilter）实现文件过滤器。
        // 如果文件名以.java结尾，或者文件对应一个路径，返回true
        String[] nameList = file.list((dir, name) -> name.endsWith(".java")
            || new File(name).isDirectory());
        for(String name : nameList)
        {
            System.out.println(name);
        }
    }
}
```

二、IO流

Java的IO流是实现输入/输出的基础，它可以方便地实现数据的输入/输出操作，在Java中把不同的输入/输出源（键盘、文件、网络连接等）抽象表述为“流”（stream\),通过流的方式允许Java程序使用相同的方式来访问不同的输入/输出源。stream是从起源（source\)到接收（sink\)的有序数据。

Java把所有传统的流类型（类或抽象类）都放在java.io包中，用以实现输入/输出功能。

2.1流的分类

按照不同的分类方式，可以将流分为不同的类型，下面从不同的角度来对流进行分类，它们在概念上可能存在重叠的地方。

1.输入流和输出流

按照流的流向来分，可以分为输入流和输出流。

\(1\)输入流：只能从中读取数据，而不能向其写入数据。

\(2\)输出流：只能向其写入数据，而不能从中读取数据。

这里的输入、输出都是从程序运行所在内存的角度来划分的。

Java 的输入流主要由InputStream 和 Reader 作为基类，而输出流则主要由 OutputStream 和 Writer作为基类。它们都是一些抽象基类，无法直接创建实例。

2.字节流和字符流

字节流和字符流的用法几乎完全一样，区别在于字节流和字符流所搡作的数据单元不同——字节流操作的数据单元是8位的字节，而字符流搡作的数据单元是16位的字符。

字节流主要由 InputStream 和 OutputStream 作为基类，而字符流则主要由Reader和Writer作为基类。

3.节点流和处理流

按照流的角色来分，可以分为节点流和处理流。

可以从/向一个特定的IO设备（如磁盘、网络） 读/写数据的流，称为节点流，节点流也被称为低级流。

处理流则用于对一个己存在的流进行连接或封装，通过封装后的流来实现数据读功能。处理流也被称为高级流。

实际上 ，Java 使用处理流来包装节点流是一种典型的装饰器设计橒式，通过使用处理流来包装不同的节点流，既可以消除不同节点流的实现差异，也可以提供更方便的方法来完成输入/输出功能.因此处理流也被称为包装流.

2 . 2 流的概念模型

Java把所有设备里的有序数据抽象成流模型，简化了输入/输出处理，理解了流的概念模型也就了解了JavaIO.

Java的IO流共涉及40多个类。但是Java的IO流的40多个类都是从如下4个抽象基类派生的。

\(1\)InputStrcam/Reader:所冇输入流的基类，前者是字节输入流，后者是字符输入流.

\(2\)OutputStreamAVritcr:所有输出流的\*类，前者是字节输出流，后者是字符输出流.

对于InpmStream和Reader而言，它们把输入设备抽象成一个"水管”，这个水管里的毎个“水滴"依次抹列.

对于OuiputStream和Writer而言，它们同样把输出设备抽象成一个"水管”.只是这个水管里没有任何水滴.

Java的处理流模型则体现了Java

输入/输出流设计的灵活性。处理流的功能主要体现在以下两个方面：

\(1\)性能的提髙：主要以增加缓冲的方式来提高输入出的效率。

\(2\)操作的便捷：处理流可能提供了一系列便捷的方认来一次输入/大批量的内容.而不是输入/输出一个成多个“水滴'处理流可以“嫁接”在任何已存在的流的基础之上。这就允许Java应用程序采用相同的代码、透明的方式来访N不阀的输入/输出设各的数据流。

三、字节流和字符流

3.1 InputStream和Reader

InputStream 和 Reader 是所有输入流的抽象基类，本身并不能创建实例来执行输入，但它们将成为所有输入流的模板，所以它们的方法是所冇输入流都可使用的方法。

在InputStream 里包含如下三个方法。

\(1\)int read\(\):从输入流中读取单个字节。返回所读取的字节数据（字节数据可直接转换为 int 类型\)。

\(2\)int read\(byte\[\] b \):从输入流中最多读取 b.length 个字节的数据，并将其存储在字节数组 b中，返回实际读取的字节数。

\(3\)int read\(byte\[\] b ，int off,int len\):从输入流中最多读取len个字节的数据，并将其存储在数组b中，放入数组 b 中时，并不是从数组起点开始，而是从 off 位置开始，返回实际读取的字节数。

在 Reader 里包含如下三个方法。

\(1\)int read\(\) :从输入流中读取单个字符（相当于水管中取出一滴水），返回所读取的字符数据（字符数据可直接转换为 int 类型）

\(2\)int read \( char\[\]  cbuf \):从输入流中最多读取 cbuf.length 个字符的数据，并将其存储在字符数组 cbuf中，返冋实际读取的字符数。

\(3\)int read \( char\[\] cbuf,int off,int len \):从输入流中最多读取len个字符的数据，并将其存储在字符数组cbuf中，放入数组 cbuf 中时，并不楚从数组起点开始，而是从 off 位置开始，返回实际读取的字符数。

```
public class FileInputStreamTest
{
    public static void main(String[] args) throws IOException
    {
        // 创建字节输入流
        FileInputStream fis = new FileInputStream(
            "FileInputStreamTest.java");
        // 创建一个长度为1024的“竹筒”
        byte[] bbuf = new byte[1024];
        // 用于保存实际读取的字节数
        int hasRead = 0;
        // 使用循环来重复“取水”过程
        while ((hasRead = fis.read(bbuf)) > 0 )
        {
            // 取出“竹筒”中水滴（字节），将字节数组转换成字符串输入！
            System.out.print(new String(bbuf , 0 , hasRead ));
        }
        // 关闭文件输入流，放在finally块里更安全
        fis.close();
    }
}
```

```
public class FileReaderTest
{
    public static void main(String[] args)
    {
        try(
            // 创建字符输入流
            FileReader fr = new FileReader("FileReaderTest.java"))
        {
            // 创建一个长度为32的“竹筒”
            char[] cbuf = new char[32];
            // 用于保存实际读取的字符数
            int hasRead = 0;
            // 使用循环来重复“取水”过程
            while ((hasRead = fr.read(cbuf)) > 0 )
            {
                // 取出“竹筒”中水滴（字符），将字符数组转换成字符串输入！
                System.out.print(new String(cbuf , 0 , hasRead));
            }
        }
        catch (IOException ex)
        {
            ex.printStackTrace();
        }
    }
}
```

除此之外，InputStream 和 Reader 还支持如下几个方法来移动记录指针。

\(1\)void mark \( int readAheadLimit \):在记录指针当前位置记录一个标记（ mark \)。

\(2\)booleanmarkSupportedO :判断此输入流是否支持 mark \(\)操作，即是否支持记录标记。

\(3\)void reset \(\):将此流的记录指针重新定位到上一次记录标记（ mark \)的位置。

\(4\)long skip\(long n\):记录指针向前移动n个字节/字符.

3.2 OutputStream 和 Writer

OutputStream 和 Writer 也非常相似，两个流都提供了

如下三个方法。

\(1\) void write \( int c \):将指定的字节/字符输出到输出流中，其中 c 既可以代表字节，也可以代表字符。

\(2\) void write \( byte\[\]/char\[\] buf \):将字节数组/字符数组中的数据输出到指定输出流中。

\(3\) void write \( byte\[\]/char\[\] buf ,  intoff ， intlen \):将字节数组/字符数组中从 off 位置开始，长度为 len的字节/字符输出到输出流中。

因为字符流直接以字符作为操作单位，所以 Writer 可以用字符串来代替字符数组，即以 String 对象作为参数， Writer 里还包含如下两个方法。

\(1\) void write \( String str \):将 str 字符串里包含的字符输出到指定输出流中。

\(2\) void write \( String str ,  int off ， int len \):将 str 字符串里从 off 位置开始，长度为len 的字符输出到指定输出流中。

```
public class FileOutputStreamTest
{
    public static void main(String[] args)
    {
        try(
            // 创建字节输入流
            FileInputStream fis = new FileInputStream(
                "FileOutputStreamTest.java");
            // 创建字节输出流
            FileOutputStream fos = new FileOutputStream("newFile.txt"))
        {
            byte[] bbuf = new byte[32];
            int hasRead = 0;
            // 循环从输入流中取出数据
            while ((hasRead = fis.read(bbuf)) > 0 )
            {
                // 每读取一次，即写入文件输出流，读了多少，就写多少。
                fos.write(bbuf , 0 , hasRead);
            }
        }
        catch (IOException ioe)
        {
            ioe.printStackTrace();
        }
    }
}
```

```
public class FileWriterTest
{
    public static void main(String[] args)
    {
        try(
            FileWriter fw = new FileWriter("poem.txt"))
        {
            fw.write("锦瑟 - 李商隐\r\n");
            fw.write("锦瑟无端五十弦，一弦一柱思华年。\r\n");
            fw.write("庄生晓梦迷蝴蝶，望帝春心托杜鹃。\r\n");
            fw.write("沧海月明珠有泪，蓝田日暖玉生烟。\r\n");
            fw.write("此情可待成追忆，只是当时已惘然。\r\n");
        }
        catch (IOException ioe)
        {
            ioe.printStackTrace();
        }
    }
}
```

四、输入输出体系

4.1处理流的使用

使用处理流时的典型思路是，使用处理流来包装节点流，程序通过处理流来执行输入/输出功能，让节点流与底层的I/O设备、文件交互。

实际上识别处理流非常简单，只要流的构造器参数不是一个物理节点，而是已经存在的流，那么这种流就一定是处理流：而所有节点流都是直接以物理IO节点作为构造器参数的。

关于使用处理流的优势，归纳起来就是两点：

①对开发人员来说，使用处理流进行输入/输出操作更简单；

②使用处理流的执行效率更高.

```
public class PrintStreamTest
{
    public static void main(String[] args)
    {
        try(
            FileOutputStream fos = new FileOutputStream("test.txt");
            PrintStream ps = new PrintStream(fos))
        {
            // 使用PrintStream执行输出
            ps.println("普通字符串");
            // 直接使用PrintStream输出对象
            ps.println(new PrintStreamTest());
        }
        catch (IOException ioe)
        {
            ioe.printStackTrace();
        }
    }
}
```

4.2输入输出流体系

![](/assets/IO.png)

字符串作为物理节点的字符输入输出流的作用：

```
public class StringNodeTest
{
    public static void main(String[] args)
    {
        String src = "从明天起，做一个幸福的人\n"
            + "喂马，劈柴，周游世界\n"
            + "从明天起，关心粮食和蔬菜\n"
            + "我有一所房子，面朝大海，春暖花开\n"
            + "从明天起，和每一个亲人通信\n"
            + "告诉他们我的幸福\n";
        char[] buffer = new char[32];
        int hasRead = 0;
        try(
            StringReader sr = new StringReader(src))
        {
            // 采用循环读取的访问读取字符串
            while((hasRead = sr.read(buffer)) > 0)
            {
                System.out.print(new String(buffer ,0 , hasRead));
            }
        }
        catch (IOException ioe)
        {
            ioe.printStackTrace();
        }
        try(
            // 创建StringWriter时，实际上以一个StringBuffer作为输出节点
            // 下面指定的20就是StringBuffer的初始长度
            StringWriter sw = new StringWriter())
        {
            // 调用StringWriter的方法执行输出
            sw.write("有一个美丽的新世界，\n");
            sw.write("她在远方等我,\n");
            sw.write("哪里有天真的孩子，\n");
            sw.write("还有姑娘的酒窝\n");
            System.out.println("----下面是sw的字符串节点里的内容----");
            // 使用toString()方法返回StringWriter的字符串节点的内容
            System.out.println(sw.toString());
        }
        catch (IOException ex)
        {
            ex.printStackTrace();
        }
    }
}
```

4.3转换流

输 入 / 输 出 流 体 系 中 还 提 供 了 两 个 转 换 流 ， 这 两 个 转 换 流 用 于 实 现 将 字 节 流 转 换 成 字 符 流 ， 其 中

InputStreamReader 将字节输入流转换成字符输入流 I OutputStreamWriter 将字节输出流转换成字符输出流。

```
public class KeyinTest
{
    public static void main(String[] args)
    {
        try(
            // 将Sytem.in对象转换成Reader对象
            InputStreamReader reader = new InputStreamReader(System.in);
            // 将普通Reader包装成BufferedReader
            BufferedReader br = new BufferedReader(reader))
        {
            String line = null;
            // 采用循环方式来一行一行的读取
            while ((line = br.readLine()) != null)
            {
                // 如果读取的字符串为"exit"，程序退出
                if (line.equals("exit"))
                {
                    System.exit(1);
                }
                // 打印读取的内容
                System.out.println("输入内容为:" + line);
            }
        }
        catch (IOException ioe)
        {
            ioe.printStackTrace();
        }
    }
}
```

4.4 推回输入流

在输入／输出流体系里有两个特殊的流，就是 Pushbacklnputslream 和 PushbackRcader ·。它们都提供了如下三个方法。

\(1\) void unread \( byte\[\]/char\[\] buf\) ：将一个字节／字符数组内容推回到推回缓冲区里，从而允许重复读取刚刚读取的内容。

\(2\)void unread\(  byte\[\]/char\[\] buf, int off . int len \) ：将 一个字节/字符数组里 从 off 开始，长度为 len 字节/字符的内容推回到推回缓冲区里，从而允许吹复读取刚刚读取的内容．

\(3\) void unread \( int b\):将一个字节／字符推回到推回缓冲区里，从而允许重复读取刚刚读取的内容

这三个方法与 Inputstream 和 Rcadcr 中的三个 read\(\)方法一一对应，这二个方法就是 Pushbacklnputstrcam 和 PushbackReader 的奥秘所在．这两个推回输入流都带有一个推回缓冲区，当程序调用这两个推回输入流的 unread（）方法时，系统将会把指定数组的内容推回到该缓冲区里，而推回输入流每次调用 read（） 方法时总是先从推回缓冲区读取，只有完全读取了推回缓冲区的内容后， 但还没有装满 read（） 所需的数组时才会从原输入流中读取。

```
public class PushbackTest
{
    public static void main(String[] args)
    {
        try(
            // 创建一个PushbackReader对象，指定推回缓冲区的长度为64
            PushbackReader pr = new PushbackReader(new FileReader(
                "PushbackTest.java") , 64))
        {
            char[] buf = new char[32];
            // 用以保存上次读取的字符串内容
            String lastContent = "";
            int hasRead = 0;
            // 循环读取文件内容
            while ((hasRead = pr.read(buf)) > 0)
            {
                // 将读取的内容转换成字符串
                String content = new String(buf , 0 , hasRead);
                int targetIndex = 0;
                // 将上次读取的字符串和本次读取的字符串拼起来，
                // 查看是否包含目标字符串, 如果包含目标字符串
                if ((targetIndex = (lastContent + content)
                    .indexOf("new PushbackReader")) > 0)
                {
                    // 将本次内容和上次内容一起推回缓冲区
                    pr.unread((lastContent + content).toCharArray());
                    // 重新定义一个长度为targetIndex的char数组
                    if(targetIndex > 32)
                    {
                        buf = new char[targetIndex];
                    }
                    // 再次读取指定长度的内容（就是目标字符串之前的内容）
                    pr.read(buf , 0 , targetIndex);
                    // 打印读取的内容
                    System.out.print(new String(buf , 0 ,targetIndex));
                    System.exit(0);
                }
                else
                {
                    // 打印上次读取的内容
                    System.out.print(lastContent);
                    // 将本次内容设为上次读取的内容
                    lastContent = content;
                }
            }
        }
        catch (IOException ioe)
        {
            ioe.printStackTrace();
        }
    }
}
```

五、重定向标准输入/输出

Java 的标准输入/输出分别通过 System.in 和 System.out 来代衣，在默认情况下它们分别代表键盘和显示器，当程序通过 System.in 来获取输入时，实际上是从键盘读取输入：当程序试图通过 System.mit 执行输出时，稈序总是输出到屏幕。

在 System 类里提供了如下三个重定向标准输入/输出的方法。

\(1\)static void setErr\(PrintStream err \):重定向“标准”错误输出流。

\(2\)static void setIn\(PrintStream in \):重定向“标准”输入流。

\(3\)static void setOut\(PrintStream out \) : 重定向 “标准”输出流。

下面程序通过重定向标准输出流，将 System.out 的输出重定向到文件输出，而不是在屏幕上输出。

```
public class RedirectOut
{
    public static void main(String[] args)
    {
        try(
            // 一次性创建PrintStream输出流
            PrintStream ps = new PrintStream(new FileOutputStream("out.txt")))
        {
            // 将标准输出重定向到ps输出流
            System.setOut(ps);
            // 向标准输出输出一个字符串
            System.out.println("普通字符串");
            // 向标准输出输出一个对象
            System.out.println(new RedirectOut());
        }
        catch (IOException ex)
        {
            ex.printStackTrace();
        }
    }
}
```

下面程序通过重定向标准输入流，将 System.in的输出重定向到指定文件，而不是键盘输入。

```
public class RedirectIn
{
    public static void main(String[] args)
    {
        try(
            FileInputStream fis = new FileInputStream("RedirectIn.java"))
        {
            // 将标准输入重定向到fis输入流
            System.setIn(fis);
            // 使用System.in创建Scanner对象，用于获取标准输入
            Scanner sc = new Scanner(System.in);
            // 增加下面一行将只把回车作为分隔符
            sc.useDelimiter("\n");
            // 判断是否还有下一个输入项
            while(sc.hasNext())
            {
                // 输出输入项
                System.out.println("键盘输入的内容是：" + sc.next());
            }
        }
        catch (IOException ex)
        {
            ex.printStackTrace();
        }
    }
}
```

六、Java 虚拟机读写其他进程的数据

使用Runtime对象的exec\(\)方法可以运行平台上的其他程序，该方法产生一个 Process 对象， Process 对象代表由该 Java 程序启动的子进程。 Process 类提供了如下三个方法，用于让程序和其子进程进行通信。

\(1\)InputStream getErrorStream\(\): 获取子进程的错误流。

\(2\)InputStream getlnputStream\(\): 获取子进程的输入流。

\(3\)OutputStream getOutputStream \(\): 获取子进程的输出流.

下面程序示范了读取其他进程的输出信息。

```
public class ReadFromProcess
{
    public static void main(String[] args)
        throws IOException
    {
        // 运行javac命令，返回运行该命令的子进程
        Process p = Runtime.getRuntime().exec("javac");
        try(
            // 以p进程的错误流创建BufferedReader对象
            // 这个错误流对本程序是输入流，对p进程则是输出流
            BufferedReader br = new BufferedReader(new
                InputStreamReader(p.getErrorStream())))
        {
            String buff = null;
            // 采取循环方式来读取p进程的错误输出
            while((buff = br.readLine()) != null)
            {
                System.out.println(buff);
            }
        }
    }
}
```

```
public class WriteToProcess
{
    public static void main(String[] args)
        throws IOException
    {
        // 运行java ReadStandard命令，返回运行该命令的子进程
        Process p = Runtime.getRuntime().exec("java ReadStandard");
        try(
            // 以p进程的输出流创建PrintStream对象
            // 这个输出流对本程序是输出流，对p进程则是输入流
            PrintStream ps = new PrintStream(p.getOutputStream()))
        {
            // 向ReadStandard程序写入内容，这些内容将被ReadStandard读取
            ps.println("普通字符串");
            ps.println(new WriteToProcess());
        }
    }
}
// 定义一个ReadStandard类，该类可以接受标准输入，
// 并将标准输入写入out.txt文件。
class ReadStandard
{
    public static void main(String[] args)
    {
        try(
            // 使用System.in创建Scanner对象，用于获取标准输入
            Scanner sc = new Scanner(System.in);
            PrintStream ps = new PrintStream(
            new FileOutputStream("out.txt")))
        {
            // 增加下面一行将只把回车作为分隔符
            sc.useDelimiter("\n");
            // 判断是否还有下一个输入项
            while(sc.hasNext())
            {
                // 输出输入项
                ps.println("键盘输入的内容是：" + sc.next());
            }
        }
        catch(IOException ioe)
        {
            ioe.printStackTrace();
        }
    }
}
```

七、RandomAccessFile

RandomAccessFile是Java输入/输出流体系中功能最丰富的文件内容访问类，它提供了众多的方法来访问文件内容，它既可以读取文件内容，也可以向文件输出数据.与普通的输入/输出流不同的是，RandomAccessFile 支持“随机访问”的方式，程序接跳转到文件的任意地方来读写数据。

由于RandomAccessFile 可以自由访问文件的任意位置,所以如果只需要访问文件部分内容，而不是把文件从头读到尾 ，使用 RandomAccessFile 将是更的选择。与 OutputStream 、 Writer 等输出流不同的是， RandomAccessFile 允许自由定位文件记录指针，RandomAccessFile可以不从开始的地方开始输出，因此 RandomAccessFile 可以向已存在的文件后追加内容。如果程序需要向已存在的文件后追加内容，则应该使用 RandomAccessFile 。

RandomAccessFile 的方法虽然多，但它有一个最大的局限，就是只能读写文件，不能读写其他 IO节点。

RandomAccessFile 对象也包含了一个记录指针，用以标识当前读写处的位背，当程序新创迚一个RandomAccessFile 对象时，该对象的文件记录指针位于文件头（也就是0处），当读/写了n个字节后，文件记录指针将会向后移动 n 个字节•除此之外， RandomAccessFile 可以自由移动该记录指针，既可以向前移动，也可以向后移动。 RandomAccessFile 也含了如下两个方法来操作文件记录指针。

\(1\)long getFilePointer\(\):返回文件记录指针的当前位置。

\(2\)void seek\(long pos\):将文件记录指针定位到pos位置。

RandomAccessFile 既可以读文件，也可以写，所以它既包含了完全类似于 InputStream 的三个read \(\)方法，其用法和InputStream 的三个 read\(\)方法完全一样：也包含了完全类似 OutputStream 的三个 write\(\)方法，其用法和 OutputStream 的三个 write\(\)方法完全一样。除此之外， RandomAccessFile 还包含了一系列的 readXxx\(\)和 writeXxx \(\)方法来完成输入、输出。

RandomAccessFilc 类有两个构造器，其实这两个构造器基本相同，只足指定文件的形式不同而已一个使用 String 参数来指定文件名，一个使用 File 参数来指定文件本身。除此之外，创建 RandomAccessFile对象时还需要指定一个 mode 参数，该参数指定 RandomAccessFile 的访问模式，该参数有如下4个值。

\(1\)"r":以只读方式打幵指定文件.如果试图对该 RandomAccessFile 执行写入方法，都将抛出IOException 异常。

\(2\)" rw ":以读、写方式打开指定文件。如果该文件尚不存在，则尝试创建该文件。

\(3\)" rws " : 以 读 、写方式打开指定文件。相对于" rw ”模式，还要求对文件的内容或元数据的每个更新都同步写入到底层存储设备。

\(4\)" rwd ":以读、写方式打幵指定文件。相对 于" rw "模式，还要求对文件内容的每个更新都同步写入到底层存储设备。

下面程序使用了RandomAccessFile 来访问指定的中间部分数据。

```
public class RandomAccessFileTest
{
    public static void main(String[] args)
    {
        try(
            RandomAccessFile raf =  new RandomAccessFile(
                "RandomAccessFileTest.java" , "r"))
        {
            // 获取RandomAccessFile对象文件指针的位置，初始位置是0
            System.out.println("RandomAccessFile的文件指针的初始位置："
                + raf.getFilePointer());
            // 移动raf的文件记录指针的位置
            raf.seek(300);
            byte[] bbuf = new byte[1024];
            // 用于保存实际读取的字节数
            int hasRead = 0;
            // 使用循环来重复“取水”过程
            while ((hasRead = raf.read(bbuf)) > 0 )
            {
                // 取出“竹筒”中水滴（字节），将字节数组转换成字符串输入！
                System.out.print(new String(bbuf , 0 , hasRead ));
            }
        }
        catch (IOException ex)
        {
            ex.printStackTrace();
        }
    }
}
```

八、对象序列化

对象序列化的目标是将对象保存到磁盘中，或允许在网络中直接传输对象.对象序列化机制允许把

内存中的 Java 对象转换成平台无关的二进制流，从而允许把这种二进制流持久地保存在磁盘，通过

网络将这种二进制流传输到另一个网络节点。其他程序一旦获得了这种二进制流（无论是从磁盘中获取

的，还是通过网络获取的），都可以将这种二进制流恢复成原来的 Java 对象。

8.1序列化的含义和意义

序列化机制允许将实现序列化的Java 对象转换成字序列，这些字节序列可以保存在磁盘上 ,或

通过网络传输，以备以后重新恢复成原来的对象.序列化机制使得对象可以脱离程序的运行而独立存在.

对象的序列化（ Serialize \) 指将一个Java对象写入IO流中 ，与此对应的是 ，对象的反序列化

\( Deserialize \)则指从IO流中恢复该 Java 对象，

如果需要让某个对象支持序列化机制，则必须让它的类是可序列化的（ serializable 〉.为了让某个

类是可序列化的，该类必须实现如下两个接口之一：

\(1\)Serializable

\(2\)Extemalizable

Java的很多类己经实现了Serializable ，该接口是一个标记接口，实现该接口无须实现任何方法，它

只是表明该类的实例是可序列化的。

所有可能在网络上传输的对象的类都应该是可序列化的。否则程序将会出现异常。 比如RMI

\( Remote Method Invoke ,即远程方法调用，是JavaEE的基础）过程中的參数和返回值;所有需要保存

到磁盘电的对象的类都必须可序列化，比如 Web 应用中需要保存到 HttpSession 或 ServletContext 属性的Java对象。

因为序列化是RMI过程的参数和返回值都必须实现的机制，而 RMI又是Java EE 技术的基础——

所有的分布式应用常常需要跨平台、跨网络.所以要求所有传递的参数、返回值必须实现序列化。因此

序列化机制是 Java EE平台的基础.通常建议：程序创建的每个 JavaBean 类都实现 Serializable .

8.2使用对象流实现序列化

如采需要将菜个对象保存到磁盘上或者通过网络传输，那么这个类应该实现 Serializable 接口或者

Extemalizable 接口之一。

使用Serializable 来实现序列化非常简中.主要让目标类实现 Serializable 标记接口即可，无须实现

任何方法。

一旦某个类实现了  Serializable 接口，该类的对象就是可序列化的，程序可以通过如下两个步骤来

序列化该对象。

\(1\)创建一个 ObjectOutputStream, 这个输出流是一个处理流，所以必须建立在其他节点流的基础之上。

\(2\)调用 ObjectOutputStream 对象的 writeObject \(\)方法输出可序列化对象.

```
public class Person implements java.io.Serializable
{
    private String name;
    private int age;
    // 注意此处没有提供无参数的构造器!
    public Person(String name , int age)
    {
        System.out.println("有参数的构造器");
        this.name = name;
        this.age = age;
    }
    // 省略name与age的setter和getter方法

    // name的setter和getter方法
    public void setName(String name)
    {
        this.name = name;
    }
    public String getName()
    {
        return this.name;
    }

    // age的setter和getter方法
    public void setAge(int age)
    {
        this.age = age;
    }
    public int getAge()
    {
        return this.age;
    }
}


public class WriteObject
{
    public static void main(String[] args)
    {
        try(
            // 创建一个ObjectOutputStream输出流
            ObjectOutputStream oos = new ObjectOutputStream(
                new FileOutputStream("object.txt")))
        {
            Person per = new Person("孙悟空", 500);
            // 将per对象写入输出流
            oos.writeObject(per);
        }
        catch (IOException ex)
        {
            ex.printStackTrace();
        }
    }
}
```

如果希望从二进制流中恢复Java对象，则需要使用反序列化。反序列化的步骤如下：

\(1\)创建一个ObjectInputStream输入流，这个输入流是一个处理流，所以必须建立在其他节点流的基础之上。

\(2\)调用ObjectInputStream对象的readObjet\(\)方法读取流中的对象，该方法返回一个Object类型的Java对象，如果程序知道该Java

对象的类型，则可以将该对象强制类型装换成其真实的类型。

```
public class ReadObject
{
    public static void main(String[] args)
    {
        try(
            // 创建一个ObjectInputStream输入流
            ObjectInputStream ois = new ObjectInputStream(
                new FileInputStream("object.txt")))
        {
            // 从输入流中读取一个Java对象，并将其强制类型转换为Person类
            Person p = (Person)ois.readObject();
            System.out.println("名字为：" + p.getName()
                + "\n年龄为：" + p.getAge());
        }
        catch (Exception ex)
        {
            ex.printStackTrace();
        }
    }
}
```

8.3对象引用的序列化

```
public class Teacher
    implements java.io.Serializable
{
    private String name;
    private Person student;
    public Teacher(String name , Person student)
    {
        this.name = name;
        this.student = student;
    }
    // 此处省略了name和student的setter和getter方法

    // name的setter和getter方法
    public void setName(String name)
    {
        this.name = name;
    }
    public String getName()
    {
        return this.name;
    }

    // student的setter和getter方法
    public void setStudent(Person student)
    {
        this.student = student;
    }
    public Person getStudent()
    {
        return this.student;
    }
}


public class WriteTeacher
{
    public static void main(String[] args)
    {
        try(
            // 创建一个ObjectOutputStream输出流
            ObjectOutputStream oos = new ObjectOutputStream(
                new FileOutputStream("teacher.txt")))
        {
            Person per = new Person("孙悟空", 500);
            Teacher t1 = new Teacher("唐僧" , per);
            Teacher t2 = new Teacher("菩提祖师" , per);
            // 依次将四个对象写入输出流
            oos.writeObject(t1);
            oos.writeObject(t2);
            oos.writeObject(per);
            oos.writeObject(t2);
        }
        catch (IOException ex)
        {
            ex.printStackTrace();
        }
    }
}


public class ReadTeacher
{
    public static void main(String[] args)
    {
        try(
            // 创建一个ObjectInputStream输出流
            ObjectInputStream ois = new ObjectInputStream(
                new FileInputStream("teacher.txt")))
        {
            // 依次读取ObjectInputStream输入流中的四个对象
            Teacher t1 = (Teacher)ois.readObject();
            Teacher t2 = (Teacher)ois.readObject();
            Person p = (Person)ois.readObject();
            Teacher t3 = (Teacher)ois.readObject();
            // 输出true
            System.out.println("t1的student引用和p是否相同："
                + (t1.getStudent() == p));
            // 输出true
            System.out.println("t2的student引用和p是否相同："
                + (t2.getStudent() == p));
            // 输出true
            System.out.println("t2和t3是否是同一个对象："
                + (t2 == t3));
        }
        catch (Exception ex)
        {
            ex.printStackTrace();
        }
    }
}
```

8.4 自定义序列化

```
public class Person
    implements java.io.Serializable
{
    private String name;
    private transient int age;
    // 注意此处没有提供无参数的构造器!
    public Person(String name , int age)
    {
        System.out.println("有参数的构造器");
        this.name = name;
        this.age = age;
    }
    // 省略name与age的setter和getter方法

    // name的setter和getter方法
    public void setName(String name)
    {
        this.name = name;
    }
    public String getName()
    {
        return this.name;
    }

    // age的setter和getter方法
    public void setAge(int age)
    {
        this.age = age;
    }
    public int getAge()
    {
        return this.age;
    }
}



public class TransientTest
{
    public static void main(String[] args)
    {
        try(
            // 创建一个ObjectOutputStream输出流
            ObjectOutputStream oos = new ObjectOutputStream(
                new FileOutputStream("transient.txt"));
            // 创建一个ObjectInputStream输入流
            ObjectInputStream ois = new ObjectInputStream(
                new FileInputStream("transient.txt")))
        {
            Person per = new Person("孙悟空", 500);
            // 系统会per对象转换字节序列并输出
            oos.writeObject(per);
            Person p = (Person)ois.readObject();
            System.out.println(p.getAge());
        }
        catch (Exception ex)
        {
            ex.printStackTrace();
        }
    }
}
```



