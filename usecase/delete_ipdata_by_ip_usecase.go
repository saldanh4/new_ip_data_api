package usecase

func (ipUseCase *IpDataUsecase) DeleteIpDataByIp(ipNumber string) (string, error) {
	message, err := ipUseCase.repository.DeleteIpDataByIp(ipNumber)
	if err != nil {
		return "deu ruim", err
	}
	return message, nil

}
