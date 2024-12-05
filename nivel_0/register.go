package nivel0

import "fmt"

type Register struct {
	Name   string
	Addr   byte
	Value  byte
	Device *device
	Update bool
}

// solo tiene como argumento el valor del registro a modificar
func (r *Register) Write_Error(value []byte) {
	data := []byte{r.Addr}
	data = append(data, value...)
	r.Device.Write_Error(data)
}

func (r *Register) Read_Error() []byte {
	r.Value = r.Device.Read_Error([]byte{r.Addr})[0]
	return []byte{r.Value}
}

func (r *Register) Actualizar() {
	if r.Update {
		r.Write_Error([]byte{r.Value})
		r.Update = false
	} else {
		fmt.Println("no hay cambios")
	}
}
