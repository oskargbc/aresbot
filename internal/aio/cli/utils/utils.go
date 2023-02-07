package utils

import (
	"aresbot/pkg/clear"
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"runtime"
	"strings"
	"time"
)

var (
	White    = color.New(color.FgWhite)
	HiYellow = color.New(color.FgHiRed)
	Yellow   = color.New(color.FgRed)
	Red      = color.New(color.FgRed)
	Green    = color.New(color.FgGreen)
)

func AwaitInput(tag string, message string, in []string) *string {
	PrintTag("MENU")
	Red.Println("[X] - Quit")

	StaticClearLine()

	PrintTag(tag)
	reader := bufio.NewReader(os.Stdin)
	HiYellow.Print(message)

	in = append(in, "X")
	in = append(in, "B")

	text, _ := reader.ReadString('\n')

	if runtime.GOOS == "windows" {
		text = strings.Replace(text, "\n", "", -1)
	}

	canChoose := false
	for _, v := range in {
		if strings.TrimSpace(text) == v {
			canChoose = true
		}
	}

	if !canChoose {
		return nil
	}
	t := strings.TrimSpace(text)

	if t == "X" {
		PrintWithTag("AresBot", "Exiting..")
		time.Sleep(time.Second * 1)
		os.Exit(-1)
	}
	return &t
}

func AwaitTextInput(tag string, message string, minLen int) *string {
	PrintTag(tag)
	HiYellow.Print(message)

	reader := bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')

	if runtime.GOOS == "windows" {
		text = strings.Replace(text, "\n", "", -1)
	}

	t := strings.TrimSpace(text)
	if len(t) < minLen {
		return nil
	}

	return &t
}

func PrintTag(tag string) {
	White.Print("[")
	HiYellow.Print(tag)
	White.Print("]	")
}

func PrintWithTag(tag string, msg string) {
	PrintTag(tag)
	White.Println(msg)
}

func StaticClearLine() {
	fmt.Print(`

`)
}

func PrintHeader(version string) {
	yellow := color.New(color.FgRed)

	yellow.Println(`
   ▄████████    ▄████████    ▄████████    ▄████████ 
  ███    ███   ███    ███   ███    ███   ███    ███ 
  ███    ███   ███    ███   ███    █▀    ███    █▀  
  ███    ███  ▄███▄▄▄▄██▀  ▄███▄▄▄       ███        
▀███████████ ▀▀███▀▀▀▀▀   ▀▀███▀▀▀     ▀███████████ 
  ███    ███ ▀███████████   ███    █▄           ███ 
  ███    ███   ███    ███   ███    ███    ▄█    ███ 
  ███    █▀    ███    ███   ██████████  ▄████████▀  
               ███    ███                           
`)
	White.Printf("			version %s", version)
	yellow.Println(`

`)
}

func Clear() {
	clear.CallClear()
}
