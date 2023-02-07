package clear

import (
	l "aresbot/pkg/logger"
	"os"
	"os/exec"
	"runtime"
)

func CallClear() {
	var clear map[string]func() //create a map for storing clear funcs

	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //darwin example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		l.ErrorLogger.Println("Your platform is unsupported! I can't clear terminal screen :(")
	}
}
