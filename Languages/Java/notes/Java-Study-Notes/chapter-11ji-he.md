一、线程概述

几乎所有的操作系统都支持同时运行多个任务，一个任务通常就是一个程序，每个运行中的程序就是一个进程。当一个程序运行时，内部可能包含了多个顺序执行流，每个顺序执行流就是一个线程。

1.1、线程和进程

几乎所有的操作系统都支持进程的概念，所有运行中的任务通常对应一个进程（ Process )，当一个程序进入内存运行时，即变成一个进程。进程是处于运行过程中的程序，并且具有一定的独立功能，进程是系统进行资源分配和调度的一个独立单位。

一般而言，进程包含如下三个特征。

(1)独立性：进程是系统中独立存在的实体，它可以拥有自己独立的资源，每一个进程都拥有自己私有的地址空间。在没有经过进程本身允许的情况下，一个用户进程不可以直接访问其他进程的地址空间。

(2)动态性：进程与程序的区别在于，程序只是一个静态的指令集合，而进程是一个正在系统中活动的指令集合。在进程中加入了时间的概念。进程具有自己的生命周期和各种不同的状态，这些概念在程序中都是不具备的。

(3)并发性：多个进程可以在单个处理器上并发执行，多个进程之间不会互相影响。

1.2、线程于进程的关系

多线程则扩展了多进程的概念，使得同一个进程可以同时并发处理多个任务。线程（Thread)也被称作轻量级进程（Lightweight Process),线程是进程的执行单元。就像进程在操作系统中的地位一样，线程在程序中是独立的、并发的执行流。当进程被初始化后，主线程就被创建了。对于绝大多数的应用程序来说，通常仅要求有一个主线程，但也以在该进程内创建多条顺序执行流，这些顺序执行流就是线程，每个线程也是互相独立的。

线程是进程的组成部分，一个进程可以拥有多个线程，一个线程必须有一个父进程。线程可以拥有自己的堆栈、自己的程序计数器和自己的局部变屋，但小拥有系统资源，它与父进程的其他线程共享该进程所拥有的全部资源。因为多个线程共亨父进程里的全部资源，因此编程更加方便：怛必须更加小心，因为需要确保线程不会妨碍同一迸程里的其他线程。

线程时以完成一定的任务，可以与其他线程共享父进程中的共亨变量及部分环境，相互之间协间来完成进程所要完成的任务。

线程是独立运行的，它并不知道进程中是否还有其他线程存在。线程的执行是抢占式的，也就是说，当前运行的线程在任何时候都可能被挂起，以便另外一个线程可以运行。

一个线程可以创逑和撤销另一个线程，同一个进程中的多个线程之间可以并发执行。

从逻辑角度来看，多线程存在于一个应用程序中，让一个应用程序中可以有多个执行部分同时执行，但操作系统无须将多个线程看作多个独立的应用，对多线程实现调度和管理以及资源分配。线程的调度和管理由进程本身负责完成。

简而言之，一个程序运行后至少有一个进程，一个进程里可以包含多个线程，但至少要包含一个线程。

1.3、并发性和并行性的区别

并发性 （ concurrency ) 和并行性 （ parallel ) 是两个概念，并行指在同一时刻，有多条指令在多个处理器上同时执行；并发指在同一时刻只能有一条指令执行，但多个进程指令被快速轮換执行，使得在宏观上具有多个进程同时执行的效果.

1.4、多线程的优势

线程在程序中是独立的、并发的执行流，与分隔的进程相比，进程中线程之间的隔离程度要小。它们共享内存、文件句柄和其他每个进程应有的状态。

因为线程的划分尺度小于进程，使得多线程程序的并发性高。进程在执行过程中拥有独立的内存单元，而多个线程共享内存，从而极大地提高了程序的运行效率。

线程比进程具有更高的性能，这是由于同一个进程中的线程都有共性一多个线程共享同一个进程虚拟空间。线程共享的环境包括：进程代码段、进程的公有数据等。利用这些共李的数据，线程很容易实现相互之间的通信。

当操作系统创建一个进程时，必须为该进程分配独立的内存空间，并分配大量的相关资源；但创建一个线程则简中-得多，因此使用多线程来实现并发比使用多进程实现并发的性能要高得多。

总结起来，使用多线程编程具有如下几个优点。

(1)进程之间不能共享内存，但线程之间共享内存非常容易。

(2)系统创建进程时需要为该进程重新分配系统资源，但创建线程则代价小得多，因此使用多线程来实现多任务并发比多进程的效率高。

(3)Java 语言内置了多线程功能支持，而不是单纯地作为底层操作系统的调度方式，从而简化了Java的多线程编程。

二、线程创建

2.1、继承Thread类创建线程类

继承Thread类创建并启动线程的步骤：

(1)定义Thread类的子类，并重写Thread类的run()方法（线程执行体）。

(2)创建Thread子类的实例，即创建线程对象

(3)调用线程对象的start()方法来启动该线程。

继承Thread类创建线程并启动代码：

```
// 通过继承Thread类来创建线程类
public class FirstThread extends Thread
{
    private int i ;
    // 重写run方法，run方法的方法体就是线程执行体
    public void run()
    {
        for ( ; i < 100 ; i++ )
        {
            // 当线程类继承Thread类时，直接使用this即可获取当前线程
            // Thread对象的getName()返回当前该线程的名字
            // 因此可以直接调用getName()方法返回当前线程的名
            System.out.println(getName() +  " " + i);
        }
    }
    public static void main(String[] args)
    {
        for (int i = 0; i < 100;  i++)
        {
            // 调用Thread的currentThread方法获取当前线程
            System.out.println(Thread.currentThread().getName()
                +  " " + i);
            if (i == 20)
            {
                // 创建、并启动第一条线程
                new FirstThread().start();
                // 创建、并启动第二条线程
                new FirstThread().start();
            }
        }
    }
}
```

