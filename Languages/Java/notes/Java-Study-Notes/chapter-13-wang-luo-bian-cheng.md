
---

一、Java 基本网络支持

Java为网络支持提供了java.net 包，该包下的 URL 和 URLComiection 等类提供了以编程方式访问Web 服务的功能，而 URLDecoder 和 URLEncoder 则提供了普通字符串和 application/x-www-form-urlencoded MIME 字符串相互转换的静态方法.

1.1 使用 InetAddress

Java 提供了InetAddress 类来代表IP地址,InetAddress下还有两个子类：Inet4Address 、 Inet6Address ，它们分别代表Intemet Protocol version 4 \(IPv4 \) 地址和 Internet Protocol version 6 \(IPv6\) 地址 .

InetAddress 类没有提供构造器，而是提供了如下两个静态方法来获取InetAddress 实例。

\(1\)getByName\(String host\): 根据主机获取对应的InetAddress 对象.

\(2\)getByAddress\(byte\[\] addr\): 根据原始IP地址来获取对应的 InetAddress 对象。

InetAddress 还提供了如下三个方法来获取InetAddress 实例对应的 IP 地址和主机名。

\(1\)StringgetCanonicalHostName\(\): 获取此IP地址的全限定域名 •

\(2\)String getHostAddress\(\): 返回该InetAddress 实例对应的IP地址字符串（以字符串形式）。

\(3\)StringgetHostNameO: 获取此IP地址的主机名。

除此之外，InetAddress 类还提供了一个 getLocalHost\(\) 方法来获取本机IP地址对应的InetAddress实例。

InetAddress 类还提供了一个 isReachableO 方法，用于 测试是否可以到达该地址。该方法将尽最大努力试图到达主机，似防火墙和服务器配置可能阻塞请求，使得它在访问某些特定的端口时处于不可达状态。如果可以获得权限，典型的实现将使用 ICMP ECHO REQUEST• ，否则它将试图在 目标主机的端口 7\(Echo\) 上建立 TCP 连接。下面程序测试了 InetAddress 类的简单用法。

```
public class InetAddressTest
{
    public static void main(String[] args)
        throws Exception
    {
        // 根据主机名来获取对应的InetAddress实例
        InetAddress ip = InetAddress.getByName("www.crazyit.org");
        // 判断是否可达
        System.out.println("crazyit是否可达：" + ip.isReachable(2000));
        // 获取该InetAddress实例的IP字符串
        System.out.println(ip.getHostAddress());
        // 根据原始IP地址来获取对应的InetAddress实例
        InetAddress local = InetAddress.getByAddress(
            new byte[]{127,0,0,1});
        System.out.println("本机是否可达：" + local.isReachable(5000));
        // 获取该InetAddress实例对应的全限定域名
        System.out.println(local.getCanonicalHostName());
    }
}
```

1.2、URLDecoder和ULREncoder

```
public class URLDecoderTest
{
    public static void main(String[] args)
        throws Exception
    {
        // 将application/x-www-form-urlencoded字符串
        // 转换成普通字符串
        // 其中的字符串直接从图17.3所示窗口复制过来
        String keyWord = URLDecoder.decode(
            "%E7%96%AF%E7%8B%82java", "utf-8");
        System.out.println(keyWord);
        // 将普通字符串转换成
        // application/x-www-form-urlencoded字符串
        String urlStr = URLEncoder.encode(
            "疯狂Android讲义" , "GBK");
        System.out.println(urlStr);
    }
}
```

1.3、URL 、URLconnection和URLPermission

URL \(Uniform Resource Locator\)对象代表统一资源定位器，它坫指向互联网 “资源”的指针。资源可以是简单的文件或目录，也可以是对更为复杂对象的引用，例如对数据库或搜索引擎的查询。在通常情况下，URL 可以由协议名、主机、端口和资源组成，即满足如下格式：

`protocol://host:port/resourceName`

例如如下的URL地址：

`http://www.crazyit.org/index.php`

URL 类提供了多个构造器用于创建URL对象，一旦获得了URL 对象之后，就可以调用如下方法来访问该 URL 对应的资源。

\(1\)  StringgetFile\(\): 获 取 该 URL 的资源名。

\(2\)  StringgetHost \( \) : 获 取 该 URL 的主机名。

\(3\)  StringgetPathO : 获 取 该 URL 的路径部分。

\(4\)  intgetPortO : 获取 该  URL 的端口 。

\(5\) StringgetProtocol\(\):获取该 URL 的协议名称。

\(6\)  String getQuery\(\): 获取 该 URL 的查询字符串部分。

\(7\)  URLConnection openConnection \(\):返回一个  URLConnection 对象，它代表了  L j URL 所引用的远程对象的连接。

\(8\)IutStream openStream \( \) : 打 开 与 此 URL 的连接，并返回一个用子读取该 URL 资源的InputStream.

如下程序实现了一个多线程下载工具类：

