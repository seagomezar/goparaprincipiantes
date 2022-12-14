Go en Detalle

Go es un lenguaje de programación compilado moderno creado por Google para su uso interno. Está diseñado para ser fácil de usar, lo que permite a los desarrolladores crear programas de forma rápida y eficiente.

Go es un lenguaje compilado, lo que significa que el código escrito en Go se compila a lenguaje de máquina antes de que se ejecute. Esto significa que el código Go es eficiente y escalable, ya que los compiladores pueden hacer optimizaciones en el código para mejorar el rendimiento.

Go también es un lenguaje con programación concurrente, lo que significa que los programas se pueden dividir fácilmente en varias tareas que se ejecutan al mismo tiempo. Esto facilita la creación de programas complejos que pueden aprovechar al máximo el hardware disponible.

Además, Go es un lenguaje de programación de código abierto, lo que significa que los desarrolladores pueden contribuir al código fuente y mejorar el lenguaje. Esto permite que el lenguaje se mantenga actualizado y mejorado constantemente.

Finalmente, Go usa un sistema de paquetes para organizar el código. Esto ayuda a los desarrolladores a mantener el código organizado y fácil de administrar.

Go Compiler, Go Runtime y OS API detalladamente

Go Compiler es el compilador de Go que convierte el código Go en lenguaje de máquina. El compilador Go toma el código fuente escrito en Go y lo compila a lenguaje de máquina antes de que se ejecute. Esto permite que el código Go sea eficiente y escalable.

Go Runtime es el entorno de tiempo de ejecución para las aplicaciones Go. El entorno de tiempo de ejecución Go gestiona los recursos de la aplicación, como la memoria, el almacenamiento de datos, la concurrencia y la comunicación entre procesos. El entorno de tiempo de ejecución también se encarga de la ejecución del código compilado.

OS API es el conjunto de funciones de la API del sistema operativo que el código Go puede llamar para realizar tareas específicas del sistema operativo. Estas funciones permiten que el código Go interactúe con el sistema operativo para realizar tareas como la creación de procesos, la gestión de archivos y la comunicación de red.

Vamos con una metafora de Harry Potter que nos permita entender mejor como funciona Go Internamente

Go Compiler sería como la varita mágica de Harry Potter. Al igual que la varita mágica, el compilador Go toma el código escrito en Go y lo transforma en algo diferente. En este caso, el compilador Go toma el código escrito en Go y lo compila a lenguaje de máquina para que se ejecute.

Go Runtime sería como la bola de cristal de Harry Potter. Al igual que la bola de cristal, el entorno de tiempo de ejecución Go proporciona una vista en tiempo real de cómo se ejecutan las aplicaciones. El entorno de tiempo de ejecución Go gestiona los recursos de la aplicación, como la memoria, el almacenamiento de datos, la concurrencia y la comunicación entre procesos.

OS API sería como la Capa de invisibilidad de Harry Potter. Al igual que la Capa de invisibilidad, la API del sistema operativo permite al código Go interactuar con el sistema operativo para realizar tareas específicas del sistema operativo. Esto permite que el código Go realice tareas como la creación de procesos, la gestión de archivos y la comunicación de red.

Una vista a los elementos mas importantes:

┌────────────┐                                    ┌─────────────┐
│ User Code  │───────────────────────────────────▶ │ Go Compiler │
└────────────┘                                    └─────────────┘
               │                                              │
               │                                              │
               │                                              │
               │                                              │
               │                                              │
               │                                              │
               │                                              │
               │                                              │
               ▼                                              ▼
┌────────────┐                                    ┌─────────────┐
│  Assembler │◀──────────────────────────────────│  Go Runtime │
└────────────┘                                    └─────────────┘
               │                                              │
               │                                              │
               │                                              │
               │                                              │
               │                                              │
               │                                              │
               │                                              │
               │                                              │
               ▼                                              ▼
┌────────────┐                                    ┌─────────────┐
│  Machine   │◀──────────────────────────────────│    OS API   │
└────────────┘                                    └─────────────┘

Hola mundo en Go

``` go

// Esta es una implementación del programa "Hola Mundo" en Go.

package main // Esta línea indica que este programa se ejecutará como un programa independiente.

import "fmt" // Esta línea importa la biblioteca fmt, que contiene funciones útiles para imprimir datos.

func main() { // Esta línea define la función main, que es la función principal de este programa.

	fmt.Println("¡Hola mundo!") // Esta línea imprime el mensaje "¡Hola mundo!" en la pantalla.

} // Esta línea marca el fin de la función main.
```

Veamos el proceso para ejecutar el codigo:

Go Compiler: El compilador Go compila el código fuente en el archivo binario ejecutable.

Go Runtime: El tiempo de ejecución de Go proporciona los recursos necesarios para la ejecución del programa.

OS API: La API del sistema operativo se encarga de proporcionar las funciones necesarias para que el programa Go se comunique con el sistema operativo, permitiéndole realizar tareas como la importación de paquetes (import "fmt"), la impresión de la salida (fmt.Println("Hola mundo!")).

El compilador Go compilará el código fuente en un archivo binario ejecutable. Una vez compilado, el código se convertirá en una secuencia de instrucciones binarias que serán ejecutadas por el tiempo de ejecución de Go, con el apoyo de la API del sistema operativo.

El siguiente es un ejemplo de cómo el compilador Go convertirá el código fuente anterior en un archivo binario ejecutable:

