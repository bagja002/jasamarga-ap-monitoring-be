package controlles

import (
	"e-monitoring/app/models"
	"e-monitoring/pkg/database"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Pembuatan data dumy sebanyak 100 data yang sudah full di isi
// Pembuatan data dummy sebanyak 300 data yang sudah diisi
func CreateDummyKomitmen(c *fiber.Ctx) error {
	// Jumlah data dummy yang ingin dibuat (contoh: 300)
	count := 30000

	// Data untuk nama AP dan jenis pekerjaan
	namaAPData := []string{"Jasamarga Pusat", "Jasamarga Cabang Jagorawiiiii", "JMTO"}

	for i := 0; i < count; i++ {
		// Buat data komitmen dummy dengan nilai yang sesuai
		komitmen := models.Komitmen{
			IdUser:                        uint(i % 3),
			NamaAP:                        namaAPData[i%3], // Menggunakan data nama AP dalam loop
			NamaPekerjaan:                 fmt.Sprintf("Pekerjaan-%d", i+1),
			JenisPekerjaan:                "Capek",
			JenisAnggaran:                 []string{"Capex", "Japex"}[i%2],            // Bergantian antara Capex dan Japex
			KomitmenAnggaranTahunBerjalan: float64(rand.Intn(500000-100000) + 100000), // Angka acak antara 100jt dan 500jt
			NilaiKontrakKeseluruhan:       float64(rand.Intn(500000-100000) + 100000), // Angka acak antara 100jt dan 500jt
			NilaiKontrakTahun:             float64(rand.Intn(300000-100000) + 100000), // Angka acak antara 100jt dan 300jt
			NamaPenyediaBarangDanJasa:     fmt.Sprintf("Penyedia-%d", i+1),
			KualifikasiPenyedia:           []string{"UMKM", "Besar", "Kecil"}[i%3],
			StatusPencatatan:              []string{"PDN", "TKDN", "IMPOR"}[i%3], // Bergantian antara PDN, TKDN, dan Impor
			PersentasePDN:                 100.0,
			PersentaseTKDN:                100.0,
			PersentaseImpor:               100.0,
			TotalBobot:                    float64(i*2 + 30),
			RealisasiWaktuMulaiKontrak:    fmt.Sprintf("Mulai-%d", i+1),
			RealisasiWaktuBerakhirKontrak: fmt.Sprintf("Berakhir-%d", i+1),
			KeteranganLainnya:             fmt.Sprintf("Keterangan-%d", i+1),
		}

		// Jika jenis anggaran adalah Japex, biarkan 40% nilai kontrak tahun kosong
		if komitmen.JenisAnggaran == "Japex" {
			komitmen.NilaiKontrakTahun = 0.6 * komitmen.KomitmenAnggaranTahunBerjalan
		}

		// Simpan data komitmen dummy ke database
		result := database.DB.Create(&komitmen)
		if result.Error != nil {
			// Penanganan kesalahan jika operasi Create gagal
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"pesan":  "Gagal membuat data komitmen dummy",
				"pesan2": result.Error,
			})
		}
	}

	return c.JSON(fiber.Map{
		"pesan": fmt.Sprintf("Berhasil membuat %d data komitmen dummy", count),
	})
}

