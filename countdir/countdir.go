package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// CountLines - count the lines in from Reader
func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	fmt.Println(lines)
	return lines, sc.Err()
}

// CountFile write to stdout number of lines in the file
func CountFile(path string) int {
	var lines int
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
	}
	fmt.Println(lines)
	return lines
}

// CountDir - write to stdout numbers of lines for each files in directory
func CountDir(dir string) {

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), ".txt") {
			fmt.Println(f.Name())
			var buffer bytes.Buffer
			buffer.WriteString(dir)
			buffer.WriteString(f.Name())

			CountFile(buffer.String())
		}
	}
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		CountDir(arg)
	}
}
