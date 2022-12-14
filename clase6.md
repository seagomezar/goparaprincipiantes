Profundizando en las estructuras en go

En Go, una estructura es una colección de grupos de datos relacionados almacenados como un tipo único. Proporcionan una forma de almacenar y organizar datos relacionados bajo una sola estructura. Las estructuras se pueden definir usando la palabra clave struct.

Un ejemplo de una estructura en Go es la siguiente:

``` go
type persona struct {
    nombre string
    edad int
}
```

En el ejemplo anterior, tenemos una estructura llamada persona. Esta estructura tiene dos campos: nombre y edad, que son ambos de tipo string y int, respectivamente.

Ahora, para usar esta estructura en nuestro código, necesitamos crear una instancia de la misma. Esto se puede hacer de la siguiente manera:

``` go
miPersona := persona{nombre: "Juan", edad: 25}
```

En el ejemplo anterior, hemos creado una instancia de la estructura persona con el nombre de miPersona. Esta instancia contiene los valores "Juan" y 25 para los campos nombre y edad, respectivamente.

Para acceder a los campos de una estructura, se puede usar el operador punto (.). Por ejemplo, para acceder al campo nombre de miPersona, se puede usar el siguiente código:
``` go
nombre := miPersona.nombre
```

Además de los campos, las estructuras también pueden tener métodos asociados. Estos se definen de la misma manera que se definen los métodos para los tipos de datos. Por ejemplo, podemos agregar el siguiente método a la estructura persona:

``` go
func (p persona) saludar() string {
    return "Hola, mi nombre es " + p.nombre
}
```

Ahora, para usar este método, podemos llamarlo de la siguiente manera:

``` go
mensaje := miPersona.saludar()
```

En el ejemplo anterior, hemos llamado el método saludar de miPersona y guardado el resultado en la variable mensaje. El resultado sería el siguiente:

"Hola, mi nombre es Juan"


Estructuras versus Clases

Go no tiene clases. En su lugar, Go tiene estructuras que se pueden usar para modelar datos y comportamientos. Las estructuras se definen con la palabra clave struct. Aunque Go no tiene clases, algunos conceptos de la programación orientada a objetos, como la herencia, se pueden simular usando estructuras.

Otra diferencia es que las estructuras no tienen herencia, mientras que las clases sí lo hacen. Esto significa que una estructura no puede heredar propiedades o métodos de otra estructura, mientras que una clase puede heredar propiedades y métodos de otra clase.

Las estructuras son pasadas por valor. Esto significa que cualquier cambio hecho a una estructura no se reflejará en la estructura original.

10 ejemplos de estructuras en Go

1. Persona:

``` go
type persona struct {
    nombre string
    edad int
}
```

2. Coche:

``` go
type coche struct {
    marca string
    modelo string
    año int
}
```

3. Estudiante:

``` go
type estudiante struct {
    nombre string
    edad int
    notas []int
}
```

4. Ciudad:

``` go
type ciudad struct {
    nombre string
    población int
    habitantes []persona
}
```

5. Juego:

``` go
type juego struct {
    nombre string
    jugadores []persona
    duración int
}
```

6. Película:

``` go
type película struct {
    nombre string
    actores []persona
    duración int
}
```

7. Libro:

``` go
type libro struct {
    nombre string
    autor string
    páginas int
}
```

8. País:

``` go
type país struct {
    nombre string
    ciudades []ciudad
    población int
}
```

9. Empresa:

``` go
type empresa struct {
    nombre string
    empleados []persona
    fundación int
}
```

10. Música:

``` go
type música struct {
    nombre string
    artista string
    duración int
}
```

Estructuras compuestas:

En Go, se pueden crear estructuras anidadas, que son estructuras dentro de estructuras. Por ejemplo, podemos crear una estructura llamada Animal que contenga una estructura llamada Características:

``` go
type Animal struct {
    nombre string
    Características struct {
        tamaño int
        color string
    }
}
```

En el ejemplo anterior, la estructura Animal contiene un campo llamado nombre y una estructura anidada llamada Características. Esta estructura anidada contiene los campos tamaño y color.

Para acceder a los campos de la estructura anidada, necesitamos usar el operador punto (.) dos veces. Por ejemplo, para acceder al campo tamaño en una instancia de la estructura Animal, se puede usar el siguiente código:

``` go
miAnimal := Animal{nombre: "Perro", Características: struct{tamaño: 5, color: "negro"}}
tamaño := miAnimal.Características.tamaño
```

Conclusiones

En conclusión, las estructuras en Go son una forma de almacenar y organizar datos relacionados bajo una sola estructura. Estas estructuras son similares, pero no iguales, a las clases en otros lenguajes de programación. Estas estructuras también se pueden anidar para crear estructuras más complejas. Finalmente, las estructuras son pasadas por valor, mientras que las clases son pasadas por referencia.


A continuacion te dejo 10 ejercicios mas para que practiques las estructuras en Go y envies la solucion como un Pull requeste a este repositorio:

1. Crea una estructura llamada Automóvil que contenga los campos marca, modelo y año.

2. Crea una estructura llamada Película que contenga los campos nombre, director y duración.

3. Crea una estructura llamada Ciudad que contenga los campos nombre, población y ubicación.

4. Crea una estructura llamada País que contenga los campos nombre, capital y moneda.

5. Crea una estructura llamada Persona que contenga los campos nombre, edad y dirección.

6. Crea una estructura llamada Libro que contenga los campos nombre, autor y páginas.

7. Crea una estructura llamada Juego que contenga los campos nombre, jugadores y duración.

8. Crea una estructura llamada Música que contenga los campos nombre, artista y duración.

9. Crea una estructura llamada Empresa que contenga los campos nombre, empleados y fundación.

10. Crea una estructura anidada llamada Animal que contenga los campos nombre, tamaño y color.