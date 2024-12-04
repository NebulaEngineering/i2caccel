package nivel1

import n0 "cli_acc/main/nivel_0"

var PC1 = Control{Name: "PC1", Registro: &n0.CNTL1, Posicion: 0x80}
var RES = Control{Name: "RES", Registro: &n0.CNTL1, Posicion: 0x40}
var DRDYE = Control{Name: "DRDYE", Registro: &n0.CNTL1, Posicion: 0x20}
var GSEL = Control{Name: "GSEL", Registro: &n0.CNTL1, Posicion: 0x18}
var TDTE = Control{Name: "TDTE", Registro: &n0.CNTL1, Posicion: 0x04}
var WUFE = Control{Name: "WUFE", Registro: &n0.CNTL1, Posicion: 0x02}
var TPE = Control{Name: "TPE", Registro: &n0.CNTL1, Posicion: 0x01}

var SRST = Control{Name: "SRST", Registro: &n0.CNTL2, Posicion: 0x80}
var COTC = Control{Name: "COTC", Registro: &n0.CNTL2, Posicion: 0x40}
var LEM = Control{Name: "LEM", Registro: &n0.CNTL2, Posicion: 0x20}
var RIM = Control{Name: "RIM", Registro: &n0.CNTL2, Posicion: 0x10}
var DOM = Control{Name: "DOM", Registro: &n0.CNTL2, Posicion: 0x08}
var UPM = Control{Name: "UPM", Registro: &n0.CNTL2, Posicion: 0x04}
var FDM = Control{Name: "FDM", Registro: &n0.CNTL2, Posicion: 0x02}
var FUM = Control{Name: "FUM", Registro: &n0.CNTL2, Posicion: 0x01}

var OTP = Control{Name: "OTP", Registro: &n0.CNTL3, Posicion: 0xC0}
var OTDT = Control{Name: "OTDT", Registro: &n0.CNTL3, Posicion: 0x38}
var OWUF = Control{Name: "OWUF", Registro: &n0.CNTL3, Posicion: 0x07}

var IIR_BYPASS = Control{Name: "IIR_BYPASS", Registro: &n0.ODCNTL, Posicion: 0x80}
var LPR0 = Control{Name: "LPR0", Registro: &n0.ODCNTL, Posicion: 0x40}
var OSA = Control{Name: "OSA", Registro: &n0.ODCNTL, Posicion: 0x0F}

var IEN = Control{Name: "IEN", Registro: &n0.INC1, Posicion: 0x20}
var IEA = Control{Name: "IEA", Registro: &n0.INC1, Posicion: 0x10}
var IEL = Control{Name: "IEL", Registro: &n0.INC1, Posicion: 0x08}
var STPOL = Control{Name: "STPOL", Registro: &n0.INC1, Posicion: 0x02}
var SPI3E = Control{Name: "SPI3E", Registro: &n0.INC1, Posicion: 0x01}

var XNWUE = Control{Name: "XNWUE", Registro: &n0.INC2, Posicion: 0x20}
var XPWUE = Control{Name: "XPWUE", Registro: &n0.INC2, Posicion: 0x10}
var YNWUE = Control{Name: "YNWUE", Registro: &n0.INC2, Posicion: 0x08}
var YPWUE = Control{Name: "YPWUE", Registro: &n0.INC2, Posicion: 0x04}
var ZNWUE = Control{Name: "ZNWUE", Registro: &n0.INC2, Posicion: 0x02}
var ZPWUE = Control{Name: "ZPWUE", Registro: &n0.INC2, Posicion: 0x01}

var TLEM = Control{Name: "TLEM", Registro: &n0.INC3, Posicion: 0x20}
var TRIM = Control{Name: "TRIM", Registro: &n0.INC3, Posicion: 0x10}
var TDOM = Control{Name: "TDOM", Registro: &n0.INC3, Posicion: 0x08}
var TUPM = Control{Name: "TUPM", Registro: &n0.INC3, Posicion: 0x04}
var TFDM = Control{Name: "TFDM", Registro: &n0.INC3, Posicion: 0x02}
var TFUM = Control{Name: "TFUM", Registro: &n0.INC3, Posicion: 0x01}

var BFI1 = Control{Name: "BFI1", Registro: &n0.INC4, Posicion: 0x40}
var WMI1 = Control{Name: "WMI1", Registro: &n0.INC4, Posicion: 0x20}
var DRDYI1 = Control{Name: "DRDYI1", Registro: &n0.INC4, Posicion: 0x10}
var TDTI1 = Control{Name: "TDTI1", Registro: &n0.INC4, Posicion: 0x04}
var WUFI1 = Control{Name: "WUFI1", Registro: &n0.INC4, Posicion: 0x02}
var TPI1 = Control{Name: "TPI1", Registro: &n0.INC4, Posicion: 0x01}

var IEN2 = Control{Name: "IEN2", Registro: &n0.INC5, Posicion: 0x20}
var IEA2 = Control{Name: "TPI1", Registro: &n0.INC5, Posicion: 0x10}
var IEL2 = Control{Name: "TPI1", Registro: &n0.INC5, Posicion: 0x08}

var BFI2 = Control{Name: "BFI2", Registro: &n0.INC6, Posicion: 0x40}
var WMI2 = Control{Name: "WMI2", Registro: &n0.INC6, Posicion: 0x20}
var DRDYI2 = Control{Name: "DRDYI2", Registro: &n0.INC6, Posicion: 0x10}
var TDTI2 = Control{Name: "TDTI2", Registro: &n0.INC6, Posicion: 0x04}
var WUFI2 = Control{Name: "WUFI2", Registro: &n0.INC6, Posicion: 0x02}
var TPI2 = Control{Name: "TPI2", Registro: &n0.INC6, Posicion: 0x01}

var TSC = Control{Name: "TSC", Registro: &n0.TILT_TIMER, Posicion: 0xFF}

var WUFC = Control{Name: "WUFC", Registro: &n0.WUFC, Posicion: 0xFF}

//-------------------------------------------------

var ATH = Control{Name: "ATH", Registro: &n0.ATH, Posicion: 0xFF}

var AVC = Control{Name: "AVC", Registro: &n0.LP_CNTL, Posicion: 0x70}

var SMP_TH = Control{Name: "SMP_TH", Registro: &n0.BUF_CNTL1, Posicion: 0x7F}

var BUFE = Control{Name: "BUFE", Registro: &n0.BUF_CNTL2, Posicion: 0x80}
var BRES = Control{Name: "BRES", Registro: &n0.BUF_CNTL2, Posicion: 0x40}
var BFIE = Control{Name: "BFIE", Registro: &n0.BUF_CNTL2, Posicion: 0x20}
var BUF_M = Control{Name: "BUF_M", Registro: &n0.BUF_CNTL2, Posicion: 0x03}

var SMP_LEV = Control{Name: "SMP_LEV", Registro: &n0.BUF_STATUS_1, Posicion: 0xFF}

var BUF_TRIG = Control{Name: "BUF_TRIG", Registro: &n0.BUF_STATUS_2, Posicion: 0x80}

var BUF_CLEAR = Control{Name: "BUF_CLEAR", Registro: &n0.BUF_CLEAR, Posicion: 0xFF}

var BUF_READ = Control{Name: "BUF_READ", Registro: &n0.BUF_READ, Posicion: 0xFF}