``` go
package main

import "fmt"

func main() {
	// Convertir la llamada a Println a una instrucción binaria
	0x0048 0x8D 0x35 0x00 0x00 0x00 0x00 0xE8 0x00 0x00 0x00 0x00 0x48 0x8B 0xD0 0xC3

	// Proporcionar los recursos necesarios para la ejecución del programa
	0x48 0x83 0xEC 0x20 0x48 0x8B 0x05 0x00 0x00 0x00 0x00 0x48 0x8B 0x00 0x48 0x8B 0x40 0x08 0x48 0x85 0xC0 0x74 0x07 0x48 0x83 0xC4 0x20 0xC3

	// Proporcionar las funciones necesarias para que el programa Go se comunique con el sistema operativo
	0x48 0x83 0xEC 0x20 0x48 0x8B 0x05 0x00 0x00 0x00 0x00 0x48 0x8B 0x00 0x48 0x8B 0x40 0x08 0x48 0x85 0xC0 0x74 0x07 0x48 0x83 0xC4 0x20 0xC3

	fmt.Println("Hola mundo!")
}
```


Los tipos de datos en Go

En Go, los tipos de datos se usan para definir y manipular los datos que se usan en los programas. Go tiene una variedad de tipos de datos, como enteros, flotantes, cadenas, booleanos y otros. Estos tipos de datos se pueden usar para almacenar y manipular los datos.

Enteros: Los enteros son números sin decimales. Se pueden usar para almacenar valores enteros, como 5, 10, -15, etc.

Flotantes: Los flotantes son números con decimales. Se pueden usar para almacenar valores reales con decimales, como 1.5, -3.14, etc.

Cadenas: Las cadenas son secuencias de caracteres. Se pueden usar para almacenar texto, como “hola mundo”, “¡Hola!”, etc.

Booleanos: Los booleanos son valores lógicos que se pueden usar para almacenar valores verdaderos o falsos.

Además de estos tipos de datos básicos, Go también tiene otros tipos de datos, como estructuras, matrices, punteros y otros. Estos tipos de datos se usan para manipular los datos de forma más avanzada.

**Tipos de datos en Go**

- Enteros
  - Números sin decimales
  - Se usan para almacenar valores enteros
- Flotantes
  - Números con decimales
  - Se usan para almacenar valores reales con decimales
- Cadenas
  - Secuencias de caracteres
  - Se usan para almacenar texto
- Booleanos
  - Valores lógicos
  - Se usan para almacenar valores verdaderos o falsos
- Estructuras
  - Colecciones de datos relacionados
  - Se usan para manipular datos de forma avanzada
- Matrices
  - Colecciones de datos indexados
  - Se usan para manipular datos de forma avanzada
- Punteros
  - Referencias a la memoria
  - Se usan para manipular datos de forma avanzada

Funciones en Go

Las funciones en Go son bloques de código que se pueden usar para realizar una tarea específica. Las funciones permiten a los programadores dividir el código en partes más manejables, lo que facilita la creación de programas complejos.

Las funciones en Go se definen con la palabra clave "func", seguida del nombre de la función y los parámetros de entrada. Luego, se especifica el cuerpo de la función, que contiene el código que se ejecutará cuando se llame a la función.

Las funciones también se pueden usar para devolver valores. Esto se logra devolviendo uno o más valores desde el cuerpo de la función. Esto permite a los programadores manipular los datos de forma más eficiente.

Las funciones también se pueden usar para crear bloques de código reutilizables. Esto significa que una función se puede usar varias veces en un programa sin tener que escribir el mismo código varias veces. Esto hace que sea más fácil crear y mantener programas complejos.

Estos son algunos ejemplos de funciones sencillas para empezar a entender como escribirlas en go:

Ejemplo 1: Una función para calcular la suma de dos números enteros:

``` go
func suma(a int, b int) int {
    return a + b
}
```

``` go
Ejemplo 2: Una función para imprimir un mensaje en la consola:

func imprimirMensaje(mensaje string) {
    fmt.Println(mensaje)
}
```

Ejemplo 3: Una función para calcular el área de un círculo dado su radio:

``` go
func calcularAreaCirculo(radio float64) float64 {
    return math.Pi * radio * radio
}
```

Como ejecutar un archivo de go

Para ejecutar un archivo de Go, primero debe compilar el archivo en lenguaje de máquina usando el comando 'go build'. Esto creará un archivo ejecutable que se puede ejecutar con el comando 'go run'. Si se desea ejecutar el archivo sin compilarlo, se puede usar el comando 'go run'.

crea un diagrama de flujo en markdown para una funcion en go que solicite dos numeros e imprima la suma de dichos numeros.

**Diagrama de flujo de función de Go**

INICIO

- Solicitar dos números al usuario
- Calcular la suma de los dos números
- Imprimir la suma

FIN

Veamos como traducir el algoritmo anterior a GO:

``` go
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

Vamos con otro ejemplo esta vez para determinar si el numero es par o impar:

**Diagrama de flujo de programa en Go**

INICIO

- Solicitar un número al usuario
- Calcular el resto de la división del número entre 2
- Si el resto es 0
  - Imprimir "El número es par"
- Si el resto es diferente de 0
  - Imprimir "El número es impar"

FIN

Veamos como traducir el algoritmo anterior a GO:

``` go
// Este programa solicita un número al usuario e imprime si es par o impar

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

```

Si aun te cuesta entender el codigo anterior este es un diagrama que te puede ayudar:


**Diagrama visual del programa en Go**

```
main()
    |
    |
    |
    V
 Solicitar número
    |
    |
    |
    V
 Calcular resto
    |
    |
    |
    V
   if resto == 0
    |
    |
    |
    V
  Imprimir "par"
    |
    |
    |
    V
   else
    |
    |
    |
    V
  Imprimir "impar"
    |
    |
    |
    V
  Fin
```