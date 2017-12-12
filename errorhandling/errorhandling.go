package errorhandling

import (
	"bufio"
	"os"
)

// CountLines count lines in the file
func CountLines(path string) (int, error) {
	var lines int
	file, err := os.Open(path)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines++
	}
	return lines, err
}
