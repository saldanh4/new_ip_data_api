package controller

import (
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

	status, statusMessage, givenIp, err := CheckIpEntrydata(c)
	if err != nil {
		c.AbortWithStatusJSON(status, gin.H{"message": statusMessage})
		return
	}

	status, message, result, err := ipController.ipDataUsecase.GetTotalSearchByIP(givenIp.Ip)
	if err != nil {
		c.IndentedJSON(status, gin.H{"message": message})
		return
	}

	ipData = &IpData{
		IP:      result.Query,
		Isp:     result.Isp,
		Country: result.Country,
		Count:   result.Count,
	}

	c.IndentedJSON(status, gin.H{"1. message": message, "2. ipdata": ipData})
}
