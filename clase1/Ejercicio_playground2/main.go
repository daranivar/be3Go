package main

import "fmt"

//  Ejercicio 2 - Préstamo
// Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos. El banco tiene ciertas
// reglas para saber a qué cliente se le puede otorgar: solo le otorga préstamos a clientes cuya edad sea mayor a 22 años,
// se encuentren empleados y tengan más de un año de antigüedad en su trabajo. Dentro de los préstamos que otorga, no les
// cobrará interés a los que su sueldo sea mejor a $100.000.
//  Es necesario realizar una aplicación que tenga estas variables y que imprima un mensaje de acuerdo a cada caso.
// Pista: Tu código tiene que poder imprimir al menos tres mensajes diferentes.

func main() {
	// Definir las variables
	edad := 25
	empleado := true
	antiguedad := 2
	sueldo := 95000.0

	// Evaluar las condiciones para otorgar el préstamo
	if edad > 22 && empleado && antiguedad > 1 {
		if sueldo > 100000 {
			fmt.Println("Préstamo otorgado sin interés.")
		} else {
			fmt.Println("Préstamo otorgado con interés.")
		}
	} else {
		fmt.Println("No cumple con los requisitos para obtener un préstamo.")
	}

}
