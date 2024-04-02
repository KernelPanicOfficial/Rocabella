package Templates

// Structure
type IconTemplates struct {
	Name string
	Path string
}

// InitializeIconTemplates function
// InitializeIconTemplates function initializes and returns a slice of IconTemplates.
// Each IconTemplates struct represents an icon template with its name and path.
// The path points to the location of the icon executable file.
func InitializeIconTemplates() []IconTemplates {
	return []IconTemplates{
		{"access", "%PROGRAMFILES%\\Microsoft Office\\root\\vfs\\Windows\\Installer\\{90160000-000F-0000-1000-0000000FF1CE}\\accicons.exe"},
		{"excel", "%PROGRAMFILES%\\Microsoft Office\\root\\vfs\\Windows\\Installer\\{90160000-000F-0000-1000-0000000FF1CE}\\xlicons.exe"},
		{"lync", "%PROGRAMFILES%\\Microsoft Office\\root\\vfs\\Windows\\Installer\\{90160000-000F-0000-1000-0000000FF1CE}\\lyncicon.exe"},
		{"office", "%PROGRAMFILES%\\Microsoft Office\\root\\vfs\\Windows\\Installer\\{90160000-000F-0000-1000-0000000FF1CE}\\ohub32.exe"},
		{"onedrive", "%PROGRAMFILES%\\Microsoft Office\\root\\vfs\\Windows\\Installer\\{90160000-000F-0000-1000-0000000FF1CE}\\grv_icons.exe"},
		{"onenote", "%PROGRAMFILES%\\Microsoft Office\\root\\vfs\\Windows\\Installer\\{90160000-000F-0000-1000-0000000FF1CE}\\joticon.exe"},
		{"outlook", "%PROGRAMFILES%\\Microsoft Office\\root\\vfs\\Windows\\Installer\\{90160000-000F-0000-1000-0000000FF1CE}\\outicon.exe"},
		{"powerpoint", "%PROGRAMFILES%\\Microsoft Office\\root\\vfs\\Windows\\Installer\\{90160000-000F-0000-1000-0000000FF1CE}\\pptico.exe"},
		{"project", "%PROGRAMFILES%\\Microsoft Office\\root\\vfs\\Windows\\Installer\\{90160000-000F-0000-1000-0000000FF1CE}\\pj11icon.exe"},
		{"publisher", "%PROGRAMFILES%\\Microsoft Office\\root\\vfs\\Windows\\Installer\\{90160000-000F-0000-1000-0000000FF1CE}\\pubs.exe"},
		{"visio", "%PROGRAMFILES%\\Microsoft Office\\root\\vfs\\Windows\\Installer\\{90160000-000F-0000-1000-0000000FF1CE}\\visicon.exe"},
		{"word", "%PROGRAMFILES%\\Microsoft Office\\root\\vfs\\Windows\\Installer\\{90160000-000F-0000-1000-0000000FF1CE}\\wordicon.exe"},
		{"bluetooth", "%WINDIR%\\System32\\bthudtask.exe"},
		{"calc", "%WINDIR%\\System32\\calc.exe"},
		{"chm", "%WINDIR%\\hh.exe"},
		{"defender", "%PROGRAMFILES%\\Windows Defender\\MpCmdRun.exe"},
		{"defrag", "%WINDIR%\\System32\\defrag.exe"},
		{"edge", "%ProgramFiles(x86)%\\Microsoft\\Edge\\Application\\msedge.exe"},
		{"explorer", "%WINDIR%\\explorer.exe"},
		{"dir", "%WINDIR%\\explorer.exe"},
		{"directory", "%WINDIR%\\explorer.exe"},
		{"folder", "%WINDIR%\\explorer.exe"},
		{"help", "%WINDIR%\\HelpPane.exe"},
		{"hlp", "%WINDIR%\\HelpPane.exe"},
		{"internet-explorer", "%PROGRAMFILES%\\internet explorer\\iexplore.exe"},
		{"ie", "%PROGRAMFILES%\\internet explorer\\iexplore.exe"},
		{"keyboard", "%WINDIR%\\System32\\osk.exe"},
		{"magnify", "%WINDIR%\\System32\\magnify.exe"},
		{"mail", "%PROGRAMFILES%\\Windows Mail\\wab.exe"},
		{"media-player", "%PROGRAMFILES%\\Windows Media Player\\wmplayer.exe"},
		{"mobile-sync", "%WINDIR%\\System32\\mobsync.exe"},
		{"mspaint", "%WINDIR%\\System32\\mspaint.exe"},
		{"paint", "%WINDIR%\\System32\\mspaint.exe"},
		{"notepad", "%WINDIR%\\notepad.exe"},
		{"onedrive2", "%PROGRAMFILES%\\Microsoft OneDrive\\OneDrive.exe"},
		{"package", "%WINDIR%\\System32\\OptionalFeatures.exe"},
		{"pdf", "%ProgramFiles(x86)%\\Microsoft\\Edge\\Application\\msedge.exe"},
		{"perfmon", "%WINDIR%\\System32\\perfmon.exe"},
		{"python", "%WINDIR%\\py.exe"},
		{"teams", "%LOCALAPPDATA%\\Microsoft\\Teams\\current\\Teams.exe"},
		{"uac-shield", "%WINDIR%\\System32\\UserAccountControlSettings.exe"},
		{"uac", "%WINDIR%\\System32\\UserAccountControlSettings.exe"},
		{"werfault", "%WINDIR%\\System32\\werfault.exe"},
		{"windows-store", "%WINDIR%\\System32\\WSCollect.exe"},
		{"xbox", "%WINDIR%\\System32\\GamePanel.exe"},
		{"adobe-reader", "%ProgramFiles(x86)%\\Adobe\\Acrobat Reader DC\\Reader\\AcroRd32.exe"},
		{"chrome", "%PROGRAMFILES%\\Google\\Chrome\\Application\\chrome.exe"},
		{"citrix", "%ProgramFiles(x86)%\\Citrix\\ICA Client\\redirector.exe"},
		{"cyberark-epm", "%PROGRAMFILES%\\CyberArk\\Endpoint Privilege Manager\\Agent\\vf_agent.exe"},
		{"firefox", "%PROGRAMFILES%\\Mozilla Firefox\\firefox.exe"},
		{"global-protect", "%PROGRAMFILES%\\Palo Alto Networks\\GlobalProtect\\PanGPA.exe"},
		{"java", "%ProgramFiles(x86)%\\Common Files\\Java\\Java Update\\jucheck.exe"},
		{"java-update", "%ProgramFiles(x86)%\\Common Files\\Java\\Java Update\\jusched.exe"},
		{"snow-agent", "%PROGRAMFILES%\\Snow Software\\Inventory\\Agent\\snowagent.exe"},
		{"synaptics-touchpad", "%WINDIR%\\System32\\SynTPEnh.exe"},
	}
}
