package main

import "fmt"

func main() {
	fmt.Println(SaveAndMiss("123456789", 1))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))
	fmt.Println(SaveAndMiss("", 3))
	fmt.Println(SaveAndMiss("hello you all ! ", 0))
	fmt.Println(SaveAndMiss("what is your name?", 0))
	fmt.Println(SaveAndMiss("e 5£ @ 8* 7 =56 ;", 2))
	fmt.Println(SaveAndMiss("QKplq%QSw", 3))
	fmt.Println(SaveAndMiss("hello \\! n4ght cr3a8ure7 ", 5))
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 8))
	fmt.Println(SaveAndMiss("Po65 4o", 10))
}
// the function SaveAndMiss used to save num character and miss or scipe num character
func SaveAndMiss(arg string, num int) string {
	if len(arg) == 0 || num <= 0 {
		return arg
	}

	runes := []rune(arg)
	var res []rune
	for i := 0; i < len(runes); i += (num*2) {
		for j := i ; j < i+num && j < len(runes); j++ {
			res = append(res, rune(runes[j]))
		}
	}
	return string(res)
}
