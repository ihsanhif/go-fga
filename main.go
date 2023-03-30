package main

import (
	"fmt"
	"os"
)

type Teman struct {
	Absen        int
	Nama         string
	Alamat       string
	Pekerjaan    string
	AlasanGolang string
}

var temanList = []Teman{
	{1, "John Doe", "Jl. Merdeka No. 1", "Software Engineer", "Karena Golang memiliki performa yang cepat dan efisien."},
	{2, "Jane Doe", "Jl. Sudirman No. 2", "Data Scientist", "Karena Golang sangat mudah dipelajari dan memiliki sintaks yang sederhana."},
	{3, "Bob Smith", "Jl. Thamrin No. 3", "DevOps Engineer", "Karena Golang memiliki dukungan untuk concurrency dan parallelism yang kuat."},
	{4, "Alice Brown", "Jl. Gatot Subroto No. 4", "Front-end Developer", "Karena Golang dapat digunakan untuk membangun aplikasi web yang scalable dan performant."},
	{5, "Charlie Green", "Jl. Diponegoro No. 5", "Back-end Developer", "Karena Golang memiliki built-in library yang lengkap dan mudah digunakan."},
}

func main() {
	absen := os.Args[1]

	for _, teman := range temanList {
		if fmt.Sprint(teman.Absen) == absen {
			fmt.Println("Nama:", teman.Nama)
			fmt.Println("Alamat:", teman.Alamat)
			fmt.Println("Pekerjaan:", teman.Pekerjaan)
			fmt.Println("Alasan memilih kelas Golang:", teman.AlasanGolang)
			return
		}
	}

	fmt.Println("Teman dengan absen", absen, "tidak ditemukan")
}
