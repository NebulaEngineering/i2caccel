package device

// KX023 representa una instancia del dispositivo KX023.
// Contiene la configuración predefinida para este dispositivo, incluyendo
// su nombre, dirección I2C y el bus al que está conectado.
var KX023 = Device{name: "KX023", adrr: 0x1E, bus: "2", gpio: []string{"GPIO27", "GPIO169"}}
