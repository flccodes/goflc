package work

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// FreadCSV opens and reads the content from a csv file into a [][]string
func FreadCSV(csvFile string) [][]string {
	f, err := os.Open(csvFile)
	if err != nil {
		log.Printf("Error while opning the CSV file: %s\n", csvFile)
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
