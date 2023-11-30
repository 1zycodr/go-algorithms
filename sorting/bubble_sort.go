package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func bubbleSort(arr []int) []int { // O(n^2)
	for i := 0; i < len(arr)-1; i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

func main() {
	rand.NewSource(time.Now().Unix())
	n := 10000

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
	arr = bubbleSort(arr)
	fmt.Println("bubbleSort time:", time.Since(start))

	for i := 0; i < n; i++ {
		if arr[i] != arrCopy[i] {
			fmt.Println("Error while testing sorting")
			os.Exit(1)
		}
	}

	fmt.Println("Success!")
}