```
ublic class DownUtil
{
    // 定义下载资源的路径
    private String path;
    // 指定所下载的文件的保存位置
    private String targetFile;
    // 定义需要使用多少线程下载资源
    private int threadNum;
    // 定义下载的线程对象
    private DownThread[] threads;
    // 定义下载的文件的总大小
    private int fileSize;

    public DownUtil(String path, String targetFile, int threadNum)
    {
        this.path = path;
        this.threadNum = threadNum;
        // 初始化threads数组
        threads = new DownThread[threadNum];
        this.targetFile = targetFile;
    }

    public void download() throws Exception
    {
        URL url = new URL(path);
        HttpURLConnection conn = (HttpURLConnection) url.openConnection();
        conn.setConnectTimeout(5 * 1000);
        conn.setRequestMethod("GET");
        conn.setRequestProperty(
            "Accept",
            "image/gif, image/jpeg, image/pjpeg, image/pjpeg, "
            + "application/x-shockwave-flash, application/xaml+xml, "
            + "application/vnd.ms-xpsdocument, application/x-ms-xbap, "
            + "application/x-ms-application, application/vnd.ms-excel, "
            + "application/vnd.ms-powerpoint, application/msword, */*");
        conn.setRequestProperty("Accept-Language", "zh-CN");
        conn.setRequestProperty("Charset", "UTF-8");
        conn.setRequestProperty("Connection", "Keep-Alive");
        // 得到文件大小
        fileSize = conn.getContentLength();
        conn.disconnect();
        int currentPartSize = fileSize / threadNum + 1;
        RandomAccessFile file = new RandomAccessFile(targetFile, "rw");
        // 设置本地文件的大小
        file.setLength(fileSize);
        file.close();
        for (int i = 0; i < threadNum; i++)
        {
            // 计算每条线程的下载的开始位置
            int startPos = i * currentPartSize;
            // 每个线程使用一个RandomAccessFile进行下载
            RandomAccessFile currentPart = new RandomAccessFile(targetFile,
                "rw");
            // 定位该线程的下载位置
            currentPart.seek(startPos);
            // 创建下载线程
            threads[i] = new DownThread(startPos, currentPartSize,
                currentPart);
            // 启动下载线程
            threads[i].start();
        }
    }

    // 获取下载的完成百分比
    public double getCompleteRate()
    {
        // 统计多条线程已经下载的总大小
        int sumSize = 0;
        for (int i = 0; i < threadNum; i++)
        {
            sumSize += threads[i].length;
        }
        // 返回已经完成的百分比
        return sumSize * 1.0 / fileSize;
    }

    private class DownThread extends Thread
    {
        // 当前线程的下载位置
        private int startPos;
        // 定义当前线程负责下载的文件大小
        private int currentPartSize;
        // 当前线程需要下载的文件块
        private RandomAccessFile currentPart;
        // 定义已经该线程已下载的字节数
        public int length;

        public DownThread(int startPos, int currentPartSize,
            RandomAccessFile currentPart)
        {
            this.startPos = startPos;
            this.currentPartSize = currentPartSize;
            this.currentPart = currentPart;
        }

        @Override
        public void run()
        {
            try
            {
                URL url = new URL(path);
                HttpURLConnection conn = (HttpURLConnection)url
                    .openConnection();
                conn.setConnectTimeout(5 * 1000);
                conn.setRequestMethod("GET");
                conn.setRequestProperty(
                    "Accept",
                    "image/gif, image/jpeg, image/pjpeg, image/pjpeg, "
                    + "application/x-shockwave-flash, application/xaml+xml, "
                    + "application/vnd.ms-xpsdocument, application/x-ms-xbap, "
                    + "application/x-ms-application, application/vnd.ms-excel, "
                    + "application/vnd.ms-powerpoint, application/msword, */*");
                conn.setRequestProperty("Accept-Language", "zh-CN");
                conn.setRequestProperty("Charset", "UTF-8");
                InputStream inStream = conn.getInputStream();
                // 跳过startPos个字节，表明该线程只下载自己负责哪部分文件。
                inStream.skip(this.startPos);
                byte[] buffer = new byte[1024];
                int hasRead = 0;
                // 读取网络数据，并写入本地文件
                while (length < currentPartSize
                    && (hasRead = inStream.read(buffer)) != -1)
                {
                    currentPart.write(buffer, 0, hasRead);
                    // 累计该线程下载的总大小
                    length += hasRead;
                }
                currentPart.close();
                inStream.close();
            }
            catch (Exception e)
            {
                e.printStackTrace();
            }
        }
    }
}
```

程序中 DownUtils 类中的 download \(\)方法负责按如下步骤来实现多线程下栽。

\(1\)创建 URL 对象。

\(2\)获 取 指 定 URL 对象所指向资源的大小（通过 getContemLengthO 方法获得），此处用到了URLConnection 类，该类代表 Java 应用程序和 URL 之间的通信链接。

\(3\)在本地磁盘上创建—个与 网络资源具有相同大小的空文件。

\(4\)计算每个线程应该下载网络资源的哪个部分（从哪个字节开始，到哪个字节结束\)。

\(5\)依次创建、启动多个线程来下载网络资源的指定部分。

```
public class MultiThreadDown
{
    public static void main(String[] args) throws Exception
    {
        // 初始化DownUtil对象
        final DownUtil downUtil = new DownUtil("http://www.crazyit.org/"
            + "attachments/month_1403/1403202355ff6cc9a4fbf6f14a.png"
            , "ios.png", 4);
        // 开始下载
        downUtil.download();
        new Thread(() -> {
                while(downUtil.getCompleteRate() < 1)
                {
                    // 每隔0.1秒查询一次任务的完成进度，
                    // GUI程序中可根据该进度来绘制进度条
                    System.out.println("已完成："
                        + downUtil.getCompleteRate());
                    try
                    {
                        Thread.sleep(1000);
                    }
                    catch (Exception ex){}
                }
        }).start();
    }
}
```

1.4URLPermission

Java8新增了一个 URLPermission  工具类，用于管理 HttpURLConnection  的权限问题.如果在

HttpURLConnection 安装了安全管理器，通过该对象打开连接时就黹要先获得权限。

通常创建一个和 URL 的连接，并发送请求、读取此 URL 引用的资源浠要如下几个步驟 .

\(1\)通过调用URL对象的 openConncction\(\)方法来创建 URLConncction 对象 。

\(2\)设置URLConnection 的参数和普通请求属性。

\(3\)如果只是发送 GET 方式请求，则使用 connect\(\) 方法建立和远程资源之间的实际连接即可：如

果需要发送 POST 方式的请求，则需要获取 URLConnection 实例对应的输出流来发送请求参数.

