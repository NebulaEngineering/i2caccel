package view

import (
	r "i2caccel/register/kx023"
	v "i2caccel/view"
)

var XHP = v.View{Name: "XHP", Registro1: &r.XHPH, Registro2: &r.XHPL, Conversor: (9.81 / 0x4000), Str: "m/s²"}
var YHP = v.View{Name: "YHP", Registro1: &r.YHPH, Registro2: &r.YHPL, Conversor: (9.81 / 0x4000), Str: "m/s²"}
var ZHP = v.View{Name: "ZHP", Registro1: &r.ZHPH, Registro2: &r.ZHPL, Conversor: (9.81 / 0x4000), Str: "m/s²"}

var XOUT = v.View{Name: "XOUT", Registro1: &r.XOUTH, Registro2: &r.XOUTL, Conversor: (9.81 / 0x4000), Str: "m/s²"}
var YOUT = v.View{Name: "YOUT", Registro1: &r.YOUTH, Registro2: &r.YOUTL, Conversor: (9.81 / 0x4000), Str: "m/s²"}
var ZOUT = v.View{Name: "ZOUT", Registro1: &r.ZOUTH, Registro2: &r.ZOUTL, Conversor: (9.81 / 0x4000), Str: "m/s²"}
