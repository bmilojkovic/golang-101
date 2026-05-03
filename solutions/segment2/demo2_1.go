package main

import "fmt"

func clean(s string) (string, bool) {

	start := 0
	end := len(s) - 1

	for start < end && s[start] == ' ' {
		start++
	}
	for end > start && s[end] == ' ' {
		end--
	}

	// stringovi su immutable
	b := []byte(s[start : end+1])
	for i, c := range b {
		if c >= 'A' && c <= 'Z' {
			b[i] = c + 32
		}
	}
	cleaned := string(b)

	// Ako se importuje strings moze i ovako:
	// cleaned := strings.ToLower(strings.TrimSpace(s))

	return cleaned, cleaned != s

}

func isPalindrome(s string) bool {
	i := 0
	for i < len(s)/2 {
		if s[i] != s[len(s)-1-i] {
			return false
		}
		i++
	}
	return true
}

func countVowels(s string) int {
	vowels := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			vowels++
		}
	}
	return vowels
}

func removeSpaces(s string) (string, int) {
	result := ""
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			count++
		} else {
			result += string(s[i])
		}
	}
	return result, count
}

func replaceVowels(s string, replacement string) (string, int) {
	result := ""
	count := 0
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			count++
			result += replacement
		} else {
			result += string(s[i])
		}
	}
	return result, count
}

func truncate(s string, max int) (string, bool) {
	if len(s) > max {
		return s[:max], true
	} else {
		return s, false
	}
}

func capitaliseWords(s string) (string, int) {
	result := ""
	count := 0
	newWord := true

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' {
			newWord = true
			result += string(c)
		} else if newWord {
			if c >= 'a' && c <= 'z' {
				c = c - 32
			}
			result += string(c)
			count++
			newWord = false
		} else {
			result += string(c)
		}
	}
	return result, count
}

func main() {

	var str string = "   pOZDRAV SveTo   "
	fmt.Println(clean(str))
	fmt.Println("Palindrom 1, ", str, ": ", isPalindrome(str))
	fmt.Println("Palindrom 2, anavolimilovana: ", isPalindrome("anavolimilovana"))

	fmt.Println("Vowels in word string GOLANG RADIONICA: ", countVowels("golang radionica"))
	fmt.Println(removeSpaces(str))
	fmt.Println(replaceVowels("hello world", "++"))

	fmt.Println(truncate("hello world", 5))
	fmt.Println(truncate("hi", 5))
	fmt.Println(capitaliseWords("hello world"))
}
