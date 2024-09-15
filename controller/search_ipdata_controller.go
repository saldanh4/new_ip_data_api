package controller

import (
	l "new_ip_data_api/config/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IpData struct {
	IP      string
	Isp     string
	Country string
	Count   int8
}

func (ipController *IpDataController) GetTotalSearchByIP(c *gin.Context) {
	var ipData *IpData

	status, message, givenIp, err := CheckIpEntrydata(c)
	if err != nil {
		l.Logger.Warn(message, zap.Int("status", status))
		c.AbortWithStatusJSON(status, gin.H{"message": message})
		return
	}

	status, message, result, err := ipController.ipDataUsecase.GetTotalSearchByIP(givenIp.Ip)
	if err != nil {
		l.Logger.Warn(message, zap.Int("status", status))
		c.AbortWithStatusJSON(status, gin.H{"message": message})
		return
	}

	ipData = &IpData{
		IP:      result.Query,
		Isp:     result.Isp,
		Country: result.Country,
		Count:   result.Count,
	}
	l.Logger.Info(message, zap.Int("status", status))
	c.IndentedJSON(status, gin.H{"1. message": message, "2. ipdata": ipData})
}
