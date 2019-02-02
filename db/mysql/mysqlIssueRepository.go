package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mono83/charlie/model"
	"log"
	"strings"
)

type mysqlIssueRepository struct {
	Conn *sql.DB
}

func NewMysqlIssueRepository(Conn *sql.DB) *mysqlIssueRepository { // returning implementation instead of interface allows to test private methods
	return &mysqlIssueRepository{Conn: Conn}
}

func (r *mysqlIssueRepository) fetch(query string, args ...interface{}) ([]*model.Issue, error) {

	rows, err := r.Conn.Query(query, args...)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	issues := make([]*model.Issue, 0)
	for rows.Next() {
		i := new(model.Issue)
		var componentsString string
		var typeString string
		err := rows.Scan(&i.ID, &i.ReleaseID, &i.IssueID, &typeString, &componentsString, &i.Message)
		i.Type = model.ParseIssueType(typeString)
		i.Components = strings.Split(componentsString, ",")
		issues = append(issues, i)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	return issues, nil
}

func (r *mysqlIssueRepository) GetByID(id int64) (*model.Issue, error) {
	issues, err := r.fetch("SELECT id, release_id, issue_id, type, components, message FROM `issue` WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return r.getSingleResult(issues)
}

func (r *mysqlIssueRepository) GetByReleaseId(releaseId int64) ([]*model.Issue, error) {
	issues, err := r.fetch("SELECT id, release_id, issue_id, type, components, message FROM `issue` WHERE release_id = ?", releaseId)
	if err != nil {
		return nil, err
	}

	return issues, nil
}

func (r *mysqlIssueRepository) getSingleResult(issues []*model.Issue) (*model.Issue, error) {
	if len(issues) > 1 {
		return nil, errors.New(fmt.Sprintf("Expected not more than 1 issue, but was %d", len(issues)))
	}
	if len(issues) == 0 {
		return nil, nil
	}

	return issues[0], nil
}
