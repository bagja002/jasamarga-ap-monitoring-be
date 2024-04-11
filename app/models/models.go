package models

type Admin struct {
	IdAdmin  uint `gorm:"primaryKey;autoIncrement"`
	Nama     string
	Username string
	Email    string
	Password string // Pastikan menggunakan tipe yang sesuai untuk menyimpan password, biasanya []byte atau string dengan hash.
}

type Users struct {
	IdUser   uint `gorm:"primaryKey;autoIncrement"`
	NamaAP   string
	Username string //utuk login
	Pic      string	//nama PIC
	Email    string	//Email_pIC
	No_pic   int	//nomor PIC
	Password string	 `json:"-"`//pasword PIC 
	Status   string // Pastikan menggunakan tipe yang sesuai untuk menyimpan password, biasanya []byte atau string dengan hash.

}

type Komitmen struct {
	IdKontrak                     uint `gorm:"primaryKey;autoIncrement"` // Nomor kontrak (misal: 1, 2, 3)
	IdUser                        uint
	NamaAP                        string  // Nama AP/SP (Aplikasi/SPJ)
	NamaPekerjaan                 string  // Nama Pekerjaan
	JenisPekerjaan                string  // Jenis Pekerjaan( Barang/Jasa Konsultasi/Kontruksi/Jasa Lain)
	JenisAnggaran                 string  // Jenis Anggaran (CAPEX/OPEX)
	KomitmenAnggaranTahunBerjalan          float64  `gorm:"type:float;precision:10"`// Komitmen Anggaran 2023
	KomitmenKeseluruhan 	float64 `gorm:"type:float;precision:10"`
	RencanaKualifikasiPenyedia string //Non Ukmkm or UMKM 
	StatusPadi	string //Pilihan Non PaDi UMKM or PaDi UMKM 
	RencanaWaktuMulaiPekerjaan string //Untuk tahun Berjalan //(Bualan Lur)
	RencanaTahunBerakhir uint
	CatatanKomitmen string //Catatan dari Komitmen 
	//Kapan di Buat untuk filter berdasarkan Waktu Pembuatan yaitu Bulan, Tahun 
	TahunBuat 						uint //Tahun Pembuatan
	BulanBuat 						string
	//-------------------------------
	TahunRealisasi				uint 
	BulanRealisasi				string
	//-------------------------------

	NilaiKontrakKeseluruhan       float64 `gorm:"type:float;precision:10"` // Nilai Kontrak Keseluruhan
	NilaiKontrakTahun             float64  `gorm:"type:float;precision:10"`// Nilai Kontrak Tahun 2023
	NamaPenyediaBarangDanJasa     string  // Nama Penyedia Barang dan Jasa
	KualifikasiPenyedia           string  // Kualifikasi Penyedia (UMKK/Menengah/Besar)
	StatusPencatatan              string  // Status Pencatatan (PDN/TKDN/IMPOR) Berubah menjadi Jenis Rencana Nilai Anggaran
		
	PersentasePDN                 float64 `gorm:"type:float;precision:10"` // Persentase PDN (%)
	PersentaseTKDN                float64 `gorm:"type:float;precision:10"`// Persentase TKDN (%)
	PersentaseImpor               float64 `gorm:"type:float;precision:10"` // Persentase Impor (%)
	TotalBobot                    float64 `gorm:"type:float;precision:10"` // Total Bobot (%)
	RealisasiWaktuMulaiKontrak    string  // Realisasi Waktu Mulai Kontrak Berubah menjadi Rencana Waktu Mulai Pekerjaan Tahun 2024(Bulan) 
	RealisasiWaktuBerakhirKontrak string  // Realisasi Waktu Berakhir Kontrak (Rencana Tahun Berakhir Pekerjaan)
	KeteranganLainnya             string  // Keterangan Lainnya
	Status     					uint 	//Jika = 1 dia baru di buat, jika 2 dia sudah selesai di edit 
	Is_active                     uint // keterangan Jika 0 = tidak aktifve, jika 1 aktive, 2 selesai
}

//Pemisahan Kolom Komitmen

type Komitmensss struct {
	IdKontrak                     uint `gorm:"primaryKey;autoIncrement"` // Nomor kontrak (misal: 1, 2, 3)
	IdUser                        uint
	NamaAP                        string  // Nama AP/SP (Aplikasi/SPJ)
	NamaPekerjaan                 string  // Nama Pekerjaan
	JenisPekerjaan                string  // Jenis Pekerjaan
	JenisAnggaran                 string  // Jenis Anggaran
	KomitmenAnggaran2023          float64 // Komitmen Anggaran 202       
	Status 						uint
	Is_active                     uint    // keterangan Jika 0 = tidak aktifve, jika 1 aktive, 2 selesai
}


