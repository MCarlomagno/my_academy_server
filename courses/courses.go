package courses

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetCoursesRoot(c *gin.Context) {
	c.String(http.StatusOK, "Get a courses!")

	//we open the database
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//creating the statement
	sqlStatement := `SELECT id, title FROM courses WHERE id=$1;`

	// creating variables to take result
	var title string
	var id int

	//Querying
	row := db.QueryRow(sqlStatement, 1)

	//closing connection
	defer db.Close()

	// scaning result
	switch err := row.Scan(&id, &title); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, title)
	default:
		panic(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	var msg struct {
		Title string
		Id    int
	}

	msg.Title = title
	msg.Id = id

	// response
	c.JSON(http.StatusOK, msg)

}

func PostCoursesRoot(c *gin.Context) {
	c.String(http.StatusOK, "Post a courses!")
}

func PutCoursesRoot(c *gin.Context) {
	c.String(http.StatusOK, "put a courses!")
}

func DeleteCoursesRoot(c *gin.Context) {
	c.String(http.StatusOK, "delete a courses!")
}
