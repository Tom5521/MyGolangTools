package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type winmode int

const (
	WinmodeHidden winmode = iota
	WinmodeMaximized
	WinmodeMinimized
)

type Sh struct {
	input   string
	Path    string
	Windows struct {
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
	Linux struct {
		RunWithShell bool
		RunWithSudo  bool
		CustomSh     struct {
			Enable bool
			ShName string
			ShArg  string // Shell execution cmd
		}
		Bash bool // Default linux shell is sh
	}
	CustomStd struct {
		Enable bool
		Stdin  bool
		Stdout bool
		Stderr bool
	}
}

func (sh Sh) formatCmd() string {
	var (
		LinuxCommand   string
		WindowsCommand string
	)
	current_os := runtime.GOOS
	// Sel windows shell formatting
	if current_os == "windows" {
		// Start of RunWithPowerShell declaration
		if sh.Windows.RunWithPowerShell {
			PShArgs := sh.Windows.PowerShell
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
			WindowsCommand = fmt.Sprintf("powershell.exe %v%v%v%v%v%v%v%v -Command %v", SetTA, interactive, profile, encoded, nologo, exit, windowStyle_pre, windowStyle_Arg, sh.input) // This is fucking infernal lol
		} else { // End of RunWithPowerShell declaration
			WindowsCommand = "cmd.exe /C"
		}
		// Set linux shell formatting
	}
	if current_os == "linux" {
		var shell, arg string
		if sh.Linux.Bash {
			shell = "bash "
			arg = "-c "
		}
		if sh.Linux.CustomSh.Enable {
			shell = sh.Linux.CustomSh.ShName
			arg = sh.Linux.CustomSh.ShArg
		}
		if sh.Linux.RunWithSudo && sh.Linux.RunWithShell {
			LinuxCommand = "sudo " + shell + arg + sh.input
		} else if sh.Linux.RunWithSudo && !sh.Linux.RunWithShell {
			LinuxCommand = "sudo " + sh.input
		} else if sh.Linux.RunWithShell {
			LinuxCommand = shell + arg + sh.input
		}
	}
	if runtime.GOOS == "windows" {
		return WindowsCommand
	} else if runtime.GOOS == "linux" {
		return LinuxCommand
	} else {
		return ""
	}
}
func (sh Sh) setRunMode() *exec.Cmd {
	var cmd *exec.Cmd
	if sh.Windows.RunWithPowerShell || sh.Linux.RunWithShell {
		fmtcmd := sh.formatCmd()
		command := strings.Fields(fmtcmd)
		// Format the command with the respective parameters
		cmd = exec.Command(command[0], command[1:]...) // declare the *os.Cmd val
	} else {
		input := strings.Fields(sh.input)
		cmd = exec.Command(input[0], input[1:]...)
	}
	if sh.Path != "" {
		cmd.Path = sh.Path
	}
	return cmd
}

// Exec Cmd method  
func (sh Sh) Cmd(input string) error {
	sh.input = input
	cmd := sh.setRunMode()
	// Set the standar input/output/error exit
	if sh.CustomStd.Enable {
		if sh.CustomStd.Stdout {
			cmd.Stdout = os.Stdout
		}
		if sh.CustomStd.Stdin {
			cmd.Stdin = os.Stdin
		}
		if sh.CustomStd.Stderr {
			cmd.Stderr = os.Stderr
		}
	} else { // Set the default values
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
	}

	// Exec the command
	err := cmd.Run()

	// Return the error
	if err != nil {
		return err
	}
	return nil // Return nil
}

// Out method  
func (sh Sh) Out(input string) (string, error) {
	sh.input = input
	cmd := sh.setRunMode()
	out, err := cmd.Output()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

func (sh Sh) Start(input string) error {
	sh.input = input
	cmd := sh.setRunMode()
	if sh.CustomStd.Enable {
		if sh.CustomStd.Stdout {
			cmd.Stdout = os.Stdout
		}
		if sh.CustomStd.Stdin {
			cmd.Stdin = os.Stdin
		}
		if sh.CustomStd.Stderr {
			cmd.Stderr = os.Stderr
		}
	} else { // Set the default values
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
	}
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

func (sh Sh) GetCmdArg() *exec.Cmd {
	cmd := sh.setRunMode()
	return cmd
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
