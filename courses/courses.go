package courses

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetCoursesRoot function
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
		ID    int
	}

	msg.Title = title
	msg.ID = id

	// response
	c.JSON(http.StatusOK, msg)

}

// PostCoursesRoot function
func PostCoursesRoot(c *gin.Context) {
	//we open the database
	db, errA := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errA != nil {
		fmt.Println("error opening db")
	}

	//decoding body
	var bodyCourse Course

	errB := c.BindJSON(&bodyCourse)
	if errB != nil {
		fmt.Println("error binding body")
	}

	//creating the statement
	sqlStatement := `INSERT INTO courses (id_user, title, description)
		VALUES ($1, $2, $3) RETURNING id;`

	// TODO autoincremental
	var newID = 0

	//Querying
	errC := db.QueryRow(sqlStatement, bodyCourse.OwnerUserID, bodyCourse.Title, bodyCourse.Description).Scan(&newID)
	if errC != nil {
		fmt.Println(errC.Error())
	}

	bodyCourse.ID = newID

	//closing connection
	defer db.Close()
	c.JSON(http.StatusOK, bodyCourse)
}

// PutCoursesRoot function
func PutCoursesRoot(c *gin.Context) {
	c.String(http.StatusOK, "put a courses!")
}

// DeleteCoursesRoot function
func DeleteCoursesRoot(c *gin.Context) {
	c.String(http.StatusOK, "delete a courses!")
}

// Course model
type Course struct {
	ID          int    `json:"id"`
	OwnerUserID int    `json:"ownerUserId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
