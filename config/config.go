package config

import (
	//"fmt"
	"os"
)

var (
	// RootDir of your app
	RootDir = "./web"

	// SecretKey computes sg_token
	SecretKey string

	// DB addr and passwd.
	DB string
)

func init() {
	// DB = fmt.Sprintf("%v:%v@tcp(mariadb:3306)/%v?charset=utf8",
	//      os.Getenv("DB_User"),
	//      os.Getenv("DB_Passwd"),
	//      os.Getenv("DB_Db"))
	DB = "kimpv:vKdBrBZG1qIloGicpFOR@tcp(10.3.80.60:3306)/hacking?charset=utf8"
	SecretKey = os.Getenv("SecretKey")
}
