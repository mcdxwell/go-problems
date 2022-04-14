package palindrome

import "fmt"

type Checker interface {
	int | string
}

//func String[T Pal](s T) string { return fmt.Sprintf("%v", s)}

func Check[T int | string](in T) bool {
	s := fmt.Sprintf("%v", in)
	for i := 0; i < len(s)/2+1; i++ {
		j := len(s) - i - 1
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
