package user

import (
	"github.com/gin-gonic/gin"
)

func NewUserHandler(r *gin.RouterGroup, repo *UserRepository) {

	r.POST("/register", Register)
	r.POST("/login", Login)
}
