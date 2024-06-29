package main

import (
	"fmt"
	"log"
	"tictacgo/game"

	"github.com/gdamore/tcell/v2"
)

const (
	BOARD_X  = 0
	BOARD_Y  = 3
	STATUS_X = 0
	STATUS_Y = 10
)

func main() {
	defaultStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	buttonStyle := tcell.StyleDefault.Background(tcell.ColorBlue).Foreground(tcell.ColorBlue)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}

	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	s.SetStyle(defaultStyle)
	s.EnableMouse()
	s.Clear()

	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	g := game.New()
	drawHeader(s, defaultStyle)
	drawBoard(s, defaultStyle)
	drawSlots(s, g, defaultStyle, buttonStyle)
	drawStatus(s, defaultStyle, "X's turn!")

	for {
		s.Show()

		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEsc:
				return
			}
		case *tcell.EventMouse:
			if btns := ev.Buttons(); btns == tcell.Button1 {
				x, y := ev.Position()
				slot := 0

				if x == BOARD_X+1 && y == BOARD_Y {
					slot = 1
				} else if x == BOARD_X+5 && y == BOARD_Y {
					slot = 2
				} else if x == BOARD_X+9 && y == BOARD_Y {
					slot = 3
				} else if x == BOARD_X+1 && y == BOARD_Y+2 {
					slot = 4
				} else if x == BOARD_X+5 && y == BOARD_Y+2 {
					slot = 5
				} else if x == BOARD_X+9 && y == BOARD_Y+2 {
					slot = 6
				} else if x == BOARD_X+1 && y == BOARD_Y+4 {
					slot = 7
				} else if x == BOARD_X+5 && y == BOARD_Y+4 {
					slot = 8
				} else if x == BOARD_X+9 && y == BOARD_Y+4 {
					slot = 9
				}

				if slot != 0 {
					outcome, _ := g.MakeMove(slot)
					switch outcome {
					case game.SUCCESS:
						drawSlots(s, g, defaultStyle, buttonStyle)
						drawStatus(s, defaultStyle, fmt.Sprintf("%c's turn!", g.Turn))
					case game.WIN:
						drawSlots(s, g, defaultStyle, buttonStyle)
						drawStatus(s, defaultStyle, fmt.Sprintf("%c wins!!!", g.Turn))
					case game.TIE:
						drawSlots(s, g, defaultStyle, buttonStyle)
						drawStatus(s, defaultStyle, "It's a tie!")
					}
				}
			}
		}
	}
}

func drawHeader(s tcell.Screen, style tcell.Style) {
	header := "Welcome to tic-tac-go"

	for idx, char := range header {
		s.SetContent(idx, 0, char, nil, style)
		s.SetContent(idx, 1, tcell.RuneHLine, nil, style)
	}
}

func drawBoard(s tcell.Screen, style tcell.Style) {
	xOffset := 0
	for xOffset < 11 {
		s.SetContent(BOARD_X+xOffset, BOARD_Y+1, tcell.RuneHLine, nil, style)
		s.SetContent(BOARD_X+xOffset, BOARD_Y+3, tcell.RuneHLine, nil, style)
		xOffset++
	}

	s.SetContent(BOARD_X+3, BOARD_Y, tcell.RuneVLine, nil, style)
	s.SetContent(BOARD_X+7, BOARD_Y, tcell.RuneVLine, nil, style)

	s.SetContent(BOARD_X+3, BOARD_Y+2, tcell.RuneVLine, nil, style)
	s.SetContent(BOARD_X+7, BOARD_Y+2, tcell.RuneVLine, nil, style)

	s.SetContent(BOARD_X+3, BOARD_Y+4, tcell.RuneVLine, nil, style)
	s.SetContent(BOARD_X+7, BOARD_Y+4, tcell.RuneVLine, nil, style)
}

func drawSlots(s tcell.Screen, g *game.Game, defaultStyle tcell.Style, buttonStyle tcell.Style) {
	for rowIdx, row := range g.Board {
		for colIdx, slot := range row {
			style := defaultStyle
			if slot == ' ' {
				style = buttonStyle
			}
			s.SetContent((BOARD_X+1)+(4*colIdx), BOARD_Y+(2*rowIdx), rune(slot), nil, style)
		}
	}
}

func drawStatus(s tcell.Screen, style tcell.Style, status string) {
	for idx, char := range status {
		s.SetContent(STATUS_X+idx, STATUS_Y, char, nil, style)
	}
}
