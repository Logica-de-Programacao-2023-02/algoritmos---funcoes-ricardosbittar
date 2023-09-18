package main

import "errors"

func primos(x int) ([]int, error) {
	var primos []int

	if x < 0 {
		return 0, errors.New("Numero negativo")
	}
	for i := 0; i < x; i++ {

		if i >= 2 && i != x && x%i == 0 {
			primos = append(primos, i)
		}

	}
	return primos, nil

}