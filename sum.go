package sum

import "runtime"

func SumInts(ints []int) (sum int) {
	parallel := min(runtime.NumCPU(), runtime.GOMAXPROCS(0)) // degree of parallelism
	// estimated by "go test -v -run=TestCrossover" with various GOMAXPROCS settings
	if parallel == 1 || (parallel == 2 && len(ints) < 44000) || len(ints) < 22000 {
		sum = SumIntsSerial(ints)
	} else {
		sum = SumIntsParallel(ints)
	}
	return
}

func SumIntsSerial(ints []int) (sum int) {
	for _, v := range ints {
		sum += v
	}
	return
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func SumIntsParallel(ints []int) (sum int) {
	parallel := min(runtime.NumCPU(), runtime.GOMAXPROCS(0)) // degree of parallelism
	result := make(chan int, parallel)
	chunk := (len(ints) + parallel - 1) / parallel
	go func() {
		for start := 0; start < len(ints); start += chunk {
			go func(k []int) {
				result <- SumIntsSerial(k)
			}(ints[start:min(start+chunk, len(ints))])
		}
	}()

	for start := 0; start < len(ints); start += chunk {
		sum += <-result
	}
	close(result)
	return
}
