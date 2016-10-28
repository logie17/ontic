package main

import (
	"fmt"
	cli "gopkg.in/urfave/cli.v1"
	"time"
)

func backup(c *cli.Context) error {
	spinner := NewSpinner("Backing up dot files...").Start()
	t := time.Now()
	timestamp := fmt.Sprintf("%04d%02d%02d-%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	backUpDir := rootDir + "/backup/" + timestamp
	backUpListFile := rootDir + "/tmp/backup/" + timestamp + "-backup-list"

	spinner.Stop()

	fmt.Println(backUpDir)
	fmt.Println(backUpListFile)

	return nil

}
