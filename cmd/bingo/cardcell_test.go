package bingo

import "testing"

func TestCardCellCheckNumber(t *testing.T) {
	cc := NewCardCell(11)

	hit := cc.CheckNumber(12)
	if hit {
		t.Errorf("Wrong check, want: %t, got: %t", false, true)
	}
}
