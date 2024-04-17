package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

var results []result

func main() {
	var dir string

	args := os.Args
	if len(args) >= 2 {
		dir = args[1]
	} else {
		usr, _ := user.Current()
		dir = usr.HomeDir + "/Developer"
	}

	iterateDirs(dir)

	output := formatJSON()

	os.Stdout.WriteString(output)
}

// iterate over directories and their subdirectories
func iterateDirs(dir string) {
	subDirs, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, subDir := range subDirs {
		// if file is not a directory, skip to the next one
		if !subDir.IsDir() {
			continue
		}

		abs, err := filepath.Abs(dir + "/" + subDir.Name())
		if err != nil {
			log.Fatal(err)
		}

		// check directory and call function recursively if check fails
		if checkDir(abs) == true {
			results = append(results, result{
				id:    uuid.NewString(),
				title: subDir.Name(),
				path:  abs,
			})
		} else {
			iterateDirs(abs)
		}

	}
}

// check if input directory contains a git repository
func checkDir(input string) bool {
	dirs, err := os.ReadDir(input)
	if err != nil {
		log.Fatal(err)
	}

	for _, dir := range dirs {
		if dir.Name() == ".git" {
			return true
		}
	}

	return false
}

// format JSON string for Alfred workflow
func formatJSON() string {
	var items []string

	for _, result := range results {
		item := "{\"uid\": \"" + result.id + "\", \"type\": \"file\", \"title\": \"" + result.title + "\", \"arg\": \"" + result.path + "\" }"

		items = append(items, item)
	}

	json := "{\"items\": [" + strings.Join(items, ", ") + "]}"

	return json
}

type result struct {
	id    string
	title string
	path  string
}
