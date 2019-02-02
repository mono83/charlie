package mysql

import (
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/mono83/charlie/model"
	"github.com/stretchr/testify/assert"
)

func TestFetchProjectRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expected := []model.Project{
		{
			ID: 1, Name: "project-1", Description: "Charlie First Project",
		},
	}

	// In fact column names does not affect anything (specified for clarity)
	rows := sqlmock.NewRows([]string{"id", "name", "description"}).AddRow(1, "project-1", "Charlie First Project")

	// Query in fact does not affect anything
	query := "SELECT id, name, description FROM `project` WHERE id = ?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	repo := NewMysqlProjectRepository(db)

	actual, err := repo.fetch(query, 1)
	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, expected[0], *actual[0])
}
