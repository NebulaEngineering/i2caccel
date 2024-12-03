package main

import (
	nivel0 "cli_acc/main/nivel_0"
	"fmt"
)

func main() {
	nivel0.LT8640A.Init()

	nivel0.CNTL1.Write_Error([]byte{0x02})
	fmt.Println(nivel0.CNTL1.Read_Error()[0])
}
