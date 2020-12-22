package bingo

import (
	"fmt"
	"strconv"
)

// CardCell describes a BingoCard cell
type CardCell struct {
	Value int
	hit   bool
}

// NewCardCell creates a new card cell
func NewCardCell(v int) CardCell {
	cc := CardCell{
		Value: v,
	}
	return cc
}

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func (cc *CardCell) printDetails() {
	var v int
	v = cc.Value
	var s string
	if v == 0 {
		s = " *"
	} else if v < 10 {
		s = " " + strconv.Itoa(v)
	} else {
		s = strconv.Itoa(v)
	}
	if cc.IsHit() {
		s = " *" + s + "*"
		fmt.Printf(WarningColor, s)
	} else {
		s = "  " + s + " "
		fmt.Printf(InfoColor, s)
	}
}

// CheckNumber checks if number is in cell
func (cc *CardCell) CheckNumber(number int) bool {
	if number == cc.Value {
		cc.hit = true
		return true
	}
	return false
}

// IsHit returns true if cell contains number that is already hit
func (cc *CardCell) IsHit() bool {
	if cc.Value == 0 {
		cc.hit = true
	}
	return cc.hit
}
