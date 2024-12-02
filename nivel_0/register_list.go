package accelemeter

var INT_REL = Register{Name: "INT_REL", Addr: 0x17}
var CNTL1 = Register{Name: "CNTL1", Addr: 0x18}
var CNTL2 = Register{Name: "CNTL2", Addr: 0x19}
var CNTL3 = Register{Name: "CNTL3", Addr: 0x1A}

var INC1 = Register{Name: "INC1", Addr: 0x1C}
var INC2 = Register{Name: "INC2", Addr: 0x1D}
var INC3 = Register{Name: "INC3", Addr: 0x1E}
var INC4 = Register{Name: "INC4", Addr: 0x1F}

var WUFC = Register{Name: "WUFC", Addr: 0x23}

var ATH = Register{Name: "ATH", Addr: 0x30}

var XHP_L = Register{Name: "XHP_L", Addr: 0x00}
var YHP_L = Register{Name: "YHP_L", Addr: 0x02}
var ZHP_L = Register{Name: "ZHP_L", Addr: 0x04}
var XHP_H = Register{Name: "XHP_H", Addr: 0x01}
var YHP_H = Register{Name: "YHP_H", Addr: 0x03}
var ZHP_H = Register{Name: "ZHP_H", Addr: 0x05}

var XHP = Register{Name: "XHP", Addr: 0x00}
var YHP = Register{Name: "YHP", Addr: 0x02}
var ZHP = Register{Name: "ZHP", Addr: 0x04}

var XOUT = Register{Name: "XOUT", Addr: 0x06}
var YOUT = Register{Name: "YOUT", Addr: 0x08}
var ZOUT = Register{Name: "ZOUT", Addr: 0x0A}

var ODCNTL = Register{Name: "ODCNTL", Addr: 0x1B}

var STATUS_REG = Register{Name: "STATUS_REG", Addr: 0x15}
var INS1 = Register{Name: "INS1", Addr: 0x12}
var INS2 = Register{Name: "INS2", Addr: 0x13}
var INS3 = Register{Name: "INS3", Addr: 0x14}
