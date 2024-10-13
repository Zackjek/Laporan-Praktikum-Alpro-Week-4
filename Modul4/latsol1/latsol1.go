package main

import "fmt"

func main() {
    var totalBelanja, diskon, totalAkhir int
    fmt.Print("Masukkan Total Belanja Awal (dalam rupiah): ")
    fmt.Scan(&totalBelanja)
    fmt.Print("Masukkan Besarnya Diskon (dalam persen): ")
    fmt.Scan(&diskon)
    
    totalAkhir = totalBelanja - (totalBelanja * diskon / 100)
    
    fmt.Printf("Total Belanja Akhir setelah diskon: %d rupiah\n", totalAkhir)
}
