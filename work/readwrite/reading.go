package work

import (
	"encoding/csv"
	"log"
	"os"
)

// FreadCSV opens and reads the content from a csv file into a [][]string
func FreadCSV(csvFile string) [][]string {
	f, err := os.Open(csvFile)
	if err != nil {
		log.Printf("Error!%s\n", err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	reader, err := r.ReadAll()
	if err != nil {
		log.Printf("Error!%s\n", err)
	}
	return reader
}
