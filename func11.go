package main

import (
	"errors"
	"fmt"
)

func primo(x int) (bool, error) {
	primo := true
	if x < 0 {
		return false, errors.New("Numero negativo")
	}

	for i := 0; i < x; i++ {
		if x%i == 0 {
			primo = false
		}

	}
	if primo {
		fmt.Println("é primo")
	} else {
		fmt.Println("Nao é primo")
	}
	return true, nil
}
