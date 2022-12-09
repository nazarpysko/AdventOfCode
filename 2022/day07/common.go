package main

import (
	"strconv"
	"strings"
	"unicode"
)

type directories map[string]int

type fileSystem struct {
	path        string
	directories directories
}

// cd adds new path and returns the current working directory
func (f *fileSystem) cd(newDirectoryName string) {
	nDirectoriesInPath := strings.Count(f.path, "/")
	if newDirectoryName == ".." {
		if nDirectoriesInPath == 2 {
			f.path = "/"
		} else {
			f.path = f.path[:strings.LastIndex(f.path, "/")]
		}
		return
	} else if newDirectoryName == "/" {
		f.path = "/"
	} else {
		separator := ""
		if nDirectoriesInPath >= 1 {
			separator = "/"
		}
		f.path += separator + newDirectoryName
	}

	_, exists := f.directories[newDirectoryName]
	if !exists {
		f.directories[newDirectoryName] = 0
	}
}

// add sums the size passed to the current and the parents directory as well
func (f *fileSystem) add(size int) {
	for _, d := range append([]string{"/"}, strings.Split(f.path, "/")...) {
		if d != "" {
			f.directories[d] += size
		}
	}
}

func isCommand(line string) bool {
	return string(line[0]) == "$"
}

func isFile(line string) bool {
	splitLine := strings.Split(line, " ")
	return unicode.IsDigit(rune(splitLine[0][0]))
}

func getBytesFromLine(line string) int {
	splitLine := strings.Split(line, " ")
	bytes, _ := strconv.Atoi(splitLine[0])
	return bytes
}

func getSolution(d directories) int {
	totalSize := 0
	for _, v := range d {
		if v <= 100000 {
			totalSize += v
		}
	}

	return totalSize
}
