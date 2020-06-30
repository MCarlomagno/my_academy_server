package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/classes"
	"github.com/heroku/go-getting-started/courses"
	"github.com/heroku/go-getting-started/modules"
	"github.com/heroku/go-getting-started/users"
	_ "github.com/lib/pq"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	///
	/// START CLASSES
	///
	router.GET("/classes", classes.GetClassesRoot)
	router.GET("/classes/getClassesByModuleId/:moduleId", classes.GetClassesByModuleID)
	router.POST("/classes", classes.PostClassesRoot)
	router.PUT("/classes", classes.PutClassesRoot)
	router.DELETE("/classes", classes.DeleteClassesRoot)
	///
	/// FINISH CLASSES
	///

	///
	/// START USERS
	///
	router.GET("/users", users.GetUsersRoot)
	router.POST("/users", users.PostUsersRoot)
	router.PUT("/users", users.PutUsersRoot)
	router.DELETE("/users", users.DeleteUsersRoot)
	///
	/// FINISH USERS
	///

	///
	/// START MODULES
	///
	router.GET("/modules", modules.GetModulesRoot)
	router.GET("/modules/getModulesByCourseId/:courseId", modules.GetModulesByCourseID)
	router.POST("/modules", modules.PostModulesRoot)
	router.PUT("/modules", modules.PutModulesRoot)
	router.DELETE("/modules", modules.DeleteModulesRoot)
	///
	/// FINISH MODULES
	///

	///
	/// START COURSES
	///
	router.GET("/courses", courses.GetCoursesRoot)
	router.GET("/courses/getUserCreatedCourses/:ownerUserId", courses.GetUserCreatedCourses)
	router.POST("/courses", courses.PostCoursesRoot)
	router.PUT("/courses", courses.PutCoursesRoot)
	router.DELETE("/courses", courses.DeleteCoursesRoot)
	///
	/// FINISH COURSES
	///

	router.Run(":" + port)
}