type Realisasi struct {
	IdRealisasi                   uint `gorm:"primaryKey;autoIncrement"`
	IdUser                        uint
	IdKontrak                     uint
	NilaiKontrakKeseluruhan       float64 // Nilai Kontrak Keseluruhan
	NilaiKontrakTahun             float64 // Nilai Kontrak Tahun 2023
	NamaPenyediaBarangDanJasa     string  // Nama Penyedia Barang dan Jasa
	KualifikasiPenyedia           string  // Kualifikasi Penyedia (UMKK/Menengah/Besar)
	StatusPencatatan              string  // Status Pencatatan (PDN/TKDN/IMPOR)
	PersentasePDN                 float64 // Persentase PDN (%)
	PersentaseTKDN                float64 // Persentase TKDN (%)
	PersentaseImpor               float64 // Persentase Impor (%)
	TotalBobot                    float64 // Total Bobot (%)
	RealisasiWaktuMulaiKontrak    string  // Realisasi Waktu Mulai Kontrak
	RealisasiWaktuBerakhirKontrak string  // Realisasi Waktu Berakhir Kontrak
	KeteranganLainnya             string  // Keterangan Lainnya
	Status 							uint //Jika keterangan 
	Is_active                     uint    // keterangan Jika 0 = tidak aktifve, jika 1 aktive, 2 selesai

}

type Komitmen2 struct {
	IdKontrak                     uint `gorm:"primaryKey;autoIncrement"` // Nomor kontrak (misal: 1, 2, 3)
	IdUser                        uint
	NamaAP                        string  // Nama AP/SP (Aplikasi/SPJ)
	NamaPekerjaan                 string  // Nama Pekerjaan
	JenisPekerjaan                string  // Jenis Pekerjaan( Barang/Jasa Konsultasi/Kontruksi/Jasa Lain)
	JenisAnggaran                 string  // Jenis Anggaran (CAPEX/OPEX)
	KomitmenAnggaranTahunBerjalan          float64 // Komitmen Anggaran 2023
	KomitmenKeseluruhan 	float64
	RencanaKualifikasiPenyedia string //Non Ukmkm or UMKM 
	StatusPadi	string //Pilihan Non PaDi UMKM or PaDi UMKM 
	RencanaWaktuMulaiPekerjaan string //Untuk tahun Berjalan //(Bualan Lur)
	RencanaTahunBerakhir uint
	CatatanKomitmen string //Catatan dari Komitmen 
	//Kapan di Buat untuk filter berdasarkan Waktu Pembuatan yaitu Bulan, Tahun 
	TahunBuat 						uint //Tahun Pembuatan
	BulanBuat 						string
	//-------------------------------
	TahunRealisasi				uint 
	BulanRealisasi				string
	//-------------------------------

	NilaiKontrakKeseluruhan       float64 // Nilai Kontrak Keseluruhan
	NilaiKontrakTahun             float64 // Nilai Kontrak Tahun 2023
	NamaPenyediaBarangDanJasa     string  // Nama Penyedia Barang dan Jasa
	KualifikasiPenyedia           string  // Kualifikasi Penyedia (UMKK/Menengah/Besar)
	StatusPencatatan              string  // Status Pencatatan (PDN/TKDN/IMPOR) Berubah menjadi Jenis Rencana Nilai Anggaran
		
	PersentasePDN                 float64 // Persentase PDN (%)
	PersentaseTKDN                float64 // Persentase TKDN (%)
	PersentaseImpor               float64 // Persentase Impor (%)
	TotalBobot                    float64 // Total Bobot (%)
	RealisasiWaktuMulaiKontrak    string  // Realisasi Waktu Mulai Kontrak Berubah menjadi Rencana Waktu Mulai Pekerjaan Tahun 2024(Bulan) 
	RealisasiWaktuBerakhirKontrak string  // Realisasi Waktu Berakhir Kontrak (Rencana Tahun Berakhir Pekerjaan)
	KeteranganLainnya             string  // Keterangan Lainnya
	Status     					uint 	//Jika = 1 dia baru di buat, jika 2 dia sudah selesai di edit 
	Is_active                     uint    // keterangan Jika 0 = tidak aktifve, jika 1 aktive, 2 selesai
}