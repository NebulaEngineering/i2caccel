package main

import (
	c "cli_acc/main/control"
	d "cli_acc/main/device"
	r "cli_acc/main/register"
	"fmt"
)

func General(format string, reg r.Register, data []*c.Control) {
	fmt.Printf(format, reg.Read_Error())
	for _, item := range data {
		fmt.Print(item)
	}
	fmt.Println("")
}

func General1() {
	for j, i := range c.Mregister {
		fmt.Printf("datos de %v (0x%02X): %v\n", j.Name, j.Read_Error()[0], i)

	}

	// fmt.Printf("datos de %v", reg.Read_Error())
	// for _, item := range data {
	// 	fmt.Print(item)
	// }
	// fmt.Println("")
}

func Previo() {

}

func main() {

	d.LT8640A.Init()
	defer d.LT8640A.Close()

	General1()

	// General("datos de CNTL1: (%02X) ", r.CNTL1, []*c.Control{&c.PC1, &c.RES, &c.DRDYE, &c.GSEL, &c.WUFE, &c.TPE})
	// General("datos de CNTL3: (%02X) ", r.CNTL3, []*c.Control{&c.OWUF})
	// General("datos de ODCNTL: (%02X) ", r.ODCNTL, []*c.Control{&c.IIR_BYPASS, &c.LPR0, &c.OSA})
	// General("datos de INC1: (%02X) ", r.INC1, []*c.Control{&c.IEN, &c.IEA, &c.IEL, &c.STPOL, &c.SPI3E})
	// General("datos de INC2: (%02X) ", r.INC2, []*c.Control{&c.XNWUE, &c.XPWUE, &c.YNWUE, &c.YPWUE, &c.ZNWUE, &c.ZPWUE})
	// General("datos de INC3: (%02X) ", r.INC3, []*c.Control{&c.TLEM, &c.TRIM, &c.TDOM, &c.TUPM, &c.TFDM, &c.TFUM})
	// General("datos de INC4: (%02X) ", r.INC4, []*c.Control{&c.BFI1, &c.WMI1, &c.DRDYI1, &c.TDTI1, &c.WUFI1, &c.TPI1})
	// fmt.Printf("datos de WUFC: (%02X)\n", r.WUFC.Read_Error())
	// fmt.Printf("datos de ATH: (%02X)\n", r.ATH.Read_Error())
	// General("datos de LP_CNTL:: (%02X) ", r.LP_CNTL, []*c.Control{&c.AVC})

}
