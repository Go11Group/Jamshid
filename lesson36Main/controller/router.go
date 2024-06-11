package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"my_project/handlers"
	"my_project/storage/postgres"
)

func RouterController(router *gin.Engine) *gin.Engine {
	db, err := postgres.ConnectionDB()
	if err != nil {
		fmt.Println("===================", err)
		panic(err)
	}
	companyRepo := postgres.NewCompanyRepository(db)
	companyHandler := handlers.NewConnectionWithCompany(companyRepo)
	router.POST("/api/company/create", companyHandler.CreateCompanyHandler)
	router.GET("/api/company/get", companyHandler.GetAllCompanyHandler)
	router.DELETE("/api/company/delete/:id", companyHandler.DeletedCompanyHandler)
	router.PUT("/api/company/update/:id", companyHandler.UpdateCompanyHandler)

	interviewRepo := postgres.NewInterviewRepository(db)
	interviewHandler := handlers.NewConnectionWithInterview(interviewRepo)

	router.POST("/api/interview/create", interviewHandler.CreateInterviewHandler)
	router.GET("/api/interview/get", interviewHandler.GetAllInterviewHandler)
	router.DELETE("/api/interview/delete/:id", interviewHandler.DeletedInterviewHandler)
	router.PUT("/api/interview/update/:id", interviewHandler.UpdateInterviewHandler)

	recruiterRepo := postgres.NewRecruiterRepository(db)
	recruiterHandler := handlers.NewConnectionWithRecruiter(recruiterRepo)

	router.POST("/api/recruiter/create", recruiterHandler.CreateRecruiterHandler)
	router.GET("/api/recruiter/get", recruiterHandler.GetAllRecruiterHandler)
	router.DELETE("/api/recruiter/delete/:id", recruiterHandler.DeletedRecruiterHandler)
	router.PUT("/api/recruiter/update/:id", recruiterHandler.UpdateRecruiterHandler)

	resumeRepo := postgres.NewResumeRepository(db)
	resumeHandler := handlers.NewConnectionWithResume(resumeRepo)
	router.POST("/api/resume/create", resumeHandler.CreateResumeHandler)
	router.GET("/api/resume/get", resumeHandler.GetAllResumeHandler)
	router.DELETE("/api/resume/delete/:id", resumeHandler.DeletedResumeHandler)
	router.PUT("/api/resume/update/:id", resumeHandler.UpdateResumeHandler)

	userRepo := postgres.NewUserRepository(db)
	userHandler := handlers.NewConnectionWithUser(userRepo)
	router.POST("/api/user/create", userHandler.CreateUserHandler)
	router.GET("/api/user/get", userHandler.GetAllUserHandler)
	router.DELETE("/api/user/delete/:id", userHandler.DeletedUserHandler)
	router.PUT("/api/user/update/:id", userHandler.UpdateUserHandler)

	vacancyRepo := postgres.NewVacancyRepository(db)
	vacancyHandler := handlers.NewConnectionWithVacancy(vacancyRepo)
	router.POST("/api/vacancy/create", vacancyHandler.CreateVacancyHandler)
	router.GET("/api/vacancy/get", vacancyHandler.GetAllVacancyHandler)
	router.POST("/api/vacancy/delete/:id", vacancyHandler.DeletedVacancyHandler)
	router.GET("/api/vacancy/update/:id", vacancyHandler.UpdateVacancyHandler)

	return router

}
