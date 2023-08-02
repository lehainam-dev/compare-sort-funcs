package main

import (
	"math/rand"
	"sort"
	"testing"
)

var denseArr []int
var spareArr []int

func BenchmarkSortSlice(b *testing.B) {
	var testcases = []struct {
		name  string
		input []int
	}{
		{name: "large array small values", input: denseArr},
		{name: "large array large values", input: spareArr},
	}

	for _, t := range testcases {
		b.Run(t.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := make([]int, len(t.input))
				copy(arr, t.input)
				b.StartTimer()
				sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
			}
		})
	}
}

func BenchmarkSortStable(b *testing.B) {
	var testcases = []struct {
		name  string
		input []int
	}{
		{name: "large array small values", input: denseArr},
		{name: "large array large values", input: spareArr},
	}

	for _, t := range testcases {
		b.Run(t.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := make([]int, len(t.input))
				copy(arr, t.input)
				b.StartTimer()
				sort.SliceStable(arr, func(i, j int) bool { return arr[i] < arr[j] })
			}
		})
	}
}

func BenchmarkSortIndex(b *testing.B) {
	var testcases = []struct {
		name  string
		input []int
	}{
		{name: "large array small values", input: denseArr},
		{name: "large array large values", input: spareArr},
	}

	for _, t := range testcases {
		b.Run(t.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := make([]int, len(t.input))
				copy(arr, t.input)
				b.StartTimer()
				sort.Slice(arr, func(i, j int) bool {
					if arr[i] == arr[j] {
						return i < j
					}
					return arr[i] < arr[j]
				})
			}
		})
	}
}

func init() {
	n := 500_000
	for i := 0; i < n; i++ {
		denseArr = append(denseArr, rand.Intn(1_000))
		spareArr = append(spareArr, rand.Intn(1_000_000_000))
	}
}
