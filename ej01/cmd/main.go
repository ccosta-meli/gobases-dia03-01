/*
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

Nombre: [Nombre del alumno]

Apellido: [Apellido del alumno]

DNI: [DNI del alumno]

Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.

Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

package main

import "fmt"

type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

type AlumnoInterface interface {
	imprimirAlumnos()
}

func (a *Alumno) imprimirAlumnos() {
	fmt.Println(a.Nombre)
	fmt.Println(a.Apellido)
	fmt.Println(a.DNI)
	fmt.Println(a.Fecha)
}

func details(a AlumnoInterface) {
	a.imprimirAlumnos()
}

func main() {
	alumno := Alumno{"Cristian", "Costa", 38056171, "26-01-1994"}
	details(&alumno)
}
