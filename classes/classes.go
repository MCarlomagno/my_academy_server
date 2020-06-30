package classes

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetClassesRoot function
func GetClassesRoot(c *gin.Context) {
	c.String(http.StatusOK, "Get a classes!")
}

// GetClassesByModuleID function
func GetClassesByModuleID(c *gin.Context) {
	//we open the database
	db, errA := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errA != nil {
		fmt.Println("error opening db")
	}
	//closing connection
	defer db.Close()

	//creating the statement
	sqlStatement := `SELECT id, id_module, title, description, video_url FROM classes WHERE id_module=$1;`

	//Querying
	rows, errB := db.Query(sqlStatement, c.Param("moduleId"))
	if errB != nil {
		// handle this error better than this
		fmt.Println("error opening rows")
	}
	defer rows.Close()

	classes := make([]*Class, 0)

	for rows.Next() {
		var id int
		var title string
		var moduleID int
		var description string
		var videoURL string
		class := new(Class)
		errC := rows.Scan(&id, &moduleID, &title, &description, &videoURL)
		if errC != nil {
			fmt.Println(errC)
		}
		class.ID = id
		class.ModuleID = moduleID
		class.Title = title
		class.Description = description
		class.VideoURL = videoURL
		classes = append(classes, class)
	}
	// response
	c.JSON(http.StatusOK, classes)
}

// PostClassesRoot function
func PostClassesRoot(c *gin.Context) {
	//we open the database
	db, errA := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if errA != nil {
		fmt.Println("error opening db")
	}

	//decoding body
	var bodyClass Class

	errB := c.BindJSON(&bodyClass)
	if errB != nil {
		fmt.Println("error binding body")
	}

	//creating the statement
	sqlStatement := `INSERT INTO classes (id_module, title, description, video_url)
				VALUES ($1, $2, $3, $4) RETURNING id;`

	// TODO autoincremental
	var newID = 0

	//Querying
	errC := db.QueryRow(sqlStatement, bodyClass.ModuleID, bodyClass.Title, bodyClass.Description, bodyClass.VideoURL).Scan(&newID)
	if errC != nil {
		fmt.Println(errC.Error())
	}

	bodyClass.ID = newID

	//closing connection
	defer db.Close()
	c.JSON(http.StatusOK, bodyClass)
}

// PutClassesRoot function
func PutClassesRoot(c *gin.Context) {
	c.String(http.StatusOK, "put a classes!")
}

// DeleteClassesRoot function
func DeleteClassesRoot(c *gin.Context) {
	c.String(http.StatusOK, "delete a classes!")
}

// Class model
type Class struct {
	ID          int    `json:"id"`
	ModuleID    int    `json:"moduleId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoURL    string `json:"videoUrl"`
}
