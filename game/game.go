package game

type Player byte

const (
	PLAYER_ONE = 'X'
	PLAYER_TWO = 'O'
)

type TurnOutcome int

const (
	FAIL    = 0
	SUCCESS = 1
	WIN     = 2
	TIE     = 3
)

type Game struct {
	Board [][]byte
	Turn  Player
}

func New() *Game {
	return &Game{
		Board: [][]byte{
			{' ', ' ', ' '},
			{' ', ' ', ' '},
			{' ', ' ', ' '},
		},
		Turn: PLAYER_ONE,
	}
}

func (g *Game) MakeMove(slot int) (TurnOutcome, string) {
	if slot < 1 || slot > 9 {
		return FAIL, "Invalid slot number\n"
	}

	var row, col int

	if slot%3 == 0 {
		row = (slot / 3) - 1
		col = 2
	} else {
		row = slot / 3
		col = (slot % 3) - 1
	}

	if g.Board[row][col] != ' ' {
		return FAIL, "This slot has been taken\n"
	}

	g.Board[row][col] = byte(g.Turn)

	result := g.checkWin()

	if result == SUCCESS {
		if g.Turn == PLAYER_ONE {
			g.Turn = PLAYER_TWO
		} else {
			g.Turn = PLAYER_ONE
		}
	}

	return result, ""
}

func (g *Game) checkWin() TurnOutcome {
	if g.Board[0][0] != ' ' && g.Board[0][0] == g.Board[0][1] && g.Board[0][0] == g.Board[0][2] {
		return WIN
	}
	if g.Board[1][0] != ' ' && g.Board[1][0] == g.Board[1][1] && g.Board[1][0] == g.Board[1][2] {
		return WIN
	}
	if g.Board[2][0] != ' ' && g.Board[2][0] == g.Board[2][1] && g.Board[2][0] == g.Board[2][2] {
		return WIN
	}

	if g.Board[0][0] != ' ' && g.Board[0][0] == g.Board[1][0] && g.Board[0][0] == g.Board[2][0] {
		return WIN
	}
	if g.Board[0][1] != ' ' && g.Board[0][1] == g.Board[1][1] && g.Board[0][1] == g.Board[2][1] {
		return WIN
	}
	if g.Board[0][2] != ' ' && g.Board[0][2] == g.Board[1][2] && g.Board[0][2] == g.Board[2][2] {
		return WIN
	}

	if g.Board[0][0] != ' ' && g.Board[0][0] == g.Board[1][1] && g.Board[0][0] == g.Board[2][2] {
		return WIN
	}
	if g.Board[0][2] != ' ' && g.Board[0][2] == g.Board[1][1] && g.Board[0][2] == g.Board[2][0] {
		return WIN
	}

	if g.checkTie() {
		return TIE
	}

	return SUCCESS
}

func (g *Game) checkTie() bool {
	if g.Board[0][0] == ' ' {
		return false
	}
	if g.Board[0][1] == ' ' {
		return false
	}
	if g.Board[0][2] == ' ' {
		return false
	}
	if g.Board[1][0] == ' ' {
		return false
	}
	if g.Board[1][1] == ' ' {
		return false
	}
	if g.Board[1][2] == ' ' {
		return false
	}
	if g.Board[2][0] == ' ' {
		return false
	}
	if g.Board[2][1] == ' ' {
		return false
	}
	if g.Board[2][2] == ' ' {
		return false
	}

	return true
}
