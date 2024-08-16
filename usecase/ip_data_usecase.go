package usecase

import (
	"new_ip_data_api/repository"
)

type IpDataUsecase struct {
	repository repository.IpDataRepository
}

func NewIpDataUsecase(repo repository.IpDataRepository) IpDataUsecase {
	return IpDataUsecase{
		repository: repo,
	}
}

// func (ipUseCase *IpDataUsecase) StoreIpData(ipData model.IpDataInfo) (model.IpDataInfo, error) {
// 	//pegar o ID gerado ao salvar o IP no banco e retornar em um tipo de Ipdata para exibição
// 	ipDataID, err := ipUseCase.repository.StoreIpData(ipData)
// 	if err != nil {
// 		return model.IpDataInfo{}, err
// 	}

// 	ipData.Id = ipDataID

// 	return ipData, nil
// }
