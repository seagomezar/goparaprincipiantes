package main

import (
	"fmt"
	"math"
)

type Producto struct {
	Nombre string
	Precio float64
}

func main() {
	var numeros = []float64{1, 2, 3, 4, 5}
	fmt.Println("Promedio", promedio(numeros))
	fmt.Println("Producto", producto(2, 3))
	fmt.Println("Sumar", sumar(2.0, 3.0))
	fmt.Println("Factorial", factorial(5))
	fmt.Println("Área Circulo", areaCirculo(2))
	fmt.Println("Perímetro Cuadrado", perimetroCuadrado(2))
	fmt.Println("Volumen Esfera", volumenEsfera(2))
	fmt.Println("Area Triangulo", areaTriangulo(2, 3))
	fmt.Println("Hipotenusa", hipotenusa(2, 3))
	ingresoDatos()
}

func promedio(numeros []float64) float64 {
	var sum = 0.0
	for _, numero := range numeros {
		sum += numero
	}
	return sum / float64(len(numeros))
}

func producto(a float64, b float64) float64 {
	var sum = a * b
	return sum
}

func sumar(a float64, b float64) float64 {
	var sum = a * b
	return sum
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func areaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}

func perimetroCuadrado(lado float64) float64 {
	return lado * 4
}

func volumenEsfera(radio float64) float64 {
	return (4 / 3) * math.Pi * radio * radio * radio
}

func areaTriangulo(base float64, altura float64) float64 {
	return base * altura / 2
}

func hipotenusa(cateto1 float64, cateto2 float64) float64 {
	return math.Sqrt(cateto1*cateto1 + cateto2*cateto2)
}

func ingresoDatos() {
	fmt.Println("Introduzca dos números: ")
	var a, b float64
	fmt.Println("Introduzca el número 1: ")
	fmt.Scanln(&a)
	fmt.Println("Introduzca el número 2: ")
	fmt.Scanln(&b)
	fmt.Println("Los numeros que ingresaste son: ", a, b)
}
