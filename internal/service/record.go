package service

import (
	"gitlab.digital-spirit.ru/study/artem_crud/internal/models"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/repository"
)

type RecordService struct {
	repo repository.Record
}

func NewRecordService(repo repository.Record) *RecordService {
	return &RecordService{repo: repo}
}

func (s *RecordService) Create(record models.Record) (string, error) {
	return s.repo.Create(record)
}

func (s *RecordService) GetById(recordUid string) (models.Record, error) {
	return s.repo.GetByUid(recordUid)
}

func (s *RecordService) GetByFilter(params models.RecordInput) ([]models.Record, error) {
	return s.repo.GetByFilter(params)
}

func (s *RecordService) Update(recordUid string, input models.RecordInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(recordUid, input)
}

func (s *RecordService) Delete(recordUid string) error {
	return s.repo.Delete(recordUid)
}
