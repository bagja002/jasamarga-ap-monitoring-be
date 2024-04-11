package database

import (
	"e-monitoring/app/models"
	//"e-monitoring/app/models/debugs"
	"fmt"

	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/joho/godotenv"
	"os"

)

var DB *gorm.DB

/*
func Connect() {

		dsn := "root:@tcp(127.0.0.1:3306)/Learning_online?charset=utf8mb4&parseTime=True&loc=Local"

		con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Println("Gagal Meng koneksikan database")
		}
		fmt.Println("database terkoneksi")
		DB = con
		err1:= DB.AutoMigrate(models.User{}, models.Admin{}, models.Course{}, models.Categori{})
		if err1!= nil {
	        fmt.Println("gagal Auto Migration")
	    } else {
			fmt.Println("berhasil Auto Migration")
		}

}*
*/
func Connect() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Menggunakan variabel untuk membuat DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// Membuat koneksi ke database
	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Melakukan migration pada mode
	con.AutoMigrate(&models.Admin{}, &models.Users{}, &models.Komitmen{}, &models.Komitmen2{})

	fmt.Println("Database terkoneksi")
	DB = con
}

func Connectdev() {

	dbUser := "root"
	dbPass := ""
	dbHost := "localhost"
	dbPort := "3306"
	dbName := "apmonitdev"

	// Menggunakan variabel untuk membuat DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)

	// Membuat koneksi ke database
	con, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Melakukan migration pada mode
	//con.AutoMigrate(models.Admin{}, models.Users{}, models.Komitmen{})

	con.AutoMigrate(models.Admin{}, models.Users{}, models.Komitmen{}, models.Komitmen2{}) //debugs.Realisasi{} )

	fmt.Println("Database terkoneksi")
	DB = con
}
