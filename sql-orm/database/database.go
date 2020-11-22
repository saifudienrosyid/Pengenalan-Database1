package database

import (
	"log"

	"gorm.io/gorm"
)

type CustomerORM struct {
	ID           int    `gorm:"primary_key" json:"customer_id"`
	NpwpId       string `json:"npwp_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Age          int    `json:"age"`
	CustomerType string `json:"customer_type"`
	Street       string `json:"street"`
	City         string `json:"city"`
	State        string `json:"state"`
	ZipCode      string `json:"zip_code"`
	PhoneNumber  string `json:"phone_number"`
	// 1 akun CustomerORM dapat memiliki banyak akun AccountORM
	AccountORM []AccountORM `gorm:"ForeignKey:IdCustomerRefer" json:"account_orm"`
}

type AccountORM struct {
	ID              int    `gorm:"primary_key" json:"-"`
	IdCustomerRefer int    `json:"-"`
	Balance         int    `json: "balance"`
	AccountType     string `json:"account_type"`
}

func InsertCustomer(customer CustomerORM, db *gorm.DB) {
	if err := db.Create(&customer).Error; err != nil {
		log.Println("Failed to insert : ", err.Error())
		return
	}
	log.Println("Success insert data!")
}

func GetCustomers(db *gorm.DB) {
	var customer []CustomerORM
	if err := db.Preload("AccountORM").Find(&customer).Error; err != nil {
		log.Println("Failed to get data", err.Error())
		return
	}
	log.Println(customer)
}

func DeleteCustomer(id int, db *gorm.DB) {
	var customer CustomerORM

	if err := db.Where(&CustomerORM{ID: id}).Delete(&customer).Error; err != nil {
		log.Println("Failed to delete data : ", err.Error())
		return
	}
	log.Println("Success delete data!")
}

func UpdateCustomer(customer CustomerORM, id int, db *gorm.DB) {
	if err := db.Model(&CustomerORM{}).Where(&CustomerORM{ID: id}).Updates(customer).Error; err != nil {
		log.Println("Failed to update data : ", err.Error())
		return
	}
	log.Println("Success update data!")
}
