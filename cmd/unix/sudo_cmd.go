//go:build unix
// +build unix

package unix

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type UnixSudoCmd struct {
	UnixCmd
	sudo_pars struct {
		getted bool
		Passwd string
	}
}

// Runs a command as sudo
func Sudo_Cmd(command string, optional_password ...string) UnixSudoCmd {
	sudoSh := UnixSudoCmd{}
	sudoSh.input = command
	if len(optional_password) >= 1 {
		sudoSh.SetSudoPasswd(optional_password[0])
	}
	return sudoSh
}

// Sudo parameters
func (sh *UnixSudoCmd) SetSudoPasswd(password string) {
	sh.sudo_pars.getted = true
	sh.sudo_pars.Passwd = password
}

// Internal sudo functions
func (sh UnixSudoCmd) getExec() *exec.Cmd {
	var cmd *exec.Cmd
	if sh.runWithShell.Enabled {
		if sh.runWithShell.bash {
			command := strings.Fields(fmt.Sprintf("bash -c \"sudo -S %v\"", sh.input))
			cmd = exec.Command(command[0], command[1:]...)
		}
		if sh.runWithShell.customSh.Enable {
			cshell := sh.runWithShell.customSh
			command := strings.Fields(fmt.Sprintf("sudo -S %v %v %v", cshell.ShName, cshell.ShArg, sh.input))
			cmd = exec.Command(command[0], command[1:]...)
		}
	} else {
		command := strings.Fields("sudo -S " + sh.input)
		cmd = exec.Command(command[0], command[1:]...)
	}
	return cmd
}

func (sh UnixSudoCmd) writePasswd(cmd *exec.Cmd) {
	stdin, _ := cmd.StdinPipe()
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, sh.sudo_pars.Passwd)
	}()
}

// sudo running funcions

func (sh UnixSudoCmd) Run() error {
	cmd := sh.getExec()
	sh.setStd(cmd)
	sh.writePasswd(cmd)
	return cmd.Run()
}
func (sh UnixSudoCmd) Out() (string, error) {
	cmd := sh.getExec()
	sh.writePasswd(cmd)
	out, err := cmd.Output()
	return string(out), err
}

func (sh UnixSudoCmd) CombinedOut() (string, error) {
	cmd := sh.getExec()
	sh.setStd(cmd)
	sh.writePasswd(cmd)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func (sh UnixSudoCmd) Start() error {
	cmd := sh.getExec()
	sh.setStd(cmd)
	sh.writePasswd(cmd)
	return cmd.Start()
}
