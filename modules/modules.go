package modules

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetModulesRoot function
func GetModulesRoot(c *gin.Context) {
	c.String(http.StatusOK, "Get a modules!")
}

// PostModulesRoot function
func PostModulesRoot(c *gin.Context) {
	//we open the database
	db, errA := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errA != nil {
		fmt.Println("error opening db")
	}

	//decoding body
	var bodyModule Module

	errB := c.BindJSON(&bodyModule)
	if errB != nil {
		fmt.Println("error binding body")
	}

	//creating the statement
	sqlStatement := `INSERT INTO modules (id_course, title, description)
			VALUES ($1, $2, $3) RETURNING id;`

	// TODO autoincremental
	var newID = 0

	//Querying
	errC := db.QueryRow(sqlStatement, bodyModule.CourseID, bodyModule.Title, bodyModule.Description).Scan(&newID)
	if errC != nil {
		fmt.Println(errC.Error())
	}

	bodyModule.ID = newID

	//closing connection
	defer db.Close()
	c.JSON(http.StatusOK, bodyModule)
}

// GetModulesByCourseID function
func GetModulesByCourseID(c *gin.Context) {
	//we open the database
	db, errA := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errA != nil {
		fmt.Println("error opening db")
	}
	//closing connection
	defer db.Close()

	//creating the statement
	sqlStatement := `SELECT id, id_course, title, description FROM modules WHERE id_course=$1;`

	//Querying
	rows, errB := db.Query(sqlStatement, c.Param("courseId"))
	if errB != nil {
		// handle this error better than this
		fmt.Println("error opening rows")
	}
	defer rows.Close()

	modules := make([]*Module, 0)

	for rows.Next() {
		var id int
		var title string
		var courseID int
		var description string
		module := new(Module)
		errC := rows.Scan(&id, &courseID, &title, &description)
		if errC != nil {
			fmt.Println(errC)
		}
		module.ID = id
		module.CourseID = courseID
		module.Title = title
		module.Description = description
		modules = append(modules, module)
	}
	// response
	c.JSON(http.StatusOK, modules)
}

// PutModulesRoot function
func PutModulesRoot(c *gin.Context) {
	c.String(http.StatusOK, "put a modules!")
}

// DeleteModulesRoot function
func DeleteModulesRoot(c *gin.Context) {
	c.String(http.StatusOK, "delete a modules!")
}

// Module model
type Module struct {
	ID          int    `json:"id"`
	CourseID    int    `json:"courseId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
