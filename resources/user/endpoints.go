package user

import (
	"macadadi/stagetwo/db"
	"macadadi/stagetwo/services"

	"github.com/gin-gonic/gin"
)

func Endpoints(route *gin.Engine, db db.DB, s *services.UserService) {
	route.POST("/api", AddUser(db, s))
	route.GET("/api", GetAllUsers(db, s))
	route.GET("/api/:user_id", FindUserById(db, s))
	route.PUT("/api/:user_id", UpdateUser(db, s))
	route.DELETE("/api/:user_id", DeleteUser(db, s))
}
