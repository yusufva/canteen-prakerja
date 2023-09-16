package handler

import (
	"canteen-prakerja/dto"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/pkg/helpers"
	"canteen-prakerja/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type barangHandler struct {
	barangService service.BarangService
}

func NewBarangHandler(barangService service.BarangService) barangHandler {
	return barangHandler{
		barangService: barangService,
	}
}

func (bh *barangHandler) GetAllBarang(c *gin.Context) {
	allBarangs, err := bh.barangService.GetAllBarang()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(allBarangs.Statuscode, allBarangs)
}

func (bh *barangHandler) GetBarangById(c *gin.Context) {
	barangId, err := helpers.GetParamsId(c, "barangId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
	}

	response, err := bh.barangService.GetBarangById(barangId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, &response)
}

func (bh *barangHandler) CreateNewBarang(c *gin.Context) {
	var barangRequest dto.NewBarangRequest

	if err := c.ShouldBindJSON(&barangRequest); err != nil {
		errBindJson := custerrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	newBarang, err := bh.barangService.CreateBarang(barangRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(newBarang.StatusCode, newBarang)
}

func (bh *barangHandler) UpdateBarangById(c *gin.Context) {
	var barangRequest dto.NewBarangRequest

	if err := c.ShouldBindJSON(&barangRequest); err != nil {
		errBindJson := custerrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	barangId, err := helpers.GetParamsId(c, "barangId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := bh.barangService.UpdateBarangById(barangId, barangRequest)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}

func (bh *barangHandler) DeleteBarangById(c *gin.Context) {
	barangId, err := helpers.GetParamsId(c, "barangId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := bh.barangService.DeleteBarangById(barangId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}
