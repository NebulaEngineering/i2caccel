// El paquete register define la estructura Register para interactuar con los registros del acelerómetro
// LT8640A, proporcionando métodos para leer y escribir valores, y forma parte de una CLI para configurar
// y leer datos del dispositivo.
package register

import (
	D "goi2caccel/device" // Importa el paquete "device" para interactuar con el dispositivo I2C.
)

// Register representa un registro de un dispositivo I2C, que facilita
// la interaccion con un registro definido
type Register struct {
	Name   string    // Nombre del registro (ej: "XHP_L", "CNTL1").
	Addr   byte      // Dirección del registro en el dispositivo I2C.
	Value  byte      // Valor actual del registro, actualizado tras una lectura.
	Device *D.Device // Puntero al dispositivo I2C (LT8640A) al que pertenece el registro.
	Update bool      // Indica si el valor del registro ha sido actualizado.
}

// Write_Error escribe un valor en el registro.  Recibe el valor como un
// slice de bytes.  En caso de error durante la escritura, termina el programa
// con un mensaje de error, simplificando la gestión de errores en la CLI.
func (r *Register) Write_Error(value []byte) {
	data := []byte{r.Addr}
	data = append(data, value...)
	r.Device.Write_Error(data)
}

// Write escribe un valor en el registro. Recibe el valor como un slice de
// bytes.  Retorna un error si ocurre algún problema durante la escritura,
// permitiendo un manejo de errores más preciso.
func (r *Register) Write(value []byte) error {
	data := []byte{r.Addr}
	data = append(data, value...)
	return r.Device.Write(data)
}

// Read_Error lee el valor del registro y lo almacena en el campo `Value`.
//
//	Utiliza el método `Read_Error` del dispositivo, que termina el programa
//
// en caso de error, simplificando la gestión de errores en la CLI.  Retorna
// el valor leído como un slice de bytes.
func (r *Register) Read_Error() []byte {
	r.Value = r.Device.Read_Error([]byte{r.Addr})[0]
	return []byte{r.Value}
}

// Read lee el valor del registro y lo almacena en el campo `Value`.
// Utiliza el método `Read` del dispositivo y retorna el valor leído como un
// slice de bytes, junto con un error si la lectura falla, permitiendo un manejo
// de errores más preciso.
func (r *Register) Read() ([]byte, error) {
	data, err := r.Device.Read([]byte{r.Addr})
	if err != nil {
		return nil, err
	}
	r.Value = data[0]
	return []byte{r.Value}, nil

}