func CreateKomitmen(c *fiber.Ctx) error {
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	if role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	if id_admin == 0 {
		return c.JSON(fiber.Map{
			"pesan": "Users tidak terdaftar ",
		})
	}

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(fiber.Map{
			"pesan": err.Error(),
		})
	}

	IdUser, err := strconv.Atoi(data["id_user"])
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi Id_User ke tipe int",
		})
	}

	komitmen_anggaran_2023, err := strconv.Atoi(data["komitmen_anggaran_thn_berjalan"])
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi Komitmen_anggaran ke tipe int",
		})
	}
	//Penambahakn waktu

	now := time.Now()              // Get the current local time
	year, month, day := now.Date() // Extract the year, month, and day

	fmt.Printf("Year: %d, Month: %s, Day: %d\n", year, month, day)

	komitmen := models.Komitmen{
		IdUser:                        uint(IdUser),
		NamaAP:                        data["nama_ap"],
		NamaPekerjaan:                 data["nama_pekerjaan"],
		JenisPekerjaan:                data["jenis_pekerjaan"],
		JenisAnggaran:                 data["jenis_anggaran"],
		KomitmenAnggaranTahunBerjalan: float64(komitmen_anggaran_2023),
		StatusPencatatan:              data["status_pencatatan"],
		BulanBuat:                     month.String(),
		TahunBuat:                     uint(year),
		Status:                        1, //baru di buat
		Is_active:                     1,
	}

	result := database.DB.Create(&komitmen)
	if result.Error != nil {
		// Penanganan kesalahan jika operasi Create gagal
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"pesan": "Gagal membuat entri komitmen",
		})
	}

	return c.JSON(fiber.Map{
		"pesan": "komitmen telah berhasil di buat",
	})
}

func GetKomitmen(c *fiber.Ctx) error {
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	if role != 1 && role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "Role tidak sesuai, masukkan role dengan benar",
		})
	}

	if id_admin == 0 {
		return c.JSON(fiber.Map{
			"pesan": "User tidak terdaftar",
		})
	}

	idUserQueryParam := c.Query("id_user")
	idKontrak := c.Query("id_kontrak")
	is_active := c.Query("is_active")
	status := c.Query("status")

	bulanBuat := c.Query("bulanBuat")
	tahunBuat := c.Query("tahunBuat")
	statusPencatatan := c.Query("status_pencatatan")

	//pengecekan isis qurey 
			

	var komitmen []models.Komitmen

	/*
			//untuk realisasi users
			if idUserQueryParam != "" && is_active != "" && status !="" && bulanBuat != "" && tahunBuat !=""{
				database.DB.Where("id_user = ? AND status = ? AND is_active = ? AND bulan_buat = ? AND tahun_buat = ?", idUserQueryParam, status, is_active, bulanBuat, tahunBuat).Find(&komitmen)
				fmt.Println("Get Untuk Realisasi berdasarkan waktu dan tahun")
			} else if idUserQueryParam != "" && is_active != "" && status !="" && bulanBuat != ""{
				database.DB.Where("id_user = ? AND status = ? AND is_active = ? AND bulan_buat = ? ", idUserQueryParam, status, is_active, bulanBuat).Find(&komitmen)
				fmt.Println("Get Untuk Realisasi berdasarkan bulan")
			}else if idUserQueryParam != "" && is_active != "" && status !=""{
		        // If the "id_user" query parameter is provided, filter the results by id_user
				//get Komitmen Masing Masing AP
		        database.DB.Where("id_user = ? AND status = ? AND is_active = ? ", idUserQueryParam, status, is_active).Find(&komitmen)
				fmt.Println("Get Untuk Realisasi")
		    } else if idUserQueryParam != "" && is_active != "" {
				database.DB.Where("id_user = ? AND is_active = ? ", idUserQueryParam, is_active).Find(&komitmen)
				fmt.Println("Get Untuk Komitmen")
			}else if idKontrak != "" && is_active != ""{
				database.DB.Where("id_kontrak = ? AND is_active = ? ", idKontrak, is_active).Find(&komitmen)
				fmt.Println("ini pengambilan berdasarkan id kontrak nya adalah", idKontrak)
			}else {
		        // If "id_user" query parameter is not provided, retrieve all komitmen

		        database.DB.Find(&komitmen)
				fmt.Println("INI mengambilan semua")
		    }

	*/

	// Inisialisasi kueri
	query := database.DB.Where("komitmen_anggaran_tahun_berjalan >0")

	if idUserQueryParam != "" {
		query = query.Where("id_user = ?", idUserQueryParam)
	}
	if idKontrak != "" {
		query = query.Where("id_kontrak = ?", idKontrak)
	}
	if is_active != "" {
		query = query.Where("is_active = ?", is_active)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if bulanBuat != "" {
		query = query.Where("bulan_buat = ?", bulanBuat)
	}
	if tahunBuat != "" {
		query = query.Where("tahun_buat = ?", tahunBuat)
	}

	if statusPencatatan != "" {
		query = query.Where("status_pencatatan = ?", statusPencatatan)
	}


	

	// Definisikan variabel untuk menyimpan hasil
	// Ganti 'YourModel' dengan nama model data Anda

	// Eksekusi kueri dan simpan hasilnya ke dalam variabel 'results'
	query.Find(&komitmen)



	return c.JSON(fiber.Map{
		"jumlah_data": len(komitmen),
		"data":        komitmen,
	})
}

