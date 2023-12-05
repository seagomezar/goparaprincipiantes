
### 📦Variables

Las variables en Go son contenedores para almacenar datos 📄. Una variable contiene un valor 💡, que puede ser cualquier cosa, desde un número 🔢 hasta una cadena de texto 📝 o un objeto 🎁. Cuando se define una variable, se especifica el nombre de la variable 🏷️, el tipo de datos 🔎 al que pertenece y el valor que se le asigna 🎯. Una vez que se ha definido una variable, se puede usar su nombre para referirse a ella y para acceder a su valor 🔗.

Existen dos tipos principales de variables en Go 🟢: las variables locales 🏠 y las variables globales 🌐. Las variables locales se definen dentro de una función 🧩 y solo están disponibles dentro de esa función. Las variables globales se definen fuera de cualquier función y están disponibles en todo el programa 🌍.

#### **Variables locales 🏠:**

Las variables locales son aquellas variables definidas dentro de una función 🧩 o bloque de código 📜. Por ejemplo, en el siguiente código, la variable "a" es una variable local:

```go
func main() {       
   var a = 5        
}

// Esta línea sacará error porque 'a' no esta definida fuera de la función
print(a);

```

#### **Variables globales 🌐:**

Las variables globales son variables definidas fuera de cualquier función 🧩 o bloque de código 📜. Por ejemplo, en el siguiente código, la variable "b" es una variable global:

```go
// Una variable global
var b = 10              

func main() { 
   // La variable B esta disponible aqui
   print(b) 
}

// Aquí también estara definida
func otra (){
	println(variable)
}

```

Las variables también se pueden definir como constantes 💎. Las constantes son variables que tienen un valor fijo 🔒 y no pueden cambiarse durante la ejecución del programa 🚀.

#### **Constantes 💎:**

Las constantes son valores que un programa no puede cambiar durante su ejecución 🚀. Por ejemplo, en el siguiente código, la constante "PI" es una constante:

```go
const PI = 3.14159

//Esto ocacionará un error
PI = 88
```

#### Formas de inicializar variables 📦:

Las variables son contenedores para almacenar valores durante la ejecución de un programa 🚀. Por ejemplo, en el siguiente código, la variable "x" es una variable que inicializamos junto con su tipo de manera vacía, y luego le asignamos su valor.

```go
// Declaración de la variable 'x' en Go, vacía
var x int
// Asignación del valor entero 5 a la variable 'x'  
x = 5          
```

Pero también es posible asignarla sin el tipo siempre y cuando establescamos su valor, esto se conoce como inferencia de tipo:

```go
// Declaración de la variable 'y' en go en una sola línea donde no expresamos el tipo
var y = 10
// También funcionaría aunque es innecesario expresar el tipo
var y int = 10 
```

#### **Asignación de una variable a otra 📍:**

Es posible asignarle el contenido de una variable a otra variable, esto se conoce como paso por valor o copia.

```go
// Declaración de la variable 'x' en Go
var x int
// Asignación del valor entero 5 a la variable 'x'
x = 5
fmt.Println("Valor de x: ", x)
var obj = x
fmt.Println("Valor de obj: ", obj)
x = 99
// El valor de obj seguira siendo 5 a pesar que el valor de x cambió a 99
fmt.Println("Valor de obj: ", obj)


```

#### 🎨Tipos de datos en Detalle🔍:

Los tipos de datos en Go 🟢 determinan el tipo de valores que se pueden almacenar y manipular dentro de un programa 💾. Go tiene una variedad de tipos de datos primitivos 🧱, incluidos enteros (integers)🔢, flotantes (floats)🔟, cadenas (strings)📝 y booleanos (booleans)🔘, así como también tipos de datos compuestos como arreglos (arrays)🔢🔢, estructuras (structs)🏗️ y punteros (pointers)📍. También puedes definir tus propios tipos de datos usando la palabra clave `type` 📚.

**Enteros (Integers)🔢:**

Los enteros se usan para almacenar valores numéricos enteros, como 3, -2 y 0. Hay varios tipos de enteros, incluidos `int8`, `int16`, `int32`, `int64`, `uint8`, `uint16`, `uint32` y `uint64`.

**Valores de punto flotante (Floats)🔟:**

Los valores de punto flotante se usan para almacenar valores numéricos con decimales, como 3.14, -2.5 y 0.0. Hay dos tipos de punto flotante: `float32` y `float64`.

**Cadenas (Strings)📝:**

Las cadenas se usan para almacenar secuencias de caracteres, como "hola mundo" 🌍. Una cadena es una secuencia de caracteres codificada en UTF-8 🔠.

**Booleanos (Booleans)🔘:**

Los booleanos se usan para almacenar valores lógicos, como `true` o `false`.

**Arreglos (Arrays)🔢🔢:**

Los arreglos se usan para almacenar un conjunto de valores del mismo tipo, como `1, 2, 3, 4`. Los arreglos no tienen un tamaño fijo, por lo que una vez creado, se pueden añadir o modificar elementos 📏. Por Ejemplo:

```go
var numeros = []float64{5, 10, 9, 8, 7}
// Añadimos un elemento a nuestro arreglo
numeros = append(numeros, 18)
```

**Estructuras (Structs)🏗️:**

Las estructuras se usan para almacenar un conjunto de variables relacionadas, como una el nombre y el precio de un producto en el mismo lugar. Las estructuras son similares a los arreglos, pero pueden contener diferentes tipos de datos 🔀.

Por ejemplo, una estructura que representa un producto puede ser así:

```go
type Producto struct {
	Nombre string
	Precio float64
}

```

