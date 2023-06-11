package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/scratchingmycranium/golang-rest/interfaces"
)

type userHandler struct {
	userRepo interfaces.IUserRepo
}

func InitUserHandler(userRepo interfaces.IUserRepo) *userHandler {
	return &userHandler{
		userRepo: userRepo,
	}
}

func (h *userHandler) RegisterUserRoutes(r *gin.Engine) {
	r.GET("/users", h.getUser)
}

// Handler for GET /users/:id
func (h *userHandler) getUser(c *gin.Context) {
	// id := c.Param("id")

	user, err := h.userRepo.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
