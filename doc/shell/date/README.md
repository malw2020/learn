[*back to shell*](https://github.com/malw2020/learn/tree/master/doc/shell#contents)<br>

# Shell: date

- [date to timestamp](#date-to-unix-timestamp)
- [timestamp to date](#unix-timestamp-to-date)

[↑ top](#shell-date)
<br><br><br><br><hr>


#### date to unix timestamp

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


#### unix timestamp to date

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
