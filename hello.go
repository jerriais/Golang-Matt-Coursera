package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	var x int = 1
	var ip *int
	ip = &x
	fmt.Println(*ip)
}
