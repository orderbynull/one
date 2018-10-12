package core

import "fmt"

const errorTmpl = "Error: %s\n"

// Error prints error message to stdout
func Error(err error)  {
	fmt.Printf(errorTmpl, err.Error())
}