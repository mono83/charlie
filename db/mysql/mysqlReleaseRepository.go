package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mono83/charlie/db"
	"github.com/mono83/charlie/model"
	"log"
	"time"
	"strconv"
)

type mysqlReleaseRepository struct {
	Conn        *sql.DB
	ProjectRepo db.ProjectRepository
	IssueRepo   db.IssueRepository
}

func NewMysqlReleaseRepository(Conn *sql.DB) *mysqlReleaseRepository { // returning implementation instead of interface allows to test private methods
	return &mysqlReleaseRepository{Conn: Conn, ProjectRepo: NewMysqlProjectRepository(Conn), IssueRepo: NewMysqlIssueRepository(Conn)}
}

func NewMysqlReleaseRepositoryWithIssueRepo(Conn *sql.DB, IssueRepo db.IssueRepository) *mysqlReleaseRepository { // returning implementation instead of interface allows to test private methods
	return &mysqlReleaseRepository{Conn: Conn, ProjectRepo: NewMysqlProjectRepository(Conn), IssueRepo: IssueRepo}
}

func (repo *mysqlReleaseRepository) fetch(query string, args ...interface{}) ([]*model.Release, error) {

	rows, err := repo.Conn.Query(query, args...)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	releases := make([]*model.Release, 0)
	for rows.Next() {
		r := new(model.Release)
		var unixSeconds int64
		var major sql.NullString
		var minor sql.NullString
		var patch sql.NullString
		var label sql.NullString
		var build sql.NullString

		err := rows.Scan(&r.ID, &r.ProjectID, &major, &minor, &patch, &label, &build, &unixSeconds)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		r.Version.Major = major.String
		r.Version.Minor = minor.String
		r.Version.Patch = patch.String
		r.Version.Label = label.String
		r.Version.Build = build.String
		r.Date = time.Unix(unixSeconds, 0)
		issues, err := repo.IssueRepo.GetByReleaseId(r.ID)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		r.Issues = issues
		releases = append(releases, r)
	}

	return releases, nil
}

func (repo *mysqlReleaseRepository) GetByID(id int64) (*model.Release, error) {
	releases, err := repo.fetch("SELECT id, project_id, major, minor, patch, label, build, date FROM `release` WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return repo.getSingleRelease(releases)
}

func (repo *mysqlReleaseRepository) GetByProjectID(projectId int64) ([]*model.Release, error) {
	releases, err := repo.fetch("SELECT id, project_id, major, minor, patch, label, build, date FROM `release` WHERE project_id = ?", projectId)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

func (repo *mysqlReleaseRepository) getSingleRelease(releases []*model.Release) (*model.Release, error) {
	if len(releases) > 1 {
		return nil, errors.New(fmt.Sprintf("Expected not more than 1 release, but was %d", len(releases)))
	}
	if len(releases) == 0 {
		return nil, nil
	}

	return releases[0], nil
}

func (repo *mysqlReleaseRepository) Store(release model.Release) (int64, error) {

	// Making sure ProjectId is set correctly
	project, err := repo.ProjectRepo.GetByID(release.ProjectID)
	if err != nil || project == nil {
		return 0, errors.New("Project not found by Id " + strconv.FormatInt(release.ProjectID, 10))
	}

	// Inserting Release into DB and fetching its assigned ID
	query := "INSERT INTO `release` (`project_id`, `major`, `minor`, `patch`, `label`, `build`, `date`) VALUES (?, ?, ?, ?, ?, ?, ?)"
	result, err := repo.Conn.Exec(query, release.ProjectID, release.Major, release.Minor, release.Patch, release.Label, release.Build, release.Date.Unix())
	if err != nil {
		return 0, err
	}

	releaseId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Setting releaseID to issues and storing them
	for _, issue := range release.Issues {
		issue.ReleaseID = releaseId
	}
	err = repo.IssueRepo.Store(release.Issues)

	return releaseId, err
}