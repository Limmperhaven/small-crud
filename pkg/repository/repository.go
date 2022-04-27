package repository

import (
	"gitlab.digital-spirit.ru/study/artem_crud/models"
	"gitlab.digital-spirit.ru/study/artem_crud/pkg/repository/in-memory"
)

type Record interface {
	Create(record models.Record) (string, error)
	GetByUid(recordUid string) (models.Record, error)
	GetByFilter(params models.RecordInput) ([]models.Record, error)
	Update(recordUid string, record models.RecordInput) error
	Delete(recordUid string) error
}

type Repository struct {
	Record
}

func NewInMemoryRepository(db map[string]models.Record) *Repository {
	return &Repository{
		Record: in_memory.NewRecordInMemory(db),
	}
}
