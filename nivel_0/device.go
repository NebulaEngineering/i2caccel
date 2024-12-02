package accelemeter

import (
	"fmt"
	"log"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

func Device() {

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
	read := make([]byte, 1)

	// escritura
	if _, err := d.Write([]byte{CNTL1.Addr, 0x02}); err != nil {
		log.Fatalf("No se envia el mensaje Tx: %v", err)
	}

	// lectura
	if err := d.Tx([]byte{CNTL1.Addr}, read); err != nil {
		log.Fatalf("No se envia el mensaje Write : %v", err)
	}
	fmt.Printf("%2X\n", read)

}
