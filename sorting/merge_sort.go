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

func main() {
	rand.NewSource(time.Now().Unix())
	n := 10000000

	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(100)+1)
	}

	arrCopy := make([]int, n)
	copy(arrCopy, arr)
	start := time.Now()
	sort.Ints(arrCopy)
	fmt.Println("sort.Sort time:", time.Since(start))

	start = time.Now()
	arr = mergeSort(arr)
	fmt.Println("mergeSort time:", time.Since(start))

	for i := 0; i < n; i++ {
		if arr[i] != arrCopy[i] {
			fmt.Println("Error while testing sorting")
			os.Exit(1)
		}
	}

	fmt.Println("Success!")
}
