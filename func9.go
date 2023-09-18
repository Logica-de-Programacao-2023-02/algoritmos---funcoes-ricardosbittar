package main

import (
	"errors"
	"strings"
)

func extrairPalavras(texto string) ([]string, error) {
	if texto == "" {
		return nil, errors.New("A string está vazia")
	}

	palavras := strings.Fields(texto)

	return palavras, nil
}