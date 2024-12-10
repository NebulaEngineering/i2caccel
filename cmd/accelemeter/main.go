package main

import (
	"fmt"
	c0 "goi2caccel/control/kx023"
	d "goi2caccel/device"
)

// funcion que explora un Mapping para una lectura e interpretacion entre los registros
// y sus propiedades especificas.
func Interpretar() {
	for _, i := range c0.Seq_Register {
		var StrUp string = " "
		if i.Update {
			StrUp = "*"
		}

		fmt.Printf("%v%v (0x%02X): ", StrUp, i.Name, i.Value) // i.Read_error altera la interpretacion ojo
		for _, j := range c0.Mregister[i] {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

func main() {

	d.KX023.Init()
	defer d.KX023.Close()

	// lectura inicial de los registros
	for _, i := range c0.Seq_Register {
		i.Value = i.Read_Error()[0]
	}

	Interpretar()
	fmt.Println()

	c0.Config()

	fmt.Println()
	Interpretar()

}
