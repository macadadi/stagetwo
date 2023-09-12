package user

import (
	"net/http"
	"macadadi/stagetwo/db"
	"macadadi/stagetwo/form"
	"macadadi/stagetwo/services"

	"github.com/gin-gonic/gin"
)

func AddUser(db db.DB, s *services.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var form *form.User
		ctx := c.Request.Context()
		if err := c.BindJSON(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "You provided an invalid form"})
			return
		}

		if err := s.AddUser(ctx, db, form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "successfully added user"})
	}
}

func GetAllUsers(db db.DB, s *services.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		users, err := s.GetAllUsers(ctx, db)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
