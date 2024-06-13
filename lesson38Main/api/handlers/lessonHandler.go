package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_project/model"
	"net/http"
	"strconv"
	"time"
)

func (handle *Handler) CreateLesson(gn *gin.Context) {
	lesson := model.Lesson{}
	err := gn.BindJSON(&lesson)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.LessonRepo.CreateLesson(lesson)
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
func (handle *Handler) UpdateLesson(gn *gin.Context) {
	lesson := model.Lesson{}
	err := gn.BindJSON(&lesson)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.LessonRepo.UpdateLesson(gn.Param("id"), lesson)
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

func (handle *Handler) DeleteLesson(gn *gin.Context) {
	err := handle.LessonRepo.DeleteLesson(gn.Param("id"))
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

func (handle *Handler) GetLesson(gn *gin.Context) {
	lessonFilter := model.LessonFilter{}
	lessonFilter.CourseId = gn.Query("course_id")
	lessonFilter.Title = gn.Query("title")
	lessonFilter.Content = gn.Query("content")
	limit, err := strconv.Atoi(gn.Query("limit"))
	offset, err := strconv.Atoi(gn.Query("offset"))
	lessonFilter.Limit = limit
	lessonFilter.Offset = offset
	lessons, err := handle.LessonRepo.GetLessons(lessonFilter)
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(lessons); i++ {

			gn.JSON(http.StatusOK, gin.H{
				"message":    "Success",
				"status":     http.StatusOK,
				"time":       time.Now(),
				"id":         lessons[i].Id,
				"course_id":  lessons[i].CourseId,
				"title":      lessons[i].Title,
				"content":    lessons[i].Content,
				"created_at": lessons[i].CreatedAt,
				"updated_at": lessons[i].UpdatedAt,
				"deleted_at": lessons[i].DeletedAt,
			})
		}
	}
}

func (handle *Handler) GetLessonById(gn *gin.Context) {
	lesson, err := handle.LessonRepo.GetById(gn.Param("id"))
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
			"id":         lesson.Id,
			"content":    lesson.Content,
			"course_id":  lesson.CourseId,
			"title":      lesson.Title,
			"created_at": lesson.CreatedAt,
			"updated_at": lesson.UpdatedAt,
			"deleted_at": lesson.DeletedAt,
		})
	}

}
