package control

func Booleano(num byte) map[byte]any {
	return map[byte]any{num: 1, 0x00: 0}
}

// GESEL
const (
	Gsel_sens_2g byte = 0x00
	Gsel_sens_4g byte = 0x08
	Gsel_sens_8g byte = 0x10
)

var GSEL_C = map[byte]any{
	Gsel_sens_2g: 2,
	Gsel_sens_4g: 4,
	Gsel_sens_8g: 8,
}

// OWUF
const (
	OWUF_0_781HZ byte = 0x00
	OWUF_1_563HZ byte = 0x01
	OWUF_3_125HZ byte = 0x02
	OWUF_6_25HZ  byte = 0x03
	OWUF_12_5HZ  byte = 0x04
	OWUF_25HZ    byte = 0x05
	OWUF_50HZ    byte = 0x06
	OWUF_100HZ   byte = 0x07
)

var OWUF_C = map[byte]any{
	OWUF_0_781HZ: 0.781,
	OWUF_1_563HZ: 1.563,
	OWUF_3_125HZ: 3.125,
	OWUF_6_25HZ:  6.25,
	OWUF_12_5HZ:  12.5,
	OWUF_25HZ:    25,
	OWUF_50HZ:    50,
	OWUF_100HZ:   100,
}

// OSA
const (
	OSA_12_5HZ  byte = 0x00
	OSA_25HZ    byte = 0x01
	OSA_50HZ    byte = 0x02
	OSA_100HZ   byte = 0x03
	OSA_200HZ   byte = 0x04
	OSA_400HZ   byte = 0x05
	OSA_800HZ   byte = 0x06
	OSA_1600HZ  byte = 0x07
	OSA_0_781HZ byte = 0x08
	OSA_1_563HZ byte = 0x09
	OSA_3_125HZ byte = 0x0A
	OSA_6_25HZ  byte = 0x0B
)

var OSA_C = map[byte]any{
	OSA_12_5HZ:  12.5,
	OSA_25HZ:    25,
	OSA_50HZ:    50,
	OSA_100HZ:   100,
	OSA_200HZ:   200,
	OSA_400HZ:   400,
	OSA_800HZ:   800,
	OSA_1600HZ:  1600,
	OSA_0_781HZ: 0.781,
	OSA_1_563HZ: 1.563,
	OSA_3_125HZ: 3.125,
	OSA_6_25HZ:  6.25,
}

// AVC

const (
	AVC_0_samples   = 0x00
	AVC_2_samples   = 0x10
	AVC_4_samples   = 0x20
	AVC_8_samples   = 0x30
	AVC_16_samples  = 0x40
	AVC_32_samples  = 0x50
	AVC_64_samples  = 0x60
	AVC_128_samples = 0x70
)

var AVC_C = map[byte]any{
	AVC_0_samples:   0,
	AVC_2_samples:   2,
	AVC_4_samples:   4,
	AVC_8_samples:   8,
	AVC_16_samples:  16,
	AVC_32_samples:  32,
	AVC_64_samples:  64,
	AVC_128_samples: 128,
}

// BUF_M_C
const (
	BUF_M_FIFO = iota
	BUF_M_Stream
	BUF_M_Trigger
	BUF_M_FILO
)

var BUF_M_C = map[byte]any{

	BUF_M_FIFO:    0x00,
	BUF_M_Stream:  0x01,
	BUF_M_Trigger: 0x02,
	BUF_M_FILO:    0x03,
}
