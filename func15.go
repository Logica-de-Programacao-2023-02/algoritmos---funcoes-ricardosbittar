package main

import "errors"

func contem(x int, slice []int) (bool, error) {

	if len(slice) == 0 {
		return false, errors.New("Slice vazia")
	}

	for i := 0; i < len(slice); i++ {
		if i == x {
			return true, nil
		}

	}
	return false, nil
}