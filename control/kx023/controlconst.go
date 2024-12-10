package control

import c "goi2caccel/control"

// Booleano genera un mapa para valores booleanos, donde la clave 'hex' representa el bit
// activado y 0x00 el bit desactivado, simplificando la interpretación de registros
// con valores booleanos.
func Booleano(hex byte) map[byte]c.MappableValue {
	return map[byte]c.MappableValue{hex: c.ByteValue(1), 0x00: c.ByteValue(0)}
}

// Constantes que representan los diferentes rangos de sensibilidad del acelerómetro.
const (
	Gsel_sens_2g byte = 0x00
	Gsel_sens_4g byte = 0x08
	Gsel_sens_8g byte = 0x10
)

// GSEL_C mapea los valores de los registros a los rangos de sensibilidad en g.
var GSEL_C = map[byte]c.MappableValue{
	Gsel_sens_2g: c.ByteValue(2),
	Gsel_sens_4g: c.ByteValue(4),
	Gsel_sens_8g: c.ByteValue(8),
}

// Constantes que representan las diferentes frecuencias de Motion Detect.
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

// Mapea los valores de los registros a las frecuencias de Motion Detect.
var OWUF_C = map[byte]c.MappableValue{
	OWUF_0_781HZ: c.Float32Value(0.781),
	OWUF_1_563HZ: c.Float32Value(1.563),
	OWUF_3_125HZ: c.Float32Value(3.125),
	OWUF_6_25HZ:  c.Float32Value(6.25),
	OWUF_12_5HZ:  c.Float32Value(12.5),
	OWUF_25HZ:    c.Float32Value(25),
	OWUF_50HZ:    c.Float32Value(50),
	OWUF_100HZ:   c.Float32Value(100),
}

// Constantes que representan las diferentes tasas de muestreo de salida.
// Define las constantes para las diferentes tasas de muestreo de salida (Output Sample Rate) del acelerómetro.
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

// Mapea los valores de los registros a las tasas de muestreo de salida.
var OSA_C = map[byte]c.MappableValue{
	OSA_12_5HZ:  c.Float32Value(12.5),
	OSA_25HZ:    c.Float32Value(25),
	OSA_50HZ:    c.Float32Value(50),
	OSA_100HZ:   c.Float32Value(100),
	OSA_200HZ:   c.Float32Value(200),
	OSA_400HZ:   c.Float32Value(400),
	OSA_800HZ:   c.Float32Value(800),
	OSA_1600HZ:  c.Float32Value(1600),
	OSA_0_781HZ: c.Float32Value(0.781),
	OSA_1_563HZ: c.Float32Value(1.563),
	OSA_3_125HZ: c.Float32Value(3.125),
	OSA_6_25HZ:  c.Float32Value(6.25),
}

// Constantes que representan el control de muestras promediado.
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

// Mapea los valores de los registros para el control de muestras promediado.
var AVC_C = map[byte]c.MappableValue{
	AVC_0_samples:   c.ByteValue(0),
	AVC_2_samples:   c.ByteValue(2),
	AVC_4_samples:   c.ByteValue(4),
	AVC_8_samples:   c.ByteValue(8),
	AVC_16_samples:  c.ByteValue(16),
	AVC_32_samples:  c.ByteValue(32),
	AVC_64_samples:  c.ByteValue(64),
	AVC_128_samples: c.ByteValue(128),
}

// Constantes que representan los diferentes modo de operacion del buffer.
const (
	BUF_M_FIFO = iota
	BUF_M_Stream
	BUF_M_Trigger
	BUF_M_FILO
)

// BUF_M_C mapea los modo de operacion del buffer.
var BUF_M_C = map[byte]c.MappableValue{
	BUF_M_FIFO:    c.StringValue("FIFO"),
	BUF_M_Stream:  c.StringValue("Stream"),
	BUF_M_Trigger: c.StringValue("Trigger"),
	BUF_M_FILO:    c.StringValue("FILO"),
}
