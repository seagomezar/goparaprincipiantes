
Arreglos en Go

Los arreglos en Go son una estructura de datos que se usa para almacenar un conjunto de valores del mismo tipo. Un arreglo es una secuencia ordenada de elementos que comparten el mismo tipo. Los arreglos tienen un tamaño fijo, por lo que una vez creado, no se puede cambiar.

Por ejemplo, el siguiente código crea un arreglo de enteros de tamaño 3:

``` go
var arr [3]int
```

Los elementos de un arreglo se pueden acceder usando su índice, que es el número de la posición del elemento en el arreglo. Los índices comienzan desde 0. Por ejemplo, el siguiente código asigna el valor 5 al primer elemento del arreglo anterior:

``` go
arr[0] = 5
```

Los arreglos también se pueden inicializar con valores predefinidos. Por ejemplo, el siguiente código crea un arreglo de enteros de tamaño 3 con los valores 5, 10 y 15:

``` go
var arr = [3]int{5, 10, 15}
```

Los arreglos se pueden usar para iterar sobre sus elementos usando un bucle for. Por ejemplo, el siguiente código imprime los elementos del arreglo anterior:

``` go
for i := 0; i < len(arr); i++ {
  fmt.Println(arr[i])
}

// Esto imprimirá 5, 10 y 15 en la pantalla.
```

Como viste el codigo anterior utiliza un for que es una estructura que no hemos revisado aun, hablemos de los ciclos en go:

1. Ciclo for: El ciclo for es el ciclo más comúnmente usado en Go. Se usa para iterar sobre una secuencia de elementos. Un ejemplo de su uso sería:

La sintaxis de for en Go es la siguiente:

for [inicialización]; [condición]; [incremento] {
    // código a ejecutar
}

Estructura del ciclo for

+---------+
|  Start  |
+---------+
     |
	 V
+---------+
|  i = 1  |
+---------+
     |
	 V
	+----------+
----|  i < 10  |<----------------
|	+----------+				|
|		| True					|
|		V						|
|	+----------------------+	|
|	|  Program Statements  |-----
|	+----------------------+
|False
|
V
+---------+
|  End	  |
+---------+


Este es el enlace a la documentacion oficial de Go que habla sobre el ciclo For: https://go.dev/ref/spec#For_statements

por ejemplo:

```go
package main

import "fmt"

func main() {
   // Declarando un arreglo de enteros
   array := []int{10, 20, 30, 40, 50, 60}

   // Usando el ciclo for para iterar sobre los elementos del arreglo
   for i := 0; i < len(array); i++ {
      fmt.Println("Elemento en la posición", i, "es", array[i])
   }
}
```

El código anterior imprimirá lo siguiente:

Elemento en la posición 0 es 10
Elemento en la posición 1 es 20
Elemento en la posición 2 es 30
Elemento en la posición 3 es 40
Elemento en la posición 4 es 50
Elemento en la posición 5 es 60

2. Ciclo while: El ciclo while es una estructura de control que se usa para ejecutar un bloque de código repetidamente mientras una condición se cumpla. Su sintaxis es


while [condición] {
    // código a ejecutar
}

Estructura del ciclo while
+---------+
|  Start  |
+---------+
     |
	   V
	  +----------+
----| condition|<----------------
|	  +----------+				         |
|		| True					            |
|		V						                |
|	+----------------------+	    |
|	|  Program Statements  |-----
|	+----------------------+
|False
|
V
+---------+
|  End	  |
+---------+


Un ejemplo de su uso sería:

```go
package main

import "fmt"

func main() {
   // Declarando una variable para contar
   var i int = 0

   // Usando el ciclo while para iterar 5 veces
   for i < 5 {
      fmt.Println("Contador:", i)
      i++
   }
}
```

El código anterior imprimirá lo siguiente:

Contador: 0
Contador: 1
Contador: 2
Contador: 3
Contador: 4

