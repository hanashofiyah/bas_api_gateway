package handler

import (
	"apigateway/models"
	"apigateway/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountInterface interface {
	GetAccount(*gin.Context)
	CreateAccount(*gin.Context)
	UpdateAccount(*gin.Context)
	DeleteAccount(*gin.Context)
	BalanceAccount(*gin.Context)
}

type accountImplement struct{}

func NewAccount() AccountInterface {
	return &accountImplement{}
}

func (a *accountImplement) GetAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	name := queryParam.Get("name")

	accounts := []models.Account{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	// q := orm
	// if name != "" {
	// 	q = q.Where("name = ?", name)
	// }

	result := orm.Find(&accounts, "name = ?", name)

	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    accounts,
	})
}

type BodyPayloadAccount struct {
	// 	AccountID string
	// 	Name      string
	// 	Address   string
}

func (a *accountImplement) CreateAccount(g *gin.Context) {
	bodyPayload := models.Account{}

	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
	}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Create(&bodyPayload)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    bodyPayload,
	})
}

func (a *accountImplement) UpdateAccount(g *gin.Context) {

	bodyPayload := models.Account{}
	err := g.BindJSON(&bodyPayload)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	user := models.Account{}

	// queryParam := g.Request.URL.Query()

	// name := queryParam.Get("name")
	orm.First(&user, "account_id = ?", id)

	if user.AccountID == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Account not found",
		})
		return
	}

	user.Name = bodyPayload.Name
	user.Username = bodyPayload.Username
	orm.Save(user)

	g.JSON(http.StatusOK, gin.H{
		"message": "Update account successfully",
		"data":    user,
	})
}

func (a *accountImplement) DeleteAccount(g *gin.Context) {
	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Where("account_id = ? ", id).Delete(&models.Account{})

	fmt.Println(id, "check")
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Delete account successfully",
		"data":    id,
	})
}

type BodyPayloadBalance struct {
	Account_ID string
	Month      string
}

func (a *accountImplement) BalanceAccount(g *gin.Context) {
	BodyPayloadBalance := BodyPayloadBalance{}

	err := g.BindJSON(&BodyPayloadBalance)
	if err != nil {
		g.AbortWithError(http.StatusBadRequest, err)

	}
	transactionResult := struct {
		Total int
	}{}
	transaction := []models.Transaction{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	q := orm
	result := q.Find(&transaction)
	orm.Model(&models.Transaction{}).Select("sum(amount) as total").Where("account_id = ? AND date_part( 'Month' , transaction_date) = ?", BodyPayloadBalance.Account_ID, BodyPayloadBalance.Month).Group("account_id").Scan(&transactionResult)

	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Hello guys this API rest for later",
		"data":    transaction,
	})
}
