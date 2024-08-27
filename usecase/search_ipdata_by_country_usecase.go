package usecase

import "new_ip_data_api/model"

func (ipUseCase *IpDataUsecase) GetTotalSearchByCountry(givenCountry string) ([]model.IpDataInfo, error) {

	countryList, err := ipUseCase.repository.GetTotalSearchByCountry(givenCountry)
	if err != nil {
		return nil, err
	}

	return countryList, nil
}
