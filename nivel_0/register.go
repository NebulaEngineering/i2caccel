package accelemeter

import (
	"periph.io/x/conn/v3/i2c"
)

type Register struct {
	Name   string
	Addr   byte
	Value  byte
	Update bool
}

func (r *Register) Write(d *i2c.Dev, n byte) error {
	if _, err := d.Write([]byte{CNTL1.Addr, n}); err != nil {
		return err
	}
	return nil
}

func (r *Register) Read(d *i2c.Dev) ([]byte, error) {
	read := make([]byte, 1)
	if err := d.Tx([]byte{CNTL1.Addr}, read); err != nil {
		return nil, err
	}
	return read, nil
}
