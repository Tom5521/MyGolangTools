package commands

import (
	"os"
	"os/exec"
	"runtime"
)

type Sh struct {
	Windows struct {
		PowerShell bool // Default terminal/shell is cmd
	}
	Linux struct {
		CustomSh struct {
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

func (sh Sh) formatCmd() [2]string {
	var (
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
			arg = args[1]
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
	}
	return [2]string{shell, arg}
}

// Exec Cmd method  
func (sh Sh) Cmd(input string) error {
	fmtcmd := sh.formatCmd()
	cmd := exec.Command(fmtcmd[0], fmtcmd[1], input)
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
	} else {
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
	}

	err := cmd.Run()

	if err != nil {
		return err
	}
	return nil
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