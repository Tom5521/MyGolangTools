package commands

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Sh struct {
	RunWithShell bool
	Windows      struct {
		PowerShell bool // Default terminal/shell is cmd
	}
	Linux struct {
		RunWithSudo bool
		CustomSh    struct {
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

func (sh Sh) formatCmd() [4]string {
	var (
		sudo       string
		shell, arg string
		shells     = [4]string{"sh", "bash", "PowerShell.exe", "cmd"}
		args       = [2]string{"-c", "/C"}
	)
	current_os := runtime.GOOS
	// Sel windows shell formatting
	if current_os == "windows" {
		arg = args[1]
		if sh.Windows.PowerShell {
			shell = shells[2]
		} else {
			shell = shells[3]
		}
		// Set linux shell formatting
	} else if current_os == "linux" {
		if sh.Linux.Bash {
			shell = shells[1]
			arg = args[0]
		} else {
			shell = shells[0]
			arg = args[0]
		}
		if sh.Linux.CustomSh.Enable {
			shell = sh.Linux.CustomSh.ShName
			arg = sh.Linux.CustomSh.ShArg
		}
		if sh.Linux.RunWithSudo {
			sudo = "sudo"
		}
	}
	return [4]string{shell,arg,sudo}
}

// Exec Cmd method  
func (sh Sh) Cmd(input string) error {
	var cmd *exec.Cmd
	if !sh.RunWithShell {
		fmtcmd := sh.formatCmd()                                              // Format the command with the respective parameters
		cmd = exec.Command(fmtcmd[0], fmtcmd[1], fmtcmd[2], fmtcmd[3], input) // declare the *os.Cmd val
	} else {
		input := strings.Fields(input)
		cmd = exec.Command(input[0], input[1:]...)
	}
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
	fmtCmd := sh.formatCmd()
	cmd := exec.Command(fmtCmd[0], fmtCmd[1], input)
	out, err := cmd.Output()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}
