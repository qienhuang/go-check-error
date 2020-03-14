# A general go check error function


```
EXAMPLES:
	checkError.With(err)
	checkError.With(err, "Fatal")                        //Specifies log style
	checkError.With(err, "Fatal", "READ CONFIG FILE")     //Specifies log style and tag
	checkError.With(err, "READ CONFIG FILE")             //Only specifies tag
	
OUTPUT:
2020/01/14 17:13:15 [READ CONFIG FILE] open filepath: The system cannot find the file specified.
```
