package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func selectionSort(arr []int) []int { // O(n^2)
	for i := 0; i < len(arr); i++ { // O(n)
		idxMin := i
		for j := i + 1; j < len(arr); j++ { // O((n - 1) / 2 - 1)
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
	return arr
}

func main() {
	rand.NewSource(time.Now().Unix())
	n := 100000

	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(1000))
	}

	arrCopy := make([]int, n)
	copy(arrCopy, arr)
	start := time.Now()
	sort.Ints(arrCopy)
	fmt.Println("sort.Sort time:", time.Since(start))

	start = time.Now()
	arr = selectionSort(arr)
	fmt.Println("selectionSort time:", time.Since(start))

	for i := 0; i < n; i++ {
		if arr[i] != arrCopy[i] {
			fmt.Println("Error while testing sorting")
			os.Exit(1)
		}
	}

	fmt.Println("Success!")
}
