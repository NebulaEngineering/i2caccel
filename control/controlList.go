package control

import (
	r "cli_acc/main/register"
)

var PC1 = Control{Name: "PC1", Registro: &r.CNTL1, Posicion: 0x80, Mapping: Booleano(0x80)}
var RES = Control{Name: "RES", Registro: &r.CNTL1, Posicion: 0x40, Mapping: Booleano(0x40)}
var DRDYE = Control{Name: "DRDYE", Registro: &r.CNTL1, Posicion: 0x20, Mapping: Booleano(0x20)}
var GSEL = Control{Name: "GSEL", Registro: &r.CNTL1, Posicion: 0x18, Mapping: GSEL_C, Str: " G"}
var TDTE = Control{Name: "TDTE", Registro: &r.CNTL1, Posicion: 0x04, Mapping: Booleano(0x04)}
var WUFE = Control{Name: "WUFE", Registro: &r.CNTL1, Posicion: 0x02, Mapping: Booleano(0x02)}
var TPE = Control{Name: "TPE", Registro: &r.CNTL1, Posicion: 0x01, Mapping: Booleano(0x01)}

var SRST = Control{Name: "SRST", Registro: &r.CNTL2, Posicion: 0x80, Mapping: Booleano(0x80)}
var COTC = Control{Name: "COTC", Registro: &r.CNTL2, Posicion: 0x40, Mapping: Booleano(0x40)}
var LEM = Control{Name: "LEM", Registro: &r.CNTL2, Posicion: 0x20, Mapping: Booleano(0x20)}
var RIM = Control{Name: "RIM", Registro: &r.CNTL2, Posicion: 0x10, Mapping: Booleano(0x10)}
var DOM = Control{Name: "DOM", Registro: &r.CNTL2, Posicion: 0x08, Mapping: Booleano(0x08)}
var UPM = Control{Name: "UPM", Registro: &r.CNTL2, Posicion: 0x04, Mapping: Booleano(0x04)}
var FDM = Control{Name: "FDM", Registro: &r.CNTL2, Posicion: 0x02, Mapping: Booleano(0x02)}
var FUM = Control{Name: "FUM", Registro: &r.CNTL2, Posicion: 0x01, Mapping: Booleano(0x01)}

var OTP = Control{Name: "OTP", Registro: &r.CNTL3, Posicion: 0xC0}
var OTDT = Control{Name: "OTDT", Registro: &r.CNTL3, Posicion: 0x38}
var OWUF = Control{Name: "OWUF", Registro: &r.CNTL3, Posicion: 0x07, Mapping: OWUF_C, Str: " Hz"}

var IIR_BYPASS = Control{Name: "IIR_BYPASS", Registro: &r.ODCNTL, Posicion: 0x80, Mapping: Booleano(0x80)}
var LPR0 = Control{Name: "LPR0", Registro: &r.ODCNTL, Posicion: 0x40, Mapping: Booleano(0x40)}
var OSA = Control{Name: "OSA", Registro: &r.ODCNTL, Posicion: 0x0F, Mapping: OSA_C, Str: " Hz"}

var IEN = Control{Name: "IEN", Registro: &r.INC1, Posicion: 0x20, Mapping: Booleano(0x20)}
var IEA = Control{Name: "IEA", Registro: &r.INC1, Posicion: 0x10, Mapping: Booleano(0x10)}
var IEL = Control{Name: "IEL", Registro: &r.INC1, Posicion: 0x08, Mapping: Booleano(0x08)}
var STPOL = Control{Name: "STPOL", Registro: &r.INC1, Posicion: 0x02, Mapping: Booleano(0x02)}
var SPI3E = Control{Name: "SPI3E", Registro: &r.INC1, Posicion: 0x01, Mapping: Booleano(0x01)}

var XNWUE = Control{Name: "XNWUE", Registro: &r.INC2, Posicion: 0x20, Mapping: Booleano(0x20)}
var XPWUE = Control{Name: "XPWUE", Registro: &r.INC2, Posicion: 0x10, Mapping: Booleano(0x10)}
var YNWUE = Control{Name: "YNWUE", Registro: &r.INC2, Posicion: 0x08, Mapping: Booleano(0x08)}
var YPWUE = Control{Name: "YPWUE", Registro: &r.INC2, Posicion: 0x04, Mapping: Booleano(0x04)}
var ZNWUE = Control{Name: "ZNWUE", Registro: &r.INC2, Posicion: 0x02, Mapping: Booleano(0x02)}
var ZPWUE = Control{Name: "ZPWUE", Registro: &r.INC2, Posicion: 0x01, Mapping: Booleano(0x01)}

