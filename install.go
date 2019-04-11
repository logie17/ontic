package main

import (
	"fmt"
	cli "gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type method int
const (
	hardlink method = iota
	symlink
)

var defaultMethod = hardlink

func install(c *cli.Context) error {
	spinner := NewSpinner("Installing dot files...").Start()

	backup(c)

	files := allFiles()
	method := defaultMethod

	count := 0
	skipCount := 0

	for file, path := range files {
		src := path + "/" + file
		dst := homeDir + "/" + file
		if upToDate(src, dst, method) {
			skipCount++
			continue
		}
		count++
		dir := filepath.Dir(dst)
		if _, statErr := os.Stat(dir); statErr != nil {
			os.MkdirAll(dir, os.ModePerm)
		}

		os.Remove(dst)

		if method == hardlink {
			cmd := fmt.Sprintf("cp -f %s %s", src, dst)
			_, err := exec.Command("/bin/sh", "-c", cmd).Output()
			if err != nil {
				log.Fatalf("Unable to run the %s %v", cmd, err)
			}
		} else if method == symlink {
			if err := os.Symlink(src, dst); err != nil {
				log.Panicf("Unable create symlink! %v\n", err)
			}
		} else {
			panic("Invalid method")
		}

	}

	if method == symlink {
		watchDirs := []string{"bin", ".zsh", ".bash", ".vim", ".sh"}
		existingDirs := []string{}

		for _, dir := range watchDirs {
			fullDir := homeDir + "/" + dir
			if s, statErr := os.Stat(fullDir); statErr == nil && s.IsDir() {
				existingDirs = append(existingDirs, fullDir)
			}

		}

		badSyms := findDeadSyms(existingDirs)
		if len(badSyms) == 0 {
			fmt.Println("No bad symlinks found")
		} else {
			for _, badSym := range badSyms {
				// Check err here?
				os.Remove(badSym)
				fmt.Printf("Removing bad link %s\n", badSym)
			}
		}

	}

	spinner.Stop()

	return nil
}

// Todo write a unit test for this
func findDeadSyms(dirs []string) []string {
	reaped := []string{}
	for _, dir := range dirs {
		files, _ := ioutil.ReadDir(dir)
		for _, file := range files {
			path := dir + "/" + file.Name()
			fi, _ := os.Lstat(path)
			s, sErr := os.Stat(path)
			if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
				if os.IsNotExist(sErr) {
					reaped = append(reaped, path)
				}
			} else if sErr == nil && s.IsDir() {
				reaped = append(reaped, findDeadSyms([]string{path})...)
			}
		}
	}
	return reaped
}
