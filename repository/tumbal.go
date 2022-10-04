package repository

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	Id     int                   `uri:"id"`
	Name   string                `form:"name"`
	Email  string                `form:"email"`
	Avatar *multipart.FileHeader `form:"avatar" binding:"required"`
}

func main() {
	r := gin.Default()
	r.PUT("/user/:id", func(c *gin.Context) {
		var userObj user
		if err := c.ShouldBind(&userObj); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}

		if err := c.ShouldBindUri(&userObj); err != nil {
			c.String(http.StatusBadRequest, "bad request")
			return
		}

		err := c.SaveUploadedFile(userObj.Avatar, userObj.Avatar.Filename)
		if err != nil {
			c.String(http.StatusInternalServerError, "unknown error")
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"data":   userObj,
		})
	})
	r.Static("assets", "./assets")

	r.Run("localhost:8080")
}
