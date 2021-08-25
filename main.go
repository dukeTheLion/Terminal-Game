package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func()

var lines = 15
var columns = 40

func init() {
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
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func boxDrawer(text string, x int, y int) {
	var outputString = ""

	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			switch {
			case i == 0 && j == 0:
				outputString = "┏"
				print(outputString)

			case j > 0 && j < x-1 && (i == 0 || i == y-1):
				outputString = "━"
				print(outputString)

			case i == 0 && j == x-1:
				outputString = "┓"
				print(outputString)

			case i == y-1 && j == 0:
				outputString = "┗"
				print(outputString)

			case i == y-1 && j == x-1:
				outputString = "┛"
				print(outputString)

			default:
				if j == 0 || j == x-1 {
					outputString = "┃"
				} else {
					outputString = " "
				}

				print(outputString)
			}
		}
		print("\n")
	}

}

func drawer(text string) {
	boxDrawer(text, columns, lines)
	boxDrawer(text, columns, 5)
}

func loop() {
	for {
		CallClear()

		drawer("")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
		time.Sleep(500 * time.Millisecond)
		CallClear()
	}
}

func main() {
	loop()
	CallClear()
	fmt.Println("End Game!")
}
