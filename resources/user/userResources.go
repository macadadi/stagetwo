package user

import (
	"fmt"
	"macadadi/stagetwo/db"
	"macadadi/stagetwo/form"
	"macadadi/stagetwo/services"
	"net/http"
	"strconv"

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

func FindUserById(db db.DB, s *services.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id, _ := strconv.ParseInt(c.Param("id"), 0, 64)
		users, err := s.FindUserByID(ctx, db, id)

		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, users)
	}
}

func UpdateUser(db db.DB, s *services.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user *form.User
		ctx := c.Request.Context()
		id, _ := strconv.ParseInt(c.Param("user_id"), 0, 64)
		c.ShouldBind(&user)
		user.Id = id

		users, err := s.UpdateUser(ctx, db, user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err_message": fmt.Sprintf("Error when updating %s", err.Error())})
			return
		}
		c.JSON(http.StatusOK, users)
	}
}
func DeleteUser(db db.DB, s *services.UserService) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		id, _ := strconv.ParseInt(c.Param("user_id"), 0, 64)
		err := s.DeleteUser(ctx, db, id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
	}
}
