package modules

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetModulesRoot(c *gin.Context) {
	c.String(http.StatusOK, "Get a modules!")
}

func PostModulesRoot(c *gin.Context) {
	c.String(http.StatusOK, "Post a modules!")
}

func PutModulesRoot(c *gin.Context) {
	c.String(http.StatusOK, "put a modules!")
}

func DeleteModulesRoot(c *gin.Context) {
	c.String(http.StatusOK, "delete a modules!")
}
