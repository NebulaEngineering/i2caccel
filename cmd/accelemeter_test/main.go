package main

import (
	accelemeter "cli_acc/main/nivel_0"
	"fmt"
	"log"
	"time"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

type Read_register interface {
	Read(d *i2c.Dev) ([]byte, error)
}

func Error_register(r Read_register, d *i2c.Dev) []byte {
	if read, err := r.Read(d); err != nil {
		log.Fatalf("No se envia el mensaje Write : %v", err)
		return nil
	} else {
		return read
	}
}

func main() {
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

	var data Read_register

	eje_X = &accelemeter.XHP_H
	numero := Error_register(data, d)
	fmt.Print(numero)

	for true {

		// XOUT:= Error_register(accelemeter.YHP_H,d)

		if read, err := accelemeter.YHP_H.Read(d); err != nil {
			log.Fatalf("No se envia el mensaje Write : %v", err)
		} else {
			fmt.Print(read[0])
		}

		if read, err := accelemeter.YHP_L.Read(d); err != nil {
			log.Fatalf("No se envia el mensaje Write : %v", err)
		} else {
			fmt.Print(read[0], " ")
		}

		if read, err := accelemeter.ZHP_H.Read(d); err != nil {
			log.Fatalf("No se envia el mensaje Write : %v", err)
		} else {
			fmt.Print(read[0])
		}

		if read, err := accelemeter.ZHP_L.Read(d); err != nil {
			log.Fatalf("No se envia el mensaje Write : %v", err)
		} else {
			fmt.Println(read[0])
		}

		time.Sleep(400 * time.Millisecond)
	}
}
