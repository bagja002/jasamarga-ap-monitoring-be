package controlles

import (
	"e-monitoring/app/models"
	"e-monitoring/pkg/database"

	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const role_admin = 1
const role_user = 2


func GetAdmin(c *fiber.Ctx)error{
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

	var admin models.Admin
	database.DB.Where("id_admin = ?", id_admin).Find(&admin)

	if admin.IdAdmin == 0 {
		c.SendStatus(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"pesan": "Tidak di temumukan admin di database ",
		})
	}

	return c.JSON(fiber.Map{
		"data":admin,
	})
}


func RegistAdmin(c *fiber.Ctx)error {

	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)

	admin:= models.Admin{
		Nama: "admin",
		Password: string(password),
		Email: "admin@gmail.com",
		Username: "admin",

	}
	database.DB.Create(&admin)


	return c.JSON(fiber.Map{
		"pesan":"Registrasi telah berhasil", 
	})
}


func LoginAdmin(c *fiber.Ctx)error{

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.JSON(fiber.Map{
			"pesan": err.Error(),
		})
	}

	var admin models.Admin

	database.DB.Where("username = ? ", data["username"]).First(&admin)
	if admin.IdAdmin == 0 {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"pesan": "Username tidak di temukan",
		})
	}

	if err:= bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"pesan": "Incorrect password!",
		})
	} else{
		claims := jwt.MapClaims{
			"name":     admin.Nama,
			"id_admin": admin.IdAdmin,
			"role":     role_admin,
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
