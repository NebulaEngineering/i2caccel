package view

import (
	"fmt"
	r "i2caccel/register"
)

type View struct {
	Name      string
	Registro1 *r.Register
	Registro2 *r.Register
	Conversor float64
	Str       string
}

func (v *View) String() string {
	// conversion de complemento a 2 hacia normal
	value := int16((uint16(v.Registro1.Read_Error()[0]) << 8) | uint16(v.Registro2.Read_Error()[0]))
	return fmt.Sprintf("%v: %05.3f ", v.Name, float64(value)*v.Conversor)
}
