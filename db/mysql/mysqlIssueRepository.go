package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mono83/charlie/model"
	"log"
	"strings"
	"strconv"
)

type mysqlIssueRepository struct {
	Conn *sql.DB
}

func NewMysqlIssueRepository(Conn *sql.DB) *mysqlIssueRepository { // returning implementation instead of interface allows to test private methods
	return &mysqlIssueRepository{Conn: Conn}
}

func (repo *mysqlIssueRepository) fetch(query string, args ...interface{}) ([]*model.Issue, error) {

	rows, err := repo.Conn.Query(query, args...)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	issues := make([]*model.Issue, 0)
	for rows.Next() {
		i := new(model.Issue)
		var componentsString sql.NullString
		var message sql.NullString
		var typeString sql.NullString
		err := rows.Scan(&i.ID, &i.ReleaseID, &i.IssueID, &typeString, &componentsString, &message)
		i.Type = model.ParseIssueType(typeString.String)
		i.Message = message.String
		if len(componentsString.String) > 0 {
			i.Components = strings.Split(componentsString.String, ",")
		}
		issues = append(issues, i)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
	}

	return issues, nil
}

func (repo *mysqlIssueRepository) GetByID(id int64) (*model.Issue, error) {
	issues, err := repo.fetch("SELECT id, release_id, issue_id, type, components, message FROM `issue` WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	return repo.getSingleResult(issues)
}

func (repo *mysqlIssueRepository) GetByReleaseId(releaseId int64) ([]*model.Issue, error) {
	issues, err := repo.fetch("SELECT id, release_id, issue_id, type, components, message FROM `issue` WHERE release_id = ?", releaseId)
	if err != nil {
		return nil, err
	}

	return issues, nil
}

func (repo *mysqlIssueRepository) getSingleResult(issues []*model.Issue) (*model.Issue, error) {
	if len(issues) > 1 {
		return nil, errors.New(fmt.Sprintf("Expected not more than 1 issue, but was %d", len(issues)))
	}
	if len(issues) == 0 {
		return nil, nil
	}

	return issues[0], nil
}

func (repo *mysqlIssueRepository) Store(issues []*model.Issue) (error) {

	insertQuery := "INSERT INTO `issue` (`release_id`, `issue_id`, `type`, `components`, `message`) VALUES "

	placeHolders := make([]string, 0)
	args := make([]interface{}, 0)
	for i := 0; i < len(issues); i++ {
		// Placeholders
		placeHolders = append(placeHolders, "(?, ?, ?, ?, ?)")

		// Arguments
		args = append(args, strconv.FormatInt(issues[i].ReleaseID, 10))
		args = append(args, issues[i].IssueID)
		args = append(args, issues[i].Type.String())
		args = append(args, strings.Join(issues[i].Components, ","))
		args = append(args, issues[i].Message)
	}

	query := insertQuery + strings.Join(placeHolders, ",")

	_, err := repo.Conn.Exec(query, args...)

	return err
}