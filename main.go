package main

import (
	"fmt"

	"github.com/Shakkuuu/typinggame-golang/typinggame"
)

func main() {
	filename := "words.txt"
	if err := typinggame.FileLoad(filename); err != nil {
		fmt.Println(err)
	}

	time_limit := 20
	typinggame.Typing(time_limit)
}
