package main

import (
	"gopkg.in/urfave/cli.v1"
	"log"
	"os"
)

var (
	installCommand = cli.Command{
		Name:   "install",
		Usage:  "Gets the latest dot files and installs them",
		Action: install,
	}

	updateCommand = cli.Command{
		Name:   "update",
		Usage:  "Downloads all dot files form repos specified in conf file",
		Action: update,
	}
)

func main() {
	ReadConfig()
	app := cli.NewApp()

	app.Name = "ontic"
	app.Usage = "install update"
	//app.Version = build.Info().String()
	app.Commands = []cli.Command{installCommand, updateCommand}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
