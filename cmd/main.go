package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("Hellow Go!!")
	debug_str := "Hellow Go!!"
	spew.Dump(debug_str)
}
