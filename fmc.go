package fmc

import (
	"fmt"
	"strings"

	"github.com/jwalton/gchalk/pkg/ansistyles"
	"github.com/jwalton/go-supportscolor"
)

var (
	//aqua    = rgb(0, 255, 255)
	//black   = rgb(0, 0, 0)
	//blue    = rgb(0, 0, 255)
	//fuchsia = rgb(255, 0, 255)
	//gray    = rgb(128, 128, 128)
	//green   = rgb(0, 128, 0)
	//lime    = rgb(0, 255, 0)
	//	maroon  = rgb(128, 0, 0)
	//navy    = rgb(0, 0, 128)
	//olive   = rgb(128, 128, 0)
	//purple  = rgb(128, 0, 128)
	//red     = rgb(255, 0, 0)
	//silver  = rgb(192, 192, 192)
	//teal    = rgb(0, 128, 128)
	//white   = rgb(255, 255, 255)
	//yellow  = rgb(255, 255, 0)
	aqua    = ansistyles.Ansi256(51)
	black   = ansistyles.Ansi256(16)
	blue    = ansistyles.Ansi256(21)
	fuchsia = ansistyles.Ansi256(201)
	gray    = ansistyles.Ansi256(244)
	green   = ansistyles.Ansi256(34)
	lime    = ansistyles.Ansi256(46)
	maroon  = ansistyles.Ansi256(124)
	navy    = ansistyles.Ansi256(19)
	olive   = ansistyles.Ansi256(142)
	purple  = ansistyles.Ansi256(127)
	red     = ansistyles.Ansi256(196)
	silver  = ansistyles.Ansi256(250)
	teal    = ansistyles.Ansi256(37)
	white   = ansistyles.Ansi256(231)
	yellow  = ansistyles.Ansi256(226)
	reset   = "\x1b[39m"
)

func format(c, word string) string {
	return c + word + reset
}
func bformat(c, word string) string {
	return c + "\033[1m" + word + "\033[0m" + reset
}
func setcolor(item string) string {
	color := ""
	mess := ""
	itemLen := len(item)
	//	fmt.Printf("item:%s, itemLen: '%d'\n", item, itemLen)
	if 3 < itemLen {
		//		fmt.Printf("item[3]: '%s'\n", string(item[3]))
		color = item[0:3]
		if 4 <= itemLen {
			if string(item[3]) == " " {
				//				fmt.Println("item[3]")
				if 4 < len(item) {
					mess = item[4:]
				} else {
					mess = item[3:]
				}

			} else {
				//			fmt.Println("!item[3]")
				mess = item[3:]
			}

		} else {
			mess = item[3:]
		}
	} else {

		color = ""
		mess = ""
	}
	//	fmt.Printf("color: %s\n", color)
	//	fmt.Printf("mess: %s\n", mess)
	world := "	"
	switch color {
	case "yst":
		world = format(yellow, mess)
	case "ybt":
		world = bformat(yellow, mess)
	case "rst":
		world = format(red, mess)
	case "rbt":
		world = bformat(red, mess)
	case "gst":
		world = format(lime, mess)
	case "gbt":
		world = bformat(lime, mess)
	case "bst":
		world = format(aqua, mess)
	case "bbt":
		world = bformat(aqua, mess)
	case "wst":
		world = format(white, mess)
	case "wbt":
		world = bformat(white, mess)
	default:
		world = mess
	}
	//fmt.Printf("wordset:%s\n", world)
	return world
}
func Sprint(form string) string {
	items := strings.Split(form, "#")
	ilen := len(items)
	msg := ""
	for i := 0; i < ilen; i++ {
		item := items[i]
		if i == 0 {
			msg = msg + item
		} else {
			if 0 < len(items[i-1]) {
				lastSymbol := string((items[i-1])[len(items[i-1])-1])
				if lastSymbol != "!" {
					rt := setcolor(item)
					format(aqua, "aqua")
					msg = msg + rt
				} else {

					msg = msg + item
				}
				//fmt.Println("last:", lastSymbol)
			} else {
				rt := setcolor(item)
				msg = msg + rt
			}
		}

	}
	//fmt.Printf("msg: %s\n", msg)
	return msg
}

//Println string
func Println(form string) {
	//fmt.Print()
	fmt.Printf("%s\n", Sprint(form))
}

//Printfln +and line
func Printfln(wformat string, a ...interface{}) {
	fmt.Printf(Sprint(wformat), a...)
	fmt.Printf("\n")
}

//Printf interface
func Printf(wformat string, a ...interface{}) {
	//pr := Sprint(wformat)
	//t := fmt.Sprintf(pr, a...)
	fmt.Printf(Sprint(wformat), a...)
	//fmt.Printf("color: %s\n", format(aqua, "aqua"))
}
func init() {
	supportscolor.Stdout()
}