\(4\)远程资源变为可用,程序可以访问远程资源的头字段或通过输入流读取远程资源的数据 .

向Web站点发送GET请求、POST强求，并从Web站点取得响应代码如下：

```
ublic class GetPostTest
{
    /**
     * 向指定URL发送GET方法的请求
     * @param url 发送请求的URL
     * @param param 请求参数，格式满足name1=value1&name2=value2的形式。
     * @return URL所代表远程资源的响应
     */
    public static String sendGet(String url , String param)
    {
        String result = "";
        String urlName = url + "?" + param;
        try
        {
            URL realUrl = new URL(urlName);
            // 打开和URL之间的连接
            URLConnection conn = realUrl.openConnection();
            // 设置通用的请求属性
            conn.setRequestProperty("accept", "*/*");
            conn.setRequestProperty("connection", "Keep-Alive");
            conn.setRequestProperty("user-agent"
                , "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)");
            // 建立实际的连接
            conn.connect();
            // 获取所有响应头字段
            Map<String, List<String>> map = conn.getHeaderFields();
            // 遍历所有的响应头字段
            for (String key : map.keySet())
            {
                System.out.println(key + "--->" + map.get(key));
            }
            try(
                // 定义BufferedReader输入流来读取URL的响应
                BufferedReader in = new BufferedReader(
                    new InputStreamReader(conn.getInputStream() , "utf-8")))
            {
                String line;
                while ((line = in.readLine())!= null)
                {
                    result += "\n" + line;
                }
            }
        }
        catch(Exception e)
        {
            System.out.println("发送GET请求出现异常！" + e);
            e.printStackTrace();
        }
        return result;
    }
    /**
     * 向指定URL发送POST方法的请求
     * @param url 发送请求的URL
     * @param param 请求参数，格式应该满足name1=value1&name2=value2的形式。
     * @return URL所代表远程资源的响应
     */
    public static String sendPost(String url , String param)
    {
        String result = "";
        try
        {
            URL realUrl = new URL(url);
            // 打开和URL之间的连接
            URLConnection conn = realUrl.openConnection();
            // 设置通用的请求属性
            conn.setRequestProperty("accept", "*/*");
            conn.setRequestProperty("connection", "Keep-Alive");
            conn.setRequestProperty("user-agent",
            "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)");
            // 发送POST请求必须设置如下两行
            conn.setDoOutput(true);
            conn.setDoInput(true);
            try(
                // 获取URLConnection对象对应的输出流
                PrintWriter out = new PrintWriter(conn.getOutputStream()))
            {
                // 发送请求参数
                out.print(param);
                // flush输出流的缓冲
                out.flush();
            }
            try(
                // 定义BufferedReader输入流来读取URL的响应
                BufferedReader in = new BufferedReader(new InputStreamReader
                    (conn.getInputStream() , "utf-8")))
            {
                String line;
                while ((line = in.readLine())!= null)
                {
                    result += "\n" + line;
                }
            }
        }
        catch(Exception e)
        {
            System.out.println("发送POST请求出现异常！" + e);
            e.printStackTrace();
        }
        return result;
    }
    // 提供主方法，测试发送GET请求和POST请求
    public static void main(String args[])
    {
        // 发送GET请求
        String s = GetPostTest.sendGet("http://localhost:8888/abc/a.jsp"
            , null);
        System.out.println(s);
        // 发送POST请求
        String s1 = GetPostTest.sendPost("http://localhost:8888/abc/login.jsp"
            , "name=crazyit.org&pass=leegang");
        System.out.println(s1);
    }
}
```

二、基于TCP协议的网络编程

TCP/IP 通信协议是一种可靠的网络协议，它在通信的两端各建立一个 Socket ,从而在通信的两端之间形成网络虚拟链路。一旦建立了虚拟的网络链路，两端的程序就可以通过虚拟链路进行通信 。 Java对基于 TCP 协议的网络通信提供了良好的封装， Java 使用 Socket 对象来代表两端的通信端口，并通过Socket 产生IO流来进行网络通信

2.1、TCP 协议基础

IP协议是Intemet上使用的一个关键协议，它的全称是Internet Protocol,即Internet协议，通常简称IP协议。通过使用IP协议，从而使Internet成为一个允许连接不同类型的计算机和不同操作系统的网络。

要使两台计算机彼此能进行通信，必须使两台计算机使用同一种“语言”， IP 协议只保证计算机能发送和接收分组数据。 IP 协议负责将消息从一个主机传送到另一个主机，消息在传送的过程中被分割成一个个的小包。

尽管计算机通过安装 IP 软件，保证了计算机之间可以发送和接收数据，但 IP 协议还不能解决数据分组在传输过程中可能出现的问题。因此，若要解决可能出现的问题，连上Internet的计算机还需要安装TCP协议来提供可靠并且无差错的通信服务。

TCP协议被称作一种端对端协议。这是因为它对两台计算机之间的连接起了重要作用 一 当一台计算机需要与另一台远程计算机连接时，TCP协议会让它们建立一个连接 ： 用于发送和接收数据的虚拟链路。TCP协议负责收集这些信息包，并将其按适当的次序放好传送，接收端收到后再将其正确地还原。TCP协议保证了数据包在传送中准确无误。TCP协议使用重发机制——当一个通信实体发送一个消息给另一个通信实体后，需要收到另一个通信实体的确认信息，如果没有收到另一个通信实体的确认信息，

则会再次軍发刚才发送的信息。

通过这种重发机制 ， TCP协议向应用程序提供了可靠的通信连接，使它能够自动适应网上的各种变化。即使在Internet暂时出现堵塞的情况, TCP也能够保证通信的可靠性。

2.2、使用ServerSocket创建TCP服务器

Java中能接收其他通信实体连接请求的类是ServerSocket, ServerSocket对象用于监听来自客户端的Socket连接，如果没有连接，它将一直处于等待状态。ServerSocket包含一个监听来自客户端连接请求的方法。

