//go:build unix
// +build unix

package unix

import (
	"os"
	"os/exec"
	"strings"
)

// global struct
type sh struct {
	input string
	path  struct {
		Enabled bool
		Path    string
	}
	runWithShell struct {
		Enabled  bool
		bash     bool // Default linux shell is sh
		customSh struct {
			Enable bool
			ShName string
			ShArg  string // Shell execution cmd
		}
	}
	customStd struct {
		Enable bool
		Stdin  bool
		Stdout bool
		Stderr bool
	}
}

// Init functions

// Runs a normal command (without sudo)
func Cmd(Command string) sh {
	sh := sh{input: Command}
	return sh
}

// General parameter funcions
func (sh *sh) SetInput(input string) {
	sh.input = input
}
func (sh *sh) SetPath(path string) {
	sh.path.Enabled = true
	sh.path.Path = path
}

// If the value is true use exec.Command([shell],[arg],input) instead of exec.Command(input[0],input[1:]...)
func (sh *sh) RunWithShell(set bool) {
	sh.runWithShell.Enabled = set
}

// Set a custom stdin,stdout or stderr. Default std is all in false
func (sh *sh) CustomStd(Stdin, Stdout, Stderr bool) {
	sh.customStd.Enable = true
	sh.customStd.Stderr = Stderr
	sh.customStd.Stdout = Stdout
	sh.customStd.Stdin = Stdin
}
func (sh *sh) Stdin(set bool) {
	sh.customStd.Enable = true
	sh.customStd.Stdin = set
}
func (sh *sh) Stderr(set bool) {
	sh.customStd.Enable = true
	sh.customStd.Stderr = set
}
func (sh *sh) Stdout(set bool) {
	sh.customStd.Enable = true
	sh.customStd.Stdout = set
}

// Set a custom shell to exec the command
func (sh *sh) CustomShell(Shell_Name, Exec_Arg string) {
	sh.RunWithShell(true)
	sh.runWithShell.customSh.Enable = true
	sh.runWithShell.customSh.ShArg = Exec_Arg
	sh.runWithShell.customSh.ShName = Shell_Name
}

func (sh *sh) UseBashShell(set bool) {
	sh.RunWithShell(true)
	sh.runWithShell.bash = true
}

// Internal funcions

func (sh sh) setStd(cmd *exec.Cmd) {
	if sh.customStd.Enable {
		std := sh.customStd
		if std.Stderr {
			cmd.Stderr = os.Stderr
		}
		if std.Stdout {
			cmd.Stdout = os.Stdout
		}
		if std.Stdin {
			cmd.Stdin = os.Stdin
		}
	}
}
func (sh sh) getExec() *exec.Cmd {
	var cmd *exec.Cmd
	if sh.runWithShell.Enabled {
		if sh.runWithShell.bash {
			cmd = exec.Command("bash", "-c", sh.input)
		}
		if sh.runWithShell.customSh.Enable {
			cshell := sh.runWithShell.customSh
			cmd = exec.Command(cshell.ShName, cshell.ShArg, sh.input)
		}
	} else {
		command := strings.Fields(sh.input)
		cmd = exec.Command(command[0], command[1:]...)
	}
	return cmd
}

// normal running funcions
func (sh sh) Run() error {
	cmd := sh.getExec()
	sh.setStd(cmd)
	return cmd.Run()
}
func (sh sh) Out() (string, error) {
	cmd := sh.getExec()
	out, err := cmd.Output()
	return string(out), err
}
func (sh sh) CombinedOut() (string, error) {
	cmd := sh.getExec()
	sh.setStd(cmd)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func (sh sh) Start() error {
	cmd := sh.getExec()
	sh.setStd(cmd)
	return cmd.Start()
}
