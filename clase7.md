**Programacion orientada a objetos en GO**

La programación orientada a objetos (POO) es un paradigma de programación que enfoca en la abstracción de los componentes de un programa para crear una relación entre ellos. Esto se logra usando objetos, y estos objetos pueden interactuar entre sí mediante mensajes. Estos objetos tienen atributos (propiedades) y comportamientos (métodos).

***La filosofía del diseño de Go***

Las raíces de Go están basadas en C y, más ampliamente, en la familia Algol. Ken Thompson medio en broma dijo que Rob Pike, Robert Granger y él mismo se reunieron y decidieron que odiaban C++. Sea una broma o no, Go es muy diferente de C++. Detalles sobre eso más adelante. Go se trata de la simplicidad máxima. Esto lo explica de manera detallada Rob Pike en Menos es exponencialmente más.

***Go frente a otros lenguajes***

Go no tiene clases, objetos, excepciones ni plantillas. Tiene recolección de basura y concurrencia integrada. La omisión más notable en lo que respecta a la orientación a objetos es que en Go no hay una jerarquía de tipos. Esto contrasta con la mayoría de los lenguajes orientados a objetos como C++, Java, C#, Scala, e incluso con los lenguajes dinámicos como Python y Ruby.

***Características del lenguaje orientado a objetos de Go***

Go no tiene clases, pero tiene tipos. En particular, tiene estructuras. Las estructuras son tipos definidos por el usuario. Los tipos de estructura (con métodos) sirven para fines similares a las clases en otros idiomas.

***Estructuras (Structs)***

Una estructura define el estado. Aquí hay una estructura Creature. Tiene un campo de nombre (Name) y una bandera booleana llamada Real, que nos dice si es una creature real o imaginaria. Las estructuras únicamente mantienen el estado y no el comportamiento.

En Go, la programación orientada a objetos se logra mediante el uso de structs para definir el estado de los objetos, y el uso de métodos para manipular el estado de los objetos. Los métodos también se pueden usar para definir comportamientos. Los métodos se definen como funciones, pero reciben un valor receiver que les permite acceder al estado de los objetos. Esto es útil para organizar el código y mejorar la reutilización.

Ejemplo 1:

``` go
type Car struct {
	Brand string
	Model string
	Speed int
}

func (c *Car) Accelerate() {
	c.Speed += 10
}

func main() {
	myCar := Car{
		Brand: "Ford",
		Model: "Mustang",
		Speed: 0,
	}

	myCar.Accelerate()
	fmt.Println(myCar.Speed) // 10
}
```

En el ejemplo anterior, hemos creado una estructura Car con los atributos Brand, Model y Speed. Hemos definido un método Accelerate para acelerar el coche. El método recibe un valor receiver que le permite acceder al estado del objeto Car. Finalmente, creamos una instancia de Car, llamada myCar, y llamamos al método Acelerate.

Ejemplo 2:

``` go
type Person struct {
	Name string
	Age  int
}

func (p *Person) Introduce() string {
	return fmt.Sprintf("Hi, my name is %s and I'm %d years old.", p.Name, p.Age)
}

func main() {
	john := Person{
		Name: "John",
		Age:  30,
	}

	fmt.Println(john.Introduce()) // Hi, my name is John and I'm 30 years old.
}
```

En este ejemplo, hemos creado una estructura Person con los atributos Name y Age. Hemos definido un método Introduce para presentarse. El método recibe un valor receiver que le permite acceder al estado del objeto Person. Finalmente, creamos una instancia de Person, llamada john, y llamamos al método Introduce.

***Métodos***

Los métodos son funciones que operan en tipos particulares. Tienen una cláusula de recepción que ordena el tipo de operación. Aquí hay un método Dump() que opera en estructuras Creature e imprime su estado:

``` go
type Creature struct {

  Name string

  Real bool

}

func (c Creature) Dump() {
  fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}
```

Esta es una sintaxis no usual, pero es muy explícita y clara (a diferencia del "this" implícito o del confuso "self" de Python).

***Incrustación***

