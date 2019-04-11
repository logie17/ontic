package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func allPaths() []string {
	ReadConfig()

	repos := Config
	paths := []string{}
	for _, repo := range repos {
		path := repo["path"]
		path = rootDir + "/" + path
		if _, err := os.Stat(path); os.IsNotExist(err) {
			log.Fatalf("The path exists but is not a git repo [%s] %v", path, err)
		} else {
			if matched, err := regexp.MatchString("^/", path); err == nil && !matched {
				path = rootDir + "/" + path
			}
			paths = append(paths, path)
		}
	}
	return paths
}

func allFiles() map[string]string {
	allFiles := map[string]string{}
	paths := allPaths()

	for _, path := range paths {
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

				// On conflicts we keep the previously found dot file
				if _, ok := allFiles[file]; !ok {
					allFiles[file] = path
				}
			}
		}

	}

	return allFiles
}

func upToDate(src, dst string, method method) bool {
	srcFi, srcFiErr := os.Stat(src)
	dstFi, dstFiErr := os.Stat(dst)

	if os.IsNotExist(srcFiErr) {
		log.Panicf("The file does not exist! " + src)
	}

	if srcFiErr != nil {
		log.Printf("There was a problem statting src file %s %v ", src, srcFiErr)
		return false
	}

	if dstFiErr != nil {
		log.Printf("There was a problem statting dst file %s %v ", dst, dstFiErr)
		return false
	}

	if srcFi.Size() != dstFi.Size() {
		return false
	}

	if method == symlink {
		linkdst, _ := os.Readlink(dst)
		return linkdst == src
	}

	return true
}
