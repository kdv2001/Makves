package usecase

import (
	"Makves/model"
	"Makves/repository"
)

type UserUseCase interface {
	GetUsersByIds([]int64) ([]model.User, error)
}

type UserUC struct {
	userRepo repository.CSVRepository
}

func NewUserUseCase(csvRepository repository.CSVRepository) UserUseCase {
	return UserUC{userRepo: csvRepository}
}

func (uc UserUC) GetUsersByIds(ids []int64) ([]model.User, error) {
	return uc.userRepo.GetItemByIds(ids)
}
