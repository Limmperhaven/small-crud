package repository

import (
	"github.com/jmoiron/sqlx"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/models"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/repository/in-memory"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/repository/postgres"
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

func NewPostgresRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Record: postgres.NewRecordPostgres(db),
	}
}
