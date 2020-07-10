package users

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"

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

// Login function
func Login(c *gin.Context) {
	//we open the database
	db, errA := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errA != nil {
		fmt.Println("error opening db")
	}
	//closing connection
	defer db.Close()

	//select user by email
	emailSQLStatement := `SELECT id, email, name, surname, image_url, password FROM users WHERE email = $1;`

	//decoding body
	var bodyUser User

	errB := c.BindJSON(&bodyUser)
	if errB != nil {
		fmt.Println("error binding body")
	}

	//Querying
	rows, errC := db.Query(emailSQLStatement, bodyUser.Email)
	if errC != nil {
		fmt.Println("error querying email")
	}

	emailResult := rows.Next()
	if !emailResult {
		responseInvalidEail := Response{Message: "invalid email"}
		c.JSON(http.StatusNotFound, responseInvalidEail)
		return
	}
	fmt.Println("En next despues de email: " + strconv.FormatBool(emailResult))

	var user User

	errD := rows.Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.ImageURL, &user.Password)
	if errD != nil {
		fmt.Println(errD)
	}

	//select user by password
	passSQLStatement := `SELECT id FROM users WHERE password = crypt($1, password) AND id = $2;`

	//Querying
	rows, errF := db.Query(passSQLStatement, bodyUser.Password, user.ID)
	if errF != nil {
		fmt.Println("error querying pass")
	}

	passResult := rows.Next()
	if !passResult {
		responseInvalidPass := Response{Message: "invalid password"}
		c.JSON(http.StatusForbidden, responseInvalidPass)
		return
	}
	fmt.Println("En next despues de pass: " + strconv.FormatBool(passResult))

	errH := rows.Scan(&user.ID)
	if errH != nil {
		fmt.Println(errD)
	}
	user.Password = ""
	c.JSON(http.StatusOK, user)
}

// SignUp function
func SignUp(c *gin.Context) {
	//we open the database
	db, errA := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errA != nil {
		fmt.Println("error opening db")
	}
	//closing connection
	defer db.Close()

	//decoding body
	var bodyUser User

	errB := c.BindJSON(&bodyUser)
	if errB != nil {
		fmt.Println("error binding body")
	}

	//creating the statement
	sqlStatement := `INSERT INTO users (email, name, surname, password)
	VALUES ($1, $2, $3, crypt($4, gen_salt('bf'))) RETURNING id;`

	// TODO autoincremental
	var newID = 0

	//Querying
	errC := db.QueryRow(sqlStatement, bodyUser.Email, bodyUser.Name, bodyUser.Surname, bodyUser.Password).Scan(&newID)
	if errC != nil {
		fmt.Println(errC.Error())
	}

	bodyUser.ID = newID

	c.JSON(http.StatusOK, bodyUser)
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
	Password string `json:"password"`
}

// Response model
type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
