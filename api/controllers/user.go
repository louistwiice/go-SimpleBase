package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/simplebase/domain"
	"github.com/louistwiice/go/simplebase/models"
)

type controller struct {
	service domain.UserService
}

// serializer to update a user
type updateUser struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// Serializer to change a password
type changePassword struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

func NewUserController(svc domain.UserService) *controller {
	return &controller{
		service: svc,
	}
}

func (c *controller) ListUser(ctx *gin.Context) {
	users, err := c.service.List()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": http.StatusBadRequest})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successful", "data": users, "code": http.StatusOK})
}

func (c *controller) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": http.StatusBadRequest})
		return
	}

	err := user.Validate()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"messages": err.Error(), "code": http.StatusBadRequest})
		return
	}

	err = c.service.Create(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": http.StatusBadRequest})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successful", "data": user, "code": http.StatusOK})
}

// Retrieve a user with his id
func (c *controller) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := c.service.Get(id)
	if err != nil || user == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": models.ErrNotFound.Error(), "code": http.StatusBadRequest})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "successful", "data": user, "code": http.StatusOK})
}

func (c *controller) UpdateUser(ctx *gin.Context) {
	var data updateUser
	var id = ctx.Param("id")

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": http.StatusBadRequest})
		return
	}

	user, err := c.service.Get(id)
	if err != nil || user == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": models.ErrNotFound.Error(), "code": http.StatusBadRequest})
		return
	}

	if data.FirstName != "" {
		user.FirstName = data.FirstName
	}
	if data.LastName != "" {
		user.LastName = data.LastName
	}
	if data.Email != "" {
		user.Email = data.Email
	}

	err = c.service.Update(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"messages": err.Error(), "code": http.StatusBadRequest})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "good", "data": user, "code": http.StatusOK})
}

// Update user password
func (c *controller) UpdatePassword(ctx *gin.Context) {
	var data changePassword
	var id = ctx.Param("id")

	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": http.StatusBadRequest})
		return
	}

	user, err := c.service.Get(id)
	if err != nil || user == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": models.ErrNotFound.Error(), "code": http.StatusBadRequest})
		return
	}

	err = user.CheckPassword(data.OldPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": gin.H{"info": "old password does not match", "details": err.Error()}, "code": http.StatusBadRequest})
		return
	}

	user.Password = data.NewPassword
	err = c.service.UpdatePassword(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "code": http.StatusBadRequest})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Password reset successfully", "code": http.StatusOK})
}

// This part is to create all endpoints relatived to users controllers
func (c *controller) MakeUserHandlers(app *gin.RouterGroup) {
	app.GET("", c.ListUser)
	app.POST("", c.CreateUser)
	app.GET(":id", c.GetUser)
	app.PATCH(":id", c.UpdateUser)
	app.POST(":id/reset_password", c.UpdatePassword)
}
