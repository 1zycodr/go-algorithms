package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

func insertionSort(arr []int) []int { // O(n^2)
	//fmt.Println(arr)
	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		j := i - 1
		for j >= 0 && curr < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = curr
		//fmt.Println(arr)
	}
	return arr
}

func main() {
	rand.NewSource(time.Now().Unix())
	n := 1000

	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(10))
	}

	arrCopy := make([]int, n)
	copy(arrCopy, arr)
	start := time.Now()
	sort.Ints(arrCopy)
	fmt.Println("sort.Sort time:", time.Since(start))

	start = time.Now()
	arr = insertionSort(arr)
	fmt.Println("insertionSort time:", time.Since(start))

	for i := 0; i < n; i++ {
		if arr[i] != arrCopy[i] {
			fmt.Println("Error while testing sorting")
			os.Exit(1)
		}
	}

	fmt.Println("Success!")
}
