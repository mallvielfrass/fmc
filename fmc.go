package fmc

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/mattn/go-isatty"
)

// Base attributes
const (
	reset attribute = iota
	bold
	faint
	italic
	underline
	blinkSlow
	blinkRapid
	reverseVideo
	concealed
	crossedOut
)

// Foreground text colors
const (
	fgBlack attribute = iota + 30
	fgRed
	fgGreen
	fgYellow
	fgBlue
	fgMagenta
	fgCyan
	fgWhite
)

const escape = "\x1b"

type attribute int
type color struct {
	params  []attribute
	nocolor *bool
}

func (c *color) add(value ...attribute) *color {
	c.params = append(c.params, value...)
	return c
}

func new(value ...attribute) *color {
	c := &color{params: make([]attribute, 0)}
	c.add(value...)
	return c
}
func (c *color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

func (c *color) isNocolorSet() bool {
	// check first if we have user setted action
	if c.nocolor != nil {
		return *c.nocolor
	}

	// if not return the global option, which is disabled by default
	return nocolor
}
func (c *color) wrap(s string) string {
	if c.isNocolorSet() {
		return s
	}

	return c.format() + s + c.unformat()
}
func (c *color) sprintfFunc() func(format string, a ...interface{}) string {
	return func(format string, a ...interface{}) string {
		return c.wrap(fmt.Sprintf(format, a...))
	}
}
func (c *color) sprintfFuncS() func(format string) string {
	return func(format string) string {

		color := fmt.Sprintf("%s[%sm", escape, c.sequence())
		//d:=
		return color + format + fmt.Sprintf("%s[%s", escape, c.unformat())
		//return c.wrap(fmt.Sprintf(format, a...))
	}
	//return func(format string) string {

	//}
}
func (c *color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}

func (c *color) unformat() string {
	return fmt.Sprintf("%s[%dm", escape, reset)
}

func checkSharp(format string) bool {
	if strings.Contains(format, "#") == true {
		return true
	}
	return false

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
		world = yst(mess)
	case "ybt":
		world = ybt(mess)
	case "rst":
		world = rst(mess)
	case "rbt":
		world = rbt(mess)
	case "gst":
		world = gst(mess)
	case "gbt":
		world = gbt(mess)
	case "bst":
		world = bst(mess)
	case "bbt":
		world = bbt(mess)
	case "wst":
		world = wst(mess)
	case "wbt":
		world = wbt(mess)
	default:
		world = mess
	}
	//fmt.Printf("wordset:%s\n", world)
	return world
}

var (
	yst     = new(fgYellow).sprintfFuncS()
	ybt     = new(fgYellow, bold).sprintfFuncS()
	rst     = new(fgRed).sprintfFuncS()
	rbt     = new(fgRed, bold).sprintfFuncS()
	gst     = new(fgGreen).sprintfFuncS()
	gbt     = new(fgGreen, bold).sprintfFuncS()
	bst     = new(fgBlue).sprintfFuncS()
	bbt     = new(fgBlue, bold).sprintfFuncS()
	wst     = new(fgWhite).sprintfFuncS()
	wbt     = new(fgWhite, bold).sprintfFuncS()
	nocolor = os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))
)

//Print interface
func Print(form string) {
	fmt.Print(Sprint(form))
}

//Println string
func Println(form string) {
	fmt.Print(Sprint(form))
	fmt.Printf("\n")
}

//Printf interface
func Printf(format string, a ...interface{}) {
	pr := Sprint(format)
	t := fmt.Sprintf(pr, a...)
	fmt.Printf(t)

}

//Printfln +and line
func Printfln(format string, a ...interface{}) {
	pr := Sprint(format)
	t := fmt.Sprintf(pr, a...)
	fmt.Printf(t)
	fmt.Printf("\n")
}

//Sprint prepare string to print or fmt.Sprintf(format, a...)
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

//PrintfT interface
func printfT(format string, a ...interface{}) {
	if checkSharp(format) != true {
		fmt.Print(format)
		for _, v := range a {

			//fmt.Println(fmt.Sprintf("%s[%sm fff", escape, d.sequence()))
			fmt.Print(v)
		}
		return
	}

	//	sep := strings.Split(format, "#")
	d := new(fgRed, bold)
	frmt := fmt.Sprintf("%s[%sm%s", escape, d.sequence(), format)
	//for _, v := range a {
	alen := len(a)
	per := strings.Split(frmt, "%")
	perc := len(per) - 1
	if perc < alen {
		fmt.Printf("%s %s\n", rbt("[error] Not enough separators in format string:"), format)
		return
	}

	fmt.Printf("a:%d | perc: %d\n", alen, perc)
	//letter := fmt.Sprintf(frmt)
	//mstring:=
	for i := 0; i < len(per); i++ {
		typeW := ""
		if i == 0 {
			typeW = "N"
		} else {
			typeW = string((per[i])[0])
		}

		fmt.Printf("per:%s, type:%s \n", per[i], typeW)
	}
	//fmt.Print(letter)

	//}
}
