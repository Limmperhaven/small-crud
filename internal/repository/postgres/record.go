package postgres

import (
	"fmt"
	uuid2 "github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"gitlab.digital-spirit.ru/study/artem_crud/models"
	"strings"
)

type RecordPostgres struct {
	db *sqlx.DB
}

func NewRecordPostgres(db *sqlx.DB) *RecordPostgres {
	return &RecordPostgres{
		db: db,
	}
}

func (r *RecordPostgres) Create(record models.Record) (string, error) {
	uuid := uuid2.New().String()

	query := fmt.Sprintf("INSERT INTO %s VALUES ($1, $2, $3, $4, $5)", recordsTable)

	_, err := r.db.Exec(query, uuid, record.FirstName, record.LastName, record.MobilePhone, record.HomePhone)

	return uuid, err
}

func (r *RecordPostgres) GetByUid(recordUid string) (models.Record, error) {
	var record models.Record

	query := fmt.Sprintf("SELECT * FROM %s WHERE uuid=$1", recordsTable)

	err := r.db.Get(&record, query, recordUid)

	return record, err
}

func (r *RecordPostgres) GetByFilter(params models.RecordInput) ([]models.Record, error) {
	records := make([]models.Record, 0)
	var setValues []string
	var args []interface{}
	argID := 1

	query := fmt.Sprintf("SELECT * FROM %s", recordsTable)

	if params.FirstName != "" {
		setValues = append(setValues, fmt.Sprintf("first_name=$%d", argID))
		args = append(args, params.FirstName)
		argID++
	}

	if params.LastName != "" {
		setValues = append(setValues, fmt.Sprintf("last_name=$%d", argID))
		args = append(args, params.LastName)
		argID++
	}

	if params.MobilePhone != "" {
		setValues = append(setValues, fmt.Sprintf("mobile_phone=$%d", argID))
		args = append(args, params.MobilePhone)
		argID++
	}

	if params.HomePhone != "" {
		setValues = append(setValues, fmt.Sprintf("home_phone=$%d", argID))
		args = append(args, params.HomePhone)
		argID++
	}

	if len(setValues) == 0 {
		err := r.db.Select(&records, query)
		return records, err
	}

	setQuery := strings.Join(setValues, " AND ")

	query = query + " WHERE " + setQuery

	err := r.db.Select(&records, query, args...)

	return records, err
}

func (r *RecordPostgres) Update(recordUid string, record models.RecordInput) error {
	var setValues []string
	var args []interface{}
	argId := 1

	if record.FirstName != "" {
		setValues = append(setValues, fmt.Sprintf("first_name=$%d", argId))
		args = append(args, record.FirstName)
		argId++
	}

	if record.LastName != "" {
		setValues = append(setValues, fmt.Sprintf("last_name=$%d", argId))
		args = append(args, record.LastName)
		argId++
	}

	if record.MobilePhone != "" {
		setValues = append(setValues, fmt.Sprintf("mobile_phone=$%d", argId))
		args = append(args, record.MobilePhone)
		argId++
	}

	if record.HomePhone != "" {
		setValues = append(setValues, fmt.Sprintf("home_phone=$%d", argId))
		args = append(args, record.HomePhone)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE uuid=$%d", recordsTable, setQuery, argId)

	args = append(args, recordUid)

	_, err := r.db.Exec(query, args...)
	return err
}

func (r *RecordPostgres) Delete(recordUid string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE uuid=$1", recordsTable)
	_, err := r.db.Exec(query, recordUid)

	return err
}
