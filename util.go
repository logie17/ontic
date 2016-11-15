package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func allFiles() map[string]string {
	repos := Config
	allFiles := map[string]string{}
	for _, repo := range repos {
		path := repo["path"]
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Fatalf("The path exists but is not a git repo %v", err)
		} else {
			cmd := fmt.Sprintf("cd %s; find -L . -type f", path)
			output, err := exec.Command("/bin/sh", "-c", cmd).Output()
			if err != nil {
				log.Fatalf("Unable to run the %s %v", cmd, err)
			}

			files := strings.Split(string(output), "\n")
			re := regexp.MustCompile("^./")

			for _, file := range files {
				file = re.ReplaceAllString(file, "")
				if matched, err := regexp.MatchString("^(?:\\.|bin/)", file); err == nil && matched {
					if matched, err := regexp.MatchString("(\\.sw.|~)$", file); err != nil || matched {
						continue
					}
					if matched, err := regexp.MatchString("\\.git/", file); err != nil || matched {
						continue
					}
					if matched, err := regexp.MatchString("\\.(git|gitignore|gitmodules?)$", file); err != nil || matched {
						continue
					}
					if matched, err := regexp.MatchString("^\\.\\.\\.(Makefile|deps)", file); err != nil || matched {
						continue
					}

					allFiles[file] = path
				}
			}

		}

	}

	return allFiles
}
