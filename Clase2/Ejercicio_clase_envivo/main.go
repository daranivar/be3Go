package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	minimum = "minimun"
	average = "average"
	maximum = "maximum"
)

func main() {
	minFun, err := operacion(minimum)
	if err != nil {
		log.Fatal(err)
	}

	maxFun, err := operacion(maximum)
	if err != nil {
		log.Fatal(err)
	}

	avgFun, err := operacion(average)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("El minimo es: ", minFun(3, 5, 1, 8, 2, 7, 4))
	fmt.Println("El maximo es: ", maxFun(3, 5, 1, 8, 2, 7, 4))
	fmt.Println("El average es: ", avgFun(3, 5, 1, 8, 2, 7, 4))

	min, max, avg, err := operacionMulti()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("minimo %d | maximo %d | promedio %d",
		min(3, 5, 1, 8, 2, 7, 4),
		max(3, 5, 1, 8, 2, 7, 4),
		avg(3, 5, 1, 8, 2, 7, 4))

	errorFunc, err := operacion("errorFunc")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("El error es: ", errorFunc(3, 5, 1, 8, 2, 7, 4))

}

func calculoMinimo(values ...int) int {
	min := values[0]
	for _, val := range values {
		if val < min {
			min = val
		}
	}
	return min
}

func calculoMaximo(values ...int) int {
	max := values[0]
	for _, val := range values {
		if val > max {
			max = val
		}
	}
	return max
}

func calculoAverage(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum / len(values)
}

func operacion(calculo string) (func(...int) int, error) {
	switch calculo {
	case minimum:
		return calculoMinimo, nil
	case maximum:
		return calculoMaximo, nil
	case average:
		return calculoAverage, nil
	default:
		return nil, errors.New("El calculo no existe.")
	}
}

func operacionMulti() (func(...int) int, func(...int) int, func(...int) int, error) {
	return calculoMinimo, calculoMaximo, calculoAverage, nil
}
