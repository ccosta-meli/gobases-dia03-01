/*
Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.

La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

Y los costos adicionales son:

- Pequeño: solo tiene el costo del producto
- Mediano: el precio del producto + un 3% de mantenerlo en la tienda
- Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

El porcentaje de mantenerlo en la tienda es en base al precio del producto.

El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.

Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga el método Precio.

Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo del producto y los adicionales en caso que los tenga
*/

package main

import "fmt"

type Producto interface {
	Precio() float64
}

type Pequeño struct {
	Costo float64
}

func (p *Pequeño) Precio() float64 {
	return p.Costo
}

type Mediano struct {
	Costo float64
}

func (p *Mediano) Precio() float64 {
	return p.Costo * 1.03
}

type Grande struct {
	Costo float64
}

func (p *Grande) Precio() float64 {
	return p.Costo*1.06 + 2500
}

func NewProducto(tipo string, costo float64) Producto {
	switch tipo {
	case "Pequeño":
		return &Pequeño{Costo: costo}
	case "Mediano":
		return &Mediano{Costo: costo}
	case "Grande":
		return &Grande{Costo: costo}
	default:
		return nil
	}
}

func main() {
	p1 := NewProducto("Pequeño", 1000)
	p2 := NewProducto("Mediano", 1000)
	p3 := NewProducto("Grande", 1000)

	fmt.Println("Precio Pequeño:", p1.Precio()) // 1000
	fmt.Println("Precio Mediano:", p2.Precio()) // 1030
	fmt.Println("Precio Grande:", p3.Precio())  // 3560
}
