package main

import (
	e1 "clase_2/modulo_ejemplo"
	"fmt"
)

// modulo personalizado

func main() {
	fmt.Println(e1.Ejemplo1())
	fmt.Println(e1.Ejemplo2("Axel"))
}

/*
func main() {
	fmt.Println("La hora actual es ", time.Now())
	fecha := time.Now()
	fmt.Println("El año es: ", fecha.Year())
	fmt.Println("El mes es: ", fecha.Month())
	fmt.Println("El año es: ", int(fecha.Month()))
	fmt.Println("El dia es: ", fecha.Day())
	fmt.Printf("El tipo es %T \n", fecha)
	fmt.Printf("%v/%v/%v \n", fecha.Day(), fecha.Month(), fecha.Year())
	minutochido := "0"
	minuto := fecha.Minute()
	min := strconv.Itoa(minuto)
	if minuto < 10 {
		min = minutochido + min
	}
	// hoy es lunes 23 de agosto del 2023 y son las 9:58 pm
	fmt.Printf("Hoy es %v %v de %v del %v y son las %v:%v pm", fecha.Weekday(), fecha.Day(), fecha.Month(), fecha.Year(), fecha.Hour(), min)
	fmt.Println("----------------Operaciones con fechas-------------------------")
	ahora := time.Now()
	fmt.Println("Fecha en este momento")
	fmt.Printf("%v/%v/%v \n", ahora.Day(), int(ahora.Month()), ahora.Year())
	fmt.Println("Mas 20 dias")
	fecha1 := ahora.Add(time.Hour * 24 * 20)
	fmt.Printf("%v/%v/%v \n", fecha1.Day(), int(fecha1.Month()), fecha1.Year())
	fecha2 := ahora.Add((time.Hour * 24) * -20)
	fmt.Printf("%v/%v/%v \n", fecha2.Day(), int(fecha2.Month()), fecha2.Year())

}
*/
