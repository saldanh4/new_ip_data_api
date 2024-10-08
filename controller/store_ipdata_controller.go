package controller

import (
	"net/http"
	l "new_ip_data_api/config/logger"
	"new_ip_data_api/model"
	"time"

	"github.com/gin-gonic/gin"
	goip "github.com/jpiontek/go-ip-api"
	"go.uber.org/zap"
)

// função POST recebendo um objeto IpDataController
func (ipController *IpDataController) StoreIpData(c *gin.Context) {

	//checagem dos dados de entrada
	status, message, givenIp, err := CheckIpEntrydata(c)
	if err != nil {
		l.Logger.Warn(message, zap.Int("status", status))
		c.AbortWithStatusJSON(status, gin.H{"message": message, "status": status})
		return
	}

	//captura do momento da requisição (formatado na struct em model/ip_data)
	h := time.Now()

	//instanciando lib de IP geolocalização
	client := goip.NewClient()

	//checando a localização para o IP informado e atribuindo à variável result
	result, err := client.GetLocationForIp(string(givenIp.Ip))
	if err != nil {
		value := "Given IP error: " + err.Error()
		l.Logger.Warn(value, zap.Int("status", http.StatusBadRequest))
		c.AbortWithStatusJSON(http.StatusBadRequest, value)
		return
	}

	//Atribuir os dados de geolocalização e hora para ipData
	ipData := model.SetIpData(result, h)

	//Chamando a função para salvar os dados no banco e retornando os dados para exibir resposta ao usuário.
	status, statusMessage, informedIp, err := ipController.ipDataUsecase.StoreIpData(ipData)
	if err != nil {
		l.Logger.Info(statusMessage, zap.Int("status", status))
		c.IndentedJSON(status, gin.H{"status": status})
		return
	}

	l.Logger.Info(statusMessage, zap.Int("status", status))
	c.IndentedJSON(status, gin.H{
		"id": informedIp.Id,
		"ip": informedIp.Query},
	)

}
