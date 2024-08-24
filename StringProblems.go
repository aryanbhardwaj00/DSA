package dsa

import (
	"fmt"
	"strings"
	"unicode"
)

func isVowel(s string) bool {
	s = strings.ToLower(s)
	switch s {
	case "a", "e", "i", "o", "u":
		return true
	default:
		return false
	}
}

func changeCase(str string) {
	// Convert string to a rune
	// Create an empty string to store indivisual strings
	// Run a loop over rune
	// If current character is vowel ,
	// If it is Upper case , change to lower case and add it to result string that we created
	// If it is lower case , change to Upper case and add it to result
	// If it is not a vowel just add it to result string
	runes := []rune(str)
	resultString := ""
	for _, value := range runes {
		if isVowel(string(value)) {
			if unicode.IsUpper(value) {
				resultString += strings.ToLower(string(value))
			} else {
				resultString += strings.ToUpper(string(value))
			}
		} else {
			resultString += string(value)
		}
	}
	fmt.Println(resultString)
}

func changeCase1(str string) {
	// Convert string to a rune
	// Create an empty string to store indivisual strings
	// Run a loop over rune
	// If current character is vowel ,
	// If it is Upper case , change to lower case and add it to result string that we created
	// If it is lower case , change to Upper case and add it to result
	// If it is not a vowel just add it to result string
	runes := []rune(str)
	for idx, value := range str {
		if isVowel(string(value)) {
			if unicode.IsUpper(value) {
				tempStr := strings.ToUpper(string(value))
				runes[idx] = []rune(tempStr)[0]
			} else {
				tempstr := strings.ToUpper(string(value))
				runes[idx] = []rune(tempstr)[0]
			}
		} else {
			runes[idx] = value
		}
	}
	fmt.Println(string(runes))
}

func strPallindrome(str string) bool {
	// Find Length
	length := len(str)
	var res string
	// Iterate through each character from Backward
	for i := length - 1; i >= 0; i-- {
		res += string(str[i])
	}
	if string(res) == str {
		return true
	}
	return false
}

func strPallindrome2(str string) bool {
	var tStr []rune = []rune(str)
	for i, j := 0, len(tStr)-1; i < len(tStr)/2; i, j = i+1, j-1 {
		if tStr[i] != tStr[j] {
			return false
		}
	}
	return true
}

func isConsonant(str string) string {
	runes := []rune(str)
	for index, char := range runes {
		if !isVowels(char) {
			runes[index] = unicode.ToUpper(char)
		}
	}
	return string(runes)
}

func isVowels(str rune) bool {
	switch str {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func consonants(str string) string {
	runes := []rune(str)
	result := ""
	for _, char := range runes {
		if !isVowels(char) {
			result += string(char)
		}
	}
	return result
}

func paranthesis(str string) string {
	result := ""
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '(' {
			if count > 0 {
				result += string(str[i])
			}
			count++
		} else {
			count--
			if count > 0 {
				result += string(str[i])
			}
		}
	}
	return result
}

func reverseStr(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		} else {
			result += string(s[i])
		}
	}
	return result
}

func returnVowels(s string) string {
	// Create an empty string
	// Run a for loop through the string
	// If the present character is vowel and its preceding character and succeeding are not Vowel
	// Then add that character to result
	// Special Cases , Two or more Continuos Vowels Surrounded by Consonants
	// If previous character is Consonant and next character is vowel ,add both to result
	// If previous character was vowel and next character is consonant add it to result
	result := ""
	runes := []rune(s)
	for i := 1; i < len(runes)-1; i++ {
		if isVowels(runes[i]) {
			if !isVowels(runes[i+1]) && !isVowels(runes[i-1]) {
				result += string(runes[i])
			} else if isVowels(runes[i+1]) && i != len(runes)-1 {
				result += string(runes[i])
			} else if !isVowels(runes[i+1]) && isVowels(runes[i-1]) {
				result += string(runes[i])
			}
		}
	}
	return result
}