Thread.currentThread():返回当前正在执行的线程对象。

getName():返回调用该方法的线程名字。

2.2、实现Runnable接口创建线程类

实现Runnable接口创建线程并启动线程的步骤：

(1)定义Runnable接口实现类，并重写Runnable接口的run()方法.

(2)创建Runnable实现类的实例，并以此实例作为Thread的target来创建Thread对象，该Thread对象是真正的线程对象。

(3)调用线程对象的start()方法来启动线程。

```
// 通过实现Runnable接口来创建线程类
public class SecondThread implements Runnable
{
    private int i ;
    // run方法同样是线程执行体
    public void run()
    {
        for ( ; i < 100 ; i++ )
        {
            // 当线程类实现Runnable接口时，
            // 如果想获取当前线程，只能用Thread.currentThread()方法。
            System.out.println(Thread.currentThread().getName()
                + "  " + i);
        }
    }

    public static void main(String[] args)
    {
        for (int i = 0; i < 100;  i++)
        {
            System.out.println(Thread.currentThread().getName()
                + "  " + i);
            if (i == 20)
            {
                SecondThread st = new SecondThread();     // ①
                // 通过new Thread(target , name)方法创建新线程
                new Thread(st , "新线程1").start();
                new Thread(st , "新线程2").start();
            }
        }
    }
}
```

2.3、使用Callable和Future创建线程

Java5以后，Java提供Callable接口，该接口提供一个call()方法作为线程执行体，但call()方法比run()方法功能更强大：

(1)call()方法可以有返回值

(2)call()方法可以声明抛出异常。

Java5提供Future接口来代表Callable接口里的call()方法的返回值，

并为 Future 接口提供了一个FutureTask 实现类，该实现类实现了 Future 接口，并实现了 Runnable 接口——可以作为Thread类的 target。

在 Future 接口串里定义了如下几个公共方法来控制它关联的Callable任务。

(1)boolean cancel(boolean  maylntemiptlfRunning ):试图取消该Future里关联的Callable任务。

(2)V get() : 返回 Callable 任务里 call()方法的返回值。 调用该方法将导致程序阻塞，必须等到子线程结束后才会得到返回值。

(3)V get(long timeout ， TimeUnit unit): 返回Callable 任务里call()方法的返回值。该方法让程序最多阻塞timeout和unit指定的时间，如果经过指定时间后 Callable 任务依然没有返回值，将会抛出TimeoutException异常。

(4)boolean isCancelled(): 如果在 Callable 任务正常完成前被取消，则返回true 。

(5)boolean isDone() : 如果 Callable 任务已究成，则返回true 。

创建并启动有返回值的线程的步骤如下：

(1)创建Callable 接口的实现类，并实现call()法，该 call()方法将作为线程执行体，且该call()方法有返回值，再创建 Callable 实现类的实例。从Java8开始，可以直接使用 Lambda 表达式创逑 Callable对象

(2)使用 FutureTask 类来包装 Callable 对象，该 FutureTask 对象封装了该 Callable 对象的 call ()方法的返回值。

(3)使用FutureTask 对象作为Thread对象的target创建并启动新线程。

(4)调用FutureTask 对象的get()方法来获得子线程执行结束后的返回值

```
public class ThirdThread
{
    public static void main(String[] args)
    {
        // 创建Callable对象
        ThirdThread rt = new ThirdThread();
        // 先使用Lambda表达式创建Callable<Integer>对象
        // 使用FutureTask来包装Callable对象
        FutureTask<Integer> task = new FutureTask<Integer>((Callable<Integer>)() -> {
            int i = 0;
            for ( ; i < 100 ; i++ )
            {
                System.out.println(Thread.currentThread().getName()
                    + " 的循环变量i的值：" + i);
            }
            // call()方法可以有返回值
            return i;
        });
        for (int i = 0 ; i < 100 ; i++)
        {
            System.out.println(Thread.currentThread().getName()
                + " 的循环变量i的值：" + i);
            if (i == 20)
            {
                // 实质还是以Callable对象来创建、并启动线程
                new Thread(task , "有返回值的线程").start();
            }
        }
        try
        {
            // 获取线程返回值
            System.out.println("子线程的返回值：" + task.get());
        }
        catch (Exception ex)
        {
            ex.printStackTrace();
        }
    }
}
```

2.4、线程创建三种方法性能分析

通过继承Thread类或实现Runnable、Callable 接口都可以实现多线程，不过实现 Runnable 接口与实现Callable接口的方式基本相间，只是 Callable 接口里定义的方法有返回值，可以声明抛出异常而已。因此可以将实现 Runnable 接口和实现 Callable 接口归为一种方式。这种方式与继承 Thread 方式之间的主要差别如下。

采用实现 Runnable 、Callable 接口的方式创建多线程的优缺点：

(1)线程类只是实现了Runnable 接口或 Callable 接口，还可以继承其他类。

(2)在这种方式下，多个线程可以共享一个target对象，所以非常适合多个相同线程来处理同一份资源的情况，从而可以将CPU、代码和数据分开，形成淸晰的模型，较好地体现了面向对象的思想。

(3)劣势是，编程稍稍复杂，如果需要访问当前线程，则必须使用 Thread.curremThread()方法。

采用继承Thread类的方式创建多线程的优缺点：

(4)劣势是因为线程类已经继承了Thread类，所以不能再继承其他父类。

(5)优势是编写简单，如果需要访问当前线程，则无须使用 Thread.curremThread()方法，直接使用this 即可获得当前线程.

鉴于上面分析,因此一般推荐采用实现 Runnable 接口、Callable接口的方式来创建多线程。

三、线程生命周期

当线程被创建并启动以后，它既不是一启动就进入了执行状态，也不是一直处于执行状态，在线程的生命周期中，它要经过新建（New)、 就绪（Runnable)、运行（Running )、 阻塞（Blocked) 和死亡（Dead)5种状态。尤其是当线程启动以后，它不可能一直“霸占”着 CPU 独自运行，所以CPU需要在多条线程之间切换，于是线程状态也会多次在运行、阻塞之间切换。

