package database

import (
	"canteen-prakerja/entity"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dialect = "mysql"

	// uncomment this for PGsql ruby on rails
	// host     = os.Getenv("POSTGRES_HOST")
	// port     = os.Getenv("POSTGRES_PORT")
	// username = os.Getenv("POSTGRES_USERNAME")
	// password = os.Getenv("POSTGRES_PASSWORD")
	// dbname   = os.Getenv("POSTGRES_DATABASE")
	// ssl      = os.Getenv("POSTGRES_SSLMODE")

	// uncomment this for MySQL
	host     = os.Getenv("MYSQL_HOST")
	port     = os.Getenv("MYSQL_PORT")
	username = os.Getenv("MYSQL_USERNAME")
	password = os.Getenv("MYSQL_PASSWORD")
	dbname   = os.Getenv("MYSQL_DATABASE")

	//uncomment variable bellow for local using .env file
	// host     = goDotEnvVariable("PG_HOST")
	// port     = goDotEnvVariable("PG_PORT")
	// username = goDotEnvVariable("PG_USERNAME")
	// password = goDotEnvVariable("PG_PASSWORD")
	// dbname   = goDotEnvVariable("PG_DBNAME")
	// ssl      = goDotEnvVariable("PG_SSLMODE")
)

var (
	db  *gorm.DB
	err error
)

func handleDatabaseConnection() {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Panicf("error while getting .env file : %s", err.Error())
	}

	//Connect Database
	// _ := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Jakarta", host, port, username, password, dbname, ssl)
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, dbname)
	dsn := os.Getenv("MYSQL_ONLINE_URL")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Discard.LogMode(logger.Silent),
	})

	if err != nil {
		log.Panicf("error while connecting to db : %s", err.Error())
	}
	fmt.Print("Server Started")
}

func handleCreateRequiredTable() {
	err = db.Debug().AutoMigrate(&entity.User{}, &entity.Barang{}, &entity.Transaksi{}, &entity.Item{})
	if err != nil {
		log.Panicln(err.Error())
	}
}

func handleNewDefaultUser() error {
	user := entity.User{
		Username: "admin_kantin",
		Password: "12345678",
	}

	err = user.HashPassword()

	if err != nil {
		return err
	}

	res := db.FirstOrCreate(&user)

	if res.Error != nil {
		return res.Error
	}

	return nil

}

func InitializeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTable()
	handleNewDefaultUser()
}

func GetDatabaseInstance() *gorm.DB {
	return db
}
