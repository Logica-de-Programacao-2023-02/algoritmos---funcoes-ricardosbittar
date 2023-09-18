package main

import "errors"

func pares(slice []int) ([]int, error) {
	var novoslice []int
	if len(slice) < 1 {
		return nil, errors.New("Slice Vazio")

	}
	for i := 0; i < len(slice); i++ {
		if i%2 == 0 {
			novoslice = append(novoslice, i)

		}
		return novoslice, nil
	}
}