3.1新建和就绪状态

当程序使用 new 关键字创建了一个线程之后，该线程就处于新建状态，此时它和其他的Java对象一样，仅仅由Java虚拟机为其分配内存，

并初始化其成员变量的值。此时的线程对象没有表现出任何线程的动态特征，程序也不会执行线程的线程执行体。

当线程对象调用start()方法之后，该线程处于就绪状态，Java虚拟机会为其创建方法调用栈和程序计数器，处这个状态中的线程并没有开始运行，只是表示该线程就可以运行了。至于该线程何时幵始运行，取决于JVM线程调度器的调度。

```
public class InvokeRun extends Thread
{
    private int i ;
    // 重写run方法，run方法的方法体就是线程执行体
    public void run()
    {
        for ( ; i < 100 ; i++ )
        {
            // 直接调用run方法时，Thread的this.getName返回的是该对象名字，
            // 而不是当前线程的名字。
            // 使用Thread.currentThread().getName()总是获取当前线程名字
            System.out.println(Thread.currentThread().getName()
                +  " " + i);   // ①
        }
    }
    public static void main(String[] args)
    {
        for (int i = 0; i < 100;  i++)
        {
            // 调用Thread的currentThread方法获取当前线程
            System.out.println(Thread.currentThread().getName()
                +  " " + i);
            if (i == 20)
            {
                // 直接调用线程对象的run方法，
                // 系统会把线程对象当成普通对象，run方法当成普通方法，
                // 所以下面两行代码并不会启动两条线程，而是依次执行两个run方法
                new InvokeRun().run();
                new InvokeRun().run();
            }
        }
    }
}
```

如上代码，启动线程的正确方法是调用Thread对象的start()方法，而不是直接调用run()方法，如果直接调用run()方法就变成单线程程序。

同时如果调用了run()方法，该线程已经不再处于新建状态，不要再次调用线程对象的start()方法。

只能对于新建状态的线程调用start()方法，否则会引发IllegalThreadStateException.

调用线程对象的start()方法后，该线程立即进入就绪状态——就绪状态相当于“等待执行”，但该线程并未真正进入运行状态。

3.2运行和阻塞状态

如果处于就绪状态的线程获得了CPU,开始执行run()方法的线程执行体，则该线程处于运行状态，如果计算机只有一个 CPU ,那么在任何时刻只有一个线程处于运行状态。当然，在一个多处理器的机器上，将会有多个线程并行（注意是并行： parallel )执行；当线程数大于处理器数时，依然会存在多个线程在同一个 CPU 上轮换的现象。

当一个线程开始运行后，它不可能一直处于运行状态（除非它的线程执行体足够短，瞬间就执行结束了），线程在运行过程中需要被中断，目的是使其他线程获得执行的机会，线程调度的细节取决于底层平台所采用的策略。对于采用抢占式策略的系统而言，系统会给每个可执行的线程一个小时间段来处理任务；当该时间段用完后，系统就会剥夺该线程所占用的资源，让其他线程获得执行的机会。在选择下一个线程时，系统会考虑线程的优先级。所有现代的桌面和服务器操作系统都采用抢占式调度策略，但—些小型设备如手机则可能采用协作式调度策略，在这样的系统中，只有当一个线程调用了它的 sleep()或 yield()方法后才会放弃所占用的资源——也就是必须由该线程主动放弃所占用的资源。

当发生如下情况时，线程将会进入阻塞状态:

(1)线程调用sleep()方法主动放弃所占用的处理器资源。

(2)线程调用了一个阻塞式IO方法，在该方法返回之前，该线程被阻塞。

(3)线程试图获得一个同步监视器，但该同步监视器正被其他线程所持有

(4)线程在等待某个通知（notify ）。

(5)程序调用了线程的 suspend()方法将该线程挂起。但这个方法容易导致死锁，所以应该尽量避免使用该方法。

当前正在执行的线程被阻塞之后，其他线程就可以获得执行的机会。被阻塞的线程会在合适的时候重新进入就绪状态，注意是就绪状态而不是运行状态。也就是说，被阻塞线程的阻塞解除后，必须重新等待线程调度器再次调度它。

针对上面几种情况，当发生如下特定的情况时可以解除上面的阻塞，让该线程重新进入就绪状态。

(1)调用 sleep ()方法的线程经过了指定时间。

(2)线程调用的阻塞式IO方法已经返回。

(3)线程成功地获得了试图取得的同步监视器。

(4)线程正在等待某个通知时，其他线程发出了一个通知。

(5)处于挂起状态的线程被调用了resume()恢复方法。

![](/assets/线程状态.png)

3.3死亡状态

线程会以如下三种方式结束，结束后就处于死亡状态。

(1)run()或call()方法执行完成，线程正常结束。

(2)线程抛出一个未捕获的Exception会Error

(3)直接调用该线程的stop()方法（该方法易死锁，不推荐使用）来结束该线程。

处于死亡状态的线程，无法再次运行（即调用start方法），如果调用start方法会引发IllegalThreadStateException.

```
public class StartDead extends Thread
{
    private int i ;
    // 重写run方法，run方法的方法体就是线程执行体
    public void run()
    {
        for ( ; i < 100 ; i++ )
        {
            System.out.println(getName() +  " " + i);
        }
    }
    public static void main(String[] args)
    {
        // 创建线程对象
        StartDead sd = new StartDead();
        for (int i = 0; i < 300;  i++)
        {
            // 调用Thread的currentThread方法获取当前线程
            System.out.println(Thread.currentThread().getName()
                +  " " + i);
            if (i == 20)
            {
                // 启动线程
                sd.start();
                // 判断启动后线程的isAlive()值，输出true
                System.out.println(sd.isAlive());
            }
            // 只有当线程处于新建、死亡两种状态时isAlive()方法返回false。
            // 当i > 20，则该线程肯定已经启动过了，如果sd.isAlive()为假时，
            // 那只能是死亡状态了。
            if (i > 20 && !sd.isAlive())

            {
                // 试图再次启动该线程
                sd.start();
            }
        }
    }
}
```

