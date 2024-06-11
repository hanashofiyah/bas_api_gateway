package models

type Account struct {
	AccountID string ` gorm:"primaryKey" ` //account_id
	Name      string ` gorm:"column:username" `
	Username  string //ini otomatis membacanya huruf password
	Password  string
}

func (a *Account) TableName() string {
	return "account"
}
