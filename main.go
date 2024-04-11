package main

import (
	"e-monitoring/app/controlles"
	"e-monitoring/pkg/database"
	"e-monitoring/pkg/middleware"

	"log"
	

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {

	app := fiber.New()
	//use Midelware

	database.Connect()


	app.Use(
		cors.New(),
		logger.New(),
	)


	//route

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Selamat Datang di aplikasi backend e-Monitoring jasamarga ")
	})
	app.Get("/1",controlles.RegistAdmin )
	
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page"}))

	//group route

	admin := app.Group("/admin")
	users := app.Group("/users")
	//admin Route
	admin.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Iniaalah halaman Admin ")
	})

	admin.Post("/login", controlles.LoginAdmin)
	admin.Get("/getadmin",middleware.JwtProtect(), controlles.GetAdmin )
	admin.Get("/buat_user", middleware.JwtProtect(), controlles.RegistUsers)
	admin.Get("/getDashboard", middleware.JwtProtect(), controlles.GetDataDHAdmin)
	admin.Get("/dummy", controlles.CreateDummyKomitmen)
	users.Get("/getDasboardap", middleware.JwtProtect(), controlles.GetDataDHAdminAP)
	admin.Get("/all_user", middleware.JwtProtect(), controlles.GetAllUser)
	admin.Post("/registerUsers", middleware.JwtProtect(), controlles.RegistUsers)
	admin.Get("/download", controlles.Download)

	//users Route
	users.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Iniaalah halaman Users ")
	})

	users.Post("/login", controlles.LoginUsers)
	users.Post("/buat_komitmen", middleware.JwtProtect(),controlles.CreateKomitmen)
	users.Get("/getKomitmen", middleware.JwtProtect(), controlles.GetKomitmen)
	users.Put("/update", middleware.JwtProtect(), controlles.UpdateKomitmen)
	users.Post("/add_realisasi", middleware.JwtProtect(), controlles.AddRealisasi)
	users.Put("/update_users", middleware.JwtProtect(), controlles.Updateuser)
	users.Post("/forgetPassword", controlles.ForgetPassword)

	//Test Route 
	app.Post("/testBuat", middleware.JwtProtect(), controlles.DevBuatKom)

	log.Fatal(app.Listen(":4000"))


}
