## go语言创建后台守护应用
```
linux 上创建 daemon 的步骤一般如下:

创建子进程，父进程退出；
调用系统调用 setsid() 脱离控制终端；
调用系统调用 umask() 清除进程 umask 确保 daemon 创建文件和目录时拥有所需权限；
修改当前工作目录为系统根目录；
关闭从父进程继承的所有文件描述符，并将标准输入/输出/错误重定向到 /dev/null。
```







## 参考
* http://litang.me/post/golang-server-design/