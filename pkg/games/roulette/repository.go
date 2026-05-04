package roulette

func BoardSetup() map[int]string {
	board := map[int]string{0: "green"}

	for i := range 37 {
		i++

		if i%2 == 0 {
			board[i] = "red"
		} else {
			board[i] = "black"
		}
	}

	return board
}

type TrialsResult struct {
	Trial []Trials
}

type Trials struct {
	IntValue int
	Color    string
}
