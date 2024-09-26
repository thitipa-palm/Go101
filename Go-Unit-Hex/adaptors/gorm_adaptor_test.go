package adaptors

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/thitipa-palm/go-Unit-Hex/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGormOrder(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database", err)
	}

	repo := NewGormRepository(gormDB)
	t.Run("Place order successfully", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery("^INSERT INTO \"orders\" (.+)$").
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		err := repo.Save(core.Order{Total: 201})
		assert.NoError(t, err)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

	t.Run("fail to add order", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery("^INSERT INTO \"orders\" (.+)$").
			WillReturnError(errors.New("database error"))
		mock.ExpectRollback()

		err := repo.Save(core.Order{Total: 201})
		assert.Error(t, err)

		assert.NoError(t, mock.ExpectationsWereMet())
	})

}
