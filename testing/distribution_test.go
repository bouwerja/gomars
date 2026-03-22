package testing

import (
	"testing"

	dist "github.com/bouwerja/gomars/pkg/distributions"
)

func TestDistribution(t *testing.T) {
	t.Log("Test has strated")

	ans := dist.NormalDistribution()

	t.Log(ans)

	t.Log("Test has stopped")
}
