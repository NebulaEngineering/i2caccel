package nivel0

var INT_REL = Register{Name: "INT_REL", Addr: 0x17, Device: &LT8640A}
var CNTL1 = Register{Name: "CNTL1", Addr: 0x18, Device: &LT8640A}
var CNTL2 = Register{Name: "CNTL2", Addr: 0x19, Device: &LT8640A}
var CNTL3 = Register{Name: "CNTL3", Addr: 0x1A, Device: &LT8640A}

var INC1 = Register{Name: "INC1", Addr: 0x1C, Device: &LT8640A}
var INC2 = Register{Name: "INC2", Addr: 0x1D, Device: &LT8640A}
var INC3 = Register{Name: "INC3", Addr: 0x1E, Device: &LT8640A}
var INC4 = Register{Name: "INC4", Addr: 0x1F, Device: &LT8640A}

var WUFC = Register{Name: "WUFC", Addr: 0x23, Device: &LT8640A}

var ATH = Register{Name: "ATH", Addr: 0x30, Device: &LT8640A}

var XHP_L = Register{Name: "XHP_L", Addr: 0x00, Device: &LT8640A}
var YHP_L = Register{Name: "YHP_L", Addr: 0x02, Device: &LT8640A}
var ZHP_L = Register{Name: "ZHP_L", Addr: 0x04, Device: &LT8640A}
var XHP_H = Register{Name: "XHP_H", Addr: 0x01, Device: &LT8640A}
var YHP_H = Register{Name: "YHP_H", Addr: 0x03, Device: &LT8640A}
var ZHP_H = Register{Name: "ZHP_H", Addr: 0x05, Device: &LT8640A}

var XHP = Register{Name: "XHP", Addr: 0x00, Device: &LT8640A}
var YHP = Register{Name: "YHP", Addr: 0x02, Device: &LT8640A}
var ZHP = Register{Name: "ZHP", Addr: 0x04, Device: &LT8640A}

var XOUT = Register{Name: "XOUT", Addr: 0x06, Device: &LT8640A}
var YOUT = Register{Name: "YOUT", Addr: 0x08, Device: &LT8640A}
var ZOUT = Register{Name: "ZOUT", Addr: 0x0A, Device: &LT8640A}

var ODCNTL = Register{Name: "ODCNTL", Addr: 0x1B, Device: &LT8640A}

var STATUS_REG = Register{Name: "STATUS_REG", Addr: 0x15, Device: &LT8640A}
var INS1 = Register{Name: "INS1", Addr: 0x12, Device: &LT8640A}
var INS2 = Register{Name: "INS2", Addr: 0x13, Device: &LT8640A}
var INS3 = Register{Name: "INS3", Addr: 0x14, Device: &LT8640A}
