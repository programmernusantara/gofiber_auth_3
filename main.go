package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func main() {
	// Inisialisasi aplikasi Fiber
	app := fiber.New()

	// Map untuk menyimpan informasi pengguna yang terdaftar
	users := map[string]string{
		"wildan": "190205",
	}

	// Route untuk halaman utama
	app.Get("/", func(c *fiber.Ctx) error {
		// Kirimkan pesan "Selamat Datang" ke browser
		return c.SendString("Selamat Datang")
	})

	// Route untuk pendaftaran pengguna baru
	app.Post("/daftar", func(c *fiber.Ctx) error {
		// Struktur data untuk menerima data pengguna dari request
		var newUser struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		// Parse data JSON dari request
		if err := c.BodyParser(&newUser); err != nil {
			// Jika gagal memproses data JSON, kembalikan pesan error
			return c.Status(fiber.StatusBadRequest).JSON("Gagal memproses data daftar")
		}

		// Tambahkan pengguna baru ke dalam map users
		users[newUser.Username] = newUser.Password

		// Kembalikan pesan sukses pendaftaran
		return c.JSON("Pengguna " + newUser.Username + " berhasil terdaftar")
	})

	// Middleware untuk autentikasi dasar
	app.Use(basicauth.New(basicauth.Config{
		Users:           users,
		ContextUsername: "user",
	}))

	// Route untuk login
	app.Get("/login", func(c *fiber.Ctx) error {
		// Dapatkan nama pengguna yang sudah login
		user, _ := c.Locals("user").(string)

		// Kembalikan pesan sukses login
		return c.JSON("Selamat " + user + ", Anda berhasil login")
	})

	// Route untuk logout
	app.Get("/logout", func(c *fiber.Ctx) error {
		// Mendapatkan nama pengguna yang berhasil logout dari konteks
		user, _ := c.Locals("user").(string)

		// Hapus informasi pengguna dari konteks untuk logout
		c.Locals("user", nil)

		// Menampilkan pesan logout dengan nama pengguna
		return c.SendString("Anda berhasil logout. Pengguna: " + user)
	})

	// Menjalankan aplikasi pada port 8080
	app.Listen(":8080")
}
