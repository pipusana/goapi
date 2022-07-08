package repositories

import (
	"github.com/pipusana/goapi/adapters"
	"github.com/pipusana/goapi/entities"
)

type NisitRepository interface {
	CreateNisit(post *entities.Nisit) (string, error)
	FindAllNisit() ([]entities.Nisit, error)
	FindOneNisit(nisit_id string) (entities.Nisit, error)
	UpdateOneNisit(nisit_id string, nisit_update entities.NisitUpdate) error
	DeleteOneNisit(nisit_id string) error
}

type nisitRepository struct {
	database adapters.MongoAdapter
}

func NewNisitRepository(database adapters.MongoAdapter) NisitRepository {
	return &nisitRepository{
		database: database,
	}
}

func (n *nisitRepository) CreateNisit(nisit *entities.Nisit) (string, error) {
	result, handleErr := n.database.InsertOne(nisit)
	return result, handleErr
}

func (n *nisitRepository) FindAllNisit() ([]entities.Nisit, error) {
	result, handleErr := n.database.FindAll()
	return result, handleErr
}

func (n *nisitRepository) FindOneNisit(nisit_id string) (entities.Nisit, error) {
	result, handleErr := n.database.FindOne(nisit_id)
	return result, handleErr
}

func (n *nisitRepository) UpdateOneNisit(nisit_id string, nisit_update entities.NisitUpdate) error {
	handleErr := n.database.UpdateOne(nisit_id, nisit_update)
	return handleErr
}

func (n *nisitRepository) DeleteOneNisit(nisit_id string) error {
	handleErr := n.database.DeleteOne(nisit_id)
	return handleErr
}
