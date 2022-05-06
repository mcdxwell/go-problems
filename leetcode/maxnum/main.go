package main

import "fmt"

func maxOperations(nums []int, k int) int {
	res, cnt := 0, map[int]int{}
	for _, n := range nums {
		if cnt[k-n] > 0 {
			cnt[k-n] = addOne(cnt[k-n], -1)
			res = addOne(res, 1)
		} else {
			cnt[n] = addOne(cnt[n], 1)
		}
	}
	return res
}

func addOne(a, b int) int {

	for b != 0 {

		// common bits of a and b go to carry
		carry := a & b

		// xor - sum bits where one is not set
		a = a ^ b

		// shift carry by 1
		b = carry << 1
	}

	return a

}

func main() {
	k := 3
	nums := []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2, 3}
	fmt.Println(maxOperations(nums, k))
	fmt.Println(addOne(10, 1))
}
