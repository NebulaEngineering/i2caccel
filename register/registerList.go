package register

import D "cli_acc/main/device"

var XHP_L = Register{Name: "XHP_L", addr: 0x00, device: &D.LT8640A}
var XHP_H = Register{Name: "XHP_H", addr: 0x01, device: &D.LT8640A}
var YHP_L = Register{Name: "YHP_L", addr: 0x02, device: &D.LT8640A}
var YHP_H = Register{Name: "YHP_H", addr: 0x03, device: &D.LT8640A}
var ZHP_L = Register{Name: "ZHP_L", addr: 0x04, device: &D.LT8640A}
var ZHP_H = Register{Name: "ZHP_H", addr: 0x05, device: &D.LT8640A}

var INT_REL = Register{Name: "INT_REL", addr: 0x17, device: &D.LT8640A}
var CNTL1 = Register{Name: "CNTL1", addr: 0x18, device: &D.LT8640A}
var CNTL2 = Register{Name: "CNTL2", addr: 0x19, device: &D.LT8640A}
var CNTL3 = Register{Name: "CNTL3", addr: 0x1A, device: &D.LT8640A}

var INC1 = Register{Name: "INC1", addr: 0x1C, device: &D.LT8640A}
var INC2 = Register{Name: "INC2", addr: 0x1D, device: &D.LT8640A}
var INC3 = Register{Name: "INC3", addr: 0x1E, device: &D.LT8640A}
var INC4 = Register{Name: "INC4", addr: 0x1F, device: &D.LT8640A}
var INC5 = Register{Name: "INC5", addr: 0x20, device: &D.LT8640A}
var INC6 = Register{Name: "INC6", addr: 0x21, device: &D.LT8640A}

var TILT_TIMER = Register{Name: "TILT_TIMER", addr: 0x22, device: &D.LT8640A}

var WUFC = Register{Name: "WUFC", addr: 0x23, device: &D.LT8640A}

var ATH = Register{Name: "ATH", addr: 0x30, device: &D.LT8640A}

var ODCNTL = Register{Name: "ODCNTL", addr: 0x1B, device: &D.LT8640A}

var STATUS_REG = Register{Name: "STATUS_REG", addr: 0x15, device: &D.LT8640A}
var INS1 = Register{Name: "INS1", addr: 0x12, device: &D.LT8640A}
var INS2 = Register{Name: "INS2", addr: 0x13, device: &D.LT8640A}
var INS3 = Register{Name: "INS3", addr: 0x14, device: &D.LT8640A}

var LP_CNTL = Register{Name: "LP_CNTL", addr: 0x35, device: &D.LT8640A}

var BUF_CNTL1 = Register{Name: "BUF_CNTL1", addr: 0x3A, device: &D.LT8640A}
var BUF_CNTL2 = Register{Name: "BUF_CNTL2", addr: 0x3B, device: &D.LT8640A}

var BUF_STATUS_1 = Register{Name: "BUF_STATUS_1", addr: 0x3C, device: &D.LT8640A}
var BUF_STATUS_2 = Register{Name: "BUF_STATUS_2", addr: 0x3D, device: &D.LT8640A}
var BUF_CLEAR = Register{Name: "BUF_CLEAR", addr: 0x3E, device: &D.LT8640A}
var BUF_READ = Register{Name: "BUF_READ", addr: 0x3F, device: &D.LT8640A}
