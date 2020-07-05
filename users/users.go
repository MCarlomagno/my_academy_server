package users

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetUsersRoot function
func GetUsersRoot(c *gin.Context) {
	//we open the database
	db, errA := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errA != nil {
		fmt.Println("error opening db")
	}
	//closing connection
	defer db.Close()

	//creating the statement
	sqlStatement := `SELECT id, email, name, surname, image_url FROM users WHERE id=$1;`

	//Querying
	row := db.QueryRow(sqlStatement, c.Param("userId"))
	defer db.Close()

	user := new(User)

	var id int
	var email string
	var name string
	var surname string
	var imageURL string

	errC := row.Scan(&id, &email, &name, &surname, &imageURL)
	if errC != nil {
		fmt.Println(errC)
	}

	user.ID = id
	user.Email = email
	user.Name = name
	user.Surname = surname
	user.ImageURL = imageURL

	c.JSON(http.StatusOK, user)
}

// PostUsersRoot function
func PostUsersRoot(c *gin.Context) {
	c.String(http.StatusOK, "Post a users!")
}

// PutUsersRoot function
func PutUsersRoot(c *gin.Context) {
	c.String(http.StatusOK, "put a users!")
}

// DeleteUsersRoot function
func DeleteUsersRoot(c *gin.Context) {
	c.String(http.StatusOK, "delete a users!")
}

// User model
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	ImageURL string `json:"imageUrl"`
}
