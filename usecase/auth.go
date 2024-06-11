package usecase

import (
	"apigateway/models"
	"apigateway/utils"
)

type autentifikasiImplement struct {
}

func NewAutentifikasi() LoginInterface {
	return &autentifikasiImplement{}
}

type LoginInterface interface {
	Autentifikasi(username, password string) bool
}

func (a *autentifikasiImplement) Autentifikasi(username, password string) bool {
	// if username == "admin" && password == "admin123" {
	// 	return true
	// }
	// return false
	account := models.Account{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	orm.Find(&account, "username = ? AND password = ?", username, password)
	return account.AccountID != ""
}
