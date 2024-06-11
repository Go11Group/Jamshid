package handlers

import (
	"github.com/gin-gonic/gin"
	"my_project/model"
	"my_project/storage/postgres"
	"net/http"
)

func ConnectionWithRecruiterHandler() {
	db, err := postgres.ConnectionDB()
	if err != nil {
		panic(err)
	}
	postgres.NewRecruiterRepository(db)

}

type ConnectionWithRecruiter struct {
	recruiterHandler *postgres.RecruiterRepository
}

func NewConnectionWithRecruiter(repo *postgres.RecruiterRepository) *ConnectionWithRecruiter {
	return &ConnectionWithRecruiter{recruiterHandler: repo}
}

func (rch *ConnectionWithRecruiter) CreateRecruiterHandler(gn *gin.Context) {
	recruiter := model.Recruiter{}
	err := gn.BindJSON(&recruiter)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
		})
	}
	err = rch.recruiterHandler.CreateRecruiter(recruiter)
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
func (rch *ConnectionWithRecruiter) UpdateRecruiterHandler(gn *gin.Context) {
	recruiter := model.Recruiter{}
	err := gn.BindJSON(&recruiter)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}

	err = rch.recruiterHandler.UpdateRecruiter(gn.Param("id"), recruiter)
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
func (rch *ConnectionWithRecruiter) DeletedRecruiterHandler(gn *gin.Context) {
	err := rch.recruiterHandler.DeletedRecruiter(gn.Param("id"))
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
func (rch *ConnectionWithRecruiter) GetAllRecruiterHandler(gn *gin.Context) {
	recruiters := []model.Recruiter{}
	recruiters, err := rch.recruiterHandler.GetAllRecruiter()
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	for i := 0; i < len(recruiters); i++ {
		gn.JSON(http.StatusOK, gin.H{
			"id":           recruiters[i].Id,
			"name":         recruiters[i].Name,
			"email":        recruiters[i].Email,
			"phone_number": recruiters[i].PhoneNumber,
			"birthday":     recruiters[i].Birthday,
			"gender":       recruiters[i].Gender,
			"created_at":   recruiters[i].CreatedAt,
			"updated_at":   recruiters[i].UpdatedAt,
			"deleted_at":   recruiters[i].DeletedAt,
		})
	}

}
