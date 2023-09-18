package main

import (
	"errors"
	"strings"
)

func asterisco(x string) (string, error) {
	if x == "" {
		return "", errors.New("Texto vazio")
	}
	vogais := "aeiouAEIOU"
	for _, vogal := range vogais {
		x = strings.ReplaceAll(x, string(vogal), "*")
	}
	return x, nil

}
