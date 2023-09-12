package main

import "fmt"

func main() {
	var v int = 19
	fmt.Println("El valor de v es de:", v)
	Increase(&v)
	fmt.Println("El valor de v luego del increase es de:", v)
	fmt.Println("La direccion de memoria de v es:", &v)
}

func Increase(v *int) {
	*v++
}
