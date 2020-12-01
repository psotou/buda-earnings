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
	ganaciaFloat, _ := strconv.ParseFloat(ganancia, 32)

	gananciaPorcentual2 := ganaciaFloat / 200000
	gananciaPorcentual3 := ganaciaFloat / 300000
	gananciaPorcentual5 := ganaciaFloat / 500000

	precioVenta2 := (1 + gananciaPorcentual2) / math.Pow((1-0.0035), 2) * precioCompraFloat
	precioVenta3 := (1 + gananciaPorcentual3) / math.Pow((1-0.0035), 2) * precioCompraFloat
	precioVenta5 := (1 + gananciaPorcentual5) / math.Pow((1-0.0035), 2) * precioCompraFloat

	fmt.Println("----------------------------------")
	fmt.Printf("PRECIO VENTA PARA GANAR %.f CLP\n", ganaciaFloat)
	fmt.Printf("PRECIO COMPRA CRIPTO  %.f CLP\n", precioCompraFloat)
	fmt.Println("----------------------------------")
	fmt.Println("Monto invertido    Precio Venta")
	fmt.Println("----------------------------------")
	fmt.Printf("   200000 CLP       %.f CLP\n", precioVenta2)
	fmt.Printf("   300000 CLP       %.f CLP\n", precioVenta3)
	fmt.Printf("   500000 CLP       %.f CLP\n", precioVenta5)
}
