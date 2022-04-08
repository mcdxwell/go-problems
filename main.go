package main

import (
	"fmt"
	
	"github.com/mcdxwell/go-problems/adder"
)


func main() {
	
      x := adder.Adder(1, 2)
      z := adder.Adder("hello", "world")
	fmt.Println(x, z)
}
