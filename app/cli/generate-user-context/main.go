package main

import (
	"errors"
	valid "github.com/asaskevich/govalidator"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:      "setup",
			Usage:     "enable auto complete for bash, zsh or PowerShell",
			ArgsUsage: "[shell type (bash|zsh|powershell)]",
			Action:    setup,
		},
		{
			Name:         "teardown",
			Aliases:      []string{"td"},
			Usage:        "disable auto complete for bash, zsh or PowerShell",
			ArgsUsage:    "[shell type (bash|zsh|powershell)]",
			Action:       teardown,
			BashComplete: teardownComplete,
		},
		{
			Name:      "create",
			Aliases:   []string{"c"},
			Usage:     "create a namespace and user account for that user",
			ArgsUsage: "[user-name]",
			Action: func(context *cli.Context) error {
				name := context.Args().Get(0)
				if valid.IsDNSName(name) {
					return errors.New("please use a dns-friendly name")
				}
				action, err := HelmInit()
				if err != nil {
					log.Println("failed to initialize helm")
					return err
				}
				list, err := List(action)
				if err != nil {
					log.Println("failed to list actions")
					return err
				}
				exist := ExistInHelm(name, list)
				if exist {
					log.Println("namespace and account already exist")
					return errors.New("namespace and account already exist")
				}
				return nil
			},
		},
	}
	app.UseShortOptionHandling = true
	app.EnableBashCompletion = true
	app.Name = "Atomi Ops Kit"
	app.Description = "Allow for AtomiCloud's ops tooling"
	app.Version = "0.0.1"
	app.Usage = "Operations"
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		{
			Name:  "Kirinnee",
			Email: "kirinnee97@gmail.com",
		},
	}

	cli.AppHelpTemplate = `{{.Name}} - {{.Usage}}
VERSION: {{.Version}}
USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}

   {{end}}
`

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
