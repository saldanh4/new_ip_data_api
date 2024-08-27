package controller

import (
	"fmt"
	"net/http"
	"new_ip_data_api/model"
	"time"

	"github.com/gin-gonic/gin"
	goip "github.com/jpiontek/go-ip-api"
)

// função POST recebendo um objeto IpDataController
func (ipController *IpDataController) StoreIpData(c *gin.Context) {

	//checagem dos dados de entrada
	givenIp, err := CheckIpEntrydata(c)
	if err != nil {
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
		c.AbortWithStatusJSON(http.StatusBadRequest, value)
		return
	}

	//Atribuir os dados de geolocalização e hora para ipData
	ipData := model.SetIpData(result, h)

	//Chamando a função para salvar os dados no banco e retornando os dados para exibir resposta ao usuário.
	informedIp, err := ipController.ipDataUsecase.StoreIpData(ipData)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)

	}
	mID := fmt.Sprintf("%d", informedIp.Id)
	mIP := informedIp.Query
	message := ("Criado o ID: " + mID + " com os dados do IP: " + mIP)
	c.IndentedJSON(http.StatusCreated, message)

}
