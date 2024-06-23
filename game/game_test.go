package game

import (
	"reflect"
	"testing"
)

type GameTest struct {
	moves     []Move
	expecteds []ExpectedBoardResult
}

type Move struct {
	slot   int
	player Player
}

type ExpectedBoardResult struct {
	board  [][]byte
	result TurnOutcome
	msg    string
}

func TestFillSlot(t *testing.T) {
	tests := []GameTest{
		{
			moves: []Move{{1, PLAYER_ONE}},
			expecteds: []ExpectedBoardResult{
				{
					board:  [][]byte{{PLAYER_ONE, ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
			},
		},
		{
			moves: []Move{{1, PLAYER_ONE}, {5, PLAYER_TWO}, {3, PLAYER_ONE}, {9, PLAYER_TWO}},
			expecteds: []ExpectedBoardResult{
				{
					board:  [][]byte{{PLAYER_ONE, ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
				{
					board:  [][]byte{{PLAYER_ONE, ' ', ' '}, {' ', PLAYER_TWO, ' '}, {' ', ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
				{
					board:  [][]byte{{PLAYER_ONE, ' ', PLAYER_ONE}, {' ', PLAYER_TWO, ' '}, {' ', ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
				{
					board:  [][]byte{{PLAYER_ONE, ' ', PLAYER_ONE}, {' ', PLAYER_TWO, ' '}, {' ', ' ', PLAYER_TWO}},
					result: SUCCESS,
					msg:    "",
				},
			},
		},
		{
			moves: []Move{{1, PLAYER_ONE}, {1, PLAYER_TWO}},
			expecteds: []ExpectedBoardResult{
				{
					board:  [][]byte{{PLAYER_ONE, ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
				{
					board:  [][]byte{{PLAYER_ONE, ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}},
					result: FAIL,
					msg:    "This slot has been taken\n",
				},
			},
		},
		{
			moves: []Move{{10, PLAYER_ONE}, {0, PLAYER_ONE}},
			expecteds: []ExpectedBoardResult{
				{
					board:  [][]byte{{' ', ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}},
					result: FAIL,
					msg:    "Invalid slot number\n",
				},
				{
					board:  [][]byte{{' ', ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}},
					result: FAIL,
					msg:    "Invalid slot number\n",
				},
			},
		},
	}

	for _, tt := range tests {
		g := New()

		for idx, move := range tt.moves {
			result, msg := g.MakeMove(move.player, move.slot)
			if result != tt.expecteds[idx].result {
				t.Errorf("Unexpected turn result: Expected = %d. Got = %d. Message = %q", tt.expecteds[idx].result, result, msg)
			}

			if msg != tt.expecteds[idx].msg {
				t.Errorf("Unexpected message: Expected = %q. Got = %q", tt.expecteds[idx].msg, msg)
			}

			if !reflect.DeepEqual(g.board, tt.expecteds[idx].board) {
				t.Errorf("Wrong board values: Expected = %v. Got = %v", tt.expecteds[idx].board, g.board)
			}
		}
	}
}
