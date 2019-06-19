package main

import (
	"fmt"
	"os"

	"gopkg.in/urfave/cli.v2"

	"github.com/wrfly/golang-template/config"
)

var appName = "template-app"

func main() {
	app := &cli.App{
		Name:        appName,
		Usage:       "Some template application",
		Authors:     author,
		Version:     simpleVersionInfo,
		HideVersion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Usage:   "config file path",
				Aliases: []string{"c"},
				Value:   "./config.yml",
			},
			&cli.BoolFlag{
				Name:    "example",
				Usage:   "config file example",
				Aliases: []string{"e"},
				Value:   false,
			},
			&cli.BoolFlag{
				Name:  "env-list",
				Usage: "config environment lists",
				Value: false,
			},
			&cli.BoolFlag{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "print the version",
			},
		},
		CustomAppHelpTemplate: helpTemplate,
		Action: func(c *cli.Context) error {
			if c.NumFlags() == 0 {
				return cli.ShowAppHelp(c)
			}
			switch {
			case c.Bool("version"):
				fmt.Println(versionInfo)
				return nil
			case c.Bool("env-list"):
				config.EnvList()
				return nil
			case c.Bool("example"):
				if err := config.Example(); err != nil {
					panic(err)
				}
				return nil
			}

			// main logic here
			cfg, err := config.New(c.String("config"))
			if err != nil {
				panic(err)
			}
			fmt.Println("got port", cfg.Port)

			return nil
		},
	}

	app.Run(os.Args)
}
