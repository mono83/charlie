package mysql

import (
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/mono83/charlie/model"
	"github.com/stretchr/testify/assert"
)

func TestFetchIssueRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	expected := []model.Issue{
		{
			ID:   1, ReleaseID: 2, IssueID: "issue-1",
			Type: model.Added, Components: []string{"backend", "frontend"}, Message: "First Issue",
		},
	}

	// In fact column names does not affect anything (specified for clarity)
	rows := sqlmock.NewRows([]string{"id", "release_id", "issue_id", "type", "components", "message"}).
		AddRow(1, 2, "issue-1", "added", "backend,frontend", "First Issue")

	// Query in fact does not affect anything
	query := "SELECT id, release_id, issue_id, type, components, message FROM `issue` WHERE id = ?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	repo := NewMysqlIssueRepository(db)

	actual, err := repo.fetch(query, 1)
	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, expected[0], *actual[0])
}
