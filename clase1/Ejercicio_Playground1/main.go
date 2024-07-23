package main

//Ejercicio 1 - Descuento
/*Una tienda de ropa quiere ofrecer a sus clientes un descuento sobre sus productos. Para ello necesitan una aplicación que
les permita calcular el descuento basándose en dos variables: su precio y el descuento en porcentaje. La tienda espera obtener
como resultado el valor con el descuento aplicado y luego imprimirlo en consola.
Crear la aplicación de acuerdo a los requerimientos.*/
import "fmt"

var precio int = 20
var descuento int = 10
var resultado int = 0

func main() {
	resultado = precio * descuento / 100
	total := precio - descuento
	fmt.Println("el descuento es: ", resultado)
	fmt.Println("el precio total es: ", total)

}
