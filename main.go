package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	var precioCompra string = os.Args[1]
	var ganancia string = os.Args[2]

	precioCompraFloat, _ := strconv.ParseFloat(precioCompra, 32)
	gananciaFloat, _ := strconv.ParseFloat(ganancia, 32)

	gananciaPorcentual2 := gananciaFloat / 200000
	gananciaPorcentual3 := gananciaFloat / 300000
	gananciaPorcentual4 := gananciaFloat / 400000
	gananciaPorcentual5 := gananciaFloat / 500000
	gananciaPorcentual1 := gananciaFloat / 1000000

	precioVenta2 := (1 + gananciaPorcentual2) / math.Pow((1-0.0035), 2) * precioCompraFloat
	precioVenta3 := (1 + gananciaPorcentual3) / math.Pow((1-0.0035), 2) * precioCompraFloat
	precioVenta4 := (1 + gananciaPorcentual4) / math.Pow((1-0.0035), 2) * precioCompraFloat
	precioVenta5 := (1 + gananciaPorcentual5) / math.Pow((1-0.0035), 2) * precioCompraFloat
	precioVenta1 := (1 + gananciaPorcentual1) / math.Pow((1-0.0035), 2) * precioCompraFloat

	fmt.Println("----------------------------------")
	fmt.Printf("PRECIO VENTA PARA GANAR %.f CLP\n", gananciaFloat)
	fmt.Printf("PRECIO COMPRA CRIPTO  %.f CLP\n", precioCompraFloat)
	fmt.Println("----------------------------------")
	fmt.Println("Monto invertido    Precio Venta")
	fmt.Println("----------------------------------")
	fmt.Printf("   200000 CLP       %.f CLP\n", precioVenta2)
	fmt.Printf("   300000 CLP       %.f CLP\n", precioVenta3)
	fmt.Printf("   400000 CLP       %.f CLP\n", precioVenta4)
	fmt.Printf("   500000 CLP       %.f CLP\n", precioVenta5)
	fmt.Printf("  1000000 CLP       %.f CLP\n", precioVenta1)
}
