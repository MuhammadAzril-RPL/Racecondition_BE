# Simulasi Antrian Pelanggan di Kasir

## Deskripsi
Program ini mensimulasikan proses antrian pelanggan di sebuah toko dengan beberapa kasir. Setiap pelanggan akan memiliki sejumlah barang yang dibeli, dan kasir akan melayani mereka secara paralel menggunakan goroutine. Program ini juga menghitung total pemasukan toko setelah semua pelanggan telah dilayani.

## Alur Program
1. **Inisialisasi Data**
   - Program mendefinisikan daftar barang yang tersedia di toko beserta harganya.
   - Menggunakan channel untuk merepresentasikan antrian pelanggan.
   
2. **Simulasi Kedatangan Pelanggan**
   - Sebanyak 100 pelanggan dibuat secara otomatis dengan ID dan nama acak menggunakan pustaka `gofakeit`.
   - Setiap pelanggan mendapatkan nomor antrian dan waktu mulai transaksi.
   - Pelanggan memilih antara 1 hingga 3 barang secara acak dari daftar barang yang tersedia.
   - Pelanggan kemudian dimasukkan ke dalam antrian (channel) untuk diproses oleh kasir.

3. **Pemrosesan oleh Kasir**
   - Program menjalankan 5 kasir secara paralel menggunakan goroutine.
   - Setiap kasir mengambil pelanggan dari antrian dan mensimulasikan proses pembayaran dengan delay acak antara 1 hingga 5 detik.
   - Kasir menghitung total harga belanjaan pelanggan dan mencatat waktu tunggu mereka.
   - Total pemasukan toko diperbarui secara aman menggunakan `sync.Mutex` untuk menghindari kondisi balapan (race condition).
   - Informasi transaksi ditampilkan ke layar dalam format:
     ```
     Kasir X melayani [Nama Pelanggan] (Antrian: N) - Total: RpXXXXX.XX - Waktu tunggu: Xs
     ```

4. **Penyelesaian dan Rekapitulasi**
   - Setelah semua pelanggan dilayani, program mencetak total pemasukan toko.
   - Pesan penutup menandakan bahwa semua pelanggan telah selesai dilayani.
