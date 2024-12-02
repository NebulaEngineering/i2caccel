package accelemeter

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

// type Device struct {
// 	Name string
// 	dev  *i2c.Bus
// }

func Device_01() {

	// Inicializacion de drive de host
	if _, err := host.Init(); err != nil {
		log.Fatalf("failed to initialize periph: %v", err)
	}

	// Ocupar el bus i2c
	bus, err := i2creg.Open("2")
	if err != nil {
		log.Fatalf("error al carga bus i2c : %v", err)
	}
	defer bus.Close()

	d := &i2c.Dev{Addr: 0x1E, Bus: bus}

	if err := CNTL1.Write(d, 0x82); err != nil {
		log.Fatalf("No se envia el mensaje Tx: %v", err)
	}

	if read, err := CNTL1.Read(d); err != nil {
		log.Fatalf("No se envia el mensaje Write : %v", err)
	} else {
		fmt.Printf("t %2X\n", read)
	}
}
