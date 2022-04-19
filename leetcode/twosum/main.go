package main

import "fmt"

func twoSum(array []int, target int) []int {
	hm := make(map[int]int)

	for i, num := range array {
		_, ok := hm[num]
		if ok {
			return []int{i, hm[num]}
		}

		hm[target-num] = i
	}

	return nil
}

func main() {

	target := 8
	array := []int{5, 2, 1, -3, 3, 5, 10}
	fmt.Println(twoSum(array, target))
}
