package main

import (
	"errors"
	"strings"
	"unicode"
)

func Maiuscula(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("slice vazio")
	}

	var string []string

	for _, str := range slice {
		if len(str) > 0 && unicode.IsUpper([]rune(str)[0]) {
			string = append(string, str)
		}
	}

	if len(string) == 0 {
		return "", errors.New("Nenhuma string começa com letra maiúscula")
	}

	resultado := strings.Join(string, " ")
	return resultado, nil
}
