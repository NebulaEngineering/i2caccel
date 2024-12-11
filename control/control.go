// Package control proporciona la lógica para interpretar los valores de los registros del
// acelerómetro, traduciendo los datos binarios en representaciones legibles y significativas
// para el usuario.
//
// Define constantes y mapas que se utilizan para interpretar y configurar los valores de los
// registros del acelerómetro.  Forma parte de una CLI que permite leer y configurar
// las propiedades del acelerómetro.
package control

import (
	"fmt"
	r "i2caccel/register"
)

// MappableValue es una interfaz para los tipos de datos que pueden ser mapeados a
// valores de registro.  Permite la flexibilidad de usar diferentes tipos de datos
// (float32, byte, string) como valores en el mapeo.
type MappableValue interface {
	isMappable()
}

// Float32Value representa un valor de punto flotante de 32 bits que puede ser mapeado.
type Float32Value float32

// ByteValue representa un valor de byte que puede ser mapeado.
type ByteValue byte

// StringValue representa un valor de cadena que puede ser mapeado.
type StringValue string

// Métodos dummy para implementar la interfaz MappableValue.
func (Float32Value) isMappable() {}
func (ByteValue) isMappable()    {}
func (StringValue) isMappable()  {}

// Control representa una propiedad configurable del acelerómetro.
// Almacena el nombre de la propiedad, el registro asociado, la máscara de bits
// que representan la propiedad dentro del registro, y un mapa para traducir el valor
// binario a una representación legible. El campo `Str` permite añadir una unidad
// o sufijo a la representación de la propiedad.  El campo `Update` indica si
// la propiedad ha sido modificada.
type Control struct {
	Name     string
	Registro *r.Register
	Posicion byte // Máscara de bits para la propiedad
	Mapping  map[byte]MappableValue
	Str      string
	Update   bool
}

// String devuelve una representación en string de la propiedad del acelerómetro.
// Incluye el nombre de la propiedad, el valor interpretado (si existe un mapeo),
// y un asterisco (*) si la propiedad ha sido modificada.
func (C *Control) String() string {
	var strUp string = ""
	var data string

	// Indica si la propiedad ha sido modificada.
	if C.Update {
		strUp = "*"
	}

	// Si no hay un mapeo definido, no se imprime un valor interpretado.
	if C.Mapping == nil {
		data = ""
	} else {
		value := (C.Registro.Value & C.Posicion)
		data = fmt.Sprint(C.Mapping[value])
	}

	return fmt.Sprintf("%v%v:%v%v ", strUp, C.Name, data, C.Str)
}

// Previo modifica el registro del acelerómetro con el nuevo valor proporcionado.
// Si existe un mapeo, busca la clave correspondiente al nuevo valor y actualiza
// el registro en consecuencia.  Si no hay mapeo, actualiza el registro directamente
// con el nuevo valor (asumiendo que es de tipo ByteValue).
func (C *Control) Previo(n MappableValue) {
	if C.Mapping != nil {
		for key, val := range C.Mapping {
			if val == n {
				C.Registro.Value = (C.Registro.Value & ^C.Posicion)
				C.Registro.Value = (C.Registro.Value | key)
				C.Update = true
				C.Registro.Update = true
				return // Salir del bucle una vez que se encuentra la coincidencia
			}
		}
	} else {
		switch v := n.(type) {
		case ByteValue:
			C.Registro.Value = (C.Registro.Value & ^C.Posicion)
			C.Registro.Value = (C.Registro.Value | byte(v))
			C.Update = true
			C.Registro.Update = true
		}
	}
}
