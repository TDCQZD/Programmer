# VScode配置C语言开发环境

- 安装VSCode
- 在VSCode内安装c++插件
- 安装编译、调试环境
- 修改VSCode调试配置文件

## 安装VSCode
* 下载地址：https://code.visualstudio.com/?utm_expid=101350005-25.TcgI322oRoCwQD7KJ5t8zQ.0
## 在VSCode内安装c++插件
- C/C++ for Visual Studio Code
## 安装编译、调试环境
- MinGW
- Cygwin
### MinGW
- http://mingw-w64.org/doku.php/download
> MingW-W64-builds
1. 下载软件并安装
2. 配置环境变量
3. 测试安装
```
gcc -v
```
## 修改VSCode调试配置文件
1. 工作目录下新建.vscode文件夹
2. 新建三个配置文件launch.json、tasks.json、c_cpp_properties.json

## 快速开始
```
gcc hello.c -o hello
hello
```
## 参考资料
* https://segmentfault.com/a/1190000014800106