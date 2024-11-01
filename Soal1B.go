package main

import "fmt"


func cekVoucher(noTiket_2311102098 int) bool {
	var digitPertama, digitKedua, digitAkhir1, digitAkhir2, digitTengah1, digitTengah2 int
	panjang := 0
	temp := noTiket_2311102098
	for temp > 0 {
		temp /= 10
		panjang++
	}

	if panjang != 5 && panjang != 6 {
		return false
	}

	if digitPertama*digitKedua != digitAkhir1*digitAkhir2 {
		return false
	}

	if digitTengah1 != digitTengah2 {
		return false
	}

	return true
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scanln(&N)

	jumlahValid := 0
	jumlahTidakValid := 0

	for i := 1; i <= N; i++ {
		var noTiket_2311102098 int
		fmt.Printf("Masukkan nomor seri voucher mahasiswa : ")
		fmt.Scanln(&noTiket_2311102098)

		if cekVoucher(noTiket_2311102098) {
			jumlahValid++
		} else {
			jumlahTidakValid++
		}
	}

	fmt.Println("Jumlah voucher valid:", jumlahValid)
	fmt.Println("Jumlah voucher tidak valid:", jumlahTidakValid)
}
