package main

import "fmt"

func IterBSearch(arr []int, l, r, x int) int {
	m := 0
	for l <= r {

		m = l + (r-l)/2

		if arr[m] == x {
			return m
		}

		if arr[m] < x {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}

func RecurBSearch(arr []int, l, r, x int) int {
	m := 0
	if r >= l {

		m = l + (r-l)/2

		if arr[m] == x {
			return m
		}

		if arr[m] > x {
			return RecurBSearch(arr, l, m-1, x)
		}

		return RecurBSearch(arr, m+1, r, x)
	}

	return -1
}

func main() {
	arr := []int{3, 4, 5, 6, 7, 8, 9}
	fmt.Println(arr)
	n := len(arr)
	x := 9
	l := 0
	r := n - 1
	result := IterBSearch(arr, l, r, x)
	fmt.Println(result)
	result = RecurBSearch(arr, l, r, x)
	fmt.Println(result)
}
