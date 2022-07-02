package numgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinspace(t *testing.T) {
	exp := []float64{0, 1, 2}
	res := Linspace(0, 2, 3, true)
	assert.Equal(t, exp, res, "they should be equal")
}
