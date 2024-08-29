package controller

import (
	"github.com/gin-gonic/gin"
)

func (ipDataController *IpDataController) DeleteIpDataByIp(c *gin.Context) {
	// givenIp, err := CheckIpEntrydata(c)
	// if err != nil {
	// 	return
	// }
	status, statusMessage, givenIp, err := CheckIpEntrydata(c)
	if err != nil {
		c.AbortWithStatusJSON(status, gin.H{"message": statusMessage})
		return
	}

	status, message, err := ipDataController.ipDataUsecase.DeleteIpDataByIp(givenIp.Ip)
	if err != nil {
		c.IndentedJSON(status, gin.H{"status": status, "message": message})
		return
	}

	c.IndentedJSON(status, gin.H{"message": message})
}
