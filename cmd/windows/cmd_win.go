//go:build windows
// +build windows

package commands

import "fmt"

type winmode int

const (
	WinmodeHidden winmode = iota
	WinmodeMaximized
	WinmodeMinimized
)

type exec struct {
	input             bool
	Whitout_Shell     bool
	RunWithPowerShell bool
	// To better understand this type in your favorite command interpreter (in Windows) "powershell.exe -h".
	PowerShell struct {
		windowStyle struct {
			// The first one in the list that is true will be the one chosen to be implemented if there is more than one true value.
			Enabled   bool
			Hidden    bool
			Minimized bool
			Maximized bool
		}
		NonInteractive bool
		Mta            bool
		NoProfile      bool
		EncodedCommand bool
		Sta            bool
		NoExit         bool
		NoLogo         bool
	}
}

func (exec exec) formatCmd() string {
	var WindowsCommand string
	if exec.RunWithPowerShell {
		PShArgs := exec.PowerShell
		var SetTA, interactive, profile, encoded, nologo, exit, windowStyle_pre, windowStyle_Arg string
		if PShArgs.Mta && !PShArgs.Sta { // MTA set
			SetTA = "-Mta "
		}
		if PShArgs.Sta && !PShArgs.Mta { // STA set
			SetTA = "-Sta "
		}
		if PShArgs.NonInteractive {
			interactive = "-NonInteractive "
		}
		if PShArgs.NoProfile {
			profile = "-NoProfile "
		}
		if PShArgs.EncodedCommand {
			encoded = "-EncodedCommand "
		}
		if PShArgs.NoLogo {
			nologo = "-NoLogo "
		}
		if PShArgs.NoExit {
			exit = "-NoExit "
		}
		if PShArgs.windowStyle.Enabled {
			func() {
				WSpar := PShArgs.windowStyle
				windowStyle_pre = "-WindowStyle "
				if WSpar.Hidden {
					windowStyle_Arg = "Hidden"
					return
				}
				if WSpar.Maximized {
					windowStyle_Arg = "Maximized"
					return
				}
				if WSpar.Minimized {
					windowStyle_Arg = "Minimized"
				}
			}()
		}
		WindowsCommand = fmt.Sprintf("powershell.exe %v%v%v%v%v%v%v%v -Command %v", SetTA, interactive, profile, encoded, nologo, exit, windowStyle_pre, windowStyle_Arg, exec.input) // This is fucking infernal lol
	} else { // End of RunWithPowerShell declaration
		WindowsCommand = "cmd.exe /C"
	}
	return WindowsCommand
}

// window mode config
func (sh Sh) SetWindowPSMode(mode winmode) {
	sh.Windows.PowerShell.windowStyle.Enabled = true
	switch mode {
	case 0:
		sh.Windows.PowerShell.windowStyle.Hidden = true
	case 1:
		sh.Windows.PowerShell.windowStyle.Maximized = true
	case 2:
		sh.Windows.PowerShell.windowStyle.Minimized = true
	}
}
