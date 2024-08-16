package controller

import (
	"fmt"
	"net"
	"net/http"
	"new_ip_data_api/model"
	"time"

	"github.com/gin-gonic/gin"
	goip "github.com/jpiontek/go-ip-api"
)

// função POST recebendo um objeto IpDataController
func (ipController *IpDataController) StoreIpData(c *gin.Context) {

	//captura do momento da requisição (formatado na struct em model/ip_data)
	h := time.Now()

	//Variável do tipo IpDataInfo que contém os campos da struct com os dados do IP necessários
	var ipData model.IpDataInfo

	//Checagem do body da requisição onde aponta o resultado para o endereço de memória de givenIp
	if err := c.ShouldBindJSON(&givenIp); err != nil {
		value := "Given data error: " + err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, value)
		return
	}

	//checagem do IP informado para confirmar se é um padrão de IPV4
	checkIP, _, err := net.ParseCIDR(givenIp.Ip + "/32")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "O parâmetro informado não possui o padrão de um IP válido")
		return
	} else if checkIP.To4() == nil {
		fmt.Printf("O parâmetro informado não é um IP válido: %v", err)
		return
	}

	//instanciando lib de IP geolocalização
	client := goip.NewClient()

	//checando a localização para o IP informado e atribuindo à variável result
	result, err := client.GetLocationForIp(string(givenIp.Ip))
	if err != nil {
		value := "Given IP error: " + err.Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, value)
		return
	}

	//atribuido os dados de geolocalização e hora para ipData
	ipData = model.SetIpData(result, h)

	//Chamando a função para salvar os dados no banco e retornando os dados para exeibir resposta ao usuário.
	informedIp, err := ipController.ipDataUsecase.StoreIpData(ipData)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)

	}
	mID := fmt.Sprintf("%d", informedIp.Id)
	mIP := informedIp.Query
	message := ("Criado o ID: " + mID + " com os dados do IP: " + mIP)
	c.IndentedJSON(http.StatusCreated, message)

}