Y se podría usar asi:

```go
var producto Producto
producto.Nombre = "Producto 1"
producto.Precio = 100
```

En resumen, Go ofrece una amplia variedad de tipos de datos y estructuras para adaptarse a las necesidades de tu programa 🎯. Conocerlos y utilizarlos adecuadamente te permitirá escribir programas más eficientes y mejor organizados 🚀.

#### Nota sobre las comillas 🎯:

En Go, las comillas dobles y simples tienen usos específicos y distintos:

1.  **Comillas Dobles (" ")**:
    
    -   Se utilizan para representar cadenas de texto (strings). Una cadena en Go es una secuencia de bytes inmutable y puede contener cualquier dato que puedan representar los bytes.
        
    -   Ejemplo: `"Hola Mundo"` es una cadena de texto.
        
    -   Las cadenas delimitadas por comillas dobles pueden contener caracteres especiales escapados, como `\n` (salto de línea) o `\t` (tabulación).
        
    
2.  **Comillas Simples (' ')**:
    
    -   Se usan para representar un único carácter, conocido como un rune en Go. Un rune es un tipo de dato que representa un punto de código Unicode.
        
    -   Ejemplo: `'A'` es un rune que representa la letra A.
        
    -   A diferencia de las cadenas, un rune es un valor entero y ocupa 4 bytes (32 bits) en Go.
        
    

#### Veamos diez ejemplos con su respectiva solución:

Calcular la suma de dos números: Solución:

```go
package main

import "fmt"

func main() {
	fmt.Println("Calcular la suma de dos números: Solución:")
	var a = 5
	var b = 10
	var resultado = sumar(a, b)
	fmt.Println("El resultado es: ", resultado)
}

func sumar(a int, b int) int {
	var sum = a + b
	return sum
}
```

Calcular el producto de dos números: Solución:

```go
package main

import "fmt"

func main() {
	fmt.Println("Calcular el producto de dos números: Solución:")
	var a = 5.0
	var b = 10.1
	var resultado = producto(a, b)
	fmt.Println("El resultado es: ", resultado)
}

// Al usar numeros flotantes automaticamente se asigna el tipo float64
func producto(a float64, b float64) float64 {
	var sum = a * b
	return sum
}


```

Calcular el promedio de una lista de números: Solución:

```go
func main() {
	fmt.Println("Calcular el producto de dos números: Solución:")
        // Esto es un arreglo en Golang de numeros, es decir una colección de numeros
	var numeros = []float64{5, 10, 9, 8, 7}
	var promedio = promedio(numeros)
	fmt.Println("El resultado es: ", promedio)
}

func promedio(numeros []float64) float64 {
	var sum = 0.0
	for _, numero := range numeros {
		sum += numero
	}
	return sum / float64(len(numeros))
}


```

Calcular el factorial de un número: Solución:

```go
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

```

Calcular el área de un círculo: Solución:

```go
import "math"

func areaCirculo(radio float64) float64 {
	return math.Pi * radio * radio
}


```

Calcular el perímetro de un cuadrado: Solución:

```go
func perimetroCuadrado(lado float64) float64 {
	return lado * 4
}
```

Calcular el volumen de una esfera: Solución:

```go
func volumenEsfera(radio float64) float64 {
	return (4 / 3) * math.Pi * radio * radio * radio
}


```

Calcular el área de un triángulo: Solución:

```go
func areaTriangulo(base float64, altura float64) float64 {
	return base * altura / 2
}


```

Calcular la hipotenusa de un triángulo rectángulo: Solución:

```go
func hipotenusa(cateto1 float64, cateto2 float64) float64 {
	return math.Sqrt(cateto1*cateto1 + cateto2*cateto2)
}
```

Pedir al usuario dos numeros y visualizar el resultado por pantalla

```go
func ingresoDatos(){
	fmt.Println("Introduzca dos números: ")
	var a, b float64
	fmt.Println("Introduzca el número 1: ")
	fmt.Scanln(&a)
	fmt.Println("Introduzca el número 2: ")
	fmt.Scanln(&b)
	fmt.Println("Los numeros que ingresaste son: ", a, b)
}
```

A continuacion te dejo 20 ejercicios mas para que practiques y envies la solucion como un Pull requests a este repo: [**https://github.com/seagomezar/goparaprincipiantes**](https://github.com/seagomezar/goparaprincipiantes)

1.  Imprimir la longitud de una cadena de texto.
    
2.  Comparar dos números enteros y mostrar el mayor.
    
3.  Calcular el cuadrado de un número.
    
4.  Crear una función para convertir una cadena a mayúsculas.
    
5.  Crear una función para convertir una cadena a minúsculas.
    
6.  Determinar si una cadena contiene una subcadena determinada.
    
7.  Determinar si un número es par o impar.
    
8.  Determinar si una cadena es palíndroma.
    
9.  Incrementar un número en una unidad.
    
10.  Determinar si un número está dentro de un rango determinado.
    
11.  Obtener el índice de la primera aparición de una subcadena en una cadena.
    
12.  Obtener la suma de dos números enteros.
    
13.  Comprobar si un número es divisible por otro.
    
14.  Determinar si una cadena es un número.
    
15.  Calcular el máximo común divisor (MCD) de dos números.
    
16.  Determinar si una cadena comienza con una subcadena determinada.
    
17.  Obtener el tamaño de una matriz.
    
18.  Obtener el número de elementos únicos en una matriz.
    
19.  Determinar si un número es primo.
    
20.  Obtener los índices de los elementos duplicados en una matriz.