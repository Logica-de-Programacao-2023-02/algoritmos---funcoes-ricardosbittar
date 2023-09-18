package main

import (
	"errors"
	"sort"
)

func alfabetica(x string) (string, error) {
	if x == "" {
		return "", errors.New("texto vazio")

	}
	sort.Strings(x)
	return x, nil

}
