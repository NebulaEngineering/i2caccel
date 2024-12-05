package main

import (
	n0 "cli_acc/main/nivel_0"
	n1 "cli_acc/main/nivel_1"
	"fmt"
)

func main() {

	n0.LT8640A.Init()

	fmt.Printf("datos de CNTL1: (%02X) ", n0.CNTL1.Read_Error())
	for _, item := range []*n1.Control{&n1.PC1, &n1.RES, &n1.DRDYE, &n1.GSEL, &n1.WUFE, &n1.TPE} {
		item.Interpretar()
	}
	fmt.Println("")

	fmt.Printf("datos de CNTL3: (%02X)", n0.CNTL3.Read_Error())
	n1.OWUF.Interpretar()
	fmt.Println("")

	fmt.Printf("datos de ODCNTL: (%02X)", n0.ODCNTL.Read_Error())
	for _, item := range []*n1.Control{&n1.IIR_BYPASS, &n1.LPR0, &n1.OSA} {
		item.Interpretar()
	}
	fmt.Println("")

	fmt.Printf("datos de INC1: (%02X)", n0.INC1.Read_Error())
	for _, item := range []*n1.Control{&n1.IEN, &n1.IEA, &n1.IEL, &n1.STPOL, &n1.SPI3E} {
		item.Interpretar()
	}
	fmt.Println("")

	fmt.Printf("datos de INC2: (%02X)", n0.INC2.Read_Error())
	for _, item := range []*n1.Control{&n1.XNWUE, &n1.XPWUE, &n1.YNWUE, &n1.YPWUE, &n1.ZNWUE, &n1.ZPWUE} {
		item.Interpretar()
	}
	fmt.Println("")

	fmt.Printf("datos de INC3: (%02X)", n0.INC3.Read_Error())
	for _, item := range []*n1.Control{&n1.TLEM, &n1.TRIM, &n1.TDOM, &n1.TUPM, &n1.TFDM, &n1.TFUM} {
		item.Interpretar()
	}

	fmt.Println("")

	fmt.Printf("datos de INC4: (%02X)", n0.INC4.Read_Error())
	for _, item := range []*n1.Control{&n1.BFI1, &n1.WMI1, &n1.DRDYI1, &n1.TDTI1, &n1.WUFI1, &n1.TPI1} {
		item.Interpretar()
	}
	fmt.Println("")

	fmt.Printf("datos de WUFC: (%02X)", n0.WUFC.Read_Error())
	fmt.Println("")

	fmt.Printf("datos de ATH: (%02X)", n0.ATH.Read_Error())
	fmt.Println("")

	fmt.Printf("datos de LP_CNTL: (%02X)", n0.LP_CNTL.Read_Error())
	n1.AVC.Interpretar()

	fmt.Println("")

}
