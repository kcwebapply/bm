package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// GetTerminalInput return std.in
func GetTerminalInput(text string) (string, error) {
	var err error
	fmt.Print(">> " + text + ":")
	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		err = errors.New("get input error")
	}
	return scanner.Text(), err
}

// TextCounter counts input string length
func TextCounter(text string) int {
	textCounter := 0
	befPos := 0
	for pos := range text {
		if pos-befPos == 3 {
			textCounter += 2 // to treat japanese character as 2byte.
			befPos = pos
		} else {
			textCounter++
			befPos = pos
		}
	}
	return textCounter
}