四、线程操作——线程控制

4.1join线程

join():Thread提供的一个让线程等待另一个线程完成的方法。当某个程序执行流中调用其它线程的join()方法时。调用线程将被阻塞，直到被join()方法加入的join线程执行完为止。

```
public class JoinThread extends Thread{
    // 提供一个有参数的构造器，用于设置该线程的名字
    public JoinThread(String name){
        super(name);
    }
    // 重写run()方法，定义线程执行体
    public void run()
    {
        for (int i = 0; i < 100 ; i++ )
        {
            System.out.println(getName() + "  " + i);
        }
    }
    public static void main(String[] args)throws Exception
    {
        // 启动子线程
        new JoinThread("新线程").start();
        for (int i = 0; i < 100 ; i++ )
        {
            if (i == 20)
            {
                JoinThread jt = new JoinThread("被Join的线程");
                jt.start();
                // main线程调用了jt线程的join()方法，main线程
                // 必须等jt执行结束才会向下执行
                jt.join();
            }
            System.out.println(Thread.currentThread().getName()
                + "  " + i);
        }
    }
}
```

join()方法有如下三种重载方法：

(1)join():等待被join的线程执行完成

(2)join(long millis):等待被join的线程的时间最长为millis毫秒，如果在millis毫秒内被join的线程还没有执行结束，则不再等待。

(3)join(long millis,int nanos):等待被join的线程的时间最长为millis毫秒加nanos毫微秒。(不推荐，时间无法精确到毫微秒)

4.2后台线程

在后台运行且为其它的线程提供服务的线程称为后台线程又称守护线程或精灵线程。

JVM的垃圾回收线程就是典型的后台线程。

注：如果所有的前台线程都死亡，后台线程会自动死亡。

Thread类提供一个isDaemon()方法，用于判断线程是否是后台线程。

```
public class DaemonThread extends Thread
{
    // 定义后台线程的线程执行体与普通线程没有任何区别
    public void run()
    {
        for (int i = 0; i < 1000 ; i++ )
        {
            System.out.println(getName() + "  " + i);
        }
    }
    public static void main(String[] args)
    {
        DaemonThread t = new DaemonThread();
        // 将此线程设置成后台线程
        t.setDaemon(true);
        // 启动后台线程
        t.start();
        for (int i = 0 ; i < 10 ; i++ )
        {
            System.out.println(Thread.currentThread().getName()
                + "  " + i);
        }
        // -----程序执行到此处，前台线程（main线程）结束------
        // 后台线程也应该随之结束
    }
}
```

4.3线程睡眠：sleep

如果需要让当前正在执行的线程暂停一段时间，并进入阻塞状态，则可以通过调用 Thread 类的静态sleep()方法来实现。sleep()方法有两种重载形式。

(1)static void sleep ( long millis ):让当前正在执行的线程暂停 millis 毫秒，并进入阻塞状态，该方法受到系统计时器和线程调度器的精度与准确度的影响。

(2)static void sleep(long miHis，int nanos ):让当前正在执行的线程暂停 millis 毫秒加 nanos 毫微秒，并进入阻塞状态，该方法受到系统计时器和线程调度器的精度与准确度的影响。

与前面类似的是，程序很少调用第二种形式的sleep()方法。当当前线程调用 sleepO 方法进入阻寒状态后，在其睡眠时间段内，该线程不会获得执行的机会，即使系统中没冇芄他可执行的线程，处于sleep()中的线程也不会执行，因此sleep()方法常用来暂停程序的

执行。

```
public class SleepTest
{
    public static void main(String[] args)
        throws Exception
    {
        for (int i = 0; i < 10 ; i++ )
        {
            System.out.println("当前时间: " + new Date());
            // 调用sleep方法让当前线程暂停1s。
            Thread.sleep(1000);
        }
    }
}
```

4.4线程让步：yield

yield()方法是一个和sleep()方法相似的方法，它也是 Thread 类提供的一个静态方法，它也可以让当前正在执行的线程暂停，但它不会阻塞该线程，它只是将该线程转入就绪状态。 yield()只是让当前线程暂停一下，让系统的线程调度器重新调度一次，完全可能的情况足：当某个线程调用 yield ()方法暂停之后，线程调度器又将其调度出来重新执行。

实际上，当某个线程调用了  yield ()方法暂停之后，只有优先级与当前线程相同，或存优先级比当前线程更高的处于就绪状态的线程才会获得执行的机会。

```
public class YieldTest extends Thread{
    public YieldTest(String name){
        super(name);
    }
    // 定义run方法作为线程执行体
    public void run(){
        for (int i = 0; i < 50 ; i++ )
        {
            System.out.println(getName() + "  " + i);
            // 当i等于20时，使用yield方法让当前线程让步
            if (i == 20)
            {
                Thread.yield();
            }
        }
    }
    public static void main(String[] args)throws Exception
    {
        // 启动两条并发线程
        YieldTest yt1 = new YieldTest("高级");
        // 将ty1线程设置成最高优先级
        yt1.setPriority(Thread.MAX_PRIORITY);
        yt1.start();
        YieldTest yt2 = new YieldTest("低级");
        // 将yt2线程设置成最低优先级
        yt2.setPriority(Thread.MIN_PRIORITY);
        yt2.start();
    }
}
```

关于sleep()方法和 yield()方法的区别如下：

(1)sleep()方法暂停当前线程后，会给其他线程执行机会，不会理会其他线程的优先级；但yield()方法只会给优先级相同 ,或优先级更高的线程执行机会。

