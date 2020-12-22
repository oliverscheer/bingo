package bingo

import "testing"

func applyAndTestValues(t *testing.T, card Card, numbers []int) {
	numberCount := len(numbers)
	for idx, number := range numbers {
		hit := card.CheckNumber(number)
		if !hit {
			t.Errorf("Hit check, got %t, wanted %t", hit, true)
		}
		win := card.CheckForWin()
		if idx < numberCount-1 {
			if win {
				t.Errorf("Win check, got %t, wanted %t", win, false)
			}
		} else {
			if !win {
				t.Errorf("Win check, got %t, wanted %t", win, true)
			}
		}
	}
}

func TestFirstHorizontalLineWin(t *testing.T) {
	cards := []Card{
		{
			ID: 1,
			Cell: [5][5]CardCell{
				{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
				{{Value: 16}, {Value: 17}, {Value: 18}, {Value: 19}, {Value: 20}},
				{{Value: 31}, {Value: 32}, {Value: 0}, {Value: 33}, {Value: 34}},
				{{Value: 46}, {Value: 47}, {Value: 48}, {Value: 49}, {Value: 50}},
				{{Value: 61}, {Value: 62}, {Value: 63}, {Value: 64}, {Value: 65}},
			},
		},
	}
	card := cards[0]
	numbers := []int{1, 2, 3, 4, 5}
	applyAndTestValues(t, card, numbers)
}

func TestSecondHorizontalLineWin(t *testing.T) {
	cards := []Card{
		{
			ID: 1,
			Cell: [5][5]CardCell{
				{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
				{{Value: 16}, {Value: 17}, {Value: 18}, {Value: 19}, {Value: 20}},
				{{Value: 31}, {Value: 32}, {Value: 0}, {Value: 33}, {Value: 34}},
				{{Value: 46}, {Value: 47}, {Value: 48}, {Value: 49}, {Value: 50}},
				{{Value: 61}, {Value: 62}, {Value: 63}, {Value: 64}, {Value: 65}},
			},
		},
	}
	card := cards[0]
	numbers := []int{16, 17, 18, 19, 20}
	applyAndTestValues(t, card, numbers)
}

func TestThirdHorizontalLineWin(t *testing.T) {
	cards := []Card{
		{
			ID: 1,
			Cell: [5][5]CardCell{
				{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
				{{Value: 16}, {Value: 17}, {Value: 18}, {Value: 19}, {Value: 20}},
				{{Value: 31}, {Value: 32}, {Value: 0}, {Value: 33}, {Value: 34}},
				{{Value: 46}, {Value: 47}, {Value: 48}, {Value: 49}, {Value: 50}},
				{{Value: 61}, {Value: 62}, {Value: 63}, {Value: 64}, {Value: 65}},
			},
		},
	}
	card := cards[0]
	numbers := []int{31, 32, 33, 34}
	applyAndTestValues(t, card, numbers)
}
func TestFourthHorizontalLineWin(t *testing.T) {
	cards := []Card{
		{
			ID: 1,
			Cell: [5][5]CardCell{
				{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
				{{Value: 16}, {Value: 17}, {Value: 18}, {Value: 19}, {Value: 20}},
				{{Value: 31}, {Value: 32}, {Value: 0}, {Value: 33}, {Value: 34}},
				{{Value: 46}, {Value: 47}, {Value: 48}, {Value: 49}, {Value: 50}},
				{{Value: 61}, {Value: 62}, {Value: 63}, {Value: 64}, {Value: 65}},
			},
		},
	}
	card := cards[0]
	numbers := []int{46, 47, 48, 49, 50}
	applyAndTestValues(t, card, numbers)
}

func TestFifthHorizontalLineWin(t *testing.T) {
	cards := []Card{
		{
			ID: 1,
			Cell: [5][5]CardCell{
				{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
				{{Value: 16}, {Value: 17}, {Value: 18}, {Value: 19}, {Value: 20}},
				{{Value: 31}, {Value: 32}, {Value: 0}, {Value: 33}, {Value: 34}},
				{{Value: 46}, {Value: 47}, {Value: 48}, {Value: 49}, {Value: 50}},
				{{Value: 61}, {Value: 62}, {Value: 63}, {Value: 64}, {Value: 65}},
			},
		},
	}
	card := cards[0]
	numbers := []int{61, 62, 63, 64, 65}
	applyAndTestValues(t, card, numbers)
}

func TestFirstVerticalLineWin(t *testing.T) {
	cards := []Card{
		{
			ID: 1,
			Cell: [5][5]CardCell{
				{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
				{{Value: 16}, {Value: 17}, {Value: 18}, {Value: 19}, {Value: 20}},
				{{Value: 31}, {Value: 32}, {Value: 0}, {Value: 33}, {Value: 34}},
				{{Value: 46}, {Value: 47}, {Value: 48}, {Value: 49}, {Value: 50}},
				{{Value: 61}, {Value: 62}, {Value: 63}, {Value: 64}, {Value: 65}},
			},
		},
	}
	card := cards[0]
	numbers := []int{1, 16, 31, 46, 61}
	applyAndTestValues(t, card, numbers)
}

func TestSecondVerticalLineWin(t *testing.T) {
	cards := []Card{
		{
			ID: 1,
			Cell: [5][5]CardCell{
				{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
				{{Value: 16}, {Value: 17}, {Value: 18}, {Value: 19}, {Value: 20}},
				{{Value: 31}, {Value: 32}, {Value: 0}, {Value: 33}, {Value: 34}},
				{{Value: 46}, {Value: 47}, {Value: 48}, {Value: 49}, {Value: 50}},
				{{Value: 61}, {Value: 62}, {Value: 63}, {Value: 64}, {Value: 65}},
			},
		},
	}
	card := cards[0]
	numbers := []int{2, 17, 32, 47, 62}
	applyAndTestValues(t, card, numbers)
}

func TestThirdVerticalLineWin(t *testing.T) {
	cards := []Card{
		{
			ID: 1,
			Cell: [5][5]CardCell{
				{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
				{{Value: 16}, {Value: 17}, {Value: 18}, {Value: 19}, {Value: 20}},
				{{Value: 31}, {Value: 32}, {Value: 0}, {Value: 33}, {Value: 34}},
				{{Value: 46}, {Value: 47}, {Value: 48}, {Value: 49}, {Value: 50}},
				{{Value: 61}, {Value: 62}, {Value: 63}, {Value: 64}, {Value: 65}},
			},
		},
	}
	card := cards[0]
	numbers := []int{3, 18, 19, 48, 63}
	applyAndTestValues(t, card, numbers)
}
func TestFourthVerticalLineWin(t *testing.T) {
	cards := []Card{
		{
			ID: 1,
			Cell: [5][5]CardCell{
				{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
				{{Value: 16}, {Value: 17}, {Value: 18}, {Value: 19}, {Value: 20}},
				{{Value: 31}, {Value: 32}, {Value: 0}, {Value: 33}, {Value: 34}},
				{{Value: 46}, {Value: 47}, {Value: 48}, {Value: 49}, {Value: 50}},
				{{Value: 61}, {Value: 62}, {Value: 63}, {Value: 64}, {Value: 65}},
			},
		},
	}
	card := cards[0]
	numbers := []int{4, 19, 33, 49, 64}
	applyAndTestValues(t, card, numbers)
}

func TestFifthVerticalLineWin(t *testing.T) {
	cards := []Card{
		{
			ID: 1,
			Cell: [5][5]CardCell{
				{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
				{{Value: 16}, {Value: 17}, {Value: 18}, {Value: 19}, {Value: 20}},
				{{Value: 31}, {Value: 32}, {Value: 0}, {Value: 33}, {Value: 34}},
				{{Value: 46}, {Value: 47}, {Value: 48}, {Value: 49}, {Value: 50}},
				{{Value: 61}, {Value: 62}, {Value: 63}, {Value: 64}, {Value: 65}},
			},
		},
	}
	card := cards[0]
	numbers := []int{5, 20, 34, 50, 65}
	applyAndTestValues(t, card, numbers)
}

func TestLeftTopToRightBottomLineWin(t *testing.T) {
	cards := []Card{
		{
			ID: 1,
			Cell: [5][5]CardCell{
				{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
				{{Value: 16}, {Value: 17}, {Value: 18}, {Value: 19}, {Value: 20}},
				{{Value: 31}, {Value: 32}, {Value: 0}, {Value: 33}, {Value: 34}},
				{{Value: 46}, {Value: 47}, {Value: 48}, {Value: 49}, {Value: 50}},
				{{Value: 61}, {Value: 62}, {Value: 63}, {Value: 64}, {Value: 65}},
			},
		},
	}
	card := cards[0]
	numbers := []int{1, 17, 49, 65}
	applyAndTestValues(t, card, numbers)
}

func TestRightTopToLeftBottomLineWin(t *testing.T) {
	cards := []Card{
		{
			ID: 1,
			Cell: [5][5]CardCell{
				{{Value: 1}, {Value: 2}, {Value: 3}, {Value: 4}, {Value: 5}},
				{{Value: 16}, {Value: 17}, {Value: 18}, {Value: 19}, {Value: 20}},
				{{Value: 31}, {Value: 32}, {Value: 0}, {Value: 33}, {Value: 34}},
				{{Value: 46}, {Value: 47}, {Value: 48}, {Value: 49}, {Value: 50}},
				{{Value: 61}, {Value: 62}, {Value: 63}, {Value: 64}, {Value: 65}},
			},
		},
	}
	card := cards[0]
	numbers := []int{5, 19, 47, 61}
	applyAndTestValues(t, card, numbers)
}
