package control

import (
	c "i2caccel/control"
	x "i2caccel/register"
	r "i2caccel/register/kx023"
)

// Reporta si la interrupcion lo origino el evento buffer lleno.
var BFI = c.Control{Name: "BFI", Registro: &r.INS2, Posicion: 0x40, Mapping: Booleano(0x40)}

// Reporta el Watermark del buffer.
var WMI = c.Control{Name: "WMI", Registro: &r.INS2, Posicion: 0x20, Mapping: Booleano(0x20)}

// Reporta nuevo datos disponibles.
var DRDY = c.Control{Name: "DRDY", Registro: &r.INS2, Posicion: 0x10, Mapping: Booleano(0x10)}

// Reporta los eventos de tap/double tap.
var TDTS = c.Control{Name: "TDTS", Registro: &r.INS2, Posicion: 0x0C, Mapping: Booleano(0x0C)}

// Reporta los eventos de deteccion de movimiento.
var WUFS = c.Control{Name: "WUFS", Registro: &r.INS2, Posicion: 0x02, Mapping: Booleano(0x02)}

// Reporta los eventos de deteccion de movimiento.
var TPS = c.Control{Name: "TPS", Registro: &r.INS2, Posicion: 0x01, Mapping: Booleano(0x01)}

// Alterna entre modo sleep y run de acelerometro.
var PC1 = c.Control{Name: "PC1", Registro: &r.CNTL1, Posicion: 0x80, Mapping: Booleano(0x80)}

// Alterna la resolucion entre bajo consumo y alta resolucion.
var RES = c.Control{Name: "RES", Registro: &r.CNTL1, Posicion: 0x40, Mapping: Booleano(0x40)}

// Alterna la deteccion de datos nuevos en los registro en la salida de los datos.
var DRDYE = c.Control{Name: "DRDYE", Registro: &r.CNTL1, Posicion: 0x20, Mapping: Booleano(0x20)}

// Alterna entre diferente valores de sensibilidad del acelerometro (2g,4g,8g).
var GSEL = c.Control{Name: "GSEL", Registro: &r.CNTL1, Posicion: 0x18, Mapping: GSEL_C, Str: " G"}

// Alterna la deteccion de evento tap y double tap.
var TDTE = c.Control{Name: "TDTE", Registro: &r.CNTL1, Posicion: 0x04, Mapping: Booleano(0x04)}

// Alterna la deteccion de movimiento.
var WUFE = c.Control{Name: "WUFE", Registro: &r.CNTL1, Posicion: 0x02, Mapping: Booleano(0x02)}

// Alterna la deteccion de la orientacion relativia a la gravedad.
var TPE = c.Control{Name: "TPE", Registro: &r.CNTL1, Posicion: 0x01, Mapping: Booleano(0x01)}

// Inicializa el reinicio de sensor y su RAM.
var SRST = c.Control{Name: "SRST", Registro: &r.CNTL2, Posicion: 0x80, Mapping: Booleano(0x80)}

// Inicializa un test de comprobacion en la opreacion de sensor
var COTC = c.Control{Name: "COTC", Registro: &r.CNTL2, Posicion: 0x40, Mapping: Booleano(0x40)}

// Alterna la detecion del evento de la orientacion del sensor en el eje -x.
var LEM = c.Control{Name: "LEM", Registro: &r.CNTL2, Posicion: 0x20, Mapping: Booleano(0x20)}

// Alterna la detecion del evento de la orientacion del sensor en el eje +x.
var RIM = c.Control{Name: "RIM", Registro: &r.CNTL2, Posicion: 0x10, Mapping: Booleano(0x10)}

// Alterna la detecion del evento de la orientacion del sensor en el eje -y.
var DOM = c.Control{Name: "DOM", Registro: &r.CNTL2, Posicion: 0x08, Mapping: Booleano(0x08)}

// Alterna la detecion del evento de la orientacion del sensor en el eje +y.
var UPM = c.Control{Name: "UPM", Registro: &r.CNTL2, Posicion: 0x04, Mapping: Booleano(0x04)}

// Alterna la detecion del evento de la orientacion del sensor en el eje -z.
var FDM = c.Control{Name: "FDM", Registro: &r.CNTL2, Posicion: 0x02, Mapping: Booleano(0x02)}

