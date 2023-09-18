package main

import "strings"

func conc(slice []string) string {
	concatenacao := strings.Join(slice, "")
	return concatenacao
}
