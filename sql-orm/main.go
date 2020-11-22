package main

import (
	"fmt"
	"log"

	"github.com/b-rachman/pengenalan-database1/sql-orm/config"
	"github.com/b-rachman/pengenalan-database1/sql-orm/database"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg, err := getConfig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := initDB(cfg.Database)
	if err != nil {
		log.Println(err)
		return
	}

	/*database.InsertCustomer(database.CustomerORM{
		NpwpId:       "npwp123",
		FirstName:    "David",
		LastName:     "Sebastian",
		Age:          25,
		CustomerType: "Premium",
		Street:       "Jl. Let. Karjono",
		City:         "Banjarnegara",
		State:        "Indonesia",
		ZipCode:      "53412",
		PhoneNumber:  "082200000000",
		AccountORM: []database.AccountORM{
			{
				Balance:     10000,
				AccountType: "Pertamax",
			}, {
				Balance:     5000,
				AccountType: "Premium",
			},
		},
	}, db)*/

	//database.GetCustomers(db)
	database.UpdateCustomer(database.CustomerORM{
		FirstName: "Clare",
		Age:       30,
		City:      "Banyumas",
	}, 1, db)
}

func getConfig() (config.Config, error) {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return config.Config{}, err
	}

	var cfg config.Config
	err := viper.Unmarshal(&cfg)
	if err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func initDB(dbConfig config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.DbName, dbConfig.Config)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Membuat table baru di database pengenalan_database, nama tabel merefer ke nama struct (CustomerORM, dan AccountORM)
	// yang nantinya nama tabel yang terbuat berubah menjadi huruf kecil semua dan berupa kata jamak (ada tamabhan huruf 's')
	db.AutoMigrate(&database.CustomerORM{}, &database.AccountORM{})

	log.Println("db successfully connected")

	return db, nil
}
