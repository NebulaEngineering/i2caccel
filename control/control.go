package control

import (
	r "cli_acc/main/register"
	"fmt"
)

type Control struct {
	Name     string
	Registro *r.Register
	Posicion byte
	Mapping  map[byte]any
	A        any
}

func (C *Control) Interpretar() {
	var Value byte
	if C.Registro.Update {
		Value = (C.Registro.Value & C.Posicion)
		fmt.Printf("*%v:%v", C.Name, C.Mapping[Value])
	} else {
		Value = (C.Registro.Read_Error()[0] & C.Posicion)
		fmt.Printf(" %v:%v", C.Name, C.Mapping[Value])
	}
}

func (C *Control) Previo(n byte) {
	C.Registro.Value = n
	C.Registro.Update = true
}

// func NewControl(name string, registro *n0.Register, posicion byte, mapping map[byte]string) Control {

// 	// condicion para Mapping si usa un bit y varios bits
// 	var mapa map[byte]string
// 	if list := []byte{0x01, 0x02, 0x04, 0x08, 0x10, 0x20, 0x40, 0x80}; slices.Index(list, posicion) == -1 {
// 		mapa = mapping
// 	} else {
// 		mapa = map[byte]string{posicion: "activo", 0x00: "inactivo"}
// 	}

// 	return Control{
// 		Name:     name,
// 		Registro: registro,
// 		Posicion: posicion,
// 		Mapping:  mapa,
// 	}
// }
