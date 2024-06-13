package controller

import (
	"github.com/gin-gonic/gin"
	"my_project/handlers"
	"my_project/strorage/postgres"
)

func Router(gn *gin.Engine) *gin.Engine {
	db, err := postgres.ConnectionDb()
	if err != nil {
		panic(err)
	}
	userRepo := postgres.NewUserRepository(db)
	routeUser := handlers.NewConnectionWithUser(userRepo)
	// User
	gn.POST("/api/user/create", routeUser.CreateUserHandler)
	gn.GET("/api/user/get", routeUser.GetAllUserHandler)
	gn.GET("/api/user/filter", routeUser.GetUserFilterHandler)
	gn.GET("/api/user/courses/:id", routeUser.GetGetCourseByUserIdHandler)
	gn.GET("/api/user/get/:name/:email", routeUser.GetAllUserByEmailOrNameHandler)
	gn.PUT("/api/user/update/:id", routeUser.UpdateUserHandler)
	gn.DELETE("/api/user/delete/:id", routeUser.DeleteUserHandler)
	courseRepo := postgres.NewCourseRepository(db)
	routeCourse := handlers.NewConnectionWithCourse(courseRepo)

	gn.POST("/api/course/create", routeCourse.CreateCourseHandler)
	gn.GET("/api/course/get", routeCourse.GetAllCourseHandler)
	gn.GET("/api/course/filter", routeCourse.GetCourseFilterHandler)
	gn.GET("/api/course/lessons/:id", routeCourse.GetLessonsByCourseIdCourseHandler)
	gn.GET("/api/course/enrollments/:id", routeCourse.GetGetEnrollmentByCourseIdHandler)
	gn.PUT("/api/course/update/:id", routeCourse.UpdateCourseHandler)
	gn.DELETE("/api/course/delete/:id", routeCourse.DeleteCourseHandler)

	lessonRepo := postgres.NewLessonRepository(db)
	routeLesson := handlers.NewConnectionWithLesson(lessonRepo)

	gn.POST("/api/lesson/create", routeLesson.CreateLessonHandler)
	gn.GET("/api/lesson/get", routeLesson.GetAllLessonHandler)
	gn.GET("/api/lesson/get/filter", routeLesson.GetLessonFilterHandler)
	gn.PUT("/api/lesson/update/:id", routeLesson.UpdateLessonHandler)
	gn.DELETE("/api/lesson/delete/:id", routeLesson.DeleteLessonHandler)

	enrollmentRepo := postgres.NewEnrollmentRepository(db)
	enrollmentLesson := handlers.NewConnectionWithEnrollment(enrollmentRepo)

	gn.POST("/api/enrollment/create", enrollmentLesson.CreateEnrollmentHandler)
	gn.GET("/api/enrollment/get", enrollmentLesson.GetAllEnrollmentHandler)
	gn.GET("/api/enrollment/get/filter", enrollmentLesson.GetAEnrollmentFilterHandler)
	gn.PUT("/api/enrollment/update/:id", enrollmentLesson.UpdateEnrollmentHandler)
	gn.DELETE("/api/enrollment/delete/:id", enrollmentLesson.DeleteEnrollmentHandler)
	return gn
}
