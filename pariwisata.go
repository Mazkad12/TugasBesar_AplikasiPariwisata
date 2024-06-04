package main

import "fmt"

const NMAX int = 20

type fitur [6]string
type tempat_wisata struct {
	nama             string
	jarak, biaya, id int
	fasilitas        [6]string
}

type tabWisata [NMAX]tempat_wisata

func main() {
	var data tabWisata
	var nData int
	var fasilitas = fitur{"Toilet", "Tempat Makan", "Smoking Area", "Air", "Hewan", "Wahana"}
	data[0].nama = "jerapah"
	data[0].id = 2
	data[0].jarak = 1230
	data[0].biaya = 60000
	data[0].fasilitas[0] = "Toilet"
	data[0].fasilitas[1] = "Tempat Makan"
	data[0].fasilitas[2] = "Smoking Area"
	data[0].fasilitas[3] = "Air"
	data[0].fasilitas[4] = "Hewan"
	data[0].fasilitas[5] = "Wahana"
	data[1].nama = "harimau"
	data[1].id = 4
	data[1].jarak = 2000
	data[1].biaya = 23500
	data[1].fasilitas[0] = "Toilet"
	data[1].fasilitas[1] = "Tempat Makan"
	data[2].nama = "buaya_darat"
	data[2].id = 3
	data[2].jarak = 3550
	data[2].biaya = 10000
	data[2].fasilitas[2] = "Smoking Area"
	data[2].fasilitas[3] = "Air"
	data[2].fasilitas[4] = "Hewan"
	data[2].fasilitas[5] = "Wahana"
	data[3].nama = "histeria"
	data[3].id = 1
	data[3].jarak = 1256
	data[3].biaya = 25000
	data[3].fasilitas[0] = "Toilet"
	data[3].fasilitas[1] = "Tempat Makan"
	data[3].fasilitas[5] = "Wahana"
	nData = 4
	login(data, nData, fasilitas)
}