3. Ciclo do-while: El ciclo do-while es una estructura de control que se usa para ejecutar un bloque de código repetidamente hasta que una condición se cumpla. Su sintaxis es:

sintaxis de do while

do {
    // código a ejecutar
} while [condición];

Un ejemplo de su uso sería:

```go
package main

import "fmt"

func main() {
   // Declarando una variable para contar
   var i int = 0

   // Usando el ciclo do-while para iterar 5 veces
   do {
      fmt.Println("Contador:", i)
      i++
   } while (i < 5)
}
```

El código anterior imprimirá lo siguiente:

Contador: 0
Contador: 1
Contador: 2
Contador: 3
Contador: 4

Operaciones sobre arreglos en Go

Los arreglos en Go se pueden usar para almacenar una colección de elementos del mismo tipo. Los arreglos se utilizan para almacenar datos de forma estructurada y pueden realizar varias operaciones.

Las operaciones más comunes que se pueden realizar en arreglos en Go son:

1. Agregar elementos: Los elementos se pueden agregar a un arreglo usando la función append ().

2. Eliminar elementos: Los elementos se pueden eliminar de un arreglo usando la función delete ().

3. Buscar elementos: Los elementos se pueden buscar en un arreglo usando la función find ().

4. Ordenar elementos: Los elementos se pueden ordenar en un arreglo usando la función sort ().

5. Filtrar elementos: Los elementos se pueden filtrar en un arreglo usando la función filter ().

6. Transformar elementos: Los elementos se pueden transformar en un arreglo usando la función map ().

7. Mezclar elementos: Los elementos se pueden mezclar en un arreglo usando la función zip ().

Veamos ejemplos detallados y correctos para cada una de las operaciones descritas anteriormente

1. Agregar elementos:

``` go
package main

import "fmt"

func main() {
    // Creamos un arreglo de enteros
    numeros := []int{1, 2, 3}

    // Agregamos un elemento al arreglo
    numeros = append(numeros, 4)

    fmt.Println(numeros) // [1, 2, 3, 4]
}
```

2. Eliminar elementos:

``` go
package main

import "fmt"

func main() {
    // Creamos un arreglo de enteros
    numeros := []int{1, 2, 3, 4}

    // Eliminamos el último elemento del arreglo
    numeros = delete(numeros, len(numeros)-1)

    fmt.Println(numeros) // [1, 2, 3]
}
```

3. Buscar elementos:

``` go
package main

import "fmt"

func main() {
    // Creamos un arreglo de enteros
    numeros := []int{1, 2, 3, 4}

    // Buscamos el elemento con valor 3
    index := find(numeros, 3)

    fmt.Println(index) // 2
}
```

4. Ordenar elementos:

``` go
package main

import "fmt"

func main() {
    // Creamos un arreglo de enteros
    numeros := []int{3, 1, 4, 2}

    // Ordenamos el arreglo
    sort(numeros)

    fmt.Println(numeros) // [1, 2, 3, 4]
}
```

5. Filtrar elementos:

``` go
package main

import "fmt"

func main() {
    // Creamos un arreglo de enteros
    numeros := []int{1, 2, 3, 4, 5}

    // Filtrar los elementos pares
    numerosPares := filter(numeros, func(n int) bool {
        return n%2 == 0
    })

    fmt.Println(numerosPares) // [2, 4]
}
```

6. Transformar elementos:

``` go
package main

import "fmt"

func main() {
    // Creamos un arreglo de enteros
    numeros := []int{1, 2, 3, 4}

    // Transformamos los elementos del arreglo
    numerosDobles := map(numeros, func(n int) int {
        return n * 2
    })

    fmt.Println(numerosDobles) // [2, 4, 6, 8]
}
```

7. Mezclar elementos:

``` go
package main

import "fmt"

func main() {
    // Creamos dos arreglos de enteros
    numeros := []int{1, 2, 3}
    letras := []string{"a", "b", "c"}

    // Mezclamos los dos arreglos
    mezcla := zip(numeros, letras)

    fmt.Println(mezcla) // [(1, "a"), (2, "b"), (3, "c")]
}
```

