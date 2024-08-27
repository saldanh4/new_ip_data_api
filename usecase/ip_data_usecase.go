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
