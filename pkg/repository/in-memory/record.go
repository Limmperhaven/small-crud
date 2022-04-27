package in_memory

import (
	"errors"
	uuid2 "github.com/google/uuid"
	"gitlab.digital-spirit.ru/study/artem_crud/models"
	"sync"
)

type RecordInMemory struct {
	db map[string]models.Record
	mu *sync.RWMutex
}

func NewRecordInMemory(db map[string]models.Record) *RecordInMemory {
	return &RecordInMemory{
		db: db,
		mu: new(sync.RWMutex),
	}
}

func (r *RecordInMemory) Create(record models.Record) (string, error) {
	uuid := uuid2.New().String()
	record.Uuid = uuid

	r.mu.Lock()
	r.db[uuid] = record
	r.mu.Unlock()

	return uuid, nil
}

func (r *RecordInMemory) GetByUid(recordUid string) (models.Record, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.db[recordUid] == *new(models.Record) {
		return models.Record{}, errors.New("record with such uuid was not found")
	}

	return r.db[recordUid], nil
}

func (r *RecordInMemory) GetByFilter(params models.RecordInput) ([]models.Record, error) {
	recordsList := make([]models.Record, 0)

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, rec := range r.db {
		if rec.FirstName != params.FirstName && params.FirstName != "" {
			continue
		}
		if rec.LastName != params.LastName && params.LastName != "" {
			continue
		}
		if rec.MobilePhone != params.MobilePhone && params.MobilePhone != "" {
			continue
		}
		if rec.HomePhone != params.HomePhone && params.HomePhone != "" {
			continue
		}
		recordsList = append(recordsList, rec)
	}
	return recordsList, nil
}

func (r *RecordInMemory) Update(recordUid string, input models.RecordInput) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	record := r.db[recordUid]

	if input.FirstName != "" {
		record.FirstName = input.FirstName
	}
	if input.LastName != "" {
		record.LastName = input.LastName
	}
	if input.MobilePhone != "" {
		record.MobilePhone = input.MobilePhone
	}
	if input.HomePhone != "" {
		record.HomePhone = input.HomePhone
	}

	r.db[recordUid] = record

	return nil
}

func (r *RecordInMemory) Delete(recordUid string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.db, recordUid)

	return nil
}
