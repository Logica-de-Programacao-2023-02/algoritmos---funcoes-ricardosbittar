package main

import (
	"errors"
	"fmt"
)

func duplicar(x int) int {
	return x * 2
}

func mapearSlice(slice []int, f func(int) int) ([]int, error) {
	if len(slice) == 0 {
		return nil, errors.New("O slice está vazio")
	}

	resultado := make([]int, len(slice))

	for i, valor := range slice {
		resultado[i] = f(valor)
	}

	return resultado, nil
}

func main() {
	slice := []int{1, 2, 3, 4, 5}

	resultado, err := mapearSlice(slice, duplicar)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println(resultado)
	}

	sliceVazio := []int{}
	resultadoVazio, errVazio := mapearSlice(sliceVazio, duplicar)
	if errVazio != nil {
		fmt.Println("Erro:", errVazio)
	}
}