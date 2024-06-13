package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/model"
	"net/http"
	"strconv"
	"time"
)

func (handle *Handler) CreateEnrollment(gn *gin.Context) {
	enrollment := model.Enrollment{}
	err := gn.BindJSON(&enrollment)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.EnrollmentRepo.CreateEnrollment(enrollment)
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
func (handle *Handler) UpdateEnrollment(gn *gin.Context) {
	enrollment := model.Enrollment{}
	err := gn.BindJSON(&enrollment)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.EnrollmentRepo.UpdateEnrollment(gn.Param("id"), enrollment)
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

func (handle *Handler) DeleteEnrollment(gn *gin.Context) {
	err := handle.EnrollmentRepo.DeleteEnrollment(gn.Param("id"))
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

func (handle *Handler) GetEnrollment(gn *gin.Context) {
	enrollmentFilter := model.EnrollmentFilter{}
	enrollmentFilter.User_id = gn.Query("id")
	enrollmentFilter.User_id = gn.Query("user_id")
	enrollmentFilter.Course_id = gn.Query("course_id")
	enrollmentFilter.EnrollmentDate = gn.Query("enrollment_date")
	limit, err := strconv.Atoi(gn.Query("limit"))
	offset, err := strconv.Atoi(gn.Query("offset"))
	enrollmentFilter.Limit = limit
	enrollmentFilter.Offset = offset
	enrollments, err := handle.EnrollmentRepo.GetEnrollments(enrollmentFilter)
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

func (handle *Handler) GetEnrollmentById(gn *gin.Context) {
	enrollment, err := handle.EnrollmentRepo.GetById(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message":    "Success",
			"status":     http.StatusOK,
			"time":       time.Now(),
			"id":         enrollment.Id,
			"user_id":    enrollment.UserId,
			"course_id":  enrollment.CourseId,
			"enrollment": enrollment.EnrollmentDate,
			"created_at": enrollment.CreatedAt,
			"updated_at": enrollment.UpdatedAt,
			"deleted_at": enrollment.DeletedAt,
		})
	}

}
