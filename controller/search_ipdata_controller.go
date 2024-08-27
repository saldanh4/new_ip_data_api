package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IpData struct {
	IP      string
	Isp     string
	Country string
	Count   int8
}

func (ipController *IpDataController) GetTotalSearchByIP(c *gin.Context) {
	var ipData *IpData

	givenIp, err := CheckIpEntrydata(c)
	if err != nil {
		return
	}

	result, err := ipController.ipDataUsecase.GetTotalSearchByIP(givenIp.Ip)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	ipData = &IpData{
		IP:      result.Query,
		Isp:     result.Isp,
		Country: result.Country,
		Count:   result.Count,
	}

	c.IndentedJSON(http.StatusOK, ipData)
}
