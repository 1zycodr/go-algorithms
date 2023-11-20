package main

import "fmt"

func main() {
	set := []int{1, 2, 3}
	for mask := 0; mask < (1 << len(set)); mask++ {
		for i := 0; i < len(set); i++ {
			if mask&(1<<i) != 0 {
				fmt.Printf("%d ", set[i])
			}
		}
		fmt.Println()
	}
}
