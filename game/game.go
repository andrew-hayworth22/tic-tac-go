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
	board [][]byte
	Turn  Player
}

func New() *Game {
	return &Game{
		board: [][]byte{
			{' ', ' ', ' '},
			{' ', ' ', ' '},
			{' ', ' ', ' '},
		},
		Turn: PLAYER_ONE,
	}
}

func (g *Game) MakeMove(player Player, slot int) (TurnOutcome, string) {
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

	if g.board[row][col] != ' ' {
		return FAIL, "This slot has been taken\n"
	}

	g.board[row][col] = byte(player)

	return g.checkWin(), ""
}

//   0 1 2
// 0 _ _ _
// 1 _ _ _
// 2 _ _ _

func (g *Game) checkWin() TurnOutcome {
	if g.board[0][0] != ' ' && g.board[0][0] == g.board[0][1] && g.board[0][0] == g.board[0][2] {
		return WIN
	}
	if g.board[1][0] != ' ' && g.board[1][0] == g.board[1][1] && g.board[1][0] == g.board[1][2] {
		return WIN
	}
	if g.board[2][0] != ' ' && g.board[2][0] == g.board[2][1] && g.board[2][0] == g.board[2][2] {
		return WIN
	}

	if g.board[0][0] != ' ' && g.board[0][0] == g.board[1][0] && g.board[0][0] == g.board[2][0] {
		return WIN
	}
	if g.board[0][1] != ' ' && g.board[0][1] == g.board[1][1] && g.board[0][1] == g.board[2][1] {
		return WIN
	}
	if g.board[0][2] != ' ' && g.board[0][2] == g.board[1][2] && g.board[0][2] == g.board[2][2] {
		return WIN
	}
	return SUCCESS
}
