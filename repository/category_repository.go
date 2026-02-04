package repository

import "belajar-golang-unit-test/entity"

type CateogryRepository interface {
	FindById(id string) *entity.Category
}