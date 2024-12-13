package main

import (
	"flag"
	"fmt"
	c0 "i2caccel/control/kx023"
	d "i2caccel/device"
	v0 "i2caccel/view/kx023"
)

// funcion que explora un Mapping para una lectura e interpretacion entre los registros
// y sus propiedades especificas.
func Interpretar() {

	// lectura de los registros
	for _, reg := range c0.Seq_Register {
		reg.Value = reg.Read_Error()[0]
	}

	for _, reg := range c0.Seq_Register {
		var Update string = " "
		if reg.Update {
			Update = "*"
		}

		fmt.Printf("%v%v (0x%02X): ", Update, reg.Name, reg.Value) // i.Read_error altera la interpretacion ojo
		for _, cntl := range c0.Mregister[reg] {
			fmt.Print(cntl)
		}
		fmt.Println()
	}
}

func Run() {
	for {
		if d.KX023.ReadGPIO_Error(0) {
			// esta pendiente de la interrupcion
			fmt.Println(&v0.XOUT, &v0.YOUT, &v0.ZOUT)
		}
	}
}

func main() {

	fmt.Print("\033[H\033[2J")

	d.KX023.Init()
	defer d.KX023.Close()

	var (
		mode = flag.String("mode", "view", "escoje el modo de aplicacion de sensor (run|view|config)")
	)

	// Analizar las banderas
	flag.Parse()

	switch *mode {
	case "run":
		Run()
	case "view":
		Interpretar()
	case "config":
		c0.Config()
	}

	// Interpretar()
	// fmt.Println()

	// c0.Config()

	// fmt.Println()
	// Interpretar()
}
