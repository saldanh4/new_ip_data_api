package controller

import (
	"new_ip_data_api/usecase"
)

type IpDataController struct {
	ipDataUsecase usecase.IpDataUsecase
}

// Defini uma variável do tipo struct para receber o IP via Body no JSON, setando o padrão esperado
var givenIp struct {
	Ip string `json:"ip" binding:"required"`
}

func NewIpDataController(usecase usecase.IpDataUsecase) IpDataController {
	return IpDataController{
		ipDataUsecase: usecase,
	}
}
