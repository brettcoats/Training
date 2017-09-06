package main

import "fmt"
var x = 500
var y = 1000


func bob(b int)  int{
	return b + 10
}

func main() {
	var z int
	if  z = bob(10); z > 100 {fmt.Println(bob(y))

	}
	fmt.Println(bob(x))
}
