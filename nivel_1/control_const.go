package nivel1

type GSEL_C byte

const (
	Sensivity2g GSEL_C = 0x00
	Sensivity4g GSEL_C = 0x08
	Sensivity8g GSEL_C = 0x10
)

func (n GSEL_C) Mostrar() string {
	switch n {
	case Sensivity2g:
		return "Sensivity 2g"
	case Sensivity4g:
		return "Sensivity 4g"
	case Sensivity8g:
		return "Sensivity 8g"
	default:
		return "Unknown Sensitivity"
	}
}

type OSA_C byte

const (
	ODR_12_5HZ  OSA_C = 0x00
	ODR_25HZ    OSA_C = 0x01
	ODR_50HZ    OSA_C = 0x02
	ODR_100HZ   OSA_C = 0x03
	ODR_200HZ   OSA_C = 0x04
	ODR_400HZ   OSA_C = 0x05
	ODR_800HZ   OSA_C = 0x06
	ODR_1600HZ  OSA_C = 0x07
	ODR_0_781HZ OSA_C = 0x08
	ODR_1_563HZ OSA_C = 0x09
	ODR_3_125HZ OSA_C = 0x0A
	ODR_6_25HZ  OSA_C = 0x0B
)
