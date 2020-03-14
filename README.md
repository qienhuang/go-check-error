# A general go check error function


```
EXAMPLES:
	checkError.With(err)
	checkError.With(err, "Fatal")                        //Specifies log style
	checkError.With(err, "Fatal", "Read file error")     //Specifies log style and tag
	checkError.With(err, "Read config file")             //Only specifies tag
	
OUTPUT:
2020/01/14 17:13:15 [Read config file] open filepath: The system cannot find the file specified.
```
