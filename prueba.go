package prueba

import (
	"flag"
	"fmt"
)

func Prueba() {
	// Definir banderas
	var (
		host  = flag.String("host", "localhost", "Dirección del servidor")
		port  = flag.Int("port", 8080, "Puerto del servidor")
		debug = flag.Bool("debug", false, "Habilitar modo depuración")
	)

	// Analizar las banderas
	flag.Parse()

	// Usar las banderas
	fmt.Printf("Conectando a %s:%d (debug: %v)\n", *host, *port, *debug)
}
