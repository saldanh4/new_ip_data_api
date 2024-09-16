package controller

import (
	"net"
	"net/http"
	l "new_ip_data_api/config/logger"
	"new_ip_data_api/usecase"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
func CheckIpEntrydata(c *gin.Context) (int, string, *GivenIP, error) {
	var givenIp GivenIP
	var message string

	//Checagem do body da requisição onde aponta o resultado para o endereço de memória de givenIp
	if err := c.ShouldBindBodyWithJSON(&givenIp); err != nil {
		message = "Given data error: " + err.Error()
		l.Logger.Warn(message, zap.Error(err))
		return http.StatusBadRequest, message, &GivenIP{}, err
	}

	//checagem do IP informado para confirmar se é um padrão de IPV4
	checkIP, _, err := net.ParseCIDR(givenIp.Ip + "/32")
	if err != nil {
		message = "O parâmetro informado não possui o padrão de um IP válido"
		l.Logger.Warn(message, zap.Error(err))
		return http.StatusBadRequest, message, &GivenIP{}, err
	} else if checkIP.To4() == nil {
		message = "O parâmetro informado não é um IP válido!"
		l.Logger.Warn(message, zap.Error(err))
		return http.StatusBadRequest, message, &GivenIP{}, err
	}

	l.Logger.Info(message, zap.Int("status", http.StatusOK))
	return http.StatusOK, message, &givenIp, nil
}

// Função para checagem dos dados de pesquisa por país
func CheckCountryEntrydata(c *gin.Context) (int, string, *GivenCountry, error) {
	var givenCountry GivenCountry

	if err := c.ShouldBindBodyWithJSON(&givenCountry); err != nil {
		value := "Given data error: " + err.Error()
		l.Logger.Warn(value, zap.Int("status", http.StatusBadRequest))
		return http.StatusBadRequest, value, &GivenCountry{}, err
	}
	l.Logger.Info("ok", zap.Int("status", http.StatusOK))
	return http.StatusOK, "ok", &givenCountry, nil
}

func CheckIpsEntrydata(c *gin.Context) (int, string, *GivenIps, error) {
	var givenIps GivenIps

	if err := c.ShouldBindBodyWithJSON(&givenIps); err != nil {
		value := "Given data error: " + err.Error()
		l.Logger.Warn(value, zap.Int("status", http.StatusBadRequest))
		return http.StatusBadRequest, value, &GivenIps{}, err
	}
	l.Logger.Info("ok", zap.Int("status", http.StatusOK))
	return http.StatusOK, "ok", &givenIps, nil
}
