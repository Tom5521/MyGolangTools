//go:build windows
// +build windows

package win

import (
	"os"
	"os/exec"
	"strings"
)

type WinCmd struct {
	runWithCmd bool
	input      string
	path       struct {
		enabled bool
		path    string
	}
	customStd struct {
		enabled bool
		stdin   bool
		stderr  bool
		stdout  bool
	}
}

// Init

func Cmd(input string) WinCmd {
	sh := WinCmd{}
	sh.SetInput(input)
	return sh
}

// Running functions

func (sh WinCmd) Run() error {
	return sh.getFinal().Run()
}
func (sh WinCmd) Start() error {
	return sh.getFinal().Run()
}
func (sh WinCmd) Out() (string, error) {
	cmd := sh.getExec()
	out, err := cmd.Output()
	return string(out), err
}
func (sh WinCmd) CombinedOut() (string, error) {
	out, err := sh.getFinal().CombinedOutput()
	return string(out), err
}

// Internal functions

func (sh WinCmd) getFinal() *exec.Cmd {
	cmd := sh.getExec()
	sh.setStd(cmd)
	return cmd
}

func (sh WinCmd) setStd(cmd *exec.Cmd) {
	if !sh.customStd.enabled {
		return
	}
	cStd := sh.customStd
	if cStd.stdin {
		cmd.Stdin = os.Stdin
	}
	if cStd.stdout {
		cmd.Stdout = os.Stdout
	}
	if cStd.stderr {
		cmd.Stderr = os.Stderr
	}
}

func (sh WinCmd) getExec() *exec.Cmd {
	command := strings.Fields(sh.formatcmd())
	cmd := exec.Command(command[0], command[1:]...)
	return cmd
}

func (sh WinCmd) formatcmd() string {
	var cmd string
	if sh.runWithCmd {
		cmd = "cmd.exe /C " + sh.input
	} else {
		cmd = sh.input
	}
	return cmd
}

// Global config parameters

func (sh *WinCmd) SetPath(path string) {
	sh.path.enabled = true
	sh.path.path = path
}

func (sh *WinCmd) RunWithCmd(set bool) {
	sh.runWithCmd = set
}

func (sh *WinCmd) SetInput(input string) {
	sh.input = input
}

// Set custom std
func (sh *WinCmd) CustomStd(Stdin, Stdout, Stderr bool) {
	sh.customStd.enabled = true
	sh.customStd.stderr = Stderr
	sh.customStd.stdin = Stdin
	sh.customStd.stdout = Stdout
}

func (sh *WinCmd) Stdin(set bool) {
	sh.customStd.enabled = true
	sh.customStd.stdin = set
}
func (sh *WinCmd) Stdout(set bool) {
	sh.customStd.enabled = true
	sh.customStd.stdout = set
}
func (sh *WinCmd) Stderr(set bool) {
	sh.customStd.enabled = true
	sh.customStd.stderr = set
}
