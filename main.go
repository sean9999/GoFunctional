package main

import "fmt"

// Flice method definitions
type FSlice[T any] []T

// method signatures

type MapFunction[T any] (v T, i int, arr []T) (T)
type FilterFunction[X] func(v X, i int, arr []X) (bool)
type SomeFunction func(v int, i int, arr []int) bool
type EveryFunction func(v int, i int, arr []int) bool


type MethodSet interface {
	Map(MapFunction) []int
	Filter(FilterFunction) bool
	Some(SomeFunction) bool
	Every(EveryFunction) bool
}


func NewFSlice[T any](inputSlice []T) FSlice[T] {
	return FSlice[T](inputSlice)
}

func (fs FSlice[I]) Map(fn func[I, O](I,int,[]I) (FSlice[O])) FSlice[O] {
	var r []O
	for i, v := range fs {
		r = append(r, fn(v, i, fs))
	}
	return FSlice[O](r)
}

func (fs FSlice) Filter(fn FilterFunction) FSlice {
	var r FSlice

	for i, v := range fs {
		if fn(v, i, fs) {
			r = append(r, v)
		}
	}

	return r
}

func main() {

	slice1 := []int{1, 2, 3, 4, 5}

	var doubleIt MapFunction[int, int] = func(v int, _ int, _ []int) int {
		return v * 2
	}

	var everySecond FilterFunction[int] = func(_ int, i int, _ []int) bool {
		return (i%2 == 0)
	}

	doubles := FSlice(slice1).Map(doubleIt)
	staggered := FSlice(slice1).Filter(everySecond).Map(doubleIt).Map(doubleIt)

	fmt.Println(doubles)
	fmt.Println(staggered)
}
