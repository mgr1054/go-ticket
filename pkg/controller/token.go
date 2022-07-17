package controller

import (
	"github.com/mgr1054/go-ticket/pkg/utils"
	"github.com/mgr1054/go-ticket/pkg/db"
	"github.com/mgr1054/go-ticket/pkg/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

// struct for the incoming request
type TokenRequest struct {
	Email    string `json:"email" example:"test@online.de"`
	Password string `json:"password" example:"1234"`
}

// @Summary 		Generate Token
// @Description		Generates JWT Token based on given context, checks if username and password match
// @Description		Encode JWT with username, email and role
// @Description		allowed: unsecured
// @ID				generate-token
// @Tags 			auth
// @Produce 		json
// @Param			credentials body TokenRequest true "Create Token"
// @Success 		201 {object} string
// @Failure			400 {object} string
// @Failure			401 {object} string
// @Failure			404 {object} string
// @Failure			500 {object} string
// @Router 			/token [post]
func GenerateToken(c *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create Token"})
		c.Abort()
		return
	}
	// check if email exists and password is correct
	record := db.DB.Where("email = ?", request.Email).First(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"User not found"})
		c.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Password incorrect"})
		c.Abort()
		return
	}
	tokenString, err:= utils.GenerateJWT(user.Email, user.Username, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Could not create Token"})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"token": tokenString})
}