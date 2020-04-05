package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HandlerInterface interface {
	SaveMessage( c *gin.Context)
}

type Handler struct {
	db dbLayer.DBLayer
}

func NewHandler() (*Handler, error) {
	return new(Handler), nil
}

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}
	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()

	r.POST("/SaveMessage", h.SaveMessage)

	return r.Run(address)
}

func (h *Handler) SaveMessage(c *gin.Context) {
	if h.db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}


