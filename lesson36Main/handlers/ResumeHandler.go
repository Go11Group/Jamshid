package handlers

import (
	"github.com/gin-gonic/gin"
	"my_project/model"
	"my_project/storage/postgres"
	"net/http"
)

func ConnectionWithResumeHandler() {
	db, err := postgres.ConnectionDB()
	if err != nil {
		panic(err)
	}
	postgres.NewResumeRepository(db)

}

type ConnectionWithResume struct {
	resumeHandler *postgres.ResumeRepository
}

func NewConnectionWithResume(repo *postgres.ResumeRepository) *ConnectionWithResume {
	return &ConnectionWithResume{resumeHandler: repo}
}

func (reh *ConnectionWithResume) CreateResumeHandler(gn *gin.Context) {
	resume := model.Resume{}
	err := gn.BindJSON(&resume)
	gn.JSON(http.StatusBadRequest, gin.H{
		"message": err,
		"status":  http.StatusBadRequest,
	})
	err = reh.resumeHandler.CreateResume(resume)
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
func (reh *ConnectionWithResume) UpdateResumeHandler(gn *gin.Context) {
	resume := model.Resume{}
	err := gn.BindJSON(&resume)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
		})
	}

	err = reh.resumeHandler.UpdateResume(gn.Param("id"), resume)
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
func (reh *ConnectionWithResume) DeletedResumeHandler(gn *gin.Context) {
	err := reh.resumeHandler.DeletedResume(gn.Param("id"))
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
func (reh *ConnectionWithResume) GetAllResumeHandler(gn *gin.Context) {
	resumes := []model.Resume{}
	resumes, err := reh.resumeHandler.GetAllResume()
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	for i := 0; i < len(resumes); i++ {
		gn.JSON(http.StatusOK, gin.H{
			"id":          resumes[i].Id,
			"position":    resumes[i].Position,
			"experience":  resumes[i].Experience,
			"description": resumes[i].Description,
			"user_id":     resumes[i].UserID,
			"created_at":  resumes[i].CreatedAt,
			"updated_at":  resumes[i].UpdatedAt,
			"deleted_at":  resumes[i].DeletedAt,
		})
	}

}
