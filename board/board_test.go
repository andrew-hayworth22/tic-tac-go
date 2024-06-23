package board

import (
	"reflect"
	"testing"
)

type BoardTest struct {
	moves     []Move
	expecteds []ExpectedBoardResult
}

type Move struct {
	slot int
	char byte
}

type ExpectedBoardResult struct {
	slots [][]byte
	ok    bool
	msg   string
}

func TestFillSlot(t *testing.T) {
	tests := []BoardTest{
		{
			moves: []Move{{1, 'X'}},
			expecteds: []ExpectedBoardResult{
				{
					slots: [][]byte{{'X', '2', '3'}, {'4', '5', '6'}, {'7', '8', '9'}},
					ok:    true,
					msg:   "",
				},
			},
		},
		{
			moves: []Move{{1, 'X'}, {5, 'O'}, {3, 'X'}, {9, 'O'}},
			expecteds: []ExpectedBoardResult{
				{
					slots: [][]byte{{'X', '2', '3'}, {'4', '5', '6'}, {'7', '8', '9'}},
					ok:    true,
					msg:   "",
				},
				{
					slots: [][]byte{{'X', '2', '3'}, {'4', 'O', '6'}, {'7', '8', '9'}},
					ok:    true,
					msg:   "",
				},
				{
					slots: [][]byte{{'X', '2', 'X'}, {'4', 'O', '6'}, {'7', '8', '9'}},
					ok:    true,
					msg:   "",
				},
				{
					slots: [][]byte{{'X', '2', 'X'}, {'4', 'O', '6'}, {'7', '8', 'O'}},
					ok:    true,
					msg:   "",
				},
			},
		},
		{
			moves: []Move{{1, 'X'}, {1, 'O'}},
			expecteds: []ExpectedBoardResult{
				{
					slots: [][]byte{{'X', '2', '3'}, {'4', '5', '6'}, {'7', '8', '9'}},
					ok:    true,
					msg:   "",
				},
				{
					slots: [][]byte{{'X', '2', '3'}, {'4', '5', '6'}, {'7', '8', '9'}},
					ok:    false,
					msg:   "This slot has been taken\n",
				},
			},
		},
		{
			moves: []Move{{10, 'X'}, {0, 'X'}},
			expecteds: []ExpectedBoardResult{
				{
					slots: [][]byte{{'1', '2', '3'}, {'4', '5', '6'}, {'7', '8', '9'}},
					ok:    false,
					msg:   "Invalid slot number\n",
				},
				{
					slots: [][]byte{{'1', '2', '3'}, {'4', '5', '6'}, {'7', '8', '9'}},
					ok:    false,
					msg:   "Invalid slot number\n",
				},
			},
		},
	}

	for _, tt := range tests {
		b := New()

		for idx, move := range tt.moves {
			ok, msg := b.FillSlot(move.slot, move.char)
			if ok != tt.expecteds[idx].ok {
				t.Errorf("Wrong response returned: Expected = %t. Got = %t. Message = %q", tt.expecteds[idx].ok, ok, msg)
			}

			if msg != tt.expecteds[idx].msg {
				t.Errorf("Wrong message returned: Expected = %q. Got = %q", tt.expecteds[idx].msg, msg)
			}

			if !reflect.DeepEqual(b.Slots, tt.expecteds[idx].slots) {
				t.Error("Wrong board values")
			}
		}
	}
}
