package work

import (
	"bytes"
	"fmt"
	"testing"
)

// TestFreadCSV created / adaptated from:
// https://stackoverflow.com/a/56398447
func TestFreadCSV(t *testing.T) {
	var buffer bytes.Buffer
	buffer.WriteString("fake, csv, data")
	content, err := FreadFile(&buffer)
	if err != nil {
		t.Error("Failed to read csv data")
	}
	fmt.Print(content)

}
