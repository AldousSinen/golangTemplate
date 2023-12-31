// Package connect provides ...
package config

import (
	"Template/pkg/utils"
	"Template/pkg/utils/go-utils/database"
	"Template/pkg/utils/go-utils/encryptDecrypt"
	"fmt"

	"log"
)

func CreateConnection() {
	username, dbConfigErr := encryptDecrypt.Decrypt(utils.GetEnv("POSTGRES_USERNAME"), utils.GetEnv("SECRET_KEY"))
	if dbConfigErr != nil {
		fmt.Println("error encrypting your classified text: ", dbConfigErr)
	}
	password, dbConfigErr := encryptDecrypt.Decrypt(utils.GetEnv("POSTGRES_PASSWORD"), utils.GetEnv("SECRET_KEY"))
	if dbConfigErr != nil {
		fmt.Println("error encrypting your classified text: ", dbConfigErr)
	}
	host, dbConfigErr := encryptDecrypt.Decrypt(utils.GetEnv("POSTGRES_HOST"), utils.GetEnv("SECRET_KEY"))
	if dbConfigErr != nil {
		fmt.Println("error encrypting your classified text: ", dbConfigErr)
	}
	dbName, dbConfigErr := encryptDecrypt.Decrypt(utils.GetEnv("DATABASE_NAME"), utils.GetEnv("SECRET_KEY"))
	if dbConfigErr != nil {
		fmt.Println("error encrypting your classified text: ", dbConfigErr)
	}
	fmt.Println("username: ", username)
	fmt.Println("password: ", password)
	fmt.Println("host: ", host)
	fmt.Println("dbName: ", dbName)
	database.PostgreSQLConnect(
		utils.GetEnv("POSTGRES_USERNAME"),
		utils.GetEnv("POSTGRES_PASSWORD"),
		utils.GetEnv("POSTGRES_HOST"),
		utils.GetEnv("DATABASE_NAME"),
		utils.GetEnv("POSTGRES_PORT"),
		utils.GetEnv("POSTGRES_SSL_MODE"),
		utils.GetEnv("POSTGRES_TIMEZONE"),
	)
	err := database.DBConn.AutoMigrate()

	if err != nil {
		log.Fatal(err.Error())
	}

}
