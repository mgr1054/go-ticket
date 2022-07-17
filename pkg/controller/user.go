package controller

import (
	"github.com/mgr1054/go-ticket/pkg/db"
	"github.com/mgr1054/go-ticket/pkg/models"
	"github.com/mgr1054/go-ticket/pkg/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserUpdate struct {
	Name		string		`json:"name"`
	Username	string		`json:"username"`
	Email		string		`json:"email"`
	Password	string		`json:"password"`
	Role		string		`json:"role"`
}

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	
	user.Role = "user"
	
	if err := user.HashPassword(user.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	record := db.DB.Create(&user)
	if record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User could not be created"})
		c.Abort()
		return
	}
	c.JSON(http.StatusCreated, gin.H{"userId": user.ID, "email": user.Email, "username": user.Username, "role": user.Role})
}

func GetUserById (c *gin.Context) {

	if err := utils.CheckUserType(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}

	var user models.User

	if err := db.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserById (c *gin.Context) {

	if err := utils.CheckUserType(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}
	
	var user models.User

	if err := db.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var updateUser UserUpdate

	if err:= c.ShouldBindJSON(&updateUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "User could not be updated with provided data"})
        return
    }

	if err:= db.DB.Model(&user).Updates(models.User{
		Name: updateUser.Name,
		Username: updateUser.Username, 
		Email: updateUser.Email, 
		Password: updateUser.Password, 
		Role: updateUser.Role,
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update User"})
        return
	}

	c.JSON(http.StatusOK, user)
}

func DelteUserById (c *gin.Context) {

	if err := utils.CheckUserType(c, "admin"); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"Unauthorized for this route"})
		return
	}
	
	var user models.User

	if err := db.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found!"})
        return
    }

	if err := db.DB.Delete(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}