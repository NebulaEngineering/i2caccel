package main

import (
	"flag"
	"fmt"
	c0 "i2caccel/control/kx023"
	d "i2caccel/device"
	v0 "i2caccel/view/kx023"
	"os"
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

	var cadena string

	file, err := os.OpenFile("dataKX023.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error al abrir o crear el archivo:", err)
		return
	}
	defer file.Close() // Cerrar el archivo al finalizar

	tick := time.NewTicker(10 * time.Millisecond)
	defer tick.Stop()
	for range tick.C {
		// esta pendiente de la interrupcion
		if d.KX023.ReadGPIO_Error(0) {

			cadena = fmt.Sprintln(&v0.XOUT, &v0.YOUT, &v0.ZOUT)

			fmt.Print(cadena)
			_, err = file.WriteString(cadena)
			if err != nil {
				fmt.Println("Error al escribir en el archivo:", err)
				return
			}
		}
	}
}

func main() {

	fmt.Print("\033[H\033[2J")

	d.KX023.Init()
	defer d.KX023.Close()

	var (
		flagMODE       = flag.String("MODE", "view", "(modo) Escoje el modo de aplicacion de sensor. (run|view|config)")
		flagLPRO       = flag.Int("LPRO", 0, "(propiedad) Alternar el valor de la frecuencia de corte del filtro pasa baja. (0:=OSA/9|1:=OSA/2) (default 0)")
		flagIIR_BYPASS = flag.Int("IIR_BYPASS", 0, "(propiedad) Alternar la desactivacion de filtro pasa baja de la salida normal. (0:=active|1:=inactive) (default 0)")
		flagOSA        = flag.Float64("OSA", 12.5, "(propiedad) Alterna el valor de la frecuencia de salida de datos en la salida normal de accelerometro.\n (0.781|1.563|3.125|6.25|12.5|25|50|100|200|400|800|1600).")
		flagOWUF       = flag.Float64("OWUF", 0.781, "(propiedad) Alterna la frecuencia de salida de datos general de la detección del movimiento y salida con filtro pasa altas.\n (0.781|1.563|3.125|6.25|12.5|25|50|100)")
		flagAVC        = flag.Int("AVC", 16, "(propiedad) Valor del promedio de los datos de salida. (0|2|4|8|16|32|64|128)")
	)

	// Analizar las banderas
	flag.Parse()
	switch *flagMODE {
	case "run":
		Run()
	case "view":
		Interpretar()
	case "config":
		c0.Config(flagLPRO, flagIIR_BYPASS, flagOSA, flagOWUF, flagAVC)
	}
}
