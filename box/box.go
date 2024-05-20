package box

import (
	"fmt"
)

// I WISH I COULD CREATE ENUMS

const (
	BOX_COL_SIZE = 20
	BOX_ROW_SIZE = 10
	CLEAR_SCREEN = "\u001b[2J"
	CURSOR_HOME = "\u001b[H"
	CURSOR_LEFT = "\u001b[1000D"
)

const (
	TOP_LEFT_CORNER = "+"
	TOP_RIGHT_CORNER = "+"
	BOTTOM_LEFT_CORNER = "+"
	BOTTOM_RIGHT_CORNER = "+"
	SIDES = "|"
	TOP = "-"
	BOTTOM = "_"
)

type Box struct {
	
}

func Render() {
	// Turn to raw mode
	// Clear the screen
	fmt.Print(CLEAR_SCREEN)
	fmt.Printf("\r%s", CURSOR_HOME)

	box := [BOX_ROW_SIZE][BOX_COL_SIZE]string{}
	for rIdx, row := range box {
		for cIdx := range row {
			if (rIdx == 0 || rIdx == len(box)-1) && (cIdx != 0 || cIdx != len(row)-1) {
				box[rIdx][cIdx] = "-"
			} else {
				box[rIdx][cIdx] = " "
			}
			if rIdx != 0 && (cIdx == 0 || cIdx == len(row)-1) {
				box[rIdx][cIdx] = "|"
			}
			if (rIdx == 0 && cIdx == 0) {
				box[rIdx][cIdx] = TOP_LEFT_CORNER
			}
			if (rIdx == len(box)-1 && cIdx == 0) {
				box[rIdx][cIdx] = BOTTOM_LEFT_CORNER
			}
			if (rIdx == 0 && cIdx == len(row)-1) {
				box[rIdx][cIdx] = TOP_RIGHT_CORNER
			}
			if (rIdx == len(box)-1 && cIdx == len(row)-1) {
				box[rIdx][cIdx] = BOTTOM_RIGHT_CORNER
			}
		}
	}
	// Draw a box... I think I need to use arrays
	for _, row := range box {
		for _, col:= range row {
			fmt.Printf("%s", col)
		}
		fmt.Printf("\n%s", CURSOR_LEFT)
	}
	// Allows user to input within that box. Need locations
	// Need rerendering
}

func NewBox() *Box {
	return nil
}

func EnableRawMode() {
	
}
