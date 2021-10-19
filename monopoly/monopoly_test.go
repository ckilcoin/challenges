package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_initializeBoard(t *testing.T) {
	var board [BOARD_SIZE]int
	m := Monopoly{
		board: board,
	}
	for i := 0; i < BOARD_SIZE; i++ {
		assert.True(t, m.board[i] == 0)
	}
	m.initializeBoard()
	assert.True(t, m.board[0] == 0)
	for i := 1; i < BOARD_SIZE; i++ {
		assert.True(t, m.board[i] != 0)
	}
}
