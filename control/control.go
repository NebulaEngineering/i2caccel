// El paquete control proporciona la lógica para interpretar los valores de los registros del
// acelerómetro LT8640A, traduciendo los datos binarios en representaciones legibles y significativas
// para el usuario.
//
// Define constantes y mapas que se utilizan para interpretar y configurar los valores de los
// registros del acelerómetro LT8640A.  Forma parte de una CLI que permite leer y configurar
// las propiedades del acelerómetro.
package control

import (
	r "cli_acc/main/register"
	"fmt"
)

// Control representa un propiedad configurable del acelerómetro LT8640A.
// Almacena el nombre de la propiedad, el registro asociado, la posición del bit o bits
// que representan la propiedades dentro del registro, y un mapa para traducir el valor
// binario a una representación legible.
type Control struct {
	Name     string
	Registro *r.Register
	Posicion byte
	Mapping  map[byte]any
	Str      string
}

func (C *Control) Interpretar() {
	var Value byte
	if C.Registro.Update {
		Value = (C.Registro.Value & C.Posicion)
		fmt.Printf("*%v:%v%v", C.Name, C.Mapping[Value], C.Str)
	} else {
		Value = (C.Registro.Read_Error()[0] & C.Posicion)
		fmt.Printf(" %v:%v%v", C.Name, C.Mapping[Value], C.Str)
	}
}

func (C *Control) Previo(n byte) {
	C.Registro.Value = n
	C.Registro.Update = true
}
