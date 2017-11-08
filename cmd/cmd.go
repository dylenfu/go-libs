package cmd

import (
	"errors"
	"fmt"
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
)

func SimpleCli() {
	demo11()
}

func demo1() {
	app := cli.NewApp()
	app.Run(os.Args)
}

func demo2() {
	app := cli.NewApp()
	app.Name = "loopring"
	app.Usage = "DIAOZHATIAN DE ICO"
	app.Authors = []cli.Author{cli.Author{Name: "Dylen Fu", Email: "dylenfu@126.com"}}
	app.Action = func(c *cli.Context) error {
		fmt.Println("diao zha tian de ico,jiu wen ni pabupa")
		return nil
	}
	app.Run(os.Args)
}

/*
go run main.go --lang=spanish
2017/08/26 14:20:49 cli	- number of args: 0
2017/08/26 14:20:49 cli	- args(0):
hola Nefertiti

go run main.go --lang=spanish t
2017/08/26 14:21:29 cli	- number of args: 1
2017/08/26 14:21:29 cli	- args(0): t
hola t

go run main.go --lang=spanish t maybe
2017/08/26 14:22:58 cli	- number of args: 2
2017/08/26 14:22:58 cli	- args(0): t
hola maybe
*/
func demo3() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang",
			Value: "english",
			Usage: "language for the greeting",
		},
	}

	app.Action = func(c *cli.Context) error {
		name := "Nefertiti"

		log.Println("cli\t-", "number of args:", c.NArg())
		log.Println("cli\t-", "args(0):", c.Args().Get(0))
		if c.NArg() > 0 {
			name = c.Args().Get(1)
		}
		if c.String("lang") == "spanish" {
			fmt.Println("hola", name)
		} else {
			fmt.Println("hello", name)
		}
		return nil
	}

	app.Run(os.Args)
}

// --config or -c
func demo4() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "load configuration from `FILE`",
		},
	}

	app.Run(os.Args)
}

func demo5() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "lang, l",
			Value:  "english",
			Usage:  "language for the greeting",
			EnvVar: "LEGACY_COMPAT_LANG,APP_LANG,LANG",
		},
	}

	app.Run(os.Args)
}

// commands
/*
go run main.go add
2017/08/28 11:16:58 cli test	- command add

go run main.go complete
2017/08/28 11:17:08 cli test	- commands complete
*/
func demo6() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "language for greeting",
		},
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Load configuration from File",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task in list",
			Action: func(c *cli.Context) error {
				log.Println("cli test\t-", "commands", "complete")
				return nil
			},
		},

		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task to a list",
			Action: func(c *cli.Context) error {
				log.Println("cli test\t-", "command", "add")
				return nil
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	app.Run(os.Args)
}

// sub commands
func demo7() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name: "loop",
		},
		{
			Name:     "add",
			Category: "template",
			Usage:    "add task in list",
			Action: func(c *cli.Context) {
				log.Println("cli test\t-", "sub commands", "add")
			},
		},
		{
			Name:     "remove",
			Category: "template",
			Usage:    "remove a task from list",
			Action: func(c *cli.Context) {
				log.Println("cli test\t-", "sub commands", "remove")
			},
		},
	}

	app.Run(os.Args)
}

// exit
func demo8() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolTFlag{
			Name:  "ginger-crouton",
			Usage: "is it in the soup?",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		if !ctx.Bool("ginger-crouton") {
			return cli.NewExitError("it is not in the soup", 86)
		}
		return nil
	}

	app.Run(os.Args)
}

func demo9() {
	tasks := []string{"cook", "clean", "laundry", "eat", "sleep", "code"}

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task in list",
			Action: func(c *cli.Context) error {
				log.Println("cli test\t-", "bash completion", c.Args().First())
				return nil
			},
		},

		{
			BashComplete: func(c *cli.Context) {
				// This will pass if no args complete
				if c.NArg() > 0 {
					log.Println("cli test\t-", "bash completion", "Narg")
				}

				for _, t := range tasks {
					log.Println("cli test\t-", "bash completion", t)
				}
			},
		},
	}

	app.Run(os.Args)
}

