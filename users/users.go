package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsersRoot(c *gin.Context) {
	c.String(http.StatusOK, "Get a users!")
}

func PostUsersRoot(c *gin.Context) {
	c.String(http.StatusOK, "Post a users!")
}

func PutUsersRoot(c *gin.Context) {
	c.String(http.StatusOK, "put a users!")
}

func DeleteUsersRoot(c *gin.Context) {
	c.String(http.StatusOK, "delete a users!")
}
