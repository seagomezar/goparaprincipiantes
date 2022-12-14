Variables

Las variables en Go son contenedores para almacenar datos. Una variable contiene un valor, que puede ser cualquier cosa, desde un número hasta una cadena de texto o un objeto. Cuando se define una variable, se especifica el nombre de la variable, el tipo de datos al que pertenece y el valor que se le asigna. Una vez que se ha definido una variable, se puede usar su nombre para referirse a ella y para acceder a su valor.

Existen dos tipos principales de variables en Go: las variables locales y las variables globales. Las variables locales se definen dentro de una función y solo están disponibles dentro de esa función. Las variables globales se definen fuera de cualquier función y están disponibles en todo el programa.

Variables locales:

Las variables locales son aquellas variables definidas dentro de una función o bloque de código. Por ejemplo, en el siguiente código, la variable "a" es una variable local:

``` go
func main() {
   a := 5
   // código usando variable 'a'
}
```

Variables globales:

Las variables globales son variables definidas fuera de cualquier función o bloque de código. Por ejemplo, en el siguiente código, la variable "b" es una variable global:

``` go
b := 10

func main() {
   // código usando variable 'b'
}
```

Las variables también se pueden definir como constantes. Las constantes son variables que tienen un valor fijo y no pueden cambiarse durante la ejecución del programa.

Constantes:

Las constantes son valores que un programa no puede cambiar durante su ejecución. Por ejemplo, en el siguiente código, la constante "PI" es una constante:
``` go
const PI = 3.14159
```

Variables:

Las variables son contenedores para almacenar valores durante la ejecución de un programa. Por ejemplo, en el siguiente código, la variable "x" es una variable:
``` go
var x int
x = 5
```

Las variables también se pueden definir como punteros. Los punteros son variables que contienen la dirección de memoria de otra variable. Esto permite a los programadores acceder o modificar el valor de una variable directamente, sin tener que pasar por la función.


Los punteros son variables que almacenan direcciones de memoria. Por ejemplo, en el siguiente código, la variable "p" es un puntero:
``` go
var x int
x = 5

var p *int
p = &x

// el valor de p es la dirección de memoria de x
```

Tipos de datos en Detalle

Los tipos de datos en Go determinan el tipo de valores que se pueden almacenar y manipular dentro de un programa. Go tiene una variedad de tipos de datos primitivos, incluidos enteros, flotantes, cadenas y booleanos, así como también tipos de datos compuestos como arreglos, estructuras y punteros. También puedes definir tus propios tipos de datos usando la palabra clave type.

Los enteros se usan para almacenar valores numéricos enteros, como 3, -2 y 0. Hay varios tipos de enteros, incluidos int8, int16, int32, int64, uint8, uint16, uint32 y uint64.

Los valores de punto flotante se usan para almacenar valores numéricos con decimales, como 3.14, -2.5 y 0.0. Hay dos tipos de punto flotante: float32 y float64.

Las cadenas se usan para almacenar secuencias de caracteres, como "hola mundo". Una cadena es una secuencia de caracteres codificada en UTF-8.

Los booleanos se usan para almacenar valores lógicos, como true o false.

Los arreglos se usan para almacenar un conjunto de valores del mismo tipo, como [1, 2, 3, 4]. Los arreglos tienen un tamaño fijo, por lo que una vez creado, no se puede cambiar.

Las estructuras se usan para almacenar un conjunto de variables relacionadas, como una dirección postal que contiene una calle, una ciudad, un estado, etc. Las estructuras son similares a los arreglos, pero pueden contener diferentes tipos de datos.

Las estructuras son tipos de datos compuestos que pueden contener varios campos. Por ejemplo, en el siguiente código, la estructura "Persona" contiene los campos "nombre", "edad" y "dirección":

``` go
type Persona struct {
   Nombre string
   Edad int
   Direccion string
}
```

Veamos diez ejemplos con su respectiva solucion:

1. Calcular la suma de dos números:
Solución:

``` go
package main

import "fmt"

func main() {
	a := 5
	b := 10
	sum := a + b
	fmt.Println("La suma de", a, "y", b, "es", sum)
}
```

2. Calcular el producto de dos números:
Solución:

``` go
package main

import "fmt"

func main() {
	a := 5
	b := 10
	product := a * b
	fmt.Println("El producto de", a, "y", b, "es", product)
}
```

3. Calcular el promedio de una lista de números:
Solución:

``` go
package main

import "fmt"

func main() {
	nums := []float64{10, 20, 30, 40}
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	avg := sum / float64(len(nums))
	fmt.Println("El promedio de la lista de números es", avg)
}
```

4. Calcular el factorial de un número:
Solución:

``` go
package main

import "fmt"

func main() {
	num := 5
	factorial := 1
	for i := num; i > 0; i-- {
		factorial *= i
	}
	fmt.Println("El factorial de", num, "es", factorial)
}
```

5. Calcular el área de un círculo:
Solución:

``` go
package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 10.0
	area := math.Pi * math.Pow(radius, 2)
	fmt.Println("El área del círculo es", area)
}
```

6. Calcular el perímetro de un cuadrado:
Solución:

``` go
package main

import "fmt"

func main() {
	side := 10.0
	perimeter := side * 4
	fmt.Println("El perímetro del cuadrado es", perimeter)
}
``

7. Verificar si un número es par o impar:
Solución:

``` go
package main

import "fmt"

func main() {
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "es un número par")
	} else {
		fmt.Println(num, "es un número impar")
	}
}
```

8. Calcular el volumen de una esfera:
Solución:

``` go
package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 10.0
	volume := (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
	fmt.Println("El volumen de la esfera es", volume)
}
```

9. Calcular el área de un triángulo:
Solución:

``` go
package main

import "fmt"

func main() {
	base := 10.0
	height := 5.0
	area := 0.5 * base * height
	fmt.Println("El área del triángulo es", area)
}
```

10. Calcular la hipotenusa de un triángulo rectángulo:
Solución:


``` go
package main

import (
	"fmt"
	"math"
)

func main() {
	side1 := 5.0
	side2 := 10.0
	hypotenuse := math.Sqrt(math.Pow(side1, 2) + math.Pow(side2, 2))
	fmt.Println("La hipotenusa del triángulo rectángulo es", hypotenuse)
}
```


11. Pedir al usuario dos numeros y visualizar el resultado por pantalla

``` go
// Esta es una calculadora simple en Go.

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
```

A continuacion te dejo 20 ejercicios mas para que practiques y envies la solucion como un Pull requeste a este repo:

1. Imprimir la longitud de una cadena de texto.
2. Comparar dos números enteros y mostrar el mayor.
3. Calcular el cuadrado de un número.
4. Crear una función para convertir una cadena a mayúsculas.
5. Crear una función para convertir una cadena a minúsculas.
6. Determinar si una cadena contiene una subcadena determinada.
7. Determinar si un número es par o impar.
8. Determinar si una cadena es palíndroma.
9. Incrementar un número en una unidad.
10. Determinar si un número está dentro de un rango determinado.
11. Obtener el índice de la primera aparición de una subcadena en una cadena.
12. Obtener la suma de dos números enteros.
13. Comprobar si un número es divisible por otro.
14. Determinar si una cadena es un número.
15. Calcular el máximo común divisor (MCD) de dos números.
16. Determinar si una cadena comienza con una subcadena determinada.
17. Obtener el tamaño de una matriz.
18. Obtener el número de elementos únicos en una matriz.
19. Determinar si un número es primo.
20. Obtener los índices de los elementos duplicados en una matriz.

