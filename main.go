package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
	"unicode/utf8"
)

var clear map[string]func()

const lines = 15
const columns = 40

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

func check(e error) {
	if e != nil {
		panic(e)
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

func boxDrawer(text []string, x int, y int) {
	var outputString = ""
	var byteText = make([][]byte, len(text))

	for i, s := range text {
		byteText[i] = []byte(s)
	}

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
					print(outputString)
				} else {
					var a, s = utf8.DecodeRune(byteText[i-1])
					fmt.Printf("%c", a)

					byteText[i-1] = byteText[i-1][s:]
				}
			}
		}
		print("\n")
	}

}

func drawer(text []string) {
	CallClear()
	boxDrawer(text, columns, lines)
	boxDrawer(text, columns, 5)
}

func loop() {

	dat, err := ioutil.ReadFile("sample.txt")
	check(err)

	var text = strings.Split(string(dat), "\\\n")
	var output = make([]string, lines-2)

	for {

		for _, i2 := range text {
			output = strings.Split(i2, "\n")
			time.Sleep(100 * time.Millisecond)

			CallClear()
			drawer(output)
		}

	}

	/*var text = make([]string, lines-2)
	var val = []string{"◂", "▾", "▸", "▴"}

	for {
		CallClear()

		for k := 0; k < columns; k++ {
			for i := 0; i < lines-2; i++ {
				var temp string
				for j := 0; j < columns-2; j++ {
					if k == j {
						temp += val[2]
					} else if k > j {
						temp += "█"
					} else {
						temp += " "
					}
				}

				text[i] = temp
			}
			drawer(text)
			time.Sleep(50 * time.Millisecond)
		}
		for k := 0; k < columns; k++ {
			for i := 0; i < lines-2; i++ {
				var temp string
				for j := 0; j < columns-2; j++ {
					if k == j {
						temp += " "
					} else if k > j {
						temp += " "
					} else {
						temp += "█"
					}
				}

				text[i] = temp
			}
			drawer(text)
			time.Sleep(50 * time.Millisecond)
		}

	}*/
}

func main() {
	loop()
	CallClear()
	fmt.Println("End Game!")
}
