package Arguments

import (
	"Rocabella/Packages/Colors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	__ascii__ = `

8888888b.                            888               888 888          
888   Y88b                           888               888 888          
888    888                           888               888 888          
888   d88P .d88b.   .d8888b  8888b.  88888b.   .d88b.  888 888  8888b.  
8888888P" d88""88b d88P"        "88b 888 "88b d8P  Y8b 888 888     "88b 
888 T88b  888  888 888      .d888888 888  888 88888888 888 888 .d888888 
888  T88b Y88..88P Y88b.    888  888 888 d88P Y8b.     888 888 888  888 
888   T88b "Y88P"   "Y8888P "Y888888 88888P"   "Y8888  888 888 "Y888888 																	                                                                                       
`
	__version__      = "1.0"
	__version_name__ = "Death Star"
	__authors__      = "@nickvourd"
	__license__      = "MIT"
	__github__       = "https://github/nickvourd/Rocabella"

	text = `
Rocabella v%s - Sniffing files generator.
Rocabella is an open source tool licensed under %s.
Written with <3 by %s.
Please visit %s for more...

`

	RocabellaCli = &cobra.Command{
		Use:          "Rocabella",
		SilenceUsage: true,
		RunE:         StartRocabella,
	}
)

// init function
// init all flags.
func init() {
	// Disable default command completion for Rocabella CLI.
	RocabellaCli.CompletionOptions.DisableDefaultCmd = true

	// Add commands to the Rocabella CLI.
	RocabellaCli.Flags().SortFlags = true
	RocabellaCli.Flags().BoolP("version", "v", false, "Show Rocabella current version")
	RocabellaCli.AddCommand(scArgument)
	RocabellaCli.AddCommand(libArgument)
	RocabellaCli.AddCommand(urlArgument)
	RocabellaCli.AddCommand(lnkArgument)

	// Add flags to specific commands
	// For sc command
	scArgument.Flags().StringP("target", "t", "", "Set target host")
	scArgument.Flags().StringP("icon", "i", "", "Set icon display")
	scArgument.Flags().StringP("output", "o", "", "set output file")

	// For lib command
	libArgument.Flags().StringP("target", "t", "", "Set target host")
	libArgument.Flags().StringP("icon", "i", "", "Set icon display")
	libArgument.Flags().StringP("output", "o", "", "set output file")

	// For url command
	urlArgument.Flags().StringP("target", "t", "", "Set target host")
	urlArgument.Flags().StringP("icon", "i", "", "Set icon display")
	urlArgument.Flags().StringP("output", "o", "", "set output file")

	// For lnk command
	lnkArgument.Flags().StringP("target", "t", "", "Set target host")
	lnkArgument.Flags().StringP("icon", "i", "", "Set icon display")
	lnkArgument.Flags().StringP("output", "o", "", "set output file")
}

// StartRocabella function
func StartRocabella(cmd *cobra.Command, args []string) error {
	logger := log.New(os.Stderr, "[!] ", 0)

	// Call function named ShowAscii
	ShowAscii()

	// Check if additional arguments were provided.
	if len(os.Args) == 1 {
		// Display help message.
		err := cmd.Help()

		// If error exists
		if err != nil {
			logger.Fatal(" Error:", err)
			return err
		}
	}

	// Obtain version flag
	versionFlag, _ := cmd.Flags().GetBool("version")

	// Call function names ShowVersion
	ShowVersion(versionFlag)

	return nil
}

// ShowAscii function
func ShowAscii() {
	// Initialize RandomColor
	randomColor := Colors.RandomColor()
	fmt.Print(randomColor(__ascii__))
	fmt.Printf(text, __version__, __license__, __authors__, __github__)
}

// ShowVersion function
func ShowVersion(versionFlag bool) {
	// if one argument
	if versionFlag {
		// if version flag exists
		fmt.Printf("[+] Current version: " + Colors.BoldRed(__version__) + "\n\n[+] Version name: " + Colors.BoldRed(__version_name__) + "\n\n")
		os.Exit(0)
	}
}
