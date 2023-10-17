//go:build unix
// +build unix

package cmd

import (
	"github.com/Tom5521/MyGolangTools/cmd/unix"
)

func Cmd(input string) unix.UnixCmd {
	return unix.Cmd(input)
}
func SudoCmd(input string) unix.UnixSudoCmd {
	return unix.Sudo_Cmd(input)
}
