package main

import (
	"fmt"

	"github.com/Shakkuuu/typinggame-golang/typinggame"
)

func main() {
	if err := typinggame.FileLoad(); err != nil {
		fmt.Println(err)
	}

	typinggame.Typing()
}
