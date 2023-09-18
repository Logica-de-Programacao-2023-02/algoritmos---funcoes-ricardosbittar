package main

import "errors"

func interseccao(x []int, y []int) ([]int, error) {

	var resultado []int
	if len(x) == 0 || len(y) == 0 {
		return nil, errors.New("Um dos slices esta vazio")
		for i := 0; i < len(x); i++ {
			for j := 0; j < len(y); j++ {
				if x[i] == y[j] {
					resultado = append(resultado, i)

				}
			}
			return resultado, nil
		}
	}
}
