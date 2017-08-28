package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"fmt"
	"log"
	"sort"
	"io/ioutil"
	"io"
	"time"
	"flag"
	"errors"
)

func SimpleCli() {
	demo10()
}

func demo1() {
	app := cli.NewApp()
	app.Run(os.Args)
}

func demo2() {
	app := cli.NewApp()
	app.Name = "loopring"
	app.Usage = "DIAOZHATIAN DE ICO"
	app.Authors = []cli.Author{cli.Author{Name:"Dylen Fu",Email:"dylenfu@126.com"}}
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

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang",
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

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "config, c",
			Usage: "load configuration from `FILE`",
		},
	}

	app.Run(os.Args)
}

func demo5() {
	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang, l",
			Value: "english",
			Usage: "language for the greeting",
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

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "lang, l",
			Value: "english",
			Usage: "language for greeting",
		},
		cli.StringFlag{
			Name: "config, c",
			Usage: "Load configuration from File",
		},
	}

	app.Commands = []cli.Command {
		{
			Name: "complete",
			Aliases: []string{"c"},
			Usage: "complete a task in list",
			Action: func(c *cli.Context) error {
				log.Println("cli test\t-", "commands", "complete")
				return nil
			},
		},

		{
			Name: "add",
			Aliases: []string {"a"},
			Usage: "add a task to a list",
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
			Name: "add",
			Category: "template",
			Usage: "add task in list",
			Action: func(c *cli.Context) {
				log.Println("cli test\t-", "sub commands", "add")
			},
		},
		{
			Name: "remove",
			Category: "template",
			Usage: "remove a task from list",
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
			Name: "complete",
			Aliases: []string{"c"},
			Usage: "complete a task in list",
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

				for _,t := range tasks {
					log.Println("cli test\t-", "bash completion", t)
				}
			},
		},
	}

	app.Run(os.Args)
}

