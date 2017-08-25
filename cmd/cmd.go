package cmd

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"fmt"
)

func SimpleCli() {
	demo2()
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