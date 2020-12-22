package bingo

import "testing"

func TestPlayer(t *testing.T) {
	p := NewPlayer(1, "Oliver")

	if p.ID != 1 {
		t.Errorf("Name set wrong, got %d, want: %d", p.ID, 1)
	}
	if p.Name != "Oliver" {
		t.Errorf("Name set wrong, got %s, want: %s", p.Name, "Oliver")
	}

}

func TestAddCard(t *testing.T) {
}
