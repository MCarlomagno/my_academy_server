package enrollments

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetEnrollmentsRoot function
func GetEnrollmentsRoot(c *gin.Context) {
	c.String(http.StatusOK, "Get a enrollments!")
}

// PostEnrollmentsRoot function
func PostEnrollmentsRoot(c *gin.Context) {
	//we open the database
	db, errA := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errA != nil {
		fmt.Println("error opening db")
	}

	//decoding body
	var enrollment Enrollment

	errB := c.BindJSON(&enrollment)
	if errB != nil {
		fmt.Println("error binding body")
	}

	//creating the statement
	sqlStatement := `INSERT INTO enrollment (id_user, id_course)
		VALUES ($1, $2) RETURNING id;`

	// TODO autoincremental
	var newID = 0

	//Querying
	errC := db.QueryRow(sqlStatement, enrollment.UserID, enrollment.CourseID).Scan(&newID)
	if errC != nil {
		fmt.Println(errC.Error())
	}

	enrollment.ID = newID

	//closing connection
	defer db.Close()
	c.JSON(http.StatusOK, enrollment)
}

// PutEnrollmentsRoot function
func PutEnrollmentsRoot(c *gin.Context) {
	c.String(http.StatusOK, "put a enollments!")
}

// DeleteEnrollmentsRoot function
func DeleteEnrollmentsRoot(c *gin.Context) {
	c.String(http.StatusOK, "delete a enollments!")
}

// Enrollment model
type Enrollment struct {
	ID       int `json:"id"`
	UserID   int `json:"userId"`
	CourseID int `json:"courseId"`
}
