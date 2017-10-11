package main

import "fmt"
import "strings"
import "strconv"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (t IPAddr) String() string {
	s := make([]string, len(t))
	for i := range t {
		s[i] = strconv.Itoa(int(t[i]))
	}
	return strings.Join(s, ".")
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
