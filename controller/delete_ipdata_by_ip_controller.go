package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ipDataController *IpDataController) DeleteIpDataByIp(c *gin.Context) {
	givenIp, err := CheckIpEntrydata(c)
	if err != nil {
		return
	}

	result, err := ipDataController.ipDataUsecase.DeleteIpDataByIp(givenIp.Ip)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusOK, result)

}
