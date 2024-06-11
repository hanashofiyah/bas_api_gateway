package handler

import (
	"apigateway/models"
	"apigateway/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransferInterface interface {
	TransferBank(*gin.Context)
}

type transferImplement struct{}

func NewTransfer() TransferInterface {
	return &transferImplement{}
}

type BodyPayloadTransfer struct{}

func (b *transferImplement) TransferBank(g *gin.Context) {
	bodyPayload := models.Transaction{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Create(&bodyPayload)

	if result.Error != nil {
		fmt.Println(err)
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": `Transaction successfully`,
		"data":    bodyPayload,
	})
}
