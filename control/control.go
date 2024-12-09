// El paquete control proporciona la lógica para interpretar los valores de los registros del
// acelerómetro, traduciendo los datos binarios en representaciones legibles y significativas
// para el usuario.
//
// Define constantes y mapas que se utilizan para interpretar y configurar los valores de los
// registros del acelerómetro.  Forma parte de una CLI que permite leer y configurar
// las propiedades del acelerómetro.
package control

import (
	"fmt"
	r "goi2caccel/register"
)

// Control representa un propiedad configurable del acelerómetro.
// Almacena el nombre de la propiedad, el registro asociado, la posición del bit o bits
// que representan la propiedades dentro del registro, y un mapa para traducir el valor
// binario a una representación legible.
type Control struct {
	Name      string
	Registro  *r.Register
	Posicion  byte
	Mapping   map[byte]any
	Str       string
	Update    bool
	Conversor float32
}

// intepretacion de los datos en una propiedad de acelerometro.
func (C *Control) String() string {
	var value byte
	var convert any
	if C.Update {
		value = (C.Registro.Value & C.Posicion)
		if C.Posicion >= 0x7F {
			convert = C.Conversor * float32(value)
		} else {
			convert = C.Mapping[value]
		}
		return fmt.Sprintf("*%v:%v%v", C.Name, convert, C.Str)
	} else {
		value = (C.Registro.Read_Error()[0] & C.Posicion)
		if C.Posicion >= 0x7F {
			convert = C.Conversor * float32(value)
		} else {
			convert = C.Mapping[value]
		}
		return fmt.Sprintf(" %v:%v%v", C.Name, convert, C.Str)
	}
}

// modificar el registro, de forma previa a la actualizacion.
func (C *Control) Previo(n byte) {
	if C.Mapping[n] != nil {
		C.Registro.Value = (C.Registro.Value ^ C.Posicion)
		C.Registro.Value = (C.Registro.Value | n)
		C.Registro.Update = true
		C.Update = true
	} else {
		fmt.Println("fuera de rango")
	}
}
