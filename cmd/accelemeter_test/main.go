package main

import (
	c "cli_acc/main/control"
	d "cli_acc/main/device"
	r "cli_acc/main/register"
	"fmt"
)

func main() {

	d.LT8640A.Init()

	fmt.Printf("datos de CNTL1: (%02X) ", r.CNTL1.Read_Error())
	for _, item := range []*c.Control{&c.PC1, &c.RES, &c.DRDYE, &c.GSEL, &c.WUFE, &c.TPE} {
		item.Interpretar()
	}
	fmt.Println("")

	fmt.Printf("datos de CNTL3: (%02X)", r.CNTL3.Read_Error())
	c.OWUF.Interpretar()
	fmt.Println("")

	fmt.Printf("datos de ODCNTL: (%02X)", r.ODCNTL.Read_Error())
	for _, item := range []*c.Control{&c.IIR_BYPASS, &c.LPR0, &c.OSA} {
		item.Interpretar()
	}
	fmt.Println("")

	fmt.Printf("datos de INC1: (%02X)", r.INC1.Read_Error())
	for _, item := range []*c.Control{&c.IEN, &c.IEA, &c.IEL, &c.STPOL, &c.SPI3E} {
		item.Interpretar()
	}
	fmt.Println("")

	fmt.Printf("datos de INC2: (%02X)", r.INC2.Read_Error())
	for _, item := range []*c.Control{&c.XNWUE, &c.XPWUE, &c.YNWUE, &c.YPWUE, &c.ZNWUE, &c.ZPWUE} {
		item.Interpretar()
	}
	fmt.Println("")

	fmt.Printf("datos de INC3: (%02X)", r.INC3.Read_Error())
	for _, item := range []*c.Control{&c.TLEM, &c.TRIM, &c.TDOM, &c.TUPM, &c.TFDM, &c.TFUM} {
		item.Interpretar()
	}

	fmt.Println("")

	fmt.Printf("datos de INC4: (%02X)", r.INC4.Read_Error())
	for _, item := range []*c.Control{&c.BFI1, &c.WMI1, &c.DRDYI1, &c.TDTI1, &c.WUFI1, &c.TPI1} {
		item.Interpretar()
	}
	fmt.Println("")

	fmt.Printf("datos de WUFC: (%02X)", r.WUFC.Read_Error())
	fmt.Println("")

	fmt.Printf("datos de ATH: (%02X)", r.ATH.Read_Error())
	fmt.Println("")

	fmt.Printf("datos de LP_CNTL: (%02X)", r.LP_CNTL.Read_Error())
	c.AVC.Interpretar()

	fmt.Println("")

}
