package usecase

import "new_ip_data_api/model"

func (ipUseCase *IpDataUsecase) GetTotalSearchByIP(ipNumber string) (*model.IpDataInfo, error) {

	ipData, err := ipUseCase.repository.GetTotalSearchByIP(ipNumber)
	if err != nil {
		return nil, err
	}

	return ipData, nil
}
