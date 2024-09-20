package main

import "fmt"

func main() {
	var str1, str2 string

	var num1 int = 12345

	var str string

	str = string(rune(num1))

	for i := 0; i < len(str); i++ {
		str1 += string(str[i])

	}
	for i := len(str) - 1; i >= 0; i-- {
		str2 += string(str[i])

	}
	if str1 == str2 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}
