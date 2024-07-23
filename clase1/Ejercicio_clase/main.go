package main

import "fmt"

func main() {
	var palabra string = "palabra" // una manera de declarar la variable

	//palabra = "palabra2" // otra manera de declarar una variable

	cant := len(palabra)

	fmt.Println("la palabra ", palabra, " tiene", cant, " letras.")
	fmt.Printf("La palabra %s tiene %d letras.", palabra, cant)

}
