package handler

import (
	"canteen-prakerja/dto"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/pkg/helpers"
	"canteen-prakerja/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transaksiHandler struct {
	transaksiService service.TransaksiService
}

func NewTransaksiHandler(transaksiService service.TransaksiService) transaksiHandler {
	return transaksiHandler{
		transaksiService: transaksiService,
	}
}

func (th *transaksiHandler) GetAllTransaksi(c *gin.Context) {
	allTx, err := th.transaksiService.GetAllTransaksi()

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(allTx.StatusCode, allTx)
}

func (th *transaksiHandler) GetTransaksiById(c *gin.Context) {
	txId, err := helpers.GetParamsId(c, "txId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := th.transaksiService.GetTransaksiById(txId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (th *transaksiHandler) GetTransaksiDateBetween(c *gin.Context) {
	var getDate dto.TransaksiDateBetweenRequest

	if err := c.ShouldBindJSON(&getDate); err != nil {
		errBindJson := custerrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	dateTx, err := th.transaksiService.GetTransaksiDateBetween(&getDate)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(dateTx.StatusCode, dateTx)
}

func (th *transaksiHandler) CreateTransaksi(c *gin.Context) {
	var transaksiRequest dto.NewTransaksiRequest

	if err := c.ShouldBindJSON(&transaksiRequest); err != nil {
		errBindJson := custerrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	newTransaksi, err := th.transaksiService.CreateTransaksi(&transaksiRequest)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(newTransaksi.StatusCode, newTransaksi)
}

func (th *transaksiHandler) DeleteTransaksiById(c *gin.Context) {
	txId, err := helpers.GetParamsId(c, "txId")

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	response, err := th.transaksiService.DeleteTransaksiById(txId)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(response.StatusCode, response)
}
