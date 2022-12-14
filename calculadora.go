package main // Esta línea indica que este programa se ejecutará como un programa independiente.

import "fmt" // Esta línea importa la biblioteca fmt, que contiene funciones útiles para imprimir datos.

func main() { // Esta línea define la función main, que es la función principal de este programa.

	fmt.Println("¡Bienvenido a la calculadora!") // Esta línea imprime un mensaje de bienvenida en la pantalla.

	var num1 float64 // Esta línea crea una variable llamada num1 de tipo float64.

	fmt.Print("Ingresa el primer número: ") // Esta línea imprime un mensaje en la pantalla para solicitar un número.

	fmt.Scan(&num1) // Esta línea lee la entrada del usuario y la guarda en la variable num1.

	var num2 float64 // Esta línea crea una variable llamada num2 de tipo float64.

	fmt.Print("Ingresa el segundo número: ") // Esta línea imprime un mensaje en la pantalla para solicitar otro número.

	fmt.Scan(&num2) // Esta línea lee la entrada del usuario y la guarda en la variable num2.

	fmt.Print("El resultado es: ") // Esta línea imprime un mensaje en la pantalla para indicar el resultado.

	fmt.Println(num1 + num2) // Esta línea imprime el resultado de la suma de los dos números.

} // Esta línea marca el fin de la función main.