(2)sleep()方法会将线程转入阻塞状态，直到经过阻塞时间才会转入就绪状态；而yield()不会将线程转入阻塞状态，它只是强制当前线程进入就绪状态。因此完全有可能能某个线程调用 yield ()方法暂停之后，立即再次获得处理器资源被执行。

(3)sleep()方法声明抛出了IntemiptedException 异常，所以调用 sleep ()方法时要么捕捉该异常，要么显式声明抛出该异常：而 yield ()方法则没有声明抛出任何异常。

(4)sleep()方法比 yield()方法有更好的可移植性，通常不建议使用 yield()方法来控制并发线程的执行.

4.5设置线程优先级

每个线程执行时都具有一定的优先级，优先级高的线程获得较多的执行机会，而优先级低的线程则获得较少的执行机会。

每个线程默认的优先级都与创建它的父线程的优先级相同，在默认情况下 ,main 线程具有普通优先级，由 main线程创建的线程也具有普通优先级。

Thread 类提供了setPriority(int newPriority )、 getPriority()方法来设置和返回指定线程的优先级，其中 setPriority()方法的参数可以是一个整数，范围是1~10之间，也可以使用 Thread 类的如下三个静态常量。

(1)MAX \_ PRIORITY :其值是10 。

(2)MIN \_ PRIORITY : 其值是0

(3)NORM \_ PRIORITY: 其值是5 。

```
public class PriorityTest extends Thread
{
    // 定义一个有参数的构造器，用于创建线程时指定name
    public PriorityTest(String name)
    {
        super(name);
    }
    public void run()
    {
        for (int i = 0 ; i < 50 ; i++ )
        {
            System.out.println(getName() +  ",其优先级是："
                + getPriority() + ",循环变量的值为:" + i);
        }
    }
    public static void main(String[] args)
    {
        // 改变主线程的优先级
        Thread.currentThread().setPriority(6);
        for (int i = 0 ; i < 30 ; i++ )
        {
            if (i == 10)
            {
                PriorityTest low  = new PriorityTest("低级");
                low.start();
                System.out.println("创建之初的优先级:"
                    + low.getPriority());
                // 设置该线程为最低优先级
                low.setPriority(Thread.MIN_PRIORITY);
            }
            if (i == 20)
            {
                PriorityTest high = new PriorityTest("高级");
                high.start();
                System.out.println("创建之初的优先级:"
                    + high.getPriority());
                // 设置该线程为最高优先级
                high.setPriority(Thread.MAX_PRIORITY);
            }
        }
    }
}
```

五、线程操作——线程同步

5.1同步代码块

```
public class DrawThread extends Thread
{
    // 模拟用户账户
    private Account account;
    // 当前取钱线程所希望取的钱数
    private double drawAmount;
    public DrawThread(String name , Account account
        , double drawAmount)
    {
        super(name);
        this.account = account;
        this.drawAmount = drawAmount;
    }
    // 当多条线程修改同一个共享数据时，将涉及数据安全问题。
    public void run()
    {
        // 账户余额大于取钱数目
        if (account.getBalance() >= drawAmount)
        {
            // 吐出钞票
            System.out.println(getName()
                + "取钱成功！吐出钞票:" + drawAmount);
            try
            {
                Thread.sleep(1);
            }
            catch (InterruptedException ex)
            {
                ex.printStackTrace();
            }
            // 修改余额
            account.setBalance(account.getBalance() - drawAmount);
            System.out.println("\t余额为: " + account.getBalance());
        }
        else
        {
            System.out.println(getName() + "取钱失败！余额不足！");
        }
    }
}
```

5.2同步方法

```
public class Account
{
    // 封装账户编号、账户余额的两个成员变量
    private String accountNo;
    private double balance;
    public Account(){}
    // 构造器
    public Account(String accountNo , double balance)
    {
        this.accountNo = accountNo;
        this.balance = balance;
    }
    // 此处省略了accountNo和balance的setter和getter方法

    // accountNo的setter和getter方法
    public void setAccountNo(String accountNo)
    {
        this.accountNo = accountNo;
    }
    public String getAccountNo()
    {
        return this.accountNo;
    }

    // balance的setter和getter方法
    public void setBalance(double balance)
    {
        this.balance = balance;
    }
    public double getBalance()
    {
        return this.balance;
    }

    // 下面两个方法根据accountNo来重写hashCode()和equals()方法
    public int hashCode()
    {
        return accountNo.hashCode();
    }
    public boolean equals(Object obj)
    {
        if(this == obj)
            return true;
        if (obj !=null
            && obj.getClass() == Account.class)
        {
            Account target = (Account)obj;
            return target.getAccountNo().equals(accountNo);
        }
        return false;
    }
}
```

5.3释放同步监视器的锁定

任何线程进入同步代码块、同步方法之前，必须先获得对同步监视器的锁定，那么何时会释放对同步监视器的锁定呢？程序无法显式释放对同步监视器的锁定，线程会在如下几种情况释放对同步监视器的锁定。

(1)当前线程的同步方法、問步代码块执行结束，当前线程即释放同步监视器。

(2)当前线程在同步代码块、同步方法中遇到break、return终止了该代码块、该方法的继续执行，当前线程将会释放同步监视器。

(3)当前线程在同步代码块、同步方法中出现了末处理的 Error 或 Exception ,导致了该代码块、该方法异常结束时，当前线程将会释放同步监视器。

(4)当前线程执行同步代码块或同步方法时，程序执行了同步监视器对象的wait()方法，则当前线程暂停，并释放同步监视器。

在如下所示的怡况下，线程不会释放同步监视器。

(1)线程执行同步代码块或同步方法时，程序调用Thread.sleep()、Thread.yield()方法來暂停当前线程的执行，当前线程不会释放同步监视器。

