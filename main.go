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

func MostrarEstadisticas() {
	if len(subtotales) == 0 {
		fmt.Println("\nNo existen ventas registradas en este momento.")
		return
	}

	fmt.Println("\n--- Estadísticas de Ventas ---")
	var totalRecaudado float64

	for i := 0; i < len(nombresProductos); i++ {
		fmt.Printf("Producto: %s | Subtotal: $%.2f\n", nombresProductos[i], subtotales[i])
		totalRecaudado += subtotales[i]
	}

	fmt.Printf("------------------------------\n")
	fmt.Printf("Total recaudado: $%.2f\n", totalRecaudado)

}

	for {
		fmt.Println("\n--- Menú Principal ---")
		fmt.Println("1. Registrar una nueva venta")
		fmt.Println("2. Mostrar estadísticas")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")

		_, err := fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, ingrese un número.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}
		