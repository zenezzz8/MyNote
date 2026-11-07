package controller

import (
	"Backend/config"
	"Backend/model"
	"Backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// register
type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	// get input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Required input username and password"})
		return
	}

	//
	hashPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "fail enc password"})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: string(hashPass),
	}

	if err := config.ConnDB().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "fail create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Message": "User created successfully",
		"User":    user,
	})
}

// login
type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var input LoginInput

	//
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Required input username and password"})
		return
	}

	//
	var user model.User
	if err := config.ConnDB().Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid username or password"})
		return
	}

	//
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid username or password"})
		return
	}

	token, err := utils.GenerateToken(int(user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "fail generate token"})
		return
	}

	//
	c.SetCookie("jwt", token, 3600*24, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Login successful",
		"token":   token,
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"Message": "Logout successful",
	})
}
