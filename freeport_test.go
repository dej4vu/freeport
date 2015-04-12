package freeport

import (
	"net"
	"strconv"
	"testing"
)

func TestGet(t *testing.T) {
	port, err := Get()
	if err != nil {
		t.Fatalf("Got err from Get: %s", err)
	}
	if port == 0 {
		t.Fatal("Got port 0")
	}
}

func TestWithin(t *testing.T) {
	start := 10000
	end := 10050
	port, err := Within(start, end)
	if err != nil {
		t.Fatalf("Got err from GetRange: %s", err)
	}
	if port > end || port < start {
		t.Fatalf("Got port %d, expect port in range from %d to %d", port, start, end)
	}
}

func TestFirst(t *testing.T) {
	ps := []int{80, 8080, 8088,9000}
	port, err := First(ps...)
	if err != nil {
		t.Fatalf("Got err: %s from First", err)
	}
	var ok bool
	for _, p := range ps {
		if p == port {
			ok = true
			break
		}
	}
	if !ok {
		t.Fatalf("Got port %d, but not in %v", port, ps)
	}
}

func TestGetClose(t *testing.T) {
	port, err := Get()
	if err != nil {
		t.Fatalf("Got err: %s from Get", err)
	}
	if port == 0 {
		t.Fatal("Got port 0 from Get")
	}
	listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		t.Fatalf("Got err from net.Listen: %s", err)
	}
	listener.Close()
}
