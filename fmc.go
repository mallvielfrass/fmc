package fmc

import (
	"fmt"
	"reflect"
	"strconv"

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
	//maroon  = rgb(128, 0, 0)
	//navy    = rgb(0, 0, 128)
	//olive   = rgb(128, 128, 0)
	//purple  = rgb(128, 0, 128)
	//red     = rgb(255, 0, 0)
	//silver  = rgb(192, 192, 192)
	//teal    = rgb(0, 128, 128)
	//white   = rgb(255, 255, 255)
	//yellow  = rgb(255, 255, 0)
	aqua    = ansistyles.Ansi256(51)  //b
	black   = ansistyles.Ansi256(16)  //0
	blue    = ansistyles.Ansi256(21)  //B
	fuchsia = ansistyles.Ansi256(201) //f
	gray    = ansistyles.Ansi256(244) //1
	green   = ansistyles.Ansi256(34)  //G
	lime    = ansistyles.Ansi256(46)  //g
	maroon  = ansistyles.Ansi256(124) //m
	navy    = ansistyles.Ansi256(19)  //n
	olive   = ansistyles.Ansi256(142) //o
	purple  = ansistyles.Ansi256(127) //p
	red     = ansistyles.Ansi256(196) //r
	silver  = ansistyles.Ansi256(250) //s
	teal    = ansistyles.Ansi256(37)  //t
	white   = ansistyles.Ansi256(231) //w
	yellow  = ansistyles.Ansi256(226) //y
	reset   = "\x1b[39m"
)

func format(c, word string) string {
	return "\033[0m" + reset + c + word //+ reset
}
func bformat(c, word string) string {
	return "\033[0m" + reset + c + "\033[1m" + word //+ reset //+ "\033[0m"
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
func toString(a interface{}) interface{} {
	switch a.(type) {
	case int:
		return strconv.Itoa(a.(int))
		// v is an int here, so e.g. v + 1 is possible.
		//fmt.Printf("Integer: %v", v)
	case float64:
		return strconv.FormatInt(int64(a.(int64)), 10)
		// v is a float64 here, so e.g. v + 1.0 is possible.
	//	fmt.Printf("Float64: %v", v)
	case string:
		return a
		// v is a string here, so e.g. v + " Yeah!" is possible.
	//	fmt.Printf("String: %v", v)
	default:
		return fmt.Sprintf("%v", a)
		// And here I'm feeling dumb. ;)
		//	fmt.Printf("I don't know, ask stackoverflow.")
	}
	//return a
}
func tab(n int) string {
	var tab string
	for i := 0; i < n; i++ {
		tab = tab + "  "
	}
	return tab
}
func parse(d interface{}, n int) string {
	v := reflect.ValueOf(d)
	t := v.Type()
	var msg string
	for i := 0; i < t.NumField(); i++ {
		ty := t.Field(i).Type.Kind()
		if ty != reflect.Struct {
			name := t.Field(i).Name
			value := v.FieldByName(name).Interface()
			msg = msg + fmt.Sprintf("%s#bbt%s#tbt: #gbt%s\n", tab(n), toString(name), toString(value))
		} else {
			name := t.Field(i).Name
			msg = msg + tab(n) + "#bbt" + name + " #tbt: #ybt{\n" + parse(v.FieldByName(t.Field(i).Name).Interface(), n+1) + tab(n) + "#ybt}\n"
		}
	}
	return msg
}
func SprintStruct(d interface{}) string {
	f := reflect.ValueOf(d).Type()
	if f.Kind() != reflect.Struct {
		return fmt.Sprintf("%v", d)
	}
	v := f.Name()
	return "struct #gbt" + v + " #tbt[\n" + parse(d, 1) + "#tbt]"

}
func PrintStruct(d interface{}) {
	Printfln(SprintStruct(d))
}
