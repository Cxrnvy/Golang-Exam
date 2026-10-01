package main

import (
	"fmt"
)

var nombresProductos []string
var subtotales []float64

func RegistrarVenta(nombre string, precio float64, cantidad int) {
	subtotal := precio * float64(cantidad)
	nombresProductos = append(nombresProductos, nombre)
	subtotales = append(subtotales, subtotal)
	fmt.Printf("Venta registrada: %d x %s (Subtotal: $%.2f)\n", cantidad, nombre, subtotal)
}
