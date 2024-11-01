package main

import (
	"fmt"
)

func hitungBiaya(jumlahMenu, jumlahOrang int, adaSisa bool) int {
	var biaya_2311102098 int

	if jumlahMenu >= 10 {
		biaya_2311102098 = 150000
	} else if jumlahMenu >= 5 {
		biaya_2311102098 = jumlahMenu * 10000 + 75000
	} else {
		biaya_2311102098 = jumlahMenu * 15000
	}

	if adaSisa == true {
		biaya_2311102098 = int(float64(biaya_2311102098) * 0.20) * jumlahOrang 
	}
	return biaya_2311102098
}

func main() {
	var M int
	fmt.Print("Masukkan jumlah rombongan: ")
	fmt.Scanln(&M)

	for i := 1; i <= M; i++ {
		var jumlahMenu, jumlahOrang int
		var adaSisa bool
		fmt.Print("Masukkan jumlah jenis minuman, banyak orang, dan apakah ada sisa(false/true): ")
		fmt.Scan(&jumlahMenu, &jumlahOrang, &adaSisa)

		fmt.Printf("Total biaya untuk rombongan ke-%d: Rp %d\n", i, hitungBiaya(jumlahMenu, jumlahOrang, adaSisa))
	}
}
