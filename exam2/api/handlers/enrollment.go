package handlers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"my_project/model"
	"my_project/strorage/postgres"
	"net/http"
	"time"
)

type HandlerEnrollmentConnection struct {
	enrollmentHandler *postgres.EnrollmentRepository
}

func NewConnectionWithEnrollment(repo *postgres.EnrollmentRepository) *HandlerEnrollmentConnection {
	return &HandlerEnrollmentConnection{enrollmentHandler: repo}
}

func (handle *HandlerEnrollmentConnection) CreateEnrollmentHandler(gn *gin.Context) {
	enrollment := model.Enrollment{}
	err := gn.BindJSON(&enrollment)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.enrollmentHandler.CreateEnrollment(enrollment)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		gn.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"status":  http.StatusOK,
			"time":    time.Now(),
		})
	}
}
func (handle *HandlerEnrollmentConnection) UpdateEnrollmentHandler(gn *gin.Context) {
	enrollment := model.Enrollment{}
	err := gn.BindJSON(&enrollment)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.enrollmentHandler.UpdateEnrollment(gn.Param("id"), enrollment)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		gn.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"status":  http.StatusOK,
			"time":    time.Now(),
		})
	}
}

func (handle *HandlerEnrollmentConnection) DeleteEnrollmentHandler(gn *gin.Context) {
	err := handle.enrollmentHandler.DeleteEnrollment(gn.Param("id"))
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		gn.JSON(http.StatusOK, gin.H{
			"message": "Success",
			"status":  http.StatusOK,
			"time":    time.Now(),
		})
	}
}
func (handle *HandlerEnrollmentConnection) GetAllEnrollmentHandler(gn *gin.Context) {
	enrollments, err := handle.enrollmentHandler.GetAllEnrollments()
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(enrollments); i++ {

			gn.JSON(http.StatusOK, gin.H{
				"message":         "Success",
				"status":          http.StatusOK,
				"time":            time.Now(),
				"id":              enrollments[i].Id,
				"user_id":         enrollments[i].UserId,
				"course_id":       enrollments[i].CourseId,
				"enrollment_date": enrollments[i].EnrollmentDate,
				"created_at":      enrollments[i].CreatedAt,
				"updated_at":      enrollments[i].UpdatedAt,
				"deleted_at":      enrollments[i].DeletedAt,
			})
		}
	}
}

func (handle *HandlerEnrollmentConnection) GetAEnrollmentFilterHandler(gn *gin.Context) {
	enrollmentFilter := model.EnrollmentFilter{}
	enrollmentFilter.User_id = gn.Query("id")
	enrollmentFilter.User_id = gn.Query("user_id")
	enrollmentFilter.Course_id = gn.Query("course_id")
	enrollmentFilter.EnrollmentDate = gn.Query("enrollment_date")
	enrollments, err := handle.enrollmentHandler.GetEnrollmentFilter(enrollmentFilter)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(enrollments); i++ {

			gn.JSON(http.StatusOK, gin.H{
				"message":         "Success",
				"status":          http.StatusOK,
				"time":            time.Now(),
				"id":              enrollments[i].Id,
				"user_id":         enrollments[i].UserId,
				"course_id":       enrollments[i].CourseId,
				"enrollment_date": enrollments[i].EnrollmentDate,
				"created_at":      enrollments[i].CreatedAt,
				"updated_at":      enrollments[i].UpdatedAt,
				"deleted_at":      enrollments[i].DeletedAt,
			})
		}
	}
}
