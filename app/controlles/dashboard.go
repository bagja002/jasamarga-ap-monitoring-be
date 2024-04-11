package controlles

import (
	"e-monitoring/app/models"
	"e-monitoring/pkg/database"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetDataDHAdmin(c *fiber.Ctx) error {
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	if role != 1 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	if id_admin == 0 {
		return c.JSON(fiber.Map{
			"pesan": "Admin tidak terdaftar ",
		})
	}

	//Ambil data komitmen 2023, komitmen_keseluruhan
	type Komitmen struct {
		Nama_pekerjaan                   string
		Komitmen_anggaran_tahun_berjalan float64
		Nilai_kontrak_tahun              float64
	}
	id_users := c.Query("id_users")
	//bulan_buat := c.Query("bulanBuat")
	bulan_realisasi := c.Query("bulanRealisasi")
	//tahun_buat := c.Query("tahunBuat")
	tahun_realisasi := c.Query("tahunRealisasi")

	if id_users != "" {

		//Total Komitmen
		var komitmens []Komitmen
		if err := database.DB.Select("nama_pekerjaan, komitmen_anggaran_tahun_berjalan, nilai_kontrak_tahun").
			Where("id_user = ? ", id_users).
			Find(&komitmens).Error; err != nil {
			fmt.Println("Terjadi kesalahan:", err)
		}


		totalKomitmenAnggaran := 0.0
		for _, item := range komitmens {
			totalKomitmenAnggaran += item.Komitmen_anggaran_tahun_berjalan
		}

		totalRealisasi := 0.0
		for _, item := range komitmens {
			totalRealisasi += item.Nilai_kontrak_tahun
		}

		Persentase_realisasi := 0.0
		persentase_belum_realisasi := 0.0
		if totalKomitmenAnggaran != 0 {
			Persentase_realisasi = (totalRealisasi / totalKomitmenAnggaran) * 100
			persentase_belum_realisasi = 100 - Persentase_realisasi
		}

		// Second query
		var komitmen []models.Komitmen
		var komitmenss []models.Komitmen
		var komitmensss []models.Komitmen
		var PDN_TKDN []models.Komitmen

		if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
			Where("id_user = ? AND status_pencatatan = ? ", id_users, "TKDN").
			Find(&komitmen).Error; err != nil {
			fmt.Println("Terjadi kesalahantt:", err)
		}

		if bulan_realisasi != "" && tahun_realisasi != "" {

			database.DB.Select("nama_pekerjaan, komitmen_anggaran_tahun_berjalan,  nilai_kontrak_tahun").
				Where("id_user = ? AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, bulan_realisasi, tahun_realisasi).
				Find(&komitmens)

			database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan,  nilai_kontrak_tahun").
				Where("id_user = ? AND status_pencatatan = ? AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, "TKDN", bulan_realisasi, tahun_realisasi).
				Find(&komitmen)

	

			database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan,  nilai_kontrak_tahun").
				Where("id_user = ? AND status_pencatatan = ? AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, "PDN", bulan_realisasi, tahun_realisasi).
				Find(&komitmenss)

			database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
				Where("id_user = ? AND status_pencatatan = ? AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, "IMPOR", bulan_realisasi, tahun_realisasi).
				Find(&komitmensss)

			database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, komitmen_keseluruhan").
				Where("id_user = ? AND (status_pencatatan = ? OR status_pencatatan = ?) AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, "TKDN", "PDN", bulan_realisasi, tahun_realisasi).
				Find(&PDN_TKDN)
		}

		totalTKDN := 0.0
		for _, item := range komitmen {
			totalTKDN += item.NilaiKontrakTahun
		}

		// Third query
		if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
			Where("id_user = ? AND status_pencatatan = ?", id_users, "PDN").
			Find(&komitmenss).Error; err != nil {
			fmt.Println("Terjadi kesalahan:", err)
		}

		totalPDN := 0.0
		for _, item := range komitmenss {
			totalPDN += item.NilaiKontrakTahun
		}

		// Fourth query
		if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
			Where("id_user = ? AND status_pencatatan = ?", id_users, "IMPOR").
			Find(&komitmensss).Error; err != nil {
			fmt.Println("Terjadi kesalahan:", err)
		}

		totalImport := 0.0
		totalKomitmenImpor := 0.0
		for _, item := range komitmensss {
			totalImport += item.NilaiKontrakTahun
			totalKomitmenImpor += item.KomitmenAnggaranTahunBerjalan
		}

		Persentase_realisasi_Impor := 0.0
		persentase_belum_realisasi_impor := 0.0
		if totalKomitmenImpor != 0 {
			Persentase_realisasi_Impor = (totalImport / totalKomitmenImpor) * 100
			persentase_belum_realisasi_impor = 100 - Persentase_realisasi_Impor
		}

		// Fifth query

		if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
			Where("id_user = ? AND (status_pencatatan = ? OR status_pencatatan = ?)", id_users, "TKDN", "PDN").
			Find(&PDN_TKDN).Error; err != nil {
			fmt.Println("Terjadi kesalahan:", err)
		}

		totaTkdnPdn := 0.0
		totalKomitmenTkdnPdn := 0.0
		for _, item := range PDN_TKDN {
			totaTkdnPdn += item.NilaiKontrakTahun
			totalKomitmenTkdnPdn += item.KomitmenAnggaranTahunBerjalan
		}

		Persentase_realisasi_PDN_TKDN := 0.0
		Persentase_belum_realisasi_PDN_TKDN := 0.0
		if totalKomitmenTkdnPdn != 0 {
			Persentase_realisasi_PDN_TKDN = (totaTkdnPdn / totalKomitmenTkdnPdn) * 100
			Persentase_belum_realisasi_PDN_TKDN = 100 - Persentase_realisasi_PDN_TKDN
		}

		type Dashboard struct {
			//dashboard atas
			Komitmen        float64
			KomitmenTKDN    float64
			KomitmenImpor   float64
			Realisasi       float64
			Belum_realisasi float64
			TKDN            float64
			PDN             float64
			Import          float64
			//grafik
			Persentase_Realisasi               float64
			Persentase_Belum_Realisasi         float64
			Persentase_Realisasi_belanja_impor float64
			Persentase_Belum_belanja_impor     float64

			Persentase_Realisasi_PDN_TKDN float64
			Persentase_Belum_PDN_TKDN     float64
			Realisasi_TKDNPDN             float64
		}
		data := Dashboard{
			Komitmen:                   totalKomitmenAnggaran,
			Realisasi:                  totalRealisasi,
			TKDN:                       totalTKDN,
			PDN:                        totalPDN,
			Import:                     totalImport,
			Belum_realisasi:            totalKomitmenAnggaran - totalRealisasi,
			Persentase_Realisasi:       Persentase_realisasi,
			Persentase_Belum_Realisasi: persentase_belum_realisasi,
			//data masih ngacoo
			Persentase_Realisasi_PDN_TKDN:      Persentase_realisasi_PDN_TKDN,
			Persentase_Belum_PDN_TKDN:          Persentase_belum_realisasi_PDN_TKDN,
			Persentase_Realisasi_belanja_impor: Persentase_realisasi_Impor,
			Persentase_Belum_belanja_impor:     persentase_belum_realisasi_impor,
			KomitmenTKDN:                       totalKomitmenTkdnPdn,
			KomitmenImpor:                      totalKomitmenImpor,
			Realisasi_TKDNPDN:                  totaTkdnPdn,
		}

		return c.JSON(fiber.Map{
			"data": data,
		})

	}
	//Untuk get data semuanya
	var komitmens []Komitmen
	if err := database.DB.Select("nama_pekerjaan, komitmen_anggaran_tahun_berjalan, nilai_kontrak_tahun").Find(&komitmens).Error; err != nil {
		// Terjadi kesalahan saat mengeksekusi query
		// Anda dapat melakukan penanganan kesalahan di sini
		// Contoh: log kesalahan, kembalikan kesalahan, dll
		fmt.Println("terjadi kesalahan yang signifikat", err)

	}

	//menambah total Komitmen Anggaran Secara total
	totalKomitmenAnggaran := 0.0
	for _, item := range komitmens {
		totalKomitmenAnggaran += item.Komitmen_anggaran_tahun_berjalan

	}
	//menambah Nilai Kontrak tahun
	totalRealisasi := 0.0
	for _, item := range komitmens {
		totalRealisasi += item.Nilai_kontrak_tahun
	
	}

	Persentase_realisasi := 0.0
	persentase_belum_realisasi := 0.0

	if totalKomitmenAnggaran != 0 {
		Persentase_realisasi = (totalRealisasi / totalKomitmenAnggaran) * 100.
		persentase_belum_realisasi = 100 - Persentase_realisasi
	}

	//TKDN, PDN , IMPORT

	var komitmen []models.Komitmen

	if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
		Where("status_pencatatan = ?", "TKDN").
		Find(&komitmen).Error; err != nil {
		// Terjadi kesalahan saat mengeksekusi query
		// Anda dapat melakukan penanganan kesalahan di sini
		// Contoh: log kesalahan, kembalikan kesalahan, dll.
		fmt.Println("Terjadi kesalahan:", err)

	}

	//penambahan Total TKDN
	totalTKDN := 0.0

	for _, item := range komitmen {
		totalTKDN += item.NilaiKontrakTahun

	}
	//PDN

	var komitmenss []models.Komitmen
	if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
		Where("status_pencatatan = ?", "PDN").
		Find(&komitmenss).Error; err != nil {
		// Terjadi kesalahan saat mengeksekusi query
		// Anda dapat melakukan penanganan kesalahan di sini
		// Contoh: log kesalahan, kembalikan kesalahan, dll.
		fmt.Println("Terjadi kesalahan:", err)


	}

	//penambahan Total TKDN
	totalPDN := 0.0

	for _, item := range komitmenss {
		totalPDN += item.NilaiKontrakTahun

	}
	//untuk Impor
	var komitmensss []models.Komitmen
	if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun ").
		Where("status_pencatatan = ?", "IMPOR").
		Find(&komitmensss).Error; err != nil {
		// Terjadi kesalahan saat mengeksekusi query
		// Anda dapat melakukan penanganan kesalahan di sini
		// Contoh: log kesalahan, kembalikan kesalahan, dll.
		fmt.Println("Terjadi kesalahan:", err)

	}

	// penambahan Total TKDN
	totalImport := 0.0
	totalKomitmenImpor := 0.0

	for _, item := range komitmensss {
		totalImport += item.NilaiKontrakTahun
		totalKomitmenImpor += item.KomitmenAnggaranTahunBerjalan

	}
	// persentasenya
	Persentase_realisasi_Impor := 0.0
	persentase_belum_realisasi_impor := 0.0
	if totalKomitmenImpor != 0 {
		Persentase_realisasi_Impor = (totalImport / totalKomitmenImpor) * 100
		persentase_belum_realisasi_impor = 100 - Persentase_realisasi_Impor

	}

	var PDN_TKDN []models.Komitmen
	if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
		Where("status_pencatatan = ? OR status_pencatatan = ?", "TKDN", "PDN").
		Find(&PDN_TKDN).Error; err != nil {
		// Terjadi kesalahan saat mengeksekusi query
		// Anda dapat melakukan penanganan kesalahan di sini
		// Contoh: log kesalahan, kembalikan kesalahan, dll.
		fmt.Println("Terjadi kesalahan:", err)

	}

	//perhitungan persentase
	totaTkdnPdn := 0.0
	totalKomitmenTkdnPdn := 0.0
	//realisasi
	for _, item := range PDN_TKDN {
		totaTkdnPdn += item.NilaiKontrakTahun
		
	}
	//komitmen
	for _, item := range PDN_TKDN {
		totalKomitmenTkdnPdn += item.KomitmenAnggaranTahunBerjalan
	}

	//persentase terealisasi TKDM
	Persentase_realisasi_PDN_TKDN := 0.0
	Persentase_belum_realisasi_PDN_TKDN := 0.0
	if totalKomitmenTkdnPdn != 0 {
		Persentase_realisasi_PDN_TKDN = (totaTkdnPdn / totalKomitmenTkdnPdn) * 100
		Persentase_belum_realisasi_PDN_TKDN = 100 - Persentase_realisasi_PDN_TKDN
		
	}

	type Dashboard struct {
		//dashboard atas
		Komitmen        float64
		KomitmenTKDN    float64
		KomitmenImpor   float64
		Realisasi       float64
		Belum_realisasi float64
		TKDN            float64
		PDN             float64
		Import          float64
		//grafik
		Persentase_Realisasi               float64
		Persentase_Belum_Realisasi         float64
		Persentase_Realisasi_belanja_impor float64
		Persentase_Belum_belanja_impor     float64

		Persentase_Realisasi_PDN_TKDN float64
		Persentase_Belum_PDN_TKDN     float64
		Realisasi_TKDNPDN             float64
	}
	data := Dashboard{
		Komitmen:                   totalKomitmenAnggaran,
		Realisasi:                  totalRealisasi,
		TKDN:                       totalTKDN,
		PDN:                        totalPDN,
		Import:                     totalImport,
		Belum_realisasi:            totalKomitmenAnggaran - totalRealisasi,
		Persentase_Realisasi:       Persentase_realisasi,
		Persentase_Belum_Realisasi: persentase_belum_realisasi,
		//data masih ngacoo
		Persentase_Realisasi_PDN_TKDN:      Persentase_realisasi_PDN_TKDN,
		Persentase_Belum_PDN_TKDN:          Persentase_belum_realisasi_PDN_TKDN,
		Persentase_Realisasi_belanja_impor: Persentase_realisasi_Impor,
		Persentase_Belum_belanja_impor:     persentase_belum_realisasi_impor,
		KomitmenTKDN:                       totalKomitmenTkdnPdn,
		KomitmenImpor:                      totalKomitmenImpor,
		Realisasi_TKDNPDN:                  totaTkdnPdn,
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

func GetDataDHAdminAP(c *fiber.Ctx) error {
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	if role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	if id_admin == 0 {
		return c.JSON(fiber.Map{
			"pesan": "Admin tidak terdaftar ",
		})
	}

	//Ambil data komitmen 2023, komitmen_keseluruhan
	type Komitmen struct {
		Nama_pekerjaan                   string
		Komitmen_anggaran_tahun_berjalan float64
		Nilai_kontrak_tahun              float64
	}
	id_users := c.Query("id_users")
	//bulan_buat := c.Query("bulanBuat")
	bulan_realisasi := c.Query("bulanRealisasi")
	//tahun_buat := c.Query("tahunBuat")
	tahun_realisasi := c.Query("tahunRealisasi")

	if id_users != "" {

		//Total Komitmen
		var komitmens []Komitmen
		if err := database.DB.Select("nama_pekerjaan, komitmen_anggaran_tahun_berjalan, nilai_kontrak_tahun").
			Where("id_user = ? ", id_users).
			Find(&komitmens).Error; err != nil {
			fmt.Println("Terjadi kesalahan:", err)
		}
	

		totalKomitmenAnggaran := 0.0
		for _, item := range komitmens {
			totalKomitmenAnggaran += item.Komitmen_anggaran_tahun_berjalan
		}

		totalRealisasi := 0.0
		for _, item := range komitmens {
			totalRealisasi += item.Nilai_kontrak_tahun
		}

		Persentase_realisasi := 0.0
		persentase_belum_realisasi := 0.0
		if totalKomitmenAnggaran != 0 {
			Persentase_realisasi = (totalRealisasi / totalKomitmenAnggaran) * 100
			persentase_belum_realisasi = 100 - Persentase_realisasi
		}

		// Second query
		var komitmen []models.Komitmen
		var komitmenss []models.Komitmen
		var komitmensss []models.Komitmen
		var PDN_TKDN []models.Komitmen

		if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
			Where("id_user = ? AND status_pencatatan = ? ", id_users, "TKDN").
			Find(&komitmen).Error; err != nil {
			fmt.Println("Terjadi kesalahantt:", err)
		}

		if bulan_realisasi != "" && tahun_realisasi != "" {

			database.DB.Select("nama_pekerjaan, komitmen_anggaran_tahun_berjalan,  nilai_kontrak_tahun").
				Where("id_user = ? AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, bulan_realisasi, tahun_realisasi).
				Find(&komitmens)

			database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan,  nilai_kontrak_tahun").
				Where("id_user = ? AND status_pencatatan = ? AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, "TKDN", bulan_realisasi, tahun_realisasi).
				Find(&komitmen)

		

			database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan,  nilai_kontrak_tahun").
				Where("id_user = ? AND status_pencatatan = ? AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, "PDN", bulan_realisasi, tahun_realisasi).
				Find(&komitmenss)

			database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
				Where("id_user = ? AND status_pencatatan = ? AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, "IMPOR", bulan_realisasi, tahun_realisasi).
				Find(&komitmensss)

			database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, komitmen_keseluruhan").
				Where("id_user = ? AND (status_pencatatan = ? OR status_pencatatan = ?) AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, "TKDN", "PDN", bulan_realisasi, tahun_realisasi).
				Find(&PDN_TKDN)
		}

		totalTKDN := 0.0
		for _, item := range komitmen {
			totalTKDN += item.NilaiKontrakTahun
		}

		// Third query
		if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
			Where("id_user = ? AND status_pencatatan = ?", id_users, "PDN").
			Find(&komitmenss).Error; err != nil {
			fmt.Println("Terjadi kesalahan:", err)
		}

		totalPDN := 0.0
		for _, item := range komitmenss {
			totalPDN += item.NilaiKontrakTahun
		}

		// Fourth query
		if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
			Where("id_user = ? AND status_pencatatan = ?", id_users, "IMPOR").
			Find(&komitmensss).Error; err != nil {
			fmt.Println("Terjadi kesalahan:", err)
		}

		totalImport := 0.0
		totalKomitmenImpor := 0.0
		for _, item := range komitmensss {
			totalImport += item.NilaiKontrakTahun
			totalKomitmenImpor += item.KomitmenAnggaranTahunBerjalan
		}

		Persentase_realisasi_Impor := 0.0
		persentase_belum_realisasi_impor := 0.0
		if totalKomitmenImpor != 0 {
			Persentase_realisasi_Impor = (totalImport / totalKomitmenImpor) * 100
			persentase_belum_realisasi_impor = 100 - Persentase_realisasi_Impor
		}

		// Fifth query

		if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
			Where("id_user = ? AND (status_pencatatan = ? OR status_pencatatan = ?)", id_users, "TKDN", "PDN").
			Find(&PDN_TKDN).Error; err != nil {
			fmt.Println("Terjadi kesalahan:", err)
		}

		totaTkdnPdn := 0.0
		totalKomitmenTkdnPdn := 0.0
		for _, item := range PDN_TKDN {
			totaTkdnPdn += item.NilaiKontrakTahun
			totalKomitmenTkdnPdn += item.KomitmenAnggaranTahunBerjalan
		}

		totalRealisasiPdnTkdn := 0.0
		for _, item := range PDN_TKDN {
			totalRealisasiPdnTkdn += item.NilaiKontrakTahun
		}

		Persentase_realisasi_PDN_TKDN := 0.0
		Persentase_belum_realisasi_PDN_TKDN := 0.0
		if totalKomitmenTkdnPdn != 0 {
			Persentase_realisasi_PDN_TKDN = (totaTkdnPdn / totalKomitmenTkdnPdn) * 100
			Persentase_belum_realisasi_PDN_TKDN = 100 - Persentase_realisasi_PDN_TKDN
		}

		type Dashboard struct {
			//dashboard atas
			Komitmen        float64
			KomitmenTKDN    float64
			KomitmenImpor   float64
			Realisasi       float64
			Belum_realisasi float64
			TKDN            float64
			PDN             float64
			Import          float64
			//grafik
			Persentase_Realisasi               float64
			Persentase_Belum_Realisasi         float64
			Persentase_Realisasi_belanja_impor float64
			Persentase_Belum_belanja_impor     float64

			Persentase_Realisasi_PDN_TKDN float64
			Persentase_Belum_PDN_TKDN     float64
			Realisasi_TKDNPDN             float64
		}
		data := Dashboard{
			Komitmen:                   totalKomitmenAnggaran,
			Realisasi:                  totalRealisasi,
			TKDN:                       totalTKDN,
			PDN:                        totalPDN,
			Import:                     totalImport,
			Belum_realisasi:            totalKomitmenAnggaran - totalRealisasi,
			Persentase_Realisasi:       Persentase_realisasi,
			Persentase_Belum_Realisasi: persentase_belum_realisasi,
			//data masih ngacoo
			Persentase_Realisasi_PDN_TKDN:      Persentase_realisasi_PDN_TKDN,
			Persentase_Belum_PDN_TKDN:          Persentase_belum_realisasi_PDN_TKDN,
			Persentase_Realisasi_belanja_impor: Persentase_realisasi_Impor,
			Persentase_Belum_belanja_impor:     persentase_belum_realisasi_impor,
			KomitmenTKDN:                       totalKomitmenTkdnPdn,
			KomitmenImpor:                      totalKomitmenImpor,
			Realisasi_TKDNPDN:                  totaTkdnPdn,
		}

		return c.JSON(fiber.Map{
			"data": data,
		})

	}

	//Untuk get data semuanya
	var komitmens []Komitmen
	var komitmen []models.Komitmen
	var komitmenss []models.Komitmen
	var komitmensss []models.Komitmen
	var PDN_TKDN []models.Komitmen
	if err := database.DB.Select("nama_pekerjaan, komitmen_anggaran_tahun_berjalan, komitmen_keseluruhan").Find(&komitmens).Error; err != nil {
		// Terjadi kesalahan saat mengeksekusi query
		// Anda dapat melakukan penanganan kesalahan di sini
		// Contoh: log kesalahan, kembalikan kesalahan, dll.
		fmt.Println("Terjadi kesalahanss:", err)
	}

	if bulan_realisasi != "" && tahun_realisasi != "" {

		database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, komitmen_keseluruhan").
			Where("id_user = ? AND status_pencatatan = ? AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, "TKDN", bulan_realisasi, tahun_realisasi).
			Find(&komitmen)

		database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, komitmen_keseluruhan").
			Where("id_user = ? AND status_pencatatan = ? AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, "PDN", bulan_realisasi, tahun_realisasi).
			Find(&komitmenss)

		database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, komitmen_keseluruhan", bulan_realisasi, tahun_realisasi).
			Where("id_user = ? AND status_pencatatan = ? AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, "IMPOR").
			Find(&komitmensss)

		database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, komitmen_keseluruhan").
			Where("id_user = ? AND (status_pencatatan = ? OR status_pencatatan = ?) AND bulan_realisasi =? AND tahun_realisasi = ?", id_users, "TKDN", "PDN", bulan_realisasi, tahun_realisasi).
			Find(&PDN_TKDN)
	}

	//menambah total Komitmen Anggaran Secara total
	totalKomitmenAnggaran := 0.0
	for _, item := range komitmens {
		totalKomitmenAnggaran += item.Komitmen_anggaran_tahun_berjalan
	}
	//menambah Nilai Kontrak tahun
	totalRealisasi := 0.0
	for _, item := range komitmens {
		totalRealisasi += item.Nilai_kontrak_tahun
	}

	Persentase_realisasi := ((totalRealisasi / totalKomitmenAnggaran) * 100)
	persentase_belum_realisasi := 100 - Persentase_realisasi

	//TKDN, PDN , IMPORT

	if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, komitmen_keseluruhan").
		Where("status_pencatatan = ?", "TKDN").
		Find(&komitmen).Error; err != nil {
		// Terjadi kesalahan saat mengeksekusi query
		// Anda dapat melakukan penanganan kesalahan di sini
		// Contoh: log kesalahan, kembalikan kesalahan, dll.
		fmt.Println("Terjadi kesalahan:", err)
	}

	//penambahan Total TKDN
	totalTKDN := 0.0

	for _, item := range komitmen {
		totalTKDN += item.NilaiKontrakTahun
	}
	//PDN

	if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
		Where("status_pencatatan = ?", "PDN").
		Find(&komitmenss).Error; err != nil {
		// Terjadi kesalahan saat mengeksekusi query
		// Anda dapat melakukan penanganan kesalahan di sini
		// Contoh: log kesalahan, kembalikan kesalahan, dll.
		fmt.Println("Terjadi kesalahan:", err)
	}

	//penambahan Total TKDN
	totalPDN := 0.0

	for _, item := range komitmenss {
		totalPDN += item.NilaiKontrakTahun
	}
	//untuk Impor

	if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
		Where("status_pencatatan = ?", "IMPOR").
		Find(&komitmensss).Error; err != nil {
		// Terjadi kesalahan saat mengeksekusi query
		// Anda dapat melakukan penanganan kesalahan di sini
		// Contoh: log kesalahan, kembalikan kesalahan, dll.
		fmt.Println("Terjadi kesalahan:", err)
	}

	// penambahan Total TKDN
	totalImport := 0.0
	totalKomitmenImpor := 0.0

	for _, item := range komitmensss {
		totalImport += item.NilaiKontrakTahun
		totalKomitmenImpor += item.KomitmenAnggaranTahunBerjalan
	}
	// persentasenya
	Persentase_realisasi_Impor := 0.0
	persentase_belum_realisasi_impor := 0.0
	if totalKomitmenImpor != 0 {
		Persentase_realisasi_Impor = (totalImport / totalKomitmenImpor) * 100
		persentase_belum_realisasi_impor = 100 - Persentase_realisasi_Impor
	}

	if err := database.DB.Select("komitmen_anggaran_tahun_berjalan, nilai_kontrak_keseluruhan, nilai_kontrak_tahun").
		Where("status_pencatatan = ? OR status_pencatatan = ?", "TKDN", "PDN").
		Find(&PDN_TKDN).Error; err != nil {
		// Terjadi kesalahan saat mengeksekusi query
		// Anda dapat melakukan penanganan kesalahan di sini
		// Contoh: log kesalahan, kembalikan kesalahan, dll.
		fmt.Println("Terjadi kesalahan:", err)
	}

	//perhitungan persentase
	totaTkdnPdn := 0.0
	totalKomitmenTkdnPdn := 0.0
	//realisasi
	for _, item := range PDN_TKDN {
		totaTkdnPdn += item.NilaiKontrakTahun
	}
	//komitmen
	for _, item := range PDN_TKDN {
		totalKomitmenTkdnPdn += item.KomitmenAnggaranTahunBerjalan
	}

	//persentase terealisasi TKDM
	Persentase_realisasi_PDN_TKDN := 0.0
	Persentase_belum_realisasi_PDN_TKDN := 0.0
	if totalKomitmenTkdnPdn != 0 {
		Persentase_realisasi_PDN_TKDN = (totaTkdnPdn / totalKomitmenTkdnPdn) * 100
		Persentase_belum_realisasi_PDN_TKDN = 100 - Persentase_realisasi_PDN_TKDN
	}

	type Dashboard struct {
		//dashboard atas
		Komitmen        float64
		KomitmenTKDN    float64
		KomitmenImpor   float64
		Realisasi       float64
		Belum_realisasi float64
		TKDN            float64
		PDN             float64
		Import          float64
		//grafik
		Persentase_Realisasi               float64
		Persentase_Belum_Realisasi         float64
		Persentase_Realisasi_belanja_impor float64
		Persentase_Belum_belanja_impor     float64

		Persentase_Realisasi_PDN_TKDN float64
		Persentase_Belum_PDN_TKDN     float64
		Realisasi_TKDNPDN             float64
	}
	data := Dashboard{
		Komitmen:                   totalKomitmenAnggaran,
		Realisasi:                  totalRealisasi,
		TKDN:                       totalTKDN,
		PDN:                        totalPDN,
		Import:                     totalImport,
		Belum_realisasi:            totalKomitmenAnggaran - totalRealisasi,
		Persentase_Realisasi:       Persentase_realisasi,
		Persentase_Belum_Realisasi: persentase_belum_realisasi,
		//data masih ngacoo
		Persentase_Realisasi_PDN_TKDN:      Persentase_realisasi_PDN_TKDN,
		Persentase_Belum_PDN_TKDN:          Persentase_belum_realisasi_PDN_TKDN,
		Persentase_Realisasi_belanja_impor: Persentase_realisasi_Impor,
		Persentase_Belum_belanja_impor:     persentase_belum_realisasi_impor,
		KomitmenTKDN:                       totalKomitmenTkdnPdn,
		KomitmenImpor:                      totalKomitmenImpor,
		Realisasi_TKDNPDN:                  totaTkdnPdn,
	}

	return c.JSON(fiber.Map{
		"data": data,
	})
}

// Pastikan Anda mengimpor packages yang diperlukan.

// Komitmen adalah struktur untuk data komitmen.
type Komitmen struct {
	Nama_pekerjaan                   string
	Komitmen_anggaran_tahun_berjalan float64
	Komitmen_keseluruhan             float64
}

// Dashboard adalah struktur untuk data yang akan dikirim sebagai response.
type Dashboard struct {
	Komitmen                           float64 `json:"komitmen"`
	Realisasi                          float64 `json:"realisasi"`
	Belum_realisasi                    float64 `json:"belum_realisasi"`
	TKDN                               float64 `json:"tkdn"`
	PDN                                float64 `json:"pdn"`
	Import                             float64 `json:"import"`
	Persentase_Realisasi               float64 `json:"persentase_realisasi"`
	Persentase_Belum_Realisasi         float64 `json:"persentase_belum_realisasi"`
	Persentase_Realisasi_belanja_impor float64 `json:"persentase_realisasi_belanja_impor"`
	Persentase_Belum_belanja_impor     float64 `json:"persentase_belum_belanja_impor"`
	Persentase_Realisasi_PDN_TKDN      float64 `json:"persentase_realisasi_pdn_tkdn"`
	Persentase_Belum_PDN_TKDN          float64 `json:"persentase_belum_pdn_tkdn"`
}

/*
// GetDataDHAdminAP adalah fungsi untuk mengambil data dashboard administratif.
func GetDataDHAdminAPss(c *fiber.Ctx) error {
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	if role != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar",
		})
	}
	if id_admin == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"pesan": "Admin tidak terdaftar",
		})
	}

	id_users := c.Query("id_users")
	bulan_realisasi := c.Query("bulanRealisasi")
	tahun_realisasi := c.Query("tahunRealisasi")

	// Fungsi helper untuk mengurangi duplikasi kode
	getKomitmenData := func(statusPencatatan string) (totalKontrakTahun, totalKomitmenAnggaran float64, err error) {
		var komitmens []Komitmen
		query := database.DB.Select("komitmen_anggaran_tahun_berjalan, komitmen_keseluruhan").
			Where("id_user = ? AND status_pencatatan = ?", id_users, statusPencatatan)
		if bulan_realisasi != "" && tahun_realisasi != "" {
			query = query.Where("bulan_realisasi = ? AND tahun_realisasi = ?", bulan_realisasi, tahun_realisasi)
		}
		err = query.Find(&komitmens).Error
		if err != nil {
			return 0, 0, err
		}
		for _, item := range komitmens {
			totalKontrakTahun += item.komitmen_keseluruhan
			totalKomitmenAnggaran += item.komitmen_anggaran_tahun_berjalan
		}
		return totalKontrakTahun, totalKomitmenAnggaran, nil
	}

	// Gunakan fungsi helper untuk mendapatkan data.
	totalTKDN, _, errTKDN := getKomitmenData("TKDN")
	if errTKDN != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"pesan": "Gagal mendapatkan data TKDN"})
	}

	totalPDN, _, errPDN := getKomitmenData("PDN")
	if errPDN != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"pesan": "Gagal mendapatkan data PDN"})
	}

	totalImport, totalKomitmenImpor, errImport := getKomitmenData("IMPOR")
	if errImport != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"pesan": "Gagal mendapatkan data Impor"})
	}

	// Perhitungan persentase realisasi.
	persentaseRealisasi := calculatePercentage(totalImport, totalKomitmenImpor)

	// Sisanya dari kode Anda untuk menyiapkan data untuk dashboard.
	// ...

	// Siapkan data untuk response.
	data := Dashboard{
		Komitmen:                   // ... nilai dari total komitmen anggaran,
		Realisasi:                  // ... nilai dari total realisasi,
		TKDN:                       totalTKDN,
		PDN:                        totalPDN,
		Import:                     totalImport,
		Belum_realisasi:            // ... nilai dari total belum realisasi,
		Persentase_Realisasi:       persentaseRealisasi,
		// ... dan seterusnya
	}

	// Kirim response dengan data.
	return c.JSON(fiber.Map{
		"data": data,
	})
}

// calculatePercentage adalah fungsi helper untuk menghitung persentase.
func calculatePercentage(value, total float64) float64 {
	if total == 0 {
		return 0
	}
	return (value / total) * 100
}

*/
