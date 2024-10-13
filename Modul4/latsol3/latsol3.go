package main

import (
    "fmt"
    "math"
)

func main() {
    var x1, y1, x2, y2, x3, y3 float64
    fmt.Print("Masukkan koordinat titik A (x y): ")
    fmt.Scan(&x1, &y1)
    fmt.Print("Masukkan koordinat titik B (x y): ")
    fmt.Scan(&x2, &y2)
    fmt.Print("Masukkan koordinat titik C (x y): ")
    fmt.Scan(&x3, &y3)

    // Menghitung panjang sisi-sisi segitiga
    sisiAB := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
    sisiBC := math.Sqrt(math.Pow(x3-x2, 2) + math.Pow(y3-y2, 2))
    sisiCA := math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))

    // Menentukan sisi terpanjang
    sisiTerpanjang := sisiAB
    if sisiBC > sisiTerpanjang {
        sisiTerpanjang = sisiBC
    }
    if sisiCA > sisiTerpanjang {
        sisiTerpanjang = sisiCA
    }

    fmt.Printf("Panjang sisi terpanjang: %.2f\n", sisiTerpanjang)
}
