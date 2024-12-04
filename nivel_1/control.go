package nivel1

import (
	n0 "cli_acc/main/nivel_0"
	"fmt"
)

type Control struct {
	Name     string
	Registro *n0.Register
	Posicion byte
	Value    byte
}

func (C *Control) Interpretar() {
	value := (C.Registro.Read_Error()[0] & C.Posicion)

	//for range()

	fmt.Printf("la propiedad %v, es 0x%02X\n", C.Name, value)
}

func (C *Control) Previo() {

}
