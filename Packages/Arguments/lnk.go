package Arguments

import (
	"Rocabella/Packages/Manager"
	"Rocabella/Packages/Templates"
	"log"
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
		logger := log.New(os.Stderr, "[!] ", 0)

		// Check if additional arguments were provided.
		if len(os.Args) <= 2 {
			// Show help message.
			error := cmd.Help()
			if error != nil {
				logger.Fatal("Error:", error)
				return error
			}

			// Exit the program.
			os.Exit(0)
		}

		// Get the value of the any flag
		target, _ := cmd.Flags().GetString("target")
		icon, _ := cmd.Flags().GetString("icon")
		output, _ := cmd.Flags().GetString("output")

		// If target is empty
		if target == "" {
			logger.Fatal("The '-t' or '--target' flag is required for the command to proceed.")
		}

		// If output is empty
		if output == "" {
			logger.Fatal("The '-o' or '--output' flag is required for the command to proceed.")
		}

		// Joining arrays into a single slice
		allIcons := append(append(microsoftOfficeIcons, microsoftWindowsIcons...), thirdPartyIcons...)

		// Call function named ValidateArgument
		ValidateArgument("icons", icon, allIcons)

		// Access the variable from the Templates package
		iconTamplates := Templates.InitializeIconTemplates()

		// Call function named GetIconPath
		iconPath := GetIconPath(icon, iconTamplates)

		// Call function named CreateLNK
		Manager.CreateLNK(target, icon, output, iconPath)

		return nil
	},
}
