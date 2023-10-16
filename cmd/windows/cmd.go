//go:build windows
// +build windows

package win

type winmode int

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
	sh.input = input
	return sh
}

// Running funcions

func (sh *WinCmd) formatcmd() string {
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
