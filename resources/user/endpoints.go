package user

import(
	"github.com/gin-gonic/gin"
	"macadadi/stagetwo/db"
	"macadadi/stagetwo/services"
)

func Endpoints(route *gin.Engine, db db.DB, s *services.UserService){
	route.POST("/api",AddUser(db,s))
	route.GET("/api",GetAllUsers(db,s))
}