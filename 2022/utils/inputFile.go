package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func GenInputFile(day int) string {
	var d string
	if day < 10 {
		d = fmt.Sprintf("0%d", day)
	} else {
		d = fmt.Sprintf("%d", day)
	}

	pwd, _ := os.Getwd()
	path := fmt.Sprintf("%s\\day%s\\input.txt", pwd, d)
	fi, _ := os.Stat(path)
	if fi != nil {
		return path
	}

	fmt.Println("Creating new input file...")
	f, _ := os.OpenFile(path, os.O_CREATE, 0700)
	f.WriteString(readHttp(2022, day))
	return path
}

func readHttp(year, day int) string {
	fmt.Println("Fetching data into file...")

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	session := os.Getenv("sessionAoC")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session})
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}

func ReadInputFile(path string) []string {
	fmt.Println("Reading from input file...", path)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	arr := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		arr = append(arr, sc.Text())
	}

	return arr
}
