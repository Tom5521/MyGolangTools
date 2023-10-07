package main

import (
	"fmt"

	"github.com/Tom5521/MyGolangTools/commands"
)

var sh = commands.Sh{}

func main() {
	sh.Windows.RunWithPowerShell = true
	sh.Windows.PowerShell.WindowStyle.Enabled = true
	sh.Windows.PowerShell.WindowStyle.Minimized = true
	//sh.Windows.PowerShell.NoExit = false
	//sh.Windows.PowerShell.Mta = true
	ret, err := sh.Out("ls")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}
