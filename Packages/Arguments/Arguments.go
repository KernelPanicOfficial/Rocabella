package Arguments

import (
	"Rocabella/Packages/Colors"
	"fmt"
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
	__version__ = "1.0"
	//__version_name__ = "Death Star"
	__authors__ = "@nickvourd"
	__license__ = "MIT"
	__github__  = "https://github/nickvourd/Rocabella"

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
}

// StartRocabella function
func StartRocabella(cmd *cobra.Command, args []string) error {
	// Call function named ShowAscii
	ShowAscii()
	// Check if additional arguments were provided.
	if len(os.Args) <= 2 {
		// Display help message.
		err := cmd.Help()
		if err != nil {
			return err
		}
		// Exit the program.
		os.Exit(0)
	}
	return nil
}

// ShowAscii function
func ShowAscii() {
	// Initialize RandomColor
	randomColor := Colors.RandomColor()
	fmt.Print(randomColor(__ascii__))
	fmt.Printf(text, __version__, __license__, __authors__, __github__)
}
