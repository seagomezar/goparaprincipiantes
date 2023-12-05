
# Condicionales en Go

🔥 Aprende a controlar el flujo de tu programa con condicionales y switch en Go 🔥

Los condicionales en Go son una forma de controlar el flujo de un programa. Puedes usar  if,  else,  else if  y  switch  para determinar qué instrucciones se ejecutarán basándose en ciertas condiciones.

### 🌟 If, Else y Else If 🌟

Los condicionales "if-else if" son una forma de tomar decisiones en programación. Funcionan de la siguiente manera:

1.  Se inicia con un "if" que evalúa una condición. Si la condición es verdadera, se ejecuta un bloque de código asociado a ese "if".
2.  Si la condición del primer "if" es falsa, se pasa al siguiente "else if" (si lo hay). Se evalúa la condición del "else if", y si es verdadera, se ejecuta el bloque de código asociado a ese "else if".
3.  Este proceso se repite para cada "else if" en secuencia hasta que se encuentra una condición verdadera o hasta que no hay más "else if". Si se encuentra una condición verdadera, se ejecuta el bloque de código correspondiente y se sale de la estructura condicional.
4.  Si ninguna de las condiciones "if" o "else if" es verdadera, se ejecuta el bloque de código del "else" (si está presente), que es el último recurso.

Veamos un primer ejemplo, un condicional que evalue si una variable es mayor o menor que el numero 10:

```javascript
package main

import "fmt"

func main() {

	var x int = 11
	if x > 10 {
		fmt.Println("x es mayor a 10")
	} else {
		fmt.Println("x es menor a 10")
	}

}
```

A diferencia de otros lenguajes los condicionales no requiren parentesis en la condicion a evaluar.

Veamos ahora un ejemplo de if else if, donde yo puedo adicionar tantos casos como desee mediante la sintaxis else if:

```javascript
func ejemploIfElseIf(){
	var x int = 11
	if x > 10 {
		fmt.Println("x es mayor a 10")
	} else if x == 10 {
		fmt.Println("x es igual a 10")
	} else {
		fmt.Println("x es menor a 10")
	}
}
```

### 🎛️ Switch en Go 🎛️

La estructura switch es una forma de control de flujo en programación que permite simplificar múltiples comparaciones de una variable contra diferentes valores. Funciona de la siguiente manera:

**Inicio de switch:**  La estructura comienza con la palabra clave switch, seguida de la variable o expresión que quieres comparar.

**Casos (case):**  Dentro del bloque de switch, defines diferentes casos usando la palabra clave case, seguida de un valor específico que quieres comparar con la variable o expresión del switch. Si la variable o expresión coincide con el valor del case, se ejecuta el bloque de código asociado a ese case.

**Ejecución de un Caso:**  Cuando se encuentra una coincidencia entre la variable y un case, se ejecuta el código asociado a ese case y luego se sale de la estructura switch. Esto significa que solo se ejecuta el código de un case coincidente.

**Break:**  Aunque en algunos lenguajes es necesario usar la palabra clave break al final de cada bloque case para evitar que el control fluya hacia el siguiente case, en algunos lenguajes modernos como Go, el break es implícito y no es necesario escribirlo.

**default:**  Puedes incluir un default al final del switch, que se ejecutará si ninguno de los case coincide. Es similar a un else en una estructura if-else.

-   Ejemplo de  switch:

```javascript
func ejemploSwitch(){
	var x int = 11
	switch x {
	case 1:
		fmt.Println("x es 1")
	case 2:
		fmt.Println("x es 2")
	case 3:
		fmt.Println("x es 3")
	case 4:
		fmt.Println("x es 4")
	case 5:
		fmt.Println("x es 5")
	case 6:
		fmt.Println("x es 6")
	case 7:
		fmt.Println("x es 7")
	case 8:
		fmt.Println("x es 8")
	case 9:
		fmt.Println("x es 9")
	case 10:
		fmt.Println("x es 10")
	default:
		fmt.Println("x es mayor a 10")
	}
}

```

-   Ejemplo de  switch/case  con intervalos:

