package main

import (
	"fmt"
)

func main() {
	var a, b, h float64

	fmt.Print("Masukkan panjang sisi a: ")
	fmt.Scan(&a)

	fmt.Print("Masukkan panjang sisi b: ")
	fmt.Scan(&b)

	fmt.Print("Masukkan tinggi h: ")
	fmt.Scan(&h)

	luas := 0.5 * (a + b) * h

	fmt.Printf("Luas trapesium adalah: %.2f\n", luas)
}
