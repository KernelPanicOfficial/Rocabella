package main

import (
	"Rocabella/Packages/Arguments"
)

// main function
func main() {
	error := Arguments.RocabellaCli.Execute()
	if error != nil {
		return
	}
}
