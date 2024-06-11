package handlers

import (
	"github.com/gin-gonic/gin"
	"my_project/model"
	"my_project/storage/postgres"
	"net/http"
)

func ConnectionWithInterviewHandler() {
	db, err := postgres.ConnectionDB()
	if err != nil {
		panic(err)
	}
	postgres.NewCompanyRepository(db)

}

type ConnectionWithInterview struct {
	interviewHandler *postgres.InterviewRepository
}

func NewConnectionWithInterview(repo *postgres.InterviewRepository) *ConnectionWithInterview {
	return &ConnectionWithInterview{interviewHandler: repo}
}

func (intrv *ConnectionWithInterview) CreateInterviewHandler(gn *gin.Context) {
	interview := model.Interview{}
	err := gn.BindJSON(&interview)
	if err != nil {
		gn.JSON(http.StatusBadRequest,
			gin.H{
				"message": err,
				"status":  http.StatusBadRequest,
			})
	}
	err = intrv.interviewHandler.CreateInterview(interview)
	if err != nil {
		gn.JSON(http.StatusInternalServerError,
			gin.H{
				"message": err,
				"status":  http.StatusInternalServerError,
			})
	}
	gn.JSON(http.StatusOK, gin.H{
		"message": "created is success",
		"status":  http.StatusOK,
	})
}
func (intrv *ConnectionWithInterview) UpdateInterviewHandler(gn *gin.Context) {
	interview := model.Interview{}
	err := gn.BindJSON(&interview)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
		})
	}

	err = intrv.interviewHandler.UpdateInterview(gn.Param("id"), interview)
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
func (intrv *ConnectionWithInterview) DeletedInterviewHandler(gn *gin.Context) {
	err := intrv.interviewHandler.DeletedInterview(gn.Param("id"))
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
func (intrv *ConnectionWithInterview) GetAllInterviewHandler(gn *gin.Context) {
	interviews := []model.Interview{}
	interviews, err := intrv.interviewHandler.GetAllInterview()
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
		})
	}
	for i := 0; i < len(interviews); i++ {
		gn.JSON(http.StatusOK, gin.H{
			"id":             interviews[i].Id,
			"user_id":        interviews[i].UserId,
			"vacancy_id":     interviews[i].VacancyId,
			"recruiter_id":   interviews[i].RecruiterId,
			"interview_date": interviews[i].InterviewDate,
			"created_at":     interviews[i].CreatedAt,
			"updated_at":     interviews[i].UpdatedAt,
			"deleted_at":     interviews[i].DeletedAt,
		})
	}

}
