// Package register define la estructura Register para interactuar con los registros del acelerómetro
// Kx023, proporcionando métodos para leer y escribir valores, y forma parte de una CLI para configurar
// y leer datos del dispositivo.
package register

import (
	D "goi2caccel/device" // Importa el paquete "device" para interactuar con el dispositivo I2C.
)

// Register representa un registro de un dispositivo I2C, facilitando la interacción con un registro
// específico.  Contiene información sobre la dirección, el valor, el dispositivo al que pertenece y
// un indicador de actualización.
type Register struct {
	Name   string    // Nombre del registro (ej: "XHP_L", "CNTL1").
	Addr   byte      // Dirección del registro en el dispositivo I2C.
	Value  byte      // Valor actual del registro, actualizado tras una lectura o escritura.
	Device *D.Device // Puntero al dispositivo I2C (Kx023) al que pertenece el registro.
	Update bool      // Indica si el valor del registro ha sido modificado y necesita ser escrito en el dispositivo.
}

// Write_Error escribe un valor en el registro. Recibe el valor como un slice de bytes.
// En caso de error, termina el programa con un mensaje de error, simplificando el manejo
// de errores críticos en la CLI.
func (r *Register) Write_Error(value []byte) {
	data := []byte{r.Addr}
	data = append(data, value...)
	r.Device.Write_Error(data)
}

// Write escribe un valor en el registro. Recibe el valor como un slice de bytes.
// Retorna un error si ocurre algún problema durante la escritura, permitiendo un manejo
// de errores más granular.
func (r *Register) Write(value []byte) error {
	data := []byte{r.Addr}
	data = append(data, value...)
	return r.Device.Write(data)
}

// Read_Error lee el valor del registro y lo almacena en el campo `Value`.
// En caso de error durante la lectura, termina el programa con un mensaje de error,
// simplificando el manejo de errores críticos en la CLI. Retorna el valor leído
// como un slice de bytes.
func (r *Register) Read_Error() []byte {
	r.Value = r.Device.Read_Error([]byte{r.Addr})[0]
	return []byte{r.Value}
}

// Read lee el valor del registro y lo almacena en el campo `Value`.
// Retorna el valor leído como un slice de bytes y un error si la lectura falla,
// permitiendo un manejo de errores más preciso.
func (r *Register) Read() ([]byte, error) {
	data, err := r.Device.Read([]byte{r.Addr})
	if err != nil {
		return nil, err
	}
	r.Value = data[0]
	return []byte{r.Value}, nil
}

// Actualizar escribe el valor actual del registro (`Value`) en el dispositivo I2C si el
// registro ha sido marcado como actualizado (`Update == true`).  Después de la
// escritura, el flag `Update` se establece a `false`.  Este método asegura que los
// cambios realizados en el valor del registro se reflejen en el dispositivo físico.
func (r *Register) Actualizar() {
	if r.Update {
		r.Device.Write_Error([]byte{r.Value}) // Escribe el valor en el dispositivo.
		r.Update = false                      // Marca el registro como no actualizado.
	}
}
