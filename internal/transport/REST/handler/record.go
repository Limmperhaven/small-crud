package restHandler

import (
	"github.com/gin-gonic/gin"
	"gitlab.digital-spirit.ru/study/artem_crud/models"
	"net/http"
)

func (h *Handler) getRecordsByFilter(c *gin.Context) {
	var params models.RecordInput

	if err := c.BindQuery(&params); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	lists, err := h.services.GetByFilter(params)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lists)
}

func (h *Handler) getRecordById(c *gin.Context) {
	recordUid := c.Param("uuid")

	record, err := h.services.GetById(recordUid)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, record)
}

func (h *Handler) createRecord(c *gin.Context) {
	var input models.Record
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	uid, err := h.services.Record.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"uuid": uid,
	})
}

func (h *Handler) updateRecord(c *gin.Context) {
	recordUid := c.Param("uuid")

	var input models.RecordInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Update(recordUid, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, make(map[string]interface{}))
}

func (h *Handler) deleteRecord(c *gin.Context) {
	recordUid := c.Param("uuid")

	if err := h.services.Delete(recordUid); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, make(map[string]interface{}))
}
