package testing

import (
	"testing"

	dist "github.com/bouwerja/gomars/pkg/distributions"
)

func TestDistribution(t *testing.T) {
	t.Log("Test has strated")

	ans, err := dist.PoissonValues(0.5, 6)
	if err != nil {
		t.Log(err)
	}
	t.Log(ans)

	t.Log("Test has stopped")
}
