package repository

import (
	"belajar-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CateogryRepositoryMock struct {
	Mock mock.Mock
}

func (repo *CateogryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repo.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}