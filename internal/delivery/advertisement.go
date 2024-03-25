package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vk-task/internal/models"
)

func (h *Handler) createAdvert(c *gin.Context) {
	login, err := getUserLogin(c)
	if err != nil {
		return
	}
	var input models.Advert
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	output, err := h.services.Create(login, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, output)
}
