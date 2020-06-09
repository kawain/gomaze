package ana

import (
	"testing"
)

func TestStack(t *testing.T) {
	row := 1
	col := 1
	w := &worker{
		row:   row,
		col:   col,
		stack: [][]int{{row, col}},
	}

	w.push([]int{2, 2})
	w.push([]int{3, 3})
	w.push([]int{4, 4})

	a, err := w.pop()

	if err != nil {
		t.Fatalf("正解は %v\n", err)
	}

	if a[0] != 4 {
		t.Fatalf("正解は %v\n", a[0])
	}
	a, err = w.pop()
	a, err = w.pop()
	a, err = w.pop()
	if err != nil {
		t.Fatalf("正解は %v\n", err)
	}
	if a[0] != 1 {
		t.Fatalf("正解は %v\n", a[0])
	}
	a, err = w.pop()
	if err == nil {
		t.Fatalf("正解は %v\n", a[0])
	}
}
func TestInputCheck(t *testing.T) {
	i := inputCheck(10)
	if i != false {
		t.Fatalf("正解は %v\n", i)
	}
	i = inputCheck(4)
	if i != false {
		t.Fatalf("正解は %v\n", i)
	}
	i = inputCheck(5)
	if i != true {
		t.Fatalf("正解は %v\n", i)
	}
	i = inputCheck(101)
	if i != true {
		t.Fatalf("正解は %v\n", i)
	}
	i = inputCheck(99)
	if i != true {
		t.Fatalf("正解は %v\n", i)
	}
	i = inputCheck(20)
	if i != false {
		t.Fatalf("正解は %v\n", i)
	}
}
