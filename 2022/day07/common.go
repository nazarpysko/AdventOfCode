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
func (f *fileSystem) cd(dirName string) {
	if dirName == ".." {
		if strings.Count(f.path, "/") == 1 {
			f.path = "/"
		} else {
			f.path = f.path[:strings.LastIndex(f.path, "/")]
		}
		return
	} else if dirName == "/" {
		f.path = "/"
	} else {
		separator := "/"
		if f.path == "/" {
			separator = ""
		}
		f.path += separator + dirName
	}

	_, exists := f.directories[f.path]
	if !exists {
		f.directories[f.path] = 0
	}
}

// add sums the size to all directories from path
func (f *fileSystem) add(size int) {
	currentPath := f.path
	for {
		if currentPath == "" || currentPath == "/" {
			f.directories["/"] += size
			break
		} else {
			f.directories[currentPath] += size
			currentPath = currentPath[:strings.LastIndex(currentPath, "/")]
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
