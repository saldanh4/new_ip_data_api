package controller

import (
	l "new_ip_data_api/config/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (ipDataController *IpDataController) DeleteIpDataByIp(c *gin.Context) {

	status, statusMessage, givenIp, err := CheckIpEntrydata(c)
	if err != nil {
		l.Logger.Info(statusMessage, zap.Int("status", status))
		c.AbortWithStatusJSON(status, gin.H{"message": statusMessage})
		return
	}

	status, message, err := ipDataController.ipDataUsecase.DeleteIpDataByIp(givenIp.Ip)
	if err != nil {
		l.Logger.Info(message, zap.Int("status", status))
		c.AbortWithStatusJSON(status, gin.H{"status": status, "message": message})
		return
	}

	l.Logger.Info(message, zap.Int("status", status))
	c.IndentedJSON(status, gin.H{"message": message})
}
