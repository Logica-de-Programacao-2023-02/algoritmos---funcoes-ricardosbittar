package main

import (
	"errors"
	"strings"
)

func concatenar(slice []string) (string, error) {
	if len(slice) == 0 {
		return "", errors.New("O slice está vazio")
	}
	resultado := strings.Join(slice, ",")
	return resultado, nil
}
