package main

import (
	"flag"
	"fmt"
	c0 "i2caccel/control/kx023"
	d "i2caccel/device"
	v0 "i2caccel/view/kx023"
	"time"
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
	tick := time.NewTicker(10 * time.Millisecond)
	defer tick.Stop()
	for range tick.C {
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
		mode       = flag.String("mode", "view", "escoje el modo de aplicacion de sensor (run|view|config)")
		LPRO       = flag.Bool(c0.LPR0.Name, false, "Alternar el valor de la frecuencia de corte del fitro pasa baja")
		IIR_BYPASS = flag.Bool(c0.IIR_BYPASS.Name, false, "Alternar la desactivacion de filtro pasa baja de la salida normal.")
		OSA        = flag.Float64(c0.OSA.Name, 12.5, "Alterna el valor de la frecuencia de salida de datos en la salida normal de accelerometro.\n (0.781|1.563|3.125|6.25|12.5|25|50|100|200|400|800|1600).")
		OWUF       = flag.Float64(c0.OWUF.Name, 0.781, "Alterna la frecuencia de salida de datos general de la deccion  de movimiento y salida con filtro pasa altas.\n (0.781|1.563|3.125|6.25|12.5|25|50|100)")
		AVC        = flag.Int(c0.AVC.Name, 16, " Valor del promedio de los datos de salida (0|2|4|8|16|32|64|128)")
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

	fmt.Println(LPRO, IIR_BYPASS, OSA, OWUF, AVC)

	// Interpretar()
	// fmt.Println()

	// c0.Config()

	// fmt.Println()
	// Interpretar()
}
