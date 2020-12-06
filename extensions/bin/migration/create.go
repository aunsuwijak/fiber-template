package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"time"
)

func main() {
	// Execute ./create [filename] in command line.
	Arguments := os.Args[1:]

	if len(Arguments) < 1 {
		fmt.Println("Please specify migration name")
	}

	Filename := Arguments[0]

	// Validates migration file exists.
	// TODO: Change `.sql` extenstion to `.go` in order to use `GORM` Migrator class instead!
	ReExp := regexp.MustCompile("_(.*)\\.sql")
	ExistingMigrationFiles, err := ioutil.ReadDir("../../migrations")
	if err != nil {
		fmt.Println(err)
	}
	for _, ExistingMigrationFile := range ExistingMigrationFiles {
		ExtractedExistingMigrationFileName := ReExp.FindStringSubmatch(ExistingMigrationFile.Name())[1]
		if Filename == ExtractedExistingMigrationFileName {
			fmt.Println("Migration already exist, please double-check your filename.")
			os.Exit(1)
		}
	}

	// Generate filename prefix with YYYYMMDDHHMMSS time format.
	Time := time.Now()
	FormattedTime := fmt.Sprintf("%d%02d%02d%02d%02d%02d",
		Time.Year(), Time.Month(), Time.Day(),
		Time.Hour(), Time.Minute(), Time.Second())

	f, err := os.Create(fmt.Sprintf("../../migrations/%s_%s.sql", FormattedTime, Filename))

	if err != nil {
		fmt.Println("Cannot create your migration file, please double-check your command.")
	}

	// EOF
	_ = f
}
