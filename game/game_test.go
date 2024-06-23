package game

import (
	"bytes"
	"testing"
	"tictacgo/board"
)

func TestGame(t *testing.T) {
	g := New(board.New())

	var in bytes.Buffer
	in.WriteString("1")
	var out bytes.Buffer

	g.Run(&in, &out)
	t.Fatal(out.String())
}
