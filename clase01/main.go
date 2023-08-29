package main

import (
	"fmt"
)

// interfaces
func main() {
	e := Estructura{}
	fmt.Println(e.miFuncion())
	fmt.Println(e.calculo(3, 97))
}

type EjemploInterfaz interface {
	miFuncion() string
	calculo(n1 int, n2 int) int
}

type Estructura struct{}

func (*Estructura) miFuncion() string {
	return "texto desde mi funcion"
}

func (*Estructura) calculo(n1 int, n2 int) int {
	resultado := n1 + n2
	return resultado
}

/*
// EStructuras

func main() {
	estructura := Persona{
		Id:     1,
		Nombre: "Axel Neri",
		Correo: "axelneri81@aragon.unam.mx",
		Edad:   21,
	}

	fmt.Println(estructura)
	fmt.Printf("%+v \n", estructura) // aca se ve mas chido
	fmt.Println(reflect.TypeOf(estructura))
	fmt.Println(estructura.Nombre)
	fmt.Println("----------------------------------")
	p := new(Persona)
	fmt.Println(reflect.TypeOf(p))
	p.Id = 2
	p.Nombre = "Axel Neri lucero"
	p.Correo = "axelnerilucero@gmail.com"
	p.Edad = 21
	fmt.Println(p)
	fmt.Printf("%+v \n", p) // aca se ve mas chido
	fmt.Println(reflect.TypeOf(p))
	fmt.Println(p.Nombre)
	fmt.Println("----------------------------------")
	// estructura anidada / herencia

	categoria := Categoria{Id: 1, Nombre: "Categoria 1", Slug: "categoria-1"}
	producto := Producto{Id: 1, Nombre: "Mesa para computadora", Slug: "mesa-de-computador", Precio: 1234, CategoriaId: categoria}
	fmt.Printf("%+v \n", producto)

}

type Persona struct {
	Id     int
	Nombre string
	Correo string
	Edad   int
}

type Categoria struct {
	Id     int
	Nombre string
	Slug   string
}

type Producto struct {
	Id          int
	Nombre      string
	Slug        string
	Precio      int
	CategoriaId Categoria
}


// defer y panic
func main() {
	miFuncion()
}

func miFuncion() {
	// defer es despues de la ejecuion de la script
	defer fmt.Println("Este es el mensaje final")
	fmt.Println("Este es un primer mensaje")
	a := 1
	if a == 1 {
		panic("Fallo esta cosa perrin")
	}

}


// recursividad o recursion

func main() {
	miFuncion(0)
}

func miFuncion(valor int) {
	dato := valor + 1
	fmt.Println(dato)
	if dato < 10 {
		miFuncion(dato)
	}
}


// gorutinas

func main() {
	// Ejemplo 1
	fmt.Println(miFuncion("Axelin"))
	time.Sleep(time.Second * 5)
	fmt.Println(miFuncion("ya paso la pausa"))
	// Ejemplo 2
	miCanal := make(chan string)
	go func() {
		miCanal <- miFuncion("Juanitoalimaña")
	}()
	fmt.Println(<-miCanal)
	fmt.Println("Continuar con la ejecucion")
}

func miFuncion(parametro string) string {
	return "hola " + parametro
}


// funciones
func main() {
	miFuncion()
	fmt.Println("----------------------")
	miFuncionConParametros(2, 3)
	fmt.Println("----------------------")
	fmt.Println(miFuncionConRetorno("Axelin"))
	fmt.Println("----------------------")
	nombre, apellido, edad := miFuncionConRetornoMultiple()
	fmt.Printf("hola %v %v, tu edad es %v años \n", nombre, apellido, edad)
	fmt.Println("----------------------")
	fmt.Println("La suma es = ", suma(10, 12))
	fmt.Println("----------------------")
	Tabla := tabla(2)
	for i := 1; i <= 10; i++ {
		fmt.Printf("2 x %v = %v \n", i, Tabla())
	}
	fmt.Println("----------------------")
}

func miFuncion() {
	fmt.Println("se logro")
}

func miFuncionConParametros(n1 int, n2 int) {
	resultado := n1 + n2
	fmt.Println("La suma es ", resultado)
}

func miFuncionConRetorno(nombre string) string {
	return "Tu nombre es " + nombre
}

func miFuncionConRetornoMultiple() (string, string, int) {
	return "Axel", "Neri", 21
}

// funcion anonima

var suma = func(numero1 int, numero2 int) int {
	return numero1 + numero2
}

// funciones clousure
// funciones que regresan funciones

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}


// map
func main() {
	paises := make(map[string]int)
	paises["Argentina"] = 400000
	paises["Mexico"] = 120000000
	fmt.Println(paises)
	fmt.Println(reflect.TypeOf(paises))
	fmt.Println(paises["Mexico"])
	fmt.Println("---------------------")
	paises2 := map[int]string{
		1:  "chile",
		2:  "Mexico",
		11: "USA",
	}
	fmt.Println(paises2)
	fmt.Println(paises2[1])
	// veamos si existe algun valor en el map
	fmt.Println("--------------------")
	pais, existe := paises2[11]
	if existe {
		fmt.Println("Si existe el pais", pais)
	} else {
		fmt.Println("NO existe el pais")
	}
	// eliminar elemento de un pais
	delete(paises2, 1)
	fmt.Println("--------------------")
	fmt.Println(paises2)

	// recorrer un map con un ciclo for

	fmt.Println("------------------------")
	for id, valor := range paises2 {
		fmt.Printf("ID: %v | Nombre: %v \n", id, valor)
	}
	fmt.Println("------------------------")
	respuesta := map[string]string{
		"estado":  "ok",
		"mensaje": "cualquier mensaje",
	}
	fmt.Println(respuesta)
	fmt.Println("estado =", respuesta["estado"])
}



// Ciclos e iteraciones

func main() {
	i := 1
	for i < 11 {
		fmt.Println(i)
		i++
	}

	fmt.Println("-----")
	for i2 := 1; i2 < 11; i2++ {
		fmt.Println(i2)
	}
	for i3 := 0; i3 < 21; i3 += 5 {
		fmt.Println(i3)
	}
	for i4 := 1; i4 < 20; i4++ {
		if i4 == 18 {
			//break
			continue // se salta lo del if
		}
		fmt.Println(i4)
	}
	fmt.Println("--------")
	for i7 := 1; i7 < 11; i7++ {
		if i7%2 == 0 {
			fmt.Println(i7)
		}
	}
}


// condicionales

func main() {

		x == y | Es x igual a y?
		x != y | Es x diferente de y?
		x < y |Es x menor que y?
		x <= y | es x menor o igual que y?
		x > | Es x mayor que y?
		x >= y | Es x mayor o iguaal que y?


	edad := 21
	if edad >= 17 {
		fmt.Print("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}
	fmt.Println("---------------------")

	color := "blanco"
	if color == "rojo" {
		fmt.Println("Es rojo como la sangre")
	} else if color == "blanco" {
		fmt.Println("Es azul como el cielo")
	} else {
		fmt.Println("No tiene color patrio")
	}

	// operadores logicos && (and) || (or) ! (not)
	fmt.Println("---------------------")
	if color == "blanco" && edad == 21 {
		fmt.Println("Color es azul y edad 21")
	}

	// declarar variables dentro de una condicion
	fmt.Println("---------------------")
	if variable := 2; variable == 2 {
		fmt.Println("Variable es igual a 2")
	}

	colorchido := "rojo"
	// switch case
	switch colorchido {
	case "rojo":
		fmt.Println("-------")
		fmt.Println("es rojo")
	case "azul":
		fmt.Println("-------")
		fmt.Println("es azul")
	case "amarillo":
		fmt.Println("-------")
		fmt.Println("es amarillo")
	default:
		fmt.Println("-------")
		fmt.Println("no es complementario")
	}

}


// punteros
func main() {
	color := "rojo"
	fmt.Println(&color, color)
}


// reflect y type of
func main() {
	string1 := 32.02
	fmt.Println(reflect.TypeOf(string1))
}


// Tipos de datos
func main() {
	var string1 string = "texto"
	fmt.Println(string1)
	textoGrande := `Aprender un poco cada día marca la diferencia. Hay estudios que muestran
	que los estudiantes que hacen del aprendizaje un hábito tienen una mayor probabilidad de
	alcanzar sus objetivos. Reserva tiempo para aprender y recibe recordatorios con la
	herramienta de planificación del aprendizaje.`
	fmt.Println("---------------------")
	fmt.Println(textoGrande)
	var estado bool = false
	fmt.Println("---------------------")
	fmt.Println(estado)
	var flotante32 float32 = 33.33
	fmt.Println("---------------------")
	fmt.Println(flotante32)
	var flotante64 float64 = 64.33
	fmt.Println("---------------------")
	fmt.Println(flotante64)
	fmt.Println("---------------------")
	var entero int = 33
	fmt.Println("---------------------")
	fmt.Println(entero)
	fmt.Println("---------------------")
	var entero_int8 int8 = 123
	fmt.Println("---------------------")
	fmt.Println(entero_int8)
	var entero_int16 int16 = 123
	fmt.Println("---------------------")
	fmt.Println(entero_int16)
	var entero_int32 int32 = 45611
	fmt.Println("---------------------")
	fmt.Println(entero_int32)
	var entero_int64 int64 = 123
	fmt.Println("---------------------")
	fmt.Println(entero_int64)
	var entero_uint8 uint8 = 123
	fmt.Println("---------------------")
	fmt.Println(entero_uint8)

}

// Variabkes y constantes

const MiConstante = "Valor de mi constante"

func main() {
	// declaracion por inferencia
	var nombre string = "cesar"
	fmt.Println(nombre)
	// declaracion rapida o corta

	nombre2 := "Juan"
	fmt.Println(nombre2)
	fmt.Printf("El valor de MiConstante es %s \n", MiConstante)
}
*/
