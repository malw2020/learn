[*back to contents*](https://github.com/malw2020/learn#contents)<br>

# Shell: date

- [read file](#将日期转换为Unix时间戳)
- [write file](#将Unix时间戳转换为日期时间)

[↑ top](#python-os-io)
<br><br><br><br><hr>


#### 将日期转换为Unix时间戳

```shell
1.将当前时间以Unix时间戳表示：
  date +%s
   输出如下：
   1361542433
2.转换指定日期为Unix时间戳：
  date -d '2013-2-22 22:14' +%s
   输出如下：
   1361542440

```

[↑ top](#shell-date)
<br><br><br><br><hr>


#### write file

```shell
1.不指定日期时间的格式：
  date -d @1361542596
   输出如下：
   Fri Feb 22 22:16:36 CST 2013
2.指定日期格式的转换：
  date -d @1361542596 +"%Y-%m-%d %H:%M:%S"
   输出如下：
   2013-02-22 22:16:36

```

[↑ top](#shell-date)
<br><br><br><br><hr>
