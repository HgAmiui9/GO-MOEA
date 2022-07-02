package problems

import (
	"testing"

	"github.com/HgAmiui9/GO-MOEA/numgo"
)

func TestZDT1(t *testing.T) {
	var var_n float64 = 30
	x_array := numgo.Linspace(1, var_n, int(var_n), true)

	res1, res2 := ZDT1(x_array, var_n)
	t.Log(res1)
	t.Log(res2)
}
