# Benchmark between sort functions

TIL that `sort.SliceStable` performs poorly compare with `sort.Slice`.

I write a simple benchmark to test this.

I also wonder the performance between using `sort.SliceStable` and `sort.Slice` with compare index if elements are equal.

Here is the benchmark result:

```text
goos: darwin
goarch: arm64
pkg: compare-sort-funcs
BenchmarkSortSlice/large_array_small_values-10               459          26184017 ns/op
BenchmarkSortSlice/large_array_large_values-10               240          49681642 ns/op
BenchmarkSortStable/large_array_small_values-10               96         124684752 ns/op
BenchmarkSortStable/large_array_large_values-10               78         154809093 ns/op
BenchmarkSortIndex/large_array_small_values-10               273          43727234 ns/op
BenchmarkSortIndex/large_array_large_values-10               240          49809745 ns/op
PASS
ok      compare-sort-funcs      89.851s
```

`sort.SliceStable` slower than `sort.Slice` about 3~4 times.

`sort.Slice` with index perform near as same as default `sort.Slice`.
