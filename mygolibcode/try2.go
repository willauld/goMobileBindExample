package main

import (
	"fmt"

	"github.com/willauld/temp2/mygolib"
)

type T struct{}

var junk T

func (t T) CJUpdate(x string) {
	fmt.Printf("CJUpdate received: %s\n", x)
}

func main() {
	var str string
	var err error
	fmt.Printf("\n")
	str, err = mygolib.Callgo("Java String")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("Return string from calling java:\n\t%s\n", str)
	f := mygolib.NewUpdater(junk)
	str, err = mygolib.Callgocalljava(f)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Print(str)
	fmt.Printf("\n\n")
}
