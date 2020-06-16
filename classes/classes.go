package classes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetClassesRoot(c *gin.Context) {
	c.String(http.StatusOK, "Get a classes!")
}

func PostClassesRoot(c *gin.Context) {
	c.String(http.StatusOK, "Post a classes!")
}

func PutClassesRoot(c *gin.Context) {
	c.String(http.StatusOK, "put a classes!")
}

func DeleteClassesRoot(c *gin.Context) {
	c.String(http.StatusOK, "delete a classes!")
}
