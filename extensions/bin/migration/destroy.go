package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	// Execute ./destroy [filename] in command line.
	Arguments := os.Args[1:]

	if len(Arguments) < 1 {
		fmt.Println("Please specify migration name")
	}

	Filename := Arguments[0]

	// Validates migration file exists.
	ReExp := regexp.MustCompile("_(.*)\\.sql")
	ExistingMigrationFiles, err := ioutil.ReadDir("../../migrations")
	if err != nil {
		fmt.Println(err)
	}
	for _, ExistingMigrationFile := range ExistingMigrationFiles {
		ExtractedExistingMigrationFileName := ReExp.FindStringSubmatch(ExistingMigrationFile.Name())[1]
		if Filename == ExtractedExistingMigrationFileName {
			err := os.Remove(fmt.Sprintf("../../migrations/%s", ExistingMigrationFile.Name()))

			if err != nil {
				fmt.Println("Something went wrong, cannot delete migration file.")
			}

			os.Exit(0)
		}
	}

	fmt.Println("Migration does not exist, please double-check your filename.")
	os.Exit(1)
}