func login(data tabWisata, nData int, fasilitas fitur) {
	var opsi int
	for opsi != 3 {
		fmt.Println("===================================================")
		fmt.Println("Selamat Datang Di Aplikasi Pariwisata")
		fmt.Println("1. Login Admin")
		fmt.Println("2. Login User")
		fmt.Println("3. Exit")
		fmt.Println("===================================================")
		fmt.Scan(&opsi)
		if opsi == 1 {
			home_admin(data, nData, fasilitas)
		} else if opsi == 2 {
			home_user(data, nData, fasilitas)
		} else if opsi < 1 || opsi > 3 {
			fmt.Println("Opsi Invalid")
		}
	}
}
func home_admin(data tabWisata, nData int, fasilitas fitur) {
	var opsi int
	for opsi != 4 {
		fmt.Println("===================================================")
		fmt.Println("Menu Admin")
		fmt.Println("1. Tambah Wisata")
		fmt.Println("2. Edit Wisata")
		fmt.Println("3. Hapus Wisata")
		fmt.Println("4. Exit")
		fmt.Println("===================================================")
		fmt.Scan(&opsi)
		if opsi == 1 {
			main_tambah_wisata(&data, &nData, fasilitas)
		} else if opsi == 2 {
			main_edit_wisata(data, nData, fasilitas)
		} else if opsi == 3 {
			main_hapus_wisata(data, nData, fasilitas)
		} else if opsi < 1 || opsi > 4 {
			fmt.Println("Opsi Invalid")
		}
	}
}
func main_tambah_wisata(data *tabWisata, nData *int, fasilitas fitur) {
	var opsi int
	fmt.Println("===================================================")
	fmt.Println("Masukkan Nama Wisata")
	fmt.Scan(&data[*nData].nama)
	fmt.Println("Masukkan ID Wisata")
	fmt.Scan(&data[*nData].id)
	fmt.Println("Masukkan Jarak Wisata")
	fmt.Scan(&data[*nData].jarak)
	fmt.Println("Masukkan Biaya Wisata")
	fmt.Scan(&data[*nData].biaya)
	fmt.Println("Masukkan Fasilitas Wisata")
	list_fasilitas(fasilitas)
	for opsi != 7 {
		fmt.Scan(&opsi)
		if opsi == 1 {
			data[*nData].fasilitas[0] = fasilitas[0]
		} else if opsi == 2 {
			data[*nData].fasilitas[1] = fasilitas[1]
		} else if opsi == 3 {
			data[*nData].fasilitas[2] = fasilitas[2]
		} else if opsi == 4 {
			data[*nData].fasilitas[3] = fasilitas[3]
		} else if opsi == 5 {
			data[*nData].fasilitas[4] = fasilitas[4]
		} else if opsi == 6 {
			data[*nData].fasilitas[5] = fasilitas[5]
		}
	}

	*nData++
	fmt.Println("===================================================")
}
func list_fasilitas(fasilitas fitur) {
	for i := 0; i < 6; i++ {
		fmt.Printf("%d. %s \n", i+1, fasilitas[i])
	}
	fmt.Println("7. Selesai")
}
func main_edit_wisata(data tabWisata, nData int, fasilitas fitur) {
	var opsi, idx, y int
	var x string
	fmt.Println("===================================================")
	fmt.Println("Menu Edit Wisata")
	fmt.Println("1. Cari Wisata Berdasarkan Nama")
	fmt.Println("2. Cari Wisata Berdasarkan ID")
	fmt.Scan(&opsi)
	for opsi != 1 && opsi != 2 {
		fmt.Println("Input Invalid")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		fmt.Println("Masukkan Nama Wisata")
		fmt.Scan(&x)
		idx = cari_nama_wisata(data, nData, x)
		if idx != -1 {
			edit_wisata(&data, fasilitas, idx)
		}
	} else if opsi == 2 {
		fmt.Println("Masukkan ID Wisata")
		fmt.Scan(&y)
		fmt.Println(y)
		idx = cari_id_wisata(data, nData, y)
		if idx != -1 {
			edit_wisata(&data, fasilitas, idx)
		}
	}
	fmt.Println("===================================================")
}
func cari_nama_wisata(data tabWisata, nData int, x string) int {
	var idx int = -1
	for i := 0; i < nData; i++ {
		if data[i].nama == x {
			idx = i
			i += nData
		}
	}
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	}
	return idx

}
func cari_id_wisata(data tabWisata, nData int, y int) int {
	var idx int = -1
	for i := 0; i < nData; i++ {
		if data[i].id == y {
			idx = i
			i += nData
		}
	}
	if idx == -1 {
		fmt.Println("Data tidak ditemukan")
	}
	return idx

}
func edit_wisata(data *tabWisata, fasilitas fitur, idx int) {
	var opsi int
	fmt.Println("===================================================")
	fmt.Println("Masukkan Nama Wisata yang Baru")
	fmt.Scan(&data[idx].nama)
	fmt.Println("Masukkan ID Wisata yang Baru")
	fmt.Scan(&data[idx].id)
	fmt.Println("Masukkan Jarak Wisata yang Baru")
	fmt.Scan(&data[idx].jarak)
	fmt.Println("Masukkan Biaya Wisata yang Baru")
	fmt.Scan(&data[idx].biaya)
	fmt.Println("Masukkan Fasilitas Wisata yang Baru")
	list_fasilitas(fasilitas)
	for opsi != 7 {
		fmt.Scan(&opsi)
		if opsi == 1 {
			data[idx].fasilitas[0] = fasilitas[0]
		} else if opsi == 2 {
			data[idx].fasilitas[1] = fasilitas[1]
		} else if opsi == 3 {
			data[idx].fasilitas[2] = fasilitas[2]
		} else if opsi == 4 {
			data[idx].fasilitas[3] = fasilitas[3]
		} else if opsi == 5 {
			data[idx].fasilitas[4] = fasilitas[4]
		} else if opsi == 6 {
			data[idx].fasilitas[5] = fasilitas[5]
		}
	}

	fmt.Println("===================================================")
}
func main_hapus_wisata(data tabWisata, nData int, fasilitas fitur) {
	var opsi, idx, y int
	var x string
	fmt.Println("===================================================")
	fmt.Println("Menu Hapus Wisata")
	fmt.Println("1. Cari Wisata Berdasarkan Nama")
	fmt.Println("2. Cari Wisata Berdasarkan ID")
	fmt.Println("===================================================")
	fmt.Scan(&opsi)
	for opsi != 1 && opsi != 2 {
		fmt.Println("Input Invalid")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		fmt.Println("Masukkan Nama Wisata")
		fmt.Scan(&x)
		idx = cari_nama_wisata(data, nData, x)
		if idx != -1 {
			hapus_wisata(&data, &nData, idx)
			fmt.Println("Data telah dihapus")
		}
	} else if opsi == 2 {
		fmt.Println("Masukkan ID Wisata")
		fmt.Scan(&y)
		idx = cari_id_wisata(data, nData, y)
		if idx != -1 {
			hapus_wisata(&data, &nData, idx)
			fmt.Println("Data telah dihapus")
		}
	}
}
func hapus_wisata(data *tabWisata, nData *int, idx int) {
	for i := idx; i < *nData-1; i++ {
		data[i] = data[i+1]
	}
	*nData--
}
func home_user(data tabWisata, nData int, fasilitas fitur) {
	var opsi int
	for opsi != 7 {
		fmt.Println("===================================================")
		fmt.Println("Menu User")
		fmt.Println("1. List Wisata Berdasarkan Jarak Terdekat")
		fmt.Println("2. List Wisata Berdasarkan Jarak Terjauh")
		fmt.Println("3. List Wisata Berdasarkan Biaya Termurah")
		fmt.Println("4. List Wisata Berdasarkan Biaya Termahal")
		fmt.Println("5. List Wisata Berdasarkan Fasilitas")
		fmt.Println("6. Cari Wisata")
		fmt.Println("7. Exit")
		fmt.Println("===================================================")
		fmt.Scan(&opsi)

		if opsi == 1 {
			ascend_jarak(data, fasilitas, nData)
		} else if opsi == 2 {
			descend_jarak(data, fasilitas, nData)
		} else if opsi == 3 {
			ascend_biaya(data, fasilitas, nData)
		} else if opsi == 4 {
			descend_biaya(data, fasilitas, nData)
		} else if opsi == 5 {
			main_cari_fasilitas(data, nData, fasilitas)
		} else if opsi == 6 {
			main_cari_wisata(data, nData, fasilitas)
		} else if opsi < 1 || opsi > 7 {
			fmt.Println("Opsi Invalid")
			fmt.Scan(&opsi)
		}
	}
}
func ascend_jarak(data tabWisata, fasilitas fitur, nData int) {
	var i, pass, idx_min int
	for pass = 0; pass <= nData-2; pass++ {
		idx_min = pass
		for i = pass + 1; i <= nData-1; i++ {
			if data[i].jarak < data[idx_min].jarak {
				idx_min = i
			}
		}
		data[pass], data[idx_min] = data[idx_min], data[pass]
	}
	for i := 0; i < nData; i++ {
		display_wisata(data, fasilitas, i)
	}
}
func descend_jarak(data tabWisata, fasilitas fitur, nData int) {
	var i, pass, idx_max int
	for pass = 0; pass <= nData-2; pass++ {
		idx_max = pass
		for i = pass + 1; i <= nData-1; i++ {
			if data[i].jarak > data[idx_max].jarak {
				idx_max = i
			}
		}
		data[pass], data[idx_max] = data[idx_max], data[pass]
	}
	for i := 0; i < nData; i++ {
		display_wisata(data, fasilitas, i)
	}
}
func ascend_biaya(data tabWisata, fasilitas fitur, nData int) {
	for i := 1; i < nData; i++ {
		j := i
		for j > 0 {
			if data[j-1].biaya > data[j].biaya {
				data[j-1], data[j] = data[j], data[j-1]
			}
			j = j - 1
		}
	}
	for i := 0; i < nData; i++ {
		display_wisata(data, fasilitas, i)
	}
}
func descend_biaya(data tabWisata, fasilitas fitur, nData int) {
	for i := 1; i < nData; i++ {
		j := i
		for j > 0 {
			if data[j-1].biaya < data[j].biaya {
				data[j-1], data[j] = data[j], data[j-1]
			}
			j = j - 1
		}
	}
	for i := 0; i < nData; i++ {
		display_wisata(data, fasilitas, i)
	}
}
func display_wisata(data tabWisata, fasilitas fitur, idx int) {
	fmt.Println("===================================================")
	fmt.Println("Nama Wisata   : ", data[idx].nama)
	fmt.Println("ID Wisata     : ", data[idx].id)
	fmt.Println("Jarak Wisata  : ", data[idx].jarak)
	fmt.Println("Biaya Wisata  : ", data[idx].biaya)
	fmt.Print("Fasilitas Wisata: ")
	for i := 0; i < 6; i++ {
		if data[idx].fasilitas[i] == fasilitas[i] {
			fmt.Print(data[idx].fasilitas[i], " ")
		}

	}
	fmt.Println()
	fmt.Println("===================================================")
}
func main_cari_fasilitas(data tabWisata, nData int, fasilitas fitur) {
	var opsi int
	fmt.Println("===================================================")
	fmt.Println("Masukkan Fasilitas Wahana Yang Ingin Dicari: ")
	for i := 0; i < 6; i++ {
		fmt.Printf("%d. %s \n", i+1, fasilitas[i])
	}
	fmt.Println("===================================================")
	fmt.Scan(&opsi)
	for opsi < 1 || opsi > 6 {
		fmt.Println("Opsi Invalid. Silakan masukkan angka antara 1 dan 6.")
		fmt.Scan(&opsi)
	}

	cari_fasilitas(data, nData, fasilitas, opsi-1)
	return
}

