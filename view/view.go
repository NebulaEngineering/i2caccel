package view

import r "goi2caccel/register"

type View struct {
	Name      string
	Registro  *r.Register
	Conversor byte
	Str       string
}

func (v *View) String() string {
	return ""
}
