# 设计

感觉代码不多，所以就不用分模块了，那要分文件吗？好像也没必要，直接在一个文件 bulkdo.go 里分方法完成即可。

我想我们这个程序有几个事情要做，直接写几个方法：

1. 读取 csv 文件为一个 items 的 map

readItems

2. 根据模板和 items map 得到执行的脚本文本

parseCommands

3. 执行脚本文本

execCommands

4. 组合这些方法