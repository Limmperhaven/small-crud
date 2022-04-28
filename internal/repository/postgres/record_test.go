package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"gitlab.digital-spirit.ru/study/artem_crud/internal/models"
	"log"
	"testing"
)

const (
	user     = "postgres"
	password = "secret"
	dbname   = "postgres"
	port     = "5433"
	dialect  = "postgres"
)

var (
	db = new(sqlx.DB)
)

func TestRepository(t *testing.T) {
	//setup

	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err.Error())
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "latest",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbname,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
	}

	resource, err := pool.RunWithOptions(&opts)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err.Error())
	}

	if err = pool.Retry(func() error {
		_, err := NewPostgresDB(Config{
			Host:     "localhost",
			Port:     port,
			Username: user,
			Password: password,
			DBName:   dbname,
			SSLMode:  "disable",
		})
		return err
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err.Error())
	}

	db, _ = NewPostgresDB(Config{
		Host:     "localhost",
		Port:     port,
		Username: user,
		Password: password,
		DBName:   dbname,
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("error opening database: %s", err.Error())
	}

	upQuery := "CREATE TABLE IF NOT EXISTS records\n(\n    uuid         varchar(64)  not null unique primary key,\n    first_name   varchar(255) not null,\n    last_name    varchar(255) not null,\n    mobile_phone varchar(255) not null,\n    home_phone   varchar(255)\n);"
	_, err = db.Exec(upQuery)
	if err != nil {
		log.Fatalf("error migration up: %s", err.Error())
	}

	t.Cleanup(func() {
		if err := pool.Purge(resource); err != nil {
			log.Fatalf("Could not purge resource: %s", err)
		}
	})

	recUuid := ""
	repo := NewRecordPostgres(db)

	t.Run("CreationTests", func(t *testing.T) {

		t.Log("Creation tests started")
		t.Run("Create", func(t *testing.T) {
			uuid, err := repo.Create(models.Record{
				FirstName:   "Test",
				LastName:    "Test",
				MobilePhone: "Test",
				HomePhone:   "Test",
			})

			assert.NotEmpty(t, uuid)
			assert.Empty(t, err)
			recUuid = uuid
		})
	})

	t.Run("GetByUuidTests", func(t *testing.T) {
		t.Log("Get By UUID Tests Started")
		t.Run("GetByUuid", func(t *testing.T) {
			rec, err := repo.GetByUid(recUuid)

			assert.NoError(t, err)
			assert.NotEmpty(t, rec)
			assert.Equal(t, rec, models.Record{
				Uuid:        recUuid,
				FirstName:   "Test",
				LastName:    "Test",
				MobilePhone: "Test",
				HomePhone:   "Test",
			})
		})
		t.Run("GetByUuidNegative", func(t *testing.T) {
			_, err := repo.GetByUid("Test")

			assert.Error(t, err)
		})
	})

	t.Run("GetByFilterTests", func(t *testing.T) {
		t.Log("Get By Filter Tests Started")
		t.Run("Without params", func(t *testing.T) {
			recs, err := repo.GetByFilter(*new(models.RecordInput))

			assert.NoError(t, err)
			assert.Equal(t, recs, []models.Record{{recUuid, "Test", "Test",
				"Test", "Test"}})
		})
		t.Run("With right param", func(t *testing.T) {
			recs, err := repo.GetByFilter(models.RecordInput{FirstName: "Test"})

			assert.NoError(t, err)
			assert.Equal(t, recs, []models.Record{{recUuid, "Test", "Test",
				"Test", "Test"}})
		})
		t.Run("With invalid param", func(t *testing.T) {
			recs, err := repo.GetByFilter(models.RecordInput{FirstName: "Test1"})

			assert.NoError(t, err)
			assert.Empty(t, recs)
		})
	})

	t.Run("Update Tests", func(t *testing.T) {
		t.Log("Update Tests Started")
		t.Run("Positive", func(t *testing.T) {
			err := repo.Update(recUuid, models.RecordInput{FirstName: "Test2"})

			assert.NoError(t, err)

			rec, err := repo.GetByUid(recUuid)

			assert.NoError(t, err)
			assert.Equal(t, rec, models.Record{Uuid: recUuid, FirstName: "Test2", LastName: "Test",
				MobilePhone: "Test", HomePhone: "Test"})
		})
		t.Run("Negative", func(t *testing.T) {
			err := repo.Update("Test", models.RecordInput{FirstName: "Test2"})

			assert.Error(t, err)
		})
	})

	t.Run("Delete Tests", func(t *testing.T) {
		t.Log("Delete Tests Started")
		t.Run("Positive", func(t *testing.T) {
			err := repo.Delete(recUuid)

			assert.NoError(t, err)

			_, err = repo.GetByUid(recUuid)

			assert.Error(t, err)
		})
		t.Run("Negative", func(t *testing.T) {
			err := repo.Delete("Test")

			assert.Error(t, err)
		})
	})

}
