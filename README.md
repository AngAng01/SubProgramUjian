﻿# Sub Program Ujian
 ## SOAL 1
   Program di atas adalah program yang digunakan untuk memverifikasi nomor seri voucher berdasarkan beberapa aturan tertentu. Program ini meminta input berupa nomor seri voucher dari sejumlah mahasiswa, lalu menghitung jumlah voucher yang valid dan tidak valid.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Fungsi 'cekVoucher', Memeriksa validitas nomor voucher berdasarkan panjang digit (5 atau 6 digit) dan beberapa aturan perbandingan digit (belum sepenuhnya diimplementasikan).
      - Perhitungan Panjang Voucher, Menghitung jumlah digit voucher dengan membagi nomor secara bertahap.
      - Penghitung Voucher, Menghitung jumlah voucher valid dan tidak valid, lalu menampilkannya.
      
   ## Code Explanation
      ```go
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
      ```
   Kode di atas adalah fungsi cekVoucher dalam Go untuk memvalidasi nomor voucher. Fungsi ini menyiapkan variabel untuk digit tertentu dan menghitung panjang nomor voucher. Jika panjangnya bukan 5 atau 6, voucher dianggap tidak valid. Fungsi juga membandingkan hasil perkalian dua digit awal dan akhir serta memeriksa kesamaan dua digit tengah. Namun, logika pengisian digit belum diimplementasikan, sehingga validasi belum lengkap.
   
   ```go
      var N int
   ```
   kode diatas adalah kode yang digunakan untuk mendeklarasikan 1 variabel int yang digunakan sebagai inputan.
   
   ```go
      fmt.Print("Masukkan jumlah mahasiswa: ")
	  fmt.Scanln(&N)
   ```
   kode diatas adalah kode yang digunakan untuk menampilkan pesan untuk pengguna agar pengguna menginputkan suatu nilai untuk dieksekusi. 
   
   ```go
      jumlahValid := 0
	  jumlahTidakValid := 0
   ```
   Kode di atas menginisialisasi dua variabel, jumlahValid dan jumlahTidakValid, dengan nilai awal 0 untuk menghitung voucher valid dan tidak valid.
   
   ```go
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
   ```
   Kode di atas adalah loop for dalam bahasa Go yang meminta input nomor voucher mahasiswa sebanyak N kali. Setiap nomor voucher disimpan dalam variabel noTiket_2311102098, lalu diverifikasi melalui fungsi cekVoucher. Jika valid, variabel jumlahValid bertambah; jika tidak, jumlahTidakValid bertambah.
   
   ```go
      fmt.Println("Jumlah voucher valid:", jumlahValid)
	  fmt.Println("Jumlah voucher tidak valid:", jumlahTidakValid)
   ```
   Kode di atas menampilkan jumlah voucher yang valid dan tidak valid di layar. fmt.Println digunakan untuk mencetak nilai dari variabel jumlahValid dan jumlahTidakValid dengan teks penjelas masing-masing.
   

   ## SOAL 2
   Program di atas adalah program yang menghitung total biaya berdasarkan jumlah menu dan jumlah orang dalam suatu rombongan. 
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Fungsi 'hitungBiaya', Menghitung total biaya berdasarkan jumlah menu, jumlah orang, dan apakah ada sisa makanan.
      - Input Pengguna, Meminta input untuk jumlah rombongan, jenis minuman, dan informasi sisa makanan.
      - Output Biaya, Menampilkan total biaya untuk setiap rombongan setelah perhitungan.
      
   ## Code Explanation
      ```go
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
      ```
   Kode di atas adalah fungsi hitungBiaya dalam bahasa Go yang menghitung total biaya berdasarkan jumlah menu, jumlah orang, dan informasi tentang sisa makanan. Jika jumlah menu 10 atau lebih, biaya tetap Rp 150.000; jika antara 5 hingga 9, biaya dihitung dengan rumus tertentu; jika kurang dari 5, biaya dihitung per menu. Jika ada sisa makanan, total biaya akan dikurangi 20% dan dikalikan dengan jumlah orang sebelum dikembalikan sebagai hasil.
   
   ```go
      var M int
   ```
   kode diatas adalah kode yang digunakan untuk mendeklarasikan 1 variabel int yang digunakan sebagai inputan jumlah rombongan.
   
   ```go
      fmt.Print("Masukkan jumlah rombongan: ")
	  fmt.Scanln(&M)
   ```
   kode diatas adalah kode yang digunakan untuk menampilkan pesan untuk pengguna agar pengguna menginputkan suatu nilai untuk dieksekusi. 
   
   ```go
      for i := 1; i <= M; i++ {
    		var jumlahMenu, jumlahOrang int
    		var adaSisa bool
    		fmt.Print("Masukkan jumlah jenis minuman, banyak orang, dan apakah ada sisa(false/true): ")
    		fmt.Scan(&jumlahMenu, &jumlahOrang, &adaSisa)
    
    		fmt.Printf("Total biaya untuk rombongan ke-%d: Rp %d\n", i, hitungBiaya(jumlahMenu, jumlahOrang, adaSisa))
    	}
   ```
   Kode di atas adalah kode yang menggunakan loop for untuk meminta input dari pengguna. Dalam setiap iterasi, program meminta pengguna untuk memasukkan jumlah jenis minuman, jumlah orang, dan informasi tentang adanya sisa makanan. Setelah menerima input tersebut, program memanggil fungsi hitungBiaya untuk menghitung total biaya untuk rombongan ke-i dan mencetak hasilnya ke layar.


   ## SOAL 3
   Program di atas adalah program yang menghitung jumlah bilangan positif genap yang dimasukkan oleh pengguna.
   
   ## Overview
      Program ini terdiri dari satu file bernama 'main.go' dan mencakup komponen-komponen utama berikut:
      - Pernyataan 'package main', yang mendefinisikan paket untuk program yang dapat dieksekusi.
      - Pernyataan 'import', yang digunakan untuk menyertakan paket-paket yang diperlukan (dalam hal ini, 'fmt').
      - Fungsi 'main()', yang merupakan titik awal dari setiap program Go.
      - Fungsi 'jumlahBilPositif', Meminta pengguna untuk memasukkan bilangan dan menghitung jumlah bilangan positif genap hingga bilangan negatif dimasukkan.
      - Input Pengguna, Menggunakan loop untuk terus meminta input bilangan hingga pengguna memasukkan bilangan negatif.


      
   ## Code Explanation
      ```go
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
      ```
   Kode di atas adalah fungsi jumlahBilPositif yang menghitung jumlah bilangan positif genap dari input pengguna. Fungsi meminta pengguna untuk memasukkan bilangan hingga bilangan negatif dimasukkan, dan jika bilangan positif dan genap, akan ditambahkan ke variabel total_2311102098. Setelah itu, fungsi mengembalikan total jumlah bilangan genap yang telah dijumlahkan.
   
   ```go
      total_2311102098 := jumlahBilPositif()
	    fmt.Println("Jumlah bilangan angka genap:", total_2311102098)
   ```
   Kode di atas merupakan bagian dari fungsi main yang menyimpan hasil dari fungsi jumlahBilPositif ke dalam variabel total_2311102098. Setelah mendapatkan total jumlah bilangan genap yang dihitung, program mencetak hasilnya ke layar dengan pesan "Jumlah bilangan angka genap:". Dengan demikian, pengguna dapat melihat jumlah total bilangan positif genap yang telah dimasukkan sebelumnya.
