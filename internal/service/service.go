package service

import (
	"gitlab.digital-spirit.ru/study/artem_crud/internal/models"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/repository"
)

type Record interface {
	GetByUuid(recordUid string) (models.Record, error)
	GetByFilter(params models.RecordInput) ([]models.Record, error)
	Create(record models.Record) (string, error)
	Update(recordUid string, input models.RecordInput) error
	Delete(recordUid string) error
}

type Service struct {
	Record
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Record: NewRecordService(repo),
	}
}
