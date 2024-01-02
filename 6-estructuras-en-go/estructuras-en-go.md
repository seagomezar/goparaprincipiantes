En Go, una estructura (o `struct`) es una colección de variables (campos) agrupadas bajo un nombre. Las estructuras son fundamentales para organizar y manejar datos relacionados de forma eficiente. A diferencia de otros lenguajes que utilizan clases, Go se basa en estructuras para modelar datos y comportamientos.💡.

Un ejemplo de una estructura en Go es la siguiente:

```
// Define un nuevo tipo llamado 'Persona'.
type Persona struct {
    // Declara un campo 'nombre' de tipo string dentro del struct 'persona'.
    nombre string
    // Declara un campo 'edad' de tipo int (número entero) dentro del struct 'persona'.
    edad int
}

```

En el ejemplo anterior, tenemos una estructura llamada `persona` 😊. Esta estructura tiene dos campos: `nombre` y `edad`, que son ambos de tipo `string` e `int`, respectivamente.

Ahora, para usar esta estructura en nuestro código, necesitamos crear una instancia de la misma. Esto se puede hacer de la siguiente manera:

```
// Ambas funcionan exactamente igual

miPersona := Persona{nombre: "Juan", edad: 25}

var miPersona Persona = Persona{nombre: "Juan", edad: 25}
```

En el ejemplo anterior, hemos creado una instancia de la estructura `Persona` con el nombre de `miPersona` 🙌. Esta instancia contiene los valores "Juan" y 25 para los campos `nombre` y `edad`, respectivamente.

Para acceder a los campos de una estructura, se puede usar el operador punto (.) 🔎. Por ejemplo, para acceder al campo `nombre` de `miPersona`, se puede usar el siguiente código, si quisieramos imprimir los datos.

```
fmt.Println("Nombre:", miPersona.nombre)
fmt.Println("Edad:", miPersona.edad)
```

Además de los campos, las estructuras también pueden tener métodos asociados 🚦. Aunque no en su definición. Esto es una particularidad de Go ya que el metodo se le agrega a la estructura luego de su declaración. Por ejemplo, podemos agregar el siguiente método a la estructura `Persona`:

```
func (p Persona) saludar() string {
    return "Hola, mi nombre es " + p.nombre
}
```

Ahora, para usar este método, podemos llamarlo de la siguiente manera:

```
mensaje := miPersona.saludar()
```

**Nota**: La diferencia entre una función comun y un metodo, radica en el parentesis donde se escribe a que estructura pertenece y no lo debes confundir con los parametros que recibe el metodo. Por ejemplo: func (p Persona) saludar() es diferente de esto func saludar(p Persona). Ya que el primer parentesis luego de la palabra func indica a que estructura pertenece, de lo contrario sera una funcion normal.

En el ejemplo anterior, hemos llamado el método `saludar` de `miPersona` y guardado el resultado en la variable `mensaje` 📨. El resultado sería el siguiente:

```
"Hola, mi nombre es Juan"
```

#### **🚩 Estructuras versus Clases 🚩**

Go no tiene clases ❌. En su lugar, Go tiene estructuras que se pueden usar para modelar datos y comportamientos. Las estructuras se definen con la palabra clave `struct` 🔑. Aunque Go no tiene clases, algunos conceptos de la programación orientada a objetos, como la herencia, se pueden simular usando estructuras 🎭.

Otra diferencia es que las estructuras no tienen herencia, mientras que las clases sí lo hacen. Esto significa que una estructura no puede heredar propiedades o métodos de otra estructura, mientras que una clase puede heredar propiedades y métodos de otra clase 👨‍👩‍👦‍👦.

Las estructuras son pasadas por valor. Esto significa que cualquier cambio hecho a una estructura no se reflejará en la estructura original 🔄.

#### **Algunos Ejemplos de estructuras en Go**

Coche:

```
// Define un objeto 'coche' usando la notación literal de objetos en JavaScript.
type Coche struct {
    // Declara una propiedad 'marca' de tipo string dentro del objeto 'coche'.
    marca string
    // Declara una propiedad 'modelo' de tipo string dentro del objeto 'coche'.
    modelo string
    // Declara una propiedad 'año' de tipo número (integer) dentro del objeto 'coche'.
    año int
}

```

Estudiante:

```
type Estudiante struct {
	nombre string
	edad int
	curso string
	notas [3]float64
}


```

#### **🔩 Estructuras compuestas 🔩**

En Go, se pueden crear estructuras anidadas, que son estructuras dentro de estructuras 🎁. Por ejemplo, podemos crear una estructura llamada Tractor que contenga una estructura llamada `Carro`:

```
type Carro struct {
	brand string
	year  int
}

type Tractor struct {
	Car
	capacity int
}
```

En el ejemplo anterior, la estructura Tractor contiene un campo llamado `capacity` y una estructura anidada llamada `Car` 🎎. Esta estructura anidada contiene los campos `brand` y `year`.

