package main

import (
	"fmt"
)

func binaryStrToInt(s string) (n int) {
	for i, c := range s {
		d := 0
		if c-'0' > 0 {
			d = 1
			for j := 0; j < len(s)-1-i; j++ {
				d *= 2
			}
		}
		n += d
	}
	return
}

func addBinary(a string, b string) string {
	length := len(b)
	if len(a) > len(b) {
		length = len(a)
	}
	r := make([]byte, length)
	var extra byte
	for i := 0; i < length; i++ {
		var dib byte
		ai, bi := len(a)-1-i, len(b)-1-i
		if ai >= 0 && bi >= 0 {
			dib = a[ai] - '0' + b[bi] - '0'
		} else if ai >= 0 {
			dib = a[ai] - '0'
		} else if bi >= 0 {
			dib = b[bi] - '0'
		}
		dib += extra
		if dib > 1 {
			extra = 1
		} else {
			extra = 0
		}
		r[length-1-i] = dib%2 + '0'
	}
	if extra == 1 {
		r = append([]byte{'1'}, r...)
	}
	return string(r)
}

func main() {
	for _, s := range []string{"11", "1", "1010", "1011"} {
		fmt.Println(s, binaryStrToInt(s))
	}

	for _, array := range [][2]string{{"11", "1"}, {"1010", "1011"}} {
		a := array[0]
		b := array[1]
		r := addBinary(a, b)
		fmt.Println(a, b, r)
	}
}
