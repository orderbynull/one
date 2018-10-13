package core

import (
	"fmt"
	"os"
)

const errorTmpl = "Error: %s\n"

// Error prints error message to stdout
func Error(err error)  {
	fmt.Printf(errorTmpl, err.Error())
}

func Info(str string)  {
	fmt.Println(str)
}

func ErrorAndExit(err error)  {
	Error(err)
	os.Exit(1)
}