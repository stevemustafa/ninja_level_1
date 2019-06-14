package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	//fmt.Println("X: ", x)
	//fmt.Println("Y: ", y)
	//fmt.Println("Z: ", z)
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
}
