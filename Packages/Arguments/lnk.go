package Arguments

import (
	"Rocabella/Packages/Manager"
	"os"

	"github.com/spf13/cobra"
)

// lnkArgument represents the 'lnk' command in the CLI.
var lnkArgument = &cobra.Command{
	// Use defines how the command should be called.
	Use:          "lnk",
	Short:        "lnk shortcut file",
	SilenceUsage: true,

	// RunE defines the function to run when the command is executed.
	RunE: func(cmd *cobra.Command, args []string) error {
		// Check if additional arguments were provided.
		if len(os.Args) <= 2 {
			// Show help message.
			err := cmd.Help()
			if err != nil {
				return err
			}

			// Exit the program.
			os.Exit(0)
		}

		// Get the value of the any flag
		target, _ := cmd.Flags().GetString("target")
		icon, _ := cmd.Flags().GetString("icon")
		output, _ := cmd.Flags().GetString("output")

		// Call function named CreateLNK
		Manager.CreateLNK(target, icon, output)

		return nil
	},
}
