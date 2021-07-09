package fmc

import (
	"fmt"
	"regexp"
	"strings"
)

func screenCheck(word string) bool {
	matched, err := regexp.MatchString(`!#`, word)
	if err != nil {
		fmt.Println(err)
	}
	return matched
}
func formatWithScreen(form string) string {
	items := strings.Split(form, "!#")
	var msg string
	msg = formatWithoutScreen(items[0])
	ln := len(items)
	if 1 <= ln {
		for i := 1; i < ln; i++ {
			msg = msg + "#" + formatWithoutScreen(items[i])
		}
	} else {
		msg = msg + "#"
	}

	return msg
}
func coloring(color, mess string) string {
	var word string
	switch color {
	//aqua
	case "bst":
		word = format(aqua, mess)
	case "bbt":
		word = bformat(aqua, mess)
	//black
	case "0st":
		word = format(black, mess)
	case "0bt":
		word = bformat(black, mess)
	//blue
	case "Bst":
		word = format(blue, mess)
	case "Bbt":
		word = bformat(blue, mess)
	//fuchsia
	case "fst":
		word = format(fuchsia, mess)
	case "fbt":
		word = bformat(fuchsia, mess)
	//gray
	case "1st":
		word = format(gray, mess)
	case "1bt":
		word = bformat(gray, mess)
	//green
	case "Gst":
		word = format(green, mess)
	case "Gbt":
		word = bformat(green, mess)
	//lime
	case "gst":
		word = format(lime, mess)
	case "gbt":
		word = bformat(lime, mess)
	//maroon
	case "mst":
		word = format(maroon, mess)
	case "mbt":
		word = bformat(maroon, mess)
	//navy
	case "nst":
		word = format(navy, mess)
	case "nbt":
		word = bformat(navy, mess)
	//olive
	case "ost":
		word = format(olive, mess)
	case "obt":
		word = bformat(olive, mess)
	//purple
	case "pst":
		word = format(purple, mess)
	case "pbt":
		word = bformat(purple, mess)
	//red
	case "rst":
		word = format(red, mess)
	case "rbt":
		word = bformat(red, mess)
	//silver
	case "sst":
		word = format(silver, mess)
	case "sbt":
		word = bformat(silver, mess)
	//teal
	case "tst":
		word = format(teal, mess)
	case "tbt":
		word = bformat(teal, mess)
	//white
	case "wst":
		word = format(white, mess)
	case "wbt":
		word = bformat(white, mess)
	//yellow
	case "yst":
		word = format(yellow, mess)
	case "ybt":
		word = bformat(yellow, mess)
	//rest
	case "RRR":
		word = "\033[0m" + reset + mess
	default:
		word = mess
	}
	return word
}
func formatWithoutScreen(form string) string {
	items := strings.Split(form, "#")
	ilen := len(items)
	if ilen < 0 {
		return form
	}
	var msg string
	msg = items[0]
	for i := 1; i < ilen; i++ {
		item := items[i]
		if 4 <= len(item) {
			color := item[0:3]
			text := item[3:]
			if text[:1] == " " {
				text = item[4:]
			}
			msg = msg + coloring(color, text)
		}

	}
	//fmt.Printf("msg: %s\n", msg)
	return msg
}
func Sprint(form string) string {
	if screenCheck(form) {
		return formatWithScreen(form + "\033[0m" + reset)
	}
	return formatWithoutScreen(form + "\033[0m" + reset)

}
