package main

import "errors"

func somaDigito(numero int) (int, error) {
	if numero < 0 {
		return 0, errors.New("Número negativo")
	}

	var soma int

	for numero > 0 {
		digito := numero % 10
		soma += digito
		numero /= 10
	}

	return soma, nil
}
