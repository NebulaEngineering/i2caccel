package main

import (
	"fmt"
	c "goi2caccel/control"
	c0 "goi2caccel/control/kx023"
	d "goi2caccel/device"
	r "goi2caccel/register"
)

func General(format string, reg r.Register, data []*c.Control) {
	fmt.Printf(format, reg.Read_Error())
	for _, item := range data {
		fmt.Print(item)
	}
	fmt.Println("")
}

func General1() {
	for j, i := range c0.Mregister {
		fmt.Printf("datos de %v (0x%02X): ", j.Name, j.Read_Error()[0])
		for _, k := range i {
			fmt.Print(k)
		}
		fmt.Println()
	}

	// fmt.Printf("datos de %v", reg.Read_Error())
	// for _, item := range data {
	// 	fmt.Print(item)
	// }
	// fmt.Println("")
}

func main() {

	d.KX023.Init()
	defer d.KX023.Close()

	c0.PC1.Previo(c0.PC1.Posicion + 1)
	//fmt.Println(&m.PC1)

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
