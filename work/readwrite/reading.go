package work

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// FreadCSV opens and reads the content from a csv file into a [][]string
func FreadCSV(csvFile string) [][]string {
	f, err := os.Open(csvFile)
	if err != nil {
		log.Printf("Error while opening the CSV file: %s\n", csvFile)
	}
	defer f.Close()

	data, err := FreadFile(f)
	if err != nil {
		log.Printf("Error while reading the CSV file: %s\n", csvFile)
	}
	return data
}

// FreadFile helps the function FreadCSV() - Created / adaptated from:
// https://stackoverflow.com/a/56398447
func FreadFile(reader io.Reader) ([][]string, error) {
	r := csv.NewReader(reader)
	// r.Comma = ';' //
	data, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return data, err
}

// FreadFileLineByLine reads a file line by line and returns its content into a slice of strings
func FreadFileLineByLine(s string) []string {
	f, err := os.Open(s)
	if err != nil {
		fmt.Printf("Error while opening the file: %s\n", err)
	}

	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while reading the file:", err)
	}

	return lines
}

// FreadingDir retrieves a list of all the files and directories in a directory
func FreadingDir(s string) {
	dirs, err := os.ReadDir(s)
	if err != nil {
		fmt.Println(err)
	}

	for _, entry := range dirs {
		fmt.Println(entry.Name())
	}
}
