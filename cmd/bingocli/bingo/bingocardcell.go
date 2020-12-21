package bingo

import (
	"fmt"
	"strconv"
)

// BingoCardCell describes a BingoCard cell
type BingoCardCell struct {
	Value int
	Hit   bool
}

func (bc *BingoCardCell) printDetails() {
	var v int
	v = bc.Value
	var s string
	if v == 0 {
		s = " *"
	} else if v < 10 {
		s = " " + strconv.Itoa(v)
	} else {
		s = strconv.Itoa(v)
	}
	if bc.Hit {
		s = "*" + s + "*"
	} else {
		s = " " + s + " "
	}
	fmt.Print(" ", s)
}
