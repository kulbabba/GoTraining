package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// CountLines count lines in the reader
func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	fmt.Println(lines)
	return lines, sc.Err()
	//return 0, nil
}

// CountFile write count of lines in the file to the stdout

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
	fmt.Println(path, lines)
	return lines
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		CountFile(arg)
	}
}
