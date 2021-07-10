package fmc

import (
	"runtime"
	"strconv"
	"strings"
)

//Caller log run function
func Caller() {
	pc := make([]uintptr, 40)
	n := runtime.Callers(0, pc)

	// Printfln("#rbtn= #gbt%d", n)
	pc = pc[0:n] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)
	i := 0
	for {

		frame, more := frames.Next()

		name, line, file := frame.Func, frame.Line, frame.File
		//	fmt.Println(name)
		nm := ""
		if name == nil {
			//	fmt.Println("name==nil")
			f := runtime.FuncForPC(frame.PC)
			//file, line := f.FileLine(frame.PC)
			//fmt.Printf("%s:%d %s\n", file, line, f.Name())
			nm = f.Name()
			//Printfln("#gbt%s#wbt:#gbt%d #ybt%s", file, line, f.Name())
		} else {
			nm = name.Name()

		}

		if strings.Contains(nm, "fmc.Caller") || strings.Contains(nm, "runtime.Callers") || strings.Contains(nm, "runtime.goexit") || strings.Contains(nm, "runtime.main") {

		} else {
			if i == 0 {
				Printf("#wbtRun> ")
				i++
			} else {
				Printf("\t#rbtFrom> ")
				i++
			}
			Printfln("#gbt%s#wbt:#gbt%d #ybt%s", file, line, nm)
		}

		//	}

		if !more {
			break
		}
	}
}

//WhoCallerIs calling func name
func WhoCallerIs() string {
	pc := make([]uintptr, 40)
	n := runtime.Callers(0, pc)

	// Printfln("#rbtn= #gbt%d", n)
	pc = pc[2 : n-2] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)

	frame, _ := frames.Next()

	//name, line, file := frame.Func, frame.Line, frame.File
	//	fmt.Println(name)
	if frame.Func == nil {
		//	fmt.Println("name==nil")
		f := runtime.FuncForPC(frame.PC)
		//file, line := f.FileLine(frame.PC)
		//fmt.Printf("%s:%d %s\n", file, line, f.Name())
		//	Printfln("#gbt%s#wbt:#gbt%d #ybt%s", file, line, f.Name())
		return f.Name()
	}
	//	Printfln("#gbt%s#wbt:#gbt%d #ybt%s", file, line, name.Name())
	return frame.Func.Name()

}

//WhoCallerIs calling func name
func WhoCallerErr() (string, string, string) {
	pc := make([]uintptr, 40)
	n := runtime.Callers(0, pc)
	//var root []uintptr
	//root[0] = pc[0]
	// Printfln("#rbtn= #gbt%d", n)
	pc = pc[4 : n-2] // pass only valid pcs to runtime.CallersFrames
	frames := runtime.CallersFrames(pc)
	//runtime.CallersFrames(root)
	//fmt.Println(frames)
	frame, _ := frames.Next()

	//name, line, file := frame.Func, frame.Line, frame.File
	//	fmt.Println(name)
	if frame.Func == nil {
		//	fmt.Println("name==nil")
		f := runtime.FuncForPC(frame.PC)
		file, line := f.FileLine(frame.PC)
		//fmt.Printf("%s:%d %s\n", file, line, f.Name())
		//	Printfln("#gbt%s#wbt:#gbt%d #ybt%s", file, line, f.Name())
		return file, strconv.Itoa(line), f.Name()
	}

	//	Printfln("#gbt%s#wbt:#gbt%d #ybt%s", file, line, name.Name())
	return frame.File, strconv.Itoa(frame.Line), frame.Func.Name()

}