func AddRealisasi(c *fiber.Ctx) error {
	id_admin, _ := c.Locals("id_admin").(int)
	nama_ap, _ := c.Locals("name").(string)
	role, _ := c.Locals("role").(float64)

	if role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	if id_admin == 0 {
		return c.JSON(fiber.Map{
			"pesan": "Users tidak terdaftar ",
		})
	}

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(fiber.Map{
			"pesan": err.Error(),
		})
	}

	data3WithoutRp := strings.ReplaceAll(data["nilai_kontrak_keseluruhan"], "Rp. ", "")
	data3WithoutCommas := strings.ReplaceAll(data3WithoutRp, ",", "")

	NKK, err := strconv.Atoi(data3WithoutCommas)
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi nilai_kontrak_tipe int",
		})
	}
	data4WithoutRp := strings.ReplaceAll(data["nilai_kontrak_tahun"], "Rp. ", "")
	data4WithoutCommas := strings.ReplaceAll(data4WithoutRp, ",", "")
	NKT, err := strconv.Atoi(data4WithoutCommas)
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi nilai_kontrak_tahun int",
		})
	}
	namaPenyediaBarangDanJasa := data["nama_penyedia_barang_dan_jasa"]
	kualifikasiPenyedia := data["kualifikasi_penyedia"]
	statusPencatatan := data["status_pencatatan"]

	persentasePDN, err := strconv.ParseFloat(data["persentase_pdn"], 64)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi persentase_pdn ke tipe float64",
		})
	}

	persentaseTKDN, err := strconv.ParseFloat(data["persentase_tkdn"], 64)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi persentase_tkdn ke tipe float64",
		})
	}

	persentaseImpor, err := strconv.ParseFloat(data["persentase_impor"], 64)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi persentase_impor ke tipe float64",
		})
	}

	totalBobot, err := strconv.ParseFloat(data["total_bobot"], 64)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi total_bobot ke tipe float64",
		})
	}

	realisasiWaktuMulaiKontrak := data["realisasi_waktu_mulai_kontrak"]
	realisasiWaktuBerakhirKontrak := data["realisasi_waktu_berakhir_kontrak"]
	keteranganLainnya := data["keterangan_lainnya"]
	now := time.Now()            // Get the current local time
	year, month, _ := now.Date() // Extract the year, month, and day

	komitmen := models.Komitmen{
		IdUser:                        uint(id_admin),
		NamaAP:                        nama_ap,
		NamaPekerjaan:                 data["NamaPekerjaan"],
		JenisPekerjaan:                data["JenisPekerjaan"],
		JenisAnggaran:                 data["JenisAnggaran"],
		NilaiKontrakKeseluruhan:       float64(NKK),
		NilaiKontrakTahun:             float64(NKT),
		NamaPenyediaBarangDanJasa:     namaPenyediaBarangDanJasa,
		KualifikasiPenyedia:           kualifikasiPenyedia,
		StatusPencatatan:              statusPencatatan,
		PersentasePDN:                 persentasePDN,
		PersentaseTKDN:                persentaseTKDN,
		PersentaseImpor:               persentaseImpor,
		TotalBobot:                    totalBobot,
		RealisasiWaktuMulaiKontrak:    realisasiWaktuMulaiKontrak,
		RealisasiWaktuBerakhirKontrak: realisasiWaktuBerakhirKontrak,
		KeteranganLainnya:             keteranganLainnya,
		TahunBuat:                     uint(year),
		BulanBuat:                     month.String(),
		BulanRealisasi:                month.String(),
		TahunRealisasi:                uint(year),
		Status:                        2,
		Is_active:                     1,
	}

	database.DB.Create(&komitmen)

	return c.JSON(fiber.Map{
		"Pesan": "Sukses Mengedit data ",

		"data": komitmen,
	})
}

