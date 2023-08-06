package tools

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

// Exec Cmd method  
func (sh Sh) Cmd(input string) (string, error) {
	var (
		shell, arg string
		shells     = [4]string{"sh", "bash", "PowerShell.exe", "cmd"}
		args       = [2]string{"-c", "/C"}
	)
	getos := runtime.GOOS
	// Sel windows shell formatting
	if getos == "windows" {
		arg = args[1]
		if sh.Windows.PowerShell {
			shell = shells[2]
		} else {
			shell = shells[3]
			arg = args[1]
		}
		// Set linux shell formatting
	} else if getos == "linux" {
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

	cmd := exec.Command(shell, arg, input)
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

	out, err := cmd.CombinedOutput()
	return string(out), err
}

// Out method  
func (sh Sh) Out(input string) (string, error) {
	com := Sh{}
	com.CustomStd.Enable = true
	com.CustomStd.Stdout = false
	com.CustomStd.Stdin = false
	com.CustomStd.Stderr = false
	out, err := com.Cmd(input)
	if err != nil {
		return out, err
	}
	return out, nil
}