La incrustación en Go se basa en el concepto de la herencia. La incrustación es una forma de herencia que permite a una estructura o clase heredar los atributos y métodos de otra. Esto significa que una estructura o clase puede heredar los atributos y métodos de otra, y también puede definir sus propios atributos y métodos. Esto le permite a la estructura o clase tener acceso a los atributos y métodos de la estructura o clase de la que hereda. Esto le permite a la estructura o clase comportarse como una instancia de la estructura o clase de la que hereda.

Un ejemplo de incrustación en Go es la herencia de una estructura llamada "Animal" por otra estructura llamada "Perro". La estructura "Animal" tendría atributos como "nombre" y "edad", mientras que la estructura "Perro" tendría los mismos atributos además de otros como "raza" y "tamaño". La estructura "Perro" heredaría los métodos de la estructura "Animal", como "hablar()" y "caminar()", y tendría sus propios métodos como "ladrar()" y "morder()". Esto significa que una instancia de la estructura "Perro" tendrá acceso a los métodos de la estructura "Animal" y también a sus propios métodos.

``` go
type Animal struct {
    Name string
    Age int
}

func (a Animal) Speak() {
    fmt.Println("Hola!")
}

func (a Animal) Walk() {
    fmt.Println("Caminando...")
}

type Dog struct {
    Animal
    Breed string
    Size int
}

func (d Dog) Bark() {
    fmt.Println("¡Guau!")
}

func (d Dog) Bite() {
    fmt.Println("Mordiendo...")
}

func main() {
    d := Dog{
        Animal: Animal{
            Name: "Fido",
            Age: 5,
        },
        Breed: "Dálmata",
        Size: 10,
    }
    d.Speak()
    d.Walk()
    d.Bark()
    d.Bite()
}
```

***Incrustacion Anonima***

Puedes incrustar tipos anónimos unos dentro de otros. Si incrustas una estructura sin nombre, luego la estructura incrustada proporciona su estado (y métodos) a la estructura incrustada directamente. Por ejemplo, FlyingCreature tiene una estructura Creature sin nombre incrustada, lo que significa que FlyingCreature es una Creature.

``` go
type FlyingCreature struct {
  Creature
  WingSpan int
}
```

Ahora, si tienes una instancia de una FlyingCreature, puedes acceder directamente a sus atributos Name y Real.

``` go
dragon := &FlyingCreature{
    Creature{"Dragon", false, },
    15,
}

fmt.Println(dragon.Name)
fmt.Println(dragon.Real)
fmt.Println(dragon.WingSpan)
```

***Interfaces***

Las interfaces son el sello distintivo del soporte orientado a objetos de Go. Las interfaces son tipos que declaran conjuntos de métodos. Del mismo modo que las interfaces en otros lenguajes, no tienen implementación.


Las interfaces en Go son una forma de definir comportamiento de un objeto. Esto significa que un objeto que implementa una interfaz debe implementar ciertos comportamientos. Estos comportamientos se definen a través de la interfaz, y el objeto que implementa la interfaz debe implementar todos ellos para satisfacer la interfaz.

Un interfaz en Go se define como un grupo de métodos. Estos métodos no tienen una implementación, solamente se define el firma de los métodos.

A continuación se muestra una interfaz definida en Go llamada Animal:

``` go
type Animal interface {
    MakeNoise() string
    Move() string
}
```

Esta interfaz define dos métodos: MakeNoise() y Move(). Estos métodos no tienen una implementación, solamente se declaran.

Ahora para implementar esta interfaz en un objeto, necesitamos definir un tipo de objeto y luego implementar los métodos definidos en la interfaz. Por ejemplo, podemos definir un objeto llamado Dog:

``` go
type Dog struct {
    breed string
    age int
}
```

Esta definición del objeto Dog contiene dos campos, breed y age. Ahora para implementar la interfaz Animal, necesitamos implementar los métodos definidos en la interfaz. Esto se hace mediante la definición de los métodos con el objeto Dog.

``` go
func (d Dog) MakeNoise() string {
    return "Woof!"
}

func (d Dog) Move() string {
    return "Run!"
}
```

Ahora Dog ha implementado los métodos definidos en la interfaz Animal, por lo que Dog implementa la interfaz Animal. Esto significa que cualquier objeto Dog puede ser usado como un objeto Animal.

