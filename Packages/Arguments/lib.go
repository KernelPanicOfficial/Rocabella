package Arguments

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// libArgument represents the 'lib' command in the CLI.
var libArgument = &cobra.Command{
	// Use defines how the command should be called.
	Use:          "lib",
	Short:        "library-ms file",
	SilenceUsage: true,

	// RunE defines the function to run when the command is executed.
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := log.New(os.Stderr, "[!] ", 0)

		// Call function named ShowAscii
		ShowAscii()

		// Check if additional arguments were provided.
		if len(os.Args) <= 2 {
			// Show help message.
			err := cmd.Help()
			if err != nil {
				logger.Fatal("Error", err)
				return err
			}

			// Exit the program.
			os.Exit(0)
		}

		return nil
	},
}

/*
Documents Folder Type: {f42ee2d3-909f-4907-8871-4c22fc0bf756}
Represents the Documents folder in Windows.
Pictures Folder Type: {0ddd015d-b06c-45d5-8c4c-f59713854639}
Represents the Pictures folder in Windows.
Music Folder Type: {a0c69a99-21c8-4671-8703-7934162fcf1d}
Represents the Music folder in Windows.
Videos Folder Type: {491e922f-5643-4af4-a7eb-4e7a138d8174}
Represents the Videos folder in Windows.
Downloads Folder Type: {7d83ee9b-2244-4e70-b1f5-5393042af1e4}
Represents the Downloads folder in Windows.
Desktop Folder Type: {B4BFCC3A-DB2C-424C-B029-7FE99A87C641}
Represents the Desktop folder in Windows.
*/
