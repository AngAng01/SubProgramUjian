package main

import "fmt"

func jumlahBilPositif() int {
	var total_2311102098, bilangan int
	fmt.Println("Masukkan bilangan (negatif untuk selesai): ")
	for {
		fmt.Scanln(&bilangan)
		if bilangan < 0 {
			break
		}
		if bilangan > 0 && bilangan%2== 0 {
			total_2311102098 += bilangan
		}
	}
	return total_2311102098
}

func main() {
	total_2311102098 := jumlahBilPositif()

	fmt.Println("Jumlah bilangan angka genap:", total_2311102098)
}
