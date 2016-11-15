package main

import (
	"fmt"
	cli "gopkg.in/urfave/cli.v1"
	"log"
	"os"
	"os/exec"
	"time"
)

func backup(c *cli.Context) error {
	spinner := NewSpinner("Backing up dot files...").Start()
	t := time.Now()
	timestamp := fmt.Sprintf(
		"%04d%02d%02d-%02d%02d%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second(),
	)

	backUpDir := rootDir + "/backup/" + timestamp

	os.Mkdir(rootDir, 0700)
	os.Mkdir(backUpDir, 0700)
	os.Mkdir(rootDir+"/tmp/", 0700)
	os.Mkdir(rootDir+"/tmp/backup", 0700)

	backUpListFile := rootDir + "/tmp/backup/" + timestamp + "-backup-list"

	files := allFiles()

	f, err := os.Create(backUpListFile)
	if err != nil {
		panic(err)
	}

	count := 0
	for file := range files {
		f.WriteString(file + "\n")
		fmt.Printf("%v\n", file)
		count++
	}
	f.Close()

	cmd := fmt.Sprintf("cd %s; cat %s | cpio -dump %s", homeDir, backUpListFile, backUpDir)
	_, cmdErr := exec.Command("/bin/sh", "-c", cmd).Output()
	if cmdErr != nil {
		log.Fatal(cmdErr)
	}

	spinner.Stop()
	fmt.Printf("\033[31mBacked up %d dot files to %s\n", count, backUpDir)
	return nil
}
