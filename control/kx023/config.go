package control

import (
	c "i2caccel/control"
	r0 "i2caccel/register/kx023"
)

// configuracion por defecto para un proposito especifico.
func Config(flagLPRO *int, flagIIR_BYPASS *int, flagOSA *float64, flagOWUF *float64, flagAVC *int) {
	// apagar sensor
	PC1.Previo(c.ByteValue(0))
	r0.CNTL1.Actualizar()

	//INC1
	for i, j := range []c.ByteValue{
		1, //IEN
		1, //IEA
		0, //IEL
		0, //STPOL
		0, //SPI3E
	} {
		Mregister[&r0.INC1][i].Previo(j)
	}

	//INC2
	for i, j := range []c.ByteValue{
		1, //XNWUE
		1, //XPWUE
		1, //YNWUE
		1, //YPWUE
		1, //ZNWUE
		1, //ZPWUE
	} {
		Mregister[&r0.INC2][i].Previo(j)
	}

	//INC3
	for i, j := range []c.ByteValue{
		1, //TLEM
		1, //TRIM
		1, //TDOM
		1, //TUPM
		1, //TFDM
		1, //TFUM
	} {
		Mregister[&r0.INC3][i].Previo(j)
	}

	//INC4
	for i, j := range []c.ByteValue{
		0, //BFI1
		0, //WMI1
		1, //DRDYI1
		0, //TDTI1
		0, //WUFI1
		0, //TPI1
	} {
		Mregister[&r0.INC4][i].Previo(j)
	}

	//INC5
	for i, j := range []c.ByteValue{
		0, //IEN2
		0, //IEA2
		0, //IEL2
	} {
		Mregister[&r0.INC5][i].Previo(j)
	}

	//INC6
	for i, j := range []c.ByteValue{
		0, //BFI2
		0, //WMI2
		0, //DRDYI2
		0, //TDTI2
		0, //WUFI2
		0, //TPI2
	} {
		Mregister[&r0.INC6][i].Previo(j)
	}

	//ODCTNL
	for i, j := range []c.MappableValue{
		c.ByteValue(*flagIIR_BYPASS), //IIR_BYPASS
		c.ByteValue(*flagLPRO),       //LPR0
		c.Float32Value(*flagOSA),     //OSA (11 variantes) (0.781|1.563|3.125|6.25|12.5|25|50|100|200|400|800|1600)
	} {
		Mregister[&r0.ODCNTL][i].Previo(j)
	}

	//CNTL1
	for i, j := range []c.ByteValue{
		1, //PC1
		1, //RES
		1, //DRDY
		2, //GSEL (3 variantes)
		0, //WUFE
		0, //TPE
	} {
		Mregister[&r0.CNTL1][i].Previo(j)
	}

	//CNTL3
	for i, j := range []c.Float32Value{
		c.Float32Value(*flagOWUF), //OWUF (5 variantes)
	} {
		Mregister[&r0.CNTL3][i].Previo(j)
	}

	//ATH
	for i, j := range []c.ByteValue{
		8, //ATH
	} {
		Mregister[&r0.ATH][i].Previo(j)
	}

	//WUFC
	for i, j := range []c.ByteValue{
		15, //WUFC
	} {
		Mregister[&r0.WUFC][i].Previo(j)
	}

	//LP_CNTL
	for i, j := range []c.ByteValue{
		c.ByteValue(*flagAVC), //LP_CNTL (8 variantes)
	} {
		Mregister[&r0.LP_CNTL][i].Previo(j)
	}

	//BUF_CNTL1
	for i, j := range []c.ByteValue{
		8, //SMP_TH
	} {
		Mregister[&r0.BUF_CNTL1][i].Previo(j)
	}

	//BUF_CNTL2
	for i, j := range []c.MappableValue{
		c.ByteValue(0),        //BUFE
		c.ByteValue(0),        //BRES
		c.ByteValue(0),        //BFIE
		c.StringValue("FIFO"), //BUF_M (4 variantes)
	} {
		Mregister[&r0.BUF_CNTL2][i].Previo(j)
	}

	// actualizar todo
	for _, i := range Seq_Register {
		i.Actualizar()
		for _, j := range Mregister[i] {
			j.Update = false
		}
	}

	// encender sensor
	PC1.Previo(c.ByteValue(1))
	r0.CNTL1.Actualizar()

}
