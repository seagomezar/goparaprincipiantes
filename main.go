package main

import "fmt"

func main() {
	miPersona := Persona{nombre: "Juan", edad: 25}
	var otraPersona Persona = Persona{nombre: "Juan", edad: 25}
	fmt.Println("------------------------")
	fmt.Println("Nombre:", miPersona.nombre)
	fmt.Println("Edad:", miPersona.edad)

	fmt.Println("------------------------")
	fmt.Println("Nombre:", otraPersona.nombre)
	fmt.Println("Edad:", otraPersona.edad)

	miTractor := Truck{Car{brand: "Ford", year: 2022}, 2}
	fmt.Println("------------------------")
	fmt.Println("Tractor:", miTractor.Car.brand)
}

// Define un nuevo tipo llamado 'persona'.
type Persona struct {
	// Declara un campo 'nombre' de tipo string dentro del struct 'persona'.
	nombre string
	// Declara un campo 'edad' de tipo int (número entero) dentro del struct 'persona'.
	edad int
}

type Car struct {
	brand string
	year  int
}

type Truck struct {
	Car
	capacity int
}

type city struct {
	name string
	state string
}

type country struct {
	city
	continent string
}

type Estudiante struct {
	nombre string
	edad int
	curso string
	notas [3]float64
}




