package service

import (
	"gitlab.digital-spirit.ru/study/artem_crud/internal/repository"
	"gitlab.digital-spirit.ru/study/artem_crud/models"
)

type Record interface {
	GetById(recordUid string) (models.Record, error)
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