(2)线程执行同步代码块时，其他线程调用了该线程的suspend()方法将该线程挂起，该线程不会释放同步监视器。当然，程序应该尽量避免使用suspend()和 resume()方法来控制线程。

5.4同步锁(Lock)

从Java5开始，Java提供了一种功能更加强大的线程同步机制一通过显式定义同步锁对象来实现同步，在这种机制下，同步锁由 Lock 对象充当。

Lock 提供了比 synchronized 方法和 synchronized 代码块更广泛的锁定操作， Lock 允许实现更灵活的结构，可以具有差别很大的属性，并且支持多个相关的 Condition 对象。

Lock 是控制多个线程对共亨资源进行访问的工具。通常，锁提供了对共享资源的独占访问，每次只能有一个线程对 Lock 对象加锁，线程开始访问共享资源之前应先获得 Lock 对象。

某些锁可能允许对共享资源并发访问，如ReadWriteLock(读写锁)，Lock、ReadWriteLock是Java5提供的两个根接口并为Lock提供了ReemramLock (可重入锁）实现类，为ReadWriteLock提供了ReentrantReadWriteLock 实现类。

Java8新增了新型的 StampedLock 类，在大多数场景中它可以替代传统的 ReentrantReadWriteLock。ReentrantReadWriteLock 为读写操作提供了三种锁模式： Writing 、 ReadingOptimistic 、 Reading 。在实现线程安全的控制中，比较常用的是 ReemrantLock (可重入锁）。 使用该 Lock 对象可以显式地加锁、释放锁，通常使用 ReemramLock 的代码格式如下：

```
public class Account
{
    // 定义锁对象
    private final ReentrantLock lock = new ReentrantLock();
    // 封装账户编号、账户余额的两个成员变量
    private String accountNo;
    private double balance;
    public Account(){}
    // 构造器
    public Account(String accountNo , double balance)
    {
        this.accountNo = accountNo;
        this.balance = balance;
    }

    // accountNo的setter和getter方法
    public void setAccountNo(String accountNo)
    {
        this.accountNo = accountNo;
    }
    public String getAccountNo()
    {
        return this.accountNo;
    }
    // 因此账户余额不允许随便修改，所以只为balance提供getter方法，
    public double getBalance()
    {
        return this.balance;
    }

    // 提供一个线程安全draw()方法来完成取钱操作
    public void draw(double drawAmount)
    {
        // 加锁
        lock.lock();
        try
        {
            // 账户余额大于取钱数目
            if (balance >= drawAmount)
            {
                // 吐出钞票
                System.out.println(Thread.currentThread().getName()
                    + "取钱成功！吐出钞票:" + drawAmount);
                try
                {
                    Thread.sleep(1);
                }
                catch (InterruptedException ex)
                {
                    ex.printStackTrace();
                }
                // 修改余额
                balance -= drawAmount;
                System.out.println("\t余额为: " + balance);
            }
            else
            {
                System.out.println(Thread.currentThread().getName()
                    + "取钱失败！余额不足！");
            }
        }
        finally
        {
            // 修改完成，释放锁
            lock.unlock();
        }
    }

    // 下面两个方法根据accountNo来重写hashCode()和equals()方法
    public int hashCode()
    {
        return accountNo.hashCode();
    }
    public boolean equals(Object obj)
    {
        if(this == obj)
            return true;
        if (obj !=null
            && obj.getClass() == Account.class)
        {
            Account target = (Account)obj;
            return target.getAccountNo().equals(accountNo);
        }
        return false;
    }
}
```

5.6死锁

当两个线程相互等待对方释放同步监视器时就会发生死锁，Java虚拟机没有监测，也没有采取措施来处理死锁情况，所以多线程编程时应该采取措施避免死锁出现。一旦出现死锁，整个程序既不会发生任何异常，也不会给出任何提示，只是所有线程处于阻塞状态，无法继续。

```
class A
{
    public synchronized void foo( B b )
    {
        System.out.println("当前线程名: " + Thread.currentThread().getName()
            + " 进入了A实例的foo()方法" );     // ①
        try
        {
            Thread.sleep(200);
        }
        catch (InterruptedException ex)
        {
            ex.printStackTrace();
        }
        System.out.println("当前线程名: " + Thread.currentThread().getName()
            + " 企图调用B实例的last()方法");    // ③
        b.last();
    }
    public synchronized void last()
    {
        System.out.println("进入了A类的last()方法内部");
    }
}
class B
{
    public synchronized void bar( A a )
    {
        System.out.println("当前线程名: " + Thread.currentThread().getName()
            + " 进入了B实例的bar()方法" );   // ②
        try
        {
            Thread.sleep(200);
        }
        catch (InterruptedException ex)
        {
            ex.printStackTrace();
        }
        System.out.println("当前线程名: " + Thread.currentThread().getName()
            + " 企图调用A实例的last()方法");  // ④
        a.last();
    }
    public synchronized void last()
    {
        System.out.println("进入了B类的last()方法内部");
    }
}
public class DeadLock implements Runnable
{
    A a = new A();
    B b = new B();
    public void init()
    {
        Thread.currentThread().setName("主线程");
        // 调用a对象的foo方法
        a.foo(b);
        System.out.println("进入了主线程之后");
    }
    public void run()
    {
        Thread.currentThread().setName("副线程");
        // 调用b对象的bar方法
        b.bar(a);
        System.out.println("进入了副线程之后");
    }
    public static void main(String[] args)
    {
        DeadLock dl = new DeadLock();
        // 以dl为target启动新线程
        new Thread(dl).start();
        // 调用init()方法
        dl.init();
    }
}
```

六、线程通信

6.1传统通信方式

为了实现线程协调运行功能，可以借助于Object类提供的wait()、notify )和 notifyAll()三个方法，这三个方法并不属于 Thread 类，而是属于 Object 类。但这三个方法必须由同步监视器对象来调用，这可分成以下两种情况。

(1)对于使用 synchronized 修饰的同步方法，因为该类的默认实例（this〉就是同步监视器，所以可以在同步方法中直接调用这三个方法。

(2)对于使用 synchronized 修饰的同步代码块，同步监视器是 synchronized 后括号里的对象，所以必须使用该对象调用这三个方法。

关于这三个方法的解释如下•

(1)wait():导致当前线程等待，直到其他线程调用该同步监视器的notify()方法或notifyAll()方法来唤醒该线程。该wait()方法有三种形式——无时间参数的wait(一直等待，直到其他线程通知）、带毫秒数的wait()和带毫秒、毫微秒参数的wait()(这两种方法都是等待指定时间后自动苏醒)调用wait()方法的当前线程会释放对该同步监视器的锁定。

(2)notify():唤解在此同步监视器上等待的单个线程。如果所有线程都在此同步监视器上等待，则会选择唤醒其中一个线程。选择是任意性的。只有当前线程放弃对该同步监视器的锁定后（使用wait()方法)，才可以执行被唤醒的线程。

(3)notifyAll()方法:唤醒在此同步监视器上等待的所有线程。只有当前线程放弃对该同步监视器的锁定后，才可以执行被唤醒的线程。

6.2使用Condition控制线程通信

如果程序不使用synchronized 关键字来保证同步，而是直接使用 Lock 对象来保证间步，则系统中不存在隐式的冋步监视器，也就不能使用wait()、notify ()、 notifyAll()方法进行线程通信。

当使用 Lock 对象来保证同步时， Java 提供了一个 Condition 类来保持协调，使用 Condition 可以让那些已经得到 Lock 对象却无法继续执行的线程释放 Lock 对象， Condition 对象也可以唤解其他处于等待的线程。

Condition 将同步监视器方法（wait()、notify()和notifyAll()分解成截然不同的对象，以便通过将这些对象与 Lock 对象组合使用，为每个对象提供多个等待集（ wait - set ).在这种情况下， Lock 替代了同 步方法或同步代码块， Condition 替代了同步监视器的功能。

