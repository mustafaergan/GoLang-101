package main

import "fmt"

func main() {
	fmt.Println("first")
	var i int = 0
	for i <= 3{
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("seconds")
	i  = 0
	for ;i <= 3 ; i++ {
		fmt.Println(i)
	}

	fmt.Println("third")
	for i = 0 ;i <= 3 ; i++ {
		fmt.Println(i)
	}

	fmt.Println("break")
	for i = 0 ;i <= 3 ; i++ {
		fmt.Println(i)
		break
	}

	fmt.Println("continue")
	for i = 0 ;i <= 3 ; i++ {
		fmt.Println(i)
		continue
	}

}
