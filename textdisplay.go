package textdisplay

/*

 */

import (
	"fmt"
)

type sidebar struct {
	width  int
	border rune
	lines  []string
}

type mainDisplay struct {
	lines []string
	width int
	text  string
}

type screen struct {
	sidebarVisible        bool
	lines                 []string
	currentPage, maxPage  int
	pageWidth, pageHeight int
	sidebar               sidebar
	mainDisplay           mainDisplay
}

//------------------------------------------------

func (*sidebar) SetWidth(x int) {

}

func (scr *screen) UpdateScreen() {
	for index := 0; index < 10; index++ {
		if scr.lines[index] == "" {
			scr.lines[index] = string("->")
		}

		fmt.Printf("%s \n", scr.lines[index])

	}
}

func (*screen) ClearScreen() {
	for index := 0; index < 3; index++ {
		fmt.Printf("\n\n\n\n\n\n\n\n\n")
	}

}

func (scr *screen) SendToDisplay(s string) {
	scr.lines[0] = s
	scr.mainDisplay.text = s
}

//------------------------------------------------------

//NewTextDisplay returns a new <screen> consisting of a <sidebar> and <mainDisplay>
func NewTextDisplay() screen {
	var screen screen
	screen.pageHeight = 30
	screen.pageWidth = 80
	screen.lines = make([]string, 10, 10)
	screen.sidebar.border = '|'

	return screen
}
