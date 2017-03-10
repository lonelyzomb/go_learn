package main

import (
	"fmt"
	"strconv"

	"github.com/malashin/term/congo"
)

func main() {
	congo.Init()
	defer congo.Close()

	i := 0
	var esc rune = 27

loop:
	for {
		if congo.KeyPressed() {
			ch := congo.ReadKey()
			if ch == esc {
				fmt.Println("\n\x1b[31;1mExiting...\x1b[0m")
				break loop
			} else {
				fmt.Printf("\n%c %v\n", ch, ch)
			}
			i = 0
		} else {
			fmt.Print("\r\x1b[30;1m" + strconv.Itoa(i) + "\x1b[0m")
		}
		i++
	}
}
