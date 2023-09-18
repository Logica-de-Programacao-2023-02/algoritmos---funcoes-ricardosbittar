package main

import (
	"errors"
)

func crescente(slice []int) ([]int, error) {

	if len(slice) == 0 {
		return nil, errors.New("Slice vazio")
	}
	for i := 0; i < len(slice)-1; i++ {
		if slice[i] < slice[i+1] {
			continue
		} else if slice[i] > slice[i+1] {
			slice[i], slice[i+1] = slice[i+1], slice[i]

		}
		return slice, nil
	}

}
