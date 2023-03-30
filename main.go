package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "selamat malam"
	words := strings.Split(str, " ")

	for _, word := range words {
		for _, char := range word {
			fmt.Printf("%c\n", char)
		}
	}

	wordCount := make(map[string]int)
	for _, word := range words {
		for _, char := range word {
			charString := string(char)
			if _, ok := wordCount[charString]; ok {
				wordCount[charString]++
			} else {
				wordCount[charString] = 1
			}
		}
	}

	fmt.Println(wordCount)
}
