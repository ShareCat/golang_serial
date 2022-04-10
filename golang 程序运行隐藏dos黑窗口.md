### golang 程序运行隐藏dos黑窗口

通过go的标准库exec调用cmd命令时会闪弹黑窗口，为解决此问题在windows下可以用win32 API 的 WinExec。

此问题主要出现在带UI或无控制台的程序调用cmd时。

编译go时加入参数：
```batch
go build -ldflags=”-H windowsgui”
```
