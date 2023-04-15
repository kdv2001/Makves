package repository

import (
	"Makves/appErrors"
	"Makves/model"
	"github.com/gocarina/gocsv"
	"os"
)

type CSVRepository interface {
	GetItemByIds([]int64) ([]model.User, error)
}

type CSVRepo struct {
	data map[int64]model.User
}

func initData(filePath string) (map[int64]model.User, error) {
	userData := make(map[int64]model.User)

	in, err := os.Open(filePath)
	if err != nil {
		return userData, err
	}
	defer in.Close()

	clients := []model.User{}
	if err = gocsv.UnmarshalFile(in, &clients); err != nil {
		return userData, err
	}

	for _, client := range clients {
		userData[int64(client.Id)] = client
	}

	return userData, nil
}

func NewSCVRepo(pathToFile string) (CSVRepository, error) {
	userData, err := initData(pathToFile)
	if err != nil {
		return nil, err
	}

	return CSVRepo{data: userData}, nil
}

func (r CSVRepo) GetItemByIds(ids []int64) ([]model.User, error) {
	result := make([]model.User, 0)
	for _, id := range ids {
		val, find := r.data[id]
		if !find {
			return nil, appErrors.ErrItemNotFound
		}

		result = append(result, val)
	}

	return result, nil
}
