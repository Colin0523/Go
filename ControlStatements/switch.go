package main

import (
	"fmt"
	"runtime"
)

func main(){
	fmt.Println("Go runs on")
	switch os := runtime.GOOS; os{
		case "darwin":
			fmt.Println("OS X.")
		case "inux":
			fmt.Println("Linux.")
		default:
			//freebsd, openbsd
			//plan9, windows
			fmt.Println("%s", os)
	}
}