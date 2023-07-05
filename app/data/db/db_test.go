package db_test

import (
	"dictionary-service/app/data/db"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestFilePath = "./mock_data.txt"
)

func TestReadingData(t *testing.T) {
	setup()
	db := db.NewDb()
	err := db.LoadData(TestFilePath)
	assert.Len(t, db.GetData(), 6)
	assert.Equal(t, err, nil)
	destroy()
}

func setup() {
	file, err := os.OpenFile(TestFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Write words to the file
	words := []string{"A", "Song", "Of", "Ice", "And", "Fire"}

	for _, word := range words {
		_, err := file.WriteString(word + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func destroy() {
	os.Remove(TestFilePath)
}