Por ejemplo, podemos crear un objeto Dog y luego pasarlo como un objeto Animal a una función:

``` go
func doSomething(a Animal) {
    fmt.Println(a.MakeNoise())
    fmt.Println(a.Move())
}

func main() {
    d := Dog{breed: "Labrador", age: 5}
    doSomething(d)
}
```

En este ejemplo, hemos pasado un objeto Dog a la función doSomething(), que espera un objeto Animal como parámetro. Esto funciona porque Dog implementa la interfaz Animal, por lo que se puede usar un objeto Dog como un objeto Animal.

veamos otro ejemplo y algunas convenciones

Los objetos que implementan todos los métodos de interfaz implementan automáticamente la interfaz. No hay herencia o subclasificación o palabra clave "implements". En el siguiente fragmento de código, el tipo Foo implementa la interfaz Fooer (por convención, los nombres de la interfaz Go terminan con "er").

``` go
type Fooer interface {
  Foo1()
  Foo2()
  Foo3()
}

type Foo struct {
}

func (f Foo) Foo1() {
    fmt.Println("Foo1() here")
}

func (f Foo) Foo2() {
    fmt.Println("Foo2() here")
}

func (f Foo) Foo3() {
    fmt.Println("Foo3() here")
}
```

Diseño orientado a objetos: la forma de Go

Vamos a ver cómo se compara Go con los pilares de la programación orientada a objetos: encapsulación, herencia y polimorfismo. Esas son las características de los lenguajes de programación basados en clases, que son los lenguajes de programación orientados a objetos más populares

En lo fundamental, los objetos son construcciones de lenguaje que tienen un estado y un comportamiento que opera sobre el estado y lo expone selectivamente a otras partes del programa.

***Encapsulado***

Go encapsula las cosas a nivel de paquete. Los nombres que empiezan con una letra minúscula únicamente son visibles dentro de ese paquete. Puedes ocultar cualquier cosa en un paquete privado y solamente exponer tipos específicos, interfaces y funciones de fábrica.

Por ejemplo, para ocultar aquí el tipo Foo anterior y exponer únicamente la interfaz, puedes cambiarle el nombre a foo en minúsculas y proporcionar una función NewFoo() que devuelve la interfaz pública Fooer:

``` go
type foo struct {
}

func (f foo) Foo1() {
    fmt.Println("Foo1() here")
}

func (f foo) Foo2() {
    fmt.Println("Foo2() here")
}

func (f foo) Foo3() {
    fmt.Println("Foo3() here")
}

func NewFoo() Fooer {
    return &Foo{}
}
```

Entonces el código de otro paquete puede utilizar NewFoo() y obtener acceso a una interfaz Fooer implementada por el tipo de foo interno:

``` go
f := NewFoo()

f.Foo1()

f.Foo2()

f.Foo3()
```

***Herencia***

La herencia o subclases siempre fue un tema controversial. Hay muchos problemas con la herencia de la implementación (a diferencia de la herencia de la interfaz). La herencia múltiple, como se ha implementado por C++, Python y otros lenguajes, sufre el problema del diamante, pero incluso la herencia única no es tarea fácil con el problema frágil de la clase base.

Los lenguajes modernos y el pensamiento orientado a objetos ahora favorecen la composición sobre la herencia. Go se lo toma en serio y no tiene ningún tipo de jerarquía. Te permite compartir detalles de implementación a través de la composición. Pero Go, en un giro muy extraño (que posiblemente se originó por preocupaciones pragmáticas), permite la composición anónima mediante la incrustación.

En efecto, la composición mediante la incrustación de un tipo anónimo es equivalente a la herencia de la implementación. Una estructura incrustada es tan frágil como una clase base. También puedes incrustar una interfaz, lo que equivale a heredar de una interfaz en lenguajes como Java o C++. Incluso puede provocar un error de tiempo de ejecución que no se descubra en el momento de la compilación si el tipo de incrustación no implementa todos los métodos de interfaz.

