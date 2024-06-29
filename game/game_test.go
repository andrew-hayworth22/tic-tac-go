package game

import (
	"reflect"
	"testing"
)

type GameTest struct {
	moves     []int
	expecteds []ExpectedBoardResult
}

type ExpectedBoardResult struct {
	board  [][]byte
	result TurnOutcome
	msg    string
}

func TestMakeMove(t *testing.T) {
	tests := []GameTest{
		{
			moves: []int{1},
			expecteds: []ExpectedBoardResult{
				{
					board:  [][]byte{{PLAYER_ONE, ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
			},
		},
		{
			moves: []int{1, 5, 3, 9},
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
			moves: []int{1, 1},
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
			moves: []int{10, 0},
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
		{
			moves: []int{1, 2, 5, 3, 6, 4, 7, 9, 8},
			expecteds: []ExpectedBoardResult{
				{
					board:  [][]byte{{PLAYER_ONE, ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
				{
					board:  [][]byte{{PLAYER_ONE, PLAYER_TWO, ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
				{
					board:  [][]byte{{PLAYER_ONE, PLAYER_TWO, ' '}, {' ', PLAYER_ONE, ' '}, {' ', ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
				{
					board:  [][]byte{{PLAYER_ONE, PLAYER_TWO, PLAYER_TWO}, {' ', PLAYER_ONE, ' '}, {' ', ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
				{
					board:  [][]byte{{PLAYER_ONE, PLAYER_TWO, PLAYER_TWO}, {' ', PLAYER_ONE, PLAYER_ONE}, {' ', ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
				{
					board:  [][]byte{{PLAYER_ONE, PLAYER_TWO, PLAYER_TWO}, {PLAYER_TWO, PLAYER_ONE, PLAYER_ONE}, {' ', ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
				{
					board:  [][]byte{{PLAYER_ONE, PLAYER_TWO, PLAYER_TWO}, {PLAYER_TWO, PLAYER_ONE, PLAYER_ONE}, {PLAYER_ONE, ' ', ' '}},
					result: SUCCESS,
					msg:    "",
				},
				{
					board:  [][]byte{{PLAYER_ONE, PLAYER_TWO, PLAYER_TWO}, {PLAYER_TWO, PLAYER_ONE, PLAYER_ONE}, {PLAYER_ONE, ' ', PLAYER_TWO}},
					result: SUCCESS,
					msg:    "",
				},
				{
					board:  [][]byte{{PLAYER_ONE, PLAYER_TWO, PLAYER_TWO}, {PLAYER_TWO, PLAYER_ONE, PLAYER_ONE}, {PLAYER_ONE, PLAYER_ONE, PLAYER_TWO}},
					result: TIE,
					msg:    "",
				},
			},
		},
	}

	for tidx, tt := range tests {
		g := New()

		for idx, move := range tt.moves {
			result, msg := g.MakeMove(move)
			if result != tt.expecteds[idx].result {
				t.Errorf("TEST %d: Unexpected turn result: Expected = %s. Got = %s. Message = %q", tidx+1, getResultText(tt.expecteds[idx].result), getResultText(result), msg)
			}

			if msg != tt.expecteds[idx].msg {
				t.Errorf("TEST %d: Unexpected message: Expected = %q. Got = %q", tidx+1, tt.expecteds[idx].msg, msg)
			}

			if !reflect.DeepEqual(g.Board, tt.expecteds[idx].board) {
				t.Errorf("TEST %d: Wrong board values: Expected = %v. Got = %v", tidx+1, tt.expecteds[idx].board, g.Board)
			}
		}
	}
}

func getResultText(result TurnOutcome) string {
	switch result {
	case 0:
		return "FAIL"
	case 1:
		return "SUCCESS"
	case 2:
		return "WIN"
	case 3:
		return "TIE"
	default:
		return "UNKNOWN"
	}
}