func cari_fasilitas(data tabWisata, nData int, fasilitas fitur, x int) {
	var ada int
	for i := 0; i < nData; i++ {
		if data[i].fasilitas[x] == fasilitas[x] {
			display_wisata(data, fasilitas, i)
			ada++
		}
	}
	if ada == 0 {
		fmt.Println("Wisata Dengan Fasilitas Tersebut Tidak Ditemukan")
	}
}
func main_cari_wisata(data tabWisata, nData int, fasilitas fitur) {
	var opsi, idx, y int
	var x string
	fmt.Println("===================================================")
	fmt.Println("Menu Cari Wisata")
	fmt.Println("1. Cari Wisata Berdasarkan Nama")
	fmt.Println("2. Cari Wisata Berdasarkan ID")
	fmt.Println("===================================================")
	fmt.Scan(&opsi)
	for opsi != 1 && opsi != 2 {
		fmt.Println("Input Invalid")
		fmt.Scan(&opsi)
	}
	if opsi == 1 {
		fmt.Println("Masukkan Nama Wisata")
		fmt.Scan(&x)
		idx = cari_nama_wisata(data, nData, x)
		if idx != -1 {
			display_wisata(data, fasilitas, idx)
		}
	} else if opsi == 2 {
		fmt.Println("Masukkan ID Wisata")
		fmt.Scan(&y)
		idx = cari_id_wisata(data, nData, y)
		if idx != -1 {
			display_wisata(data, fasilitas, idx)
		}
	}
}
func cari_id_wisata_binary(data tabWisata, nData, y int) int {
	ascend_id_insertion(&data, nData)
	var left, mid, right int
	left = 0
	right = nData - 1
	idx := -1
	for left <= right && data[mid].id != y {
		mid = (left + right) / 2
		if y < data[mid].id {
			right = mid - 1
		} else if y > data[mid].id {
			left = mid + 1
		} else if mid == data[mid].id {
			idx = mid
		}
	}
	return idx
}

func ascend_id_insertion(data *tabWisata, nData int) {
	for i := 1; i < nData; i++ {
		j := i
		for j > 0 {
			if data[j-1].biaya > data[j].biaya {
				data[j-1], data[j] = data[j], data[j-1]
			}
			j = j - 1
		}
	}
}