func demo10() {
	cli.BashCompletionFlag = cli.BoolFlag{
		Name:   "compgen",
		Hidden: true,
	}

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name: "wat",
		},
	}
	app.Run(os.Args)
}

/////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// 参考eth写一个自己的demo
// 1.首先要有一个自己的service，app启动之后运行这个service
// 2.接收到相关命令后，route到相关function，function中根据解析的参数运行
// 3.这里需要注意的几个点：
// 	.cmd对应的action，入口参数一定是*cli.context
//  .cmd对应的action，可以没有返回值
//  .app.action可能会不被执行，如果直接运行某个命令的话
// go run main.go greet -e=dev --id=0x123 --desc=halo
// 2017/08/28 16:43:04 cli test	- production 0x123 halo dev
// go run main.go run start
// 2017/08/28 16:49:46 cli test	- production {0x1 hello  0xc420074600}

func demo11() {
	app := newApp("Loopring/ringminer.git", "decentralize exchange")
	app.Action = lrc
	app.Copyright = "Copyright 2013-2017 The loopring/ringminer Authors"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		greetCmd,
		runCmd,
	}
	app.Flags = []cli.Flag{
		evnFlag,
		idFlag,
		descFlag,
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	app.Before = func(cli *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}

	app.After = func(cli *cli.Context) error {
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
		os.Exit(0)
	}
}

func lrc(cli *cli.Context) {
	log.Println("cli test\t-", "welcome to urfave/cli demo for loopring production")
}

func newApp(gitCommit, usage string) *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Author = "dylenfu"
	app.Email = "dylenfu@126.com"
	app.Version = gitCommit + "0.0.1"
	app.Usage = usage
	return app
}

type Stack struct {
	ID   string
	Desc string
	Env  string
	stop chan struct{}
}

var s Stack

func Start(cli *cli.Context) {
	s.ID = "0x1"
	s.Desc = "hello"
	s.stop = make(chan struct{})

	log.Println("cli test\t-", "production", s)

	<-s.stop
}

func Stop(cli *cli.Context) {
	close(s.stop)
}

func Greeting(c *cli.Context) error {
	if !c.IsSet("id") {
		return errors.New("id should be set")
	}
	id := c.String("id")
	desc := c.String("desc")
	env := c.String("env")
	log.Println("cli test\t-", "production", id, desc, env)
	return nil
}

// 设计总共2个大的command，一个有subcommand，另一个没有
// 每个command带有1到两个flag
var (
	evnFlag = cli.StringFlag{
		Name:  "env,e",
		Usage: "select envionment for execute program",
		Value: defaultEnv(),
	}

	idFlag = cli.StringFlag{
		Name:  "id",
		Usage: "set id for stack",
	}

	descFlag = cli.StringFlag{
		Name:  "desc",
		Usage: "set description for stack",
	}

	greetCmd = cli.Command{
		Action:    Greeting,
		Name:      "greet",
		Usage:     "say hi",
		ArgsUsage: "<environment>,<Hash ID>,<description>",
		Flags: []cli.Flag{
			evnFlag,
			idFlag,
			descFlag,
		},
		Category: "IMPORT COMMANDS",
		Description: `
		Here, we build an demo struct node, import command will
		set its' id,desc and environment.
		`,
	}

	runCmd = cli.Command{
		Name:     "run",
		Usage:    "running service",
		Category: "RUNNING COMMANDS",
		Description: `
		running commands can start,wait or stop the stack,
		while wait, we can greet the others
		`,

		Subcommands: []cli.Command{
			{
				Name:   "start",
				Usage:  "start the stack",
				Action: Start,
			},
			{
				Name:   "stop",
				Usage:  "stop the stack",
				Action: Stop,
			},
		},
	}
)

func defaultEnv() string {
	return "dev"
}
