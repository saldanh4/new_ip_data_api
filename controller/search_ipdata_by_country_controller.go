package controller

import (
	"net/http"

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

	result, err := ipController.ipDataUsecase.GetTotalSearchByCountry(givenCountry.Country)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusOK, result)
}