// Alterna la detecion del evento de la orientacion del sensor en el eje +z.
var FUM = c.Control{Name: "FUM", Registro: &r.CNTL2, Posicion: 0x01, Mapping: Booleano(0x01)}

// Alterna la frecuencia de salida de datos de la orientacion del accelerometro.
var OTP = c.Control{Name: "OTP", Registro: &r.CNTL3, Posicion: 0xC0}

// Alterna la frecuencia de salida de datos de la direccion del tap.
var OTDT = c.Control{Name: "OTDT", Registro: &r.CNTL3, Posicion: 0x38}

// Alterna la frecuencia de salida de datos general de la deccion  de movimiento y salida con filtro pasa alta.
var OWUF = c.Control{Name: "OWUF", Registro: &r.CNTL3, Posicion: 0x07, Mapping: OWUF_C, Str: " Hz"}

// Alternar la desactivacion de filtro pasa baja de la salida normal.
var IIR_BYPASS = c.Control{Name: "IIR_BYPASS", Registro: &r.ODCNTL, Posicion: 0x80, Mapping: Booleano(0x80)}

// Alternar el valor de la frecuencia de contre del fitro pasa baja.
var LPR0 = c.Control{Name: "LPR0", Registro: &r.ODCNTL, Posicion: 0x40, Mapping: Booleano(0x40)}

// Alterna el valor de la frecuencia de salida de datos en la salida normal de accelerometro.
var OSA = c.Control{Name: "OSA", Registro: &r.ODCNTL, Posicion: 0x0F, Mapping: OSA_C, Str: " Hz"}

// Alterna la activacion de pin fisico para la interrupcion por INT1.
var IEN = c.Control{Name: "IEN", Registro: &r.INC1, Posicion: 0x20, Mapping: Booleano(0x20)}

// Alterna la poralidad de estado activo e inactivo en INT1.
var IEA = c.Control{Name: "IEA", Registro: &r.INC1, Posicion: 0x10, Mapping: Booleano(0x10)}

// Alterna modo de repuesta para la interupcion en INT1.
var IEL = c.Control{Name: "IEL", Registro: &r.INC1, Posicion: 0x08, Mapping: Booleano(0x08)}

// Alterna la poralidad de Self-test.
var STPOL = c.Control{Name: "STPOL", Registro: &r.INC1, Posicion: 0x02, Mapping: Booleano(0x02)}

// alterna la acivacion de al interfaz SPI 3-wire (no modificar).
var SPI3E = c.Control{Name: "SPI3E", Registro: &r.INC1, Posicion: 0x01, Mapping: Booleano(0x01)}

// Alterna la deteccion de movimiento por el eje -x.
var XNWUE = c.Control{Name: "XNWUE", Registro: &r.INC2, Posicion: 0x20, Mapping: Booleano(0x20)}

// Alterna la deteccion de movimiento por el eje +x.
var XPWUE = c.Control{Name: "XPWUE", Registro: &r.INC2, Posicion: 0x10, Mapping: Booleano(0x10)}

// Alterna la deteccion de movimiento por el eje -y.
var YNWUE = c.Control{Name: "YNWUE", Registro: &r.INC2, Posicion: 0x08, Mapping: Booleano(0x08)}

// Alterna la deteccion de movimiento por el eje -y.
var YPWUE = c.Control{Name: "YPWUE", Registro: &r.INC2, Posicion: 0x04, Mapping: Booleano(0x04)}

// Alterna la deteccion de movimiento por el eje -z.
var ZNWUE = c.Control{Name: "ZNWUE", Registro: &r.INC2, Posicion: 0x02, Mapping: Booleano(0x02)}

// Alterna la deteccion de movimiento por el eje +z.
var ZPWUE = c.Control{Name: "ZPWUE", Registro: &r.INC2, Posicion: 0x01, Mapping: Booleano(0x01)}

// Alterna la deteccion de evento tap/double en la direccion -x
var TLEM = c.Control{Name: "TLEM", Registro: &r.INC3, Posicion: 0x20, Mapping: Booleano(0x20)}

// Alterna la deteccion de evento tap/double en la direccion +x
var TRIM = c.Control{Name: "TRIM", Registro: &r.INC3, Posicion: 0x10, Mapping: Booleano(0x10)}

// Alterna la deteccion de evento tap/double en la direccion -y
var TDOM = c.Control{Name: "TDOM", Registro: &r.INC3, Posicion: 0x08, Mapping: Booleano(0x08)}

