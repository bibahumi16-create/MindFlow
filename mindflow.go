package main

import (
	"fmt"
)

type mood struct {
	Hari      int
	Bulan     int
	Tahun     int
	SkorEmosi int
	Perasaan  string
}

type tugas struct {
	Tanggal   string
	NamaTugas string
	Prioritas int
	Durasi    int
	Selesai   bool
}

type riwayat struct {
	Tanggal string
	Pesan   string
	Catatan string
}

var dataMood [1000]mood
var jumlahMood int

var dataTugas [1000]tugas
var jumlahTugas int

var dataRiwayat [1000]riwayat
var jumlahRiwayat int

func main() {
	var pilihan int

	for {
		fmt.Println("|============================|")
		fmt.Println("|        ++MINDFLOW++        |")
		fmt.Println("|============================|")
		fmt.Println("|                            |")
		fmt.Println("| Halo, Selamat Datang!      |")
		fmt.Println("|                            |")
		fmt.Println("|----------------------------|")
		fmt.Println("| Menu Utama                 |")
		fmt.Println("|----------------------------|")
		fmt.Println("| 1. Kelola Catatan Mood     |")
		fmt.Println("| 2. Kelola Tugas Harian     |")
		fmt.Println("| 3. Riwayat Percakapan      |")
		fmt.Println("| 4. Cari Data               |")
		fmt.Println("| 5. Urutkan Data            |")
		fmt.Println("| 6. Statistik               |")
		fmt.Println("| 0. Keluar                  |")
		fmt.Println("|============================|")
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			menuMood()
		case 2:
			menuTugas()
		case 3:
			menuRiwayat()
		case 4:
			menuPencarian()
		case 5:
			menuPengurutan()
		case 6:
			menuStatistik()
		case 0:
			fmt.Println("Terima kasih telah menggunakan MindFlow")
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func menuMood() {
	var pilihan int

	for {
		fmt.Println("=====================================")
		fmt.Println("         Kelola Catatan Mood         ")
		fmt.Println("=====================================")
		fmt.Println()
		fmt.Println("1. Tambah Mood")
		fmt.Println("2. Lihat semua Mood")
		fmt.Println("3. Ubah Mood")
		fmt.Println("4. Hapus Mood")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih Menu: ")

		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahMood()
		case 2:
			lihatMood()
		case 3:
			ubahMood()
		case 4:
			hapusMood()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func tambahMood() {
	var hari, bulan, tahun, skorEmosi int
	var perasaan string

	fmt.Print("Hari: ")
	fmt.Scan(&hari)

	fmt.Print("Bulan: ")
	fmt.Scan(&bulan)

	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	for i := 0; i < jumlahMood; i++ {
		if dataMood[i].Hari == hari &&
			dataMood[i].Bulan == bulan &&
			dataMood[i].Tahun == tahun {
			fmt.Println("Mood pada tanggal tersebut sudah ada")
			return
		}
	}

	for {
		fmt.Print("Skor Emosi (1-10): ")
		fmt.Scan(&skorEmosi)

		if skorEmosi >= 1 && skorEmosi <= 10 {
			break
		}

		fmt.Println("Skor emosi harus antara 1 sampai 10!")
	}
	fmt.Print("Perasaan: ")
	fmt.Scan(&perasaan)

	dataMood[jumlahMood] = mood{
		Hari:      hari,
		Bulan:     bulan,
		Tahun:     tahun,
		SkorEmosi: skorEmosi,
		Perasaan:  perasaan,
	}

	jumlahMood++
	fmt.Println("Data berhasil ditambahkan")
}

func lihatMood() {
	if jumlahMood == 0 {
		fmt.Println("Belum ada data yang tersedia")
		return
	}
	for i := 0; i < jumlahMood; i++ {
		fmt.Println("\nData Mood", i+1)
		fmt.Printf("Tanggal   : %02d-%02d-%d\n", dataMood[i].Hari, dataMood[i].Bulan, dataMood[i].Tahun)
		fmt.Println("Skor Emosi: ", dataMood[i].SkorEmosi)
		fmt.Println("Perasaan  : ", dataMood[i].Perasaan)
	}
}

func ubahMood() {
	var nomor int

	if jumlahMood == 0 {
		fmt.Println("Belum ada data yang tersedia")
		return
	}

	lihatMood()

	fmt.Print("\nPilih nomor data yang diubah : ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > jumlahMood {
		fmt.Println("Nomor tidak valid")
		return
	}
	nomor--

	fmt.Print("Hari Baru: ")
	fmt.Scan(&dataMood[nomor].Hari)

	fmt.Print("Bulan Baru: ")
	fmt.Scan(&dataMood[nomor].Bulan)

	fmt.Print("Tahun Baru: ")
	fmt.Scan(&dataMood[nomor].Tahun)

	for {
		fmt.Print("Skor Baru (1-10): ")
		fmt.Scan(&dataMood[nomor].SkorEmosi)

		if dataMood[nomor].SkorEmosi >= 1 &&
			dataMood[nomor].SkorEmosi <= 10 {
			break
		}

		fmt.Println("Skor harus antara 1 sampai 10!")
	}

	fmt.Print("Perasaan Baru : ")
	fmt.Scan(&dataMood[nomor].Perasaan)
	fmt.Println("Data berhasil diubah")
}

func hapusMood() {
	var nomor int

	if jumlahMood == 0 {
		fmt.Println("Belum ada data yang tersedia")
		return
	}

	lihatMood()

	fmt.Print("\nPilih nomor data yang dihapus : ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > jumlahMood {
		fmt.Println("Nomor tidak valid")
		return
	}
	nomor--

	for i := nomor; i < jumlahMood-1; i++ {
		dataMood[i] = dataMood[i+1]
	}

	jumlahMood--

	fmt.Println("Data berhasil dihapus")
}

func menuTugas() {
	var pilihan int

	for {
		fmt.Println("|=====================================|")
		fmt.Println("|         Kelola Tugas Harian         |")
		fmt.Println("|=====================================|")
		fmt.Println("|                                     |")
		fmt.Println("| 1. Tambah Tugas                     |")
		fmt.Println("| 2. Lihat semua Tugas                |")
		fmt.Println("| 3. Ubah Tugas                       |")
		fmt.Println("| 4. Hapus Tugas                      |")
		fmt.Println("| 5. Status Tugas                     |")
		fmt.Println("| 0. Kembali                          |")
		fmt.Println("|=====================================|")
		fmt.Println()
		fmt.Print("Pilih Menu: ")

		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahTugas()
		case 2:
			lihatTugas()
		case 3:
			ubahTugas()
		case 4:
			hapusTugas()
		case 5:
			selesaikanTugas()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func tambahTugas() {
	var tanggal, namaTugas string
	var prioritas, durasi int

	fmt.Print("Tanggal: ")
	fmt.Scan(&tanggal)

	fmt.Print("Nama Tugas: ")
	fmt.Scan(&namaTugas)

	for {
		fmt.Print("Prioritas (1-5): ")
		fmt.Scan(&prioritas)

		if prioritas >= 1 && prioritas <= 5 {
			break
		}

		fmt.Println("Prioritas harus antara 1 sampai 5!")

	}

	for {
		fmt.Print("Durasi (menit): ")
		fmt.Scan(&durasi)

		if durasi > 0 {
			break
		}

		fmt.Println("Durasi harus lebih dari 0!")

	}

	dataTugas[jumlahTugas] = tugas{
		Tanggal:   tanggal,
		NamaTugas: namaTugas,
		Prioritas: prioritas,
		Durasi:    durasi,
		Selesai:   false,
	}

	jumlahTugas++

	fmt.Println("Data berhasil ditambahkan")

}

func lihatTugas() {
	if jumlahTugas == 0 {
		fmt.Println("Belum ada data yang tersedia")
		return
	}

	for i := 0; i < jumlahTugas; i++ {
		status := "Belum Selesai"

		if dataTugas[i].Selesai {
			status = "Selesai"
		}

		fmt.Println("\nData Tugas", i+1)
		fmt.Println("Tanggal   : ", dataTugas[i].Tanggal)
		fmt.Println("Nama Tugas: ", dataTugas[i].NamaTugas)
		fmt.Println("Prioritas : ", dataTugas[i].Prioritas)
		fmt.Println("Durasi    : ", dataTugas[i].Durasi, "menit")
		fmt.Println("Status    : ", status)
	}

}

func ubahTugas() {
	var nomor int

	if jumlahTugas == 0 {
		fmt.Println("Belum ada data tugas")
		return
	}

	lihatTugas()

	fmt.Print("\nPilih nomor tugas yang di ubah: ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > jumlahTugas {
		fmt.Println("Nomor tidak valid")
		return
	}

	nomor--

	fmt.Print("Tanggal Baru: ")
	fmt.Scan(&dataTugas[nomor].Tanggal)
	fmt.Print("Nama Tugas Baru : ")
	fmt.Scan(&dataTugas[nomor].NamaTugas)

	for {
		fmt.Print("Prioritas Baru (1-5): ")
		fmt.Scan(&dataTugas[nomor].Prioritas)

		if dataTugas[nomor].Prioritas >= 1 &&
			dataTugas[nomor].Prioritas <= 5 {
			break
		}

		fmt.Println("Prioritas harus antara 1 sampai 5!")
	}

	for {
		fmt.Print("Durasi Baru (menit): ")
		fmt.Scan(&dataTugas[nomor].Durasi)

		if dataTugas[nomor].Durasi > 0 {
			break
		}

		fmt.Println("Durasi harus lebih dari 0!")
	}
	fmt.Println("Data berhasil diubah")
}

func hapusTugas() {
	var nomor int

	if jumlahTugas == 0 {
		fmt.Println("Belum ada data tugas")
		return
	}

	lihatTugas()

	fmt.Print("\nPilih nomor tugas yang dihapus: ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > jumlahTugas {
		fmt.Println("Nomor tidak valid")
		return
	}

	nomor--

	for i := nomor; i < jumlahTugas-1; i++ {
		dataTugas[i] = dataTugas[i+1]
	}

	jumlahTugas--

	fmt.Println("Data berhasil dihapus")
}

func selesaikanTugas() {
	var nomor int

	if jumlahTugas == 0 {
		fmt.Println("Belum ada data tugas")
		return
	}

	lihatTugas()

	fmt.Print("\nPilih nomor tugas yang telah selesai: ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > jumlahTugas {
		fmt.Println("Nomor tugas tidak valid")
		return
	}

	dataTugas[nomor-1].Selesai = true

	fmt.Println("Status tugas berhasil diubah menjadi selesai")

}

func menuRiwayat() {
	var pilihan int

	for {
		fmt.Println("|=====================================|")
		fmt.Println("|        Menu Riwayat Percakapan      |")
		fmt.Println("|=====================================|")
		fmt.Println("| 1. Tambah Percakapan                |")
		fmt.Println("| 2. Lihat Percakapan                 |")
		fmt.Println("| 3. Hapus Percakapan                 |")
		fmt.Println("| 0. Keluar                           |")
		fmt.Println("|=====================================|")
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahPercakapan()
		case 2:
			lihatPercakapan()
		case 3:
			hapusPercakapan()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func tambahPercakapan() {
	var tanggal, pesan, catatan string

	fmt.Print("Tanggal: ")
	fmt.Scan(&tanggal)

	for i := 0; i < jumlahRiwayat; i++ {
		if dataRiwayat[i].Tanggal == tanggal {
			fmt.Println("Percakapan pada tanggal tersebut sudah ada")
			return
		}
	}

	fmt.Print("Masukan Pesan: ")
	fmt.Scan(&pesan)

	fmt.Print("Masukan Catatan: ")
	fmt.Scan(&catatan)

	dataRiwayat[jumlahRiwayat] = riwayat{
		Tanggal: tanggal,
		Pesan:   pesan,
		Catatan: catatan,
	}

	jumlahRiwayat++
	fmt.Println("Data berhasil ditambahkan")
}

func lihatPercakapan() {
	if jumlahRiwayat == 0 {
		fmt.Println("Belum ada riwayat")
		return
	}

	for i := 0; i < jumlahRiwayat; i++ {
		fmt.Println("\nData Riwayat", i+1)
		fmt.Println("Tanggal: ", dataRiwayat[i].Tanggal)
		fmt.Println("Pesan  : ", dataRiwayat[i].Pesan)
		fmt.Println("Catatan: ", dataRiwayat[i].Catatan)
	}
}

func hapusPercakapan() {
	var nomor int

	if jumlahRiwayat == 0 {
		fmt.Println("Belum ada riwayat percakapan")
		return
	}

	lihatPercakapan()

	fmt.Print("\nPilih nomor percakapan yang dihapus: ")
	fmt.Scan(&nomor)

	if nomor < 1 || nomor > jumlahRiwayat {
		fmt.Println("Nomor tidak valid")
		return
	}

	nomor--

	for i := nomor; i < jumlahRiwayat-1; i++ {
		dataRiwayat[i] = dataRiwayat[i+1]
	}

	jumlahRiwayat--

	fmt.Println("Riwayat berhasil dihapus")
}

func menuPencarian() {
	var pilihan int

	for {
		fmt.Println("|============================|")
		fmt.Println("|        Menu Pencarian      |")
		fmt.Println("|============================|")
		fmt.Println("| 1. Cari Mood Tanggal       |")
		fmt.Println("| 2. Cari Tugas Tanggal      |")
		fmt.Println("| 3. Cari Mood Perasaan      |")
		fmt.Println("| 4. Cari Tugas Prioritas    |")
		fmt.Println("| 5. Cari Tugas Durasi       |")
		fmt.Println("| 0. Keluar                  |")
		fmt.Println("|============================|")
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			cariMoodSequential()
		case 2:
			cariTugasSequential()
		case 3:
			cariMoodPerasaan()
		case 4:
			cariTugasPrioritas()
		case 5:
			cariTugasDurasi()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func sequentialSearchMood(hari, bulan, tahun int) int {
	for i := 0; i < jumlahMood; i++ {
		if dataMood[i].Hari == hari &&
			dataMood[i].Bulan == bulan &&
			dataMood[i].Tahun == tahun {
			return i
		}
	}
	return -1
}

func cariMoodSequential() {
	var hari, bulan, tahun int

	if jumlahMood == 0 {
		fmt.Println("Belum ada data mood")
		return
	}

	fmt.Print("Hari: ")
	fmt.Scan(&hari)

	fmt.Print("Bulan: ")
	fmt.Scan(&bulan)

	fmt.Print("Tahun: ")
	fmt.Scan(&tahun)

	hasil := sequentialSearchMood(hari, bulan, tahun)

	if hasil == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}

	fmt.Println("\nData ditemukan")
	fmt.Printf("Tanggal   : %02d-%02d-%d\n", dataMood[hasil].Hari, dataMood[hasil].Bulan, dataMood[hasil].Tahun)
	fmt.Println("Skor Emosi: ", dataMood[hasil].SkorEmosi)
	fmt.Println("Perasaan  : ", dataMood[hasil].Perasaan)
}

func sequentialSearchTugas(nama string) {
	var ditemukan bool

	for i := 0; i < jumlahTugas; i++ {

		if dataTugas[i].NamaTugas == nama {

			status := "Belum Selesai"

			if dataTugas[i].Selesai {
				status = "Selesai"
			}

			fmt.Println("\nData ditemukan")
			fmt.Println("Nomor      :", i+1)
			fmt.Println("Nama Tugas :", dataTugas[i].NamaTugas)
			fmt.Println("Prioritas  :", dataTugas[i].Prioritas)
			fmt.Println("Durasi     :", dataTugas[i].Durasi)
			fmt.Println("Status     :", status)

			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Data tidak ditemukan")
	}
}

func cariTugasSequential() {
	var nama string

	if jumlahTugas == 0 {
		fmt.Println("Belum ada data tugas")
		return
	}

	fmt.Print("Masukkan nama tugas: ")
	fmt.Scan(&nama)

	sequentialSearchTugas(nama)
}

func cariMoodPerasaan() {
	var perasaan string

	if jumlahMood == 0 {
		fmt.Println("Belum ada data Mood")
		return
	}

	fmt.Println("Masukan perasaan yang dicari: ")
	fmt.Scan(&perasaan)

	binarySearchMood(perasaan)
}

func binarySearchMood(perasaan string) {
	var Mood [1000]mood

	for i := 0; i < jumlahMood; i++ {
		Mood[i] = dataMood[i]
	}

	for i := 0; i < jumlahMood-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahMood; j++ {
			if Mood[j].Perasaan < Mood[minIdx].Perasaan {
				minIdx = j
			}
		}
		Mood[i], Mood[minIdx] = Mood[minIdx], Mood[i]
	}

	kiri := 0
	kanan := jumlahMood - 1
	tengahDitemukan := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2

		if Mood[tengah].Perasaan == perasaan {
			tengahDitemukan = tengah
			break
		} else if Mood[tengah].Perasaan < perasaan {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if tengahDitemukan == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}

	awal := tengahDitemukan
	for awal > 0 && Mood[awal-1].Perasaan == perasaan {
		awal--
	}

	akhir := tengahDitemukan
	for akhir < jumlahMood-1 && Mood[akhir+1].Perasaan == perasaan {
		akhir++
	}

	fmt.Printf("\n%d data ditemukan dengan perasaan '%s':\n", akhir-awal+1, perasaan)
	for i := awal; i <= akhir; i++ {
		fmt.Printf("\nData %d:\n", i-awal+1)
		fmt.Printf("Tanggal   : %02d-%02d-%d\n", Mood[i].Hari, Mood[i].Bulan, Mood[i].Tahun)
		fmt.Println("Skor Emosi: ", Mood[i].SkorEmosi)
		fmt.Println("Perasaan  : ", Mood[i].Perasaan)
	}
}

func binarySearchPrioritas(prioritas int) {
	var temp [1000]tugas
	for i := 0; i < jumlahTugas; i++ {
		temp[i] = dataTugas[i]
	}

	for i := 0; i < jumlahTugas-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahTugas; j++ {
			if temp[j].Prioritas < temp[minIdx].Prioritas {
				minIdx = j
			}
		}
		temp[i], temp[minIdx] = temp[minIdx], temp[i]
	}

	kiri := 0
	kanan := jumlahTugas - 1
	tengahDitemukan := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if temp[tengah].Prioritas == prioritas {
			tengahDitemukan = tengah
			break
		} else if temp[tengah].Prioritas < prioritas {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if tengahDitemukan == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}

	awal := tengahDitemukan
	for awal > 0 && temp[awal-1].Prioritas == prioritas {
		awal--
	}

	akhir := tengahDitemukan
	for akhir < jumlahTugas-1 && temp[akhir+1].Prioritas == prioritas {
		akhir++
	}

	fmt.Printf("\n%d data ditemukan dengan prioritas %d:\n", akhir-awal+1, prioritas)
	for i := awal; i <= akhir; i++ {
		status := "Belum Selesai"
		if temp[i].Selesai {
			status = "Selesai"
		}
		fmt.Printf("\nData %d:\n", i-awal+1)
		fmt.Println("Tanggal   :", temp[i].Tanggal)
		fmt.Println("Nama Tugas:", temp[i].NamaTugas)
		fmt.Println("Prioritas :", temp[i].Prioritas)
		fmt.Println("Durasi    :", temp[i].Durasi, "menit")
		fmt.Println("Status    :", status)
	}
}

func cariTugasPrioritas() {
	var prioritas int

	if jumlahTugas == 0 {
		fmt.Println("Belum ada data tugas")
		return
	}

	fmt.Print("Masukkan prioritas yang dicari (1-5): ")
	fmt.Scan(&prioritas)

	if prioritas < 1 || prioritas > 5 {
		fmt.Println("Prioritas harus antara 1 sampai 5")
		return
	}

	binarySearchPrioritas(prioritas)
}

func binarySearchDurasi(durasi int) {
	var temp [1000]tugas
	for i := 0; i < jumlahTugas; i++ {
		temp[i] = dataTugas[i]
	}

	for i := 1; i < jumlahTugas; i++ {
		tmp := temp[i]
		j := i - 1
		for j >= 0 && temp[j].Durasi > tmp.Durasi {
			temp[j+1] = temp[j]
			j--
		}
		temp[j+1] = tmp
	}

	kiri := 0
	kanan := jumlahTugas - 1
	tengahDitemukan := -1

	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if temp[tengah].Durasi == durasi {
			tengahDitemukan = tengah
			break
		} else if temp[tengah].Durasi < durasi {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	if tengahDitemukan == -1 {
		fmt.Println("Data tidak ditemukan")
		return
	}

	awal := tengahDitemukan
	for awal > 0 && temp[awal-1].Durasi == durasi {
		awal--
	}

	akhir := tengahDitemukan
	for akhir < jumlahTugas-1 && temp[akhir+1].Durasi == durasi {
		akhir++
	}

	fmt.Printf("\n%d data ditemukan dengan durasi %d menit:\n", akhir-awal+1, durasi)
	for i := awal; i <= akhir; i++ {
		status := "Belum Selesai"
		if temp[i].Selesai {
			status = "Selesai"
		}
		fmt.Printf("\nData %d:\n", i-awal+1)
		fmt.Println("Tanggal   :", temp[i].Tanggal)
		fmt.Println("Nama Tugas:", temp[i].NamaTugas)
		fmt.Println("Prioritas :", temp[i].Prioritas)
		fmt.Println("Durasi    :", temp[i].Durasi, "menit")
		fmt.Println("Status    :", status)
	}
}

func cariTugasDurasi() {
	var durasi int

	if jumlahTugas == 0 {
		fmt.Println("Belum ada data tugas")
		return
	}

	fmt.Print("Masukkan durasi yang dicari (menit): ")
	fmt.Scan(&durasi)

	if durasi <= 0 {
		fmt.Println("Durasi harus lebih dari 0")
		return
	}

	binarySearchDurasi(durasi)
}

func menuPengurutan() {
	var pilihan int

	for {
		fmt.Println("|============================|")
		fmt.Println("|        Menu Pengurutan     |")
		fmt.Println("|============================|")
		fmt.Println("| 1. Urutkan Tugas Prioritas |")
		fmt.Println("|    (Selection Sort)        |")
		fmt.Println("| 2. Urutkan Tugas Durasi    |")
		fmt.Println("|    (Insertion sort)        |")
		fmt.Println("| 0. Kembali                 |")
		fmt.Println("|============================|")
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			urutkanPrioritas()
		case 2:
			urutkanDurasi()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func urutkanPrioritas() {
	if jumlahTugas == 0 {
		fmt.Println("Belum ada data tugas")
		return
	}

	for i := 0; i < jumlahTugas-1; i++ {
		minIdx := i
		for j := i + 1; j < jumlahTugas; j++ {
			if dataTugas[j].Prioritas < dataTugas[minIdx].Prioritas {
				minIdx = j
			}
		}
		dataTugas[i], dataTugas[minIdx] = dataTugas[minIdx], dataTugas[i]
	}

	fmt.Println("\nData tugas berhasil diurutkan berdasarkan prioritas (Selection Sort):")
	lihatTugas()

}

func urutkanDurasi() {
	if jumlahTugas == 0 {
		fmt.Println("Belum ada data tugas")
		return
	}

	for i := 1; i < jumlahTugas; i++ {
		tmp := dataTugas[i]
		j := i - 1
		for j >= 0 && dataTugas[j].Durasi > tmp.Durasi {
			dataTugas[j+1] = dataTugas[j]
			j--
		}
		dataTugas[j+1] = tmp
	}

	fmt.Println("\nData tugas berhasil diurutkan berdasarkan durasi (Insertion Sort):")
	lihatTugas()
}

func menuStatistik() {
	var pilihan int

	for {
		fmt.Println("|============================|")
		fmt.Println("|        Menu Statistik      |")
		fmt.Println("|============================|")
		fmt.Println("| 1. Tren suasana hati       |")
		fmt.Println("| 2. Statistik tugas harian  |")
		fmt.Println("| 0. Kembali                 |")
		fmt.Println("|============================|")
		fmt.Println()
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			trenSuasanaHati()
		case 2:
			statistikTugas()
		case 0:
			return
		default:
			fmt.Println("Pilihan tidak tersedia")
		}
	}
}

func trenSuasanaHati() {
	if jumlahMood == 0 {
		fmt.Println("Belum ada data mood")
		return
	}

	var bulan, tahun int
	fmt.Print("Masukkan bulan (1-12): ")
	fmt.Scan(&bulan)
	fmt.Print("Masukkan tahun: ")
	fmt.Scan(&tahun)

	ada := false
	for i := 0; i < jumlahMood; i++ {
		if dataMood[i].Bulan == bulan && dataMood[i].Tahun == tahun {
			ada = true
			break
		}
	}

	if !ada {
		fmt.Println("Tidak ada data mood pada bulan dan tahun tersebut")
		return
	}

	fmt.Println("++++++++++++ MindFlow ++++++++++++")
	fmt.Println("Tren Suasana Hati Mingguan")
	fmt.Printf("Bulan: %02d-%d\n\n", bulan, tahun)

	mingguAwal := [4]int{1, 8, 15, 22}
	mingguAkhir := [4]int{7, 14, 21, 31}

	for m := 0; m < 4; m++ {
		total := 0
		jumlah := 0

		for i := 0; i < jumlahMood; i++ {
			if dataMood[i].Bulan == bulan &&
				dataMood[i].Tahun == tahun &&
				dataMood[i].Hari >= mingguAwal[m] &&
				dataMood[i].Hari <= mingguAkhir[m] {
				total += dataMood[i].SkorEmosi
				jumlah++
			}
		}

		fmt.Printf("Minggu %d (%d-%d)\n", m+1, mingguAwal[m], mingguAkhir[m])
		if jumlah == 0 {
			fmt.Println("  Tidak ada data")
		} else {
			rataRata := float64(total) / float64(jumlah)
			fmt.Printf("  Jumlah Data  : %d\n", jumlah)
			fmt.Printf("  Rata-rata    : %.2f\n", rataRata)
		}
		fmt.Println()
	}

	fmt.Println("++++++++++++++++++++++++++++++++++")
}

func statistikTugas() {
	if jumlahTugas == 0 {
		fmt.Println("Belum ada data tugas")
		return
	}

	var tanggalUnik [1000]string
	jumlahTanggal := 0

	for i := 0; i < jumlahTugas; i++ {
		found := false
		for j := 0; j < jumlahTanggal; j++ {
			if tanggalUnik[j] == dataTugas[i].Tanggal {
				found = true
				break
			}
		}
		if !found {
			tanggalUnik[jumlahTanggal] = dataTugas[i].Tanggal
			jumlahTanggal++
		}
	}

	fmt.Println("++++++++++++ MindFlow ++++++++++++")
	fmt.Println("Statistik Tingkat Penyelesaian Tugas Harian")
	fmt.Println()

	for i := 0; i < jumlahTanggal; i++ {
		total := 0
		selesai := 0

		for j := 0; j < jumlahTugas; j++ {
			if dataTugas[j].Tanggal == tanggalUnik[i] {
				total++
				if dataTugas[j].Selesai {
					selesai++
				}
			}
		}

		belumSelesai := total - selesai
		tingkat := float64(selesai) / float64(total) * 100

		fmt.Println("Tanggal      :", tanggalUnik[i])
		fmt.Println("Total Tugas  :", total)
		fmt.Println("Selesai      :", selesai)
		fmt.Println("Belum Selesai:", belumSelesai)
		fmt.Printf("Tingkat      : %.2f%%\n", tingkat)
		fmt.Println()
	}
	fmt.Println("++++++++++++++++++++++++++++++++++")
}
