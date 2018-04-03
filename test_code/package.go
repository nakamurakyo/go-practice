package main

import (
	"log"
	"./test-lib"
)

func main() {
	sample.Foo()
	
  // x := sample.Mydata{}
	x := sample.Mydata{ Num1: 2 }
	
  // x.Num1 = 1
	x.Str1  = "test"

	log.Printf("x.Num1=%d, x.Str1=%s\n", x.Num1, x.Str1)
}