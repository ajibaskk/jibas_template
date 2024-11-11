package http

import (
	"jibas-template/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	RegisterRoutes(router *gin.Engine)
}

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) UserHandlerInterface {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/users", h.GetAllUsers)
	router.POST("/users", h.CreateUser)
}

func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.userUsecase.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