func DeleteKomitmen(c *fiber.Ctx) error {

	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	if role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	if id_admin == 0 {
		return c.JSON(fiber.Map{
			"pesan": "Users tidak terdaftar ",
		})
	}

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(fiber.Map{
			"pesan": err.Error(),
		})
	}

	//id_kontrak := c.Query("id_kontrak")

	return c.JSON(fiber.Map{})
}

//Test Untuk buat omitmen yang terbaru

func UpdateKomitmen(c *fiber.Ctx) error {
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	if role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	if id_admin == 0 {
		return c.JSON(fiber.Map{
			"pesan": "Users tidak terdaftar ",
		})
	}

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(fiber.Map{
			"pesan": err.Error(),
		})
	}

	id_kontrak := c.Query("id_kontrak")

	var komitmen models.Komitmen
	result := database.DB.Where("id_kontrak = ?", id_kontrak).
		Select("nama_ap, nama_pekerjaan, jenis_pekerjaan, jenis_anggaran, komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun, nama_penyedia_barang_dan_jasa, kualifikasi_penyedia, status_pencatatan, persentase_pdn, persentase_tkdn, persentase_impor, total_bobot, realisasi_waktu_mulai_kontrak, realisasi_waktu_berakhir_kontrak, keterangan_lainnya").
		First(&komitmen)
	if result.Error != nil {
		// Penanganan kesalahan jika terjadi kesalahan dalam eksekusi perintah.
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengambil data",
		})
	}

	nilai_kontrak := data["nilai_kontrak_keseluruhan"]
	if nilai_kontrak == "" {
		return c.JSON(fiber.Map{
			"pesan": "Masukan Nilai kontrak keseluruhan",
		})
	}

	data5WithoutRp := strings.ReplaceAll(nilai_kontrak, "Rp. ", "")
	data5WithoutCommas := strings.ReplaceAll(data5WithoutRp, ",", "")

	NKK, err := strconv.Atoi(data5WithoutCommas)
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi nilai_kontrak_tipe int",
		})
	}
	data6WithoutRp := strings.ReplaceAll(data["nilai_kontrak_tahun"], "Rp. ", "")
	data6WithoutCommas := strings.ReplaceAll(data6WithoutRp, ",", "")
	NKT, err := strconv.Atoi(data6WithoutCommas)
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi nilai_kontrak_tahun int",
		})
	}
	namaPenyediaBarangDanJasa := data["nama_penyedia_barang_dan_jasa"]
	kualifikasiPenyedia := data["kualifikasi_penyedia"]
	statusPencatatan := data["status_pencatatan"]

	persentasePDN, err := strconv.ParseFloat(data["persentase_pdn"], 64)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi persentase_pdn ke tipe float64",
		})
	}

	persentaseTKDN, err := strconv.ParseFloat(data["persentase_tkdn"], 64)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi persentase_tkdn ke tipe float64",
		})
	}

	persentaseImpor, err := strconv.ParseFloat(data["persentase_impor"], 64)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi persentase_impor ke tipe float64",
		})
	}

	totalBobot, err := strconv.ParseFloat(data["total_bobot"], 64)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi total_bobot ke tipe float64",
		})
	}

	realisasiWaktuMulaiKontrak := data["realisasi_waktu_mulai_kontrak"]
	realisasiWaktuBerakhirKontrak := data["realisasi_waktu_berakhir_kontrak"]
	keteranganLainnya := data["keterangan_lainnya"]
	now := time.Now()            // Get the current local time
	year, month, _ := now.Date() // Extract the year, month, and day

	updates := models.Komitmen{
		NilaiKontrakKeseluruhan:       float64(NKK),
		NilaiKontrakTahun:             float64(NKT),
		NamaPenyediaBarangDanJasa:     namaPenyediaBarangDanJasa,
		KualifikasiPenyedia:           kualifikasiPenyedia,
		StatusPencatatan:              statusPencatatan,
		PersentasePDN:                 persentasePDN,
		PersentaseTKDN:                persentaseTKDN,
		PersentaseImpor:               persentaseImpor,
		TotalBobot:                    totalBobot,
		RealisasiWaktuMulaiKontrak:    realisasiWaktuMulaiKontrak,
		RealisasiWaktuBerakhirKontrak: realisasiWaktuBerakhirKontrak,
		KeteranganLainnya:             keteranganLainnya,
		BulanRealisasi:                month.String(),
		TahunRealisasi:                uint(year),
		Status:                        2,
	}

	database.DB.Model(&komitmen).Where("id_kontrak = ?", id_kontrak).Updates(updates)

	return c.JSON(fiber.Map{
		"Pesan": "Sukses Mengedit data ",

		"data": komitmen,
	})
}

