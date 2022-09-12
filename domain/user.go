package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/simplebase/models"
)

type UserRepository interface {
	List() ([]*models.User, error)
	Create(u *models.User, id ...string) error
	Get(id string) (*models.User, error)
	Update(u *models.User) error
	UpdatePassword(u *models.User) error
	Delete(id int) error
}

type UserService interface {
	List() ([]*models.User, error)
	Create(u *models.User) error
	Get(id string) (*models.User, error)
	Update(u *models.User) error
	UpdatePassword(u *models.User) error
	Delete(id int) error
}

type UserController interface {
	ListUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
}
