package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scratchingmycranium/golang-rest/interfaces"
)

type userHandler struct {
	userSvc interfaces.IUserService
}

func InitUserHandler(userSvc interfaces.IUserService) *userHandler {
	return &userHandler{
		userSvc: userSvc,
	}
}

func (h *userHandler) RegisterUserRoutes(r *gin.Engine) {
	r.GET("/users", h.getUser)
}

// Handler for GET /users/:id
func (h *userHandler) getUser(c *gin.Context) {
	// id := c.Param("id")

	user, err := h.userSvc.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
