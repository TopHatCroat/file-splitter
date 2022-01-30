package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestFiles(filePaths []string) {
	var testFilePaths []string
	for _, path := range filePaths {
		testFilePaths = append(testFilePaths, path)
	}

	for _, filePath := range testFilePaths {
		// remove "./", we don't care for that
		var cleanFilePath = strings.Replace(filePath, "./", "", 1)
		dirs := strings.Split(cleanFilePath, "/")

		dirPath := strings.Join(dirs[:len(dirs)-1], "/")
		_, err := os.Stat(dirPath)
		if os.IsNotExist(err) {
			err := os.MkdirAll(dirPath, 0777)
			if err != nil {
				fmt.Printf("Error creating directory path: %s", err)
			}
		}

		file, err := os.Create(cleanFilePath)

		if err != nil {
			fmt.Println(err)
		}

		defer func(file *os.File) {
			_ = file.Close()
		}(file)
	}
}

func clearTestFiles() {
	err := os.RemoveAll("__mock__")
	if err != nil {
		fmt.Println(err)
	}
}

func TestDetectSingleFile(t *testing.T) {
	var singleFile = "__mock__/test.txt"

	createTestFiles([]string{singleFile})

	files := detectFiles(singleFile, 0, 1)

	assert.Equal(t, []string{singleFile}, files)

	clearTestFiles()
}

func TestDetect2LevelFiles(t *testing.T) {
	var files = []string{"__mock__/lvl2/test.txt", "__mock__/lvl2/test2.txt"}

	createTestFiles(files)

	actual := detectFiles("__mock__/**", 0, 1)

	assert.Equal(t, files, actual)

	clearTestFiles()
}

func TestDetectFilesWithIndex(t *testing.T) {
	var files = []string{"__mock__/test.txt", "__mock__/test2.txt"}

	createTestFiles(files)

	actual := detectFiles("__mock__/**", 0, 2)

	assert.Equal(t, []string{"__mock__/test.txt"}, actual)

	clearTestFiles()
}

func TestDetectFilesWithIndexEvenAndOdd(t *testing.T) {
	var files = []string{"__mock__/test.txt", "__mock__/test2.txt"}

	createTestFiles(files)

	actual := detectFiles("__mock__/**", 0, 2)

	assert.Equal(t, []string{"__mock__/test.txt"}, actual)

	actual2 := detectFiles("__mock__/**", 1, 2)

	assert.Equal(t, []string{"__mock__/test2.txt"}, actual2)

	clearTestFiles()
}
