package register

import (
	d "i2caccel/device"
	r "i2caccel/register"
)

var XHPL = r.Register{Name: "XHPL", Addr: 0x00, Device: &d.KX023}
var XHPH = r.Register{Name: "XHPH", Addr: 0x01, Device: &d.KX023}
var YHPL = r.Register{Name: "YHPL", Addr: 0x02, Device: &d.KX023}
var YHPH = r.Register{Name: "YHPH", Addr: 0x03, Device: &d.KX023}
var ZHPL = r.Register{Name: "ZHPL", Addr: 0x04, Device: &d.KX023}
var ZHPH = r.Register{Name: "ZHPH", Addr: 0x05, Device: &d.KX023}

var XOUTL = r.Register{Name: "XOUTL", Addr: 0x06, Device: &d.KX023}
var XOUTH = r.Register{Name: "XOUTH", Addr: 0x07, Device: &d.KX023}
var YOUTL = r.Register{Name: "YOUTL", Addr: 0x08, Device: &d.KX023}
var YOUTH = r.Register{Name: "YOUTH", Addr: 0x09, Device: &d.KX023}
var ZOUTL = r.Register{Name: "ZOUTL", Addr: 0x0A, Device: &d.KX023}
var ZOUTH = r.Register{Name: "ZOUTH", Addr: 0x0B, Device: &d.KX023}

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
