package handler

import (
	"apigateway/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AutentifikasiInterface interface {
	AutentifikasiAccount(*gin.Context)
}

type autentifikasiImplement struct{}

func NewAutentifikasi() AutentifikasiInterface {
	return &autentifikasiImplement{}
}

type BodyPayloadAutentifikasi struct {
	Username string
	Password string
}

func (a *autentifikasiImplement) AutentifikasiAccount(g *gin.Context) {

	bodyPayloadAuth := BodyPayloadAutentifikasi{}
	err := g.BindJSON(&bodyPayloadAuth)

	login := usecase.NewAutentifikasi().Autentifikasi(bodyPayloadAuth.Username, bodyPayloadAuth.Password)

	if login {
		g.JSON(http.StatusOK, gin.H{
			"message": "Anda berhasil login",
			"data":    bodyPayloadAuth,
		})
	} else {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Anda gagal login",
			"data":    err,
		})
	}
}
