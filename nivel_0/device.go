package nivel0

import (
	"log"
	"sync"

	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

type device struct {
	name  string
	adrr  uint16
	bus   string
	iobus i2c.BusCloser
	dev   *i2c.Dev
	mux   sync.Mutex
}

// Inicializar periperico y controladores I2C
func (d *device) Init() {
	// Inicializacion de drive de host
	if _, err := host.Init(); err != nil {
		log.Fatalf("failed to initialize periph: %v", err)
	}

	// Ocupar el iobus i2c
	iobus, err := i2creg.Open(d.bus)
	if err != nil {
		log.Fatalf("error al carga bus i2c : %v", err)
	}

	d.dev = &i2c.Dev{Addr: d.adrr, Bus: iobus}
}

func (d *device) Close() error {
	return d.iobus.Close()
}

// Escritura I2C
// Escritura sin manejo de error
func (d *device) Write(w []byte) error {
	d.mux.Lock()
	defer d.mux.Unlock()
	if _, err := d.dev.Write(w); err != nil {
		return err
	}
	return nil
}

func (d *device) Write_Error(w []byte) {
	if err := d.Write(w); err != nil {
		log.Fatalf("Error es del registro 0x%2X, Error: %v", w, err)
	}
}

// Lectura I2C
// Lectura sin manejo de error
func (d *device) Read(w []byte) ([]byte, error) {
	d.mux.Lock()
	defer d.mux.Unlock()
	read := make([]byte, 1)
	if err := d.dev.Tx(w, read); err != nil {
		return nil, err
	}
	return read, nil
}

// Lectura con manejo estandar del  error
func (d *device) Read_Error(w []byte) []byte {
	if read, err := d.Read(w); err != nil {
		log.Fatalf("Error lectura del registro 0x%2X, Error: %v", w, err)
		return nil
	} else {
		return read
	}
}
