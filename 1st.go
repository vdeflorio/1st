package main // questo il pacchetto main

import (
    "fmt"    // fmt e' general purpose
    "math"   // math e' per matematica (math.Pi, math.Pow, math.Cos ...)
             // NOTA: poi anche fare "math/rand"
)
func main() {
    fmt.Println("Circle Area calculation") // fmt contiene fmt.Print e fmt.Println
    fmt.Print("Enter a radius value: ")
    var radius float64                     // definisco radius come float64
    fmt.Scanf("%f", &radius)               // #clang
    area := math.Pi * math.Pow(radius,2)   // ':='   dichiara-e-inizializza
    fmt.Printf("Circle area with radius %.2f = %.2f \n",radius, area)  // fmt contiene la Printf !! #clang
}

