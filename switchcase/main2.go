package main

import "fmt"

func SwitchCase(c byte) byte {
	c := 0
	if '0' <= c && c <= '9' {
		return c - '0'
	}
	if  'a' <= c && c <= 'f' {
		return c - 'a' + 32
	}
	if  'A' <= c && c <= 'F' {
		return c - 'A' + 32
	}
	return 0
}
