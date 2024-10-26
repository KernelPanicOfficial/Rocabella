package Arguments

import (
	"Rocabella/Packages/Manager"
	"Rocabella/Packages/Output"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// libArgument represents the 'lib' command in the CLI.
var libArgument = &cobra.Command{
	// Use defines how the command should be called.
	Use:          "LIB",
	Short:        "library-ms file",
	SilenceUsage: true,
	Aliases:      []string{"lib"},

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
				logger.Fatal("Error ", err)
				return err
			}

			// Exit the program.
			os.Exit(0)
		}

		// Get the value of the any flag
		target, _ := cmd.Flags().GetString("target")
		output, _ := cmd.Flags().GetString("output")
		port, _ := cmd.Flags().GetString("port")
		share, _ := cmd.Flags().GetString("share")

		// If target is empty
		if target == "" {
			logger.Fatal("The '-t' or '--target' flag is required for the command to proceed.")
		}

		// If output is empty
		if output == "" {
			logger.Fatal("The '-o' or '--output' flag is required for the command to proceed.")
		}

		// Call function named OutputValidate
		output = Output.OutputValidate(output, 3)

		// Call function named CreateLIB
		Manager.CreateLIB(target, output, port, share)

		return nil
	},
}
