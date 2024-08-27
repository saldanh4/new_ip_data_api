package controller

import (
	"fmt"
	"net"
	"net/http"
	"new_ip_data_api/usecase"

	"github.com/gin-gonic/gin"
)

type IpDataController struct {
	ipDataUsecase usecase.IpDataUsecase
}
type GivenIP struct {
	Ip string `json:"ip" binding:"required"`
}

func NewIpDataController(usecase usecase.IpDataUsecase) IpDataController {
	return IpDataController{
		ipDataUsecase: usecase,
	}
}

// Função para checagem dos dados de pesquisa por IP
func CheckIpEntrydata(c *gin.Context) (*GivenIP, error) {
	var givenIp GivenIP
	//Checagem do body da requisição onde aponta o resultado para o endereço de memória de givenIp
	if err := c.ShouldBindBodyWithJSON(&givenIp); err != nil {
		value := "Given data error: " + err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, value)
		return &GivenIP{}, err
	}

	//checagem do IP informado para confirmar se é um padrão de IPV4
	checkIP, _, err := net.ParseCIDR(givenIp.Ip + "/32")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "O parâmetro informado não possui o padrão de um IP válido")
		return &GivenIP{}, err
	} else if checkIP.To4() == nil {
		fmt.Printf("O parâmetro informado não é um IP válido: %v", err)
		return &GivenIP{}, err
	}

	return &givenIp, nil
}

// Função para checagem dos dados de pesquisa por país
func CheckCountryEntrydata(c *gin.Context) (*GivenCountry, error) {
	var givenCountry GivenCountry

	if err := c.ShouldBindBodyWithJSON(&givenCountry); err != nil {
		value := "Given data error: " + err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, value)
		return &GivenCountry{}, err
	}
	return &givenCountry, nil
}
