package commands

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Sh struct {
	Path    string
	Windows struct {
		RunWithPowerShell bool
		// To better understand this type in your favorite command interpreter (in Windows) "powershell.exe -h".
		PowerShell struct {
			WindowStyle struct {
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

func (sh Sh) formatCmd() []string {
	var (
		command string
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
			if PShArgs.WindowStyle.Enabled {
				func() {
					WSpar := PShArgs.WindowStyle
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
			command = fmt.Sprintf("powershell.exe %v%v%v%v%v%v%v%v /C ", SetTA, interactive, profile, encoded, nologo, exit, windowStyle_pre, windowStyle_Arg) // This is fucking infernal lol
			return strings.Fields(command)
		}
		// End of RunWithPowerShell declaration
		return strings.Fields("cmd.exe /C ")

		// Set linux shell formatting
	} else if current_os == "linux" {
		var shell, sudo, arg string
		if sh.Linux.Bash {
			shell = "bash "
			arg = "-c "
		}
		if sh.Linux.CustomSh.Enable {
			shell = sh.Linux.CustomSh.ShName
			arg = sh.Linux.CustomSh.ShArg
		}
		if sh.Linux.RunWithSudo {
			sudo = "sudo "
		}
		command = sudo + shell + arg
		return strings.Fields(command)
	}
	return []string{}
}
func (sh Sh) setRunMode(input string) *exec.Cmd {
	var cmd *exec.Cmd
	if sh.Windows.RunWithPowerShell || sh.Linux.RunWithShell {
		fmtcmd := sh.formatCmd()                     // Format the command with the respective parameters
		cmd = exec.Command(fmtcmd[0], fmtcmd[1:]...) // declare the *os.Cmd val
	} else {
		input := strings.Fields(input)
		cmd = exec.Command(input[0], input[1:]...)
	}
	cmd.Path = sh.Path
	return cmd
}

// Exec Cmd method  
func (sh Sh) Cmd(input string) error {
	cmd := sh.setRunMode(input)
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
	cmd := sh.setRunMode(input)
	out, err := cmd.Output()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

func (sh Sh) Start(input string) error {
	cmd := sh.setRunMode(input)
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

func (sh Sh) GetCmdArg(input string) *exec.Cmd {
	cmd := sh.setRunMode(input)
	return cmd
}
