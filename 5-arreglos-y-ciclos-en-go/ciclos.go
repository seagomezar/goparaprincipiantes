package main

import "fmt"

func main() {

	// ejemploCicloFor()
	// ejemploCicloWhile()
	operacionesArreglos()
}

func ejemploCicloWhile() {
	// Declarando una variable para contar
	var i int = 0

	// Usando el ciclo for para iterar 5 veces como si fuera un while
	for i < 5 {
		fmt.Println("Contador:", i)
		i++
	}

	var j int = 0

	// Emulacion del ciclo do-while
	for {
		fmt.Println("Contador:", j)
		j++
		if j > 4 {
			break
		}
	}
}

func ejemploCicloFor() {
	var arr = [3]int{5, 10, 15}

	// Ambos bloques son similares i contiene el indice del arreglo
	// v contiene el valor del elemento del arreglo
	for i, v := range arr {
		fmt.Println(i, v)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	/* En este caso en ambos ciclos se imprimirá:
	0 5
	1 10
	2 15
	*/
}

func operacionesArreglos() {
		// Creamos un arreglo de enteros
	    numeros := []int{1, 2, 3}

	    // Agregamos un elemento al arreglo
	    numeros = append(numeros, 4)

	    fmt.Println(numeros) // [1, 2, 3, 4]

		// Eliminar el ultimo elemento del arreglo
		numeros = numeros[:len(numeros)-1]

		fmt.Println(numeros) // [1, 2, 3]

		// Eliminar el primer elemento del arreglo
		numeros = numeros[1:]
		fmt.Println(numeros) // [2, 3]

		// Eliminar el elemento en la posicion 1
		numeros = append(numeros[:1], numeros[2:]...)
		fmt.Println(numeros) // [2]

		// Eliminar el elemento en la posicion 0
		numeros = append(numeros[:0], numeros[1:]...)
		fmt.Println(numeros) // []

		// Volvamos a llenar el arreglo con 5 numeros
		numeros = []int{1, 6, 3, 4, 2}

		//Buscar el elemento 3 en el arreglo
		index := 0
		for index < len(numeros) {
			if numeros[index] == 3 {
				break
			}
			index++
		}

		fmt.Println(index) // 2

		// Ordenar el arreglo ascendentemente
		ordenarMayor(numeros, len(numeros))

		// Ordenar el arreglo descendentemente
		ordenarMenor(numeros, len(numeros))

}

// La función ordenarMayor() es la encargada de tomar el array y ordenar sus elementos de mayor a menor. Recibe dos argumentos, el primero es un array de tipo int y el segundo es la cantidad de elementos del array, que también es de tipo int.
func ordenarMayor(listNum []int, Cant int) {
    tmp := 0
    for x := 0; x < Cant; x++ {
        for y := 0; y < Cant; y++ {
            if listNum[x] > listNum[y] {
                tmp = listNum[x]
                listNum[x] = listNum[y]
                listNum[y] = tmp
            }
        }
    }
    fmt.Print("\nArray dinamico ordenado: ")
    for i := 0; i < Cant; i++ {
        fmt.Print("[",listNum[i],"]")
    }
    fmt.Println()
}

func ordenarMenor(listNum []int, Cant int) {
    tmp := 0
    for x := 0; x < Cant; x++ {
        for y := 0; y < Cant; y++ {
            if listNum[x] < listNum[y] {
                tmp = listNum[y]
                listNum[y] = listNum[x]
                listNum[x] = tmp
            }
        }
    }
    fmt.Print("\nArray dinamico ordenado: ")
    for i := 0; i < Cant; i++ {
        fmt.Print("[",listNum[i],"]")
    }
    fmt.Println()
}