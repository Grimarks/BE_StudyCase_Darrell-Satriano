package main

import "fmt"

func main() {
	input := 0
	fmt.Scanln(&input) //Kita input dulu int mau berapa kali input gedegok

	gedegok := []string{} //Simpan gedegok pakai list biar mudah

	//Loop input gedegok
	for i := 0; i < input; i++ {
		var s string
		fmt.Scanln(&s)
		gedegok = append(gedegok, s)
	}

	// loop pengecekan
	for i := 0; i < input; i++ {
		g := 0
		c := 0
		valid := true
		//ada g dan c di atas, jadi nanti di data dan dibandingin nanti

		s := gedegok[i] //untuk cek per huruf

		for j := 0; j < len(s); j++ {
			// hitung g dan c nya
			if s[j] == 'G' {
				g++
			}
			if s[j] == 'C' {
				c++
			}

			// cek DGD
			if s[j] == 'G' {
				if j-1 >= 0 && j+1 < len(s) {
					if s[j-1] == 'D' && s[j+1] == 'D' {
						valid = false
						break
					}
				}
			}
		}

		// cek jumlah g dan c
		if g != c {
			valid = false
		}

		if valid {
			fmt.Println("VALID")
		} else {
			fmt.Println("TIDAK VALID")
		}
	}
}
