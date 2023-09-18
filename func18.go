package main

import "errors"

func Somar(slice []int, f func(int) int) (int, error) {
	if len(slice) == 0 {
		return 0, errors.New("O slice está vazio")
	}

	var soma int

	for _, valor := range slice {
		resultado := f(valor)
		soma += resultado
	}

	return soma, nil
}
