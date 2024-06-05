package usecase

type login struct {
}

type LoginInterface interface {
	Authentifikasi(username, password string) bool
}

func (a *login) Authentifikasi(username, password string) bool {
	if username == "admin" && password == "admin123" {
		return true
	}
	return false
}

func NewLogin() LoginInterface {
	return &login{}
}
