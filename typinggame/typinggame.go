package typinggame

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

var Words []string

func Typing(ti int) {
	var point int
	var qword string
	var inputword string
	fmt.Printf("タイピングゲームをはじめます。制限時間は%d 秒です。\n", ti)
	time.Sleep(3 * time.Second)
	fmt.Println("スタート!")
	timeout := time.After(time.Duration(ti) * time.Second)
	for tf := true; tf && point < len(Words); {
		qword = Words[point]
		fmt.Print(qword)
		fmt.Print(">")
		select {
		case <-timeout:
			fmt.Printf("\nタイムオーバー\n")
			tf = false
		default:
			fmt.Scan(&inputword)
			if qword == inputword {
				fmt.Println("OK")
				point++
			}
		}
	}
	fmt.Printf("%d点中%d点です。", len(Words), point)
}

func FileLoad(filename string) error {
	t, err := os.Open(filename)
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
