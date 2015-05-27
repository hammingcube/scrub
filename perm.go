package main

import _ "fmt"
import "math/rand"

func NewJumbler() func(rune) rune {
	const LOWER = "abcdefghijklmnopqrstuvwxyz"
	const UPPER = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const DIGITS = "0123456789"
	arr := rand.Perm(len(LOWER))
	arr2 := rand.Perm(len(DIGITS))
	//fmt.Println(arr)
	//fmt.Println(arr2)
	fn := func(ch rune) rune {
		switch {
		  case 'a' <= ch && ch <= 'z': return rune(LOWER[arr[ch-'a']])
		  case 'A' <= ch && ch <= 'Z': return rune(UPPER[arr[ch-'A']])
		  case '0' <= ch && ch <= '9': return rune(DIGITS[arr2[ch-'0']])
		  default: return ch
		}
	}
	return fn
}