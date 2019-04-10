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
		log.Panicf("Unable to create backup list file %s %v", backUpListFile, err)
	}

	count := 0
	for file := range files {
		_, sErr := os.Stat(homeDir + "/" + file)
		if os.IsNotExist(sErr) {
			continue
		}

		if sErr != nil {
			fmt.Printf("Unexpected error %v\n", sErr)
			continue
		}

		f.WriteString(file + "\n")
		count++
	}
	f.Close()

	cmd := fmt.Sprintf("cd %s; cat %s | cpio -dump %s", homeDir, backUpListFile, backUpDir)
	_, cmdErr := exec.Command("/bin/sh", "-c", cmd).Output()
	if cmdErr != nil {
		log.Panicf("Unable to backup files with cmd: %s %v", cmd, cmdErr)
	}

	spinner.Stop()
	fmt.Printf("\033[31mBacked up %d dot files to %s\033[m\n", count, backUpDir)
	return nil
}