func DevBuatKom(c *fiber.Ctx) error {
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	if role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	if id_admin == 0 {
		return c.JSON(fiber.Map{
			"pesan": "Users tidak terdaftar ",
		})
	}

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(fiber.Map{
			"pesan": err.Error(),
		})
	}

	IdUser, err := strconv.Atoi(data["id_user"])
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi Id_User ke tipe int",
		})
	}

	data1WithoutRp := strings.ReplaceAll(data["komitmen_anggaran_thn_berjalan"], "Rp. ", "")
	data1WithoutCommas := strings.ReplaceAll(data1WithoutRp, ",", "")

	komitmen_anggaran_thn_berjalan, err := strconv.Atoi(data1WithoutCommas)
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi Komitmen_anggaran ke tipe int",
		})
	}

	data2WithoutRp := strings.ReplaceAll(data["komitmen_keseluruhan_anggaran_thn_berjalan"], "Rp. ", "")
	data2WithoutCommas := strings.ReplaceAll(data2WithoutRp, ",", "")

	komitmen_keseluruhan_anggaran_thn_berjalan, err := strconv.Atoi(data2WithoutCommas)
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi Komitmen keseluruhan anggaran ke tipe int",
		})
	}

	rencana_waktu_thn_berakhir, err := strconv.Atoi(data["rencana_tahun_berakhir"])
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi Waktu tahun berikahir ke tipe int",
		})
	}
	//Penambahakn waktu
	database.Connect()
	now := time.Now()              // Get the current local time
	year, month, day := now.Date() // Extract the year, month, and day

	fmt.Printf("Year: %d, Month: %s, Day: %d\n", year, month, day)

	komitmen := models.Komitmen{
		IdUser:                        uint(IdUser),
		NamaAP:                        data["nama_ap"],
		NamaPekerjaan:                 data["nama_pekerjaan"],
		JenisPekerjaan:                data["jenis_pekerjaan"],
		JenisAnggaran:                 data["jenis_anggaran"],
		KomitmenAnggaranTahunBerjalan: float64(komitmen_anggaran_thn_berjalan),
		KomitmenKeseluruhan:           float64(komitmen_keseluruhan_anggaran_thn_berjalan),
		StatusPadi:                    data["status_padi"],
		RencanaKualifikasiPenyedia:    data["rencana_kualifikasi_penyedia"],
		RencanaWaktuMulaiPekerjaan:    data["rencana_waktu_mulai_pekerjaan"],
		RencanaTahunBerakhir:          uint(rencana_waktu_thn_berakhir),
		StatusPencatatan:              data["status_pencatatan"],
		BulanBuat:                     month.String(),
		TahunBuat:                     uint(year),
		Status:                        1, //baru di buat
		Is_active:                     1,
	}

	result := database.DB.Create(&komitmen)
	if result.Error != nil {
		// Penanganan kesalahan jika operasi Create gagal
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"pesan": "Gagal membuat entri komitmen",
		})
	}

	return c.JSON(fiber.Map{
		"pesan": "komitmen telah berhasil di buat",
	})
}
