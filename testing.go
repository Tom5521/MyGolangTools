package main

import (
	"github.com/Tom5521/MyGolangTools/commands"
)

var sh = commands.Sh{}

func main() {
	sh.SetWindowPSMode(commands.WinmodeHidden)
	sh.SetWindowPSMode(commands.WinmodeMinimized)
}
