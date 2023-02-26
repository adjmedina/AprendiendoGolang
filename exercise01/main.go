// Paquete al que corresponde este directorio
package main

// importacion del las liberias requeridas entre comillas

// Funcion Principal o entrypoint
func main() {

	//Declaracion de variables requeridas
	var license bool = true
	var age int = 19

	if (age >= 15) && (license == true) {
		println("Puede Seguir avanzando")
	} else {
		println("No puede seguir circulando")
	}
	println("Programa terminado")
	//fmt.Printf("Probando una variable flotante %.2f", precio)

}
