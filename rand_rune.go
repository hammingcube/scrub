package main

import "fmt"
import "unicode"
import "math/rand"
import "time"

func isAlphaNumeric(r rune) bool {
	return r == '_' || unicode.IsLetter(r) || unicode.IsDigit(r)
}

var upperCaseLetters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var lowerCaseLetters = []rune("abcdefghijklmnopqrstuvwxyz")
var digits = []rune("0123456789")

func randomChar(r *rand.Rand, runes []rune) rune {
	return runes[r.Intn(len(runes))]
}


func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var str string
	fmt.Scanf("%s", &str)
	fmt.Printf("Got %s\n", str)
	
	sl := []rune{}
	for _, ch := range str {
		var runes []rune
		if unicode.IsUpper(ch) {
			runes = upperCaseLetters
		} else if unicode.IsLower(ch) {
			runes = lowerCaseLetters
		} else if unicode.IsDigit(ch) {
			runes = digits
		} else {
			runes = []rune{}
		}
		var x rune
		if len(runes) > 0 {
			x = randomChar(r, runes)
		} else {
			x = ch
		}
		fmt.Printf("%s -> %s\n", 
			string(ch),
			string(x))
		sl = append(sl, x)
	}
	fmt.Println(string(sl))
}