Condition 实例被绑定在一个 Lock 对象上。要获得特定 Lock 实例的 Condition 实例，调用 Lock 对象的 newCondition ()方法即可。 Condition 类提供了如下三个方法:

(1)await():类似于隐式同步监视器上的 wait()方法，导致当前线程等待，直到其他线程调用该 Condition的signal()方法或 signalAll()方法来唤醒该线程。该await()方法有更多变体，如Long awaitNanos(long nanosTimeout)、 void awaitUninterruptibly()、 awaitUntil(Date deadLine )等，可以完成更丰富的等待操作。

(2)signal(): 唤醒在此 Lock 对象上等待的单个线程。如果所有线程都在该 Lock 对象上等待，则会选择唤醒其中一个线程。选择是任意性的。只有当前线程放弃对该 Lock 对象的锁定后（使用await()方法），才可以执行被唤醒的线程。

(3)signalAll():唤醒在此 Lock 对象上等待的所有线程。只有当前线程放弃对该 Lock 对象的锁定后，才可以执行被唤醒的线程。

```
public class Account
{
    // 显式定义Lock对象
    private final Lock lock = new ReentrantLock();
    // 获得指定Lock对象对应的Condition
    private final Condition cond  = lock.newCondition();
    // 封装账户编号、账户余额的两个成员变量
    private String accountNo;
    private double balance;
    // 标识账户中是否已有存款的旗标
    private boolean flag = false;

    public Account(){}
    // 构造器
    public Account(String accountNo , double balance)
    {
        this.accountNo = accountNo;
        this.balance = balance;
    }

    // accountNo的setter和getter方法
    public void setAccountNo(String accountNo)
    {
        this.accountNo = accountNo;
    }
    public String getAccountNo()
    {
        return this.accountNo;
    }
    // 因此账户余额不允许随便修改，所以只为balance提供getter方法，
    public double getBalance()
    {
        return this.balance;
    }

    public void draw(double drawAmount)
    {
        // 加锁
        lock.lock();
        try
        {
            // 如果flag为假，表明账户中还没有人存钱进去，取钱方法阻塞
            if (!flag)
            {
                cond.await();
            }
            else
            {
                // 执行取钱
                System.out.println(Thread.currentThread().getName()
                    + " 取钱:" +  drawAmount);
                balance -= drawAmount;
                System.out.println("账户余额为：" + balance);
                // 将标识账户是否已有存款的旗标设为false。
                flag = false;
                // 唤醒其他线程
                cond.signalAll();
            }
        }
        catch (InterruptedException ex)
        {
            ex.printStackTrace();
        }
        // 使用finally块来释放锁
        finally
        {
            lock.unlock();
        }
    }
    public void deposit(double depositAmount)
    {
        lock.lock();
        try
        {
            // 如果flag为真，表明账户中已有人存钱进去，则存钱方法阻塞
            if (flag)             // ①
            {
                cond.await();
            }
            else
            {
                // 执行存款
                System.out.println(Thread.currentThread().getName()
                    + " 存款:" +  depositAmount);
                balance += depositAmount;
                System.out.println("账户余额为：" + balance);
                // 将表示账户是否已有存款的旗标设为true
                flag = true;
                // 唤醒其他线程
                cond.signalAll();
            }
        }
        catch (InterruptedException ex)
        {
            ex.printStackTrace();
        }
        // 使用finally块来释放锁
        finally
        {
            lock.unlock();
        }
    }

    // 下面两个方法根据accountNo来重写hashCode()和equals()方法
    public int hashCode()
    {
        return accountNo.hashCode();
    }
    public boolean equals(Object obj)
    {
        if(this == obj)
            return true;
        if (obj !=null
            && obj.getClass() == Account.class)
        {
            Account target = (Account)obj;
            return target.getAccountNo().equals(accountNo);
        }
        return false;
    }
}
```