Las interfaces en Go son una forma de definir comportamiento de un objeto. Esto significa que un objeto que implementa una interfaz debe implementar ciertos comportamientos. Estos comportamientos se definen a través de la interfaz, y el objeto que implementa la interfaz debe implementar todos ellos para satisfacer la interfaz.

Un interfaz en Go se define como un grupo de métodos. Estos métodos no tienen una implementación, solamente se define el firma de los métodos.

A continuación se muestra una interfaz definida en Go llamada Animal:

``` go
type Animal interface {
    MakeNoise() string
    Move() string
}
```

Esta interfaz define dos métodos: MakeNoise() y Move(). Estos métodos no tienen una implementación, solamente se declaran.

Ahora para implementar esta interfaz en un objeto, necesitamos definir un tipo de objeto y luego implementar los métodos definidos en la interfaz. Por ejemplo, podemos definir un objeto llamado Dog:

``` go
type Dog struct {
    breed string
    age int
}
```

Esta definición del objeto Dog contiene dos campos, breed y age. Ahora para implementar la interfaz Animal, necesitamos implementar los métodos definidos en la interfaz. Esto se hace mediante la definición de los métodos con el objeto Dog.

``` go
func (d Dog) MakeNoise() string {
    return "Woof!"
}

func (d Dog) Move() string {
    return "Run!"
}
```

Ahora Dog ha implementado los métodos definidos en la interfaz Animal, por lo que Dog implementa la interfaz Animal. Esto significa que cualquier objeto Dog puede ser usado como un objeto Animal.

Por ejemplo, podemos crear un objeto Dog y luego pasarlo como un objeto Animal a una función:

``` go
func doSomething(a Animal) {
    fmt.Println(a.MakeNoise())
    fmt.Println(a.Move())
}

func main() {
    d := Dog{breed: "Labrador", age: 5}
    doSomething(d)
}
```

En este ejemplo, hemos pasado un objeto Dog a la función doSomething(), que espera un objeto Animal como parámetro. Esto funciona porque Dog implementa la interfaz Animal, por lo que se puede usar un objeto Dog como un objeto Animal.

veamos otro ejemplo

Aquí SuperFoo incrusta la interfaz Fooer, pero no implementa sus métodos. El compilador de Go encantadamente te permitirá crear un nuevo SuperFoo y llamar a los métodos Fooer, pero obviamente fallará en tiempo de ejecución. Esto compila:

``` go
type SuperFooer struct {
  Fooer
}

func main() {
  s := SuperFooer{}
  s.Foo2()
```

Ejecutar este código nos mostrará el siguiente error debido a que se violaron las reglas de encapsulamiento

``` bash
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xb code=0x1 addr=0x28 pc=0x2a78]

goroutine 1 [running]:
panic(0xde180, 0xc82000a0d0)
  /usr/local/Cellar/go/1.6/libexec/src/runtime/panic.go:464 +0x3e6
main.main()
  /Users/gigi/Documents/dev/go/src/github.com/oop_test/main.go:104 +0x48
exit status 2

Process finished with exit code 1
```

***Polimorfismo***

El polimorfismo es la esencia de la programación orientada a objetos: la capacidad de tratar objetos de diferentes tipos de forma uniforme siempre que se adhieran a la misma interfaz. Las interfaces de Go brindan esta capacidad de una manera muy directa e intuitiva.

Aquí tienes un ejemplo elaborado en el que se crean múltiples creatures (¡y una door!) que implementan la interfaz Dumper se crean y se almacenan en una fracción y luego se llama al método Dump() para cada una. También notarás diferentes estilos de instanciar los objetos.

Veamos el codigo completo

