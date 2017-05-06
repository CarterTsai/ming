package main

import (
	"os"

	"github.com/urfave/cli"

	command "github.com/CarterTsai/ming/command/service"
	config "github.com/CarterTsai/ming/config"
)

func main() {

	app := cli.NewApp()
	app.Name = "ming"
	app.Usage = "ming toolkit"
	app.VisibleFlags()
	app.Version = "0.0.1"
	app.EnableBashCompletion = true

	cli.AppHelpTemplate = config.Usage()
	// cli.CommandHelpTemplate = config.CommandUsage()

	app.Commands = []cli.Command{
		{
			Name: "img2pdf",
			// Aliases:  []string{"i"},
			Usage:    "img to pdf",
			HelpName: "img",
			Subcommands: cli.Commands{
				command.Create(),
			},
			Action: func(c *cli.Context) error {
				// init variable
				cli.ShowSubcommandHelp(c)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
