package Output

import (
	"log"
	"os"
	"path/filepath"
)

// OutputValidate function
func OutputValidate(output string, statement int) {
	logger := log.New(os.Stderr, "[!] ", 0)
	extension := filepath.Ext(output)

	switch statement {
	case 1:
		if extension != ".lnk" {
			logger.Fatal("The output file must have a '.lnk' extension.")
		}
	case 2:
		if extension != ".searchConnector-ms" {
			logger.Fatal("The output file must have a '.searchConnector-ms' extension.")
		}
	case 3:
		if extension != ".library-ms" {
			logger.Fatal("The output file must have a '.library-ms' extension.")
		}
	case 4:
		if extension != ".url" {
			logger.Fatal("The output file must have a '.url' extension.")
		}
	default:
		logger.Fatal("Invalid statement.")
	}

}