``` go
package main

import "fmt"

type Creature struct {
  Name string
  Real bool
}

func Dump(c*Creature) {
  fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

func (c Creature) Dump() {
  fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)
}

type FlyingCreature struct {
  Creature
  WingSpan int
}

func (fc FlyingCreature) Dump() {
  fmt.Printf("Name: '%s', Real: %t, WingSpan: %d\n",
    fc.Name,
    fc.Real,
    fc.WingSpan)
}

type Unicorn struct {
  Creature
}

type Dragon struct {
  FlyingCreature
}

type Pterodactyl struct {
  FlyingCreature
}

func NewPterodactyl(wingSpan int) *Pterodactyl {
  pet := &Pterodactyl{
    FlyingCreature{
      Creature{
        "Pterodactyl",
        true,
      },
      wingSpan,
    },
  }
  return pet
}

type Dumper interface {
  Dump()
}

type Door struct {
  Thickness int
  Color     string
}

func (d Door) Dump() {
  fmt.Printf("Door => Thickness: %d, Color: %s", d.Thickness, d.Color)
}

func main() {
  creature := &Creature{
    "some creature",
    false,
  }

  uni := Unicorn{
    Creature{
      "Unicorn",
      false,
    },
  }

  pet1 := &Pterodactyl{
    FlyingCreature{
      Creature{
        "Pterodactyl",
        true,
      },
      5,
    },
  }

  pet2 := NewPterodactyl(8)

  door := &Door{3, "red"}

  Dump(creature)
  creature.Dump()
  uni.Dump()
  pet1.Dump()
  pet2.Dump()

  creatures := []Creature{
    *creature,
    uni.Creature,
    pet1.Creature,
    pet2.Creature}
  fmt.Println("Dump() through Creature embedded type")
  for _, creature := range creatures {
    creature.Dump()
  }

  dumpers := []Dumper{creature, uni, pet1, pet2, door}
  fmt.Println("Dump() through Dumper interface")
  for _, dumper := range dumpers {
    dumper.Dump()
  }
}
```

**Conclusión**

Go es un verdadero lenguaje de programación orientado a objetos. Permite el modelado basado en objetos y promueve la mejor práctica para utilizar interfaces en vez de jerarquías de tipos concretos. Go tomó algunas decisiones sintácticas atípicas, pero en general, trabajar con tipos, métodos e interfaces se siente simple, ligero y natural.

La herencia es uno de los principales conceptos relacionados con la programación orientada a objetos, y Go lo soporta a través de la sintaxis de la composición. Esto significa que los tipos de Go pueden tener uno o más campos asociados a otros tipos. Al mismo tiempo, los métodos asociados a un tipo pueden llamar a los métodos de los tipos asociados. Esto permite una forma de herencia más simple, sin la necesidad de definir clases.

La incrustación es una forma ligeramente diferente de herencia. En Go, un tipo puede incluir otros tipos dentro de su definición. Esto permite aprovechar el comportamiento de los tipos incrustados, al tiempo que permite a los tipos tener su propio comportamiento adicional. La incrustación es una herramienta poderosa que ayuda a simplificar la programación orientada a objetos.

Las interfaces son una de las principales características de Go. Las interfaces permiten a los programadores definir un contrato para los tipos que implementan la interfaz. Esto proporciona una forma estandarizada para tratar con diferentes tipos de objetos de la misma forma. Las interfaces presentan una forma clara de definir la funcionalidad de los objetos sin necesidad de tener una jerarquía de tipos explícita.

En resumen, Go es un lenguaje de programación orientado a objetos que soporta la herencia, la incrustación y las interfaces. La herencia permite la composición de tipos y métodos, la incrustación permite la reutilización de comportamiento y las interfaces definen un contrato para los tipos que implementan la interfaz. Estas características hacen de Go un lenguaje de programación orientado a objetos moderno y poderoso.


A continuacion te dejo 10 ejercicios mas para que practiques las estructuras en Go y envies la solucion como un Pull requeste a este repositorio:

1. Crea una interfaz en Go llamada Shape con métodos Area() y Perimeter().
2. Cree dos clases en Go, Rectangle y Square, que implementen la interfaz Shape.
3. Agregue algunos atributos, como width y height, a la clase Rectangle para calcular el área y el perímetro.
4. Agregue un atributo, side, a la clase Square para calcular el área y el perímetro.
5. Escriba un método en la clase Rectangle que calcule el área.
6. Escriba un método en la clase Rectangle que calcule el perímetro.
7. Escriba un método en la clase Square que calcule el área.
8. Escriba un método en la clase Square que calcule el perímetro.
9. Cree una instancia de cada clase Rectangle y Square.
10. Llame a los métodos Area() y Perimeter() para cada instancia, y verifique si los resultados son correctos.