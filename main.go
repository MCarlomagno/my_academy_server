package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	//_ "github.com/lib/pq"
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

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	///
	/// START CLASSES
	///
	router.GET("/classes", func(c *gin.Context) {
		c.String(http.StatusOK, "Get a classes!")
	})
	router.POST("/classes", func(c *gin.Context) {
		c.String(http.StatusOK, "post a classes!")
	})
	router.PUT("/classes", func(c *gin.Context) {
		c.String(http.StatusOK, "put a classes!")
	})
	router.DELETE("/classes", func(c *gin.Context) {
		c.String(http.StatusOK, "delete a classes!")
	})
	///
	/// FINISH CLASSES
	///

	///
	/// START USERS
	///
	router.GET("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "Get a users!")
	})
	router.POST("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "post a users!")
	})
	router.PUT("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "put a users!")
	})
	router.DELETE("/users", func(c *gin.Context) {
		c.String(http.StatusOK, "delete a users!")
	})
	///
	/// FINISH USERS
	///

	///
	/// START MODULES
	///
	router.GET("/modules", func(c *gin.Context) {
		c.String(http.StatusOK, "Get a modules!")
	})
	router.POST("/modules", func(c *gin.Context) {
		c.String(http.StatusOK, "post a modules!")
	})
	router.PUT("/modules", func(c *gin.Context) {
		c.String(http.StatusOK, "put a modules!")
	})
	router.DELETE("/modules", func(c *gin.Context) {
		c.String(http.StatusOK, "delete a modules!")
	})
	///
	/// FINISH MODULES
	///

	///
	/// START COURSES
	///
	router.GET("/courses", func(c *gin.Context) {
		c.String(http.StatusOK, "Get a courses!")

		fmt.Println("1")
		//we open the database
		db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

		fmt.Println("2")
		err = db.Ping()
		if err != nil {
			panic(err)
		}

		//creating the statement
		sqlStatement := "SELECT id, title FROM courses WHERE id=%d;"

		fmt.Println("3")
		// creating variables to take result
		var title string
		var id int

		fmt.Println("4")
		//Querying
		row := db.QueryRow(sqlStatement, 1)

		// scaning result
		switch err := row.Scan(&id, &title); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
		case nil:
			fmt.Println(id, title)
		default:
			panic(err)
		}

		fmt.Println("5")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("6")
		var msg struct {
			Ttile string
			Id    int
		}

		fmt.Println("7")
		msg.Ttile = title
		msg.Id = id

		fmt.Println("8")
		// response
		c.JSON(http.StatusOK, msg)

		//closing connection
		defer db.Close()
	})
	router.POST("/courses", func(c *gin.Context) {
		c.String(http.StatusOK, "post a courses!")
	})
	router.PUT("/courses", func(c *gin.Context) {
		c.String(http.StatusOK, "put a courses!")
	})
	router.DELETE("/courses", func(c *gin.Context) {
		c.String(http.StatusOK, "delete a courses!")
	})
	///
	/// FINISH COURSES
	///

	router.Run(":" + port)
}
