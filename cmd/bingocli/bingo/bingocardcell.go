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

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

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
		s = " *" + s + "*"
		fmt.Printf(WarningColor, s)
	} else {
		s = "  " + s + " "
		fmt.Printf(InfoColor, s)
	}
}
