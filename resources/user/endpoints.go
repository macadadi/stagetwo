package user

import(
	"github.com/gin-gonic/gin"
	"macadadi/stagetwo/db"
	"macadadi/stagetwo/services"
)

func Endpoints(route *gin.Engine, db db.DB, s *services.UserService){
	route.GET("/product",AddUser(db,s))

}