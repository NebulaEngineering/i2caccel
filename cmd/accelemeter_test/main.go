package main

import (
	n0 "cli_acc/main/nivel_0"
	n1 "cli_acc/main/nivel_1"
	"fmt"
)

func main() {

	n0.LT8640A.Init()

	// fmt.Printf("datos de CNTL1: %02X\n", n0.CNTL1.Read_Error())
	// n1.PC1.Interpretar()
	// n1.RES.Interpretar()
	// n1.DRDYE.Interpretar()

	// n1.GSEL.Interpretar()
	// n1.GSEL.Previo(n1.Gsel_sens_8g)
	// n1.GSEL.Interpretar()

	// apagar sensor
	n0.CNTL1.Write_Error([]byte{0x82})

	n1.PC1.Interpretar()
	fmt.Printf(" CTNL1: 0x%02X\n", n0.CNTL1.Read_Error()[0])
	n0.CNTL1.Write_Error([]byte{n1.PC1.Posicion ^ n0.CNTL1.Read_Error()[0]})

	// encender sensor
	n1.PC1.Interpretar()
	fmt.Printf(" CTNL1: 0x%02X\n", n0.CNTL1.Read_Error()[0])
	n0.CNTL1.Write_Error([]byte{n1.PC1.Posicion | n0.CNTL1.Read_Error()[0]})

	n1.PC1.Interpretar()
	fmt.Printf(" CTNL1: 0x%02X\n", n0.CNTL1.Read_Error()[0])

	// n1.TDTE.Interpretar()
	// n1.WUFE.Interpretar()
	// n1.TPE.Interpretar()

	// fmt.Printf("datos de CNTL3: %02X\n", n0.CNTL3.Read_Error())
	// n1.OTP.Interpretar()
	// n1.OTDT.Interpretar()
	// n1.OWUF.Interpretar()

	// fmt.Printf("datos de ODCNTL: %02X\n", n0.ODCNTL.Read_Error())
	// n1.IIR_BYPASS.Interpretar()
	// n1.LPR0.Interpretar()
	// n1.OSA.Interpretar()

	// nivel0.CNTL1.Write_Error([]byte{0x82})
	// fmt.Printf("%2X\n", nivel0.CNTL1.Read_Error()[0])

	// n1.PC1.Interpretar()

}
