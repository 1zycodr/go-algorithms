package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func radixSort(arr []int) []int {
	buckets := [10][]int{}
	for pow := 1; pow <= 10000; pow *= 10 {
		buckets = [10][]int{}
		for _, n := range arr {
			idx := n / pow % 10
			buckets[idx] = append(buckets[idx], n)
		}
		k := 0
		for i := 0; i < 10; i++ {
			for _, n := range buckets[i] {
				arr[k] = n
				k++
			}
			buckets[i] = []int{}
		}
	}
	return arr
}

func main() {
	rand.NewSource(time.Now().Unix())
	n := 1000000

	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100000))
	}

	arrCopy := make([]int, n)
	copy(arrCopy, arr)

	start := time.Now()
	sort.Ints(arrCopy)
	fmt.Println("sort.Sort time:", time.Since(start))

	start = time.Now()
	arr = radixSort(arr)
	fmt.Println("radixSort time:", time.Since(start))

	for i := 0; i < n; i++ {
		if arr[i] != arrCopy[i] {
			fmt.Println("Error while testing sorting")
			os.Exit(1)
		}
	}

	fmt.Println("Success!")
}
