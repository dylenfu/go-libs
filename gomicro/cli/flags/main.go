package main

import (
	"fmt"
	"github.com/micro/go-micro/util/log"
	"os"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(flags)
	service.Init(action)
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

var flags = micro.Flags(
	cli.StringFlag{
		Name:  "string_flag",
		Usage: "string_flag是字符串flag",
		Value: "我要被覆盖",
	},
	cli.IntFlag{
		Name:  "int_flag",
		Usage: "int_flag是整形flag",
		// 可以使用EnvVar来声明支持使用环境变量，但是如果flag和env同时使用时，以flag为准
		EnvVar: "INT_FLAG",
	},
	cli.BoolFlag{
		Name:  "bool_flag",
		Usage: "bool_flag是布尔值的flag",
	},
	cli.StringFlag{
		Name:  "string_flag_default",
		Usage: "我是缺省值",
		Value: "我是缺省值",
	},
)

var action = micro.Action(func(c *cli.Context) {
	log.Logf("字符串flag值: %s\n", c.String("string_flag"))
	log.Logf("字符串缺省值: %s\n", c.String("string_flag_default"))
	log.Logf("整形flag值: %d\n", c.Int("int_flag"))
	log.Logf("布尔值flag值: %t\n", c.Bool("bool_flag"))
	//打印完退出
	os.Exit(0)
})
