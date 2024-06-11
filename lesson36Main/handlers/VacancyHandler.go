package handlers

import (
	"github.com/gin-gonic/gin"
	"my_project/model"
	"my_project/storage/postgres"
	"net/http"
)

func ConnectionWithVacancyHandler() {
	db, err := postgres.ConnectionDB()
	if err != nil {
		panic(err)
	}
	postgres.NewVacancyRepository(db)

}

type ConnectionWithVacancy struct {
	vacancyHandler *postgres.VacancyRepository
}

func NewConnectionWithVacancy(repo *postgres.VacancyRepository) *ConnectionWithVacancy {
	return &ConnectionWithVacancy{vacancyHandler: repo}
}

func (vh *ConnectionWithVacancy) CreateVacancyHandler(gn *gin.Context) {
	vacancy := model.Vacancy{}
	err := gn.BindJSON(&vacancy)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadGateway,
		})
	}
	err = vh.vacancyHandler.CreateVacancy(vacancy)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	gn.JSON(http.StatusOK, gin.H{
		"message": "created is success",
		"status":  http.StatusOK,
	})
}
func (vh *ConnectionWithVacancy) UpdateVacancyHandler(gn *gin.Context) {
	vacancy := model.Vacancy{}
	err := gn.BindJSON(&vacancy)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
		})
	}

	err = vh.vacancyHandler.UpdateVacancy(gn.Param("id"), vacancy)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	gn.JSON(http.StatusOK, gin.H{
		"message": "updated is success",
		"status":  http.StatusOK,
	})

}
func (vh *ConnectionWithVacancy) DeletedVacancyHandler(gn *gin.Context) {
	err := vh.vacancyHandler.DeletedVacancy(gn.Param("id"))
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	gn.JSON(http.StatusOK, gin.H{
		"message": "deleted is success",
		"status":  http.StatusOK,
	})
}
func (vh *ConnectionWithVacancy) GetAllVacancyHandler(gn *gin.Context) {
	vacancies := []model.Vacancy{}
	vacancies, err := vh.vacancyHandler.GetAllVacancy()
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	for i := 0; i < len(vacancies); i++ {
		gn.JSON(http.StatusOK, gin.H{
			"id":             vacancies[i].Id,
			"name":           vacancies[i].Name,
			"min_experience": vacancies[i].MinExperience,
			"company_id":     vacancies[i].CompanyID,
			"description":    vacancies[i].Description,
			"created_at":     vacancies[i].CreatedAt,
			"updated_at":     vacancies[i].UpdatedAt,
			"deleted_at":     vacancies[i].DeletedAt,
		})
	}

}