\(1\)Socket accept\(\):如果接收到一个客户端Socket的连接请求，该方法将返回一个与客户端Socket对应的Socket ;否则该方法将一直处于等待状态，线程也被阻塞。

为了创建ServerSocket对象，ServerSocket类提供了如下几个构造器。

\(1\)ServerSocket\(int port\):用指定的端口 port来创建一个ServerSocket。该端口应该有一个有效的端口整数值，即0~65535。

\(2\)ServerSocket\(int port,int backlog\):增加一个用来改变连接队列长度的参数backlog。

\(3\)ServerSocket\(intport,int backlog,InetAddress localAddr\):在机器存在多个 IP 地址的情况下，允许通过localAddr参数来指定将ServerSocket绑定到指定的IP地址。当ServerSocket使用完毕后，应使用ServerSocket的close\(\)方法来关闭该ServerSockeU在通常情

况下，服务器不应该只接收一个客户端请求，而应该不断地接收来自客户端的所有请求，所以Java程序通常会通过循环不断地调用ServerSocket的acceptO方法。

2.3使用Socket通信

客户端通常可以使用 Socket 的构造器来连接到指定服务器， Socket 通常可以使用如下两个构造器。

\(1\)Socket \(InetAddress/String remoteAddress，int port \):创建连接到指定远程主机、远程端口的  Socket ，该构造器没有指定本地地址、本地端口，默认使用本地主机的默认 IP 地址，默认使用系统动态分配的端口。

\(2\)Socket \( InetAddress/String remoteAddress ,  int port ,  InetAddress localAddr ,  int localPort \):创建连接到指定远程主机、远程端口的 Socket ,并指定本地IP地址和本地端口，适用于本地主机有多个 IP地址的情形。当客户端、服务器端产生了对应的Socket之后，，程序无须再区分服务器端、客户端，而是通过各自的Socket进行通信。

Socket提供了如下两个方法来获取输入流和输出流。

\(1\)  InputStream getlnputStream\(\) : 返 回 该 Socket 对象对应的输入流，让程序通过该输入流从 Socket中収出数据。

\(2\) OutputStream getOutputStream\(\):返 M 该Socket对象对应的输出流，让程序通过该输出流向Socket中输出数据。

下面的服务器端程序非常简单，它仅仅建立ServerSocket监听，并使用Socket获取输出流输出。

```
public class Server
{
    public static void main(String[] args)
        throws IOException
    {
        // 创建一个ServerSocket，用于监听客户端Socket的连接请求
        ServerSocket ss = new ServerSocket(30000);
        // 采用循环不断接受来自客户端的请求
        while (true)
        {
            // 每当接受到客户端Socket的请求，服务器端也对应产生一个Socket
            Socket s = ss.accept();
            // 将Socket对应的输出流包装成PrintStream
            PrintStream ps = new PrintStream(s.getOutputStream());
            // 进行普通IO操作
            ps.println("您好，您收到了服务器的新年祝福！");
            // 关闭输出流，关闭Socket
            ps.close();
            s.close();
        }
    }
}
```

下面的服务器端程序非常简单，它仅仅使用Socket建立与指定IP地址、指定端口的连接，并使用Socket获取输入流读取数据：

```
public class Client
{
    public static void main(String[] args)
        throws IOException
    {
        Socket socket = new Socket("127.0.0.1" , 30000);   // ①
        // 将Socket对应的输入流包装成BufferedReader
        BufferedReader br = new BufferedReader(
        new InputStreamReader(socket.getInputStream()));
        // 进行普通IO操作
        String line = br.readLine();
        System.out.println("来自服务器的数据：" + line);
        // 关闭输入流、socket
        br.close();
        socket.close();
    }
}
```

2.4加入多线程

```
线程通信类
public class MyServer
{
    // 定义保存所有Socket的ArrayList，并将其包装为线程安全的
    public static List<Socket> socketList
        = Collections.synchronizedList(new ArrayList<>());
    public static void main(String[] args)
        throws IOException
    {
        ServerSocket ss = new ServerSocket(30000);
        while(true)
        {
            // 此行代码会阻塞，将一直等待别人的连接
            Socket s = ss.accept();
            socketList.add(s);
            // 每当客户端连接后启动一条ServerThread线程为该客户端服务
            new Thread(new ServerThread(s)).start();
        }
    }
}

// 负责处理每个线程通信的线程类
public class ServerThread implements Runnable
{
    // 定义当前线程所处理的Socket
    Socket s = null;
    // 该线程所处理的Socket所对应的输入流
    BufferedReader br = null;
    public ServerThread(Socket s)
    throws IOException
    {
        this.s = s;
        // 初始化该Socket对应的输入流
        br = new BufferedReader(new InputStreamReader(s.getInputStream()));
    }
    public void run()
    {
        try
        {
            String content = null;
            // 采用循环不断从Socket中读取客户端发送过来的数据
            while ((content = readFromClient()) != null)
            {
                // 遍历socketList中的每个Socket，
                // 将读到的内容向每个Socket发送一次
                for (Socket s : MyServer.socketList)
                {
                    PrintStream ps = new PrintStream(s.getOutputStream());
                    ps.println(content);
                }
            }
        }
        catch (IOException e)
        {
            e.printStackTrace();
        }
    }
    // 定义读取客户端数据的方法
    private String readFromClient()
    {
        try
        {
            return br.readLine();
        }
        // 如果捕捉到异常，表明该Socket对应的客户端已经关闭
        catch (IOException e)
        {
            // 删除该Socket。
            MyServer.socketList.remove(s);      // ①
        }
        return null;
    }
}
```

