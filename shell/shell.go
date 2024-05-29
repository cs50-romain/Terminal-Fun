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
	NEW_LINE = 13
	CLEAR_LINE = "\u001b[2K"
	DEL = 127
	DELETE_CHAR = "\u001b[P"
)

type Cursor struct {
	HorizontalPos int
	VerticalPos   int
	PreviousPos   int
}

type StatusBar struct {
	CursorPosition string
	Line string
}

var previousCommands = []string{}
var cursor Cursor

func Run() error {
	fmt.Print(CURSOR_LEFT)
	fmt.Print("> ")
	var line_buffer = make([]byte, 0)
	cursor.VerticalPos = 0
	cursor.HorizontalPos = 0
	for {
		// Calculate current line length
		var buf [3]byte
		os.Stdin.Read(buf[:])
		fmt.Print(string(buf[0]))
		if buf[0] == '\u001b' {
			if buf[1] == '[' {
				if buf[2] == 'A' || buf[2] == 'B' || buf[2] == 'C' || buf[2] == 'D' { // RIGHT
					handleArrows(buf[2], len(line_buffer))
				}
			}
		} else if buf[0] == NEW_LINE {
			previousCommands = append(previousCommands, string(line_buffer[:]))
			fmt.Println(CURSOR_LEFT)
			fmt.Print(cursor)
			fmt.Printf("Echoing: %s\n%s", line_buffer, CURSOR_LEFT)
			line_buffer = make([]byte, 0)
			cursor.HorizontalPos = 0
			cursor.VerticalPos = len(previousCommands) - 1
			fmt.Print("> ")
		} else if buf[0] == 'q' || buf[0] == CTRL_C {
			fmt.Print(cursor)
			fmt.Println(CURSOR_LEFT)
			break
		} else if buf[0] == DEL {
			if len(line_buffer) / 3 == 0 {
				continue
			}
			line_buffer = line_buffer[:len(line_buffer)-3]
			fmt.Printf("%s%s", cursorlf, DELETE_CHAR)
		} else {
			line_buffer = append(line_buffer, buf[:]...)
			cursor.HorizontalPos++
		}
	}
	return nil
}

func handleArrows(char byte, length int) string {
	if char == 'C' { // RIGHT
		// Only want to go as far there are characters
		if cursor.HorizontalPos <= length / 3 - 1{
			fmt.Print(cursorrt)
			cursor.HorizontalPos++
		}
	} else if char == 'D' { // LEFT
		if cursor.HorizontalPos > 0 {
			fmt.Print(cursorlf)
			if cursor.HorizontalPos > 0 {
				cursor.HorizontalPos--
			}
		}
	} else if char == 'A' { // UP
		if len(previousCommands) == 0 {
			return ""
		}
		if cursor.VerticalPos == len(previousCommands) {
			return ""
		}

		fmt.Printf("%s%s> ", CLEAR_LINE, CURSOR_LEFT)
		fmt.Print(previousCommands[cursor.VerticalPos])

		if cursor.VerticalPos == 0 {
			return ""
		}
		cursor.VerticalPos--
	} else if char == 'B' { // DOWN
		if length == 0 && cursor.VerticalPos == len(previousCommands)-1{
			return ""
		}

		if cursor.VerticalPos == len(previousCommands) - 1{
			fmt.Print(" ")
			return ""
		}
	
		if cursor.VerticalPos != len(previousCommands) {
			cursor.VerticalPos++
		}

		fmt.Printf("%s%s> ", CLEAR_LINE, CURSOR_LEFT)
		fmt.Print(previousCommands[cursor.VerticalPos])
	}
	return ""
}
