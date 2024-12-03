package nivel0

type Register struct {
	Name   string
	Addr   byte
	Value  byte
	Device *device
	Update bool
}

func (r *Register) Write_Error(value []byte) {
	data := []byte{r.Addr}
	data = append(data, value...)
	r.Device.Write_Error(data)
}

func (r *Register) Read_Error() []byte {
	return r.Device.Read_Error([]byte{r.Addr})
}
