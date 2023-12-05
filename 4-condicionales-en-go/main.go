package main

import "fmt"

func main() {
	ejemploIfElse()
	ejemploIfElseIf()
	ejemploSwitch()
	ejemploSwitchRango()
	ejemploSwitchConTexto()
}

func ejemploIfElse() {
	var x int = 11
	if x > 10 {
		fmt.Println("x es mayor a 10")
	} else {
		fmt.Println("x es menor a 10")
	}
}

func ejemploIfElseIf() {
	var x int = 11
	if x > 10 {
		fmt.Println("x es mayor a 10")
	} else if x == 10 {
		fmt.Println("x es igual a 10")
	} else {
		fmt.Println("x es menor a 10")
	}
}

func ejemploSwitch() {
	var x int = 11
	switch x {
	case 1:
		fmt.Println("x es 1")
	case 2:
		fmt.Println("x es 2")
	case 3:
		fmt.Println("x es 3")
	case 4:
		fmt.Println("x es 4")
	case 5:
		fmt.Println("x es 5")
	case 6:
		fmt.Println("x es 6")
	case 7:
		fmt.Println("x es 7")
	case 8:
		fmt.Println("x es 8")
	case 9:
		fmt.Println("x es 9")
	case 10:
		fmt.Println("x es 10")
	default:
		fmt.Println("x es mayor a 10")
	}
}

func ejemploSwitchRango() {
	var numeroMes int = 3
	var estacion string

	switch {
	case numeroMes >= 1 && numeroMes < 3:
		estacion = "Invierno"
	case numeroMes >= 3 && numeroMes < 6:
		estacion = "Primavera"
	case numeroMes >= 6 && numeroMes < 9:
		estacion = "Verano"
	case numeroMes >= 9 && numeroMes < 12:
		estacion = "Otoño"
	default:
		estacion = "Ese valor no corresponde a ninguna Estación del Año"
	}
	fmt.Println(estacion)
}

func ejemploSwitchConTexto() {
	var x string = "def"
	// Evalúa el valor de la variable 'x' en una estructura 'switch'
	switch x {
	// Si 'x' es igual a "abc"
	case "abc":
		// Imprime en la consola el mensaje "x es abc"
		fmt.Println("x es abc")
	case "def":
		// Imprime en la consola el mensaje "x es def"
		fmt.Println("x es def")
	// Si 'x' no es igual a "abc" ni a "def"
	default:
		// Imprime en la consola el mensaje "x es otra cadena"
		fmt.Println("x es otra cadena")
	}
}
