package checkerror

import (
	"fmt"
	"log"
	"runtime"
)

// A general go check error function
/* EXAMPLES:
checkError.With(err)
checkError.With(err, "Fatal")  //Specifies log style
checkError.With(err, "Fatal", "Read config.yaml file")    //Specifies log style and output message
checkError.With(err, "Read config file")      //Only specifies log output message
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

	//If doesn't specify a tag, then use the caller function name instead
	if tag == "" {
		fpcs := make([]uintptr, 1)
		n := runtime.Callers(3, fpcs)
		if n != 0 {
			fun := runtime.FuncForPC(fpcs[0] - 1)
			if fun != nil {
				tag = "From " + fun.Name() + ": "
			}
		}

	}

	//Print out
	msg := fmt.Sprintf("[%s] %s", tag, err)
	switch logStyle {
	case "fatal":
		log.Fatal(msg)
	case "panic":
		log.Panic(msg)
	default:
		log.Println(msg)
	}

}
