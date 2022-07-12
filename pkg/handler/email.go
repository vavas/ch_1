package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addEmail(c *gin.Context) {
	email := c.Param("email")

	err := h.services.Create(email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, statusResponse{""})
}

func (h *Handler) removeEmail(c *gin.Context) {
	email := c.Param("email")
	err := h.services.Email.Delete(email)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) getAllEmails(c *gin.Context) {
	emails, err := h.services.Email.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, emails)
}
