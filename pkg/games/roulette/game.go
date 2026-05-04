package roulette

import (
	"math/rand/v2"
)

func RunGame() Trials {
	return getEntry()
}

func pickRandom() int {
	return rand.IntN(38)
}

func getEntry() Trials {
	board := BoardSetup()
	value := pickRandom()

	return Trials{
		IntValue: value,
		Color:    board[value],
	}
}
