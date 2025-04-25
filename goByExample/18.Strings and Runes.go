package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	fmt.Println(len(s))
	// 打印出来 "18" 表示"len()"返还字符串的字节大小

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()
	// fmt.Println()打印了一个空行没有实际含义

	fmt.Println(utf8.RuneCountInString(s))
	// 用"utf8.RuneCountInString(s)"打印出来了字符的长度,6个,还有这个泰语真难,看半天才知道6个字符是哪六个字符

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	fmt.Println("\n")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {

	if r == 't' {
		fmt.Println("Found tee")
	} else if r == 'ส' {
		fmt.Println("Found so sua")

	}
}