```
public class MyClient
{
    public static void main(String[] args)throws Exception
    {
        Socket s = new Socket("127.0.0.1" , 30000);
        // 客户端启动ClientThread线程不断读取来自服务器的数据
        new Thread(new ClientThread(s)).start();   // ①
        // 获取该Socket对应的输出流
        PrintStream ps = new PrintStream(s.getOutputStream());
        String line = null;
        // 不断读取键盘输入
        BufferedReader br = new BufferedReader(
            new InputStreamReader(System.in));
        while ((line = br.readLine()) != null)
        {
            // 将用户的键盘输入内容写入Socket对应的输出流
            ps.println(line);
        }
    }
}

public class ClientThread implements Runnable
{
    // 该线程负责处理的Socket
    private Socket s;
    // 该线程所处理的Socket所对应的输入流
    BufferedReader br = null;
    public ClientThread(Socket s)
        throws IOException
    {
        this.s = s;
        br = new BufferedReader(
            new InputStreamReader(s.getInputStream()));
    }
    public void run()
    {
        try
        {
            String content = null;
            // 不断读取Socket输入流中的内容，并将这些内容打印输出
            while ((content = br.readLine()) != null)
            {
                System.out.println(content);
            }
        }
        catch (Exception e)
        {
            e.printStackTrace();
        }
    }
}
```

2.5记录用户信息

三、基于UDP协议的网络编程

UDP 协议是一种不可靠的网络协议，它在通信实例的两端各建立一个 Socket ,但这两个 Socket 之间并没有虚拟链路，这两个 Socket 只是发送、接收数据报的对象。 Java 提供了  DatagramSocket 对象作为基于 UDP 协议的 Socket ，使用 DatagramPacket 代表 DatagramSocket 发送、接收的数据报。

3.1UDP 协议基础

UDP 协议是英文 User Datagram Protocol 的缩写，即用户数据报协议，主要用来支持那些需要在计算机之间传输数据的网络连接。 UDP 协议从问世至今已经被使用了很多年，虽然 UDP 协议目前应用不如 TCP 协议广泛，但 UDP 协议依然是一个非常实用和可行的网络传输层协议。尤其是在一些实时性很强的应用场说中，比如网络游戏、视频会议等， UDP 协议的快速更具有独特的魅力。

UDP 协议是一种面向非连接的协议，面向非连接指的是在正式通信前不必与对方先建立连接，不管对方状态就直接发送。罕于对方是否可以接收到这些数据内容， UDP 协议无法控制， 因此说 UDP 协议是一种不可靠的协议。 UDP 协议适用于一次只传送少量数据、对可靠性要求不高的应用环境。

与前面介绍的 TCP 协议一样， UDP 协议直接位于IP 协议之上。实际上，IP协议属于OSI 参考模型的网络层协议，而 UDP 协议和 TCP 协议都 属于传输 层 协议。因为 UDP 协议是面向非连接的协议，没有建立连接的过程，因此它的通信效率很高；但也 IHW 为

如此，它的可靠性不如 TCP 协议。UDP 协议的主要作用是完成网络数据流和数据报之间的转换——在信息的发送端， UDP 协议将网

络数据流封装成数据报，然后将数据报发送出去；在信息的接收端， UDP 协议将数据报转换成实际数据内容。

UDP协议和TCP协议简单对比如下。

\(1\)TCP协议：可靠，传输大小无限制，但是需要连接建立时间，差错控制开销大。

\(2\)UDP协议：不可靠，差错控制开销较小，传输大小限制在64KB以下，不需要建立连接。

3.2 使用 DatagramSocket 发送、接收数据

Java使用DatagramSocket代表UDP协议的Socket，DatagramSocket本身只是码头，不维护状态，不能产生IO流，它的唯一作用就是接收和发送数据报，Java使用DatagramPacket来代表数据报，DatagramSocket接收和发送的数据都是通过DatagramPacket对象完成的。

DatagramSocket 的构造器。

\(1\)DatagramSocket\(\):创建一个DatagramSocket实例，并将该对象绑足到本机默认丨P地址、本机所有可用端口中随机选择的某个端口。

\(2\)DatagramSocket\(improt\):创建一个DatagramSocket实例，并将该对象绑定到本机默认IP地址、指定端口。

\(3\)DatagramSocket\(int port， InetAddress laddr\):创建一个 DatagramSocket 实例，并将该对象绑定到指定IP地址、指定端口。

通过上面三个构造器中的任意一个构造器即可创建一个DatagramSocket实例，通常在创建服务器时，创建指定端口的DatagramSocket实例——这样保证其他客户端可以将数据发送到该服务器。一旦得到了 DatagramSocket实例之后，就可以通过如下两个方法来接收和发送数据。

\(1\)receive\(DatagramPacket p\) : 从 该 DatagramSocket 中接收数据报。

\(2\)send\(DatagramPacket p\) : 以 该 DatagramSocket 对象向外发送数据报。

使用DatagramSocket发送数据报时，DatagramSocket并不知道将该数据报发送到哪里，而是由DatagramPacket自身决定数据报的目的地。就像码头并不知道每个集装箱的目的地，码头只是将这些集装箱发送出去，而集装箱本身包含了该集装箱的目的地。

下面看一下DatagramPacket的构造器。

\(1\)DatagramPacket\(byte\[\]buf,int length\):以一个空数组来创建 DatagramPacket 对象，该对象的作用是接收DatagramSocket中的数据。

\(2\)DatagramPacket\(byte\[\] buf， int length, InelAddress addr， int port \):以一个包含数据的数组来创建DatagramPacket 对象，创建该 DatagramPacket 对象时还指定了IP地址和端口 ,这就决定了该数据报的目的地。

\(3\)DatagramPacket\(byte\[\] buf， int offset, int length \):以一个空数组来创建 DatagramPacket  对象，并指定接收到的数据放入 buf 数组中时从 offset 开始，最多放Length 个字节。

