package sum

import "runtime"

func SumInts(ints []int) int {
	if len(ints) < 28728 { // estimated by TestCrossover using "go test -v"
		return SumIntsSerial(ints)
	}
	return SumIntsParallel(ints)
}

func SumIntsSerial(ints []int) int {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func SumIntsParallel(ints []int) int {
	ncpus := runtime.NumCPU() // degree of parallelism
	result := make(chan int, ncpus)
	chunk := (len(ints) + ncpus - 1) / ncpus
	go func() {
		for start := 0; start < len(ints); start += chunk {
			go func(k []int) {
				result <- SumIntsSerial(k)
			}(ints[start:min(start+chunk, len(ints))])
		}
	}()

	sum := 0
	for start := 0; start < len(ints); start += chunk {
		sum += <-result
	}
	close(result)

	return sum
}
