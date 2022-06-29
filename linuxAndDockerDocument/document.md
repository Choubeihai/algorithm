# 1. Namespaces

命令空间是Linux内核用来隔离内核资源的一种方式，通过命令空间可以让一些进程只能看到与自己相关的一部分资源，而另外一些进程也只能看到与它们自己相关的资源，这两拨进程根本就感觉不到对方的存在。具体实现方法就是把不同的进程放进不同的命令空间里，则会让他们相互看不到。

在日常使用 Linux 或者 macOS 时，我们并没有运行多个完全分离的服务器的需要，但是如果我们在服务器上启动了多个服务，这些服务其实会相互影响的，每一个服务都能看到其他服务的进程，也可以访问宿主机器上的任意文件，这是很多时候我们都不愿意看到的，我们更希望运行在同一台机器上的不同服务能做到完全隔离，就像运行在多台不同的机器上一样。

有六种Namespaces，即提供了六类系统资源的隔离方式。

1. Mount: 隔离文件系统挂载点
2. UTS: 隔离主机名和域名信息
3. IPC: 隔离进程间通信
4. PID: 隔离进程的ID
5. Network: 隔离网络资源
6. User: 隔离用户和用户组的ID

## 1. API使用

涉及的API：

1. clone()

2. setns()

3. unshare()

Namespaces的参数：

1. CLONE_NEWNS: 用于指定Mount Namespace 文件系统挂载点

2. CLONE_NEWUTS: 用于指定UTS Namespace  主机名和域名信息

3. CLONE_NEWIPC: 用于指定IPC Namespace	 进程间通信

4. CLONE_NEWPID: 用于指定PID Namespace	 进程ID

5. CLONE_NEWNET: 用于指定Network Namespace	网络资源
6. CLONE_NEWUSER: 用于指定User Namespace  用户和用户组

示例解释API：

```c
int clone(int (*child_func)(void *), void *child_stack, int flags, void *arg)
```

flags = CLONE_NEWNS | CLONE_NEWUTS | CLONE_NEWIPC，通过flag参数控制子进程的命名空间，传入这个flags那么新创建的进程将同时拥有独立的Mount Namespace、UTS Namespace和IPC Namespace。

```c
int setns(int fd, int nstype)
```

setns()函数可以把进程加入到指定的Namespace中，其中fd是对应的Namespaces，	nstype设置为0表示不检测该Namespaces类型，直接将进程加入Namespaces。

```c
int unshare(int flags)
```

unshare()系统调用用于将当前进程和所在的Namespace分离，并加入到一个新的	Namespace中，相对于setns()系统调用来说，unshare()不用关联之前存在的Namespace，	只需要指定需要分离的Namespace就行，该调用会自动创建一个新的Namespace。

