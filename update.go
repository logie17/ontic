package main

import (
	"fmt"
	cli "gopkg.in/urfave/cli.v1"
	"time"
)

func update(c *cli.Context) error {
	spinner := NewSpinner("Updating...").Start()
	time.Sleep(3 * time.Second)
	spinner.Stop()

	return nil
}
