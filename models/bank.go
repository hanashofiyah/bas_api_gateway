package models

type Bank struct {
	BankID  string `gorm:"primarykey"`
	Name    string `gorm:"column:name"`
	Address string
}

func (a *Bank) TableName() string {
	return "bank"
}
