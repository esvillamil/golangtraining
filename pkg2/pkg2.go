package pkg2

import (
"fmt"
"math/rand"
"time"
)

func init () {
fmt.Println("Subsistema␣2␣activado␣...")
rand.Seed(time.Now().Unix())
}

func Translate (cadena string) string {

runes := []rune(cadena)
rand.Shuffle(len( runes ), func( i , j int) {
runes [ i ], runes [ j ] = runes [ j ], runes [ i ]
})
return string( runes )
}
