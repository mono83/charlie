package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mono83/charlie/db"
	"github.com/mono83/charlie/model"
	"log"
	"time"
)

type mysqlReleaseRepository struct {
	Conn      *sql.DB
	IssueRepo db.IssueRepository
}

func NewMysqlReleaseRepository(Conn *sql.DB) *mysqlReleaseRepository { // returning implementation instead of interface allows to test private methods
	return &mysqlReleaseRepository{Conn: Conn, IssueRepo: NewMysqlIssueRepository(Conn)}
}

func NewMysqlReleaseRepositoryWithIssueRepo(Conn *sql.DB, IssueRepo db.IssueRepository) *mysqlReleaseRepository { // returning implementation instead of interface allows to test private methods
	return &mysqlReleaseRepository{Conn: Conn, IssueRepo: IssueRepo}
}

func (releaseRepo *mysqlReleaseRepository) fetch(query string, args ...interface{}) ([]*model.Release, error) {

	rows, err := releaseRepo.Conn.Query(query, args...)
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
		issues, err := releaseRepo.IssueRepo.GetByReleaseId(r.ID)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		r.Issues = issues
		releases = append(releases, r)
	}

	return releases, nil
}

func (r *mysqlReleaseRepository) GetByID(id int64) (*model.Release, error) {
	releases, err := r.fetch("SELECT id, project_id, major, minor, patch, label, build, date FROM `release` WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return r.getSingleRelease(releases)
}

func (r *mysqlReleaseRepository) GetByProjectID(projectId int64) ([]*model.Release, error) {
	releases, err := r.fetch("SELECT id, project_id, major, minor, patch, label, build, date FROM `release` WHERE project_id = ?", projectId)
	if err != nil {
		return nil, err
	}

	return releases, nil
}

func (r *mysqlReleaseRepository) getSingleRelease(releases []*model.Release) (*model.Release, error) {
	if len(releases) > 1 {
		return nil, errors.New(fmt.Sprintf("Expected not more than 1 release, but was %d", len(releases)))
	}
	if len(releases) == 0 {
		return nil, nil
	}

	return releases[0], nil
}
