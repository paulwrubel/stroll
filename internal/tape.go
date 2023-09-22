package internal

type tape struct {
	cells []int64

	zeroIndex     int
	selectedIndex int
}

func NewTape() *tape {
	return &tape{
		cells:         make([]int64, 19),
		zeroIndex:     9,
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