```javascript
func ejemploSwitchRango() {
	var numeroMes int = 3
	var estacion string

	switch {
		case numeroMes >= 1 && numeroMes < 3:
			estacion = "Invierno"
		case numeroMes >= 3 && numeroMes < 6:
			estacion = "Primavera"
		case numeroMes >= 6 && numeroMes < 9:
			estacion = "Verano"
		case numeroMes >= 9 && numeroMes < 12:
			estacion = "Otoño"
		default:
			estacion = "Ninguna Estación del Año!"
	}
	fmt.Println(estacion)
}

```

-   Ejemplo de  switch/case  con cadenas de caracteres:

```javascript
func ejemploSwitchConTexto() {
	var x string = "def"
	// Evalúa el valor de la variable 'x' en una estructura 'switch'
	switch x {
	// Si 'x' es igual a "abc"
	case "abc":
		// Imprime en la consola el mensaje "x es abc"
		fmt.Println("x es abc")
	case "def":
		// Imprime en la consola el mensaje "x es def"
		fmt.Println("x es def")
	// Si 'x' no es igual a "abc" ni a "def"
	default:
		// Imprime en la consola el mensaje "x es otra cadena"
		fmt.Println("x es otra cadena")
	}
}

```

💪 Practica con estos ejercicios de condicionales en Go y envía la solución como un Pull Request a este repositorio:  [Go Para Principiantes](https://github.com/user/repo)

🎯  **10 Ejercicios de condicionales en Go**

1.  Imprima un mensaje diferente dependiendo de si un número es par o impar.
2.  Dados dos números enteros, imprima el mayor de los dos.
3.  Dados tres números enteros, imprima el mayor de los tres.
4.  Imprima un mensaje diferente si un número es positivo o negativo.
5.  Lea un número entero y determine si es divisible entre 5.
6.  Lea un número entero y determine si es divisible entre 3 y 5 al mismo tiempo.
7.  Lea dos números enteros y determine si el primero es mayor que el segundo.
8.  Lea tres números enteros y determine si el primero es igual al segundo y al tercero.
9.  Lea dos números enteros y determine si el primero es menor que el segundo o si son iguales.
10.  Lea un número entero y determine si es par o impar y si es positivo o negativo.

🎯  **10 Ejercicios de switch en Go**

1.  Imprima un mensaje diferente dependiendo de si un número está entre 0 y 5, entre 5 y 10, entre 10 y 20, o mayor que 20.
2.  Dados dos números enteros, imprima la letra correspondiente a la calificación apropiada.
3.  Dados tres números enteros, imprima la letra correspondiente a la calificación apropiada.
4.  Lea un número entero y determine la clase de número que es.
5.  Lea un número entero entre 0 y 5 y determine a qué letra corresponde.
6.  Lea un número entero entre 5 y 10 y determine a qué letra corresponde.
7.  Lea un número entero entre 10 y 20 y determine a qué letra corresponde.
8.  Lea un número entero mayor que 20 y determine a qué letra corresponde.
9.  Lea un número entero y determine si es par o impar y si es positivo o negativo.
10.  Lea un número entero y determine si es divisible entre 2, 3, 4, 5 o 6.

🔑  **Resumen en 3 puntos:**

1.  Los condicionales en Go (if,  else,  else if) permiten controlar el flujo de un programa basándose en condiciones.
2.  La sentencia  switch  en Go facilita la comparación de una variable con múltiples expresiones y ejecuta el bloque de código asociado a la expresión coincidente.
3.  Practicar con ejercicios de condicionales y  switch  en Go ayuda a mejorar tus habilidades en el control del flujo de tus programas y a comprender mejor cómo tomar decisiones basadas en condiciones y expresiones.

🎉 Ahora que has aprendido sobre condicionales y  switch  en Go, es hora de poner en práctica tus habilidades. Trabaja en los ejercicios mencionados y envía tus soluciones como Pull Requests al repositorio  [Go Para Principiantes](https://github.com/seagomezar/goparaprincipiantes). ¡Buena suerte y diviértete programando en Go! 🎉