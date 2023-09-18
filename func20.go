package main

import "errors"

func cincoCaracteres(slice []string) ([]string, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	var resultStrings []string

	for _, str := range slice {
		if len(str) > 5 {
			resultStrings = append(resultStrings, str)
		}
	}

	if len(resultStrings) == 0 {
		return nil, errors.New("Nenhuma string possui mais de 5 caracteres")
	}

	return resultStrings, nil
}
