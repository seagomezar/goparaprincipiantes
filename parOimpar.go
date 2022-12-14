package main

import "fmt"

func main() {

    // Solicitar un número al usuario
    var numero int
    fmt.Println("Por favor ingresa un número:")
    fmt.Scan(&numero)

    // Calcular el resto de la división del número entre 2
    resto := numero % 2

    // Si el resto es 0
    if resto == 0 {
        // Imprimir "El número es par"
        fmt.Println("El número es par")
    } else {
        // Imprimir "El número es impar"
        fmt.Println("El número es impar")
    }
}