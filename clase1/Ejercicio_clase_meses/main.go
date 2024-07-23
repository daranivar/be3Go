package main

import "fmt"

func main() {
	var meses = []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Setiembre", "Octubre", "Noviembre", "Diciembre"}

	mes := 3

	if mes < 1 && mes > 12 {
		fmt.Println("Error, fuera de mes")
	}

	if mes >= 1 && mes <= 12 {
		for i := 0; i <= len(meses); i++ {
			if i == mes {
				fmt.Println(meses[i-1])

			}
		}
	}
}