6.3使用阻塞队列（BlockingQueue）控制线程通信

Java5提供了一个BlockingQueue接口，虽然 BlockingQueue 也是Queue的子接口，但它的主要用

途并不是作为容器，而是作为线程同步的工具 。 BlockingQueue 具有一个特征：当生产者线程试图向

BlockingQeue中放入元素时，如果该队列己满，则该线程被阻塞：当消费者线程试阁从 BlockingQueue

中取出元素时，如果该队列已空，则该线程被阻塞。

程序的两个线程通过交替向 BlockingQueue 中放入元素、取出元素，即可很好地控制线程的通信。

BlockingQueue 提供如下两个支持阻塞的方法。

(1)put(E e) : 尝试把 E 元素放入 BlockingQueue 中，如果该队列的元素己满，则阻塞该线程.

(2)take(): 尝试从 BlockingQueue 的头部取出元素，如果该队列的元素己空，则阻塞该线程。

BlockingQueue继承了Queue接口，当然也可使用Queue接口中的方法。这些方法归纳起来可分为

如下三组。

(1)在队列尾部插入元素。包括add(E e)、 offer(E e)和 put(E e)方法，当该队列已满时，这三个方法

分别会抛出异常、返回 false 、阻塞队列。

(2)在队列头部删除并返回刪除的元素。包括 remove() 、poll()和take()方法。当该队列已空时，这

三个方法分别会抛出异常、返冋 false 、阻塞队列。

(3)在队列头部取出但不删除元素。包括 element()和 peek()方法，当队列已空时，这两个方法分别

抛出异常、返回 false 。

```
class Producer extends Thread
{
    private BlockingQueue<String> bq;
    public Producer(BlockingQueue<String> bq)
    {
        this.bq = bq;
    }
    public void run()
    {
        String[] strArr = new String[]
        {
            "Java",
            "Struts",
            "Spring"
        };
        for (int i = 0 ; i < 999999999 ; i++ )
        {
            System.out.println(getName() + "生产者准备生产集合元素！");
            try
            {
                Thread.sleep(200);
                // 尝试放入元素，如果队列已满，线程被阻塞
                bq.put(strArr[i % 3]);
            }
            catch (Exception ex){ex.printStackTrace();}
            System.out.println(getName() + "生产完成：" + bq);
        }
    }
}
class Consumer extends Thread
{
    private BlockingQueue<String> bq;
    public Consumer(BlockingQueue<String> bq)
    {
        this.bq = bq;
    }
    public void run()
    {
        while(true)
        {
            System.out.println(getName() + "消费者准备消费集合元素！");
            try
            {
                Thread.sleep(200);
                // 尝试取出元素，如果队列已空，线程被阻塞
                bq.take();
            }
            catch (Exception ex){ex.printStackTrace();}
            System.out.println(getName() + "消费完成：" + bq);
        }
    }
}
public class BlockingQueueTest2
{
    public static void main(String[] args)
    {
        // 创建一个容量为1的BlockingQueue
        BlockingQueue<String> bq = new ArrayBlockingQueue<>(1);
        // 启动3条生产者线程
        new Producer(bq).start();
        new Producer(bq).start();
        new Producer(bq).start();
        // 启动一条消费者线程
        new Consumer(bq).start();
    }
}
```

七、线程池

7.1Executors工厂类创建线程池

使用线程池来执行线程任务的步骤如下：

(1)调用Executors类的静态工厂方法创建一个ExecutorService对象，该对象代表一个线程池。

(2)创建Runnable实现类或Callable实现类的实例，作为线程执行任务。

(3)调用ExecutorService对象的submit()方法来提交Runnable实例或Callable实例。

(4)当不想提交任何任务时，调用ExecutorService对象的shutdown()方法来关闭线程池。

```
public class ThreadPoolTest
{
	public static void main(String[] args)
		throws Exception
	{
		// 创建足够的线程来支持4个CPU并行的线程池
		// 创建一个具有固定线程数（6）的线程池
		ExecutorService pool = Executors.newFixedThreadPool(6);
		// 使用Lambda表达式创建Runnable对象
		Runnable target = () -> {
			for (int i = 0; i < 100 ; i++ )
			{
				System.out.println(Thread.currentThread().getName()
					+ "的i值为:" + i);
			}
		};
		// 向线程池中提交两个线程
		pool.submit(target);
		pool.submit(target);
		// 关闭线程池
		pool.shutdown();
	}
}


```

7.2ForkJoinPool

```
// 继承RecursiveAction来实现"可分解"的任务
class PrintTask extends RecursiveAction
{
	// 每个“小任务”只最多只打印50个数
	private static final int THRESHOLD = 50;
	private int start;
	private int end;
	// 打印从start到end的任务
	public PrintTask(int start, int end)
	{
		this.start = start;
		this.end = end;
	}
	@Override
	protected void compute()
	{
		// 当end与start之间的差小于THRESHOLD时，开始打印
		if(end - start < THRESHOLD)
		{
			for (int i = start ; i < end ; i++ )
			{
				System.out.println(Thread.currentThread().getName()
					+ "的i值：" + i);
			}
		}
		else
		{
			// 如果当end与start之间的差大于THRESHOLD时，即要打印的数超过50个
			// 将大任务分解成两个小任务。
			int middle = (start + end) / 2;
			PrintTask left = new PrintTask(start, middle);
			PrintTask right = new PrintTask(middle, end);
			// 并行执行两个“小任务”
			left.fork();
			right.fork();
		}
	}
}
public class ForkJoinPoolTest
{
	public static void main(String[] args)
		throws Exception
	{
		ForkJoinPool pool = new ForkJoinPool();
		// 提交可分解的PrintTask任务
		pool.submit(new PrintTask(0 , 300));
		pool.awaitTermination(2, TimeUnit.SECONDS);
		// 关闭线程池
		pool.shutdown();
	}
}


```



