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

type HandlerCourseConnection struct {
	courseHandler *postgres.CourseRepository
}

func NewConnectionWithCourse(repo *postgres.CourseRepository) *HandlerCourseConnection {
	return &HandlerCourseConnection{courseHandler: repo}
}

func (handle *HandlerCourseConnection) CreateCourseHandler(gn *gin.Context) {
	course := model.Course{}
	err := gn.BindJSON(&course)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.courseHandler.CreateCourse(course)
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
func (handle *HandlerCourseConnection) UpdateCourseHandler(gn *gin.Context) {
	course := model.Course{}
	err := gn.BindJSON(&course)
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,
			"status":  http.StatusBadRequest,
			"time":    time.Now(),
		})
	}
	err = handle.courseHandler.UpdateCourse(gn.Param("id"), course)
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

func (handle *HandlerCourseConnection) DeleteCourseHandler(gn *gin.Context) {
	fmt.Println(gn.Param("id"))
	err := handle.courseHandler.DeleteCourse(gn.Param("id"))
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
func (handle *HandlerCourseConnection) GetAllCourseHandler(gn *gin.Context) {
	courses, err := handle.courseHandler.GetAllCourses()
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(courses); i++ {

			gn.JSON(http.StatusOK, gin.H{
				"message":     "Success",
				"status":      http.StatusOK,
				"time":        time.Now(),
				"id":          courses[i].Id,
				"title":       courses[i].Title,
				"description": courses[i].Description,
				"created_at":  courses[i].CreatedAt,
				"updated_at":  courses[i].UpdatedAt,
				"deleted_at":  courses[i].DeletedAt,
			})
		}
	}
}

func (handle *HandlerCourseConnection) GetCourseFilterHandler(gn *gin.Context) {
	courseFilter := model.CourseFilter{}
	courseFilter.Id = gn.Query("id")
	courseFilter.Title = gn.Query("title")
	courseFilter.Description = gn.Query("description")
	courses, err := handle.courseHandler.GetCourseFilter(courseFilter)
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(courses); i++ {

			gn.JSON(http.StatusOK, gin.H{
				"message":     "Success",
				"status":      http.StatusOK,
				"time":        time.Now(),
				"id":          courses[i].Id,
				"title":       courses[i].Title,
				"description": courses[i].Description,
				"created_at":  courses[i].CreatedAt,
				"updated_at":  courses[i].UpdatedAt,
				"deleted_at":  courses[i].DeletedAt,
			})
		}
	}
}

func (handle *HandlerCourseConnection) GetLessonsByCourseIdCourseHandler(gn *gin.Context) {
	courses, err := handle.courseHandler.GetLessonByCourseId(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(courses); i++ {
			gn.JSON(200, gin.H{
				"message": "Success",
				"status":  http.StatusOK,
				"time":    time.Now(),
				"id":      courses[i].Id,
				"title":   courses[i].Title,
				"Lessons": gin.H{
					"lesson_id":  courses[i].Lessons.Id,
					"course_id":  courses[i].Lessons.CourseId,
					"title":      courses[i].Lessons.Title,
					"content":    courses[i].Lessons.Content,
					"created_at": courses[i].Lessons.CreatedAt,
					"updated_at": courses[i].Lessons.UpdatedAt,
					"deleted_at": courses[i].Lessons.DeletedAt,
				},
				"description": courses[i].Description,
				"created_at":  courses[i].CreatedAt,
				"updated_at":  courses[i].UpdatedAt,
				"deleted_at":  courses[i].DeletedAt,
			})
		}
	}
}

func (handle *HandlerCourseConnection) GetGetEnrollmentByCourseIdHandler(gn *gin.Context) {
	courses, err := handle.courseHandler.GetEnrollmentByCourseId(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {

		for i := 0; i < len(courses); i++ {
			gn.JSON(http.StatusOK, gin.H{
				"message":     "Success",
				"status":      http.StatusOK,
				"time":        time.Now(),
				"id":          courses[i].Id,
				"title":       courses[i].Title,
				"description": courses[i].Description,
				"User": gin.H{
					"message":    "Success",
					"status":     http.StatusOK,
					"id":         courses[i].Users.Id,
					"name":       courses[i].Users.Name,
					"email":      courses[i].Users.Email,
					"birthday":   courses[i].Users.Birthday,
					"created_at": courses[i].Users.CreatedAt,
					"updated_at": courses[i].Users.UpdatedAt,
					"deleted_at": courses[i].Users.DeletedAt,
				},
				"created_at": courses[i].CreatedAt,
				"updated_at": courses[i].UpdatedAt,
				"deleted_at": courses[i].DeletedAt,
			})
		}
	}
}
