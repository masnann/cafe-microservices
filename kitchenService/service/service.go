package service

import "kitchen-service/repository"

type KitchenService struct {
	repo repository.KitchenRepositoryInterface
}

func NewKitchenService(repo repository.KitchenRepositoryInterface) KitchenServiceInterface {
	return KitchenService{
		repo: repo,
	}
}

