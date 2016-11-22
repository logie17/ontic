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

	backupCommand = cli.Command{
		Name:   "backup",
		Usage:  "Backs dot files to the directory",
		Action: backup,
	}

	initCommand = cli.Command{
		Name:   "init",
		Usage:  "Initiales ontic with all configuration files/directories",
		Action: initOntic,
	}

	homeDir = os.Getenv("HOME")
	rootDir = homeDir + "/.ontic"
)

func main() {
	app := cli.NewApp()

	app.Name = "ontic"
	app.Usage = "install update"
	//app.Version = build.Info().String()
	app.Commands = []cli.Command{
		installCommand,
		updateCommand,
		backupCommand,
		initCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
