package fmc

import (
	"fmt"
	"log"
	"os"
)

var std = log.New(os.Stderr, "", log.LstdFlags)

func errorHandlePrepare(e error, a ...interface{}) string {
	//if e != nil {
	path, line, fn := WhoCallerErr()
	if a != nil {
		return fmt.Sprintf(Sprint("#rbtError #RRR[#gbt%s#RRR:#gbt%s #ybt%s()#RRR]: \n\tDescription: #RRR[#ybt %s#RRR] \n\t#RRRMessage: #ybt%s"), path, line, fn, fmt.Sprintf("%v", a[0]), e.Error())
	}
	return fmt.Sprintf(Sprint("#rbtError#RRR[#gbt%s#RRR:#gbt%s #ybt%s()#RRR]: \n\t#RRRMessage: #ybt%s"), path, line, fn, e.Error())
	//}
	//return ""
}
func ErrorHandle(e error, a ...interface{}) {
	if e != nil {
		var msg string
		if a != nil {
			msg = errorHandlePrepare(e, a[0])
		} else {
			//	fmt.Println("else")
			msg = errorHandlePrepare(e)
		}
		fmt.Println(msg)

	}

}
func ErrorHandleFatal(e error, a ...interface{}) {
	if e != nil {
		var msg string
		if a != nil {
			msg = errorHandlePrepare(e, a[0])
		} else {
			msg = errorHandlePrepare(e)
		}
		std.Output(2, msg)
		os.Exit(1)
	}
}
