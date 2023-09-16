package handler

import (
	"canteen-prakerja/dto"
	"canteen-prakerja/pkg/custerrs"
	"canteen-prakerja/service"

	"github.com/gin-gonic/gin"
)

type reportHandler struct {
	reportService service.ReportService
}

func NewReportHandler(reportService service.ReportService) reportHandler {
	return reportHandler{
		reportService: reportService,
	}
}

func (rh *reportHandler) GetReportDate(c *gin.Context) {
	var getDate dto.DateRangeReportRequest

	if err := c.ShouldBindJSON(&getDate); err != nil {
		errBindJson := custerrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	report, err := rh.reportService.GetReportDateBetween(&getDate)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(report.StatusCode, report)
}

func (rh *reportHandler) GetReportDateBetween(c *gin.Context) {
	var getDate dto.DateRangeReportRequest

	if err := c.ShouldBindJSON(&getDate); err != nil {
		errBindJson := custerrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	report, err := rh.reportService.GetReportDateBetween(&getDate)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(report.StatusCode, report)
}

func (rh *reportHandler) GetTotalReportDate(c *gin.Context) {
	var getDate dto.DateRangeReportRequest

	if err := c.ShouldBindJSON(&getDate); err != nil {
		errBindJson := custerrs.NewUnprocessibleEntityError("invalid request body")
		c.JSON(errBindJson.Status(), errBindJson)
		return
	}

	total, err := rh.reportService.GetTotalReportDateBetween(&getDate)

	if err != nil {
		c.AbortWithStatusJSON(err.Status(), err)
		return
	}

	c.JSON(total.StatusCode, total)
}