var TLEM = Control{Name: "TLEM", Registro: &r.INC3, Posicion: 0x20, Mapping: Booleano(0x20)}
var TRIM = Control{Name: "TRIM", Registro: &r.INC3, Posicion: 0x10, Mapping: Booleano(0x10)}
var TDOM = Control{Name: "TDOM", Registro: &r.INC3, Posicion: 0x08, Mapping: Booleano(0x08)}
var TUPM = Control{Name: "TUPM", Registro: &r.INC3, Posicion: 0x04, Mapping: Booleano(0x04)}
var TFDM = Control{Name: "TFDM", Registro: &r.INC3, Posicion: 0x02, Mapping: Booleano(0x02)}
var TFUM = Control{Name: "TFUM", Registro: &r.INC3, Posicion: 0x01, Mapping: Booleano(0x01)}

var BFI1 = Control{Name: "BFI1", Registro: &r.INC4, Posicion: 0x40, Mapping: Booleano(0x40)}
var WMI1 = Control{Name: "WMI1", Registro: &r.INC4, Posicion: 0x20, Mapping: Booleano(0x20)}
var DRDYI1 = Control{Name: "DRDYI1", Registro: &r.INC4, Posicion: 0x10, Mapping: Booleano(0x10)}
var TDTI1 = Control{Name: "TDTI1", Registro: &r.INC4, Posicion: 0x04, Mapping: Booleano(0x04)}
var WUFI1 = Control{Name: "WUFI1", Registro: &r.INC4, Posicion: 0x02, Mapping: Booleano(0x02)}
var TPI1 = Control{Name: "TPI1", Registro: &r.INC4, Posicion: 0x01, Mapping: Booleano(0x01)}

var IEN2 = Control{Name: "IEN2", Registro: &r.INC5, Posicion: 0x20, Mapping: Booleano(0x20)}
var IEA2 = Control{Name: "IEA2", Registro: &r.INC5, Posicion: 0x10, Mapping: Booleano(0x10)}
var IEL2 = Control{Name: "IEL2", Registro: &r.INC5, Posicion: 0x08, Mapping: Booleano(0x08)}

var BFI2 = Control{Name: "BFI2", Registro: &r.INC6, Posicion: 0x40, Mapping: Booleano(0x40)}
var WMI2 = Control{Name: "WMI2", Registro: &r.INC6, Posicion: 0x20, Mapping: Booleano(0x20)}
var DRDYI2 = Control{Name: "DRDYI2", Registro: &r.INC6, Posicion: 0x10, Mapping: Booleano(0x10)}
var TDTI2 = Control{Name: "TDTI2", Registro: &r.INC6, Posicion: 0x04, Mapping: Booleano(0x04)}
var WUFI2 = Control{Name: "WUFI2", Registro: &r.INC6, Posicion: 0x02, Mapping: Booleano(0x02)}
var TPI2 = Control{Name: "TPI2", Registro: &r.INC6, Posicion: 0x01, Mapping: Booleano(0x01)}

var TSC = Control{Name: "TSC", Registro: &r.TILT_TIMER, Posicion: 0xFF, Mapping: nil}

var WUFC = Control{Name: "WUFC", Registro: &r.WUFC, Posicion: 0xFF, Mapping: nil}

//-------------------------------------------------

var ATH = Control{Name: "ATH", Registro: &r.ATH, Posicion: 0xFF, Mapping: nil}

var AVC = Control{Name: "AVC", Registro: &r.LP_CNTL, Posicion: 0x70, Mapping: AVC_C, Str: " Samples"}

var SMP_TH = Control{Name: "SMP_TH", Registro: &r.BUF_CNTL1, Posicion: 0x7F, Mapping: nil}

var BUFE = Control{Name: "BUFE", Registro: &r.BUF_CNTL2, Posicion: 0x80, Mapping: Booleano(0x80)}
var BRES = Control{Name: "BRES", Registro: &r.BUF_CNTL2, Posicion: 0x40, Mapping: Booleano(0x40)}
var BFIE = Control{Name: "BFIE", Registro: &r.BUF_CNTL2, Posicion: 0x20, Mapping: Booleano(0x20)}
var BUF_M = Control{Name: "BUF_M", Registro: &r.BUF_CNTL2, Posicion: 0x03, Mapping: BUF_M_C}

var SMP_LEV = Control{Name: "SMP_LEV", Registro: &r.BUF_STATUS_1, Posicion: 0xFF, Mapping: nil}

var BUF_TRIG = Control{Name: "BUF_TRIG", Registro: &r.BUF_STATUS_2, Posicion: 0x80, Mapping: Booleano(0x80)}

var BUF_CLEAR = Control{Name: "BUF_CLEAR", Registro: &r.BUF_CLEAR, Posicion: 0xFF, Mapping: nil}

var BUF_READ = Control{Name: "BUF_READ", Registro: &r.BUF_READ, Posicion: 0xFF, Mapping: nil}

var Mregister = map[*r.Register][]*Control{
	&r.CNTL1:     {&PC1, &RES, &DRDYE, &GSEL, &WUFE, &TPE},
	&r.CNTL3:     {&OWUF},
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
