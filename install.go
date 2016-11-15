package main

import (
	"fmt"
	cli "gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"regexp"
)

var defaultMethod = "symlink"

func install(c *cli.Context) error {
	spinner := NewSpinner("Installing dot files...").Start()

	backup(c)

	files := allFiles()
	method := defaultMethod

	count := 0
	skipCount := 0
	re := regexp.MustCompile("(\\.*/)\\.*")

	for file, path := range files {
		src := path + "/" + file
		dst := homeDir + "/" + file
		if upToDate(src, dst, method) {
			skipCount++
			continue
		}
		count++
		dir := re.ReplaceAllString(dst, "$1")
		if _, statErr := os.Stat(dir); statErr != nil {
			os.MkdirAll(dir, os.ModePerm)
		}

		fmt.Printf("Removing %v\n", dst)
		os.Remove(dst)

		if method == "symlink" {
			if err := os.Symlink(src, dst); err != nil {
				log.Panicf("Unable create symlink! %v\n", err)
			}
		} else {
			panic("Invalid method")
		}

	}

	spinner.Stop()

	return nil
}
