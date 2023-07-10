package main

import "fmt"

type sum interface {
	add() int
}

type numsInt struct {
	a int
	b int
}

func newNumsInt(a, b int) *numsInt {
	return &numsInt{
		a: a,
		b: b,
	}
}

func (i numsInt) add() int {
	return i.a + i.b
}

type numsFloat struct {
	a float64
	b float64
}

func newNumsFloat(a, b float64) *numsFloat {
	return &numsFloat{
		a: a,
		b: b,
	}
}

func (f *numsFloat) convertToInt() (int, int) {
	return int(f.a), int(f.b)
}

type numsFloatAdapter struct {
	f *numsFloat
}

func newNumsFloatAdapter(a, b float64) *numsFloatAdapter {
	return &numsFloatAdapter{f: newNumsFloat(a, b)}
}

func (adapt *numsFloatAdapter) add() int {
	a, b := adapt.f.convertToInt()
	return a + b
}

func main() {
	i := newNumsInt(1, 6)
	fmt.Println(i.add())
	f := newNumsFloatAdapter(1.6, 5.9)
	fmt.Println(f.add())
}
