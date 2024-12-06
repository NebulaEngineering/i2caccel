package register

import (
	d "cli_acc/main/device"
)

type Register struct {
	name   string
	addr   byte
	Value  byte
	device *d.Device
	Update bool
}

func (r *Register) Write_Error(value []byte) {
	data := []byte{r.addr}
	data = append(data, value...)
	r.device.Write_Error(data)
}

func (r *Register) Write(value []byte) error {
	data := []byte{r.addr}
	data = append(data, value...)
	if err := r.device.Write(data); err != nil {
		return err
	}
	return nil
}

func (r *Register) Read_Error() []byte {
	r.Value = r.device.Read_Error([]byte{r.addr})[0]
	return []byte{r.Value}
}

func (r *Register) Read() ([]byte, error) {
	if data, err := r.device.Read([]byte{r.addr}); err != nil {
		r.Value = data[0]
		return nil, err
	}
	return []byte{r.Value}, nil
}
