package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func bubbleSort(arr []int) []int { // O(n^2)
	for i := 0; i < len(arr); i++ {
		sorted := true
		for j := 1; j < len(arr); j++ {
			if arr[j] < arr[j-1] {
				sorted = false
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		if sorted {
			break
		}
	}
	return arr
}

func main() {
	rand.NewSource(time.Now().Unix())
	n := 1000

	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100))
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
