package main

import (
	"fmt"

	"github.com/jwalton/gchalk/pkg/ansistyles"
	"github.com/jwalton/go-supportscolor"
)

///func rgb(a, x, y int) uint8 {
//	s := ansistyles.RGBToAnsi256(uint8(a), uint8(x), uint8(y))
//return ansistyles.Ansi256(s)
//	return s
//}

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

func bold(word string) string {
	return "\033[1m" + word + "\033[0m"

}
func format(c, word string) string {

	return c + word + reset
}
func main() {
	supportscolor.Stdout()

	fmt.Printf("color: %s\n", format(aqua, "aqua"))
	fmt.Printf("color: %s\n", format(black, "black"))
	fmt.Printf("color: %s\n", format(blue, "blue"))
	fmt.Printf("color: %s\n", format(fuchsia, "fuchsia"))
	fmt.Printf("color: %s\n", format(gray, "gray"))
	fmt.Printf("color: %s\n", format(green, "green"))
	fmt.Printf("color: %s\n", format(lime, "lime"))
	fmt.Printf("color: %s\n", format(maroon, "maroon"))
	fmt.Printf("color: %s\n", format(navy, "navy"))
	fmt.Printf("color: %s\n", format(olive, "olive"))
	fmt.Printf("color: %s\n", format(purple, "purple"))
	fmt.Printf("color: %s\n", format(red, "red"))
	fmt.Printf("color: %s\n", format(silver, "silver"))
	fmt.Printf("color: %s\n", format(teal, "teal"))
	fmt.Printf("color: %s\n", format(white, "white \033[1mwhite\033[0m"))
	fmt.Printf("color: %s\n", format(yellow, "yellow "+bold("yellow")))

	fmt.Printf("color: %s", format(teal, "kek "+bold("kek")))

}
