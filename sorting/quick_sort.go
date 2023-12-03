package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func partition(arr []int, i, j int) int {
	p, w := i, i

	for k := i + 1; k < j; k++ {
		if arr[k] < arr[p] {
			w++
			arr[k], arr[w] = arr[w], arr[k]
		}
	}
	arr[p], arr[w] = arr[w], arr[p]

	return w
}

func quickSort(arr []int, i, j int) {
	if i == j || len(arr) == 1 {
		return
	}
	pivot := partition(arr, i, j)
	quickSort(arr, i, pivot)
	quickSort(arr, pivot+1, j)
}

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr))
}

func main() {
	rand.NewSource(time.Now().Unix())
	n := 1000

	//arr := make([]int, 0, n)
	//for i := 0; i < n; i++ {
	//	arr = append(arr, rand.Intn(100)+1)
	//}
	arr := []int{50, 52, 82, 85, 25, 34}

	arrCopy := make([]int, n)
	copy(arrCopy, arr)
	start := time.Now()
	sort.Ints(arrCopy)
	fmt.Println("sort.Sort time:", time.Since(start))

	start = time.Now()
	QuickSort(arr)
	fmt.Println("QuickSort time:", time.Since(start))

	for i := 0; i < n; i++ {
		if arr[i] != arrCopy[i] {
			fmt.Println("Error while testing sorting")
			os.Exit(1)
		}
	}

	fmt.Println("Success!")
}
