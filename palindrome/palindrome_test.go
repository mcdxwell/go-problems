package palindrome_test

import (
	"testing"

	"github.com/mcdxwell/go-problems/palindrome"
)

// WIP
func IsPalindromeTest(t *testing.T) {

	items := []int{1001, 1000, 3333}
	answ := []bool{true, false, true}
	for i, v := range items {
		res := palindrome.Check(v)

		if res != answ[i] {
			t.Errorf("Error! Expected %v. Got: %v", answ[i], res)
		}
	}
}
