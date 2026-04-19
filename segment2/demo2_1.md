## Functions

Implement the following string manipulation functions and test each of them separately with a simple main function:

```golang
// clean takes a string and returns a sanitised version of it with leading and
// trailing whitespace removed and all characters lowercased, along with a bool
// indicating whether the string was actually changed.
//
// Example: clean("  Hello  ") → ("hello", true)
//          clean("hello")     → ("hello", false)
func clean(s string) (string, bool)

// isPalindrome takes a string and returns true if it reads the same forwards
// and backwards, false otherwise. Assumes the input is already clean (no
// spaces, all lowercase).
//
// Example: isPalindrome("racecar") → true
//          isPalindrome("hello")   → false
func isPalindrome(s string) bool
```

### Extra work

```golang
// countVowels takes a string and returns the number of vowels (a, e, i, o, u)
// it contains.
//
// Example: countVowels("hello") → 2
//          countVowels("rhythm") → 0
func countVowels(s string) int

// removeSpaces takes a string and returns a version with all spaces removed,
// along with the number of spaces that were removed.
//
// Example: removeSpaces("the fox") → ("thefox", 2)
//          removeSpaces("hello")   → ("hello", 0)
func removeSpaces(s string) (string, int)

// replaceVowels takes a string and a replacement string, and returns a version
// of the input where every vowel has been swapped out for the replacement,
// along with the number of replacements made.
//
// Example: replaceVowels("hello", "*") → ("h*ll*", 2)
//          replaceVowels("rhythm", "*") → ("rhythm", 0)
func replaceVowels(s string, replacement string) (string, int)

// truncate takes a string and a maximum length. If the string is longer than
// max, it returns a version cut to max characters and true. If the string is
// already short enough, it returns the original string unchanged and false.
//
// Example: truncate("hello world", 5) → ("hello", true)
//          truncate("hi", 5)          → ("hi", false)
func truncate(s string, max int) (string, bool)

// capitaliseWords takes a string and returns a version where the first letter
// of every word is uppercased, along with the number of words that were
// processed.
//
// Example: capitaliseWords("the quick fox") → ("The Quick Fox", 3)
//          capitaliseWords("hello")         → ("Hello", 1)
func capitaliseWords(s string) (string, int)
```
