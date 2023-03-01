package main

import "fmt"

func main() {
	lengthOfNonRepeatingSubStr("abcabcbb")
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	//fmt.Println(lengthOfNonRepeatingSubStr("aaaa"))
	//fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	//fmt.Println(lengthOfNonRepeatingSubStr(""))
	//fmt.Println(lengthOfNonRepeatingSubStr("这里是慕课网"))
}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	return maxLength
}
