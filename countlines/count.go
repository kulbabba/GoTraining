package count

import (
	"bufio"
	"log"
	"os"
)

// ContLines count lines of content in the file
func CountLines(path string) int {

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
	return lines
}
