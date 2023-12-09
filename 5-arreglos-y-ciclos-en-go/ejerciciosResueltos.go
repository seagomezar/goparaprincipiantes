package main

import ("fmt";"sort")

func main() {
	ejercicio1()
	ejercicio2()
	ejercicio3()
	ejercicio4()
	ejercicio5()
	ejercicio6()
	ejercicio7()
	ejercicio8()
	ejercicio9()
	ejercicio10()
}

// 1. Escriba un programa en Go para imprimir todos los elementos de un arreglo:
func ejercicio1() {
	// declarar arreglo
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// iterar por el arreglo
	for _, num := range arr {
		fmt.Println(num)
	}
}

// 2. Escriba un programa en Go para encontrar la suma de todos los números de un arreglo:
func ejercicio2() {
	// declarar arreglo
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// inicializar la suma en 0
	sum := 0

	// iterar por el arreglo
	for _, num := range arr {
		sum += num
	}

	fmt.Println("La suma de los elementos del arreglo es", sum)
}

// 3. Escriba un programa en Go para encontrar el máximo de un arreglo:
func ejercicio3() {
	// declarar arreglo
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// inicializar el máximo en 0
	max := 0

	// iterar por el arreglo
	for _, num := range arr {
		if num > max {
			max = num
		}
	}

	fmt.Println("El elemento máximo del arreglo es", max)
}

// 4. Escriba un programa en Go para encontrar el índice de un elemento en un arreglo (Usa otra manera de iterar por el arreglo):
func ejercicio4() {
	// declarar arreglo
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// declarar el elemento a buscar
	element := 4

	// inicializar el índice
	index := -1

	// iterar por el arreglo
	for i, num := range arr {
		if num == element {
			index = i
			break
		}
	}

	fmt.Println("El índice del elemento", element, "es", index)
}

// 5. Escriba un programa en Go para ordenar los elementos de un arreglo (usando la función sort.Ints<):
func ejercicio5(){
	// declarar arreglo
    arr := []int{3,6,2,1,4,9,7,8,5,10}

    // ordenar el arreglo
    sort.Ints(arr)

    fmt.Println("Arreglo ordenado:", arr)
}

// 6. Escriba un programa en Go para eliminar un elemento de un arreglo:
func ejercicio6(){
	// declarar arreglo
    arr := []int{1,2,3,4,5,6,7,8,9,10}

    // declarar el elemento a eliminar
    element := 5

    // inicializar el índice
    index := -1

    // iterar por el arreglo
    for i, num := range arr {
        if num == element {
            index = i
            break
        }
    }

    // eliminar elemento del arreglo
    if index != -1 {
        arr = append(arr[:index], arr[index+1:]...)
    }

    fmt.Println("Arreglo después de eliminar elemento:", arr)
}

//7. Escriba un programa en Go para contar el número de veces que se repite un elemento en un arreglo:

func ejercicio7() {

    // declarar arreglo
    arr := []int{1,2,3,1,4,1,5,6,7,8,9,10,1}

    // declarar el elemento a buscar
    element := 1

    // inicializar el contador
    count := 0

    // iterar por el arreglo
    for _, num := range arr {
        if num == element {
            count++
        }
    }

    fmt.Println("El elemento", element, "se repite", count, "veces.")
}

// 8. Escriba un programa en Go para insertar un elemento en una posición específica de un arreglo:
func ejercicio8() {

    // declarar arreglo
    arr := []int{1,2,3,4,5,6,7,8,9,10}

    // declarar el elemento y la posición
    element := 11
    position := 5

    // insertar el elemento
    arr = append(arr, 0)
    copy(arr[position+1:], arr[position:])
    arr[position] = element

    fmt.Println("Arreglo después de insertar elemento:", arr)
}

//9. Escriba un programa en Go para invertir los elementos de un arreglo:
func ejercicio9(){

    // declarar arreglo
    arr := []int{1,2,3,4,5,6,7,8,9,10}

    // invertir los elementos
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }

    fmt.Println("Arreglo invertido:", arr)
}

// 10. Escriba un programa en Go para eliminar los elementos duplicados de un arreglo:
func ejercicio10()  {
	// declarar arreglo
    arr := []int{1,2,3,4,4,5,6,7,7,8,9,9,10}

    // inicializar una nueva matriz
    newArr := []int{}

    // iterar por el arreglo
    for _, num := range arr {

        // verificar si el elemento ya existe en la nueva matriz
        isExist := false
        for _, newNum := range newArr {
            if num == newNum {
                isExist = true
                break
            }
        }

        // agregar el elemento si no existe
        if !isExist {
            newArr = append(newArr, num)
        }
    }

    fmt.Println("Arreglo sin elementos duplicados:", newArr)
}