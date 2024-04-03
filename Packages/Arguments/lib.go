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

		// Check if additional arguments were provided.
		if len(os.Args) <= 2 {
			// Show help message.
			error := cmd.Help()
			if error != nil {
				logger.Fatal("Error", error)
				return error
			}

			// Exit the program.
			os.Exit(0)
		}

		return nil
	},
}
