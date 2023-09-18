package main

func media(slice []int) float64 {
	var soma int
	var media float64
	for i := 0; i < len(slice); i++ {
		soma += i

	}
	media = float64(soma) / float64(len(slice))
	return media
}
