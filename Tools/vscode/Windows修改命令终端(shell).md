# Windows修改命令终端(shell)

1. 在文件 -> 首选项 -> 设置中打开settings页面, 搜索shell或则找到Terminal>Integrated>Shell:Windows,
2. 把地址切换成你需要以下的某个shell的地址（powershell, cmd, 和bash）
```
// Command Prompt
"terminal.integrated.shell.windows": "C:\\Windows\\System32\\cmd.exe"
// PowerShell
"terminal.integrated.shell.windows": "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe"
// Git Bash
"terminal.integrated.shell.windows": "C:\\Program Files\\Git\\bin\\bash.exe"
// Bash on Ubuntu (on Windows)
"terminal.integrated.shell.windows": "C:\\Windows\\System32\\bash.exe"
```
## 参考资料
* https://code.visualstudio.com/docs/editor/integrated-terminal#_configuration
