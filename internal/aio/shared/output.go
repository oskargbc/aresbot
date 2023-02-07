package shared

import (
	"github.com/fatih/color"
	"time"
)

var (
	white    = color.New(color.FgWhite)
	hiYellow = color.New(color.FgHiRed)
	yellow   = color.New(color.FgRed)
	red      = color.New(color.FgRed)
	green    = color.New(color.FgGreen)
)

func OStep(iD int, s string) {
	t := time.Now()

	color.Blue("[Task %d] - %s - Step: %s", iD, t.Format("15:04:05"), s)
}

func OInfo(iD int, s string) {
	t := time.Now()

	color.White("[Task %d] - %s - Info: %s", iD, t.Format("15:04:05"), s)
}

func OInfoF(s string) {
	t := time.Now()

	PrintTag()
	color.White(" - %s - Info: %s", t.Format("15:04:05"), s)
}

func OWarning(iD int, s string) {
	t := time.Now()

	color.Yellow("[Task %d] - %s - Warning: %s", iD, t.Format("15:04:05"), s)
}

func OError(iD int, s string) {
	t := time.Now()

	color.Red("[Task %d] - %s - Error: %s", iD, t.Format("15:04:05"), s)
}

func OErrorF(s string) {
	t := time.Now()

	PrintTag()
	color.Red(" - %s - Error: %s", t.Format("15:04:05"), s)
}

func OSuccess(iD int, s string) {
	t := time.Now()

	color.Green("[Task %d] - %s - Success: %s", iD, t.Format("15:04:05"), s)
}

func PrintTag() {
	white.Print("[")
	hiYellow.Print("AresBot")
	white.Print("]")
}
