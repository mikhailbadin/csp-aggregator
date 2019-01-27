package controllers

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mikhailbadin/csp-aggregator/forms"
	"github.com/mikhailbadin/csp-aggregator/models"
)

// WriteReportHandler write csp_report handler
func WriteReportHandler(c *gin.Context) {
	var body forms.CSPReportMeta
	headers := forms.Headers{
		Referer: c.GetHeader("Referer"),
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "")
		log.Println("cannot bing csp report: ", err.Error())
		return
	}
	if err := writeReport(&body.Report, &headers, false); err != nil {
		c.String(http.StatusInternalServerError, "")
		log.Println("cannot write csp report: ", err.Error())
		return
	}
	c.String(http.StatusOK, "")
}

// WriteReportOnlyHandler write csp_report only handler
func WriteReportOnlyHandler(c *gin.Context) {
	var body forms.CSPReportMeta
	headers := forms.Headers{
		Referer: c.GetHeader("Referer"),
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "")
		log.Println("cannot bing csp report:", err.Error())
		return
	}
	if err := writeReport(&body.Report, &headers, true); err != nil {
		c.String(http.StatusInternalServerError, "")
		log.Println("cannot write csp report:", err.Error())
		return
	}
	c.String(http.StatusOK, "")
}

func writeReport(report *forms.CSPReport, headers *forms.Headers, reportOnly bool) error {
	var errRow, errLog error
	wg := sync.WaitGroup{}
	time := time.Now()

	wg.Add(2)
	go func() {
		errRow = models.WriteCSPRow(report, headers, reportOnly, time)
		wg.Done()
	}()
	go func() {
		errLog = models.WriteCSPLog(report, headers, reportOnly, time)
		wg.Done()
	}()
	wg.Wait()
	if errRow != nil {
		return fmt.Errorf("cannot write Row report: %s", errRow.Error())
	}
	if errLog != nil {
		return fmt.Errorf("cannot write Row report: %s", errLog.Error())
	}
	return nil
}
