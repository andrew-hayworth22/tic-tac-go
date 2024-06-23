package board

import (
	"bytes"
)

type Board struct {
	Slots [][]byte
}

func New() *Board {
	return &Board{
		Slots: [][]byte{
			{'1', '2', '3'},
			{'4', '5', '6'},
			{'7', '8', '9'},
		},
	}
}

func (b *Board) Draw() string {
	var out bytes.Buffer

	for r, row := range b.Slots {
		for c, slot := range row {
			out.WriteByte(slot)

			if c < len(row)-1 {
				out.WriteString(" | ")
			}
		}

		if r < len(b.Slots)-1 {
			out.WriteString("\n---------\n")
		}
	}

	out.WriteString("\n")

	return out.String()
}

func (b *Board) FillSlot(slot int, char byte) (bool, string) {
	if slot < 1 || slot > 9 {
		return false, "Invalid slot number\n"
	}

	var row, col int

	if slot%3 == 0 {
		row = (slot / 3) - 1
		col = 2
	} else {
		row = slot / 3
		col = (slot % 3) - 1
	}

	if b.Slots[row][col] < '0' || b.Slots[row][col] > '9' {
		return false, "This slot has been taken\n"
	}

	b.Slots[row][col] = char

	return true, ""
}
