package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"my_project/model"
	"net/http"
	"strconv"
	"time"
)

/* create course method dida gin contextni oladi  bu yerda request ni qabul qilib create methoddga berib yuboradi u method table yozadi*/
func (handle *Handler) CreateCourse(gn *gin.Context) {
	course := model.Course{}    // yangi course struct ochish
	err := gn.BindJSON(&course) // requesdan kelayotgan malumot ni course ga parse qilish
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,                   // error qayatadi response
			"status":  http.StatusBadRequest, // error qayatadi status
			"time":    time.Now(),            // error qayatadi qaytgan vaqti
		})
	}
	// create course bu postgres ichidagi method borgan malumot ni tablega yozadi
	err = handle.CourseRepo.CreateCourse(course)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,                            // error qayatadi response
			"status":  http.StatusInternalServerError, // internalserver error  qayatadi status
			"time":    time.Now(),                     // error qayatadi qaytgan vaqti
		})
	} else {
		gn.JSON(http.StatusOK, gin.H{
			"message": "Success",     // success qayatadi response
			"status":  http.StatusOK, // 200 qayatadi status
			"time":    time.Now(),    //  qayatadi qaytgan vaqti
		})
	}
}
func (handle *Handler) UpdateCourse(gn *gin.Context) {
	course := model.Course{}    //// yangi course struct ochish
	err := gn.BindJSON(&course) // requesdan kelayotgan malumot ni course ga parse qilish
	if err != nil {
		gn.JSON(http.StatusBadRequest, gin.H{
			"message": err,                   // error qayatadi responsega
			"status":  http.StatusBadRequest, // badrequest error  qayatadi status
			"time":    time.Now(),            // qayatadi qaytgan vaqti
		})
	}
	// bu yerda param da olib method ga berib yuboriladi paramda id  keladi
	err = handle.CourseRepo.UpdateCourse(gn.Param("id"), course)
	if err != nil {
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,                            // error response
			"status":  http.StatusInternalServerError, // status 5000
			"time":    time.Now(),                     // time hozirgi vaqt
		})
	} else {
		gn.JSON(http.StatusOK, gin.H{
			"message": "Success",     /// success response
			"status":  http.StatusOK, // status 200
			"time":    time.Now(),    //time now
		})
	}
}

func (handle *Handler) DeleteCourse(gn *gin.Context) {
	fmt.Println(gn.Param("id"))
	err := handle.CourseRepo.DeleteCourse(gn.Param("id"))
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

func (handle *Handler) GetCourse(gn *gin.Context) {
	courseFilter := model.CourseFilter{}
	courseFilter.Title = gn.Query("title")
	limit, err := strconv.Atoi(gn.Query("limit"))
	offset, err := strconv.Atoi(gn.Query("offset"))
	courseFilter.Limit = limit
	courseFilter.Offset = offset
	courseFilter.Description = gn.Query("description")
	courses, err := handle.CourseRepo.GetCourses(courseFilter)
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

func (handle *Handler) GetLessonsByCourseId(gn *gin.Context) {
	courseId, lesssons, err := handle.CourseRepo.GetLessonByCourseId(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {
		for i := 0; i < len(lesssons); i++ {
			gn.JSON(200, gin.H{
				"message": "Success",
				"status":  http.StatusOK,
				"time":    time.Now(),
				"id":      courseId,
				"Lessons": gin.H{
					"lesson_id": lesssons[i].Id,
					"course_id": lesssons[i].CourseId,
					"title":     lesssons[i].Title,
					"content":   lesssons[i].Content,
				},
			})
		}
	}
}

func (handle *Handler) GetUserWithEnrollmentByCourseId(gn *gin.Context) {
	courseId, users, err := handle.CourseRepo.GetUserByCourseIdWithEnrollment(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {

		for i := 0; i < len(users); i++ {
			gn.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"status":  http.StatusOK,
				"time":    time.Now(),
				"id":      courseId,
				"User": gin.H{
					"message": "Success",
					"status":  http.StatusOK,
					"id":      users[i].Id,
					"name":    users[i].Name,
					"email":   users[i].Email,
				},
			})
		}
	}
}

func (handle *Handler) GetCourseById(gn *gin.Context) {
	course, err := handle.CourseRepo.GetById(gn.Param("id"))
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
	} else {

		gn.JSON(http.StatusOK, gin.H{
			"message":     "Success",
			"status":      http.StatusOK,
			"time":        time.Now(),
			"id":          course.Id,
			"title":       course.Title,
			"description": course.Description,
			"created_at":  course.CreatedAt,
			"updated_at":  course.UpdatedAt,
			"deleted_at":  course.DeletedAt,
		})
	}

}

func (handle *Handler) ShowPopularCourse(gn *gin.Context) {
	// Default values for start and end time
	startTime := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endTime := time.Date(2024, 12, 31, 23, 59, 59, 999, time.UTC)

	// Query parameters
	time1 := gn.Query("start_time")
	time2 := gn.Query("end_time")

	// Parse start_time if provided
	if time1 != "" {
		parsedStartTime, err := time.Parse("2006-01-02", time1)
		if err == nil {
			startTime = parsedStartTime
		} else {
			gn.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid start_time format. Use YYYY-MM-DD.",
				"status":  http.StatusBadRequest,
			})
			return
		}
	}

	// Parse end_time if provided
	if time2 != "" {
		parsedEndTime, err := time.Parse("2006-01-02", time2)
		if err == nil {
			endTime = parsedEndTime
		} else {
			gn.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid end_time format. Use YYYY-MM-DD.",
				"status":  http.StatusBadRequest,
			})
			return
		}
	}

	// Get the most popular course
	count, course, err := handle.CourseRepo.GetPopularyCourse(startTime, endTime)
	if err != nil {
		fmt.Println("+++++++++++", err)
		gn.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"status":  http.StatusInternalServerError,
			"time":    time.Now(),
		})
		return
	}

	gn.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"status":  http.StatusOK,
		"time":    time.Now(),
		"id":      course.Id,
		"title":   course.Title,
		"count":   count,
	})
}
