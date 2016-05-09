package sum

import (
	"runtime"
	"testing"
)

const LIMIT = 10000000

var increasing []int

func init() {
	increasing = make([]int, LIMIT)
	for i := range increasing {
		increasing[i] = i + 1
	}
}

func testSerial(b *testing.T, n int) {
	var ints []int
	if n <= LIMIT {
		ints = increasing[:n]
	} else {
		ints = make([]int, n)
		for i := range ints {
			ints[i] = i + 1
		}
	}

	computed := SumIntsSerial(ints)
	expected := (n * (n + 1)) / 2
	if computed != expected {
		b.Errorf("sum of 1..%d, value is %d; want %d", n, computed, expected)
	}
}

func TestSerial10(t *testing.T)      { testSerial(t, 10) }
func TestSerial100(t *testing.T)     { testSerial(t, 100) }
func TestSerial1000(t *testing.T)    { testSerial(t, 1000) }
func TestSerial10000(t *testing.T)   { testSerial(t, 10000) }
func TestSerial100000(t *testing.T)  { testSerial(t, 100000) }
func TestSerial1000000(t *testing.T) { testSerial(t, 1000000) }

func testParallel(b *testing.T, n int) {
	var ints []int
	if n <= LIMIT {
		ints = increasing[:n]
	} else {
		ints = make([]int, n)
		for i := range ints {
			ints[i] = i + 1
		}
	}

	computed := SumIntsParallel(ints)
	expected := (n * (n + 1)) / 2
	if computed != expected {
		b.Errorf("sum of 1..%d, value is %d; want %d", n, computed, expected)
	}
}

func TestParallel10(t *testing.T)      { testParallel(t, 10) }
func TestParallel100(t *testing.T)     { testParallel(t, 100) }
func TestParallel1000(t *testing.T)    { testParallel(t, 1000) }
func TestParallel10000(t *testing.T)   { testParallel(t, 10000) }
func TestParallel100000(t *testing.T)  { testParallel(t, 100000) }
func TestParallel1000000(t *testing.T) { testParallel(t, 1000000) }

func testSum(b *testing.T, n int) {
	var ints []int
	if n <= LIMIT {
		ints = increasing[:n]
	} else {
		ints = make([]int, n)
		for i := range ints {
			ints[i] = i + 1
		}
	}

	computed := SumInts(ints)
	expected := (n * (n + 1)) / 2
	if computed != expected {
		b.Errorf("sum of 1..%d, value is %d; want %d", n, computed, expected)
	}
}

func TestSum10(t *testing.T)      { testSum(t, 10) }
func TestSum100(t *testing.T)     { testSum(t, 100) }
func TestSum1000(t *testing.T)    { testSum(t, 1000) }
func TestSum10000(t *testing.T)   { testSum(t, 10000) }
func TestSum100000(t *testing.T)  { testSum(t, 100000) }
func TestSum1000000(t *testing.T) { testSum(t, 1000000) }

func benchmarkSerial(b *testing.B, n int) {
	var ints []int
	if n <= LIMIT {
		ints = increasing[:n]
	} else {
		ints = make([]int, n)
		for i := range ints {
			ints[i] = i + 1
		}
		b.ResetTimer()
	}

	for i := 0; i < b.N; i++ {
		_ = SumIntsSerial(ints)
	}
}

func BenchmarkSerial10(b *testing.B)         { benchmarkSerial(b, 10) }
func BenchmarkSerial100(b *testing.B)        { benchmarkSerial(b, 100) }
func BenchmarkSerial1000(b *testing.B)       { benchmarkSerial(b, 1000) }
func BenchmarkSerial10000(b *testing.B)      { benchmarkSerial(b, 10000) }
func BenchmarkSerial100000(b *testing.B)     { benchmarkSerial(b, 100000) }
func BenchmarkSerial1000000(b *testing.B)    { benchmarkSerial(b, 1000000) }
func BenchmarkSerial10000000(b *testing.B)   { benchmarkSerial(b, 10000000) }
func BenchmarkSerial100000000(b *testing.B)  { benchmarkSerial(b, 100000000) }
func BenchmarkSerial1000000000(b *testing.B) { benchmarkSerial(b, 1000000000) }

