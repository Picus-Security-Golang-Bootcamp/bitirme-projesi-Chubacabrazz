package user

import (
	"net/http"

	"github.com/Chubacabrazz/picus-storeApp/storage/helper"
	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {

	user_id, err := helper.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

type LoginInput struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := User{}

	u.Email = input.Email
	u.Password = input.Password

	helper, err := LoginCheck(u.Email, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"helper": helper})

}

type RegisterInput struct {
	ID         int    `json:"ID"`
	Email      string `json:"Email" binding:"required"`
	Password   string `json:"Password" binding:"required"`
	First_name string `json:"Firstname"`
	Last_name  string `json:"Lastname"`
	IsAdmin    bool
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := User{}
	u.ID = input.ID
	u.Email = input.Email
	u.Password = input.Password
	u.First_name = input.First_name
	u.Last_name = input.Last_name
	u.IsAdmin = input.IsAdmin

	_, err := u.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}
