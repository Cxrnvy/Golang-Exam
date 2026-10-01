package main

import "fmt"

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

	fmt.Println("\nEstadísticas de Ventas")
	var totalRecaudado float64

	for i := 0; i < len(nombresProductos); i++ {
		fmt.Printf("Producto: %s | Subtotal: $%.2f\n", nombresProductos[i], subtotales[i])
		totalRecaudado += subtotales[i]
	}

	fmt.Printf("Total recaudado: $%.2f\n", totalRecaudado)
}

func main() {
	var opcion int

	for {
		fmt.Println("\n Menú Principal ")
		fmt.Println("1: Registrar una nueva venta")
		fmt.Println("2: Mostrar estadísticas")
		fmt.Println("3: Salir")
		fmt.Print("Seleccione una opción: ")

		_, err := fmt.Scanln(&opcion)
		if err != nil {
			fmt.Println("Entrada inválida. Por favor, ingrese un número.")
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		switch opcion {
		case 1:
			registrarNuevaVenta()
		case 2:
			MostrarEstadisticas()
		case 3:
			fmt.Println("\nSaliendo del programa...")
			return
		default:
			fmt.Println("\nOpción no válida. Intente de nuevo.")
		}
	}
}

func registrarNuevaVenta() {
	fmt.Println("\nProductos")
	fmt.Println("1: Arroz  = $1.25")
	fmt.Println("2: Leche  = $0.95")
	fmt.Println("3: Pan    = $0.50")
	fmt.Print("Seleccione el número del producto: ")

	var opcionProducto int
	fmt.Scanln(&opcionProducto)

	var nombre string
	var precio float64

	switch opcionProducto {
	case 1:
		nombre = "Arroz"
		precio = 1.25
	case 2:
		nombre = "Leche"
		precio = 0.95
	case 3:
		nombre = "Pan"
		precio = 0.50
	default:
		fmt.Println("Producto no válido - Venta Cancelada")
		return
	}

	fmt.Print("Ingrese la cantidad vendida: ")
	var cantidad int
	fmt.Scanln(&cantidad)

	if cantidad <= 0 {
		fmt.Println("La cantidad debe ser mayor a 0 - Venta Cancelada")
		return
	}

	RegistrarVenta(nombre, precio, cantidad)
}
