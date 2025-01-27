// Paquete al que corresponde el conjunto de instrucciones
package main

// Todos los nombres de paquetes que se importan en el código deberán ir entre comillas dobles
// Si un paquete aue se importe no se utiliza, el código no compilará
import (
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"
)

// obtenerTipoDeCambio obtiene el tipo de cambio de USD a MXN
// Entradas: ninguna
// Salidas: string con el tipo de cambio o error
func obtenerTipoDeCambio() (string, error) {
	client := resty.New()
	client.SetTimeout(10 * time.Second)
	resp, err := client.R().Get("https://api.exchangeratesapi.io/latest?base=USD&symbols=MXN")
	if err != nil {
		return "", err
	}

	return resp.String(), nil
}

// Función Principal o entrypoint

func main() {
	resultado, err := obtenerTipoDeCambio()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("El producto tiene un valor de $s  al valor de cambio del dia de hoy", resultado)
}
