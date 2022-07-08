package usecases

import (
	"github.com/pipusana/goapi/entities"
	"github.com/pipusana/goapi/repositories"
)

type NisitUsecase interface {
	CreateNisit(input *entities.Nisit) (string, error)
	FindAllNisit() ([]entities.Nisit, error)
	FindOneNisit(nisit_id string) (entities.Nisit, error)
	UpdateOneNisit(nisit_id string, nisit_update entities.NisitUpdate) error
	DeleteOneNisit(nisit_id string) error
}

type nisitUsecase struct {
	nisitRepo repositories.NisitRepository
	logRepo   repositories.LoggerRepository
}

// #ถ้า return interface ค่าที่จะ return ออกไปจะเป็น addr ของ struct นั้นๆ
func NewNisitUseCase(repoNisit repositories.NisitRepository, repoLogger repositories.LoggerRepository) NisitUsecase {
	return &nisitUsecase{
		nisitRepo: repoNisit,
		logRepo:   repoLogger,
	}
}

func (nu *nisitUsecase) CreateNisit(nisit *entities.Nisit) (string, error) {
	nu.logRepo.Log("[Create] nisits")
	return nu.nisitRepo.CreateNisit(nisit)
}

func (nu *nisitUsecase) FindAllNisit() ([]entities.Nisit, error) {
	nu.logRepo.Log("[Find] all nisits")
	return nu.nisitRepo.FindAllNisit()
}

func (nu *nisitUsecase) FindOneNisit(nisit_id string) (entities.Nisit, error) {
	nu.logRepo.Log("[Find] one nisits")
	return nu.nisitRepo.FindOneNisit(nisit_id)
}

func (nu *nisitUsecase) UpdateOneNisit(nisit_id string, nisit_updates entities.NisitUpdate) error {
	nu.logRepo.Log("[Update] one nisits")
	return nu.nisitRepo.UpdateOneNisit(nisit_id, nisit_updates)
}

func (nu *nisitUsecase) DeleteOneNisit(nisit_id string) error {
	nu.logRepo.Log("[Delete] one nisits")
	return nu.nisitRepo.DeleteOneNisit(nisit_id)
}
