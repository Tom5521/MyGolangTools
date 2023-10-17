//go:build windows
// +build windows

package cmd

import win "github.com/Tom5521/MyGolangTools/cmd/windows"

func Cmd(input string) win.WinCmd {
	cmd := win.Cmd(input)
	return cmd
}

func PSCmd(input string) win.WinPS {
	cmd := win.PSCmd(input)
	return cmd
}
