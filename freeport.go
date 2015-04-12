// Package freeport provide an API to find a free port to bind to.
package freeport

import (
	"net"
	"strconv"
)

// Get a free port from the os randomly.
func Get() (port int, err error) {
	return within(0, 0)
}

// First return the  first of free port at choise
func First(ps ...int) (port int, err error) {
	for _,p := range ps {
		tp, err := within(p, p)
		if err == nil {
			port=tp
			break
		}
	}
	return port,err
}

// Within return the first free post from start to end.
func Within(start, end int) (port int, err error) {
	return within(start, end)
}

func within(start, end int) (port int, err error) {
	if start > end {
		start, end = end, start
	}
	var tempPort string
	for i := start; i <= end; i++ {
		tempPort = strconv.Itoa(i)
		listener, lErr := net.Listen("tcp", "127.0.0.1:"+tempPort)
		// break if find a free port
		if lErr == nil {
			_, portStr,sErr := net.SplitHostPort(listener.Addr().String())
			listener.Close()
			if sErr == nil {
				return strconv.Atoi(portStr)
			}else{
				err=sErr
			}
		}else{
			err=lErr
		}
	}
	return port,err
}
