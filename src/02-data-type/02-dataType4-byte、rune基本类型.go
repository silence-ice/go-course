package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	/**
	 * 字符串转byte、rune
	 * byte 型(uint8 类型)，代表了 ASCII 码的一个字符
	 * rune 类型，代表一个 UTF-8 字符，当需要处理中文、日文或者其他复合字符时，则需要用到 rune 类型
	 */
	a := "aaa"
	fmt.Println(a)         //输出 aaa
	fmt.Println([]rune(a)) //a的ASCII码是97，输出[97 97 97]
	fmt.Println([]byte(a)) //a的ASCII码是97，输出[97 97 97]

	for i, ch := range []byte(a) {
		fmt.Println(i, string(ch))
		//(0 a)(1 a)(2 a)
	}

	b := "黄斌黄"
	fmt.Println(b)         //输出 黄斌黄
	fmt.Println([]byte(b)) //[233 187 132 230 150 140 233 187 132]，一个中文字符当成3个字符
	fmt.Println([]rune(b)) //[40644 25996 40644]

	for i, ch := range []rune(b) {
		fmt.Println(i, string(ch))
		//fmt.Printf("(%d %c)",i,ch)//(0 黄)(1 斌)(2 黄)
	}

	//输出字符长度
	fmt.Println(utf8.RuneCountInString(b)) //输出3

	//中文字中的 byte转rune
	bBytes := []byte(b)
	//fmt.Println(len(bBytes))//9
	ch, size := utf8.DecodeRune(bBytes)
	fmt.Println(ch, size) //40644 3，输出第一个中文字符，字符的长度占3个

	for len(bBytes) > 0 {
		ch, size := utf8.DecodeRune(bBytes)
		bBytes = bBytes[size:]
		//40644 黄
		//25996 斌
		//40644 黄
		fmt.Println(ch)
	}
}
