package main

import "fmt"

func maxOperations(nums []int, k int) int {
	res, cnt := 0, map[int]int{}
	for _, n := range nums {
		if cnt[k-n] > 0 {
			cnt[k-n] -= 1
			res += 1
		} else {
			cnt[n] += 1
		}
	}
	return res
}

func main() {
	k := 3
	nums := []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2, 3}
	fmt.Println(maxOperations(nums, k))
}
