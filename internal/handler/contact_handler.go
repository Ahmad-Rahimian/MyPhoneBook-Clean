package handler

import (
	"net/http"

	"phonebook/internal/domain"
	"phonebook/internal/service"

	"github.com/gin-gonic/gin"
)

type _ = domain.Contact

// GetContactsHandler godoc
// @Summary      Get all contacts
// @Description  Returns list of contacts
// @Tags         contacts
// @Produce      json
// @Success      200  {array}   domain.Contact
// @Failure      500  {object}  map[string]string
// @Router       /contacts [get]
func GetContactsHandler(c *gin.Context) {
	contacts, err := service.GetContacts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contacts)
}
