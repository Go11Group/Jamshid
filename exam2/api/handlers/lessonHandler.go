package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"my_project/model"
	"my_project/strorage/postgres"
	"net/http"
	"time"
)

type HandlerLessonConnection struct {
	lessonHandler *postgres.LessonRepository
}

func NewConnectionWithLesson(repo *postgres.LessonRepository) *HandlerLessonConnection {
	return &HandlerLessonConnection{lessonHandler: repo}
}

func (handle *HandlerLessonConnection) CreateLessonHandler(gn *gin.Context) {
	lesson := model.Lesson{}
	err := gn.BindJSON(&lesson)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.lessonHandler.CreateLesson(lesson)
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
func (handle *HandlerLessonConnection) UpdateLessonHandler(gn *gin.Context) {
	lesson := model.Lesson{}
	err := gn.BindJSON(&lesson)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.lessonHandler.UpdateLesson(gn.Param("id"), lesson)
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

func (handle *HandlerLessonConnection) DeleteLessonHandler(gn *gin.Context) {
	err := handle.lessonHandler.DeleteLesson(gn.Param("id"))
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
func (handle *HandlerLessonConnection) GetAllLessonHandler(gn *gin.Context) {
	lessons, err := handle.lessonHandler.GetAllLessons()
	if err != nil {
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

func (handle *HandlerLessonConnection) GetLessonFilterHandler(gn *gin.Context) {
	lessonFilter := model.LessonFilter{}
	lessonFilter.Id = gn.Query("id")
	lessonFilter.CourseId = gn.Query("course_id")
	lessonFilter.Title = gn.Query("title")
	lessonFilter.Content = gn.Query("content")
	lessons, err := handle.lessonHandler.GetLessonFilter(lessonFilter)
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
