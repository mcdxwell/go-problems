package main

import (
	"fmt"
	"math/bits"
)

//Paulo Freire
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

func main() {
	var n byte = 0x0F
	fmt.Printf("%016b\n", n)
	n = ^n
	fmt.Printf("%016b\n", n)
	m := uint16(n)

	fmt.Printf("Len16(%016b) = %d\n", 8, bits.Len16(8))
	fmt.Printf("Len16(%016b) = %d\n", m, bits.Len16(m))

	client1 := client{status: connected}
	client2 := client{status: disconnected}

	fmt.Println(client1.IsConnected())
	fmt.Println(client2.IsConnected())

	client1.status = disconnected
	fmt.Println(client1.IsConnected())
	client1.status = connected
	fmt.Println(client1.IsConnected())

}
