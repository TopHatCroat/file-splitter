package main

import (
	"flag"
	"fmt"
	"github.com/bmatcuk/doublestar/v4"
	"io/fs"
	"os"
	"strings"
)

var glob = flag.String("glob", "*", "Glob pattern for files to split")
var index = flag.Int("index", 0, "Index of current split")
var total = flag.Int("total", 2, "Total number of splits")
var help = flag.Bool("help", false, "Show help")

func detectFiles(pattern string, index, total int) []string {
	var filePaths []string
	var count = 0

	fileSystem := os.DirFS(".")
	err := doublestar.GlobWalk(fileSystem, pattern, func(path string, d fs.DirEntry) error {
		if d.IsDir() {
			return nil
		}

		if count%total == index {
			filePaths = append(filePaths, path)
		}

		count = count + 1

		return nil
	})

	if err != nil {
		fmt.Printf("Failed to load files with pattern: %v\n", err)
		os.Exit(1)
	}

	return filePaths
}

func main() {
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *index < 0 || *index >= *total {
		fmt.Println("Index must be higher than 0 and less then total")
		os.Exit(1)
	}

	fmt.Println(*glob)
	fmt.Println(*index)
	fmt.Println(*total)

	files := detectFiles(*glob, *index, *total)

	fmt.Println(strings.Join(files, " "))
}
