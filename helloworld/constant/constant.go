package main

import "fmt"

const s string = "constant"

func main() {
	//s = "test"
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(d)


}
