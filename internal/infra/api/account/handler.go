package account

import (
	"net/http"

	"github.com/deBeloper-code/goFinance/internal/pkg/ports"
	"github.com/gin-gonic/gin"
)

type accountHandler struct {
	cardService ports.AccountService
}

func NewHandler(service ports.AccountService) *accountHandler {
	return &accountHandler{
		cardService: service,
	}
}

func (s *accountHandler) Add(c *gin.Context) {
	// s.cardService.Add()
	c.JSON(http.StatusOK, nil)
}
func (s *accountHandler) GetUserAccount(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
	// s.cardService.Add()
}
