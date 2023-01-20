package typinggame

import (
	"bufio"
	"fmt"
	"os"
)

var Words []string

func Typing() {
	var point int
	var qword string
	var inputword string
	for point < len(Words) {
		qword = Words[point]
		fmt.Print(qword)
		fmt.Print(">")
		fmt.Scan(&inputword)
		if qword == inputword {
			fmt.Println("OK")
			point++
		}
	}
}

func FileLoad() error {
	t, err := os.Open("./words.txt")
	if err != nil {
		return err
	}
	defer t.Close()
	scanner := bufio.NewScanner(t)
	for scanner.Scan() {
		Words = append(Words, scanner.Text())
	}
	return nil
}
