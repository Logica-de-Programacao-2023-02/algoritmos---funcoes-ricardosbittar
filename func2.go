package main

import "strings"

func qtdVogais(x string) map[rune]int {
	vogais := "aeiouAEIOU"
	m := make(map[rune]int)

	for _, i := range x {
		if strings.ContainsRune(vogais, i) {
			m[i]++
		}
		return m
	}
}
