# A general check error function for Golang

```
EXAMPLES:
	checkError.With(err)
	checkError.With(err, "Fatal")                        //Specifies log style
	checkError.With(err, "Fatal", "READ CONFIG FILE")     //Specifies log style and tag
	checkError.With(err, "READ CONFIG FILE")             //Only specifies tag
	
OUTPUT:
2020/01/14 17:13:15 [READ CONFIG FILE] open filepath: The system cannot find the file specified.
```

#### INSTALL PACKAGE:
```
go get github.com/qienhuang/go-check-error
```

#### USE:
```go
package main

import(
	"io/ioutil"
	"github.com/qienhuang/go-check-error"
)

func main() {
	_,err:=ioutil.ReadFile("config.ini")
	checkerror.With(err, "Fatal", "READ CONFIG FILE")
}
```
#### OUTPUT:
2020/01/14 17:13:15 [READ CONFIG FILE] open config.ini: The system cannot find the file specified.

