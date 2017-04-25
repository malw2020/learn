[*back to shell*](https://github.com/malw2020/learn/tree/master/doc/shell#contents)<br>

# Shell: df

- [df description](#df-description)
- [df exam](#crontab-df)
- [crontab service](#crontab-service)

[↑ top](#shell-df)
<br><br><br><br><hr>


#### df description

```shell
	基本格式 : 
		df [选项] [文件]
	
	必要参数：
		-a 全部文件系统列表
		-h 方便阅读方式显示
		-H 等于“-h”，但是计算式，1K=1000，而不是1K=1024
		-i 显示inode信息
		-k 区块为1024字节
		-l 只显示本地文件系统
		-m 区块为1048576字节
		--no-sync 忽略 sync 命令
		-P 输出格式为POSIX
		--sync 在取得磁盘信息前，先执行sync命令
		-T 文件系统类型
	
	显示指定磁盘文件的可用空间。如果没有文件名被指定，则所有当前被挂载的文件系统的可用空间将被显示。默认情况下，磁盘空间将以 1KB 为单位进行显示，除非环境变量 POSIXLY_CORRECT 被指定，那样将以512字节为单位进行显示

```

[↑ top](#shell-df)
<br><br><br><br><hr>

#### df exam

```shell
	1. df
	   显示磁盘使用情况
	2. df -T 
	   列出文件系统的类型
	3. df -h
	   以更易读的方式显示目前磁盘空间和使用情况  
	4.  df -H 
	    以更易读的方式显示目前磁盘空间和使用情况，和上面的-h参数相同,不过在格式化的时候,采用1000而不是1024进行容量转换 
	5. df -t ext3
	   显示指定类型磁盘 

```

[↑ top](#shell-df)
<br><br><br><br><hr>



