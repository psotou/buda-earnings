package main

import (
	"flag"
	"fmt"
	"math"
)

var (
	gananciaFloat     *float64
	precioCompraFloat *float64
)

func main() {
	precioCompraFloat = flag.Float64("c", 1, "Precio de compra")
	gananciaFloat = flag.Float64("g", 1, "Retorno esperado")
	flag.Parse()

	gananciaPorcentual5 := *gananciaFloat / 500000
	gananciaPorcentual1 := *gananciaFloat / 1000000
	gananciaPorcentual2M := *gananciaFloat / 2000000
	gananciaPorcentual3M := *gananciaFloat / 3000000
	gananciaPorcentual4M := *gananciaFloat / 4000000
	gananciaPorcentual5M := *gananciaFloat / 5000000
	gananciaPorcentual6M := *gananciaFloat / 6000000

	precioVenta5 := (1 + gananciaPorcentual5) / math.Pow((1-0.0035), 2) * *precioCompraFloat
	precioVenta1 := (1 + gananciaPorcentual1) / math.Pow((1-0.0035), 2) * *precioCompraFloat
	precioVenta2M := (1 + gananciaPorcentual2M) / math.Pow((1-0.0035), 2) * *precioCompraFloat
	precioVenta3M := (1 + gananciaPorcentual3M) / math.Pow((1-0.0035), 2) * *precioCompraFloat
	precioVenta4M := (1 + gananciaPorcentual4M) / math.Pow((1-0.0035), 2) * *precioCompraFloat
	precioVenta5M := (1 + gananciaPorcentual5M) / math.Pow((1-0.0035), 2) * *precioCompraFloat
	precioVenta6M := (1 + gananciaPorcentual6M) / math.Pow((1-0.0035), 2) * *precioCompraFloat

	fmt.Println("---------------------------------------------------------------------")
	fmt.Printf("%30s %.f CLP según el monto invertido\n", "Precio de venta para ganar", *gananciaFloat)
	fmt.Printf("%37s %.f CLP\n", "Precio de compra", *precioCompraFloat)
	fmt.Println("---------------------------------------------------------------------")
	fmt.Printf("%35s %s %s\n", "Monto invertido ", " ", "Precio Venta")
	fmt.Printf("%52s \n", "-----------------------------------")
	fmt.Printf("%32s %5s %.f CLP\n", "500000 CLP", " ", precioVenta5)
	fmt.Printf("%32s %5s %.f CLP\n", "1000000 CLP", " ", precioVenta1)
	fmt.Printf("%32s %5s %.f CLP\n", "2000000 CLP", " ", precioVenta2M)
	fmt.Printf("%32s %5s %.f CLP\n", "3000000 CLP", " ", precioVenta3M)
	fmt.Printf("%32s %5s %.f CLP\n", "4000000 CLP", " ", precioVenta4M)
	fmt.Printf("%32s %5s %.f CLP\n", "5000000 CLP", " ", precioVenta5M)
	fmt.Printf("%32s %5s %.f CLP\n", "6000000 CLP", " ", precioVenta6M)
	fmt.Println("---------------------------------------------------------------------")
}