\(4\)DatagramPacket\(byte\[\] buf, int offset, int length, InetAddress address, int port\): 创建一个用于发送的DatagramPacket 对象，指定发送 buf 数组中从 offset 开始，总共 length 个字节。

当服务器端（也以是客户端）接收到一个 DatagramPacket 对象后，如果想向该数据报的发送者“反馈”一些信息，但由于 UDP 协议是面向非连接的，所以接收者并不知道每个数据报由谁发送过来，

程序可以调用 DatagramPacket 的如下三个方法来获取发送者的 IP 地址和端口。

\(1\) InetAddress getAddressO :当程序准备发送此数据报时，该方法返回此数据报的目标机器的 IP地址；当程序刚接收到一个数据报时，该方法返回该数据报的发送主机的 IP 地址。

\(2\) int getPort\(\): 当程序准备发送此数据报时，该方法返回此数据报的目标机器的端口：当程序刚接收到一个数据报时，该方法返回该数据报的发送主机的端口。

\(3\)  SocketAddress getSocketAddressO :当程序准备发送此数据报时，该方法返回此数据报的 目标SocketAddress ; 当 程 序 刚 接 收 到 一 个 数 据 报 时 ， 该 方 法 返 回 该 数 据 报 的 发 送 主 机 的SocketAddress 。

```
public class UdpServer
{
    public static final int PORT = 30000;
    // 定义每个数据报的最大大小为4K
    private static final int DATA_LEN = 4096;
    // 定义接收网络数据的字节数组
    byte[] inBuff = new byte[DATA_LEN];
    // 以指定字节数组创建准备接受数据的DatagramPacket对象
    private DatagramPacket inPacket =
        new DatagramPacket(inBuff , inBuff.length);
    // 定义一个用于发送的DatagramPacket对象
    private DatagramPacket outPacket;
    // 定义一个字符串数组，服务器发送该数组的的元素
    String[] books = new String[]
    {
        "疯狂Java讲义",
        "轻量级Java EE企业应用实战",
        "疯狂Android讲义",
        "疯狂Ajax讲义"
    };
    public void init()throws IOException
    {
        try(
            // 创建DatagramSocket对象
            DatagramSocket socket = new DatagramSocket(PORT))
        {
            // 采用循环接受数据
            for (int i = 0; i < 1000 ; i++ )
            {
                // 读取Socket中的数据，读到的数据放入inPacket封装的数组里。
                socket.receive(inPacket);
                // 判断inPacket.getData()和inBuff是否是同一个数组
                System.out.println(inBuff == inPacket.getData());
                // 将接收到的内容转成字符串后输出
                System.out.println(new String(inBuff
                    , 0 , inPacket.getLength()));
                // 从字符串数组中取出一个元素作为发送的数据
                byte[] sendData = books[i % 4].getBytes();
                // 以指定字节数组作为发送数据、以刚接受到的DatagramPacket的
                // 源SocketAddress作为目标SocketAddress创建DatagramPacket。
                outPacket = new DatagramPacket(sendData
                    , sendData.length , inPacket.getSocketAddress());
                // 发送数据
                socket.send(outPacket);
            }
        }
    }
    public static void main(String[] args)
        throws IOException
    {
        new UdpServer().init();
    }
}
```

```
public class UdpClient
{
    // 定义发送数据报的目的地
    public static final int DEST_PORT = 30000;
    public static final String DEST_IP = "127.0.0.1";
    // 定义每个数据报的最大大小为4K
    private static final int DATA_LEN = 4096;
    // 定义接收网络数据的字节数组
    byte[] inBuff = new byte[DATA_LEN];
    // 以指定字节数组创建准备接受数据的DatagramPacket对象
    private DatagramPacket inPacket =
        new DatagramPacket(inBuff , inBuff.length);
    // 定义一个用于发送的DatagramPacket对象
    private DatagramPacket outPacket = null;
    public void init()throws IOException
    {
        try(
            // 创建一个客户端DatagramSocket，使用随机端口
            DatagramSocket socket = new DatagramSocket())
        {
            // 初始化发送用的DatagramSocket，它包含一个长度为0的字节数组
            outPacket = new DatagramPacket(new byte[0] , 0
                , InetAddress.getByName(DEST_IP) , DEST_PORT);
            // 创建键盘输入流
            Scanner scan = new Scanner(System.in);
            // 不断读取键盘输入
            while(scan.hasNextLine())
            {
                // 将键盘输入的一行字符串转换字节数组
                byte[] buff = scan.nextLine().getBytes();
                // 设置发送用的DatagramPacket里的字节数据
                outPacket.setData(buff);
                // 发送数据报
                socket.send(outPacket);
                // 读取Socket中的数据，读到的数据放在inPacket所封装的字节数组里。
                socket.receive(inPacket);
                System.out.println(new String(inBuff , 0
                    , inPacket.getLength()));
            }
        }
    }
    public static void main(String[] args)
        throws IOException
    {
        new UdpClient().init();
    }
}
```

3.3 使用 MulticastSocket 实现多点广播

DatagramSocket 只允许数据报发送给指定的目标地址，而MulticastSocket 可以将数据报以广播方式发送到多个客户端。若要使用多点广播，则需要让一个数据报标有一组目标主机地址，当数据报发出后，整个组的所有主机都能收到该数据报。 IP 多点广播（或多点发送）实现了将单一信息发送到多个接收者的广播，其思想是设置一组特殊网络地址作为多点广播地址，每一个多点广播地址都被看做一个组，当客户端需要发送、接收广播信息时，加入到该组即可。

IP 协议为多点广播提供了这批特殊的IP地址，这些IP 地址的范围是224.0,0.0至239.255.255.255。

![](/assets/MulticastSocket.png)

