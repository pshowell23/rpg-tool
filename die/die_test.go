package die

import (
	"testing"
)

//TestRoller tests Roller function
func TestRoller(t *testing.T) {
	dieSides := 20
	var rolls []int

	for i := 0; i < 100; i++ {
		roll := Roller(dieSides)
		rolls = append(rolls, roll)
	}

	for _, val := range rolls {
		if val > dieSides || val < 1 {
			t.Errorf("Die has inappropriate numbers")
		}
	}
}
