package canvas_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"sketch/internal/canvas"
	"sketch/tests/faker"
	"testing"
)

func TestRepository_GetByID(t *testing.T) {
	const query = "select id, drawing, created_at from drawings where id = $1"
	setup := func() (canvas.Repository, sqlmock.Sqlmock) {
		mockDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		db := sqlx.NewDb(mockDB, "sqlmock")
		return canvas.NewRepository(db), mock
	}

	t.Run("when there is a result, should return it", func(t *testing.T) {
		repository, mock := setup()
		fakeDraw := canvas.NewCanvas(":)")
		rows := sqlmock.
			NewRows([]string{"id", "drawing", "created_at"}).
			AddRow(fakeDraw.ID, fakeDraw.Drawing, fakeDraw.CreatedAt)

		mock.ExpectQuery(query).
			WithArgs(fakeDraw.ID).
			WillReturnRows(rows)

		result, err := repository.GetByID(context.Background(), fakeDraw.ID)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
		assert.Nil(t, err)
		assert.EqualValues(t, fakeDraw, result)
	})

	t.Run("when there are no results, should return not found error", func(t *testing.T) {
		repository, mock := setup()

		mock.ExpectQuery(query).WithArgs("123").WillReturnError(sql.ErrNoRows)
		result, err := repository.GetByID(context.Background(), "123")

		assert.Empty(t, result)
		assert.ErrorIs(t, err, canvas.ErrNotFound)
	})

	t.Run("when there is an error querying the result, should return it", func(t *testing.T) {
		repository, mock := setup()

		mock.ExpectQuery(query).WithArgs("123").WillReturnError(faker.NewError())
		result, err := repository.GetByID(context.Background(), "123")

		assert.Empty(t, result)
		assert.ErrorIs(t, err, faker.NewError())
	})
}

func TestRepository_Save(t *testing.T) {
	const query = "insert into drawings (id, drawing, created_at) values (?, ?, ?)"
	setup := func() (canvas.Repository, sqlmock.Sqlmock) {
		mockDB, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		db := sqlx.NewDb(mockDB, "sqlmock")
		return canvas.NewRepository(db), mock
	}
	t.Run("when there is no error saving the drawing", func(t *testing.T) {
		repository, mock := setup()

		fakeCanvas := faker.NewCanvas(t)
		mock.ExpectExec(query).
			WithArgs(fakeCanvas.ID, fakeCanvas.Drawing, fakeCanvas.CreatedAt).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err := repository.Save(context.Background(), fakeCanvas)

		assert.NoError(t, err)
		if err := mock.ExpectationsWereMet(); err != nil {
			assert.Fail(t, err.Error())
		}
	})

	t.Run("when there is an error saving the drawing, should return it", func(t *testing.T) {
		repository, mock := setup()

		fakeCanvas := faker.NewCanvas(t)
		mock.ExpectExec(query).
			WithArgs(fakeCanvas.ID, fakeCanvas.Drawing, fakeCanvas.CreatedAt).
			WillReturnError(faker.NewError())

		err := repository.Save(context.Background(), fakeCanvas)

		assert.ErrorIs(t, err, faker.NewError())
		if err := mock.ExpectationsWereMet(); err != nil {
			assert.Fail(t, err.Error())
		}
	})
}