MulticastSocket 类是实现多点广播的关键，当 MulticastSocket 把一个DatagramPacket 发送到多点广播 IP 地址时，该数据报将被自动广播到加入该地址的所有 MulticastSocket 。

MulticastSocket 既可以将数据报发送到多点广播地址，也可以接收其他主机的广播倍息。

MulticastSocket 有点像  DatagramSocket ，事实上  MulticastSocket 是  DatagramSocket 的一个子类，也就是说， MulticastSocket 是特殊的 DatagramSocket 。当要发送一个数据报时，可以使用随机端口创建MulticastSocket ，也可以在指定端口创建 MulticastSocket 。 MulticastSocket 提供了如下三个构造器。

\(1\)  public MulticastSocket \(\):使用本机默认地址、随机端口来创建 MulticastSocket 对象。

\(2\)  public MulticastSocket \( intportNumber \):使用本机默认地址、指定端口来创建  MulticastSocket对象。

\(3\) public MulticastSocket\(SocketAddress bindaddr \):使用本机指定IP地址、指定端口来创建MulticastSocket 对象。

创建 MulticastSocket 对象后，还需要将该 MuhicastSocket 加入到指定的多点广播地址，MulticastSocIcet 使用 joinGroup \(\)方法加入指定组：使用 leaveGroup（）方法脱离一个组。

\(1\)joinGroup \(InetAddress multicastAddr \) : 将 该  MulticastSocket 加入指定的多点广播地址。

\(2\)leaveGroup \(InetAddress multicastAddr \) : 让 该  MulticastSocket 离开指定的多点广播地址。

在某些系统中，可能有多个网络接口。这可能会给多点广播带来问题，这时候程序需要在一个指定的网络接口 上监听，通过调用 setlnterface \(\)方法可以强制 MulticastSocket 使用指定的网络接口；也可以使用getlnterface\(\)方法查询 MulticastSocket 的网络接口。

MulticastSocket 用于发送、接收数据报的方法是 DatagramSocket 完全 一 样。 MulticastSocket 比DatagramSocket 多了 一 个 setTimeToLive\(int ttl \)方法，该 ttl参数用于设置数据最多可以跨过多少个 网络 ，当ttl 的 值 为 0 时 ， 指 定 数 椐 报 应 停 留在 本 地 主 机 ； 当ttl 的 值 为 丨 时 ， 指 定 数 椐  报 发 送 到 本 地 局 域网; 当ttl 的 值 为 0 时， 意 味  只 能 发 送 到 本 站 点 的 网 络 上; 当 ttl 的 值 为 6 4 时 ， 意 味  数 据 报 应 保 留在本地区 ； 当ttl 的 值 为128时， 味数据报应保留在本大洲：当 ttl 的值为255时，意味着数据报可发送到所有地方; 默认情况下 ，该 ttl 的值为1。

```
// 让该类实现Runnable接口，该类的实例可作为线程的target
public class MulticastSocketTest implements Runnable
{
    // 使用常量作为本程序的多点广播IP地址
    private static final String BROADCAST_IP
        = "230.0.0.1";
    // 使用常量作为本程序的多点广播目的的端口
    public static final int BROADCAST_PORT = 30000;
    // 定义每个数据报的最大大小为4K
    private static final int DATA_LEN = 4096;
    //定义本程序的MulticastSocket实例
    private MulticastSocket socket = null;
    private InetAddress broadcastAddress = null;
    private Scanner scan = null;
    // 定义接收网络数据的字节数组
    byte[] inBuff = new byte[DATA_LEN];
    // 以指定字节数组创建准备接受数据的DatagramPacket对象
    private DatagramPacket inPacket
        = new DatagramPacket(inBuff , inBuff.length);
    // 定义一个用于发送的DatagramPacket对象
    private DatagramPacket outPacket = null;
    public void init()throws IOException
    {
        try(
            // 创建键盘输入流
            Scanner scan = new Scanner(System.in))
        {
            // 创建用于发送、接收数据的MulticastSocket对象
            // 由于该MulticastSocket对象需要接收数据，所以有指定端口
            socket = new MulticastSocket(BROADCAST_PORT);
            broadcastAddress = InetAddress.getByName(BROADCAST_IP);
            // 将该socket加入指定的多点广播地址
            socket.joinGroup(broadcastAddress);
            // 设置本MulticastSocket发送的数据报会被回送到自身
            socket.setLoopbackMode(false);
            // 初始化发送用的DatagramSocket，它包含一个长度为0的字节数组
            outPacket = new DatagramPacket(new byte[0]
                , 0 , broadcastAddress , BROADCAST_PORT);
            // 启动以本实例的run()方法作为线程体的线程
            new Thread(this).start();
            // 不断读取键盘输入
            while(scan.hasNextLine())
            {
                // 将键盘输入的一行字符串转换字节数组
                byte[] buff = scan.nextLine().getBytes();
                // 设置发送用的DatagramPacket里的字节数据
                outPacket.setData(buff);
                // 发送数据报
                socket.send(outPacket);
            }
        }
        finally
        {
            socket.close();
        }
    }
    public void run()
    {
        try
        {
            while(true)
            {
                // 读取Socket中的数据，读到的数据放在inPacket所封装的字节数组里。
                socket.receive(inPacket);
                // 打印输出从socket中读取的内容
                System.out.println("聊天信息：" + new String(inBuff
                    , 0 , inPacket.getLength()));
            }
        }
        // 捕捉异常
        catch (IOException ex)
        {
            ex.printStackTrace();
            try
            {
                if (socket != null)
                {
                    // 让该Socket离开该多点IP广播地址
                    socket.leaveGroup(broadcastAddress);
                    // 关闭该Socket对象
                    socket.close();
                }
                System.exit(1);
            }
            catch (IOException e)
            {
                e.printStackTrace();
            }
        }
    }
    public static void main(String[] args)
        throws IOException
    {
        new MulticastSocketTest().init();
    }
}
```

