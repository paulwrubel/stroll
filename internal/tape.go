package internal

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type tape struct {
	cells []int64

	zeroIndex     int
	selectedIndex int
}

func NewTape() *tape {
	return &tape{
		cells:         make([]int64, 10),
		zeroIndex:     0,
		selectedIndex: 0,
	}
}

func (t *tape) Read() int64 {
	return t.cells[t.zeroIndex+t.selectedIndex]
}

func (t *tape) Write(val int64) {
	t.cells[t.zeroIndex+t.selectedIndex] = val
}

func (t *tape) MoveTo(i int) {
	t.selectedIndex = i

	trueIndex := t.zeroIndex + t.selectedIndex
	if trueIndex >= len(t.cells) {
		t.expandRight(trueIndex - (len(t.cells) - 1))
	} else if trueIndex < 0 {
		t.expandLeft(-trueIndex)
	}
}

func (t *tape) MoveForward() {
	t.selectedIndex++

	trueIndex := t.zeroIndex + t.selectedIndex
	if trueIndex >= len(t.cells) {
		t.expandRight(1)
	}
}

func (t *tape) MoveBackward() {
	t.selectedIndex--

	trueIndex := t.zeroIndex + t.selectedIndex
	if trueIndex < 0 {
		t.expandLeft(1)
	}
}

func (t *tape) expandRight(n int) {
	t.cells = append(t.cells, make([]int64, n)...)
}

func (t *tape) expandLeft(n int) {
	t.cells = append(make([]int64, n), t.cells...)
	t.zeroIndex += n
}

func (t *tape) DebugPrint() {
	selectedStrings := []string{}
	indexStrings := []string{}
	valueStrings := []string{}
	charStrings := []string{}
	for i := 0; i < len(t.cells); i++ {
		idxString := strconv.FormatInt(int64(i-t.zeroIndex), 10)
		valString := strconv.FormatInt(t.cells[i], 10)
		charVal := rune(t.cells[i])
		var charString string
		if !unicode.IsPrint(charVal) {
			charString = "_NP_"
		} else {
			charString = fmt.Sprintf("'%c'", t.cells[i])
		}

		maxLen := max(len(idxString), len(valString), len(charString))
		if len(valString) < maxLen {
			valString = strings.Repeat(" ", maxLen-len(valString)) + valString
		}
		if len(idxString) < maxLen {
			idxString = strings.Repeat(" ", maxLen-len(idxString)) + idxString
		}
		if len(charString) < maxLen {
			charString = strings.Repeat(" ", maxLen-len(charString)) + charString
		}
		var selString string
		if i == t.zeroIndex+t.selectedIndex {
			selString = strings.Repeat(" ", maxLen-1) + "*"
		} else {
			selString = strings.Repeat(" ", maxLen)
		}

		selectedStrings = append(selectedStrings, selString)
		indexStrings = append(indexStrings, idxString)
		valueStrings = append(valueStrings, valString)
		charStrings = append(charStrings, charString)
	}

	fmt.Println()
	fmt.Printf("sel: [  %s  ]\n", strings.Join(selectedStrings, "   "))
	fmt.Printf("idx: [  %s  ]\n", strings.Join(indexStrings, "   "))
	fmt.Printf("val: [  %s  ]\n", strings.Join(valueStrings, ",  "))
	fmt.Printf("cha: [  %s  ]\n", strings.Join(charStrings, ",  "))
}
