package main

import "fmt"

func plusten(a *int) {
	*a += 10
}

func main() {
	num := 1
	plusten(&num)

	fmt.Println(num)
}
