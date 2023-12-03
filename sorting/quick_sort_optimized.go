package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func insertionSort(arr []int, start, end int) {
	for i := start + 1; i < end; i++ {
		curr := arr[i]
		j := i - 1
		for j >= start && arr[j] > curr {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = curr
	}
}

func partitionOptimized(arr []int, i, j int) int {
	p, w := rand.Intn(j-i)+i, i
	arr[p], arr[i] = arr[i], arr[p]
	p = i

	for k := i + 1; k < j; k++ {
		if arr[k] < arr[p] {
			w++
			arr[k], arr[w] = arr[w], arr[k]
		}
	}
	arr[p], arr[w] = arr[w], arr[p]

	return w
}

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

func quickSortOptimized(arr []int, i, j int) {
	if i == j || len(arr) == 1 {
		return
	}
	if j-i <= 12 {
		insertionSort(arr, i, j)
		return
	}
	pivot := partitionOptimized(arr, i, j)
	quickSortOptimized(arr, i, pivot)
	quickSortOptimized(arr, pivot+1, j)
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

func QuickSortOptimized(arr []int) {
	quickSortOptimized(arr, 0, len(arr))
}

func main() {
	rand.NewSource(time.Now().Unix())
	n := 1000000

	arr := make([]int, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100)+1)
	}

	arrCopy := make([]int, n)
	arr2 := make([]int, n)
	copy(arrCopy, arr)
	copy(arr2, arr)

	start := time.Now()
	sort.Ints(arrCopy)
	fmt.Println("sort.Sort time:", time.Since(start))

	start = time.Now()
	QuickSort(arr)
	fmt.Println("QuickSort time:", time.Since(start))

	start = time.Now()
	QuickSortOptimized(arr2)
	fmt.Println("QuickSortOptimized time:", time.Since(start))

	for i := 0; i < n; i++ {
		if arr2[i] != arrCopy[i] {
			fmt.Println("Error while testing sorting")
			os.Exit(1)
		}
	}

	fmt.Println("Success!")
}
