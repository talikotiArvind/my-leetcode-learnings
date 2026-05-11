package main

import "fmt"

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	result := []byte{}

	for i >= 0 || j >= 0 || carry > 0 {
		a, b := 0, 0
		if i >= 0 {
			a = int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			b = int(num2[j] - '0')
			j--
		}
		total := a + b + carry
		carry = total / 10
		result = append(result, byte(total%10)+'0')
	}

	for l, r := 0, len(result)-1; l < r; l, r = l+1, r-1 {
		result[l], result[r] = result[r], result[l]
	}

	return string(result)
}

func main() {
	fmt.Println(addStrings("11", "123"))   // "134"
	fmt.Println(addStrings("456", "77"))   // "533"
	fmt.Println(addStrings("999", "1"))    // "1000"
	fmt.Println(addStrings("0", "0"))      // "0"
	fmt.Println(addStrings("1", "9999"))   // "10000"
}