四、使用代理服务器

从Java5开始，Java在java.net包下提供了 Proxy和ProxySelector两个类，其中Proxy代表一个代理服务器，可以在打开URLConnection连接时指定Proxy,创建Socket连接时也可以指定Proxy:而ProxySelector代表一个代理选择器，它提供了对代理服务器更加灵活的控制，它可以对HTTP、HTTPS、FTP、SOCKS等进行分别设置，而且还可以设置不需要通过代理服务器的主机和地址。通过使用ProxySelector，可以实现像在Intemet Explorer、Firefox等软件中设置代理服务器类似的效果。

4.1直接使用Proxy创建连接

Proxy 存一个构造器： Proxy \( Proxy.Type type ,  SocketAddress sa \)，用于创建表示代理服务器的 Proxy

对象。其中 sa 参数指定代理服务器的地址，type 表示该代理服务器的类型，该服务器类型有如下三种：

\(1\) Proxy.Type.DIRECH 表示直接连接，不使用代理。

\(2\) Proxy.Type.HTTP :表示支持高级协议代理，如 HTTP 或 FTP 。

\(3\) Proxy.Type.SOCKS :表示 SOCKS\(V4或V5\)代理。

一旦创建了Proxy对象之后，程序就可以在使用 URLConnection 打开连接时，或者创建 Socket 连

接时传入一个 Proxy 对象，作为本次连接所使用的代理服务器。

其中 URL 包含了一个 URLConnection openConnection\(Proxy proxy \)方法，该方法使用指定的代理服

务器来打开连接：而 Socket 则提供了一个 Socket \( Proxy proxy \)构造器，该构造器使用指定的代理服务器创建一个没有连接的 Socket 对象。

下面以 URLConnection 为例来介绍如何在 URLConnection 中使用代理服务器。

```
public class ProxyTest
{
    // 下面是代理服务器的地址和端口，
    // 换成实际有效的代理服务器的地址和端口
    final String PROXY_ADDR = "129.82.12.188";
    final int PROXY_PORT = 3124;
    // 定义需要访问的网站地址
    String urlStr = "http://www.crazyit.org";
    public void init()
        throws IOException , MalformedURLException
    {
        URL url = new URL(urlStr);
        // 创建一个代理服务器对象
        Proxy proxy = new Proxy(Proxy.Type.HTTP
            , new InetSocketAddress(PROXY_ADDR , PROXY_PORT));
        // 使用指定的代理服务器打开连接
        URLConnection conn = url.openConnection(proxy);
        // 设置超时时长。
        conn.setConnectTimeout(5000);
        try(
            // 通过代理服务器读取数据的Scanner
            Scanner scan = new Scanner(conn.getInputStream(), "utf-8");
            PrintStream ps = new PrintStream("index.htm"))
        {
            while (scan.hasNextLine())
            {
                String line = scan.nextLine();
                // 在控制台输出网页资源内容
                System.out.println(line);
                // 将网页资源内容输出到指定输出流
                ps.println(line);
            }
        }
    }
    public static void main(String[] args)
        throws IOException , MalformedURLException
    {
        new ProxyTest().init();
    }
}
```

4.2使用ProxySelector自动选择代理服务器

lMll 介绍的1\*1接使用 Proxy 对象吋以在打开 URLConnection 或 Socket 时指定代理服务器，怛使用

ProxySelector代表一个代理选择器，它本身是一个抽象类，程序无法创建它的实例，幵发者可以考

虑继承 ProxySelector来实现自己的代理选择器。实现 ProxySelector的步骤非常简单，程序只要定义一

个继承 ProxySelector 的类，并让该类实现如下两个抽象方法。

\(1\) List&lt;Proxy &gt; select \(URI uri\): 根据业务需要返回代理服务器列表，如果该方法返回的集介中只包

含一个 Proxy ,该 Proxy 将会作为默认的代理服务器。

\(2\) connectFailed\(URI uri ,  SocketAddress sa ,IOException ioe \):连接代理服务器失敗时回调该方法。

```
public class ProxySelectorTest
{
    // 下面是代理服务器的地址和端口，
    // 随便一个代理服务器的地址和端口
    final String PROXY_ADDR = "139.82.12.188";
    final int PROXY_PORT = 3124;
    // 定义需要访问的网站地址
    String urlStr = "http://www.crazyit.org";
    public void init()
        throws IOException , MalformedURLException
    {
        // 注册默认的代理选择器
        ProxySelector.setDefault(new ProxySelector()
        {
            @Override
            public void connectFailed(URI uri
                , SocketAddress sa, IOException ioe)
            {
                System.out.println("无法连接到指定代理服务器！");
            }
            // 根据"业务需要"返回特定的对应的代理服务器
            @Override
            public List<Proxy> select(URI uri)
            {
                // 本程序总是返回某个固定的代理服务器。
                List<Proxy> result = new ArrayList<>();
                result.add(new Proxy(Proxy.Type.HTTP
                    , new InetSocketAddress(PROXY_ADDR , PROXY_PORT)));
                return result;
            }
        });
        URL url = new URL(urlStr);
        // 没有指定代理服务器、直接打开连接
        URLConnection conn = url.openConnection();   //①
        // 设置超时时长。
        conn.setConnectTimeout(3000);
        try(
            // 通过代理服务器读取数据的Scanner
            Scanner scan = new Scanner(conn.getInputStream());
            PrintStream ps = new PrintStream("index.htm"))
        {
            while (scan.hasNextLine())
            {
                String line = scan.nextLine();
                // 在控制台输出网页资源内容
                System.out.println(line);
                // 将网页资源内容输出到指定输出流
                ps.println(line);
            }
        }
    }
    public static void main(String[] args)
        throws IOException , MalformedURLException
    {
        new ProxySelectorTest().init();
    }
}
```

