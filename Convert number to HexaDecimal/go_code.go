package main

import "fmt"

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	if num < 0 {
		num = num + (1 << 32) // two's complement
	}

	chars := "0123456789abcdef"
	res := []byte{}

	for num > 0 {
		res = append([]byte{chars[num&0xf]}, res...) // prepend
		num >>= 4
	}

	return string(res)
}

func main() {
	fmt.Println(toHex(26))   // 1a
	fmt.Println(toHex(-1))   // ffffffff
	fmt.Println(toHex(0))    // 0
	fmt.Println(toHex(255))  // ff
}
