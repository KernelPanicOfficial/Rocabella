package main

import (
	"Rocabella/Packages/Arguments"
	"Rocabella/Packages/Utils"
	"log"
	"os"
)

// main function
func main() {
	logger := log.New(os.Stderr, "[!] ", 0)

	// Call function named CheckGoVersion
	Utils.CheckGoVersion()

	err := Arguments.RocabellaCli.Execute()
	if err != nil {
		logger.Fatal("Error", err)
		return
	}
}
