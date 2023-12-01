package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"sort"
	"time"
)

func countingSort(arr []int) []int { // O(n + m), m = max(arr) - min(arr) + 1
	max, min := math.MinInt32, math.MaxInt32
	for _, v := range arr { // O(n)
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	temp := make([]int, max-min+1)
	for _, v := range arr { // O(n)
		temp[v-min]++
	}
	k := 0
	for i, v := range temp { // O(m), m = max(arr) - min(arr) + 1
		for v != 0 {
			v--
			arr[k] = i + min
			k++
		}
	}
	return arr
}

func main() {
	rand.NewSource(time.Now().Unix())
	n := 100000

	arr := make([]int, 0)
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(10000000)+1)
	}

	arrCopy := make([]int, n)
	copy(arrCopy, arr)
	start := time.Now()
	sort.Ints(arrCopy)
	fmt.Println("sort.Sort time:", time.Since(start))

	start = time.Now()
	arr = countingSort(arr)
	fmt.Println("countingSort time:", time.Since(start))

	for i := 0; i < n; i++ {
		if arr[i] != arrCopy[i] {
			fmt.Println("Error while testing sorting")
			os.Exit(1)
		}
	}

	fmt.Println("Success!")
}
