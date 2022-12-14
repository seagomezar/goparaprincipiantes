Condicionales en Go

Los condicionales en Go son una forma de controlar el flujo de un programa. Esto se logra con la sentencia if (si), la cual ejecuta una o más instrucciones si la condición especificada se evalúa como verdadera. La estructura de un condicional if en Go es la siguiente:

``` go
if condición {
    // instrucciones
}
```

If: La sentencia if (si) se usa para ejecutar una o más instrucciones si la condición especificada se evalúa como verdadera. La sentencia if consta de una condición y un bloque de código. Si la condición es verdadera, se ejecutará el bloque de código.

Else: La sentencia else (de lo contrario) se usa para ejecutar una o más instrucciones si la condición especificada se evalúa como falsa. La sentencia else consta de una condición y un bloque de código. Si la condición es falsa, se ejecutará el bloque de código.

Else if: La sentencia else if (de lo contrario, si) se usa para comprobar varias condiciones. Si la primera condición se evalúa como falsa, se evaluará la siguiente condición. Si esta segunda condición se evalúa como verdadera, se ejecutará el bloque de código asociado a ella. Esto se puede hacer tantas veces como sea necesario.

Los switchs en Go son una forma concisa de controlar el flujo de un programa. Esto se logra mediante una sentencia switch, que se utiliza para comparar una variable con varias expresiones. Si la variable coincide con una de las expresiones, se ejecutará el bloque de código asociado. La estructura de un switch en Go es la siguiente:

``` go
switch {
    case expresión1:
        // instrucciones
    case expresión2:
        // instrucciones
    ...
    case expresiónN:
        // instrucciones
    default:
        // instrucciones
}
```

Te dejo algunos ejemplos sobre el uso de condicionales en Go:

1. Ejemplo de If:

``` go
if x > 10 {
    fmt.Println("x es mayor que 10")
}
```

2. Ejemplo de If/Else:

``` go
if x > 10 {
    fmt.Println("x es mayor que 10")
} else {
    fmt.Println("x es menor o igual a 10")
}
```

3. Ejemplo de If/Else If/Else:

``` go
if x > 10 {
    fmt.Println("x es mayor que 10")
} else if x == 10 {
    fmt.Println("x es igual a 10")
} else {
    fmt.Println("x es menor que 10")
}
```

4. Ejemplo de If/Else anidado:

``` go
if x > 10 {
    if y > 10 {
        fmt.Println("x y y son mayores que 10")
    }
} else {
    fmt.Println("x o y son menores que 10")
}
```

5. Ejemplo de Switch:

``` go
switch x {
case 1:
    fmt.Println("x es 1")
case 2:
    fmt.Println("x es 2")
default:
    fmt.Println("x es otro número")
}
```


6. Ejemplo de Switch/Case con intervalos:

``` go
switch {
case x > 0 && x <= 10:
    fmt.Println("x está entre 0 y 10")
case x > 10 && x <= 20:
    fmt.Println("x está entre 10 y 20")
default:
    fmt.Println("x es otro número")
}
```

7. Ejemplo de Switch/Case con cadenas de caracteres:

``` go
switch x {
case "abc":
    fmt.Println("x es abc")
case "def":
    fmt.Println("x es def")
default:
    fmt.Println("x es otra cadena")
}
```

8. Ejemplo de Switch/Case con variables:

``` go
var text string = "abc"

switch text {
case "abc":
    fmt.Println("text es abc")
case "def":
    fmt.Println("text es def")
default:
    fmt.Println("text es otra cadena")
}
```

9. Ejemplo de Switch/Case con múltiples valores:

``` go
switch x {
case 1, 2, 3:
    fmt.Println("x es 1, 2 o 3")
case 4, 5, 6:
    fmt.Println("x es 4, 5 o 6")
default:
    fmt.Println("x es otro número")
}
```

10. Ejemplo de Switch/Case con expresiones:

``` go
switch {
case x > 0 && x <= 10:
    fmt.Println("x está entre 0 y 10")
case (x > 10 && x <= 20) || x % 2 == 0:
    fmt.Println("x está entre 10 y 20 o es un número par")
default:
    fmt.Println("x es otro número")
}
```

A continuacion te dejo 10 ejercicios mas para que practiques los condicionales en go y envies la solucion como un Pull requeste a este repositorio:

1. Imprima un mensaje diferente dependiendo de si un número es par o impar.
2. Dados dos números enteros, imprima el mayor de los dos.
3. Dados tres números enteros, imprima el mayor de los tres.
4. Imprima un mensaje diferente si un número es positivo o negativo.
5. Lea un número entero y determine si es divisible entre 5.
6. Lea un número entero y determine si es divisible entre 3 y 5 al mismo tiempo.
7. Lea dos números enteros y determine si el primero es mayor que el segundo.
8. Lea tres números enteros y determine si el primero es igual al segundo y al tercero.
9. Lea dos números enteros y determine si el primero es menor que el segundo o si son iguales.
10. Lea un número entero y determine si es par o impar y si es positivo o negativo.

Y también te dejo 10 ejercicios mas para que practiques los Switch en go y envies la solucion como un Pull requeste a este repositorio:

1. Imprima un mensaje diferente dependiendo de si un número está entre 0 y 5, entre 5 y 10, entre 10 y 20, o mayor que 20.
2. Dados dos números enteros, imprima la letra correspondiente a la calificación apropiada.
3. Dados tres números enteros, imprima la letra correspondiente a la calificación apropiada.
4. Lea un número entero y determine la clase de número que es.
5. Lea un número entero entre 0 y 5 y determine a qué letra corresponde.
6. Lea un número entero entre 5 y 10 y determine a qué letra corresponde.
7. Lea un número entero entre 10 y 20 y determine a qué letra corresponde.
8. Lea un número entero mayor que 20 y determine a qué letra corresponde.
9. Lea un número entero y determine si es par o impar y si es positivo o negativo.
10. Lea un número entero y determine si es divisible entre 2, 3, 4, 5 o 6.

