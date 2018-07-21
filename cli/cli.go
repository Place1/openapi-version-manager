package cli

import (
	"sort"

	"github.com/urfave/cli"

	"openapi-version-manager/commands"
)

func Run(args []string) error {
	app := cli.NewApp()
	app.Name = "openapi-version-manager"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		{
			Name:  "current",
			Usage: "show the current openapi generator version",
			Action: func(context *cli.Context) error {
				return commands.Current()
			},
		},
		{
			Name:  "list",
			Usage: "list available openapi generator versions",
			Action: func(context *cli.Context) error {
				return commands.List()
			},
		},
		{
			Name:      "use",
			Usage:     "use the specified openapi-generator-version",
			ArgsUsage: "<version>",
			Action: func(context *cli.Context) error {
				version := context.Args().First()
				if version == "" {
					return cli.NewExitError("please provide version string", 1)
				}
				return commands.Use(context.Args().First())
			},
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	return app.Run(args)
}
