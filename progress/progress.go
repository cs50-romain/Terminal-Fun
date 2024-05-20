package progress

import (
	"fmt"
	"time"
)

const (
	BOX_COL_SIZE = 20
	BOX_ROW_SIZE = 10
	CLEAR_SCREEN = "\u001b[2J"
	CURSOR_HOME = "\u001b[H"
	CURSOR_LEFT = "\u001b[1000D"
	HIDE_CURSOR = "\u001b[?25l"
	SHOW_CURSOR = "\u001b[?25h"
)

// Colors
const ( 
	Reset = "\u001b[0m"
	Yellow = "\u001b[33m"
	Green = "\u001b[35m"
	GreenBold = "\u001b[1;35m"
)

type Progress struct {
	char string
}

func ProgressRender() {
	fmt.Println(CLEAR_SCREEN)
	fmt.Printf("\r%s", CURSOR_HOME)
	fmt.Printf("%s", HIDE_CURSOR)

	progressBar := make([]string, 10)
	i := 0
	for idx := range progressBar {
		progressBar[idx] = "#"
	}

	fmt.Printf("[")
	for i < len(progressBar) {
		fmt.Printf("%s%s", Yellow, progressBar[i])
		time.Sleep(1 * time.Second)
		i++
	}
	fmt.Printf("%s] %sDONE!%s", Reset, GreenBold, Reset)
	fmt.Printf("%s\n", SHOW_CURSOR)
	fmt.Print(CURSOR_LEFT)
}

func hideCursor() {
	fmt.Printf("%s", HIDE_CURSOR)
}

func showCursor() {
	fmt.Printf("%s", SHOW_CURSOR)
}
