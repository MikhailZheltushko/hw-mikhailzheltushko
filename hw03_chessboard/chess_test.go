package main

import (
	"testing"
)

func TestGetChessboard3x3(t *testing.T) {
	size := 3
	expectedOutput := "#   # \n  #   \n#   # \n"

	chessboard := getChessboard(size)

	if chessboard != expectedOutput {
		t.Errorf("Ожидаемый результат:\n%s\nПолучили:\n%s\n", expectedOutput, chessboard)
	}
}
