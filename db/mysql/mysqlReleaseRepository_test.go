package mysql

import (
	"testing"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/mono83/charlie/model"
	"github.com/stretchr/testify/assert"
	"time"
	"github.com/mono83/charlie/db/mocks"
	"github.com/golang/mock/gomock"
)

func TestFetchReleaseRepository(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	var unixSeconds int64 = 122341223
	version := model.Version{Major: "3", Minor: "4", Patch: "5", Label: "6", Build: "7"}
	issues := []*model.Issue{
		{
			ID:   1, ReleaseID: 2, IssueID: "issue-1",
			Type: model.Added, Components: []string{"backend", "frontend"}, Message: "First Issue",
		},
	}
	expected := []model.Release{
		{
			ID: 1, ProjectID: 2, Version: version, Issues: issues, Date: time.Unix(unixSeconds, 0),
		},
	}

	// In fact column names does not affect anything (specified for clarity)
	rows := sqlmock.NewRows([]string{"id", "project_id", "major", "minor", "patch", "label", "build", "date"}).
		AddRow("1", "2", "3", "4", "5", "6", "7", unixSeconds)

	// Query in fact does not affect anything
	query := "SELECT id, project_id, major, minor, patch, label, build, date FROM `release` WHERE id = ?"

	mock.ExpectQuery(query).WillReturnRows(rows)

	// Mock for IssueRepository
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	issueRepoMock := mocks.NewMockIssueRepository(ctrl)

	issueRepoMock.
		EXPECT().
		GetByReleaseId(expected[0].ID).
		Return(issues, nil)

	repo := NewMysqlReleaseRepositoryWithIssueRepo(db, issueRepoMock)

	actual, err := repo.fetch(query, 1)
	assert.Equal(t, len(expected), len(actual))
	assert.Equal(t, expected[0], *actual[0])
}
