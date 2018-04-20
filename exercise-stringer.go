// +build ignore

package main

import (
	"fmt"
)

type IPAddr [4]byte

func (host IPAddr) String() (s string) {
	for _, b := range host {
		if s != "" {
			s += "."
		}
		s += fmt.Sprint(b)
	}
	return
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
