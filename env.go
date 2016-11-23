package main

import (
	"fmt"
	cli "gopkg.in/urfave/cli.v1"
	"strings"
)

func env(c *cli.Context) error {
	fmt.Printf("export DOTDOTDOT_ROOT='%s';\n", rootDir)
	fmt.Printf("export DOTDOTDOT_ORDER='%s';\n", strings.Join(allPaths(), " "))

	return nil
}
