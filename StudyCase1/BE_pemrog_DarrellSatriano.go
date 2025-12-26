package main

import "fmt"

func main() {
	input := 0
	fmt.Scanln(&input)

	gedegok := []string{}

	for i := 0; i < input; i++ {
		var s string
		fmt.Scanln(&s)
		gedegok = append(gedegok, s)
	}

	for i := 0; i < input; i++ {
		g := 0
		c := 0
		valid := true

		s := gedegok[i]

		for j := 0; j < len(s); j++ {
			// hitung G dan C nya
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

		// cek jumlah G dan C
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
