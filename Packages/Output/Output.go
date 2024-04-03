package Output

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// OutputValidate function
func OutputValidate(output string, statement int) string {
	extension := filepath.Ext(output)

	switch statement {
	case 1:
		if extension != ".lnk" {
			//Call function named AddExtension
			output = AddExtension(".lnk", output)
		}
	case 2:
		if extension != ".searchConnector-ms" {
			// Call function named AddExtension
			output = AddExtension(".searchConnector-ms", output)
		}
	case 3:
		if extension != ".library-ms" {
			// Call function named AddExtension
			output = AddExtension(".library-ms", output)
		}
	case 4:
		if extension != ".url" {
			// Call function named AddExtension
			output = AddExtension(".url", output)
		}
	default:
		// Do nothing
	}

	return output
}

// AddExtension function
func AddExtension(extension string, output string) string {
	// Decalre varaibles
	var addExtension string = output

	// Add extension to output
	addExtension = output + extension
	fmt.Printf("[!] Added the '%s' extension to %s file\n\n", extension, output)

	return addExtension
}

// WriteToFile function
func WriteToFile(file string, data string) {
	logger := log.New(os.Stderr, "[!] ", 0)
	f, err := os.Create(file)
	if err != nil {
		logger.Fatal("Error:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		logger.Fatal("Error:", err)
		return
	}
}
