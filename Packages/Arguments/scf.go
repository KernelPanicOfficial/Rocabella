package Arguments

import (
	"Rocabella/Packages/Manager"
	"Rocabella/Packages/Output"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// scfArgument represents the 'scf' command in the CLI.
var scfArgument = &cobra.Command{
	// Use defines how the command should be called.
	Use:          "scf",
	Short:        "scf shortcut file",
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
				logger.Fatal("Error:", err)
				return err
			}

			// Exit the program.
			os.Exit(0)
		}

		// Get the value of the any flag
		target, _ := cmd.Flags().GetString("target")
		output, _ := cmd.Flags().GetString("output")
		port, _ := cmd.Flags().GetString("port")

		// If target is empty
		if target == "" {
			logger.Fatal("The '-t' or '--target' flag is required for the command to proceed.")
		}

		// If output is empty
		if output == "" {
			logger.Fatal("The '-o' or '--output' flag is required for the command to proceed.")
		}

		// Call function named OutputValidate
		output = Output.OutputValidate(output, 5)

		// Call function named CreateSCF
		Manager.CreateSCF(target, output, port)

		return nil
	},
}
