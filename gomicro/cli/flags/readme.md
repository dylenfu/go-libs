flags

运行main函数:
INT_FLAG=109 go run main.go --string_flag="串串来了" --bool_flag=true

打印信息如下:
2019/10/15 10:29:38 字符串flag值: 串串来了
2019/10/15 10:29:38 字符串缺省值: 我是缺省值
2019/10/15 10:29:38 整形flag值: 109
2019/10/15 10:29:38 布尔值flag值: true

micro/flag支持环境变量，