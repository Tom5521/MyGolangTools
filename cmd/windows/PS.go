//go:build windows
// +build windows

package win

const (
	PSWinmodeHidden winmode = iota
	PSWinmodeMaximized
	PSWinmodeMinimized
)

type WinPS struct {
	WinCmd
	// To better understand this type in your favorite command interpreter (in Windows) "powershell.exe -h".
	powerShell struct {
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

// Init
func PSCmd(input string) WinPS {
	sh := WinPS{}
	sh.input = input
	return sh
}

// Powershell config parameters

func (sh *WinPS) SetPSNoLogo(set bool) {
	sh.powerShell.NoLogo = set
}

func (sh *WinPS) SetPSNoExit(set bool) {
	sh.powerShell.NoLogo = set
}

func (sh *WinPS) SetPShEncodedCmd(set bool) {
	sh.powerShell.EncodedCommand = set
}

func (sh *WinPS) SetPSNoProfile(set bool) {
	sh.powerShell.NoProfile = set
}

func (sh *WinPS) SetPSSta(set bool) {
	sh.powerShell.Sta = set
}

func (sh *WinPS) SetPSMta(set bool) {
	sh.powerShell.Mta = set
}

func (sh *WinPS) SetPShNoInteractive(set bool) {
	sh.powerShell.NonInteractive = set
}

func (sh *WinPS) SetWindowPSMode(mode winmode) {
	sh.powerShell.windowStyle.Enabled = true
	switch mode {
	case 0:
		sh.powerShell.windowStyle.Hidden = true
	case 1:
		sh.powerShell.windowStyle.Maximized = true
	case 2:
		sh.powerShell.windowStyle.Minimized = true
	}
}

// powershell logic
/*
if exec.runWithPowerShell {
		PShArgs := exec.powerShell
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
	}

*/
