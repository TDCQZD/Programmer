
//不好的解法一：只适用于单线程环境
public class Singleton_1 {
    /* 持有私有静态实例，防止被引用，此处赋值为 null，目的是实现延迟加载 */
    private static Singleton_1 instance = null;

    /* 私有构造方法，防止被实例化 */
    private Singleton_1() {
        // TODO Auto-generated constructor stub
    }

    /* 静态工程方法，创建实例 */
    public static Singleton_1 getInstance() {

        if (instance == null) {// 多个线程判断instance都为null时，在执行new操作时多线程会出现重复情况
            instance = new Singleton_1();
        }

        return instance;

    }

}


// 不好的解法二：虽然在多线程环境中能工作但效率不高(单线程+同步锁)
public class Singleton_2 {
    /* 持有私有静态实例，防止被引用，此处赋值为 null，目的是实现延迟加载 */
    private static Singleton_2 instance = null;

    /* 私有构造方法，防止被实例化 */
    private Singleton_2() {
    }

    /* 静态工程方法，创建实例 */
    public static Singleton_2 getInstance() {

        synchronized (Singleton_2.class) {
            if (instance == null) {
                instance = new Singleton_2();
            }
        }
        return instance;
    }

}
// 可行的解法：加同步锁前后两次判断实例是否已存在

public class Singleton_3 {
    /* 持有私有静态实例，防止被引用，此处赋值为 null，目的是实现延迟加载 */
    private static volatile Singleton_3 instancle = null;

    /* 私有构造方法，防止被实例化 */
    private Singleton_3() {
    }

    /* 静态工程方法，创建实例 */
    public static Singleton_3 getInstance() {
        if (instancle == null) {
            synchronized (Singleton_3.class) {
                if (instancle == null) {
                    instancle = new Singleton_3();
                }
            }
        }

        return instancle;

    }

}

// 强烈推荐的解法一：利用静态构造函数

public class Singleton_4 {
    private Singleton_4() {
    }

    private static Singleton_4 instancle = new Singleton_4();

    public static Singleton_4 getInstancle() {
        return instancle;
    }

}

// 强烈推荐的解法二：实现按需创建实例,很好地解决了Singleton_4中创建实例时机过早的问题
public class Singleton_5 {
    private Singleton_5() {
    }

    public static Singleton_5 getInstancle() {
        return Nested.instancle;
    }

    public static class Nested {
        public Nested() {

        }

        private static Singleton_5 instancle = new Singleton_5();

    }
}