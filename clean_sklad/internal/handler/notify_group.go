package handler

import (
	"net/http"

	"github.com/Alang0r/vypolnyator/clean_sklad/internal/adapter/http/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateNotifyGroup(c *gin.Context) {
	input := dto.CreateNotifyGroupDTO{}
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	//ToDo: create DTO
	// err := h.usecase.CreateNotifyGroup(input)
	// if err != nil {
	// 	newErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	c.JSON(http.StatusOK, map[string]interface{}{
		//"id": id,
	})

}
