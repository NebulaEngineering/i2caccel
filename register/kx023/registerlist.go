package register

import (
	d "goi2caccel/device"
	r "goi2caccel/register"
)

var XHP_L = r.Register{Name: "XHP_L", Addr: 0x00, Device: &d.KX023}
var XHP_H = r.Register{Name: "XHP_H", Addr: 0x01, Device: &d.KX023}
var YHP_L = r.Register{Name: "YHP_L", Addr: 0x02, Device: &d.KX023}
var YHP_H = r.Register{Name: "YHP_H", Addr: 0x03, Device: &d.KX023}
var ZHP_L = r.Register{Name: "ZHP_L", Addr: 0x04, Device: &d.KX023}
var ZHP_H = r.Register{Name: "ZHP_H", Addr: 0x05, Device: &d.KX023}

var INT_REL = r.Register{Name: "INT_REL", Addr: 0x17, Device: &d.KX023}
var CNTL1 = r.Register{Name: "CNTL1", Addr: 0x18, Device: &d.KX023}
var CNTL2 = r.Register{Name: "CNTL2", Addr: 0x19, Device: &d.KX023}
var CNTL3 = r.Register{Name: "CNTL3", Addr: 0x1A, Device: &d.KX023}

var INC1 = r.Register{Name: "INC1", Addr: 0x1C, Device: &d.KX023}
var INC2 = r.Register{Name: "INC2", Addr: 0x1D, Device: &d.KX023}
var INC3 = r.Register{Name: "INC3", Addr: 0x1E, Device: &d.KX023}
var INC4 = r.Register{Name: "INC4", Addr: 0x1F, Device: &d.KX023}
var INC5 = r.Register{Name: "INC5", Addr: 0x20, Device: &d.KX023}
var INC6 = r.Register{Name: "INC6", Addr: 0x21, Device: &d.KX023}

var TILT_TIMER = r.Register{Name: "TILT_TIMER", Addr: 0x22, Device: &d.KX023}

var WUFC = r.Register{Name: "WUFC", Addr: 0x23, Device: &d.KX023}

var ATH = r.Register{Name: "ATH", Addr: 0x30, Device: &d.KX023}

var ODCNTL = r.Register{Name: "ODCNTL", Addr: 0x1B, Device: &d.KX023}

var STATUS_REG = r.Register{Name: "STATUS_REG", Addr: 0x15, Device: &d.KX023}
var INS1 = r.Register{Name: "INS1", Addr: 0x12, Device: &d.KX023}
var INS2 = r.Register{Name: "INS2", Addr: 0x13, Device: &d.KX023}
var INS3 = r.Register{Name: "INS3", Addr: 0x14, Device: &d.KX023}

var LP_CNTL = r.Register{Name: "LP_CNTL", Addr: 0x35, Device: &d.KX023}

var BUF_CNTL1 = r.Register{Name: "BUF_CNTL1", Addr: 0x3A, Device: &d.KX023}
var BUF_CNTL2 = r.Register{Name: "BUF_CNTL2", Addr: 0x3B, Device: &d.KX023}

var BUF_STATUS_1 = r.Register{Name: "BUF_STATUS_1", Addr: 0x3C, Device: &d.KX023}
var BUF_STATUS_2 = r.Register{Name: "BUF_STATUS_2", Addr: 0x3D, Device: &d.KX023}
var BUF_CLEAR = r.Register{Name: "BUF_CLEAR", Addr: 0x3E, Device: &d.KX023}
var BUF_READ = r.Register{Name: "BUF_READ", Addr: 0x3F, Device: &d.KX023}
