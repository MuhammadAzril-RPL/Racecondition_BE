package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Barang struct {
	Nama  string
	Harga float64
}

type Pelanggan struct {
	ID         int
	Nama       string
	Keranjang  []Barang
	NomorAntri int
	WaktuMulai time.Time
}

var (
	totalPemasukan float64
	mu             sync.Mutex
)

func main() {
	var wg sync.WaitGroup

	daftarBarang := []Barang{
		{"Buku", 50000},
		{"Pensil", 3000},
		{"Tas", 100000},
		{"Botol", 25000},
		{"Tempat makan", 100000},
	}

	antrian := make(chan Pelanggan)

	go func() {
		for i := 0; i < 100; i++ {
			pelanggan := Pelanggan{
				ID:         i + 1,
				Nama:       gofakeit.Name(),
				NomorAntri: i + 1,
				WaktuMulai: time.Now(),
			}
			jumlahBeli := rand.Intn(3) + 1
			for j := 0; j < jumlahBeli; j++ {
				pelanggan.Keranjang = append(pelanggan.Keranjang, daftarBarang[rand.Intn(len(daftarBarang))])
			}

			antrian <- pelanggan
		}
		close(antrian)
	}()

	jumlahKasir := 5
	for i := 1; i <= jumlahKasir; i++ {
		wg.Add(1)
		go func(idKasir int) {
			defer wg.Done()
			for pelanggan := range antrian {
				time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
				total := 0.0
				for _, barang := range pelanggan.Keranjang {
					total += barang.Harga
				}
				waktuTunggu := time.Since(pelanggan.WaktuMulai)
				mu.Lock()
				totalPemasukan += total
				fmt.Printf("Kasir %d melayani %s (Antrian: %d) - Total: Rp%.2f - Waktu tunggu: %v\n",
					idKasir, pelanggan.Nama, pelanggan.NomorAntri, total, waktuTunggu.Round(time.Second))
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
	fmt.Println()
	fmt.Print("Semua Pelanggan telah di layani")
	fmt.Printf("\nTotal pemasukan toko: Rp%.2f\n", totalPemasukan)
}
