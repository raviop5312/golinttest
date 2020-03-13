package main

import "fmt"

type A struct {
	IntArr1 []int32
}
func main() {
	a := A{}
	fmt.Println(a.IntArr1[0])
}