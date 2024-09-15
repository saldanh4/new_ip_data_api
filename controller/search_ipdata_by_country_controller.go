package controller

import (
	l "new_ip_data_api/config/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GivenCountry struct {
	Country string `json:"country" binding:"required"`
}

func (ipController *IpDataController) GetTotalSearchByCountry(c *gin.Context) {

	status, message, givenCountry, err := CheckCountryEntrydata(c)
	if err != nil {
		l.Logger.Warn(message, zap.Int("status", status))
		c.AbortWithStatusJSON(status, message)
		return
	}

	status, message, result, err := ipController.ipDataUsecase.GetTotalSearchByCountry(givenCountry.Country)
	if err != nil {
		l.Logger.Warn(message, zap.Int("status", status))
		c.AbortWithStatusJSON(status, gin.H{"message": message, "status": status})
		return
	}

	l.Logger.Info(message, zap.Int("status", status))
	c.IndentedJSON(status, gin.H{"result": result, "message": message})
}
