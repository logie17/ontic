package main

import (
	cli "gopkg.in/urfave/cli.v1"
	"time"
)

func install(c *cli.Context) error {
	spinner := NewSpinner("Installing...").Start()
	time.Sleep(3 * time.Second)
	spinner.Stop()

	return nil
}
