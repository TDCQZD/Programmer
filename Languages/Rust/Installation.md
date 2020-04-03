# Installation

## 安装rustup

### rustup 命令
`rustup`命令行工具用于管理Rust版本和相关工具
``` 
rustc --version # Check rustup
rustup update # Updating
rustup self uninstall # Uninstalling
rustup doc # Open Local Documentation
```

### cargo  命令
```
cargo --version # check cargo
cargo new hello-rust # 
```

# Hello World!
## 1. Creating a Project Directory
For Linux, macOS, and PowerShell on Windows, enter this:
``` 
$ mkdir ~/projects
$ cd ~/projects
$ mkdir hello_world
$ cd hello_world

```
For Windows CMD, enter this:
```
> mkdir "%USERPROFILE%\projects"
> cd /d "%USERPROFILE%\projects"
> mkdir hello_world
> cd hello_world
```

## 2. Writing and Running a Rust Program
1. open the `main.rs` file
2. enter the code 
```
fn main() {
    println!("Hello, world!");
}

```
3. compile and run the file
```On Linux or macOS
$ rustc main.rs
$ ./main
Hello, world!
```
```On Windows
> rustc main.rs
> .\main.exe
Hello, world!

```

# Hello Cargo!
Cargo是Rust的构建系统和包管理器.使用此工具来管理Rust项目，Cargo会处理很多任务，例如构建代码，下载代码所依赖的库以及构建这些库。（我们称库为您的代码需要依赖项。

## 1. Creating a Project with Cargo
```
$ cargo new hello_cargo
$ cd hello_cargo

```
Cargo为我们生成了两个文件和一个目录：一个`Cargo.toml`文件和一个其中带有`main.rs`文件的 src目录。它还已经初始化了一个新的Git存储库以及一个.`gitignore`文件

* Cargo.toml:
```
[package]
name = "hello_cargo"
version = "0.1.0"
authors = ["Your Name <you@example.com>"]
edition = "2018"

[dependencies]

```
## 2. Building and Running a Cargo Project
```
$ cargo build # build a project 
$ cargo check # build a project 

$ cargo run # build and run a project

$ cargo build --release # Building for Release
```

