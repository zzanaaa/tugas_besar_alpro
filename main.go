package main

import (
	"bufio"
    "fmt"
    "sort"
    "strings"
)

type Rating struct {
    Pengguna  int
    Rating   int
}

type CoWorkingSpace struct {
    ID         int
    Nama       string
    Lokasi   string
    Fasilitas []string
    Harga      float64
    Ratings    []Rating
}

var spaces []CoWorkingSpace
var nextID = 1

func AddSpace(nama, lokasi string, fasilitas []string, harga float64) {
    space := CoWorkingSpace{
        ID:         nextID,
        Nama:       nama,
        Lokasi:   lokasi,
        FAsilitas: fasilitas,
        Harga:      harga,
    }
    nextID++
    spaces = append(spaces, space)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\n  Aplikasi Co-Working Space ")
		fmt.Println("Tambah Co-Working Space")
		fmt.Println("lihat semua")
		fmt.Println("nama pengguna (Sequential)")
		fmt.Println("rating (Binary)")
		fmt.Println("urutkan harga")
		fmt.Println("urutkan rating")
		fmt.Println("filter fasilitas")
		fmt.Println("simpan")
		fmt.Println("keluar")
		scanner.Scan()
		pilih := scanner.Text()
