# bulkdo

一个批量执行小工具。

## 安装

TODO

## 使用说明

首先需要配置命令行模板文件。到 ~/.bulkdo/ 下创建命令模板文件，例如  command1.tpl ，其中内容为 go 模板，例如：

```shell
echo {{ .v.t1 }} {{ .v.t2 }}
```

在我们希望执行 bulkdo 命令的目录下，有个文件名为 .bulkdoitems 的文件，里面的内容为：

```
t1,t2,t3
a1,a2,a3
b1,b2,b3
```

然后执行 bulkdo command1 ，就会把 command1 的变量替换后执行

```shell
echo a1 a2
echo b1 b2
```
