package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 5
	fmt.Println(arr)
	// Esto imprime [5 0 0]
	// 🚫 Esto no es permitido porque debe ser un arreglo de tamaño variable
	// arr = append(arr, 3)

	// Un arreglo de tamaño variable
	var arr2 [] string;
	// Asi podemos agregar elementos al arreglo
	arr2 = append(arr2, "hola")
	fmt.Println(arr2)
	// ☢️ Si el indice 88 aun no se ha llenado la siguiente línea sacara error
	// arr2[88] = "helloworld"
	fmt.Println(arr2)
}