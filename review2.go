package main

import (
	"fmt"
	"strings"
	"time"
)

type Proyek struct {
	ID       int
	Nama     string
	Klien    string
	Deadline time.Time
	Bayaran  float64
	Status   string // "pending", "sedang dikerjakan", "selesai"
}

var proyekList []Proyek
var nextID = 1

func TambahProyek(nama, klien, deadlineStr, status string, bayaran float64) {
	deadline, _ := time.Parse("2006-01-02", deadlineStr)
	p := Proyek{
		ID:       nextID,
		Nama:     nama,
		Klien:    klien,
		Deadline: deadline,
		Bayaran:  bayaran,
		Status:   status,
	}
	nextID++
	proyekList = append(proyekList, p)
}

func UpdateStatus(id int, status string) {
	validStatuses := map[string]bool{
		"pending":           true,
		"sedang dikerjakan": true,
		"selesai":           true,
	}

	if !validStatuses[status] {
		fmt.Println("Status tidak valid.")
		return
	}

	for i := range proyekList {
		if proyekList[i].ID == id {
			proyekList[i].Status = status
			fmt.Println("Status berhasil diperbarui.")
			return
		}
	}
	fmt.Println("Proyek dengan ID tersebut tidak ditemukan.")
}

func HapusProyek(id int) {
	for i := range proyekList {
		if proyekList[i].ID == id {
			proyekList = append(proyekList[:i], proyekList[i+1:]...)
			return
		}
	}
}

func SequentialSearch(keyword string) []Proyek {
	var hasil []Proyek
	for _, p := range proyekList {
		if strings.Contains(strings.ToLower(p.Nama), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(p.Klien), strings.ToLower(keyword)) {
			hasil = append(hasil, p)
		}
	}
	return hasil
}

func SelectionSortByDeadline() {
	n := len(proyekList)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if proyekList[j].Deadline.Before(proyekList[min].Deadline) {
				min = j
			}
		}
		proyekList[i], proyekList[min] = proyekList[min], proyekList[i]
	}
}

func InsertionSortByBayaran() {
	for i := 1; i < len(proyekList); i++ {
		key := proyekList[i]
		j := i - 1
		for j >= 0 && proyekList[j].Bayaran < key.Bayaran {
			proyekList[j+1] = proyekList[j]
			j--
		}
		proyekList[j+1] = key
	}
}

func TampilkanLaporan() {
	fmt.Println("\n100% Proyek Selesai:")
	for _, p := range proyekList {
		if p.Status == "selesai" {
			CetakProyek(p)
		}
	}

	fmt.Println("\n! Proyek Berjalan:")
	for _, p := range proyekList {
		if p.Status != "selesai" {
			CetakProyek(p)
		}
	}
}

func CetakProyek(p Proyek) {
	fmt.Printf("ID: %d | %s (%s) | Deadline: %s | Rp%.2f | Status: %s\n",
		p.ID, p.Nama, p.Klien, p.Deadline.Format("2006-01-02"), p.Bayaran, p.Status)
}

func main() {
	for {
		fmt.Println("Aplikasi Manajemen dan Tracking Kegiatan Freelance")
		fmt.Println("1. Tambah Proyek")
		fmt.Println("2. Lihat Semua Proyek")
		fmt.Println("3. Ubah Status Proyek")
		fmt.Println("4. Hapus Proyek")
		fmt.Println("5. Tampilkan Laporan")
		fmt.Println("6. Keluar")
		fmt.Print("Masukkan opsi yang Anda inginkan: ")

		var pilih int
		_, err := fmt.Scanln(&pilih)
		if err != nil {
			fmt.Println("Input tidak valid. Coba lagi.")
			continue
		}

		switch pilih {
		case 1:
			var nama, klien, deadline, status string
			var bayaran float64
			fmt.Print("Nama Proyek: ")
			fmt.Scanln(&nama)
			fmt.Print("Nama Klien: ")
			fmt.Scanln(&klien)
			fmt.Print("Deadline (YYYY-MM-DD): ")
			fmt.Scanln(&deadline)
			fmt.Print("Status: ")
			fmt.Scanln(&status)
			fmt.Print("Bayaran: ")
			fmt.Scanln(&bayaran)
			TambahProyek(nama, klien, deadline, status, bayaran)
		case 2:
			for _, p := range proyekList {
				CetakProyek(p)
			}
		case 3:
			var id int
			var status string
			fmt.Print("ID Proyek: ")
			fmt.Scanln(&id)
			fmt.Print("Status Baru: ")
			fmt.Scanln(&status)
			UpdateStatus(id, status)
		case 4:
			var id int
			fmt.Print("ID Proyek yang akan dihapus: ")
			fmt.Scanln(&id)
			HapusProyek(id)
		case 5:
			TampilkanLaporan()
		case 6:
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Opsi tidak valid, silakan pilih antara 1-6.")
		}
	}
}
