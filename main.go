// Ejercicio de inicialización de paquetes
package main

import (
"fmt"

"github.com/esvillamil/golangtraining/pkg2"
)

import _ "github.com/esvillamil/golangtraining/holamundo"

func init () {
fmt.Println("Ready to launch ...")
}


func main() {
saludo := "Liftoff!"
fmt.Println( saludo )
fmt.Println(pkg2.Translate ("Arecibo␣message␣sent"))
}