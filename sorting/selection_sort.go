package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func selectionSort(arr []int) []int { // O(n^2)
	fmt.Println(arr)
	for i := 0; i < len(arr)-1; i++ {
		indMin := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[indMin] {
				indMin = j
			}
		}
		arr[i], arr[indMin] = arr[indMin], arr[i]
		fmt.Println(arr)
	}
	return arr
}

func main() {
	rand.NewSource(time.Now().Unix())
	n := 10

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
