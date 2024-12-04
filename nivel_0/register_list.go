package nivel0

var XHP_L = Register{Name: "XHP_L", Addr: 0x00, Device: &LT8640A}
var XHP_H = Register{Name: "XHP_H", Addr: 0x01, Device: &LT8640A}
var YHP_L = Register{Name: "YHP_L", Addr: 0x02, Device: &LT8640A}
var YHP_H = Register{Name: "YHP_H", Addr: 0x03, Device: &LT8640A}
var ZHP_L = Register{Name: "ZHP_L", Addr: 0x04, Device: &LT8640A}
var ZHP_H = Register{Name: "ZHP_H", Addr: 0x05, Device: &LT8640A}

var INT_REL = Register{Name: "INT_REL", Addr: 0x17, Device: &LT8640A}
var CNTL1 = Register{Name: "CNTL1", Addr: 0x18, Device: &LT8640A}
var CNTL2 = Register{Name: "CNTL2", Addr: 0x19, Device: &LT8640A}
var CNTL3 = Register{Name: "CNTL3", Addr: 0x1A, Device: &LT8640A}

var INC1 = Register{Name: "INC1", Addr: 0x1C, Device: &LT8640A}
var INC2 = Register{Name: "INC2", Addr: 0x1D, Device: &LT8640A}
var INC3 = Register{Name: "INC3", Addr: 0x1E, Device: &LT8640A}
var INC4 = Register{Name: "INC4", Addr: 0x1F, Device: &LT8640A}
var INC5 = Register{Name: "INC5", Addr: 0x20, Device: &LT8640A}
var INC6 = Register{Name: "INC6", Addr: 0x21, Device: &LT8640A}

var TILT_TIMER = Register{Name: "TILT_TIMER", Addr: 0x22, Device: &LT8640A}

var WUFC = Register{Name: "WUFC", Addr: 0x23, Device: &LT8640A}

var ATH = Register{Name: "ATH", Addr: 0x30, Device: &LT8640A}

var ODCNTL = Register{Name: "ODCNTL", Addr: 0x1B, Device: &LT8640A}

var STATUS_REG = Register{Name: "STATUS_REG", Addr: 0x15, Device: &LT8640A}
var INS1 = Register{Name: "INS1", Addr: 0x12, Device: &LT8640A}
var INS2 = Register{Name: "INS2", Addr: 0x13, Device: &LT8640A}
var INS3 = Register{Name: "INS3", Addr: 0x14, Device: &LT8640A}

var LP_CNTL = Register{Name: "LP_CNTL", Addr: 0x35, Device: &LT8640A}

var BUF_CNTL1 = Register{Name: "BUF_CNTL1", Addr: 0x3A, Device: &LT8640A}
var BUF_CNTL2 = Register{Name: "BUF_CNTL2", Addr: 0x3B, Device: &LT8640A}

var BUF_STATUS_1 = Register{Name: "BUF_STATUS_1", Addr: 0x3C, Device: &LT8640A}
var BUF_STATUS_2 = Register{Name: "BUF_STATUS_2", Addr: 0x3D, Device: &LT8640A}
var BUF_CLEAR = Register{Name: "BUF_CLEAR", Addr: 0x3E, Device: &LT8640A}
var BUF_READ = Register{Name: "BUF_READ", Addr: 0x3F, Device: &LT8640A}
