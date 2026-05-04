package testing

import (
	"testing"

	"github.com/bouwerja/gomars/pkg/games/roulette"
)

func TestRoulette(t *testing.T) {
	t.Log("Test has strated")

	ans := roulette.Run(1000000)

	// t.Log(ans)
	t.Log(len(ans.Trial))

	t.Log("Test has stopped")
}
