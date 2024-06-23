package game

import (
	"bufio"
	"fmt"
	"io"
	"tictacgo/board"
)

type Player byte

const (
	PLAYER_ONE = 'X'
	PLAYER_TWO = 'O'
)

type Game struct {
	Board *board.Board
	Turn  Player
}

func New(b *board.Board) *Game {
	return &Game{
		Board: b,
		Turn:  PLAYER_ONE,
	}
}

func (g *Game) Run(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	var character byte

	for {
		io.WriteString(out, g.Board.Draw())
		io.WriteString(out, fmt.Sprintf("It is %c's turn! Please enter a slot to fill:", g.Turn))
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		character = scanner.Bytes()[0]
		success, msg := g.Board.FillSlot(int(character-48), byte(g.Turn))
		if !success {
			io.WriteString(out, msg)
			continue
		}

		if g.checkWin() {
			io.WriteString(out, g.Board.Draw())
			io.WriteString(out, fmt.Sprintf("%c wins!!!\n", g.Turn))
			return
		}

		if g.Turn == PLAYER_ONE {
			g.Turn = PLAYER_TWO
		} else {
			g.Turn = PLAYER_ONE
		}
	}
}

//   0 1 2
// 0 _ _ _
// 1 _ _ _
// 2 _ _ _

func (g *Game) checkWin() bool {
	if !isNum(g.Board.Slots[0][0]) && g.Board.Slots[0][0] == g.Board.Slots[0][1] && g.Board.Slots[0][0] == g.Board.Slots[0][2] {
		return true
	}
	if !isNum(g.Board.Slots[1][0]) && g.Board.Slots[1][0] == g.Board.Slots[1][1] && g.Board.Slots[1][0] == g.Board.Slots[1][2] {
		return true
	}
	if !isNum(g.Board.Slots[2][0]) && g.Board.Slots[2][0] == g.Board.Slots[2][1] && g.Board.Slots[2][0] == g.Board.Slots[2][2] {
		return true
	}

	if !isNum(g.Board.Slots[0][0]) && g.Board.Slots[0][0] == g.Board.Slots[1][0] && g.Board.Slots[0][0] == g.Board.Slots[2][0] {
		return true
	}
	if !isNum(g.Board.Slots[0][1]) && g.Board.Slots[0][1] == g.Board.Slots[1][1] && g.Board.Slots[0][1] == g.Board.Slots[2][1] {
		return true
	}
	if !isNum(g.Board.Slots[0][2]) && g.Board.Slots[0][2] == g.Board.Slots[1][2] && g.Board.Slots[0][2] == g.Board.Slots[2][2] {
		return true
	}
	return false
}

func isNum(char byte) bool {
	return char >= '0' && char <= '9'
}
