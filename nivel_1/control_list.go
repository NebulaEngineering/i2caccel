package nivel1

import n0 "cli_acc/main/nivel_0"

var PC1 = control{name: "PC1", registro: &n0.CNTL1, posicion: 0x80}
var RES = control{name: "RES", registro: &n0.CNTL1, posicion: 0x40}
var DRDYE = control{name: "DRDYE", registro: &n0.CNTL1, posicion: 0x20}
var GSEL = control{name: "GSEL", registro: &n0.CNTL1, posicion: 0x18}
var TDTE = control{name: "TDTE", registro: &n0.CNTL1, posicion: 0x04}
var WUFE = control{name: "WUFE", registro: &n0.CNTL1, posicion: 0x02}
var TPE = control{name: "TPE", registro: &n0.CNTL1, posicion: 0x01}

var SRST = control{name: "SRST", registro: &n0.CNTL2, posicion: 0x80}
var COTC = control{name: "COTC", registro: &n0.CNTL2, posicion: 0x40}
var LEM = control{name: "LEM", registro: &n0.CNTL2, posicion: 0x20}
var RIM = control{name: "RIM", registro: &n0.CNTL2, posicion: 0x10}
var DOM = control{name: "DOM", registro: &n0.CNTL2, posicion: 0x08}
var UPM = control{name: "UPM", registro: &n0.CNTL2, posicion: 0x04}
var FDM = control{name: "FDM", registro: &n0.CNTL2, posicion: 0x02}
var FUM = control{name: "FUM", registro: &n0.CNTL2, posicion: 0x01}

var OTP = control{name: "OTP", registro: &n0.CNTL3, posicion: 0xC0}
var OTDT = control{name: "OTDT", registro: &n0.CNTL3, posicion: 0x38}
var OWUF = control{name: "OWUF", registro: &n0.CNTL3, posicion: 0x07}

var IIR_BYPASS = control{name: "IIR_BYPASS", registro: &n0.ODCNTL, posicion: 0x80}
var LPR0 = control{name: "LPR0", registro: &n0.ODCNTL, posicion: 0x40}
var OSA = control{name: "OSA", registro: &n0.ODCNTL, posicion: 0x0F}

var IEN = control{name: "IEN", registro: &n0.INC1, posicion: 0x20}
var IEA = control{name: "IEA", registro: &n0.INC1, posicion: 0x10}
var IEL = control{name: "IEL", registro: &n0.INC1, posicion: 0x08}
var STPOL = control{name: "STPOL", registro: &n0.INC1, posicion: 0x02}
var SPI3E = control{name: "SPI3E", registro: &n0.INC1, posicion: 0x01}

var XNWUE = control{name: "XNWUE", registro: &n0.INC2, posicion: 0x20}
var XPWUE = control{name: "XPWUE", registro: &n0.INC2, posicion: 0x10}
var YNWUE = control{name: "YNWUE", registro: &n0.INC2, posicion: 0x08}
var YPWUE = control{name: "YPWUE", registro: &n0.INC2, posicion: 0x04}
var ZNWUE = control{name: "ZNWUE", registro: &n0.INC2, posicion: 0x02}
var ZPWUE = control{name: "ZPWUE", registro: &n0.INC2, posicion: 0x01}

var TLEM = control{name: "TLEM", registro: &n0.INC3, posicion: 0x20}
var TRIM = control{name: "TRIM", registro: &n0.INC3, posicion: 0x10}
var TDOM = control{name: "TDOM", registro: &n0.INC3, posicion: 0x08}
var TUPM = control{name: "TUPM", registro: &n0.INC3, posicion: 0x04}
var TFDM = control{name: "TFDM", registro: &n0.INC3, posicion: 0x02}
var TFUM = control{name: "TFUM", registro: &n0.INC3, posicion: 0x01}

var BFI1 = control{name: "BFI1", registro: &n0.INC4, posicion: 0x40}
var WMI1 = control{name: "WMI1", registro: &n0.INC4, posicion: 0x20}
var DRDYI1 = control{name: "DRDYI1", registro: &n0.INC4, posicion: 0x10}
var TDTI1 = control{name: "TDTI1", registro: &n0.INC4, posicion: 0x04}
var WUFI1 = control{name: "WUFI1", registro: &n0.INC4, posicion: 0x02}
var TPI1 = control{name: "TPI1", registro: &n0.INC4, posicion: 0x01}

var IEN2 = control{name: "IEN2", registro: &n0.INC5, posicion: 0x20}
var IEA2 = control{name: "TPI1", registro: &n0.INC5, posicion: 0x10}
var IEL2 = control{name: "TPI1", registro: &n0.INC5, posicion: 0x08}

var BFI2 = control{name: "BFI2", registro: &n0.INC6, posicion: 0x40}
var WMI2 = control{name: "WMI2", registro: &n0.INC6, posicion: 0x20}
var DRDYI2 = control{name: "DRDYI2", registro: &n0.INC6, posicion: 0x10}
var TDTI2 = control{name: "TDTI2", registro: &n0.INC6, posicion: 0x04}
var WUFI2 = control{name: "WUFI2", registro: &n0.INC6, posicion: 0x02}
var TPI2 = control{name: "TPI2", registro: &n0.INC6, posicion: 0x01}

var TSC = control{name: "TSC", registro: &n0.TILT_TIMER, posicion: 0xFF}

var WUFC = control{name: "WUFC", registro: &n0.WUFC, posicion: 0xFF}

//-------------------------------------------------

var ATH = control{name: "ATH", registro: &n0.ATH, posicion: 0xFF}

var AVC = control{name: "AVC", registro: &n0.LP_CNTL, posicion: 0x70}

var SMP_TH = control{name: "SMP_TH", registro: &n0.BUF_CNTL1, posicion: 0x7F}

var BUFE = control{name: "BUFE", registro: &n0.BUF_CNTL2, posicion: 0x7F}
