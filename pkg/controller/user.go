package controller

import (
	"github.com/mgr1054/go-ticket/pkg/db"
	"github.com/mgr1054/go-ticket/pkg/models"
	"github.com/mgr1054/go-ticket/pkg/utils"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserUpdate struct {
	Name		string		`json:"name" example:"Max"`
	Username	string		`json:"username" example:"mgr"`
	Email		string		`json:"email" example:"mgr@online.de"`
	Password	string		`json:"password" example:"1234"`
	Role		string		`json:"role" example:"user"`
}

// @Summary 		Register User
// @Description		Creates new User, hashes Password for DB
// @Description		role automatically set to user
// @Description		allowed: unsecured
// @ID				register-user
// @Tags 			user
// @Produce 		json
// @Param			user body UserUpdate true "Create User"
// @Success 		201 {object} string
// @Failure			400 {object} string
// @Failure			500 {object} string
// @Router 			/user/register [post]
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

// @Summary 		Get User By ID
// @Description		Sends a User with ID
// @Description		allowed:  admin
// @ID				get-user-by-id
// @Tags 			user
// @Produce 		json
// @Success 		200 {object} models.User
// @Failure			401 {object} string
// @Failure			404 {object} string
// @Router 			/secured/user/{id} [get]
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

// @Summary 		Update User By ID
// @Description		Updates User with Body and corresponding ID 
// @Description		allowed:  admin
// @ID				update-user-by-id
// @Tags 			user
// @Produce 		json
// @Accept			json
// @Param			user body UserUpdate true "Update User"
// @Success 		200 {object} models.User
// @Failure			401 {object} string
// @Failure			404 {object} string
// @Router 			/secured/user/{id} [put]
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

// @Summary 		Delete User By ID
// @Description		Deltes User with corresponding ID 
// @Description		allowed:  admin
// @ID				delete-user-by-id
// @Tags 			user
// @Produce 		json
// @Success 		200 {object} string
// @Failure			401 {object} string
// @Failure			404 {object} string
// @Router 			/secured/user/{id} [delete]
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