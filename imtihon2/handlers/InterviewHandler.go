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
	companyHandler postgres.CompanyRepository
}

func (cmh *ConnectionWithCompany) CreateCompanyHandler(cmp *gin.Context) {
	company := model.Company{}
	err := cmp.BindJSON(&company)
	if err != nil {
		fmt.Println("error blinjsonda", err)
	}
	err = cmh.companyHandler.CreateCompany(company)
	if err != nil {
		fmt.Println("error postgresda company ", err)
	}
}

func (cmh *ConnectionWithCompany) UpdateCompanyHandler(cmp *gin.Context) {
	company := model.Company{}
	err := cmp.BindJSON(&company)
	if err != nil {
		fmt.Println("error blinjsonda", err)
	}

	err = cmh.companyHandler.UpdateCompany(cmp.Param("id"), company)
	if err != nil {
		fmt.Println("error postgresda company ", err)
	}

}
func (cmh *ConnectionWithCompany) DeletedCompanyHandler(cmp *gin.Context) {
	err := cmh.companyHandler.DeletedCompany(cmp.Param("id"))
	if err != nil {
		fmt.Println("error postgresda company ", err)
	}
}
func (cmh *ConnectionWithCompany) GetAllCompanyHandler(cmp *gin.Context) {
	companies := []model.Company{}
	companies, err := cmh.companyHandler.GetAllCompany()
	if err != nil {
		fmt.Println("error postgresda", err)
	}
	for i := 0; i < len(companies); i++ {
		cmp.JSON(http.StatusOK,
			gin.H{
			  "id":companies[i].Id,
			  "name":companies[i].Name,
			  "location":companies[i].Location,
			  "workers":companies[i].Workers,
			  "created_at":companies[i].CreatedAt,
			  "updated_at":companies[i].UpdatedAt,
			  "deleted_at":companies[i].DeletedAt,
			})
	}


}
