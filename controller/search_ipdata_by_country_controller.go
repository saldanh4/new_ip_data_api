package controller

import (
	"github.com/gin-gonic/gin"
)

type GivenCountry struct {
	Country string `json:"country" binding:"required"`
}

func (ipController *IpDataController) GetTotalSearchByCountry(c *gin.Context) {

	givenCountry, err := CheckCountryEntrydata(c)
	if err != nil {
		return
	}

	status, message, result, err := ipController.ipDataUsecase.GetTotalSearchByCountry(givenCountry.Country)
	if err != nil {
		c.IndentedJSON(status, gin.H{"message": message, "status": status})
		return
	}

	c.IndentedJSON(status, gin.H{"result": result, "message": message})
}