func demo10() {
	initdemo10()
	app := cli.NewApp()
	app.Name = "kənˈtrīv"
	app.Version = "19.99.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Example Human",
			Email: "human@example.com",
		},
	}
	app.Copyright = "(c) 1999 Serious Enterprise"
	app.HelpName = "contrive"
	app.Usage = "demonstrate available API"
	app.UsageText = "contrive - demonstrating the available API"
	app.ArgsUsage = "[args and such]"
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "doo",
			Aliases:     []string{"do"},
			Category:    "motion",
			Usage:       "do the doo",
			UsageText:   "doo - does the dooing",
			Description: "no really, there is a lot of dooing to be done",
			ArgsUsage:   "[arrgh]",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "forever, forevvarr"},
			},
			Subcommands: cli.Commands{
				cli.Command{
					Name:   "wop",
					Action: wopAction,
				},
			},
			SkipFlagParsing: false,
			HideHelp:        false,
			Hidden:          false,
			HelpName:        "doo!",
			BashComplete: func(c *cli.Context) {
				fmt.Fprintf(c.App.Writer, "--better\n")
			},
			Before: func(c *cli.Context) error {
				fmt.Fprintf(c.App.Writer, "brace for impact\n")
				return nil
			},
			After: func(c *cli.Context) error {
				fmt.Fprintf(c.App.Writer, "did we lose anyone?\n")
				return nil
			},
			Action: func(c *cli.Context) error {
				c.Command.FullName()
				c.Command.HasName("wop")
				c.Command.Names()
				c.Command.VisibleFlags()
				fmt.Fprintf(c.App.Writer, "dodododododoodododddooooododododooo\n")
				if c.Bool("forever") {
					c.Command.Run(c)
				}
				return nil
			},
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, "for shame\n")
				return err
			},
		},
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{Name: "fancy"},
		cli.BoolTFlag{Name: "fancier"},
		cli.DurationFlag{Name: "howlong, H", Value: time.Second * 3},
		cli.Float64Flag{Name: "howmuch"},
		cli.GenericFlag{Name: "wat", Value: &genericType{}},
		cli.Int64Flag{Name: "longdistance"},
		cli.Int64SliceFlag{Name: "intervals"},
		cli.IntFlag{Name: "distance"},
		cli.IntSliceFlag{Name: "times"},
		cli.StringFlag{Name: "dance-move, d"},
		cli.StringSliceFlag{Name: "names, N"},
		cli.UintFlag{Name: "age"},
		cli.Uint64Flag{Name: "bigage"},
	}
	app.EnableBashCompletion = true
	app.HideHelp = false
	app.HideVersion = false
	app.BashComplete = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "lipstick\nkiss\nme\nlipstick\nringo\n")
	}
	app.Before = func(c *cli.Context) error {
		fmt.Fprintf(c.App.Writer, "HEEEERE GOES\n")
		return nil
	}
	app.After = func(c *cli.Context) error {
		fmt.Fprintf(c.App.Writer, "Phew!\n")
		return nil
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Thar be no %q here.\n", command)
	}
	app.OnUsageError = func(c *cli.Context, err error, isSubcommand bool) error {
		if isSubcommand {
			return err
		}

		fmt.Fprintf(c.App.Writer, "WRONG: %#v\n", err)
		return nil
	}
	app.Action = func(c *cli.Context) error {
		cli.DefaultAppComplete(c)
		cli.HandleExitCoder(errors.New("not an exit coder, though"))
		cli.ShowAppHelp(c)
		cli.ShowCommandCompletions(c, "nope")
		cli.ShowCommandHelp(c, "also-nope")
		cli.ShowCompletions(c)
		cli.ShowSubcommandHelp(c)
		cli.ShowVersion(c)

		categories := c.App.Categories()
		categories.AddCommand("sounds", cli.Command{
			Name: "bloop",
		})

		for _, category := range c.App.Categories() {
			fmt.Fprintf(c.App.Writer, "%s\n", category.Name)
			fmt.Fprintf(c.App.Writer, "%#v\n", category.Commands)
			fmt.Fprintf(c.App.Writer, "%#v\n", category.VisibleCommands())
		}

		fmt.Printf("%#v\n", c.App.Command("doo"))
		if c.Bool("infinite") {
			c.App.Run([]string{"app", "doo", "wop"})
		}

		if c.Bool("forevar") {
			c.App.RunAsSubcommand(c)
		}
		c.App.Setup()
		fmt.Printf("%#v\n", c.App.VisibleCategories())
		fmt.Printf("%#v\n", c.App.VisibleCommands())
		fmt.Printf("%#v\n", c.App.VisibleFlags())

		fmt.Printf("%#v\n", c.Args().First())
		if len(c.Args()) > 0 {
			fmt.Printf("%#v\n", c.Args()[1])
		}
		fmt.Printf("%#v\n", c.Args().Present())
		fmt.Printf("%#v\n", c.Args().Tail())

		set := flag.NewFlagSet("contrive", 0)
		nc := cli.NewContext(c.App, set, c)

		fmt.Printf("%#v\n", nc.Args())
		fmt.Printf("%#v\n", nc.Bool("nope"))
		fmt.Printf("%#v\n", nc.BoolT("nerp"))
		fmt.Printf("%#v\n", nc.Duration("howlong"))
		fmt.Printf("%#v\n", nc.Float64("hay"))
		fmt.Printf("%#v\n", nc.Generic("bloop"))
		fmt.Printf("%#v\n", nc.Int64("bonk"))
		fmt.Printf("%#v\n", nc.Int64Slice("burnks"))
		fmt.Printf("%#v\n", nc.Int("bips"))
		fmt.Printf("%#v\n", nc.IntSlice("blups"))
		fmt.Printf("%#v\n", nc.String("snurt"))
		fmt.Printf("%#v\n", nc.StringSlice("snurkles"))
		fmt.Printf("%#v\n", nc.Uint("flub"))
		fmt.Printf("%#v\n", nc.Uint64("florb"))
		fmt.Printf("%#v\n", nc.GlobalBool("global-nope"))
		fmt.Printf("%#v\n", nc.GlobalBoolT("global-nerp"))
		fmt.Printf("%#v\n", nc.GlobalDuration("global-howlong"))
		fmt.Printf("%#v\n", nc.GlobalFloat64("global-hay"))
		fmt.Printf("%#v\n", nc.GlobalGeneric("global-bloop"))
		fmt.Printf("%#v\n", nc.GlobalInt("global-bips"))
		fmt.Printf("%#v\n", nc.GlobalIntSlice("global-blups"))
		fmt.Printf("%#v\n", nc.GlobalString("global-snurt"))
		fmt.Printf("%#v\n", nc.GlobalStringSlice("global-snurkles"))

		fmt.Printf("%#v\n", nc.FlagNames())
		fmt.Printf("%#v\n", nc.GlobalFlagNames())
		fmt.Printf("%#v\n", nc.GlobalIsSet("wat"))
		fmt.Printf("%#v\n", nc.GlobalSet("wat", "nope"))
		fmt.Printf("%#v\n", nc.NArg())
		fmt.Printf("%#v\n", nc.NumFlags())
		fmt.Printf("%#v\n", nc.Parent())

		nc.Set("wat", "also-nope")

		ec := cli.NewExitError("ohwell", 86)
		fmt.Fprintf(c.App.Writer, "%d", ec.ExitCode())
		fmt.Printf("made it!\n")
		return ec
	}

	if os.Getenv("HEXY") != "" {
		app.Writer = &hexWriter{}
		app.ErrWriter = &hexWriter{}
	}

	app.Metadata = map[string]interface{}{
		"layers":     "many",
		"explicable": false,
		"whatever-values": 19.99,
	}

	app.Run(os.Args)
}

func wopAction(c *cli.Context) error {
	fmt.Fprintf(c.App.Writer, ":wave: over here, eh\n")
	return nil
}

func initdemo10() {
	cli.AppHelpTemplate += "\nCUSTOMIZED: you bet ur muffins\n"
	cli.CommandHelpTemplate += "\nYMMV\n"
	cli.SubcommandHelpTemplate += "\nor something\n"

	cli.HelpFlag = cli.BoolFlag{Name: "halp"}
	cli.BashCompletionFlag = cli.BoolFlag{Name: "compgen", Hidden: true}
	cli.VersionFlag = cli.BoolFlag{Name: "print-version, V"}

	cli.HelpPrinter = func(w io.Writer, templ string, data interface{}) {
		fmt.Fprintf(w, "best of luck to you\n")
	}
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "version=%s\n", c.App.Version)
	}
	cli.OsExiter = func(c int) {
		fmt.Fprintf(cli.ErrWriter, "refusing to exit %d\n", c)
	}
	cli.ErrWriter = ioutil.Discard
	cli.FlagStringer = func(fl cli.Flag) string {
		return fmt.Sprintf("\t\t%s", fl.GetName())
	}
}

type hexWriter struct{}

func (w *hexWriter) Write(p []byte) (int, error) {
	for _, b := range p {
		fmt.Printf("%x", b)
	}
	fmt.Printf("\n")

	return len(p), nil
}

type genericType struct{
	s string
}

func (g *genericType) Set(value string) error {
	g.s = value
	return nil
}

func (g *genericType) String() string {
	return g.s
}