// Alterna la deteccion de evento tap/double en la direccion +y
var TUPM = c.Control{Name: "TUPM", Registro: &r.INC3, Posicion: 0x04, Mapping: Booleano(0x04)}

// Alterna la deteccion de evento tap/double en la direccion -z
var TFDM = c.Control{Name: "TFDM", Registro: &r.INC3, Posicion: 0x02, Mapping: Booleano(0x02)}

// Alterna la deteccion de evento tap/double en la direccion +z
var TFUM = c.Control{Name: "TFUM", Registro: &r.INC3, Posicion: 0x01, Mapping: Booleano(0x01)}

// Alterna la deteccion del evento bufer lleno en el PIN INT1
var BFI1 = c.Control{Name: "BFI1", Registro: &r.INC4, Posicion: 0x40, Mapping: Booleano(0x40)}

// Alterna la deteccion del evento Watermark en el PIN INT1
var WMI1 = c.Control{Name: "WMI1", Registro: &r.INC4, Posicion: 0x20, Mapping: Booleano(0x20)}

// Alterna la deteccion del evento datos de salida nuevos en el PIN INT1
var DRDYI1 = c.Control{Name: "DRDYI1", Registro: &r.INC4, Posicion: 0x10, Mapping: Booleano(0x10)}

// Alterna la deteccion del evento Tap/Double en el PIN INT1
var TDTI1 = c.Control{Name: "TDTI1", Registro: &r.INC4, Posicion: 0x04, Mapping: Booleano(0x04)}

// Alterna la deteccion del evento Motion Detect en el PIN INT1
var WUFI1 = c.Control{Name: "WUFI1", Registro: &r.INC4, Posicion: 0x02, Mapping: Booleano(0x02)}

// Alterna la deteccion del evento posicion de la orientacion  en el PIN INT1
var TPI1 = c.Control{Name: "TPI1", Registro: &r.INC4, Posicion: 0x01, Mapping: Booleano(0x01)}

// Alterna la activacion de pin fisico para la interrupcion por INT2.
var IEN2 = c.Control{Name: "IEN2", Registro: &r.INC5, Posicion: 0x20, Mapping: Booleano(0x20)}

// Alterna la poralidad de estado activo e inactivo en INT2.
var IEA2 = c.Control{Name: "IEA2", Registro: &r.INC5, Posicion: 0x10, Mapping: Booleano(0x10)}

// Alterna modo de repuesta para la interupcion en INT2.
var IEL2 = c.Control{Name: "IEL2", Registro: &r.INC5, Posicion: 0x08, Mapping: Booleano(0x08)}

// Alterna la deteccion del evento bufer lleno en el PIN INT2.
var BFI2 = c.Control{Name: "BFI2", Registro: &r.INC6, Posicion: 0x40, Mapping: Booleano(0x40)}

// Alterna la deteccion del evento Watermark en el PIN INT2.
var WMI2 = c.Control{Name: "WMI2", Registro: &r.INC6, Posicion: 0x20, Mapping: Booleano(0x20)}

// Alterna la deteccion del evento datos de salida nuevos en el PIN INT2.
var DRDYI2 = c.Control{Name: "DRDYI2", Registro: &r.INC6, Posicion: 0x10, Mapping: Booleano(0x10)}

// Alterna la deteccion del evento Tap/Double en el PIN INT2.
var TDTI2 = c.Control{Name: "TDTI2", Registro: &r.INC6, Posicion: 0x04, Mapping: Booleano(0x04)}

// Alterna la deteccion del evento Motion Detect en el PIN INT2.
var WUFI2 = c.Control{Name: "WUFI2", Registro: &r.INC6, Posicion: 0x02, Mapping: Booleano(0x02)}

// Alterna la deteccion del evento posicion de la orientacion  en el PIN INT2.
var TPI2 = c.Control{Name: "TPI2", Registro: &r.INC6, Posicion: 0x01, Mapping: Booleano(0x01)}

// Valor de contador minimo para  la deteccion de la orientacion de acelerometro.
var TSC = c.Control{Name: "TSC", Registro: &r.TILT_TIMER, Posicion: 0xFF, Mapping: nil}

// Valor del contador minimo para la activacion de la deteccion de movimiento.
var WUFC = c.Control{Name: "WUFC", Registro: &r.WUFC, Posicion: 0xFF, Mapping: nil}

