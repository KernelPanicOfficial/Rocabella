package main

import (
	"Rocabella/Packages/Arguments"
	"Rocabella/Packages/Utils"
)

// main function
func main() {
	// Call function named CheckGoVersion
	Utils.CheckGoVersion()

	error := Arguments.RocabellaCli.Execute()
	if error != nil {
		return
	}
}
