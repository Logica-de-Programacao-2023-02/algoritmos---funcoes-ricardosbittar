package main

import "fmt"

func SegMaior(slice []int) int {

	if len(slice) < 2 {
		fmt.Println("O slice deve ter ao menos dois elementos")

		return 0

	}
	maior := slice[0]
	segundoMaior := slice[0]

	for _, i := range slice {
		if i > maior {
			maior = i
			segundoMaior = maior
		} else if i > segundoMaior && i != maior {
			segundoMaior = i
		}
	}
	return segundoMaior
}