func benchmarkParallel(b *testing.B, n int) {
	var ints []int
	if n <= LIMIT {
		ints = increasing[:n]
	} else {
		ints = make([]int, n)
		for i := range ints {
			ints[i] = i + 1
		}
		b.ResetTimer()
	}

	for i := 0; i < b.N; i++ {
		_ = SumIntsParallel(ints)
	}
}

func BenchmarkParallel10(b *testing.B)         { benchmarkParallel(b, 10) }
func BenchmarkParallel100(b *testing.B)        { benchmarkParallel(b, 100) }
func BenchmarkParallel1000(b *testing.B)       { benchmarkParallel(b, 1000) }
func BenchmarkParallel10000(b *testing.B)      { benchmarkParallel(b, 10000) }
func BenchmarkParallel100000(b *testing.B)     { benchmarkParallel(b, 100000) }
func BenchmarkParallel1000000(b *testing.B)    { benchmarkParallel(b, 1000000) }
func BenchmarkParallel10000000(b *testing.B)   { benchmarkParallel(b, 10000000) }
func BenchmarkParallel100000000(b *testing.B)  { benchmarkParallel(b, 100000000) }
func BenchmarkParallel1000000000(b *testing.B) { benchmarkParallel(b, 1000000000) }

func benchmarkSum(b *testing.B, n int) {
	var ints []int
	if n <= LIMIT {
		ints = increasing[:n]
	} else {
		ints = make([]int, n)
		for i := range ints {
			ints[i] = i + 1
		}
		b.ResetTimer()
	}

	for i := 0; i < b.N; i++ {
		_ = SumInts(ints)
	}
}

func BenchmarkSum10(b *testing.B)         { benchmarkSum(b, 10) }
func BenchmarkSum100(b *testing.B)        { benchmarkSum(b, 100) }
func BenchmarkSum1000(b *testing.B)       { benchmarkSum(b, 1000) }
func BenchmarkSum10000(b *testing.B)      { benchmarkSum(b, 10000) }
func BenchmarkSum100000(b *testing.B)     { benchmarkSum(b, 100000) }
func BenchmarkSum1000000(b *testing.B)    { benchmarkSum(b, 1000000) }
func BenchmarkSum10000000(b *testing.B)   { benchmarkSum(b, 10000000) }
func BenchmarkSum100000000(b *testing.B)  { benchmarkSum(b, 100000000) }
func BenchmarkSum1000000000(b *testing.B) { benchmarkSum(b, 1000000000) }

// Estimate slice size where parallel evaluation is faster than serial evaluation
func TestCrossover(t *testing.T) {
	if runtime.NumCPU() < 2 || runtime.GOMAXPROCS(0) < 2 {
		t.Skip("Skipping Crossover test on single processor system")
	}

	s := func(n int) func(b *testing.B) {
		return func(b *testing.B) { benchmarkSerial(b, n) }
	}
	p := func(n int) func(b *testing.B) {
		return func(b *testing.B) { benchmarkParallel(b, n) }
	}

	// locate length range where parallel becomes faster than serial
	upper := 10
	for testing.Benchmark(s(upper)).NsPerOp() < testing.Benchmark(p(upper)).NsPerOp() {
		// t.Logf("serial/parallel crossover not below slice lengths %d and %d", upper, 10*upper)
		upper *= 10
		if upper > 100000 {
			t.Skip("Skipping Crossover test on underperformant system")
		}
	}
	lower := upper / 10
	t.Logf("serial/parallel crossover between slice length %d and %d", lower, upper)

	// locate crossover
	for upper-lower > 100 { // limit fineness of search given looseness of benchmark timing
		mid := (lower + upper) / 2
		if testing.Benchmark(s(mid)).NsPerOp() < testing.Benchmark(p(mid)).NsPerOp() {
			lower = mid
		} else {
			upper = mid
		}
		// t.Logf("serial/parallel crossover between slice length %d and %d", lower, upper)
	}
	t.Logf("slice length %d estimated as serial/parallel crossover for this computer", (lower+upper)/2)
}
