package checkerror

import (
	"fmt"
	"log"

)

// A general go check error function
/* EXAMPLES:
CheckError(err)
CheckError(err, "Fatal")  //Specifies log style
CheckError(err, "Fatal", "Read config.yaml file")    //Specifies log style and output message
CheckError(err, "Read config file")      //Only specifies log output message
*/

func With(err error, args ...string) {

	if err == nil { // no error
		return
	}
	tag := ""
	logStyle := ""

	for _, arg := range args {
		switch arg {
		case "fatal", "Fatal", "FATAL":
			logStyle = "fatal"
		case "panic", "Panic", "PANIC":
			logStyle = "panic"
		default:
			tag = arg

		}
	}

	var msg string
	if tag == "" {
		msg = fmt.Sprintf("%v", err)
	} else {
		msg = fmt.Sprintf("[%s] %v", tag, err)
	}

	//Print out
	switch logStyle {
	case "fatal":
		log.Fatal(msg)
	case "panic":
		log.Panic(msg)
	default:
		log.Println(msg)
	}

}
