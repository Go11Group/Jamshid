package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"my_project/model"
	"my_project/storage/postgres"
	"net/http"
)

func ConnectionWithCompanyHandler() {
	db, err := postgres.ConnectionDB()
	if err != nil {
		panic(err)
	}
	postgres.NewCompanyRepository(db)

}

type ConnectionWithCompany struct {
	companyHandler *postgres.CompanyRepository
}

func NewConnectionWithCompany(repo *postgres.CompanyRepository) *ConnectionWithCompany {
	return &ConnectionWithCompany{companyHandler: repo}
}

func (cmh *ConnectionWithCompany) CreateCompanyHandler(gn *gin.Context) {
	company := model.Company{}
	err := gn.BindJSON(&company)
	fmt.Println(company)
	if err != nil {
		fmt.Println("---------------------", err)
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
		})
	}
	err = cmh.companyHandler.CreateCompany(company)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	gn.JSON(http.StatusOK, gin.H{
		"message": "create is success",
		"status":  http.StatusOK,
	})
}

func (cmh *ConnectionWithCompany) UpdateCompanyHandler(gn *gin.Context) {
	company := model.Company{}
	err := gn.BindJSON(&company)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
		})
		fmt.Println("error blinjsonda", err)
	}

	err = cmh.companyHandler.UpdateCompany(gn.Param("id"), company)
	if err != nil {
		gn.JSON(http.StatusInternalServerError,
			gin.H{
				"message": err,
				"status":  http.StatusInternalServerError,
			})
	}
	gn.JSON(http.StatusOK, gin.H{
		"message": "update is success",
		"status":  http.StatusOK,
	})

}
func (cmh *ConnectionWithCompany) DeletedCompanyHandler(gn *gin.Context) {
	err := cmh.companyHandler.DeletedCompany(gn.Param("id"))
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	gn.JSON(http.StatusOK, gin.H{
		"message": "delete is success",
		"status":  http.StatusOK,
	})
}
func (cmh *ConnectionWithCompany) GetAllCompanyHandler(gn *gin.Context) {
	companies := []model.Company{}
	companies, err := cmh.companyHandler.GetAllCompany()
	if err != nil {
		gn.JSON(http.StatusInternalServerError,
			gin.H{
				"message": err,
				"status":  http.StatusInternalServerError,
			})
	}

	for i := 0; i < len(companies); i++ {
		gn.JSON(http.StatusOK,
			gin.H{
				"id":         companies[i].Id,
				"name":       companies[i].Name,
				"location":   companies[i].Location,
				"workers":    companies[i].Workers,
				"created_at": companies[i].CreatedAt,
				"updated_at": companies[i].UpdatedAt,
				"deleted_at": companies[i].DeletedAt,
			})
	}

}
