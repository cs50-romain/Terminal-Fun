package shell

import (
	"fmt"
	"os"
)

const (
	CLEAR_SCREEN = "\u001b[2J"
	CURSOR_HOME = "\u001b[H"
	CURSOR_LEFT = "\u001b[1000D"
	cursorrt = "\u001b[C"
	cursorlf = "\u001b[D"
	CTRL_C = 3
)

func Run() error {
	fmt.Print(CURSOR_LEFT)
	fmt.Print("> ")
	var line_buffer = make([]byte, 64) 
	for {
		// Simple read line
		var buf [3]byte
		os.Stdin.Read(buf[:])
		fmt.Print(string(buf[0]))
		if buf[0] == '\u001b' {
			if buf[1] == '[' {
				if buf[2] == 'C' {
					// Only want to go as far there are characters
					fmt.Print(cursorrt)
				} else if buf[2] == 'D' {
					fmt.Print(cursorlf)
				}
			} else {
				continue
			}
		} else if buf[0] == 13 {
			fmt.Println(CURSOR_LEFT)
			fmt.Printf("Echoing: %s\n%s", line_buffer, CURSOR_LEFT)
			line_buffer = make([]byte, 64)
			fmt.Print("> ")
		} else if buf[0] == 'q' || buf[0] == CTRL_C {
			fmt.Println(CURSOR_LEFT)
			break
		} else {
			line_buffer = append(line_buffer, buf[:]...)	
		}
	}
	return nil
}
