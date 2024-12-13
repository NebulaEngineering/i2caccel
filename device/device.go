// El paquete device proporciona una abstracción para interactuar con dispositivos I2C,
// gestionando la comunicación a bajo nivel y ofreciendo métodos para leer y escribir datos
// en los registros del dispositivo.
package device

import (
	"fmt"
	"log"
	"sync"

	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/conn/v3/i2c"
	"periph.io/x/conn/v3/i2c/i2creg"
	"periph.io/x/host/v3"
)

// Device representa un dispositivo I2C. Contiene la información necesaria para
// interactuar con un dispositivo a través del bus I2C.
type Device struct {
	name    string        // Nombre del dispositivo I2C.
	adrr    uint16        // Dirección I2C del dispositivo.
	bus     string        // Nombre del bus I2C al que se conecta (ej: "1", "2").
	iobus   i2c.BusCloser // Instancia del bus I2C.
	devi2c  *i2c.Dev      // Representación del dispositivo I2C en el bus.
	gpio    []string
	devgpio []gpio.PinIO
	mux     sync.Mutex // Mutex para asegurar acceso exclusivo al dispositivo.
}

// Init inicializa el dispositivo I2C.  Configura la conexión con el bus I2C
// y el dispositivo especificado.
func (d *Device) Init() {
	// Inicializa el host para acceder a los periféricos.
	if _, err := host.Init(); err != nil {
		log.Fatalf("Error al inicializar el host: %v", err)
	}

	// Inicializa cada GPIO especificado.
	for _, pinName := range d.gpio {
		pin := gpioreg.ByName(pinName)
		if pin == nil {
			log.Fatalf("Error: No se encontró el GPIO %s", pinName)
		}
		if err := pin.In(gpio.PullNoChange, gpio.RisingEdge); err != nil {
			log.Fatalf("Error al configurar el GPIO %s: %v", pinName, err)
		}
		d.devgpio = append(d.devgpio, pin)
	}

	// Abre el bus I2C especificado por el nombre.
	iobus, err := i2creg.Open(d.bus)
	if err != nil {
		log.Fatalf("Error al abrir el bus I2C '%s': %v", d.bus, err)
	}
	d.iobus = iobus // Guarda el bus para poder cerrarlo posteriormente.

	// Crea una instancia del dispositivo I2C con la dirección y el bus.
	d.devi2c = &i2c.Dev{Addr: d.adrr, Bus: iobus}
}

// Close cierra la conexión con el bus I2C liberando los recursos.
func (d *Device) Close() error {
	return d.iobus.Close()
}

// Write escribe datos en el dispositivo I2C.
// Recibe un slice de bytes 'w' que contiene los datos a escribir.
// Retorna un error si ocurre algún problema durante la escritura.
func (d *Device) Write(w []byte) error {
	d.mux.Lock()         // Bloquea el mutex para evitar acceso concurrente.
	defer d.mux.Unlock() // Libera el mutex al finalizar la función.

	if _, err := d.devi2c.Write(w); err != nil {
		return err // Retorna el error si la escritura falla.
	}
	return nil // Retorna nil si la escritura es exitosa.
}

// Write_Error escribe datos en el dispositivo I2C.
// Recibe un slice de bytes 'w' que contiene los datos a escribir.
// Si ocurre un error durante la escritura, termina el programa con un mensaje de error.
func (d *Device) Write_Error(w []byte) {
	if err := d.Write(w); err != nil {
		log.Fatalf("Error al escribir en el registro 0x%2X: %v", w, err)
	}
}

// Read lee datos del dispositivo I2C.
// Recibe un slice de bytes 'w' que contiene la dirección del registro a leer.
// Retorna un slice de bytes con los datos leídos y un error si ocurre algún problema.
func (d *Device) Read(w []byte) ([]byte, error) {
	d.mux.Lock()         // Bloquea el mutex para evitar acceso concurrente.
	defer d.mux.Unlock() // Libera el mutex al finalizar la función.

	read := make([]byte, 1) // Crea un buffer para almacenar los datos leídos (1 byte).
	if err := d.devi2c.Tx(w, read); err != nil {
		return nil, err // Retorna nil y el error si la lectura/transacción falla.
	}
	return read, nil // Retorna los datos leídos y nil si la lectura es exitosa.
}

// Read_Error lee datos del dispositivo I2C.
// Recibe un slice de bytes 'w' que contiene la dirección del registro a leer.
// Si ocurre un error durante la lectura, termina el programa con un mensaje de error.
func (d *Device) Read_Error(w []byte) []byte {
	read, err := d.Read(w) // Intenta leer los datos.
	if err != nil {
		log.Fatalf("Error al leer del registro 0x%2X: %v", w, err)
		return nil // Retorna nil si ocurre un error (aunque la función termina por log.Fatalf).
	}
	return read // Retorna los datos leídos si la lectura es exitosa.
}

func (d *Device) ReadGPIO(index int) (bool, error) {
	d.mux.Lock()         // Bloquea el mutex para evitar acceso concurrente.
	defer d.mux.Unlock() // Libera el mutex al finalizar la función.
	if index < 0 || index >= len(d.devgpio) {
		return false, fmt.Errorf("índice de GPIO %d fuera de rango", index)
	}
	pin := d.devgpio[index]
	if pin == nil {
		return false, fmt.Errorf("GPIO en el índice %d no está inicializado", index)
	}
	return pin.Read() == gpio.High, nil
}

func (d *Device) ReadGPIO_Error(index int) bool {
	read, err := d.ReadGPIO(index)
	if err != nil {
		log.Fatalf("Error al configurar el GPIO%d como entrada: %v", index, err)
		return false
	}
	return read
}