Vemos 10 ejercicios clasicos con arreglos en go y su respectiva solucion:

A continuación se presentan 10 ejercicios clásicos con arreglos en Go y su respectiva solución detalladamente:

1. Escriba un programa en Go para imprimir todos los elementos de un arreglo:

Solución:

``` go
package main

import "fmt"

func main() {

    // declarar arreglo
    arr := []int{1,2,3,4,5,6,7,8,9,10}

    // iterar por el arreglo
    for _, num := range arr {
        fmt.Println(num)
    }
}
```

2. Escriba un programa en Go para encontrar la suma de todos los números de un arreglo:

Solución:

``` go
package main

import "fmt"

func main() {

    // declarar arreglo
    arr := []int{1,2,3,4,5,6,7,8,9,10}

    // inicializar la suma en 0
    sum := 0

    // iterar por el arreglo
    for _, num := range arr {
        sum += num
    }

    fmt.Println("La suma de los elementos del arreglo es", sum)
}
```

3. Escriba un programa en Go para encontrar el máximo de un arreglo:

Solución:

``` go
package main

import "fmt"

func main() {

    // declarar arreglo
    arr := []int{1,2,3,4,5,6,7,8,9,10}

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
```

4. Escriba un programa en Go para encontrar el índice de un elemento en un arreglo:

Solución:

``` go
package main

import "fmt"

func main() {

    // declarar arreglo
    arr := []int{1,2,3,4,5,6,7,8,9,10}

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
```

5. Escriba un programa en Go para ordenar los elementos de un arreglo:

Solución:

``` go
package main

import (
	"fmt"
	"sort"
)

func main() {

    // declarar arreglo
    arr := []int{3,6,2,1,4,9,7,8,5,10}

    // ordenar el arreglo
    sort.Ints(arr)

    fmt.Println("Arreglo ordenado:", arr)
}
```

6. Escriba un programa en Go para eliminar un elemento de un arreglo:

Solución:

``` go
package main

import "fmt"

func main() {

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
```

7. Escriba un programa en Go para contar el número de veces que se repite un elemento en un arreglo:

Solución:

``` go
package main

import "fmt"

func main() {

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
```

8. Escriba un programa en Go para insertar un elemento en una posición específica de un arreglo:

Solución:

``` go
package main

import "fmt"

func main() {

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
```

9. Escriba un programa en Go para invertir los elementos de un arreglo:

Solución:

``` go
package main

import "fmt"

func main() {

    // declarar arreglo
    arr := []int{1,2,3,4,5,6,7,8,9,10}

    // invertir los elementos
    for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
        arr[i], arr[j] = arr[j], arr[i]
    }

    fmt.Println("Arreglo invertido:", arr)
}
```

10. Escriba un programa en Go para eliminar los elementos duplicados de un arreglo:

Solución:

``` go
package main

import "fmt"

func main() {

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
```

A continuacion te dejo 10 ejercicios mas para que practiques los ciclos en go y envies la solucion como un Pull requeste a este repositorio:

1. Escriba un programa en Go que imprima los primeros 10 números naturales.
2. Escriba un programa en Go que calcule la suma de los primeros 10 números naturales.
3. Escriba un programa en Go que calcule el promedio de los primeros 10 números naturales.
4. Escriba un programa en Go que calcule el factorial de un número dado.
5. Escriba un programa en Go que determine si un número dado es par o impar.
6. Escriba un programa en Go que determine si un número dado es primo.
7. Escriba un programa en Go que imprima los números del 1 al 100.
8. Escriba un programa en Go que determine la suma de los números del 1 al 100.
9. Escriba un programa en Go que determine si un número dado está dentro de un rango dado.
10. Escriba un programa en Go que calcule el número de dígitos de un número dado.