// Valor de umbral minimo para la activacion de la deteccion de movimiento.
var ATH = c.Control{Name: "ATH", Registro: &r.ATH, Posicion: 0xFF, Mapping: nil}

// Valor del promedio de los datos de salida.
var AVC = c.Control{Name: "AVC", Registro: &r.LP_CNTL, Posicion: 0x70, Mapping: AVC_C, Str: " Samples"}

// Valor de umbral definido para el buffer.
var SMP_TH = c.Control{Name: "SMP_TH", Registro: &r.BUF_CNTL1, Posicion: 0x7F, Mapping: nil}

// Alterna la activacion y desactivacion de bufer.
var BUFE = c.Control{Name: "BUFE", Registro: &r.BUF_CNTL2, Posicion: 0x80, Mapping: Booleano(0x80)}

// Alterna la resolucion entre 8 bit y 16bits de la muestra de salida.
var BRES = c.Control{Name: "BRES", Registro: &r.BUF_CNTL2, Posicion: 0x40, Mapping: Booleano(0x40)}

// Alterna habilitacion de detecion de evento buffer lleno, solo en el INS2.
var BFIE = c.Control{Name: "BFIE", Registro: &r.BUF_CNTL2, Posicion: 0x20, Mapping: Booleano(0x20)}

// Alterna en el modo de operacion de buffer (filo, stream ,trigger,filo).
var BUF_M = c.Control{Name: "BUF_M", Registro: &r.BUF_CNTL2, Posicion: 0x03, Mapping: BUF_M_C}

// Reporta el numero de muestra almacenado actualmente en el buffer.
var SMP_LEV = c.Control{Name: "SMP_LEV", Registro: &r.BUF_STATUS_1, Posicion: 0xFF, Mapping: nil}

// Reporta la detencion del evento de trigger de acelerometro.
var BUF_TRIG = c.Control{Name: "BUF_TRIG", Registro: &r.BUF_STATUS_2, Posicion: 0x80, Mapping: Booleano(0x80)}

// Regitro dummy para limpiar el buffer y los evento de interrupcion relacionado con el buffer.
var BUF_CLEAR = c.Control{Name: "BUF_CLEAR", Registro: &r.BUF_CLEAR, Posicion: 0xFF, Mapping: nil}

// Leer una muestra del bufer, la ubicacion que toma los datos de buffer varia por el modo de operacion.
var BUF_READ = c.Control{Name: "BUF_READ", Registro: &r.BUF_READ, Posicion: 0xFF, Mapping: nil}

// slice que tenermina un orden especifico para mostrar los datos
var Seq_Register = []*x.Register{
	&r.INC1,
	&r.INC2,
	&r.INC3,
	&r.INC4,
	&r.INC5,
	&r.INC6,
	&r.ODCNTL,
	&r.CNTL1,
	&r.CNTL3,
	&r.ATH,
	&r.WUFC,
	&r.LP_CNTL,
	&r.BUF_CNTL1,
	&r.BUF_CNTL2,
}

// Mapeado para la relacion entre los registro y sus propiedades especificas
var Mregister = map[*x.Register][]*c.Control{
	&r.CNTL1:     {&PC1, &RES, &DRDYE, &GSEL, &WUFE, &TPE},
	&r.CNTL3:     {&OWUF},
	&r.ODCNTL:    {&IIR_BYPASS, &LPR0, &OSA},
	&r.INC1:      {&IEN, &IEA, &IEL, &STPOL, &SPI3E},
	&r.INC2:      {&XNWUE, &XPWUE, &YNWUE, &YPWUE, &ZNWUE, &ZPWUE},
	&r.INC3:      {&TLEM, &TRIM, &TDOM, &TUPM, &TFDM, &TFUM},
	&r.INC4:      {&BFI1, &WMI1, &DRDYI1, &TDTI1, &WUFI1, &TPI1},
	&r.INC5:      {&IEN2, &IEA2, &IEL2},
	&r.INC6:      {&BFI2, &WMI2, &DRDYI2, &TDTI2, &WUFI2, &TPI2},
	&r.WUFC:      {&WUFC},
	&r.ATH:       {&ATH},
	&r.LP_CNTL:   {&AVC},
	&r.BUF_CNTL1: {&SMP_TH},
	&r.BUF_CNTL2: {&BUFE, &BRES, &BFIE, &BUF_M},
}
