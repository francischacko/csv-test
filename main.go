package main

import (
	"os"

	"github.com/gocarina/gocsv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Entry struct {
	gorm.Model
	Name    string `csv:"name"`
	Email   string `csv:"email"`
	Gender  string `csv:"gender"`
	Address string `csv:"address"`
	City    string `csv:"city"`
}

func main() {
	// Open the CSV file for reading
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Parse CSV into a slice of typed data `[]Entry` (just like json.Unmarshal() does)
	// The builtin package `encoding/csv` does not support unmarshaling into a struct
	// so you need to use an external library to avoid writing for-loops.
	var entries []Entry
	err = gocsv.Unmarshal(file, &entries)
	if err != nil {
		panic(err)
	}

	// Open a postgres database connection using GORM
	db, err := gorm.Open(postgres.Open("host=localhost user=edit password=edit dbname=edit port=5432 sslmode=disable "))
	if err != nil {
		panic(err)
	}

	// Create `entries` table if not exists
	err = db.AutoMigrate(&Entry{})
	if err != nil {
		panic(err)
	}

	// Save all the records at once in the database
	result := db.Create(entries)
	if result.Error != nil {
		panic(result.Error)
	}
}
