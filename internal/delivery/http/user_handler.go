package http

import (
	"jibas-template/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandlerInterface interface {
	RegisterRoutes(router *gin.RouterGroup)
}

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) UserHandlerInterface {
	return &UserHandler{userUsecase}
}

func (h *UserHandler) RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/users", h.GetAllUsers)
	router.POST("/users", h.CreateUser)
}

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve all user records
// @Tags Users
// @Accept  json
// @Produce  json
// @Success 200 {array} domain.User
// @Failure 401 {object} domain.ErrorResponse
// @Router /internal/users [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// CreateUser godoc
// @Summary Create a new user
// @Description Add a new user to the system
// @Tags Users
// @Accept  json
// @Produce  json
// @Param   user  body  domain.User  true  "User data"
// @Success 201 {object} domain.User
// @Failure 400 {object} domain.ErrorResponse
// @Router /internal/users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Error: err.Error()})
		return
	}

	err := h.userUsecase.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
