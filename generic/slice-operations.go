package generic

import (
	"fmt"
	"strconv"
)

func Filter[T comparable](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Reduce[T1, T2 comparable](s []T1, init T2, f func(T1, T2) T2) T2 {
	var r = init
	for _, v := range s {
		r = f(v, r)
	}

	return r
}

func Map[T1, T2 comparable](s []T1, f func(T1) T2) []T2 {
	var r []T2
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

func Test9() {
	s := []string{"foo", "bar", "baz"}

	r := Filter(s, func(v string) bool {
		if v == "foo" {
			return true
		}
		return false
	})
	fmt.Println(r)

	i := []int{1, 2, 3, 4, 5}
	sum := Reduce(i, 110, func(v int, s int) int {
		return v + s
	})
	fmt.Println(sum)

	strSlice := []string{"1", "2", "3"}
	intSlice := Map(strSlice, func(s string) int {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return i
	})
	fmt.Println(intSlice)
}
