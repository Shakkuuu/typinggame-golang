package main

import "fmt"

func main() {
	var w string
	word := "あいうえお"
	for {
		fmt.Print(word)
		fmt.Print(">")
		fmt.Scan(&w)
		if word == w {
			fmt.Println("OK")
			break
		}
	}
}
