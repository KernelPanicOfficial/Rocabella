package Arguments

import (
	"Rocabella/Packages/Colors"
	"Rocabella/Packages/Templates"
	"fmt"
	"log"
	"os"
	"strings"

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
	microsoftOfficeIcons  = []string{"access", "excel", "lync", "office", "onedrive", "onenote", "outlook", "powerpoint", "project", "publisher", "visio", "word"}
	microsoftWindowsIcons = []string{"bluetooth", "calc", "chm", "defender", "defrag", "edge", "explorer", "dir", "directory", "folder", "help", "hlp", "internet-explorer", "ie", "keyboard", "magnify", "mail", "media-player", "mobile-sync", "mspaint", "paint", "notepad", "onedrive2", "package", "pdf", "perfmon", "teams", "uac-shield", "uac", "werfault", "windows-store", "xbox"}
	thirdPartyIcons       = []string{"adobe-reader", "chrome", "citrix", "cyberark-epm", "firefox", "global-protect", "java", "java-update", "python", "snow-agent", "synaptics-touchpad"}

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
	RocabellaCli.Flags().BoolP("list", "l", false, "Show the list of available icons")
	RocabellaCli.AddCommand(scArgument)
	RocabellaCli.AddCommand(libArgument)
	RocabellaCli.AddCommand(urlArgument)
	RocabellaCli.AddCommand(lnkArgument)

	// Add flags to specific commands
	// For sc command
	scArgument.Flags().StringP("target", "t", "", "Set target host")
	scArgument.Flags().StringP("icon", "i", "", "Set icon display")
	scArgument.Flags().StringP("output", "o", "", "Set output file")

	// For lib command
	libArgument.Flags().StringP("target", "t", "", "Set target host")
	libArgument.Flags().StringP("icon", "i", "", "Set icon display")
	libArgument.Flags().StringP("output", "o", "", "Set output file")

	// For url command
	urlArgument.Flags().StringP("target", "t", "", "Set target host")
	urlArgument.Flags().StringP("icon", "i", "", "Set icon display")
	urlArgument.Flags().StringP("output", "o", "", "Set output file")

	// For lnk command
	lnkArgument.Flags().StringP("target", "t", "", "Set target host")
	lnkArgument.Flags().StringP("icon", "i", "", "Set icon display")
	lnkArgument.Flags().StringP("output", "o", "", "Set output file")
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

// ValidateIcon function
func ValidateIcon(argValue string, icons []Templates.IconTemplates) string {
	for _, icon := range icons {
		if strings.EqualFold(strings.ToLower(argValue), strings.ToLower(icon.Name)) {
			iconPath := `"` + icon.Path + `"`
			switch strings.ToLower(argValue) {
			case "office":
				iconPath += "," + "56"
			case "explorer", "dir", "directory", "folder":
				iconPath += "," + "0"
			case "pdf":
				iconPath += "," + "13"
			default:
				// Do nothing, iconPath remains unchanged
			}
			return iconPath
		}
	}
	return ""
}

// ShowIconList function
func ShowIconList(iconListFlag bool, iconList1 []string, iconList2 []string, iconList3 []string) {
	// if iconListFlag is enabled
	if iconListFlag {
		microsoftOfficeIcons := strings.Join(iconList1, ", ")
		microsoftWindowsIcons := strings.Join(iconList2, ", ")
		thirdPartyIcons := strings.Join(iconList3, ", ")
		fmt.Printf("[+] Available icons:\n\n"+Colors.BoldYellow("===Microsoft Office icons===")+"\n\n%s\n\n"+Colors.BoldYellow("===Microsoft Windows icons===")+"\n\n%s\n\n"+Colors.BoldYellow("===Third-party icons (if they exist on the system)===")+"\n\n%s\n\n", Colors.BoldCyan(microsoftOfficeIcons), Colors.BoldCyan(microsoftWindowsIcons), Colors.BoldCyan(thirdPartyIcons))
		os.Exit(0)
	}

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

	// Obtain flag
	versionFlag, _ := cmd.Flags().GetBool("version")
	iconListFlag, _ := cmd.Flags().GetBool("list")

	// Call function named ShowVersion
	ShowVersion(versionFlag)

	// Call function named ShowIconList
	ShowIconList(iconListFlag, microsoftOfficeIcons, microsoftWindowsIcons, thirdPartyIcons)

	return nil
}
