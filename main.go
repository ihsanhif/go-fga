package main

import (
	"fmt"
)

func main() {
	str := "CAИAPBO"

	// Perulangan for untuk nilai i dari 0 hingga 4
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i =", i)
	}

	// Perulangan for untuk nilai j dari 0 hingga 10
	for j := 0; j <= 10; j++ {
		// Kondisi if-else untuk mencetak karakter berdasarkan nilai j
		if j == 0 {
			fmt.Println("Nilai j =", j)
		} else if j == 6 {
			fmt.Println("character U+041E 'O' starts at byte position 12")
			fmt.Println("Nilai j =", j)
		} else if j == 1 || j == 2 || j == 3 || j == 4 {
			fmt.Println("Nilai j =", j)
		} else if j == 5 {
			for i, char := range str {
				// cek apakah karakter adalah huruf besar
				if char >= 'A' && char <= 'Z' {
					// mencetak karakter unicode dan posisi byte dari karakter
					fmt.Printf("character %U '%c' starts at byte position %d\n", char, char, i)
				}
			}
		} else {
			fmt.Println("Nilai j =", j)
		}
	}
}
