package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var words []string
	t, err := os.Open("./words.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer t.Close()
	scanner := bufio.NewScanner(t)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	var w string
	var point int
	for point < len(words) {
		wo := words[point]
		fmt.Print(wo)
		fmt.Print(">")
		fmt.Scan(&w)
		if wo == w {
			fmt.Println("OK")
			point++
		}
	}
}
