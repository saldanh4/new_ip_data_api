package controller

import (
	l "new_ip_data_api/config/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GivenIps struct {
	IpOne string `json:"ip1" binding:"required"`
	IpTwo string `json:"ip2" binding:"required"`
}

func (ipController *IpDataController) DistanciaPcaSe(c *gin.Context) {
	l.Logger.Info("", zap.Int("", 1))

	status, message, givenIps, err := CheckIpsEntrydata(c)
	if err != nil {
		l.Logger.Warn(message, zap.Int("status", status))
		c.AbortWithStatusJSON(status, message)
		return
	}

	status, message, result, err := ipController.ipDataUsecase.DistanciaPcaSe(givenIps.IpOne, givenIps.IpTwo)
	if err != nil {
		l.Logger.Warn(message, zap.Int("status", status))
		c.AbortWithStatusJSON(status, gin.H{"message": message, "status": status})
		return
	}

	l.Logger.Info(message, zap.Int("status", status))
	c.IndentedJSON(status, gin.H{"result": result, "message": message})
}
