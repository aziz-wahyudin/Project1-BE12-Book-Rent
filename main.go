package main

import (
	"fmt"
	"os"
	"os/exec"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func connectGorm() (*gorm.DB, error) {
	dsn := "root:@tcp(127.0.0.1:3306)/Book_Rent?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func callClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func migrate(db *gorm.DB) {
	db.AutoMigrate()
}

func main() {
	gconn, err := connectGorm()
	migrate(gconn)
	if err != nil {
		fmt.Println("cannot connect to DB", err.Error())
	}
	var isRunning bool = true
	var inputMenu int

	for isRunning {
		fmt.Println("\t--Menu--")
		fmt.Println("1. Tampilkan Semua Buku")
		fmt.Println("2. Login Atau Register")
		fmt.Println("3. Exit")
		fmt.Println("Masukkan Input: ")
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case 1:
		case 2:
			var menuDua bool = true
			for menuDua {
				fmt.Println("==============")
				fmt.Println("1. Registrasi")
				fmt.Println("2. Login")
				fmt.Println("3. Exit")
				fmt.Scanln(&inputMenu)
				switch inputMenu {
				case 1:
					var menuRegistrasi bool = true
					for menuRegistrasi {
						fmt.Println("==============")
						fmt.Println("1. Tampilkan Buku Saya")
						fmt.Println("2. Tambah Koleksi Buku")
						fmt.Println("3. Ubah Koleksi Buku")
						fmt.Println("4. Pinjam Buku")
						fmt.Println("5. Buku Pinjaman Saya")
						fmt.Println("6. Kembalikan Buku")
						fmt.Println("7. Update Akun")
						fmt.Println("8. Hapus akun")
						fmt.Println("9. Exit")
						fmt.Scanln(&inputMenu)
						switch inputMenu {
						case 1:
						case 2:
						case 3:
						case 4:
						case 5:
						case 6:
						case 7:
						case 8:
						case 9:
							callClear()
							menuRegistrasi = false
						}
					}
				case 2:
					var menuLogin bool = true
					for menuLogin {
						fmt.Println("==============")
						fmt.Println("1. Tampilkan Buku Saya")
						fmt.Println("2. Tambah Koleksi Buku")
						fmt.Println("3. Ubah Koleksi Buku")
						fmt.Println("4. Pinjam Buku")
						fmt.Println("5. Buku Pinjaman Saya")
						fmt.Println("6. Kembalikan Buku")
						fmt.Println("7. Update Akun")
						fmt.Println("8. Hapus akun")
						fmt.Println("9. Exit")
						fmt.Scanln(&inputMenu)
						switch inputMenu {
						case 1:
						case 2:
						case 3:
						case 4:
						case 5:
						case 6:
						case 7:
						case 8:
						case 9:
							callClear()
							menuLogin = false

						}
					}
				case 3:
					callClear()
					menuDua = false
				}
			}
		case 3:
			callClear()
			isRunning = false
		}
	}
}
