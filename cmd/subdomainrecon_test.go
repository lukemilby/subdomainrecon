package main

import (
	"testing"
)


func newSubDomainsForTesting() map[string]subdomain{
	return map[string]subdomain{
		"www.lol.com": {
			ip:[]string{},
			source: []string{},
		},
	}
}

func TestPopulateIpAddresses(t *testing.T) {
	subDomains := newSubDomainsForTesting()

	populateIpAddresses(subDomains)

	got := len(subDomains["www.lol.com"].ip[0])
	want := 1
	if got <= want {
		t.Errorf("got %d is not greater that want %d ", got, want)
	}
}

func BenchmarkPopulateIpAddresses(b *testing.B) {
	subDomains := newSubDomainsForTesting()

	b.StartTimer()

	for i:= 0 ; i < b.N ;  i++ {
		populateIpAddresses(subDomains)
	}


}