Para acceder a los campos de la estructura anidada, necesitamos usar el operador punto (.) dos veces 🔍🔍. Por ejemplo, para acceder al campo `brand` en una instancia de la estructura `Carro`, se puede usar el siguiente código:

```
miTractor := Truck{Car{brand: "Ford", year: 2022}, 2}
fmt.Println("------------------------")
fmt.Println("Tractor:", miTractor.Car.brand)
```

#### Campos anónimos en structs

En Go, es posible no declarar el nombre del campo de nuestro struct y colocar únicamente el tipo de dato. Hecho así, los campos adoptarán el nombre del tipo de dato y podemos acceder a ellos usándolos.

```
type Videogame struct { 
    string 
    int 
}

myVideogame := Videogame{string: "Titulo", int: 2017}
fmt.Println(myVideogame)
// imprime {Titulo 2017}
```

#### Modificadores de Acceso en Go

A diferencia de otros lenguajes de programación, como Java y [Python](https://www.digitalocean.com/community/tutorial_series/how-to-code-in-python-3), que utilizan* modificadores de acceso* como `public`, `private` o `protected` para especificar el ámbito, Go determina si un elemento es `exported` o `unexported` según la forma en que se declara. Exportar un elemento, en este caso, lo hace `visible` fuera del paquete actual. Si no se exportó, solo es visible y utilizable dentro del paquete en el que se definió.

Esta visibilidad externa se controla escribiendo en mayúscula la primera letra del elemento declarado. Todas las declaraciones, como `Types`, `Variables`, `Constants`, `Functions` y demás que comienzan con una letra en mayúscula son visibles fuera del paquete actual.

#### **📝 Conclusiones 📝**

En conclusión, en este blog post hemos explorado el concepto de estructuras en el lenguaje de programación Go. Hemos aprendido que las estructuras son colecciones de datos relacionados que permiten almacenar y organizar información de manera eficiente. Además, hemos visto cómo se pueden definir, instanciar y utilizar estructuras en Go, así como la posibilidad de agregar métodos asociados a ellas. También hemos destacado algunas diferencias importantes entre las estructuras en Go y las clases en otros lenguajes, como la ausencia de herencia y el paso por valor en las estructuras. Además, se ha mencionado la posibilidad de crear estructuras anidadas y campos anónimos, así como el uso de mayúsculas para controlar la visibilidad de elementos en Go. En resumen, las estructuras son una parte fundamental de la programación en Go, permitiendo la organización y manipulación de datos de manera eficaz y flexible.

#### **🏋️‍♂️ Ejercicios para practicar 🏋️‍♂️**

A continuación te dejo algunos ejercicios más para que practiques las estructuras en Go y envíes la solución como un Pull request a este repositorio: [Go Para Principiantes](https://github.com/ejemplos-go-para-principiantes)

1.  Crear una estructura `EquipoDeportivo` que contenga información sobre jugadores, entrenadores y partidos jugados.
    
2.  Crear una estructura `Supermercado` que contenga información sobre productos, empleados y ventas.
    
3.  Crear una estructura `Universidad` que contenga información sobre estudiantes, profesores y cursos ofrecidos.
    
4.  Crea una estructura llamada `Libro` que contenga campos como `Título`, `Autor` y `Año de Publicación`. Luego, instancia un objeto `libro1` y muestra sus datos.
    
5.  Define un método en la estructura `Coche` que calcule la velocidad promedio en base a la distancia recorrida y el tiempo. Luego, crea una instancia de un coche y utiliza este método para calcular la velocidad promedio.
    
6.  Crea una estructura `Escuela` que contenga un campo `nombre` y un slice de estructuras `Estudiante`, donde cada estudiante tiene un `nombre`, `edad` y `curso`. Crea una instancia de `Escuela` con varios estudiantes y muestra sus datos.
    
7.  Define una estructura llamada `Producto` que tenga un campo anónimo de tipo `Precio` con un campo `valor` de tipo float64. Luego, crea un producto y muestra su precio.
    
8.  Simula la herencia en Go creando una estructura `Animal` con campos como `nombre` y `edad`, y luego crea estructuras `Perro` y `Gato` que incorporen la estructura `Animal`. Muestra los datos de un perro y un gato.
    
9.  Crea una estructura `CuentaBancaria` con campos como `saldo` y `nombreTitular`. Define métodos para depositar y retirar dinero de la cuenta. Realiza algunas transacciones y muestra el saldo final.
    
10.  Define una estructura `EquipoDeFútbol` que contenga un slice de nombres de jugadores. Agrega y elimina jugadores del equipo y muestra la lista actualizada.
    
11.  Define una estructura `Edificio` que contenga información sobre apartamentos, donde cada apartamento tiene un número y un área en metros cuadrados. Crea un método en `Edificio` para calcular el área total de todos los apartamentos.