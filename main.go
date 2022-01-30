package main

import (
	"flag"
	"fmt"
	"github.com/bmatcuk/doublestar/v4"
	"io/fs"
	"os"
	"strings"
)

var glob string
var index int
var total int
var help bool

func parseFlags() {
	flag.StringVar(&glob, "glob", "*", "Glob pattern for files to split")

	flag.IntVar(&index, "index", 0, "Index of current split")
	flag.IntVar(&total, "total", 2, "Total number of splits")

	flag.BoolVar(&help, "help", false, "Show help")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

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
	parseFlags()

	files := detectFiles(glob, index, total)

	fmt.Println(strings.Join(files, " "))
}
