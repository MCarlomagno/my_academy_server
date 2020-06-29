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
