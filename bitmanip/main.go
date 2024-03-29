package main

import (
	"fmt"
	"math/bits"
)

// rwxrob -> read Paulo Freire
const (
	disconnected = iota
	connecting
	reconnecting
	connected
)

type Client interface {
	IsConnected()
	Connect()
}

type client struct {
	status uint
}

func (c *client) IsConnected() bool {
	status := uint32(c.status)
	switch {
	case status == connected:
		return true
	case status > connecting:
		return true
	case status == connecting:
		return true
	default:
		return false
	}
}

func swap(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}

func otherswap(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func genericswap[T any](a, b T) (any, any) {
	a, b = b, a
	return a, b
}

func flipBits() {
	var n byte = 0x0F
	fmt.Printf("%016b\n", n)
	n = ^n
	fmt.Printf("%016b\n", n)
	m := uint16(n)

	fmt.Printf("Len16(%016b) = %d\n", 8, bits.Len16(8))
	fmt.Printf("Len16(%016b) = %d\n", m, bits.Len16(m))
}

func isEven(n int) bool {
	if n&1 == 1 {
		return false
	}
	return true
}

func isEvenTwo(n int) bool {
	if n&1 != 1 {
		return true
	}
	return false
}

func powerOfTwo(n int) bool {
	if n&(n-1) == 0 {
		return true
	}
	return false
}

// right shift by 2
func divByTwo(num uint) uint {
	return num >> 1
}

func mulByTwo(num uint) uint {
	return num << 1
}

func add(a, b int64) int64 {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

func main() {

	/* client1 := client{status: connected}
	client2 := client{status: disconnected}

	fmt.Println(client1.IsConnected())
	fmt.Println(client2.IsConnected())

	client1.status = disconnected
	fmt.Println(client1.IsConnected())
	client1.status = connected
	fmt.Println(client1.IsConnected())
	fmt.Println(swap(10, 12))
	fmt.Println(swap(10000, 100))
	fmt.Println(otherswap(10, 12))
	fmt.Println(genericswap(10, 12))
	fmt.Println(genericswap("hello", "good bye"))
	var n byte = 0x0F
	fmt.Println(n)
	fmt.Println(genericswap(n, n))
	fmt.Println(isEven(15))
	fmt.Println(mulByTwo(2))
	fmt.Println(divByTwo(3))
	fmt.Println(add(3, 9)) */
	for i := 0; i < 10; i++ {
		fmt.Printf("%d Is Even: %v\n", i, isEven(i))
		fmt.Printf("%d Is Even: %v\n", i, isEvenTwo(i))
	}
}
