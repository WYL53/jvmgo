package main

import "fmt"

func TestMain()  {
	a := int8(-1)
	fmt.Printf("%08b\n",uint8(a))
	b := int8(a)
	b >>= 1
	fmt.Printf("%08b\n",uint8(b))
}
