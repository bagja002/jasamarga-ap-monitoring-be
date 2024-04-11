package controlles

import (
	"e-monitoring/app/models"
	"e-monitoring/pkg/database"
	"e-monitoring/pkg/tools"


	"strconv"

	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)


func GetAllUser(c *fiber.Ctx)error {
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	if role != 1 && role != 2 {
		return c.JSON(fiber.Map{
			"pesan": "role tidak sesuai masukan role dengan benar ",
		})
	}
	if id_admin == 0 {
		return c.JSON(fiber.Map{
			"pesan": "Users tidak terdaftar ",
		})
	}

	var users []models.Users
	database.DB.Find(&users)

	

	return  c.JSON(fiber.Map{
		"data":users,
	})
}




func Getuser(c *fiber.Ctx)error{
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

	var users models.Users
	database.DB.Where("id_users = ?", id_admin).Find(&users)

	if users.IdUser == 0 {
		c.SendStatus(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"pesan": "Tidak di temumukan admin di database ",
		})
	}

	return c.JSON(fiber.Map{
		"data":users,
	})
}


func RegistUsers(c *fiber.Ctx)error {
	
	id_admin, _ := c.Locals("id_admin").(int)
	role, _ := c.Locals("role").(float64)

	if role != 	1 {
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

	No_pic, err := strconv.Atoi(data["no_pic"])
	if err != nil {
		// Penanganan kesalahan jika konversi gagal
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"pesan": "Gagal mengonversi no_pic ke tipe int",
		})
	}




	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)

	users:= models.Users{
		NamaAP: data["nama_ap"],
		Username: data["username"],
		Password: string(password),
		Email: data["email"],
		Pic: data["pic"],
		No_pic: No_pic,
		Status: "active",
	}
	database.DB.Create(&users)


	return c.JSON(fiber.Map{
		"pesan":"Pembuatan Akun AP Telah Berhasil", 
	})
}


func LoginUsers(c *fiber.Ctx)error{

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(fiber.Map{
			"pesan": err.Error(),
		})
	}

	var users models.Users

	database.DB.Where("username = ? ", data["username"]).First(&users)
	if users.IdUser == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"pesan": "Username tidak di temukan",
		})
	}

	if err:= bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"pesan": "Incorrect password!",
		})
	} else{
		claims := jwt.MapClaims{
			"name":     users.NamaAP,
			"id_admin": users.IdUser,
			"role":     role_user,
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		
		return c.JSON(fiber.Map{
			"token": t,
		})
	}
	
}


func Updateuser(c *fiber.Ctx)error{

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

	err := c.BodyParser(&data)
	if err!= nil {
		return c.JSON(fiber.Map{
			"Pesan":"Masukan data dengan benar",
		})
	}

	//mengambil modl users
	var users models.Users
	
	//ambil data user beedasarkan id adin dan simpan di var users

	database.DB.Where("id_user = ?", id_admin).Find(&users)
	
	//setelah data berada di usesr maka baut update untuk data terbaru nya 

	//siapkan data yang terbaru nya 

	no_pic, _ := strconv.ParseFloat(data["no_pic"], 64)
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)

	updates:= models.Users{
		NamaAP: data["nama_ap"],
		Pic: data["pic"],
		No_pic: int(no_pic),
		Password: string(password),
		Username: data["username"],
		Email: data["email"],
	}

	//update data setelah perubahan yang ada 

	database.DB.Model(&users).Updates(updates)

	return c.JSON(fiber.Map{
		"Pesan":"data users telah di perbarui",
		"data":users,
	})
}

func ForgetPassword(c *fiber.Ctx)error{

	/*
	Di menu Forget password logic nya seperti ini 
	1. Ada halaman Forget Password untuk user
	2. User Hanya Pemasukan Username nya dan juga email nya 
	3. jika email dan username nya ada di database maka akan di ubah password nya menjadi 1234

	*/ 
	
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(fiber.Map{
			"pesan": err.Error(),
		})
	}
	//simpan database yang udh di temukan username dan passwordnya sebanyak 5 digit 
	var users models.Users
	email := data["email"]
	username := data["username"]
	
	result := database.DB.Where("username = ? AND email = ?", username, email).Find(&users)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"Pesan": "data tidak cocok",
		})
	}
	
	//perubahan Password menjadi angka acak 
	newPassword:= tools.GenerateRandomPassword(5)
	password, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	update:= models.Users{
		Password: string(password),
	}


	//update database sesuai dengan usersnya 
	//
	database.DB.Model(&users).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan":"Password Telah di reset, silahkan memakai password sementara dan Update di profile",
		"NewPassword":newPassword,
	})	
}