package debugs


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
	Pic      string
	Email    string
	No_pic   int
	Password string
	Status   string // Pastikan menggunakan tipe yang sesuai untuk menyimpan password, biasanya []byte atau string dengan hash.

}
/*
type Komitmen struct {
	IdKontrak                     uint `gorm:"primaryKey;autoIncrement"` // Nomor kontrak (misal: 1, 2, 3)
	IdUser                        uint
	NamaAP                        string  // Nama AP/SP (Aplikasi/SPJ)
	NamaPekerjaan                 string  // Nama Pekerjaan
	JenisPekerjaan                string  // Jenis Pekerjaan
	JenisAnggaran                 string  // Jenis Anggaran
	KomitmenAnggaran2023          float64 // Komitmen Anggaran 2023
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
	Status     					uint 	//Jika = 1 dia baru di buat, jika 2 dia sudah selesai di edit 
	Is_active                     uint    // keterangan Jika 0 = tidak aktifve, jika 1 aktive, 2 selesai
}

*/


type Komitmen struct {
	IdKontrak                     uint `gorm:"primaryKey;autoIncrement"` // Nomor kontrak (misal: 1, 2, 3)
	IdUser                        uint
	NamaAP                        string  // Nama AP/SP (Aplikasi/SPJ)
	NamaPekerjaan                 string  // Nama Pekerjaan
	JenisPekerjaan                string  // Jenis Pekerjaan
	JenisAnggaran                 string  // Jenis Anggaran
	KomitmenAnggaran2023          float64 // Komitmen Anggaran 202 
	StatusPencatatan			string      
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
