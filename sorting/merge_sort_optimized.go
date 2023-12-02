package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func merge(arr1, arr2 []int) []int {
	res := []int{}
	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] > arr2[j] {
			res = append(res, arr2[j])
			j++
		} else {
			res = append(res, arr1[i])
			i++
		}
	}
	res = append(res, arr1[i:]...)
	res = append(res, arr2[j:]...)
	return res
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	left := make([]int, len(arr)/2)
	right := make([]int, len(arr)-len(arr)/2)
	copy(left, arr[:len(arr)/2])
	copy(right, arr[len(arr)/2:])
	left = mergeSort(left)
	right = mergeSort(right)
	return merge(left, right)
}

func mergeSortOptimized(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	if len(arr) <= 10 {
		return insertionSort(arr)
	}
	left := make([]int, len(arr)/2)
	right := make([]int, len(arr)-len(arr)/2)
	copy(left, arr[:len(arr)/2])
	copy(right, arr[len(arr)/2:])
	left = mergeSortOptimized(left)
	right = mergeSortOptimized(right)
	return merge(left, right)
}

func insertionSort(arr []int) []int { // O(n^2)
	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		j := i - 1
		for j >= 0 && curr < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = curr
	}
	return arr
}

func main() {
	rand.NewSource(time.Now().Unix())
	n := 1000000

	arr := make([]int, 0)
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
	arr = mergeSort(arr)
	fmt.Println("mergeSort time:", time.Since(start))

	start = time.Now()
	arr = mergeSortOptimized(arr)
	fmt.Println("mergeSortOptimized time:", time.Since(start))

	for i := 0; i < n; i++ {
		if arr[i] != arrCopy[i] {
			fmt.Println("Error while testing sorting")
			os.Exit(1)
		}
	}

	fmt.Println("Success!")
}
